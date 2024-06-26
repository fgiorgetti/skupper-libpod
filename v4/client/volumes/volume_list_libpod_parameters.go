// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewVolumeListLibpodParams creates a new VolumeListLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeListLibpodParams() *VolumeListLibpodParams {
	return &VolumeListLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeListLibpodParamsWithTimeout creates a new VolumeListLibpodParams object
// with the ability to set a timeout on a request.
func NewVolumeListLibpodParamsWithTimeout(timeout time.Duration) *VolumeListLibpodParams {
	return &VolumeListLibpodParams{
		timeout: timeout,
	}
}

// NewVolumeListLibpodParamsWithContext creates a new VolumeListLibpodParams object
// with the ability to set a context for a request.
func NewVolumeListLibpodParamsWithContext(ctx context.Context) *VolumeListLibpodParams {
	return &VolumeListLibpodParams{
		Context: ctx,
	}
}

// NewVolumeListLibpodParamsWithHTTPClient creates a new VolumeListLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeListLibpodParamsWithHTTPClient(client *http.Client) *VolumeListLibpodParams {
	return &VolumeListLibpodParams{
		HTTPClient: client,
	}
}

/*
VolumeListLibpodParams contains all the parameters to send to the API endpoint

	for the volume list libpod operation.

	Typically these are written to a http.Request.
*/
type VolumeListLibpodParams struct {

	/* Filters.

	   JSON encoded value of the filters (a map[string][]string) to process on the volumes list. Available filters:
	- driver=<volume-driver-name> Matches volumes based on their driver.
	- label=<key> or label=<key>:<value> Matches volumes based on the presence of a label alone or a label and a value.
	- name=<volume-name> Matches all of volume name.
	- opt=<driver-option> Matches a storage driver options
	- `until=<timestamp>` List volumes created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.

	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume list libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeListLibpodParams) WithDefaults() *VolumeListLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume list libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeListLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume list libpod params
func (o *VolumeListLibpodParams) WithTimeout(timeout time.Duration) *VolumeListLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume list libpod params
func (o *VolumeListLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume list libpod params
func (o *VolumeListLibpodParams) WithContext(ctx context.Context) *VolumeListLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume list libpod params
func (o *VolumeListLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume list libpod params
func (o *VolumeListLibpodParams) WithHTTPClient(client *http.Client) *VolumeListLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume list libpod params
func (o *VolumeListLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the volume list libpod params
func (o *VolumeListLibpodParams) WithFilters(filters *string) *VolumeListLibpodParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the volume list libpod params
func (o *VolumeListLibpodParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeListLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
