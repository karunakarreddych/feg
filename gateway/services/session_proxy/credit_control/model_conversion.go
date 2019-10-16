/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package credit_control

import (
	"magma/lte/cloud/go/protos"
)

func (gsu *GrantedServiceUnit) ToProto() *protos.GrantedUnits {
	if gsu == nil {
		return &protos.GrantedUnits{
			Total: &protos.CreditUnit{IsValid: false},
			Tx:    &protos.CreditUnit{IsValid: false},
			Rx:    &protos.CreditUnit{IsValid: false},
		}
	}
	return &protos.GrantedUnits{
		Total: getCreditUnit(gsu.TotalOctets),
		Tx:    getCreditUnit(gsu.InputOctets),  // Input == Tx == Uplink
		Rx:    getCreditUnit(gsu.OutputOctets), // Output == Rx == Downlink
	}
}

func getCreditUnit(volume *uint64) *protos.CreditUnit {
	if volume == nil {
		return &protos.CreditUnit{IsValid: false}
	}
	return &protos.CreditUnit{IsValid: true, Volume: *volume}
}
