// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

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

	"github.com/fgiorgetti/skupper-libpod/pkg/libpod/models"
)

// NewNetworkConnectParams creates a new NetworkConnectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkConnectParams() *NetworkConnectParams {
	return &NetworkConnectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkConnectParamsWithTimeout creates a new NetworkConnectParams object
// with the ability to set a timeout on a request.
func NewNetworkConnectParamsWithTimeout(timeout time.Duration) *NetworkConnectParams {
	return &NetworkConnectParams{
		timeout: timeout,
	}
}

// NewNetworkConnectParamsWithContext creates a new NetworkConnectParams object
// with the ability to set a context for a request.
func NewNetworkConnectParamsWithContext(ctx context.Context) *NetworkConnectParams {
	return &NetworkConnectParams{
		Context: ctx,
	}
}

// NewNetworkConnectParamsWithHTTPClient creates a new NetworkConnectParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkConnectParamsWithHTTPClient(client *http.Client) *NetworkConnectParams {
	return &NetworkConnectParams{
		HTTPClient: client,
	}
}

/*
NetworkConnectParams contains all the parameters to send to the API endpoint

	for the network connect operation.

	Typically these are written to a http.Request.
*/
type NetworkConnectParams struct {

	/* Create.

	   attributes for connecting a container to a network
	*/
	Create *models.SwagCompatNetworkConnectRequest

	/* Name.

	   the name of the network
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network connect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkConnectParams) WithDefaults() *NetworkConnectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network connect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkConnectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network connect params
func (o *NetworkConnectParams) WithTimeout(timeout time.Duration) *NetworkConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network connect params
func (o *NetworkConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network connect params
func (o *NetworkConnectParams) WithContext(ctx context.Context) *NetworkConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network connect params
func (o *NetworkConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network connect params
func (o *NetworkConnectParams) WithHTTPClient(client *http.Client) *NetworkConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network connect params
func (o *NetworkConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreate adds the create to the network connect params
func (o *NetworkConnectParams) WithCreate(create *models.SwagCompatNetworkConnectRequest) *NetworkConnectParams {
	o.SetCreate(create)
	return o
}

// SetCreate adds the create to the network connect params
func (o *NetworkConnectParams) SetCreate(create *models.SwagCompatNetworkConnectRequest) {
	o.Create = create
}

// WithName adds the name to the network connect params
func (o *NetworkConnectParams) WithName(name string) *NetworkConnectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the network connect params
func (o *NetworkConnectParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Create != nil {
		if err := r.SetBodyParam(o.Create); err != nil {
			return err
		}
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