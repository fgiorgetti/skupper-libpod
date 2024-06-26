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

// LinuxPersonality LinuxPersonality represents the Linux personality syscall input
//
// swagger:model LinuxPersonality
type LinuxPersonality struct {

	// Additional flags
	Flags []LinuxPersonalityFlag `json:"flags"`

	// domain
	Domain LinuxPersonalityDomain `json:"domain,omitempty"`
}

// Validate validates this linux personality
func (m *LinuxPersonality) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxPersonality) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(m.Flags) { // not required
		return nil
	}

	for i := 0; i < len(m.Flags); i++ {

		if err := m.Flags[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flags" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LinuxPersonality) validateDomain(formats strfmt.Registry) error {
	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := m.Domain.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("domain")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("domain")
		}
		return err
	}

	return nil
}

// ContextValidate validate this linux personality based on the context it is used
func (m *LinuxPersonality) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxPersonality) contextValidateFlags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Flags); i++ {

		if err := m.Flags[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flags" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *LinuxPersonality) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Domain.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("domain")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("domain")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinuxPersonality) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxPersonality) UnmarshalBinary(b []byte) error {
	var res LinuxPersonality
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
