// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewContainerLogsLibpodParams creates a new ContainerLogsLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerLogsLibpodParams() *ContainerLogsLibpodParams {
	return &ContainerLogsLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerLogsLibpodParamsWithTimeout creates a new ContainerLogsLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerLogsLibpodParamsWithTimeout(timeout time.Duration) *ContainerLogsLibpodParams {
	return &ContainerLogsLibpodParams{
		timeout: timeout,
	}
}

// NewContainerLogsLibpodParamsWithContext creates a new ContainerLogsLibpodParams object
// with the ability to set a context for a request.
func NewContainerLogsLibpodParamsWithContext(ctx context.Context) *ContainerLogsLibpodParams {
	return &ContainerLogsLibpodParams{
		Context: ctx,
	}
}

// NewContainerLogsLibpodParamsWithHTTPClient creates a new ContainerLogsLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerLogsLibpodParamsWithHTTPClient(client *http.Client) *ContainerLogsLibpodParams {
	return &ContainerLogsLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainerLogsLibpodParams contains all the parameters to send to the API endpoint

	for the container logs libpod operation.

	Typically these are written to a http.Request.
*/
type ContainerLogsLibpodParams struct {

	/* Follow.

	   Keep connection after returning logs.
	*/
	Follow *bool

	/* Name.

	   the name or ID of the container
	*/
	Name string

	/* Since.

	   Only return logs since this time, as a UNIX timestamp
	*/
	Since *string

	/* Stderr.

	   Return logs from stderr
	*/
	Stderr *bool

	/* Stdout.

	   Return logs from stdout
	*/
	Stdout *bool

	/* Tail.

	   Only return this number of log lines from the end of the logs

	   Default: "all"
	*/
	Tail *string

	/* Timestamps.

	   Add timestamps to every log line
	*/
	Timestamps *bool

	/* Until.

	   Only return logs before this time, as a UNIX timestamp
	*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container logs libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerLogsLibpodParams) WithDefaults() *ContainerLogsLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container logs libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerLogsLibpodParams) SetDefaults() {
	var (
		tailDefault = string("all")

		timestampsDefault = bool(false)
	)

	val := ContainerLogsLibpodParams{
		Tail:       &tailDefault,
		Timestamps: &timestampsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithTimeout(timeout time.Duration) *ContainerLogsLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithContext(ctx context.Context) *ContainerLogsLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithHTTPClient(client *http.Client) *ContainerLogsLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFollow adds the follow to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithFollow(follow *bool) *ContainerLogsLibpodParams {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithName adds the name to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithName(name string) *ContainerLogsLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetName(name string) {
	o.Name = name
}

// WithSince adds the since to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithSince(since *string) *ContainerLogsLibpodParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetSince(since *string) {
	o.Since = since
}

// WithStderr adds the stderr to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithStderr(stderr *bool) *ContainerLogsLibpodParams {
	o.SetStderr(stderr)
	return o
}

// SetStderr adds the stderr to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetStderr(stderr *bool) {
	o.Stderr = stderr
}

// WithStdout adds the stdout to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithStdout(stdout *bool) *ContainerLogsLibpodParams {
	o.SetStdout(stdout)
	return o
}

// SetStdout adds the stdout to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetStdout(stdout *bool) {
	o.Stdout = stdout
}

// WithTail adds the tail to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithTail(tail *string) *ContainerLogsLibpodParams {
	o.SetTail(tail)
	return o
}

// SetTail adds the tail to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetTail(tail *string) {
	o.Tail = tail
}

// WithTimestamps adds the timestamps to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithTimestamps(timestamps *bool) *ContainerLogsLibpodParams {
	o.SetTimestamps(timestamps)
	return o
}

// SetTimestamps adds the timestamps to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetTimestamps(timestamps *bool) {
	o.Timestamps = timestamps
}

// WithUntil adds the until to the container logs libpod params
func (o *ContainerLogsLibpodParams) WithUntil(until *string) *ContainerLogsLibpodParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the container logs libpod params
func (o *ContainerLogsLibpodParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerLogsLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Follow != nil {

		// query param follow
		var qrFollow bool

		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {

			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Since != nil {

		// query param since
		var qrSince string

		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {

			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}
	}

	if o.Stderr != nil {

		// query param stderr
		var qrStderr bool

		if o.Stderr != nil {
			qrStderr = *o.Stderr
		}
		qStderr := swag.FormatBool(qrStderr)
		if qStderr != "" {

			if err := r.SetQueryParam("stderr", qStderr); err != nil {
				return err
			}
		}
	}

	if o.Stdout != nil {

		// query param stdout
		var qrStdout bool

		if o.Stdout != nil {
			qrStdout = *o.Stdout
		}
		qStdout := swag.FormatBool(qrStdout)
		if qStdout != "" {

			if err := r.SetQueryParam("stdout", qStdout); err != nil {
				return err
			}
		}
	}

	if o.Tail != nil {

		// query param tail
		var qrTail string

		if o.Tail != nil {
			qrTail = *o.Tail
		}
		qTail := qrTail
		if qTail != "" {

			if err := r.SetQueryParam("tail", qTail); err != nil {
				return err
			}
		}
	}

	if o.Timestamps != nil {

		// query param timestamps
		var qrTimestamps bool

		if o.Timestamps != nil {
			qrTimestamps = *o.Timestamps
		}
		qTimestamps := swag.FormatBool(qrTimestamps)
		if qTimestamps != "" {

			if err := r.SetQueryParam("timestamps", qTimestamps); err != nil {
				return err
			}
		}
	}

	if o.Until != nil {

		// query param until
		var qrUntil string

		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {

			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
