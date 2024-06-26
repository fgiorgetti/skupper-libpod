// Code generated by go-swagger; DO NOT EDIT.

package volumes_compat

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

// NewVolumeCreateParams creates a new VolumeCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeCreateParams() *VolumeCreateParams {
	return &VolumeCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeCreateParamsWithTimeout creates a new VolumeCreateParams object
// with the ability to set a timeout on a request.
func NewVolumeCreateParamsWithTimeout(timeout time.Duration) *VolumeCreateParams {
	return &VolumeCreateParams{
		timeout: timeout,
	}
}

// NewVolumeCreateParamsWithContext creates a new VolumeCreateParams object
// with the ability to set a context for a request.
func NewVolumeCreateParamsWithContext(ctx context.Context) *VolumeCreateParams {
	return &VolumeCreateParams{
		Context: ctx,
	}
}

// NewVolumeCreateParamsWithHTTPClient creates a new VolumeCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeCreateParamsWithHTTPClient(client *http.Client) *VolumeCreateParams {
	return &VolumeCreateParams{
		HTTPClient: client,
	}
}

/*
VolumeCreateParams contains all the parameters to send to the API endpoint

	for the volume create operation.

	Typically these are written to a http.Request.
*/
type VolumeCreateParams struct {

	/* Create.

	     attributes for creating a volume.
	Note: If a volume by the same name exists, a 201 response with that volume's information will be generated.

	*/
	Create *models.DockerVolumeCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeCreateParams) WithDefaults() *VolumeCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume create params
func (o *VolumeCreateParams) WithTimeout(timeout time.Duration) *VolumeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume create params
func (o *VolumeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume create params
func (o *VolumeCreateParams) WithContext(ctx context.Context) *VolumeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume create params
func (o *VolumeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume create params
func (o *VolumeCreateParams) WithHTTPClient(client *http.Client) *VolumeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume create params
func (o *VolumeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreate adds the create to the volume create params
func (o *VolumeCreateParams) WithCreate(create *models.DockerVolumeCreate) *VolumeCreateParams {
	o.SetCreate(create)
	return o
}

// SetCreate adds the create to the volume create params
func (o *VolumeCreateParams) SetCreate(create *models.DockerVolumeCreate) {
	o.Create = create
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Create != nil {
		if err := r.SetBodyParam(o.Create); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
