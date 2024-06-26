// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectSecret InspectSecret contains information on secrets mounted inside the container
//
// swagger:model InspectSecret
type InspectSecret struct {

	// ID is the GID of the mounted secret file
	GID uint32 `json:"GID,omitempty"`

	// ID is the ID of the secret
	ID string `json:"ID,omitempty"`

	// ID is the ID of the mode of the mounted secret file
	Mode uint32 `json:"Mode,omitempty"`

	// Name is the name of the secret
	Name string `json:"Name,omitempty"`

	// ID is the UID of the mounted secret file
	UID uint32 `json:"UID,omitempty"`
}

// Validate validates this inspect secret
func (m *InspectSecret) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect secret based on context it is used
func (m *InspectSecret) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InspectSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectSecret) UnmarshalBinary(b []byte) error {
	var res InspectSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
