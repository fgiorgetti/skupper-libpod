// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkingConfig NetworkingConfig represents the container's networking configuration for each of its interfaces
// Carries the networking configs specified in the `docker run` and `docker network connect` commands
//
// swagger:model NetworkingConfig
type NetworkingConfig struct {

	// endpoints config
	EndpointsConfig map[string]EndpointSettings `json:"EndpointsConfig,omitempty"`
}

// Validate validates this networking config
func (m *NetworkingConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkingConfig) validateEndpointsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointsConfig) { // not required
		return nil
	}

	for k := range m.EndpointsConfig {

		if err := validate.Required("EndpointsConfig"+"."+k, "body", m.EndpointsConfig[k]); err != nil {
			return err
		}
		if val, ok := m.EndpointsConfig[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndpointsConfig" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("EndpointsConfig" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this networking config based on the context it is used
func (m *NetworkingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpointsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkingConfig) contextValidateEndpointsConfig(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.EndpointsConfig {

		if val, ok := m.EndpointsConfig[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkingConfig) UnmarshalBinary(b []byte) error {
	var res NetworkingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
