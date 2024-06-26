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

// HealthcheckResult HealthcheckResult stores information about a single run of a healthcheck probe
//
// swagger:model HealthcheckResult
type HealthcheckResult struct {

	// end
	// Format: date-time
	End strfmt.DateTime `json:"End,omitempty"`

	// exit code
	ExitCode int64 `json:"ExitCode,omitempty"`

	// output
	Output string `json:"Output,omitempty"`

	// start
	// Format: date-time
	Start strfmt.DateTime `json:"Start,omitempty"`
}

// Validate validates this healthcheck result
func (m *HealthcheckResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthcheckResult) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("End", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HealthcheckResult) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("Start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this healthcheck result based on context it is used
func (m *HealthcheckResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthcheckResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthcheckResult) UnmarshalBinary(b []byte) error {
	var res HealthcheckResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
