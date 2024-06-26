// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoreInfo StoreInfo describes the container storage and its
// attributes
//
// swagger:model StoreInfo
type StoreInfo struct {

	// config file
	ConfigFile string `json:"configFile,omitempty"`

	// graph driver name
	GraphDriverName string `json:"graphDriverName,omitempty"`

	// graph options
	GraphOptions map[string]interface{} `json:"graphOptions,omitempty"`

	// graph root
	GraphRoot string `json:"graphRoot,omitempty"`

	// graph status
	GraphStatus map[string]string `json:"graphStatus,omitempty"`

	// image copy tmp dir
	ImageCopyTmpDir string `json:"imageCopyTmpDir,omitempty"`

	// run root
	RunRoot string `json:"runRoot,omitempty"`

	// volume path
	VolumePath string `json:"volumePath,omitempty"`

	// container store
	ContainerStore *ContainerStore `json:"containerStore,omitempty"`

	// image store
	ImageStore *ImageStore `json:"imageStore,omitempty"`
}

// Validate validates this store info
func (m *StoreInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageStore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoreInfo) validateContainerStore(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerStore) { // not required
		return nil
	}

	if m.ContainerStore != nil {
		if err := m.ContainerStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerStore")
			}
			return err
		}
	}

	return nil
}

func (m *StoreInfo) validateImageStore(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageStore) { // not required
		return nil
	}

	if m.ImageStore != nil {
		if err := m.ImageStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageStore")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this store info based on the context it is used
func (m *StoreInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoreInfo) contextValidateContainerStore(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerStore != nil {
		if err := m.ContainerStore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerStore")
			}
			return err
		}
	}

	return nil
}

func (m *StoreInfo) contextValidateImageStore(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageStore != nil {
		if err := m.ImageStore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageStore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoreInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoreInfo) UnmarshalBinary(b []byte) error {
	var res StoreInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
