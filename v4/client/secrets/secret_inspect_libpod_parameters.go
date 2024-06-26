// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

// NewSecretInspectLibpodParams creates a new SecretInspectLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretInspectLibpodParams() *SecretInspectLibpodParams {
	return &SecretInspectLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretInspectLibpodParamsWithTimeout creates a new SecretInspectLibpodParams object
// with the ability to set a timeout on a request.
func NewSecretInspectLibpodParamsWithTimeout(timeout time.Duration) *SecretInspectLibpodParams {
	return &SecretInspectLibpodParams{
		timeout: timeout,
	}
}

// NewSecretInspectLibpodParamsWithContext creates a new SecretInspectLibpodParams object
// with the ability to set a context for a request.
func NewSecretInspectLibpodParamsWithContext(ctx context.Context) *SecretInspectLibpodParams {
	return &SecretInspectLibpodParams{
		Context: ctx,
	}
}

// NewSecretInspectLibpodParamsWithHTTPClient creates a new SecretInspectLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretInspectLibpodParamsWithHTTPClient(client *http.Client) *SecretInspectLibpodParams {
	return &SecretInspectLibpodParams{
		HTTPClient: client,
	}
}

/*
SecretInspectLibpodParams contains all the parameters to send to the API endpoint

	for the secret inspect libpod operation.

	Typically these are written to a http.Request.
*/
type SecretInspectLibpodParams struct {

	/* Name.

	   the name or ID of the secret
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretInspectLibpodParams) WithDefaults() *SecretInspectLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretInspectLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret inspect libpod params
func (o *SecretInspectLibpodParams) WithTimeout(timeout time.Duration) *SecretInspectLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret inspect libpod params
func (o *SecretInspectLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret inspect libpod params
func (o *SecretInspectLibpodParams) WithContext(ctx context.Context) *SecretInspectLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret inspect libpod params
func (o *SecretInspectLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret inspect libpod params
func (o *SecretInspectLibpodParams) WithHTTPClient(client *http.Client) *SecretInspectLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret inspect libpod params
func (o *SecretInspectLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the secret inspect libpod params
func (o *SecretInspectLibpodParams) WithName(name string) *SecretInspectLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the secret inspect libpod params
func (o *SecretInspectLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SecretInspectLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
