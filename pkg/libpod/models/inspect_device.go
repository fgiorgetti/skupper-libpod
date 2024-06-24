// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectDevice InspectDevice is a single device that will be mounted into the container.
//
// swagger:model InspectDevice
type InspectDevice struct {

	// CgroupPermissions is the permissions of the mounted device.
	// Presently not populated.
	// TODO.
	CgroupPermissions string `json:"CgroupPermissions,omitempty"`

	// PathInContainer is the path of the device within the container.
	PathInContainer string `json:"PathInContainer,omitempty"`

	// PathOnHost is the path of the device on the host.
	PathOnHost string `json:"PathOnHost,omitempty"`
}

// Validate validates this inspect device
func (m *InspectDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect device based on context it is used
func (m *InspectDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InspectDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectDevice) UnmarshalBinary(b []byte) error {
	var res InspectDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}