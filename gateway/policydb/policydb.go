/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package policydb

import (
	"fmt"

	"magma/feg/gateway/object_store"
	"magma/feg/gateway/registry"
	"magma/feg/gateway/streamer"
	"magma/lte/cloud/go/protos"

	"github.com/golang/glog"
)

// PolicyDBClient defines interactions with the stored policy rules
type PolicyDBClient interface {
	GetChargingKeysForRules(ruleIDs []string, ruleDefs []*protos.PolicyRule) ([]uint32, error)
	GetPolicyRuleByID(id string) (*protos.PolicyRule, error)
	GetRuleIDsForBaseNames(baseNames []string) []string
}

// RedisPolicyDBClient is a policy client that loads policies from Redis
type RedisPolicyDBClient struct {
	PolicyMap      object_store.ObjectMap
	BaseNameMap    object_store.ObjectMap
	StreamerClient streamer.Client
}

// NewRedisPolicyDBClient creates a new RedisPolicyDBClient
func NewRedisPolicyDBClient(reg registry.CloudRegistry) (*RedisPolicyDBClient, error) {
	redisClient, err := object_store.NewRedisClient()
	if err != nil {
		return nil, err
	}
	client := &RedisPolicyDBClient{
		PolicyMap: object_store.NewRedisMap(
			redisClient,
			"policydb:rules",
			getProtoSerializer(),
			getPolicyDeserializer(),
		),
		BaseNameMap: object_store.NewRedisMap(
			redisClient,
			"policydb:base_names",
			getProtoSerializer(),
			getBaseNameDeserializer(),
		),
		StreamerClient: streamer.NewStreamerClient(reg),
	}
	client.StreamerClient.AddListener(NewBaseNameStreamListener(client.BaseNameMap))
	client.StreamerClient.AddListener(NewPolicyDBStreamListener(client.PolicyMap))
	return client, nil
}

// GetPolicyRuleByID returns a policy from its ID from redis
func (client *RedisPolicyDBClient) GetPolicyRuleByID(id string) (*protos.PolicyRule, error) {
	policyRaw, err := client.PolicyMap.Get(id)
	if err != nil {
		return nil, err
	}
	policy, ok := policyRaw.(*protos.PolicyRule)
	if !ok {
		return nil, fmt.Errorf("Could not cast object to policy rule for id %s", id)
	}
	return policy, nil
}

// GetChargingKeysForRules retrieves the charging keys associated with the given
// rule names from redis.
func (client *RedisPolicyDBClient) GetChargingKeysForRules(
	ruleIDs []string,
	ruleDefs []*protos.PolicyRule,
) ([]uint32, error) {
	keys := []uint32{}
	for _, id := range ruleIDs {
		policy, err := client.GetPolicyRuleByID(id)
		if err != nil {
			glog.Errorf("Unable to get rating group for policy %s: %s", id, err)
			continue
		}
		if needsCharging(policy) {
			keys = append(keys, policy.RatingGroup)
		}
	}
	for _, policy := range ruleDefs {
		if needsCharging(policy) {
			keys = append(keys, policy.RatingGroup)
		}
	}
	return keys, nil
}

func needsCharging(rule *protos.PolicyRule) bool {
	return rule.TrackingType == protos.PolicyRule_ONLY_OCS || rule.TrackingType == protos.PolicyRule_OCS_AND_PCRF
}

// GetRuleIDsForBaseNames gets the policy rule ids for given charging rule base names.
// These base name mappings are stored into redis through the stream client
func (client *RedisPolicyDBClient) GetRuleIDsForBaseNames(baseNames []string) []string {
	policyIDs := []string{}
	for _, bn := range baseNames {
		setRaw, err := client.BaseNameMap.Get(bn)
		if err != nil {
			glog.Errorf("Failed to look up base name %s: %s", bn, err)
			continue
		}
		nameSet, ok := setRaw.(*protos.ChargingRuleNameSet)
		if !ok {
			glog.Errorf("Could not cast object to base name set for base name %s", bn)
			continue
		}
		policyIDs = append(policyIDs, nameSet.GetRuleNames()...)
	}
	return policyIDs
}
