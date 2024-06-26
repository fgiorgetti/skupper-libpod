// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ThrottleDevice ThrottleDevice is a structure that holds device:rate_per_second pair
//
// swagger:model ThrottleDevice
type ThrottleDevice struct {

	// path
	Path string `json:"Path,omitempty"`

	// rate
	Rate uint64 `json:"Rate,omitempty"`
}

// Validate validates this throttle device
func (m *ThrottleDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this throttle device based on context it is used
func (m *ThrottleDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThrottleDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThrottleDevice) UnmarshalBinary(b []byte) error {
	var res ThrottleDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
