// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SystemDataUsageLibpod(params *SystemDataUsageLibpodParams, opts ...ClientOption) (*SystemDataUsageLibpodOK, error)

	SystemEventsLibpod(params *SystemEventsLibpodParams, opts ...ClientOption) (*SystemEventsLibpodOK, error)

	SystemInfoLibpod(params *SystemInfoLibpodParams, opts ...ClientOption) (*SystemInfoLibpodOK, error)

	SystemPruneLibpod(params *SystemPruneLibpodParams, opts ...ClientOption) (*SystemPruneLibpodOK, error)

	SystemVersionLibpod(params *SystemVersionLibpodParams, opts ...ClientOption) (*SystemVersionLibpodOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SystemDataUsageLibpod shows disk usage

Return information about disk usage for containers, images, and volumes
*/
func (a *Client) SystemDataUsageLibpod(params *SystemDataUsageLibpodParams, opts ...ClientOption) (*SystemDataUsageLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemDataUsageLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemDataUsageLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/system/df",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemDataUsageLibpodReader{formats: a.formats},
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
	success, ok := result.(*SystemDataUsageLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemDataUsageLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemEventsLibpod gets events

Returns events filtered on query parameters
*/
func (a *Client) SystemEventsLibpod(params *SystemEventsLibpodParams, opts ...ClientOption) (*SystemEventsLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemEventsLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemEventsLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemEventsLibpodReader{formats: a.formats},
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
	success, ok := result.(*SystemEventsLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemEventsLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemInfoLibpod gets info

Returns information on the system and libpod configuration
*/
func (a *Client) SystemInfoLibpod(params *SystemInfoLibpodParams, opts ...ClientOption) (*SystemInfoLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemInfoLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemInfoLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemInfoLibpodReader{formats: a.formats},
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
	success, ok := result.(*SystemInfoLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemInfoLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemPruneLibpod prunes unused data
*/
func (a *Client) SystemPruneLibpod(params *SystemPruneLibpodParams, opts ...ClientOption) (*SystemPruneLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemPruneLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemPruneLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/system/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemPruneLibpodReader{formats: a.formats},
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
	success, ok := result.(*SystemPruneLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemPruneLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemVersionLibpod components version information
*/
func (a *Client) SystemVersionLibpod(params *SystemVersionLibpodParams, opts ...ClientOption) (*SystemVersionLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemVersionLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemVersionLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemVersionLibpodReader{formats: a.formats},
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
	success, ok := result.(*SystemVersionLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemVersionLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
