// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Plugins plugins
//
// swagger:model Plugins
type Plugins struct {

	// log
	Log []string `json:"log"`

	// network
	Network []string `json:"network"`

	// volume
	Volume []string `json:"volume"`
}

// Validate validates this plugins
func (m *Plugins) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plugins based on context it is used
func (m *Plugins) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Plugins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plugins) UnmarshalBinary(b []byte) error {
	var res Plugins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
