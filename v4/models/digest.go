// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Digest Digest allows simple protection of hex formatted digest strings, prefixed
// by their algorithm. Strings of type Digest have some guarantee of being in
// the correct format and it provides quick access to the components of a
// digest string.
//
// The following is an example of the contents of Digest types:
//
// sha256:7173b809ca12ec5dee4506cd86be934c4596dd234ee82c0662eac04a8c2c71dc
//
// This allows to abstract the digest behind this type and work only in those
// terms.
//
// swagger:model Digest
type Digest string

// Validate validates this digest
func (m Digest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this digest based on context it is used
func (m Digest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
