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

// NewGenerateSystemdLibpodParams creates a new GenerateSystemdLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateSystemdLibpodParams() *GenerateSystemdLibpodParams {
	return &GenerateSystemdLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateSystemdLibpodParamsWithTimeout creates a new GenerateSystemdLibpodParams object
// with the ability to set a timeout on a request.
func NewGenerateSystemdLibpodParamsWithTimeout(timeout time.Duration) *GenerateSystemdLibpodParams {
	return &GenerateSystemdLibpodParams{
		timeout: timeout,
	}
}

// NewGenerateSystemdLibpodParamsWithContext creates a new GenerateSystemdLibpodParams object
// with the ability to set a context for a request.
func NewGenerateSystemdLibpodParamsWithContext(ctx context.Context) *GenerateSystemdLibpodParams {
	return &GenerateSystemdLibpodParams{
		Context: ctx,
	}
}

// NewGenerateSystemdLibpodParamsWithHTTPClient creates a new GenerateSystemdLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateSystemdLibpodParamsWithHTTPClient(client *http.Client) *GenerateSystemdLibpodParams {
	return &GenerateSystemdLibpodParams{
		HTTPClient: client,
	}
}

/*
GenerateSystemdLibpodParams contains all the parameters to send to the API endpoint

	for the generate systemd libpod operation.

	Typically these are written to a http.Request.
*/
type GenerateSystemdLibpodParams struct {

	/* After.

	   Systemd After list for the container or pods.
	*/
	After []string

	/* ContainerPrefix.

	   Systemd unit name prefix for containers.

	   Default: "container"
	*/
	ContainerPrefix *string

	/* Name.

	   Name or ID of the container or pod.
	*/
	Name string

	/* New.

	   Create a new container instead of starting an existing one.
	*/
	New *bool

	/* NoHeader.

	   Do not generate the header including the Podman version and the timestamp.
	*/
	NoHeader *bool

	/* PodPrefix.

	   Systemd unit name prefix for pods.

	   Default: "pod"
	*/
	PodPrefix *string

	/* Requires.

	   Systemd Requires list for the container or pods.
	*/
	Requires []string

	/* RestartPolicy.

	   Systemd restart-policy.

	   Default: "on-failure"
	*/
	RestartPolicy *string

	/* RestartSec.

	   Configures the time to sleep before restarting a service.
	*/
	RestartSec *int64

	/* Separator.

	   Systemd unit name separator between name/id and prefix.

	   Default: "-"
	*/
	Separator *string

	/* StartTimeout.

	   Start timeout in seconds.
	*/
	StartTimeout *int64

	/* StopTimeout.

	   Stop timeout in seconds.

	   Default: 10
	*/
	StopTimeout *int64

	/* UseName.

	   Use container/pod names instead of IDs.
	*/
	UseName *bool

	/* Wants.

	   Systemd Wants list for the container or pods.
	*/
	Wants []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate systemd libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateSystemdLibpodParams) WithDefaults() *GenerateSystemdLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate systemd libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateSystemdLibpodParams) SetDefaults() {
	var (
		afterDefault = []string{}

		containerPrefixDefault = string("container")

		newDefault = bool(false)

		noHeaderDefault = bool(false)

		podPrefixDefault = string("pod")

		requiresDefault = []string{}

		restartPolicyDefault = string("on-failure")

		restartSecDefault = int64(0)

		separatorDefault = string("-")

		startTimeoutDefault = int64(0)

		stopTimeoutDefault = int64(10)

		useNameDefault = bool(false)

		wantsDefault = []string{}
	)

	val := GenerateSystemdLibpodParams{
		After:           afterDefault,
		ContainerPrefix: &containerPrefixDefault,
		New:             &newDefault,
		NoHeader:        &noHeaderDefault,
		PodPrefix:       &podPrefixDefault,
		Requires:        requiresDefault,
		RestartPolicy:   &restartPolicyDefault,
		RestartSec:      &restartSecDefault,
		Separator:       &separatorDefault,
		StartTimeout:    &startTimeoutDefault,
		StopTimeout:     &stopTimeoutDefault,
		UseName:         &useNameDefault,
		Wants:           wantsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithTimeout(timeout time.Duration) *GenerateSystemdLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithContext(ctx context.Context) *GenerateSystemdLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithHTTPClient(client *http.Client) *GenerateSystemdLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithAfter(after []string) *GenerateSystemdLibpodParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetAfter(after []string) {
	o.After = after
}

// WithContainerPrefix adds the containerPrefix to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithContainerPrefix(containerPrefix *string) *GenerateSystemdLibpodParams {
	o.SetContainerPrefix(containerPrefix)
	return o
}

// SetContainerPrefix adds the containerPrefix to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetContainerPrefix(containerPrefix *string) {
	o.ContainerPrefix = containerPrefix
}

