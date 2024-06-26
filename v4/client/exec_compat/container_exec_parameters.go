// Code generated by go-swagger; DO NOT EDIT.

package exec_compat

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

// NewContainerExecParams creates a new ContainerExecParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerExecParams() *ContainerExecParams {
	return &ContainerExecParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerExecParamsWithTimeout creates a new ContainerExecParams object
// with the ability to set a timeout on a request.
func NewContainerExecParamsWithTimeout(timeout time.Duration) *ContainerExecParams {
	return &ContainerExecParams{
		timeout: timeout,
	}
}

// NewContainerExecParamsWithContext creates a new ContainerExecParams object
// with the ability to set a context for a request.
func NewContainerExecParamsWithContext(ctx context.Context) *ContainerExecParams {
	return &ContainerExecParams{
		Context: ctx,
	}
}

// NewContainerExecParamsWithHTTPClient creates a new ContainerExecParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerExecParamsWithHTTPClient(client *http.Client) *ContainerExecParams {
	return &ContainerExecParams{
		HTTPClient: client,
	}
}

/*
ContainerExecParams contains all the parameters to send to the API endpoint

	for the container exec operation.

	Typically these are written to a http.Request.
*/
type ContainerExecParams struct {

	/* Control.

	   Attributes for create
	*/
	Control ContainerExecBody

	/* Name.

	   name of container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container exec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerExecParams) WithDefaults() *ContainerExecParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container exec params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerExecParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container exec params
func (o *ContainerExecParams) WithTimeout(timeout time.Duration) *ContainerExecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container exec params
func (o *ContainerExecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container exec params
func (o *ContainerExecParams) WithContext(ctx context.Context) *ContainerExecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container exec params
func (o *ContainerExecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container exec params
func (o *ContainerExecParams) WithHTTPClient(client *http.Client) *ContainerExecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container exec params
func (o *ContainerExecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithControl adds the control to the container exec params
func (o *ContainerExecParams) WithControl(control ContainerExecBody) *ContainerExecParams {
	o.SetControl(control)
	return o
}

// SetControl adds the control to the container exec params
func (o *ContainerExecParams) SetControl(control ContainerExecBody) {
	o.Control = control
}

// WithName adds the name to the container exec params
func (o *ContainerExecParams) WithName(name string) *ContainerExecParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container exec params
func (o *ContainerExecParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerExecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Control); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
