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

// AuthenticateOKBody AuthenticateOKBody authenticate o k body
//
// swagger:model AuthenticateOKBody
type AuthenticateOKBody struct {

	// An opaque token used to authenticate a user after a successful login
	// Required: true
	IdentityToken *string `json:"IdentityToken"`

	// The status of the authentication
	// Required: true
	Status *string `json:"Status"`
}

// Validate validates this authenticate o k body
func (m *AuthenticateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentityToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticateOKBody) validateIdentityToken(formats strfmt.Registry) error {

	if err := validate.Required("IdentityToken", "body", m.IdentityToken); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticateOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this authenticate o k body based on context it is used
func (m *AuthenticateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticateOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticateOKBody) UnmarshalBinary(b []byte) error {
	var res AuthenticateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