// WithName adds the name to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithName(name string) *GenerateSystemdLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetName(name string) {
	o.Name = name
}

// WithNew adds the new to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithNew(new *bool) *GenerateSystemdLibpodParams {
	o.SetNew(new)
	return o
}

// SetNew adds the new to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetNew(new *bool) {
	o.New = new
}

// WithNoHeader adds the noHeader to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithNoHeader(noHeader *bool) *GenerateSystemdLibpodParams {
	o.SetNoHeader(noHeader)
	return o
}

// SetNoHeader adds the noHeader to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetNoHeader(noHeader *bool) {
	o.NoHeader = noHeader
}

// WithPodPrefix adds the podPrefix to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithPodPrefix(podPrefix *string) *GenerateSystemdLibpodParams {
	o.SetPodPrefix(podPrefix)
	return o
}

// SetPodPrefix adds the podPrefix to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetPodPrefix(podPrefix *string) {
	o.PodPrefix = podPrefix
}

// WithRequires adds the requires to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithRequires(requires []string) *GenerateSystemdLibpodParams {
	o.SetRequires(requires)
	return o
}

// SetRequires adds the requires to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetRequires(requires []string) {
	o.Requires = requires
}

// WithRestartPolicy adds the restartPolicy to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithRestartPolicy(restartPolicy *string) *GenerateSystemdLibpodParams {
	o.SetRestartPolicy(restartPolicy)
	return o
}

// SetRestartPolicy adds the restartPolicy to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetRestartPolicy(restartPolicy *string) {
	o.RestartPolicy = restartPolicy
}

// WithRestartSec adds the restartSec to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithRestartSec(restartSec *int64) *GenerateSystemdLibpodParams {
	o.SetRestartSec(restartSec)
	return o
}

// SetRestartSec adds the restartSec to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetRestartSec(restartSec *int64) {
	o.RestartSec = restartSec
}

// WithSeparator adds the separator to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithSeparator(separator *string) *GenerateSystemdLibpodParams {
	o.SetSeparator(separator)
	return o
}

// SetSeparator adds the separator to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetSeparator(separator *string) {
	o.Separator = separator
}

// WithStartTimeout adds the startTimeout to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithStartTimeout(startTimeout *int64) *GenerateSystemdLibpodParams {
	o.SetStartTimeout(startTimeout)
	return o
}

// SetStartTimeout adds the startTimeout to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetStartTimeout(startTimeout *int64) {
	o.StartTimeout = startTimeout
}

// WithStopTimeout adds the stopTimeout to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithStopTimeout(stopTimeout *int64) *GenerateSystemdLibpodParams {
	o.SetStopTimeout(stopTimeout)
	return o
}

// SetStopTimeout adds the stopTimeout to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetStopTimeout(stopTimeout *int64) {
	o.StopTimeout = stopTimeout
}

// WithUseName adds the useName to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithUseName(useName *bool) *GenerateSystemdLibpodParams {
	o.SetUseName(useName)
	return o
}

// SetUseName adds the useName to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetUseName(useName *bool) {
	o.UseName = useName
}

// WithWants adds the wants to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) WithWants(wants []string) *GenerateSystemdLibpodParams {
	o.SetWants(wants)
	return o
}

