// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new volumes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for volumes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	VolumeCreateLibpod(params *VolumeCreateLibpodParams, opts ...ClientOption) (*VolumeCreateLibpodCreated, error)

	VolumeDeleteLibpod(params *VolumeDeleteLibpodParams, opts ...ClientOption) (*VolumeDeleteLibpodNoContent, error)

	VolumeExistsLibpod(params *VolumeExistsLibpodParams, opts ...ClientOption) (*VolumeExistsLibpodNoContent, error)

	VolumeInspectLibpod(params *VolumeInspectLibpodParams, opts ...ClientOption) (*VolumeInspectLibpodOK, error)

	VolumeListLibpod(params *VolumeListLibpodParams, opts ...ClientOption) (*VolumeListLibpodOK, error)

	VolumePruneLibpod(params *VolumePruneLibpodParams, opts ...ClientOption) (*VolumePruneLibpodOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
VolumeCreateLibpod creates a volume
*/
func (a *Client) VolumeCreateLibpod(params *VolumeCreateLibpodParams, opts ...ClientOption) (*VolumeCreateLibpodCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeCreateLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeCreateLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/volumes/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeCreateLibpodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeCreateLibpodCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeCreateLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VolumeDeleteLibpod removes volume
*/
func (a *Client) VolumeDeleteLibpod(params *VolumeDeleteLibpodParams, opts ...ClientOption) (*VolumeDeleteLibpodNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeDeleteLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeDeleteLibpod",
		Method:             "DELETE",
		PathPattern:        "/libpod/volumes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeDeleteLibpodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeDeleteLibpodNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeDeleteLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VolumeExistsLibpod volumes exists

Check if a volume exists
*/
func (a *Client) VolumeExistsLibpod(params *VolumeExistsLibpodParams, opts ...ClientOption) (*VolumeExistsLibpodNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeExistsLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeExistsLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/volumes/{name}/exists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeExistsLibpodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeExistsLibpodNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeExistsLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VolumeInspectLibpod inspects volume
*/
func (a *Client) VolumeInspectLibpod(params *VolumeInspectLibpodParams, opts ...ClientOption) (*VolumeInspectLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeInspectLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeInspectLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/volumes/{name}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeInspectLibpodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeInspectLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeInspectLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VolumeListLibpod lists volumes

Returns a list of volumes
*/
func (a *Client) VolumeListLibpod(params *VolumeListLibpodParams, opts ...ClientOption) (*VolumeListLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeListLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeListLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/volumes/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeListLibpodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeListLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeListLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VolumePruneLibpod prunes volumes
*/
func (a *Client) VolumePruneLibpod(params *VolumePruneLibpodParams, opts ...ClientOption) (*VolumePruneLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumePruneLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumePruneLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/volumes/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumePruneLibpodReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumePruneLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumePruneLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}