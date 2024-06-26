// Code generated by go-swagger; DO NOT EDIT.

package exec_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new exec compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for exec compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ContainerExec(params *ContainerExecParams, opts ...ClientOption) (*ContainerExecCreated, error)

	ExecInspect(params *ExecInspectParams, opts ...ClientOption) (*ExecInspectOK, error)

	ExecResize(params *ExecResizeParams, opts ...ClientOption) (*ExecResizeCreated, error)

	ExecStart(params *ExecStartParams, opts ...ClientOption) (*ExecStartOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ContainerExec creates an exec instance

Create an exec session to run a command inside a running container. Exec sessions will be automatically removed 5 minutes after they exit.
*/
func (a *Client) ContainerExec(params *ContainerExecParams, opts ...ClientOption) (*ContainerExecCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerExecParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerExec",
		Method:             "POST",
		PathPattern:        "/containers/{name}/exec",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerExecReader{formats: a.formats},
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
	success, ok := result.(*ContainerExecCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerExec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecInspect inspects an exec instance

Return low-level information about an exec instance.
*/
func (a *Client) ExecInspect(params *ExecInspectParams, opts ...ClientOption) (*ExecInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecInspect",
		Method:             "GET",
		PathPattern:        "/exec/{id}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExecInspectReader{formats: a.formats},
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
	success, ok := result.(*ExecInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecResize resizes an exec instance

Resize the TTY session used by an exec instance. This endpoint only works if tty was specified as part of creating and starting the exec instance.
*/
func (a *Client) ExecResize(params *ExecResizeParams, opts ...ClientOption) (*ExecResizeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecResizeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecResize",
		Method:             "POST",
		PathPattern:        "/exec/{id}/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExecResizeReader{formats: a.formats},
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
	success, ok := result.(*ExecResizeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecResize: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecStart starts an exec instance

Starts a previously set up exec instance. If detach is true, this endpoint returns immediately after starting the command. Otherwise, it sets up an interactive session with the command.
*/
func (a *Client) ExecStart(params *ExecStartParams, opts ...ClientOption) (*ExecStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecStartParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ExecStart",
		Method:             "POST",
		PathPattern:        "/exec/{id}/start",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExecStartReader{formats: a.formats},
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
	success, ok := result.(*ExecStartOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ExecStart: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
