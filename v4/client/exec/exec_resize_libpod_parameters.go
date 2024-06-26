// Code generated by go-swagger; DO NOT EDIT.

package exec

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
	"github.com/go-openapi/swag"
)

// NewExecResizeLibpodParams creates a new ExecResizeLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecResizeLibpodParams() *ExecResizeLibpodParams {
	return &ExecResizeLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecResizeLibpodParamsWithTimeout creates a new ExecResizeLibpodParams object
// with the ability to set a timeout on a request.
func NewExecResizeLibpodParamsWithTimeout(timeout time.Duration) *ExecResizeLibpodParams {
	return &ExecResizeLibpodParams{
		timeout: timeout,
	}
}

// NewExecResizeLibpodParamsWithContext creates a new ExecResizeLibpodParams object
// with the ability to set a context for a request.
func NewExecResizeLibpodParamsWithContext(ctx context.Context) *ExecResizeLibpodParams {
	return &ExecResizeLibpodParams{
		Context: ctx,
	}
}

// NewExecResizeLibpodParamsWithHTTPClient creates a new ExecResizeLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecResizeLibpodParamsWithHTTPClient(client *http.Client) *ExecResizeLibpodParams {
	return &ExecResizeLibpodParams{
		HTTPClient: client,
	}
}

/*
ExecResizeLibpodParams contains all the parameters to send to the API endpoint

	for the exec resize libpod operation.

	Typically these are written to a http.Request.
*/
type ExecResizeLibpodParams struct {

	/* H.

	   Height of the TTY session in characters
	*/
	H *int64

	/* ID.

	   Exec instance ID
	*/
	ID string

	/* W.

	   Width of the TTY session in characters
	*/
	W *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the exec resize libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecResizeLibpodParams) WithDefaults() *ExecResizeLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the exec resize libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecResizeLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the exec resize libpod params
func (o *ExecResizeLibpodParams) WithTimeout(timeout time.Duration) *ExecResizeLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exec resize libpod params
func (o *ExecResizeLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exec resize libpod params
func (o *ExecResizeLibpodParams) WithContext(ctx context.Context) *ExecResizeLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exec resize libpod params
func (o *ExecResizeLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exec resize libpod params
func (o *ExecResizeLibpodParams) WithHTTPClient(client *http.Client) *ExecResizeLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exec resize libpod params
func (o *ExecResizeLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithH adds the h to the exec resize libpod params
func (o *ExecResizeLibpodParams) WithH(h *int64) *ExecResizeLibpodParams {
	o.SetH(h)
	return o
}

// SetH adds the h to the exec resize libpod params
func (o *ExecResizeLibpodParams) SetH(h *int64) {
	o.H = h
}

// WithID adds the id to the exec resize libpod params
func (o *ExecResizeLibpodParams) WithID(id string) *ExecResizeLibpodParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the exec resize libpod params
func (o *ExecResizeLibpodParams) SetID(id string) {
	o.ID = id
}

// WithW adds the w to the exec resize libpod params
func (o *ExecResizeLibpodParams) WithW(w *int64) *ExecResizeLibpodParams {
	o.SetW(w)
	return o
}

// SetW adds the w to the exec resize libpod params
func (o *ExecResizeLibpodParams) SetW(w *int64) {
	o.W = w
}

// WriteToRequest writes these params to a swagger request
func (o *ExecResizeLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.H != nil {

		// query param h
		var qrH int64

		if o.H != nil {
			qrH = *o.H
		}
		qH := swag.FormatInt64(qrH)
		if qH != "" {

			if err := r.SetQueryParam("h", qH); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.W != nil {

		// query param w
		var qrW int64

		if o.W != nil {
			qrW = *o.W
		}
		qW := swag.FormatInt64(qrW)
		if qW != "" {

			if err := r.SetQueryParam("w", qW); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
