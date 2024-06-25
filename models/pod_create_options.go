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

// PodCreateOptions PodCreateOptions provides all possible options for creating a pod and its infra container.
//
// The JSON tags below are made to match the respective field in ContainerCreateOptions for the purpose of mapping.
//
// swagger:model PodCreateOptions
type PodCreateOptions struct {

	// cgroup parent
	CgroupParent string `json:"cgroup_parent,omitempty"`

	// cpus
	Cpus float64 `json:"cpus,omitempty"`

	// cpuset cpus
	CpusetCpus string `json:"cpuset_cpus,omitempty"`

	// create command
	CreateCommand []string `json:"create_command"`

	// device read b ps
	DeviceReadBPs []string `json:"device_read_bps"`

	// devices
	Devices []string `json:"devices"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// infra
	Infra bool `json:"infra,omitempty"`

	// infra command
	InfraCommand string `json:"container_command,omitempty"`

	// infra conmon pid file
	InfraConmonPidFile string `json:"container_conmon_pidfile,omitempty"`

	// infra image
	InfraImage string `json:"infra_image,omitempty"`

	// infra name
	InfraName string `json:"container_name,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// pid
	Pid string `json:"pid,omitempty"`

	// security opt
	SecurityOpt []string `json:"security_opt"`

	// share
	Share []string `json:"share"`

	// share parent
	ShareParent bool `json:"share_parent,omitempty"`

	// sysctl
	Sysctl []string `json:"sysctl"`

	// volume
	Volume []string `json:"volume"`

	// volumes from
	VolumesFrom []string `json:"volumes_from"`

	// net
	Net *NetOptions `json:"net,omitempty"`
}

// Validate validates this pod create options
func (m *PodCreateOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodCreateOptions) validateNet(formats strfmt.Registry) error {
	if swag.IsZero(m.Net) { // not required
		return nil
	}

	if m.Net != nil {
		if err := m.Net.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("net")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pod create options based on the context it is used
func (m *PodCreateOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodCreateOptions) contextValidateNet(ctx context.Context, formats strfmt.Registry) error {

	if m.Net != nil {
		if err := m.Net.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("net")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("net")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodCreateOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodCreateOptions) UnmarshalBinary(b []byte) error {
	var res PodCreateOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}