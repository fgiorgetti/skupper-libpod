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

// ContainerCreateCreatedBody ContainerCreateCreatedBody OK response to ContainerCreate operation
//
// swagger:model ContainerCreateCreatedBody
type ContainerCreateCreatedBody struct {

	// The ID of the created container
	// Required: true
	ID *string `json:"Id"`

	// Warnings encountered when creating the container
	// Required: true
	Warnings []string `json:"Warnings"`
}

// Validate validates this container create created body
func (m *ContainerCreateCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerCreateCreatedBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ContainerCreateCreatedBody) validateWarnings(formats strfmt.Registry) error {

	if err := validate.Required("Warnings", "body", m.Warnings); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this container create created body based on context it is used
func (m *ContainerCreateCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res ContainerCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
