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

// PodSpecGenerator PodSpecGenerator describes options to create a pod
//
// swagger:model PodSpecGenerator
type PodSpecGenerator struct {

	// CNINetworks is a list of CNI networks to join the container to.
	// If this list is empty, the default CNI network will be joined
	// instead. If at least one entry is present, we will not join the
	// default network (unless it is part of this list).
	// Only available if NetNS is set to bridge.
	// Optional.
	// Deprecated: as of podman 4.0 use "Networks" instead.
	CNINetworks []string `json:"cni_networks"`

	// CPU period of the cpuset, determined by --cpus
	CPUPeriod uint64 `json:"cpu_period,omitempty"`

	// CPU quota of the cpuset, determined by --cpus
	CPUQuota int64 `json:"cpu_quota,omitempty"`

	// CgroupParent is the parent for the Cgroup that the pod will create.
	// This pod cgroup will, in turn, be the default cgroup parent for all
	// containers in the pod.
	// Optional.
	CgroupParent string `json:"cgroup_parent,omitempty"`

	// DNSOption is a set of DNS options that will be used in the infra
	// container's resolv.conf, which will, by default, be shared with all
	// containers in the pod.
	// Conflicts with NoInfra=true.
	// Optional.
	DNSOption []string `json:"dns_option"`

	// DNSSearch is a set of DNS search domains that will be used in the
	// infra container's resolv.conf, which will, by default, be shared with
	// all containers in the pod.
	// If not provided, DNS search domains from the host's resolv.conf will
	// be used.
	// Conflicts with NoInfra=true.
	// Optional.
	DNSSearch []string `json:"dns_search"`

	// DNSServer is a set of DNS servers that will be used in the infra
	// container's resolv.conf, which will, by default, be shared with all
	// containers in the pod.
	// If not provided, the host's DNS servers will be used, unless the only
	// server set is a localhost address. As the container cannot connect to
	// the host's localhost, a default server will instead be set.
	// Conflicts with NoInfra=true.
	// Optional.
	DNSServer []IP `json:"dns_server"`

	// Devices contains user specified Devices to be added to the Pod
	Devices []string `json:"pod_devices"`

	// HostAdd is a set of hosts that will be added to the infra container's
	// etc/hosts that will, by default, be shared with all containers in
	// the pod.
	// Conflicts with NoInfra=true and NoManageHosts.
	// Optional.
	HostAdd []string `json:"hostadd"`

	// Hostname is the pod's hostname. If not set, the name of the pod will
	// be used (if a name was not provided here, the name auto-generated for
	// the pod will be used). This will be used by the infra container and
	// all containers in the pod as long as the UTS namespace is shared.
	// Optional.
	Hostname string `json:"hostname,omitempty"`

	// Image volumes bind-mount a container-image mount into the pod's infra container.
	// Optional.
	ImageVolumes []*ImageVolume `json:"image_volumes"`

	// InfraCommand sets the command that will be used to start the infra
	// container.
	// If not set, the default set in the Libpod configuration file will be
	// used.
	// Conflicts with NoInfra=true.
	// Optional.
	InfraCommand []string `json:"infra_command"`

	// InfraConmonPidFile is a custom path to store the infra container's
	// conmon PID.
	InfraConmonPidFile string `json:"infra_conmon_pid_file,omitempty"`

	// InfraImage is the image that will be used for the infra container.
	// If not set, the default set in the Libpod configuration file will be
	// used.
	// Conflicts with NoInfra=true.
	// Optional.
	InfraImage string `json:"infra_image,omitempty"`

	// InfraName is the name that will be used for the infra container.
	// If not set, the default set in the Libpod configuration file will be
	// used.
	// Conflicts with NoInfra=true.
	// Optional.
	InfraName string `json:"infra_name,omitempty"`

	// Labels are key-value pairs that are used to add metadata to pods.
	// Optional.
	Labels map[string]string `json:"labels,omitempty"`

	// Mounts are mounts that will be added to the pod.
	// These will supersede Image Volumes and VolumesFrom volumes where
	// there are conflicts.
	// Optional.
	Mounts []*Mount `json:"mounts"`

	// Name is the name of the pod.
	// If not provided, a name will be generated when the pod is created.
	// Optional.
	Name string `json:"name,omitempty"`

	// NetworkOptions are additional options for each network
	// Optional.
	NetworkOptions map[string][]string `json:"network_options,omitempty"`

	// Map of networks names to ids the container should join to.
	// You can request additional settings for each network, you can
	// set network aliases, static ips, static mac address  and the
	// network interface name for this container on the specific network.
	// If the map is empty and the bridge network mode is set the container
	// will be joined to the default network.
	Networks map[string]PerNetworkOptions `json:"Networks,omitempty"`

	// NoInfra tells the pod not to create an infra container. If this is
	// done, many networking-related options will become unavailable.
	// Conflicts with setting any options in PodNetworkConfig, and the
	// InfraCommand and InfraImages in this struct.
	// Optional.
	NoInfra bool `json:"no_infra,omitempty"`

	// NoManageHosts indicates that /etc/hosts should not be managed by the
	// pod. Instead, each container will create a separate /etc/hosts as
	// they would if not in a pod.
	// Conflicts with HostAdd.
	NoManageHosts bool `json:"no_manage_hosts,omitempty"`

	// NoManageResolvConf indicates that /etc/resolv.conf should not be
	// managed by the pod. Instead, each container will create and manage a
	// separate resolv.conf as if they had not joined a pod.
	// Conflicts with NoInfra=true and DNSServer, DNSSearch, DNSOption.
	// Optional.
	NoManageResolvConf bool `json:"no_manage_resolv_conf,omitempty"`

	// Overlay volumes are named volumes that will be added to the pod.
	// Optional.
	OverlayVolumes []*OverlayVolume `json:"overlay_volumes"`

	// pod create command
	PodCreateCommand []string `json:"pod_create_command"`

	// PortMappings is a set of ports to map into the infra container.
	// As, by default, containers share their network with the infra
	// container, this will forward the ports to the entire pod.
	// Only available if NetNS is set to Bridge or Slirp.
	// Optional.
	PortMappings []*PortMapping `json:"portmappings"`

	// security opt
	SecurityOpt []string `json:"security_opt"`

	// PodCreateCommand is the command used to create this pod.
	// This will be shown in the output of Inspect() on the pod, and may
	// also be used by some tools that wish to recreate the pod
	// (e.g. `podman generate systemd --new`).
	// Optional.
	// ShareParent determines if all containers in the pod will share the pod's cgroup as the cgroup parent
	ShareParent bool `json:"share_parent,omitempty"`

	// SharedNamespaces instructs the pod to share a set of namespaces.
	// Shared namespaces will be joined (by default) by every container
	// which joins the pod.
	// If not set and NoInfra is false, the pod will set a default set of
	// namespaces to share.
	// Conflicts with NoInfra=true.
	// Optional.
	SharedNamespaces []string `json:"shared_namespaces"`

	// Sysctl sets kernel parameters for the pod
	Sysctl map[string]string `json:"sysctl,omitempty"`

	// ThrottleReadBpsDevice contains the rate at which the devices in the pod can be read from/accessed
	ThrottleReadBpsDevice map[string]LinuxThrottleDevice `json:"throttleReadBpsDevice,omitempty"`

	// Volumes are named volumes that will be added to the pod.
	// These will supersede Image Volumes and VolumesFrom  volumes where
	// there are conflicts.
	// Optional.
	Volumes []*NamedVolume `json:"volumes"`

	// VolumesFrom is a set of containers whose volumes will be added to
	// this pod. The name or ID of the container must be provided, and
	// may optionally be followed by a : and then one or more
	// comma-separated options. Valid options are 'ro', 'rw', and 'z'.
	// Options will be used for all volumes sourced from the container.
	VolumesFrom []string `json:"volumes_from"`

	// netns
	Netns *Namespace `json:"netns,omitempty"`

	// pidns
	Pidns *Namespace `json:"pidns,omitempty"`

	// resource limits
	ResourceLimits *LinuxResources `json:"resource_limits,omitempty"`

	// userns
	Userns *Namespace `json:"userns,omitempty"`
}

