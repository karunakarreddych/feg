/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package object_store

import (
	"fmt"

	"magma/feg/gateway/registry"

	"github.com/go-redis/redis"
)

// RedisClient defines an interface to interact with Redis. Only hash functions
// are used for now
type RedisClient interface {
	HSet(hash string, field string, value string) error
	HGet(hash string, field string) (string, error)
	HGetAll(hash string) (map[string]string, error)
	HDel(hash string, field string) error
}

// RedisClientImpl is the implementation of the redis client using an actual connection
// to redis using go-redis
type RedisClientImpl struct {
	rawClient *redis.Client
}

// NewRedisClient gets the redis configuration from the service config and returns
// a new client or an error if something went wrong
func NewRedisClient() (RedisClient, error) {
	address, err := registry.GetServiceAddress(registry.REDIS)
	if err != nil {
		return nil, err
	}
	return &RedisClientImpl{
		rawClient: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf(address),
		}),
	}, nil
}

// HSet sets a value at a hash,field pair
func (client *RedisClientImpl) HSet(hash string, field string, value string) error {
	return client.rawClient.HSet(hash, field, value).Err()
}

// HGet gets a value at a hash,field pair
func (client *RedisClientImpl) HGet(hash string, field string) (string, error) {
	return client.rawClient.HGet(hash, field).Result()
}

// HGetAll gets all the possible values for fields in a hash
func (client *RedisClientImpl) HGetAll(hash string) (map[string]string, error) {
	return client.rawClient.HGetAll(hash).Result()
}

// HDel deletes a value at a hash,field pair
func (client *RedisClientImpl) HDel(hash string, field string) error {
	return client.rawClient.HDel(hash, field).Err()
}
