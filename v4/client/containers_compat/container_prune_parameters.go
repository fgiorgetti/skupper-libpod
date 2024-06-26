// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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
)

// NewContainerPruneParams creates a new ContainerPruneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerPruneParams() *ContainerPruneParams {
	return &ContainerPruneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerPruneParamsWithTimeout creates a new ContainerPruneParams object
// with the ability to set a timeout on a request.
func NewContainerPruneParamsWithTimeout(timeout time.Duration) *ContainerPruneParams {
	return &ContainerPruneParams{
		timeout: timeout,
	}
}

// NewContainerPruneParamsWithContext creates a new ContainerPruneParams object
// with the ability to set a context for a request.
func NewContainerPruneParamsWithContext(ctx context.Context) *ContainerPruneParams {
	return &ContainerPruneParams{
		Context: ctx,
	}
}

// NewContainerPruneParamsWithHTTPClient creates a new ContainerPruneParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerPruneParamsWithHTTPClient(client *http.Client) *ContainerPruneParams {
	return &ContainerPruneParams{
		HTTPClient: client,
	}
}

/*
ContainerPruneParams contains all the parameters to send to the API endpoint

	for the container prune operation.

	Typically these are written to a http.Request.
*/
type ContainerPruneParams struct {

	/* Filters.

	    Filters to process on the prune list, encoded as JSON (a `map[string][]string`).  Available filters:
	- `until=<timestamp>` Prune containers created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	- `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune containers with (or without, in case `label!=...` is used) the specified labels.

	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container prune params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerPruneParams) WithDefaults() *ContainerPruneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container prune params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerPruneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container prune params
func (o *ContainerPruneParams) WithTimeout(timeout time.Duration) *ContainerPruneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container prune params
func (o *ContainerPruneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container prune params
func (o *ContainerPruneParams) WithContext(ctx context.Context) *ContainerPruneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container prune params
func (o *ContainerPruneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container prune params
func (o *ContainerPruneParams) WithHTTPClient(client *http.Client) *ContainerPruneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container prune params
func (o *ContainerPruneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the container prune params
func (o *ContainerPruneParams) WithFilters(filters *string) *ContainerPruneParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the container prune params
func (o *ContainerPruneParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerPruneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
