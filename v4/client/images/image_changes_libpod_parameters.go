// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewImageChangesLibpodParams creates a new ImageChangesLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageChangesLibpodParams() *ImageChangesLibpodParams {
	return &ImageChangesLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageChangesLibpodParamsWithTimeout creates a new ImageChangesLibpodParams object
// with the ability to set a timeout on a request.
func NewImageChangesLibpodParamsWithTimeout(timeout time.Duration) *ImageChangesLibpodParams {
	return &ImageChangesLibpodParams{
		timeout: timeout,
	}
}

// NewImageChangesLibpodParamsWithContext creates a new ImageChangesLibpodParams object
// with the ability to set a context for a request.
func NewImageChangesLibpodParamsWithContext(ctx context.Context) *ImageChangesLibpodParams {
	return &ImageChangesLibpodParams{
		Context: ctx,
	}
}

// NewImageChangesLibpodParamsWithHTTPClient creates a new ImageChangesLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageChangesLibpodParamsWithHTTPClient(client *http.Client) *ImageChangesLibpodParams {
	return &ImageChangesLibpodParams{
		HTTPClient: client,
	}
}

/*
ImageChangesLibpodParams contains all the parameters to send to the API endpoint

	for the image changes libpod operation.

	Typically these are written to a http.Request.
*/
type ImageChangesLibpodParams struct {

	/* DiffType.

	   select what you want to match, default is all
	*/
	DiffType *string

	/* Name.

	   the name or id of the image
	*/
	Name string

	/* Parent.

	   specify a second layer which is used to compare against it instead of the parent layer
	*/
	Parent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image changes libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageChangesLibpodParams) WithDefaults() *ImageChangesLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image changes libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageChangesLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image changes libpod params
func (o *ImageChangesLibpodParams) WithTimeout(timeout time.Duration) *ImageChangesLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image changes libpod params
func (o *ImageChangesLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image changes libpod params
func (o *ImageChangesLibpodParams) WithContext(ctx context.Context) *ImageChangesLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image changes libpod params
func (o *ImageChangesLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image changes libpod params
func (o *ImageChangesLibpodParams) WithHTTPClient(client *http.Client) *ImageChangesLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image changes libpod params
func (o *ImageChangesLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiffType adds the diffType to the image changes libpod params
func (o *ImageChangesLibpodParams) WithDiffType(diffType *string) *ImageChangesLibpodParams {
	o.SetDiffType(diffType)
	return o
}

// SetDiffType adds the diffType to the image changes libpod params
func (o *ImageChangesLibpodParams) SetDiffType(diffType *string) {
	o.DiffType = diffType
}

// WithName adds the name to the image changes libpod params
func (o *ImageChangesLibpodParams) WithName(name string) *ImageChangesLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image changes libpod params
func (o *ImageChangesLibpodParams) SetName(name string) {
	o.Name = name
}

// WithParent adds the parent to the image changes libpod params
func (o *ImageChangesLibpodParams) WithParent(parent *string) *ImageChangesLibpodParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the image changes libpod params
func (o *ImageChangesLibpodParams) SetParent(parent *string) {
	o.Parent = parent
}

// WriteToRequest writes these params to a swagger request
func (o *ImageChangesLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DiffType != nil {

		// query param diffType
		var qrDiffType string

		if o.DiffType != nil {
			qrDiffType = *o.DiffType
		}
		qDiffType := qrDiffType
		if qDiffType != "" {

			if err := r.SetQueryParam("diffType", qDiffType); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Parent != nil {

		// query param parent
		var qrParent string

		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {

			if err := r.SetQueryParam("parent", qParent); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
