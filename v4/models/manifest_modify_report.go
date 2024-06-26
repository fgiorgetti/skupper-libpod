// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ManifestModifyReport ManifestModifyReport provides the model for removed digests and changed manifest
//
// swagger:model ManifestModifyReport
type ManifestModifyReport struct {

	// Errors associated with operation
	Errors []string `json:"errors"`

	// Manifest List ID
	ID string `json:"Id,omitempty"`

	// Images to removed from manifest list, otherwise not provided.
	Images []string `json:"images"`
}

// Validate validates this manifest modify report
func (m *ManifestModifyReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manifest modify report based on context it is used
func (m *ManifestModifyReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManifestModifyReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManifestModifyReport) UnmarshalBinary(b []byte) error {
	var res ManifestModifyReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
