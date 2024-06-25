// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectHostPort InspectHostPort provides information on a port on the host that a container's
// port is bound to.
//
// swagger:model InspectHostPort
type InspectHostPort struct {

	// IP on the host we are bound to. "" if not specified (binding to all
	// IPs).
	HostIP string `json:"HostIp,omitempty"`

	// Port on the host we are bound to. No special formatting - just an
	// integer stuffed into a string.
	HostPort string `json:"HostPort,omitempty"`
}

// Validate validates this inspect host port
func (m *InspectHostPort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect host port based on context it is used
func (m *InspectHostPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InspectHostPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectHostPort) UnmarshalBinary(b []byte) error {
	var res InspectHostPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}