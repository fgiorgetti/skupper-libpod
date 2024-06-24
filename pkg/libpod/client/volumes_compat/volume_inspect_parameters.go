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

// NewVolumeInspectParams creates a new VolumeInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeInspectParams() *VolumeInspectParams {
	return &VolumeInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeInspectParamsWithTimeout creates a new VolumeInspectParams object
// with the ability to set a timeout on a request.
func NewVolumeInspectParamsWithTimeout(timeout time.Duration) *VolumeInspectParams {
	return &VolumeInspectParams{
		timeout: timeout,
	}
}

// NewVolumeInspectParamsWithContext creates a new VolumeInspectParams object
// with the ability to set a context for a request.
func NewVolumeInspectParamsWithContext(ctx context.Context) *VolumeInspectParams {
	return &VolumeInspectParams{
		Context: ctx,
	}
}

// NewVolumeInspectParamsWithHTTPClient creates a new VolumeInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeInspectParamsWithHTTPClient(client *http.Client) *VolumeInspectParams {
	return &VolumeInspectParams{
		HTTPClient: client,
	}
}

/*
VolumeInspectParams contains all the parameters to send to the API endpoint

	for the volume inspect operation.

	Typically these are written to a http.Request.
*/
type VolumeInspectParams struct {

	/* Name.

	   the name or ID of the volume
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInspectParams) WithDefaults() *VolumeInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume inspect params
func (o *VolumeInspectParams) WithTimeout(timeout time.Duration) *VolumeInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume inspect params
func (o *VolumeInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume inspect params
func (o *VolumeInspectParams) WithContext(ctx context.Context) *VolumeInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume inspect params
func (o *VolumeInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume inspect params
func (o *VolumeInspectParams) WithHTTPClient(client *http.Client) *VolumeInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume inspect params
func (o *VolumeInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the volume inspect params
func (o *VolumeInspectParams) WithName(name string) *VolumeInspectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the volume inspect params
func (o *VolumeInspectParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
