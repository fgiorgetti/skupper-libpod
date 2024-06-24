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

// DockerVolumeCreate docker volume create
//
// swagger:model DockerVolumeCreate
type DockerVolumeCreate struct {
	VolumeCreateBody
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DockerVolumeCreate) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VolumeCreateBody
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VolumeCreateBody = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DockerVolumeCreate) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.VolumeCreateBody)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this docker volume create
func (m *DockerVolumeCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this docker volume create based on the context it is used
func (m *DockerVolumeCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VolumeCreateBody
	if err := m.VolumeCreateBody.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}