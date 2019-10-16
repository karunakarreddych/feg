// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GatewayFederationConfigs Federation configuration for a gateway
// swagger:model gateway_federation_configs
type GatewayFederationConfigs struct {

	// aaa server
	// Required: true
	AaaServer *AaaServer `json:"aaa_server"`

	// eap aka
	// Required: true
	EapAka *EapAka `json:"eap_aka"`

	// gx
	// Required: true
	Gx *Gx `json:"gx"`

	// gy
	// Required: true
	Gy *Gy `json:"gy"`

	// health
	// Required: true
	Health *Health `json:"health"`

	// hss
	// Required: true
	Hss *Hss `json:"hss"`

	// s6a
	// Required: true
	S6a *S6a `json:"s6a"`

	// served network ids
	// Required: true
	ServedNetworkIds ServedNetworkIds `json:"served_network_ids"`

	// swx
	// Required: true
	Swx *Swx `json:"swx"`
}

// Validate validates this gateway federation configs
func (m *GatewayFederationConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAaaServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEapAka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS6a(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServedNetworkIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayFederationConfigs) validateAaaServer(formats strfmt.Registry) error {

	if err := validate.Required("aaa_server", "body", m.AaaServer); err != nil {
		return err
	}

	if m.AaaServer != nil {
		if err := m.AaaServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aaa_server")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateEapAka(formats strfmt.Registry) error {

	if err := validate.Required("eap_aka", "body", m.EapAka); err != nil {
		return err
	}

	if m.EapAka != nil {
		if err := m.EapAka.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_aka")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateGx(formats strfmt.Registry) error {

	if err := validate.Required("gx", "body", m.Gx); err != nil {
		return err
	}

	if m.Gx != nil {
		if err := m.Gx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gx")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateGy(formats strfmt.Registry) error {

	if err := validate.Required("gy", "body", m.Gy); err != nil {
		return err
	}

	if m.Gy != nil {
		if err := m.Gy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gy")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateHealth(formats strfmt.Registry) error {

	if err := validate.Required("health", "body", m.Health); err != nil {
		return err
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateHss(formats strfmt.Registry) error {

	if err := validate.Required("hss", "body", m.Hss); err != nil {
		return err
	}

	if m.Hss != nil {
		if err := m.Hss.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hss")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateS6a(formats strfmt.Registry) error {

	if err := validate.Required("s6a", "body", m.S6a); err != nil {
		return err
	}

	if m.S6a != nil {
		if err := m.S6a.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s6a")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateServedNetworkIds(formats strfmt.Registry) error {

	if err := validate.Required("served_network_ids", "body", m.ServedNetworkIds); err != nil {
		return err
	}

	if err := m.ServedNetworkIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("served_network_ids")
		}
		return err
	}

	return nil
}

func (m *GatewayFederationConfigs) validateSwx(formats strfmt.Registry) error {

	if err := validate.Required("swx", "body", m.Swx); err != nil {
		return err
	}

	if m.Swx != nil {
		if err := m.Swx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayFederationConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayFederationConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayFederationConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
