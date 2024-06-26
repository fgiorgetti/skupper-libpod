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
)

// NewContainerUnpauseLibpodParams creates a new ContainerUnpauseLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerUnpauseLibpodParams() *ContainerUnpauseLibpodParams {
	return &ContainerUnpauseLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerUnpauseLibpodParamsWithTimeout creates a new ContainerUnpauseLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerUnpauseLibpodParamsWithTimeout(timeout time.Duration) *ContainerUnpauseLibpodParams {
	return &ContainerUnpauseLibpodParams{
		timeout: timeout,
	}
}

// NewContainerUnpauseLibpodParamsWithContext creates a new ContainerUnpauseLibpodParams object
// with the ability to set a context for a request.
func NewContainerUnpauseLibpodParamsWithContext(ctx context.Context) *ContainerUnpauseLibpodParams {
	return &ContainerUnpauseLibpodParams{
		Context: ctx,
	}
}

// NewContainerUnpauseLibpodParamsWithHTTPClient creates a new ContainerUnpauseLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerUnpauseLibpodParamsWithHTTPClient(client *http.Client) *ContainerUnpauseLibpodParams {
	return &ContainerUnpauseLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainerUnpauseLibpodParams contains all the parameters to send to the API endpoint

	for the container unpause libpod operation.

	Typically these are written to a http.Request.
*/
type ContainerUnpauseLibpodParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container unpause libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerUnpauseLibpodParams) WithDefaults() *ContainerUnpauseLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container unpause libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerUnpauseLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) WithTimeout(timeout time.Duration) *ContainerUnpauseLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) WithContext(ctx context.Context) *ContainerUnpauseLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) WithHTTPClient(client *http.Client) *ContainerUnpauseLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) WithName(name string) *ContainerUnpauseLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container unpause libpod params
func (o *ContainerUnpauseLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerUnpauseLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
