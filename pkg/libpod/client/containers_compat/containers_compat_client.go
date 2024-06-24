// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new containers compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for containers compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ContainerArchive(params *ContainerArchiveParams, writer io.Writer, opts ...ClientOption) (*ContainerArchiveOK, error)

	ContainerArchiveLibpod(params *ContainerArchiveLibpodParams, writer io.Writer, opts ...ClientOption) (*ContainerArchiveLibpodOK, error)

	ContainerAttach(params *ContainerAttachParams, opts ...ClientOption) error

	ContainerCreate(params *ContainerCreateParams, opts ...ClientOption) (*ContainerCreateCreated, error)

	ContainerDelete(params *ContainerDeleteParams, opts ...ClientOption) (*ContainerDeleteNoContent, error)

	ContainerExport(params *ContainerExportParams, opts ...ClientOption) (*ContainerExportOK, error)

	ContainerInspect(params *ContainerInspectParams, opts ...ClientOption) (*ContainerInspectOK, error)

	ContainerKill(params *ContainerKillParams, opts ...ClientOption) (*ContainerKillNoContent, error)

	ContainerList(params *ContainerListParams, opts ...ClientOption) (*ContainerListOK, error)

	ContainerLogs(params *ContainerLogsParams, opts ...ClientOption) (*ContainerLogsOK, error)

	ContainerPause(params *ContainerPauseParams, opts ...ClientOption) (*ContainerPauseNoContent, error)

	ContainerPrune(params *ContainerPruneParams, opts ...ClientOption) (*ContainerPruneOK, error)

	ContainerRename(params *ContainerRenameParams, opts ...ClientOption) (*ContainerRenameNoContent, error)

	ContainerResize(params *ContainerResizeParams, opts ...ClientOption) (*ContainerResizeOK, error)

	ContainerRestart(params *ContainerRestartParams, opts ...ClientOption) (*ContainerRestartNoContent, error)

	ContainerStart(params *ContainerStartParams, opts ...ClientOption) (*ContainerStartNoContent, error)

	ContainerStats(params *ContainerStatsParams, opts ...ClientOption) (*ContainerStatsOK, error)

	ContainerStop(params *ContainerStopParams, opts ...ClientOption) (*ContainerStopNoContent, error)

	ContainerTop(params *ContainerTopParams, opts ...ClientOption) (*ContainerTopOK, error)

	ContainerUnpause(params *ContainerUnpauseParams, opts ...ClientOption) (*ContainerUnpauseNoContent, error)

	ContainerWait(params *ContainerWaitParams, opts ...ClientOption) (*ContainerWaitOK, error)

	ImageCommit(params *ImageCommitParams, opts ...ClientOption) (*ImageCommitCreated, error)

	PutContainerArchive(params *PutContainerArchiveParams, opts ...ClientOption) (*PutContainerArchiveOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ContainerArchive gets files from a container

Get a tar archive of files from a container
*/
func (a *Client) ContainerArchive(params *ContainerArchiveParams, writer io.Writer, opts ...ClientOption) (*ContainerArchiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerArchiveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerArchive",
		Method:             "GET",
		PathPattern:        "/containers/{name}/archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerArchiveReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*ContainerArchiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerArchive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerArchiveLibpod copies files from a container

Copy a tar archive of files from a container
*/
func (a *Client) ContainerArchiveLibpod(params *ContainerArchiveLibpodParams, writer io.Writer, opts ...ClientOption) (*ContainerArchiveLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerArchiveLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerArchiveLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/containers/{name}/archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerArchiveLibpodReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*ContainerArchiveLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerArchiveLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerAttach attaches to a container

Hijacks the connection to forward the container's standard streams to the client.
*/
func (a *Client) ContainerAttach(params *ContainerAttachParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerAttachParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerAttach",
		Method:             "POST",
		PathPattern:        "/containers/{name}/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerAttachReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
ContainerCreate creates a container
*/
func (a *Client) ContainerCreate(params *ContainerCreateParams, opts ...ClientOption) (*ContainerCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerCreate",
		Method:             "POST",
		PathPattern:        "/containers/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerCreateReader{formats: a.formats},
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
	success, ok := result.(*ContainerCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerDelete removes a container
*/
func (a *Client) ContainerDelete(params *ContainerDeleteParams, opts ...ClientOption) (*ContainerDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerDelete",
		Method:             "DELETE",
		PathPattern:        "/containers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerDeleteReader{formats: a.formats},
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
	success, ok := result.(*ContainerDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerExport exports a container

Export the contents of a container as a tarball.
*/
func (a *Client) ContainerExport(params *ContainerExportParams, opts ...ClientOption) (*ContainerExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerExport",
		Method:             "GET",
		PathPattern:        "/containers/{name}/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerExportReader{formats: a.formats},
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
	success, ok := result.(*ContainerExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerExport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerInspect inspects container

Return low-level information about a container.
*/
func (a *Client) ContainerInspect(params *ContainerInspectParams, opts ...ClientOption) (*ContainerInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerInspect",
		Method:             "GET",
		PathPattern:        "/containers/{name}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerInspectReader{formats: a.formats},
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
	success, ok := result.(*ContainerInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerKill kills container

Signal to send to the container as an integer or string (e.g. SIGINT)
*/
func (a *Client) ContainerKill(params *ContainerKillParams, opts ...ClientOption) (*ContainerKillNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerKillParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerKill",
		Method:             "POST",
		PathPattern:        "/containers/{name}/kill",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerKillReader{formats: a.formats},
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
	success, ok := result.(*ContainerKillNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerKill: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerList lists containers

Returns a list of containers
*/
func (a *Client) ContainerList(params *ContainerListParams, opts ...ClientOption) (*ContainerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerList",
		Method:             "GET",
		PathPattern:        "/containers/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerListReader{formats: a.formats},
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
	success, ok := result.(*ContainerListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerLogs gets container logs

Get stdout and stderr logs from a container.
*/
func (a *Client) ContainerLogs(params *ContainerLogsParams, opts ...ClientOption) (*ContainerLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerLogs",
		Method:             "GET",
		PathPattern:        "/containers/{name}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerLogsReader{formats: a.formats},
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
	success, ok := result.(*ContainerLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerPause pauses container

Use the cgroups freezer to suspend all processes in a container.
*/
func (a *Client) ContainerPause(params *ContainerPauseParams, opts ...ClientOption) (*ContainerPauseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerPauseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerPause",
		Method:             "POST",
		PathPattern:        "/containers/{name}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerPauseReader{formats: a.formats},
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
	success, ok := result.(*ContainerPauseNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerPause: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerPrune deletes stopped containers

Remove containers not in use
*/
func (a *Client) ContainerPrune(params *ContainerPruneParams, opts ...ClientOption) (*ContainerPruneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerPruneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerPrune",
		Method:             "POST",
		PathPattern:        "/containers/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerPruneReader{formats: a.formats},
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
	success, ok := result.(*ContainerPruneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerPrune: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerRename renames an existing container

Change the name of an existing container.
*/
func (a *Client) ContainerRename(params *ContainerRenameParams, opts ...ClientOption) (*ContainerRenameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerRenameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerRename",
		Method:             "POST",
		PathPattern:        "/containers/{name}/rename",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerRenameReader{formats: a.formats},
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
	success, ok := result.(*ContainerRenameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerRename: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerResize resizes a container s t t y

Resize the terminal attached to a container (for use with Attach).
*/
func (a *Client) ContainerResize(params *ContainerResizeParams, opts ...ClientOption) (*ContainerResizeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerResizeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerResize",
		Method:             "POST",
		PathPattern:        "/containers/{name}/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerResizeReader{formats: a.formats},
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
	success, ok := result.(*ContainerResizeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerResize: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerRestart restarts container
*/
func (a *Client) ContainerRestart(params *ContainerRestartParams, opts ...ClientOption) (*ContainerRestartNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerRestartParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerRestart",
		Method:             "POST",
		PathPattern:        "/containers/{name}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerRestartReader{formats: a.formats},
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
	success, ok := result.(*ContainerRestartNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerRestart: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerStart starts a container
*/
func (a *Client) ContainerStart(params *ContainerStartParams, opts ...ClientOption) (*ContainerStartNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerStartParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerStart",
		Method:             "POST",
		PathPattern:        "/containers/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerStartReader{formats: a.formats},
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
	success, ok := result.(*ContainerStartNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerStart: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerStats gets stats for a container

This returns a live stream of a container’s resource usage statistics.
*/
func (a *Client) ContainerStats(params *ContainerStatsParams, opts ...ClientOption) (*ContainerStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerStats",
		Method:             "GET",
		PathPattern:        "/containers/{name}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerStatsReader{formats: a.formats},
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
	success, ok := result.(*ContainerStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerStats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerStop stops a container

Stop a container
*/
func (a *Client) ContainerStop(params *ContainerStopParams, opts ...ClientOption) (*ContainerStopNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerStopParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerStop",
		Method:             "POST",
		PathPattern:        "/containers/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerStopReader{formats: a.formats},
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
	success, ok := result.(*ContainerStopNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerStop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerTop lists processes running inside a container
*/
func (a *Client) ContainerTop(params *ContainerTopParams, opts ...ClientOption) (*ContainerTopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerTopParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerTop",
		Method:             "GET",
		PathPattern:        "/containers/{name}/top",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerTopReader{formats: a.formats},
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
	success, ok := result.(*ContainerTopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerTop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerUnpause unpauses container

Resume a paused container
*/
func (a *Client) ContainerUnpause(params *ContainerUnpauseParams, opts ...ClientOption) (*ContainerUnpauseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerUnpauseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerUnpause",
		Method:             "POST",
		PathPattern:        "/containers/{name}/unpause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerUnpauseReader{formats: a.formats},
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
	success, ok := result.(*ContainerUnpauseNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerUnpause: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ContainerWait waits on a container

Block until a container stops or given condition is met.
*/
func (a *Client) ContainerWait(params *ContainerWaitParams, opts ...ClientOption) (*ContainerWaitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerWaitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContainerWait",
		Method:             "POST",
		PathPattern:        "/containers/{name}/wait",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerWaitReader{formats: a.formats},
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
	success, ok := result.(*ContainerWaitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ContainerWait: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImageCommit news image

Create a new image from a container
*/
func (a *Client) ImageCommit(params *ImageCommitParams, opts ...ClientOption) (*ImageCommitCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageCommitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImageCommit",
		Method:             "POST",
		PathPattern:        "/commit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageCommitReader{formats: a.formats},
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
	success, ok := result.(*ImageCommitCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImageCommit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutContainerArchive puts files into a container

Put a tar archive of files into a container
*/
func (a *Client) PutContainerArchive(params *PutContainerArchiveParams, opts ...ClientOption) (*PutContainerArchiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutContainerArchiveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutContainerArchive",
		Method:             "PUT",
		PathPattern:        "/containers/{name}/archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutContainerArchiveReader{formats: a.formats},
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
	success, ok := result.(*PutContainerArchiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutContainerArchive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}