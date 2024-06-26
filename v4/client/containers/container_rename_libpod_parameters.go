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

// NewContainerRenameLibpodParams creates a new ContainerRenameLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerRenameLibpodParams() *ContainerRenameLibpodParams {
	return &ContainerRenameLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRenameLibpodParamsWithTimeout creates a new ContainerRenameLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerRenameLibpodParamsWithTimeout(timeout time.Duration) *ContainerRenameLibpodParams {
	return &ContainerRenameLibpodParams{
		timeout: timeout,
	}
}

// NewContainerRenameLibpodParamsWithContext creates a new ContainerRenameLibpodParams object
// with the ability to set a context for a request.
func NewContainerRenameLibpodParamsWithContext(ctx context.Context) *ContainerRenameLibpodParams {
	return &ContainerRenameLibpodParams{
		Context: ctx,
	}
}

// NewContainerRenameLibpodParamsWithHTTPClient creates a new ContainerRenameLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerRenameLibpodParamsWithHTTPClient(client *http.Client) *ContainerRenameLibpodParams {
	return &ContainerRenameLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainerRenameLibpodParams contains all the parameters to send to the API endpoint

	for the container rename libpod operation.

	Typically these are written to a http.Request.
*/
type ContainerRenameLibpodParams struct {

	/* Name.

	   Full or partial ID or full name of the container to rename
	*/
	PathName string

	/* Name.

	   New name for the container
	*/
	QueryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container rename libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerRenameLibpodParams) WithDefaults() *ContainerRenameLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container rename libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerRenameLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container rename libpod params
func (o *ContainerRenameLibpodParams) WithTimeout(timeout time.Duration) *ContainerRenameLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container rename libpod params
func (o *ContainerRenameLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container rename libpod params
func (o *ContainerRenameLibpodParams) WithContext(ctx context.Context) *ContainerRenameLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container rename libpod params
func (o *ContainerRenameLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container rename libpod params
func (o *ContainerRenameLibpodParams) WithHTTPClient(client *http.Client) *ContainerRenameLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container rename libpod params
func (o *ContainerRenameLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPathName adds the name to the container rename libpod params
func (o *ContainerRenameLibpodParams) WithPathName(name string) *ContainerRenameLibpodParams {
	o.SetPathName(name)
	return o
}

// SetPathName adds the name to the container rename libpod params
func (o *ContainerRenameLibpodParams) SetPathName(name string) {
	o.PathName = name
}

// WithQueryName adds the name to the container rename libpod params
func (o *ContainerRenameLibpodParams) WithQueryName(name string) *ContainerRenameLibpodParams {
	o.SetQueryName(name)
	return o
}

// SetQueryName adds the name to the container rename libpod params
func (o *ContainerRenameLibpodParams) SetQueryName(name string) {
	o.QueryName = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRenameLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.PathName); err != nil {
		return err
	}

	// query param name
	qrName := o.QueryName
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
