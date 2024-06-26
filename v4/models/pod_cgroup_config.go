// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodCgroupConfig PodCgroupConfig contains configuration options about a pod's cgroups.
//
// This will be expanded in future updates to pods.
//
// swagger:model PodCgroupConfig
type PodCgroupConfig struct {

	// CgroupParent is the parent for the Cgroup that the pod will create.
	// This pod cgroup will, in turn, be the default cgroup parent for all
	// containers in the pod.
	// Optional.
	CgroupParent string `json:"cgroup_parent,omitempty"`
}

// Validate validates this pod cgroup config
func (m *PodCgroupConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod cgroup config based on context it is used
func (m *PodCgroupConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodCgroupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodCgroupConfig) UnmarshalBinary(b []byte) error {
	var res PodCgroupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}