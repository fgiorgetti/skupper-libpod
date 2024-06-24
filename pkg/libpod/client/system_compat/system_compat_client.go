// Code generated by go-swagger; DO NOT EDIT.

package system_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SystemAuth(params *SystemAuthParams, opts ...ClientOption) (*SystemAuthOK, error)

	SystemDataUsage(params *SystemDataUsageParams, opts ...ClientOption) (*SystemDataUsageOK, error)

	SystemEvents(params *SystemEventsParams, opts ...ClientOption) (*SystemEventsOK, error)

	SystemInfo(params *SystemInfoParams, opts ...ClientOption) (*SystemInfoOK, error)

	SystemPing(params *SystemPingParams, opts ...ClientOption) (*SystemPingOK, error)

	SystemVersion(params *SystemVersionParams, opts ...ClientOption) (*SystemVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SystemAuth checks auth configuration
*/
func (a *Client) SystemAuth(params *SystemAuthParams, opts ...ClientOption) (*SystemAuthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemAuth",
		Method:             "POST",
		PathPattern:        "/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemAuthReader{formats: a.formats},
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
	success, ok := result.(*SystemAuthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemDataUsage shows disk usage

Return information about disk usage for containers, images, and volumes
*/
func (a *Client) SystemDataUsage(params *SystemDataUsageParams, opts ...ClientOption) (*SystemDataUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemDataUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemDataUsage",
		Method:             "GET",
		PathPattern:        "/system/df",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemDataUsageReader{formats: a.formats},
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
	success, ok := result.(*SystemDataUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemDataUsage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemEvents gets events

Returns events filtered on query parameters
*/
func (a *Client) SystemEvents(params *SystemEventsParams, opts ...ClientOption) (*SystemEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemEvents",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemEventsReader{formats: a.formats},
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
	success, ok := result.(*SystemEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemInfo gets info

Returns information on the system and libpod configuration
*/
func (a *Client) SystemInfo(params *SystemInfoParams, opts ...ClientOption) (*SystemInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemInfo",
		Method:             "GET",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemInfoReader{formats: a.formats},
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
	success, ok := result.(*SystemInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SystemPing pings service

	Return protocol information in response headers.

`HEAD /libpod/_ping` is also supported.
`/_ping` is available for compatibility with other engines.
The '_ping' endpoints are not versioned.
*/
func (a *Client) SystemPing(params *SystemPingParams, opts ...ClientOption) (*SystemPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemPing",
		Method:             "GET",
		PathPattern:        "/libpod/_ping",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemPingReader{formats: a.formats},
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
	success, ok := result.(*SystemPingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemPing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SystemVersion components version information
*/
func (a *Client) SystemVersion(params *SystemVersionParams, opts ...ClientOption) (*SystemVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSystemVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SystemVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SystemVersionReader{formats: a.formats},
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
	success, ok := result.(*SystemVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SystemVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}