// SetWants adds the wants to the generate systemd libpod params
func (o *GenerateSystemdLibpodParams) SetWants(wants []string) {
	o.Wants = wants
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateSystemdLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// binding items for after
		joinedAfter := o.bindParamAfter(reg)

		// query array param after
		if err := r.SetQueryParam("after", joinedAfter...); err != nil {
			return err
		}
	}

	if o.ContainerPrefix != nil {

		// query param containerPrefix
		var qrContainerPrefix string

		if o.ContainerPrefix != nil {
			qrContainerPrefix = *o.ContainerPrefix
		}
		qContainerPrefix := qrContainerPrefix
		if qContainerPrefix != "" {

			if err := r.SetQueryParam("containerPrefix", qContainerPrefix); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.New != nil {

		// query param new
		var qrNew bool

		if o.New != nil {
			qrNew = *o.New
		}
		qNew := swag.FormatBool(qrNew)
		if qNew != "" {

			if err := r.SetQueryParam("new", qNew); err != nil {
				return err
			}
		}
	}

	if o.NoHeader != nil {

		// query param noHeader
		var qrNoHeader bool

		if o.NoHeader != nil {
			qrNoHeader = *o.NoHeader
		}
		qNoHeader := swag.FormatBool(qrNoHeader)
		if qNoHeader != "" {

			if err := r.SetQueryParam("noHeader", qNoHeader); err != nil {
				return err
			}
		}
	}

	if o.PodPrefix != nil {

		// query param podPrefix
		var qrPodPrefix string

		if o.PodPrefix != nil {
			qrPodPrefix = *o.PodPrefix
		}
		qPodPrefix := qrPodPrefix
		if qPodPrefix != "" {

			if err := r.SetQueryParam("podPrefix", qPodPrefix); err != nil {
				return err
			}
		}
	}

	if o.Requires != nil {

		// binding items for requires
		joinedRequires := o.bindParamRequires(reg)

		// query array param requires
		if err := r.SetQueryParam("requires", joinedRequires...); err != nil {
			return err
		}
	}

	if o.RestartPolicy != nil {

		// query param restartPolicy
		var qrRestartPolicy string

		if o.RestartPolicy != nil {
			qrRestartPolicy = *o.RestartPolicy
		}
		qRestartPolicy := qrRestartPolicy
		if qRestartPolicy != "" {

			if err := r.SetQueryParam("restartPolicy", qRestartPolicy); err != nil {
				return err
			}
		}
	}

	if o.RestartSec != nil {

		// query param restartSec
		var qrRestartSec int64

		if o.RestartSec != nil {
			qrRestartSec = *o.RestartSec
		}
		qRestartSec := swag.FormatInt64(qrRestartSec)
		if qRestartSec != "" {

			if err := r.SetQueryParam("restartSec", qRestartSec); err != nil {
				return err
			}
		}
	}

	if o.Separator != nil {

		// query param separator
		var qrSeparator string

		if o.Separator != nil {
			qrSeparator = *o.Separator
		}
		qSeparator := qrSeparator
		if qSeparator != "" {

			if err := r.SetQueryParam("separator", qSeparator); err != nil {
				return err
			}
		}
	}

	if o.StartTimeout != nil {

		// query param startTimeout
		var qrStartTimeout int64

		if o.StartTimeout != nil {
			qrStartTimeout = *o.StartTimeout
		}
		qStartTimeout := swag.FormatInt64(qrStartTimeout)
		if qStartTimeout != "" {

			if err := r.SetQueryParam("startTimeout", qStartTimeout); err != nil {
				return err
			}
		}
	}

	if o.StopTimeout != nil {

		// query param stopTimeout
		var qrStopTimeout int64

		if o.StopTimeout != nil {
			qrStopTimeout = *o.StopTimeout
		}
		qStopTimeout := swag.FormatInt64(qrStopTimeout)
		if qStopTimeout != "" {

			if err := r.SetQueryParam("stopTimeout", qStopTimeout); err != nil {
				return err
			}
		}
	}

	if o.UseName != nil {

		// query param useName
		var qrUseName bool

		if o.UseName != nil {
			qrUseName = *o.UseName
		}
		qUseName := swag.FormatBool(qrUseName)
		if qUseName != "" {

			if err := r.SetQueryParam("useName", qUseName); err != nil {
				return err
			}
		}
	}

	if o.Wants != nil {

		// binding items for wants
		joinedWants := o.bindParamWants(reg)

		// query array param wants
		if err := r.SetQueryParam("wants", joinedWants...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGenerateSystemdLibpod binds the parameter after
func (o *GenerateSystemdLibpodParams) bindParamAfter(formats strfmt.Registry) []string {
	afterIR := o.After

	var afterIC []string
	for _, afterIIR := range afterIR { // explode []string

		afterIIV := afterIIR // string as string
		afterIC = append(afterIC, afterIIV)
	}

	// items.CollectionFormat: ""
	afterIS := swag.JoinByFormat(afterIC, "")

	return afterIS
}

// bindParamGenerateSystemdLibpod binds the parameter requires
func (o *GenerateSystemdLibpodParams) bindParamRequires(formats strfmt.Registry) []string {
	requiresIR := o.Requires

	var requiresIC []string
	for _, requiresIIR := range requiresIR { // explode []string

		requiresIIV := requiresIIR // string as string
		requiresIC = append(requiresIC, requiresIIV)
	}

	// items.CollectionFormat: ""
	requiresIS := swag.JoinByFormat(requiresIC, "")

	return requiresIS
}

// bindParamGenerateSystemdLibpod binds the parameter wants
func (o *GenerateSystemdLibpodParams) bindParamWants(formats strfmt.Registry) []string {
	wantsIR := o.Wants

	var wantsIC []string
	for _, wantsIIR := range wantsIR { // explode []string

		wantsIIV := wantsIIR // string as string
		wantsIC = append(wantsIC, wantsIIV)
	}

	// items.CollectionFormat: ""
	wantsIS := swag.JoinByFormat(wantsIC, "")

	return wantsIS
}
