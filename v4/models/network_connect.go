// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkConnect NetworkConnect represents the data to be used to connect a container to the network
//
// swagger:model NetworkConnect
type NetworkConnect struct {

	// container
	Container string `json:"Container,omitempty"`

	// endpoint config
	EndpointConfig *EndpointSettings `json:"EndpointConfig,omitempty"`
}

// Validate validates this network connect
func (m *NetworkConnect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkConnect) validateEndpointConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointConfig) { // not required
		return nil
	}

	if m.EndpointConfig != nil {
		if err := m.EndpointConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndpointConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network connect based on the context it is used
func (m *NetworkConnect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpointConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkConnect) contextValidateEndpointConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EndpointConfig != nil {
		if err := m.EndpointConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndpointConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkConnect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkConnect) UnmarshalBinary(b []byte) error {
	var res NetworkConnect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
