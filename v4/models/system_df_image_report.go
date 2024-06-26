// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SystemDfImageReport SystemDfImageReport describes an image for use with df
//
// swagger:model SystemDfImageReport
type SystemDfImageReport struct {

	// containers
	Containers int64 `json:"Containers,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// image ID
	ImageID string `json:"ImageID,omitempty"`

	// repository
	Repository string `json:"Repository,omitempty"`

	// shared size
	SharedSize int64 `json:"SharedSize,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// tag
	Tag string `json:"Tag,omitempty"`

	// unique size
	UniqueSize int64 `json:"UniqueSize,omitempty"`
}

// Validate validates this system df image report
func (m *SystemDfImageReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemDfImageReport) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this system df image report based on context it is used
func (m *SystemDfImageReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemDfImageReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDfImageReport) UnmarshalBinary(b []byte) error {
	var res SystemDfImageReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
