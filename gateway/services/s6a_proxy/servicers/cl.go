/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

// package service implements S6a GRPC proxy service which sends AIR, ULR messages over diameter connection,
// waits (blocks) for diameter's AIAs, ULAs & returns their RPC representation
// It also handles CLR, sends sync rpc request to gateway, then returns a CLA over diameter connection.
package servicers

import (
	"magma/feg/cloud/go/protos"
	"magma/feg/gateway/services/s6a_proxy"

	"github.com/fiorix/go-diameter/diam"
	"github.com/fiorix/go-diameter/diam/avp"
	"github.com/fiorix/go-diameter/diam/datatype"
	"github.com/golang/glog"
)

const (
	MaxSyncRPCRetries = 3
	MaxDiamClRetries  = 3
)

// S6a CLR
func handleCLR(s *s6aProxy) diam.HandlerFunc {
	return func(c diam.Conn, m *diam.Message) {
		glog.V(2).Infof("handling CLR\n")
		var code uint32 //result-code
		var clr CLR
		err := m.Unmarshal(&clr)
		if err != nil {
			glog.Errorf("CLR Unmarshal failed for remote %s & message %s: %s", c.RemoteAddr(), m, err)
			return
		}
		var retries = MaxSyncRPCRetries
		for ; retries >= 0; retries-- {
			code, err = forwardCLRToGateway(&clr)
			if err != nil {
				glog.Errorf("Failed to forward CLR to gateway. err: %v. Retries left: %v\n", err, retries)
			} else {
				break
			}
		}
		err = s.sendCLA(c, m, code, &clr, MaxDiamClRetries)
		if err != nil {
			glog.Errorf("Failed to send CLA: %s", err.Error())
		} else {
			glog.V(2).Infof("Successfully sent CLA\n")
		}
	}
}

func forwardCLRToGateway(clr *CLR) (uint32, error) {
	cancelLocationType := protos.CancelLocationRequest_CancellationType(clr.CancellationType)
	in := &protos.CancelLocationRequest{UserName: clr.UserName, CancellationType: cancelLocationType}
	res, err := s6a_proxy.GWS6AProxyCancelLocation(in)
	if err != nil {
		if res != nil && res.ErrorCode == protos.ErrorCode_USER_UNKNOWN {
			return diam.Success, err
		}
		return diam.UnableToDeliver, err
	}
	if res.ErrorCode != 0 {
		//todo: once gateway side is implemented, check how errCode is populated, and how
		// to translate them into go-diameter result-code.
		return diam.UnableToComply, nil
	}
	return diam.Success, nil
}

func (s *s6aProxy) sendCLA(c diam.Conn, m *diam.Message, code uint32, clr *CLR, retries uint) error {
	ans := m.Answer(code)
	// SessionID is required to be the AVP in position 1
	ans.InsertAVP(diam.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String(clr.SessionID)))
	ans.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(clr.AuthSessionState))
	s.addDiamOriginAVPs(m)

	_, err := ans.WriteToWithRetry(c, retries)
	return err

}
