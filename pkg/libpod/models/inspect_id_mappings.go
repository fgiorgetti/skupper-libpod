// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectIDMappings inspect ID mappings
//
// swagger:model InspectIDMappings
type InspectIDMappings struct {

	// g ID map
	GIDMap []string `json:"GidMap"`

	// UID map
	UIDMap []string `json:"UidMap"`
}

// Validate validates this inspect ID mappings
func (m *InspectIDMappings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect ID mappings based on context it is used
func (m *InspectIDMappings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InspectIDMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectIDMappings) UnmarshalBinary(b []byte) error {
	var res InspectIDMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}