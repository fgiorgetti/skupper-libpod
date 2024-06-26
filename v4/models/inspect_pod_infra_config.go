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
	"github.com/go-openapi/validate"
)

// InspectPodInfraConfig InspectPodInfraConfig contains the configuration of the pod's infra
// container.
//
// swagger:model InspectPodInfraConfig
type InspectPodInfraConfig struct {

	// CPUPeriod contains the CPU period of the pod
	CPUPeriod uint64 `json:"cpu_period,omitempty"`

	// CPUQuota contains the CPU quota of the pod
	CPUQuota int64 `json:"cpu_quota,omitempty"`

	// CPUSetCPUs contains linux specific CPU data for the container
	CPUSetCPUs string `json:"cpuset_cpus,omitempty"`

	// DNSOption is a set of DNS options that will be used by the infra
	// container's resolv.conf and shared with the remainder of the pod.
	DNSOption []string `json:"DNSOption"`

	// DNSSearch is a set of DNS search domains that will be used by the
	// infra container's resolv.conf and shared with the remainder of the
	// pod.
	DNSSearch []string `json:"DNSSearch"`

	// DNSServer is a set of DNS Servers that will be used by the infra
	// container's resolv.conf and shared with the remainder of the pod.
	DNSServer []string `json:"DNSServer"`

	// HostAdd adds a number of hosts to the infra container's resolv.conf
	// which will be shared with the rest of the pod.
	HostAdd []string `json:"HostAdd"`

	// HostNetwork is whether the infra container (and thus the whole pod)
	// will use the host's network and not create a network namespace.
	HostNetwork bool `json:"HostNetwork,omitempty"`

	// NetworkOptions are additional options for each network
	NetworkOptions map[string][]string `json:"NetworkOptions,omitempty"`

	// Networks is a list of CNI networks the pod will join.
	Networks []string `json:"Networks"`

	// NoManageHosts indicates that the pod will not manage /etc/hosts and
	// instead each container will handle their own.
	NoManageHosts bool `json:"NoManageHosts,omitempty"`

	// NoManageResolvConf indicates that the pod will not manage resolv.conf
	// and instead each container will handle their own.
	NoManageResolvConf bool `json:"NoManageResolvConf,omitempty"`

	// Pid is the PID namespace mode of the pod's infra container
	PidNS string `json:"pid_ns,omitempty"`

	// PortBindings are ports that will be forwarded to the infra container
	// and then shared with the pod.
	PortBindings map[string][]InspectHostPort `json:"PortBindings,omitempty"`

	// static IP
	StaticIP IP `json:"StaticIP,omitempty"`

	// StaticMAC is a static MAC address that will be assigned to the infra
	// container and then used by the pod.
	StaticMAC string `json:"StaticMAC,omitempty"`

	// UserNS is the usernamespace that all the containers in the pod will join.
	UserNS string `json:"userns,omitempty"`
}

// Validate validates this inspect pod infra config
func (m *InspectPodInfraConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePortBindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectPodInfraConfig) validatePortBindings(formats strfmt.Registry) error {
	if swag.IsZero(m.PortBindings) { // not required
		return nil
	}

	for k := range m.PortBindings {

		if err := validate.Required("PortBindings"+"."+k, "body", m.PortBindings[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.PortBindings[k]); i++ {

			if err := m.PortBindings[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PortBindings" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PortBindings" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *InspectPodInfraConfig) validateStaticIP(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticIP) { // not required
		return nil
	}

	if err := m.StaticIP.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StaticIP")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StaticIP")
		}
		return err
	}

	return nil
}

// ContextValidate validate this inspect pod infra config based on the context it is used
func (m *InspectPodInfraConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePortBindings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectPodInfraConfig) contextValidatePortBindings(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PortBindings {

		for i := 0; i < len(m.PortBindings[k]); i++ {

			if err := m.PortBindings[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PortBindings" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("PortBindings" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *InspectPodInfraConfig) contextValidateStaticIP(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticIP.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StaticIP")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StaticIP")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InspectPodInfraConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectPodInfraConfig) UnmarshalBinary(b []byte) error {
	var res InspectPodInfraConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
