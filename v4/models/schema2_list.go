// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Schema2List Schema2List is a list of platform-specific manifests.
//
// swagger:model Schema2List
type Schema2List struct {

	// manifests
	Manifests []*Schema2ManifestDescriptor `json:"manifests"`

	// media type
	MediaType string `json:"mediaType,omitempty"`

	// schema version
	SchemaVersion int64 `json:"schemaVersion,omitempty"`
}

// Validate validates this schema2 list
func (m *Schema2List) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManifests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema2List) validateManifests(formats strfmt.Registry) error {
	if swag.IsZero(m.Manifests) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifests); i++ {
		if swag.IsZero(m.Manifests[i]) { // not required
			continue
		}

		if m.Manifests[i] != nil {
			if err := m.Manifests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this schema2 list based on the context it is used
func (m *Schema2List) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateManifests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema2List) contextValidateManifests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Manifests); i++ {

		if m.Manifests[i] != nil {
			if err := m.Manifests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("manifests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Schema2List) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schema2List) UnmarshalBinary(b []byte) error {
	var res Schema2List
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
