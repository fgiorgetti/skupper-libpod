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

// NewContainersStatsAllLibpodParams creates a new ContainersStatsAllLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainersStatsAllLibpodParams() *ContainersStatsAllLibpodParams {
	return &ContainersStatsAllLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainersStatsAllLibpodParamsWithTimeout creates a new ContainersStatsAllLibpodParams object
// with the ability to set a timeout on a request.
func NewContainersStatsAllLibpodParamsWithTimeout(timeout time.Duration) *ContainersStatsAllLibpodParams {
	return &ContainersStatsAllLibpodParams{
		timeout: timeout,
	}
}

// NewContainersStatsAllLibpodParamsWithContext creates a new ContainersStatsAllLibpodParams object
// with the ability to set a context for a request.
func NewContainersStatsAllLibpodParamsWithContext(ctx context.Context) *ContainersStatsAllLibpodParams {
	return &ContainersStatsAllLibpodParams{
		Context: ctx,
	}
}

// NewContainersStatsAllLibpodParamsWithHTTPClient creates a new ContainersStatsAllLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainersStatsAllLibpodParamsWithHTTPClient(client *http.Client) *ContainersStatsAllLibpodParams {
	return &ContainersStatsAllLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainersStatsAllLibpodParams contains all the parameters to send to the API endpoint

	for the containers stats all libpod operation.

	Typically these are written to a http.Request.
*/
type ContainersStatsAllLibpodParams struct {

	/* Containers.

	   names or IDs of containers
	*/
	Containers []string

	/* Interval.

	   Time in seconds between stats reports

	   Default: 5
	*/
	Interval *int64

	/* Stream.

	   Stream the output

	   Default: true
	*/
	Stream *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the containers stats all libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainersStatsAllLibpodParams) WithDefaults() *ContainersStatsAllLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the containers stats all libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainersStatsAllLibpodParams) SetDefaults() {
	var (
		intervalDefault = int64(5)

		streamDefault = bool(true)
	)

	val := ContainersStatsAllLibpodParams{
		Interval: &intervalDefault,
		Stream:   &streamDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) WithTimeout(timeout time.Duration) *ContainersStatsAllLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) WithContext(ctx context.Context) *ContainersStatsAllLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) WithHTTPClient(client *http.Client) *ContainersStatsAllLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContainers adds the containers to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) WithContainers(containers []string) *ContainersStatsAllLibpodParams {
	o.SetContainers(containers)
	return o
}

// SetContainers adds the containers to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) SetContainers(containers []string) {
	o.Containers = containers
}

// WithInterval adds the interval to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) WithInterval(interval *int64) *ContainersStatsAllLibpodParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) SetInterval(interval *int64) {
	o.Interval = interval
}

// WithStream adds the stream to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) WithStream(stream *bool) *ContainersStatsAllLibpodParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the containers stats all libpod params
func (o *ContainersStatsAllLibpodParams) SetStream(stream *bool) {
	o.Stream = stream
}

// WriteToRequest writes these params to a swagger request
func (o *ContainersStatsAllLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Containers != nil {

		// binding items for containers
		joinedContainers := o.bindParamContainers(reg)

		// query array param containers
		if err := r.SetQueryParam("containers", joinedContainers...); err != nil {
			return err
		}
	}

	if o.Interval != nil {

		// query param interval
		var qrInterval int64

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := swag.FormatInt64(qrInterval)
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	if o.Stream != nil {

		// query param stream
		var qrStream bool

		if o.Stream != nil {
			qrStream = *o.Stream
		}
		qStream := swag.FormatBool(qrStream)
		if qStream != "" {

			if err := r.SetQueryParam("stream", qStream); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamContainersStatsAllLibpod binds the parameter containers
func (o *ContainersStatsAllLibpodParams) bindParamContainers(formats strfmt.Registry) []string {
	containersIR := o.Containers

	var containersIC []string
	for _, containersIIR := range containersIR { // explode []string

		containersIIV := containersIIR // string as string
		containersIC = append(containersIC, containersIIV)
	}

	// items.CollectionFormat: ""
	containersIS := swag.JoinByFormat(containersIC, "")

	return containersIS
}
