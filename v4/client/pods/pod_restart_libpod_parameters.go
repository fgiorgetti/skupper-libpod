// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// NewPodRestartLibpodParams creates a new PodRestartLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodRestartLibpodParams() *PodRestartLibpodParams {
	return &PodRestartLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodRestartLibpodParamsWithTimeout creates a new PodRestartLibpodParams object
// with the ability to set a timeout on a request.
func NewPodRestartLibpodParamsWithTimeout(timeout time.Duration) *PodRestartLibpodParams {
	return &PodRestartLibpodParams{
		timeout: timeout,
	}
}

// NewPodRestartLibpodParamsWithContext creates a new PodRestartLibpodParams object
// with the ability to set a context for a request.
func NewPodRestartLibpodParamsWithContext(ctx context.Context) *PodRestartLibpodParams {
	return &PodRestartLibpodParams{
		Context: ctx,
	}
}

// NewPodRestartLibpodParamsWithHTTPClient creates a new PodRestartLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodRestartLibpodParamsWithHTTPClient(client *http.Client) *PodRestartLibpodParams {
	return &PodRestartLibpodParams{
		HTTPClient: client,
	}
}

/*
PodRestartLibpodParams contains all the parameters to send to the API endpoint

	for the pod restart libpod operation.

	Typically these are written to a http.Request.
*/
type PodRestartLibpodParams struct {

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod restart libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodRestartLibpodParams) WithDefaults() *PodRestartLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod restart libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodRestartLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod restart libpod params
func (o *PodRestartLibpodParams) WithTimeout(timeout time.Duration) *PodRestartLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod restart libpod params
func (o *PodRestartLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod restart libpod params
func (o *PodRestartLibpodParams) WithContext(ctx context.Context) *PodRestartLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod restart libpod params
func (o *PodRestartLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod restart libpod params
func (o *PodRestartLibpodParams) WithHTTPClient(client *http.Client) *PodRestartLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod restart libpod params
func (o *PodRestartLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod restart libpod params
func (o *PodRestartLibpodParams) WithName(name string) *PodRestartLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod restart libpod params
func (o *PodRestartLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PodRestartLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