// Validate validates this pod spec generator
func (m *PodSpecGenerator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverlayVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottleReadBpsDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePidns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodSpecGenerator) validateDNSServer(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSServer) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSServer); i++ {

		if err := m.DNSServer[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_server" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_server" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateImageVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.ImageVolumes); i++ {
		if swag.IsZero(m.ImageVolumes[i]) { // not required
			continue
		}

		if m.ImageVolumes[i] != nil {
			if err := m.ImageVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("image_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("image_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateMounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Mounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Mounts); i++ {
		if swag.IsZero(m.Mounts[i]) { // not required
			continue
		}

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for k := range m.Networks {

		if err := validate.Required("Networks"+"."+k, "body", m.Networks[k]); err != nil {
			return err
		}
		if val, ok := m.Networks[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Networks" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Networks" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateOverlayVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.OverlayVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.OverlayVolumes); i++ {
		if swag.IsZero(m.OverlayVolumes[i]) { // not required
			continue
		}

		if m.OverlayVolumes[i] != nil {
			if err := m.OverlayVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overlay_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("overlay_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validatePortMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.PortMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.PortMappings); i++ {
		if swag.IsZero(m.PortMappings[i]) { // not required
			continue
		}

		if m.PortMappings[i] != nil {
			if err := m.PortMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portmappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("portmappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateThrottleReadBpsDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottleReadBpsDevice) { // not required
		return nil
	}

	for k := range m.ThrottleReadBpsDevice {

		if err := validate.Required("throttleReadBpsDevice"+"."+k, "body", m.ThrottleReadBpsDevice[k]); err != nil {
			return err
		}
		if val, ok := m.ThrottleReadBpsDevice[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleReadBpsDevice" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleReadBpsDevice" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) validateNetns(formats strfmt.Registry) error {
	if swag.IsZero(m.Netns) { // not required
		return nil
	}

	if m.Netns != nil {
		if err := m.Netns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netns")
			}
			return err
		}
	}

	return nil
}

func (m *PodSpecGenerator) validatePidns(formats strfmt.Registry) error {
	if swag.IsZero(m.Pidns) { // not required
		return nil
	}

	if m.Pidns != nil {
		if err := m.Pidns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pidns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pidns")
			}
			return err
		}
	}

	return nil
}

func (m *PodSpecGenerator) validateResourceLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceLimits) { // not required
		return nil
	}

	if m.ResourceLimits != nil {
		if err := m.ResourceLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_limits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_limits")
			}
			return err
		}
	}

	return nil
}

func (m *PodSpecGenerator) validateUserns(formats strfmt.Registry) error {
	if swag.IsZero(m.Userns) { // not required
		return nil
	}

	if m.Userns != nil {
		if err := m.Userns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userns")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pod spec generator based on the context it is used
func (m *PodSpecGenerator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverlayVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottleReadBpsDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePidns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodSpecGenerator) contextValidateDNSServer(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServer); i++ {

		if err := m.DNSServer[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_server" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_server" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateImageVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImageVolumes); i++ {

		if m.ImageVolumes[i] != nil {
			if err := m.ImageVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("image_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("image_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mounts); i++ {

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Networks {

		if val, ok := m.Networks[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateOverlayVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OverlayVolumes); i++ {

		if m.OverlayVolumes[i] != nil {
			if err := m.OverlayVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("overlay_volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("overlay_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidatePortMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PortMappings); i++ {

		if m.PortMappings[i] != nil {
			if err := m.PortMappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portmappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("portmappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateThrottleReadBpsDevice(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ThrottleReadBpsDevice {

		if val, ok := m.ThrottleReadBpsDevice[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodSpecGenerator) contextValidateNetns(ctx context.Context, formats strfmt.Registry) error {

	if m.Netns != nil {
		if err := m.Netns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netns")
			}
			return err
		}
	}

	return nil
}

func (m *PodSpecGenerator) contextValidatePidns(ctx context.Context, formats strfmt.Registry) error {

	if m.Pidns != nil {
		if err := m.Pidns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pidns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pidns")
			}
			return err
		}
	}

	return nil
}

func (m *PodSpecGenerator) contextValidateResourceLimits(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceLimits != nil {
		if err := m.ResourceLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_limits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_limits")
			}
			return err
		}
	}

	return nil
}

func (m *PodSpecGenerator) contextValidateUserns(ctx context.Context, formats strfmt.Registry) error {

	if m.Userns != nil {
		if err := m.Userns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userns")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodSpecGenerator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodSpecGenerator) UnmarshalBinary(b []byte) error {
	var res PodSpecGenerator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
