// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Propagation Propagation represents the propagation of a mount.
//
// swagger:model Propagation
type Propagation string

// Validate validates this propagation
func (m Propagation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this propagation based on context it is used
func (m Propagation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}