// Code generated by go-swagger; DO NOT EDIT.

package volumes_compat

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

// NewVolumePruneParams creates a new VolumePruneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumePruneParams() *VolumePruneParams {
	return &VolumePruneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumePruneParamsWithTimeout creates a new VolumePruneParams object
// with the ability to set a timeout on a request.
func NewVolumePruneParamsWithTimeout(timeout time.Duration) *VolumePruneParams {
	return &VolumePruneParams{
		timeout: timeout,
	}
}

// NewVolumePruneParamsWithContext creates a new VolumePruneParams object
// with the ability to set a context for a request.
func NewVolumePruneParamsWithContext(ctx context.Context) *VolumePruneParams {
	return &VolumePruneParams{
		Context: ctx,
	}
}

// NewVolumePruneParamsWithHTTPClient creates a new VolumePruneParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumePruneParamsWithHTTPClient(client *http.Client) *VolumePruneParams {
	return &VolumePruneParams{
		HTTPClient: client,
	}
}

/*
VolumePruneParams contains all the parameters to send to the API endpoint

	for the volume prune operation.

	Typically these are written to a http.Request.
*/
type VolumePruneParams struct {

	/* Filters.

	     JSON encoded value of filters (a map[string][]string) to match volumes against before pruning.
	Available filters:
	  - `until=<timestamp>` Prune volumes created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	  - `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune volumes with (or without, in case `label!=...` is used) the specified labels.

	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume prune params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumePruneParams) WithDefaults() *VolumePruneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume prune params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumePruneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume prune params
func (o *VolumePruneParams) WithTimeout(timeout time.Duration) *VolumePruneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume prune params
func (o *VolumePruneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume prune params
func (o *VolumePruneParams) WithContext(ctx context.Context) *VolumePruneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume prune params
func (o *VolumePruneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume prune params
func (o *VolumePruneParams) WithHTTPClient(client *http.Client) *VolumePruneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume prune params
func (o *VolumePruneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the volume prune params
func (o *VolumePruneParams) WithFilters(filters *string) *VolumePruneParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the volume prune params
func (o *VolumePruneParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *VolumePruneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
