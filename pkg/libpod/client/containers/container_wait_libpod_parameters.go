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
	"github.com/go-openapi/swag"
)

// NewContainerWaitLibpodParams creates a new ContainerWaitLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerWaitLibpodParams() *ContainerWaitLibpodParams {
	return &ContainerWaitLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerWaitLibpodParamsWithTimeout creates a new ContainerWaitLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerWaitLibpodParamsWithTimeout(timeout time.Duration) *ContainerWaitLibpodParams {
	return &ContainerWaitLibpodParams{
		timeout: timeout,
	}
}

// NewContainerWaitLibpodParamsWithContext creates a new ContainerWaitLibpodParams object
// with the ability to set a context for a request.
func NewContainerWaitLibpodParamsWithContext(ctx context.Context) *ContainerWaitLibpodParams {
	return &ContainerWaitLibpodParams{
		Context: ctx,
	}
}

// NewContainerWaitLibpodParamsWithHTTPClient creates a new ContainerWaitLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerWaitLibpodParamsWithHTTPClient(client *http.Client) *ContainerWaitLibpodParams {
	return &ContainerWaitLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainerWaitLibpodParams contains all the parameters to send to the API endpoint

	for the container wait libpod operation.

	Typically these are written to a http.Request.
*/
type ContainerWaitLibpodParams struct {

	/* Condition.

	   Conditions to wait for. If no condition provided the 'exited' condition is assumed.
	*/
	Condition []string

	/* Interval.

	   Time Interval to wait before polling for completion.

	   Default: "250ms"
	*/
	Interval *string

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container wait libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerWaitLibpodParams) WithDefaults() *ContainerWaitLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container wait libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerWaitLibpodParams) SetDefaults() {
	var (
		intervalDefault = string("250ms")
	)

	val := ContainerWaitLibpodParams{
		Interval: &intervalDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container wait libpod params
func (o *ContainerWaitLibpodParams) WithTimeout(timeout time.Duration) *ContainerWaitLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container wait libpod params
func (o *ContainerWaitLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container wait libpod params
func (o *ContainerWaitLibpodParams) WithContext(ctx context.Context) *ContainerWaitLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container wait libpod params
func (o *ContainerWaitLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container wait libpod params
func (o *ContainerWaitLibpodParams) WithHTTPClient(client *http.Client) *ContainerWaitLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container wait libpod params
func (o *ContainerWaitLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCondition adds the condition to the container wait libpod params
func (o *ContainerWaitLibpodParams) WithCondition(condition []string) *ContainerWaitLibpodParams {
	o.SetCondition(condition)
	return o
}

// SetCondition adds the condition to the container wait libpod params
func (o *ContainerWaitLibpodParams) SetCondition(condition []string) {
	o.Condition = condition
}

// WithInterval adds the interval to the container wait libpod params
func (o *ContainerWaitLibpodParams) WithInterval(interval *string) *ContainerWaitLibpodParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the container wait libpod params
func (o *ContainerWaitLibpodParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithName adds the name to the container wait libpod params
func (o *ContainerWaitLibpodParams) WithName(name string) *ContainerWaitLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container wait libpod params
func (o *ContainerWaitLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerWaitLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Condition != nil {

		// binding items for condition
		joinedCondition := o.bindParamCondition(reg)

		// query array param condition
		if err := r.SetQueryParam("condition", joinedCondition...); err != nil {
			return err
		}
	}

	if o.Interval != nil {

		// query param interval
		var qrInterval string

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
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

// bindParamContainerWaitLibpod binds the parameter condition
func (o *ContainerWaitLibpodParams) bindParamCondition(formats strfmt.Registry) []string {
	conditionIR := o.Condition

	var conditionIC []string
	for _, conditionIIR := range conditionIR { // explode []string

		conditionIIV := conditionIIR // string as string
		conditionIC = append(conditionIC, conditionIIV)
	}

	// items.CollectionFormat: ""
	conditionIS := swag.JoinByFormat(conditionIC, "")

	return conditionIS
}
