// Code generated by go-swagger; DO NOT EDIT.

package networks

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

	"github.com/fgiorgetti/skupper-libpod/v4/models"
)

// NewNetworkConnectLibpodParams creates a new NetworkConnectLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkConnectLibpodParams() *NetworkConnectLibpodParams {
	return &NetworkConnectLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkConnectLibpodParamsWithTimeout creates a new NetworkConnectLibpodParams object
// with the ability to set a timeout on a request.
func NewNetworkConnectLibpodParamsWithTimeout(timeout time.Duration) *NetworkConnectLibpodParams {
	return &NetworkConnectLibpodParams{
		timeout: timeout,
	}
}

// NewNetworkConnectLibpodParamsWithContext creates a new NetworkConnectLibpodParams object
// with the ability to set a context for a request.
func NewNetworkConnectLibpodParamsWithContext(ctx context.Context) *NetworkConnectLibpodParams {
	return &NetworkConnectLibpodParams{
		Context: ctx,
	}
}

// NewNetworkConnectLibpodParamsWithHTTPClient creates a new NetworkConnectLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkConnectLibpodParamsWithHTTPClient(client *http.Client) *NetworkConnectLibpodParams {
	return &NetworkConnectLibpodParams{
		HTTPClient: client,
	}
}

/*
NetworkConnectLibpodParams contains all the parameters to send to the API endpoint

	for the network connect libpod operation.

	Typically these are written to a http.Request.
*/
type NetworkConnectLibpodParams struct {

	/* Create.

	   attributes for connecting a container to a network
	*/
	Create *models.SwagNetworkConnectRequest

	/* Name.

	   the name of the network
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network connect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkConnectLibpodParams) WithDefaults() *NetworkConnectLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network connect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkConnectLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network connect libpod params
func (o *NetworkConnectLibpodParams) WithTimeout(timeout time.Duration) *NetworkConnectLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network connect libpod params
func (o *NetworkConnectLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network connect libpod params
func (o *NetworkConnectLibpodParams) WithContext(ctx context.Context) *NetworkConnectLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network connect libpod params
func (o *NetworkConnectLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network connect libpod params
func (o *NetworkConnectLibpodParams) WithHTTPClient(client *http.Client) *NetworkConnectLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network connect libpod params
func (o *NetworkConnectLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreate adds the create to the network connect libpod params
func (o *NetworkConnectLibpodParams) WithCreate(create *models.SwagNetworkConnectRequest) *NetworkConnectLibpodParams {
	o.SetCreate(create)
	return o
}

// SetCreate adds the create to the network connect libpod params
func (o *NetworkConnectLibpodParams) SetCreate(create *models.SwagNetworkConnectRequest) {
	o.Create = create
}

// WithName adds the name to the network connect libpod params
func (o *NetworkConnectLibpodParams) WithName(name string) *NetworkConnectLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the network connect libpod params
func (o *NetworkConnectLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkConnectLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
