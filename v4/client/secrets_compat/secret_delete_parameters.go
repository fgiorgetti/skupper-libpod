// Code generated by go-swagger; DO NOT EDIT.

package secrets_compat

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

// NewSecretDeleteParams creates a new SecretDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretDeleteParams() *SecretDeleteParams {
	return &SecretDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretDeleteParamsWithTimeout creates a new SecretDeleteParams object
// with the ability to set a timeout on a request.
func NewSecretDeleteParamsWithTimeout(timeout time.Duration) *SecretDeleteParams {
	return &SecretDeleteParams{
		timeout: timeout,
	}
}

// NewSecretDeleteParamsWithContext creates a new SecretDeleteParams object
// with the ability to set a context for a request.
func NewSecretDeleteParamsWithContext(ctx context.Context) *SecretDeleteParams {
	return &SecretDeleteParams{
		Context: ctx,
	}
}

// NewSecretDeleteParamsWithHTTPClient creates a new SecretDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretDeleteParamsWithHTTPClient(client *http.Client) *SecretDeleteParams {
	return &SecretDeleteParams{
		HTTPClient: client,
	}
}

/*
SecretDeleteParams contains all the parameters to send to the API endpoint

	for the secret delete operation.

	Typically these are written to a http.Request.
*/
type SecretDeleteParams struct {

	/* Name.

	   the name or ID of the secret
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretDeleteParams) WithDefaults() *SecretDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret delete params
func (o *SecretDeleteParams) WithTimeout(timeout time.Duration) *SecretDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret delete params
func (o *SecretDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret delete params
func (o *SecretDeleteParams) WithContext(ctx context.Context) *SecretDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret delete params
func (o *SecretDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret delete params
func (o *SecretDeleteParams) WithHTTPClient(client *http.Client) *SecretDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret delete params
func (o *SecretDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the secret delete params
func (o *SecretDeleteParams) WithName(name string) *SecretDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the secret delete params
func (o *SecretDeleteParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SecretDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
