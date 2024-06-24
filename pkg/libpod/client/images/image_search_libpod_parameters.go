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
	"github.com/go-openapi/swag"
)

// NewImageSearchLibpodParams creates a new ImageSearchLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageSearchLibpodParams() *ImageSearchLibpodParams {
	return &ImageSearchLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageSearchLibpodParamsWithTimeout creates a new ImageSearchLibpodParams object
// with the ability to set a timeout on a request.
func NewImageSearchLibpodParamsWithTimeout(timeout time.Duration) *ImageSearchLibpodParams {
	return &ImageSearchLibpodParams{
		timeout: timeout,
	}
}

// NewImageSearchLibpodParamsWithContext creates a new ImageSearchLibpodParams object
// with the ability to set a context for a request.
func NewImageSearchLibpodParamsWithContext(ctx context.Context) *ImageSearchLibpodParams {
	return &ImageSearchLibpodParams{
		Context: ctx,
	}
}

// NewImageSearchLibpodParamsWithHTTPClient creates a new ImageSearchLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageSearchLibpodParamsWithHTTPClient(client *http.Client) *ImageSearchLibpodParams {
	return &ImageSearchLibpodParams{
		HTTPClient: client,
	}
}

/*
ImageSearchLibpodParams contains all the parameters to send to the API endpoint

	for the image search libpod operation.

	Typically these are written to a http.Request.
*/
type ImageSearchLibpodParams struct {

	/* Filters.

	     A JSON encoded value of the filters (a `map[string][]string`) to process on the images list. Available filters:
	- `is-automated=(true|false)`
	- `is-official=(true|false)`
	- `stars=<number>` Matches images that has at least 'number' stars.

	*/
	Filters *string

	/* Limit.

	   maximum number of results

	   Default: 25
	*/
	Limit *int64

	/* ListTags.

	   list the available tags in the repository
	*/
	ListTags *bool

	/* Term.

	   term to search
	*/
	Term *string

	/* TLSVerify.

	   skip TLS verification for registries
	*/
	TLSVerify *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image search libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageSearchLibpodParams) WithDefaults() *ImageSearchLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image search libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageSearchLibpodParams) SetDefaults() {
	var (
		limitDefault = int64(25)

		listTagsDefault = bool(false)

		tLSVerifyDefault = bool(false)
	)

	val := ImageSearchLibpodParams{
		Limit:     &limitDefault,
		ListTags:  &listTagsDefault,
		TLSVerify: &tLSVerifyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the image search libpod params
func (o *ImageSearchLibpodParams) WithTimeout(timeout time.Duration) *ImageSearchLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image search libpod params
func (o *ImageSearchLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image search libpod params
func (o *ImageSearchLibpodParams) WithContext(ctx context.Context) *ImageSearchLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image search libpod params
func (o *ImageSearchLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image search libpod params
func (o *ImageSearchLibpodParams) WithHTTPClient(client *http.Client) *ImageSearchLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image search libpod params
func (o *ImageSearchLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the image search libpod params
func (o *ImageSearchLibpodParams) WithFilters(filters *string) *ImageSearchLibpodParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the image search libpod params
func (o *ImageSearchLibpodParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the image search libpod params
func (o *ImageSearchLibpodParams) WithLimit(limit *int64) *ImageSearchLibpodParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the image search libpod params
func (o *ImageSearchLibpodParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithListTags adds the listTags to the image search libpod params
func (o *ImageSearchLibpodParams) WithListTags(listTags *bool) *ImageSearchLibpodParams {
	o.SetListTags(listTags)
	return o
}

// SetListTags adds the listTags to the image search libpod params
func (o *ImageSearchLibpodParams) SetListTags(listTags *bool) {
	o.ListTags = listTags
}

// WithTerm adds the term to the image search libpod params
func (o *ImageSearchLibpodParams) WithTerm(term *string) *ImageSearchLibpodParams {
	o.SetTerm(term)
	return o
}

// SetTerm adds the term to the image search libpod params
func (o *ImageSearchLibpodParams) SetTerm(term *string) {
	o.Term = term
}

// WithTLSVerify adds the tLSVerify to the image search libpod params
func (o *ImageSearchLibpodParams) WithTLSVerify(tLSVerify *bool) *ImageSearchLibpodParams {
	o.SetTLSVerify(tLSVerify)
	return o
}

// SetTLSVerify adds the tlsVerify to the image search libpod params
func (o *ImageSearchLibpodParams) SetTLSVerify(tLSVerify *bool) {
	o.TLSVerify = tLSVerify
}

// WriteToRequest writes these params to a swagger request
func (o *ImageSearchLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.ListTags != nil {

		// query param listTags
		var qrListTags bool

		if o.ListTags != nil {
			qrListTags = *o.ListTags
		}
		qListTags := swag.FormatBool(qrListTags)
		if qListTags != "" {

			if err := r.SetQueryParam("listTags", qListTags); err != nil {
				return err
			}
		}
	}

	if o.Term != nil {

		// query param term
		var qrTerm string

		if o.Term != nil {
			qrTerm = *o.Term
		}
		qTerm := qrTerm
		if qTerm != "" {

			if err := r.SetQueryParam("term", qTerm); err != nil {
				return err
			}
		}
	}

	if o.TLSVerify != nil {

		// query param tlsVerify
		var qrTLSVerify bool

		if o.TLSVerify != nil {
			qrTLSVerify = *o.TLSVerify
		}
		qTLSVerify := swag.FormatBool(qrTLSVerify)
		if qTLSVerify != "" {

			if err := r.SetQueryParam("tlsVerify", qTLSVerify); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
