// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPAMConfig IPAMConfig represents IPAM configurations
//
// swagger:model IPAMConfig
type IPAMConfig struct {

	// aux address
	AuxAddress map[string]string `json:"AuxiliaryAddresses,omitempty"`

	// gateway
	Gateway string `json:"Gateway,omitempty"`

	// IP range
	IPRange string `json:"IPRange,omitempty"`

	// subnet
	Subnet string `json:"Subnet,omitempty"`
}

// Validate validates this IP a m config
func (m *IPAMConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP a m config based on context it is used
func (m *IPAMConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAMConfig) UnmarshalBinary(b []byte) error {
	var res IPAMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
