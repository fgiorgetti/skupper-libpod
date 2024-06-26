// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeCreateLibpodReader is a Reader for the VolumeCreateLibpod structure.
type VolumeCreateLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCreateLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVolumeCreateLibpodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVolumeCreateLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeCreateLibpodCreated creates a VolumeCreateLibpodCreated with default headers values
func NewVolumeCreateLibpodCreated() *VolumeCreateLibpodCreated {
	return &VolumeCreateLibpodCreated{}
}

/*
VolumeCreateLibpodCreated describes a response with status code 201, with default header values.

Volume create response
*/
type VolumeCreateLibpodCreated struct {
	Payload *VolumeCreateLibpodCreatedBody
}

// IsSuccess returns true when this volume create libpod created response has a 2xx status code
func (o *VolumeCreateLibpodCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume create libpod created response has a 3xx status code
func (o *VolumeCreateLibpodCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume create libpod created response has a 4xx status code
func (o *VolumeCreateLibpodCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume create libpod created response has a 5xx status code
func (o *VolumeCreateLibpodCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this volume create libpod created response a status code equal to that given
func (o *VolumeCreateLibpodCreated) IsCode(code int) bool {
	return code == 201
}

func (o *VolumeCreateLibpodCreated) Error() string {
	return fmt.Sprintf("[POST /libpod/volumes/create][%d] volumeCreateLibpodCreated  %+v", 201, o.Payload)
}

func (o *VolumeCreateLibpodCreated) String() string {
	return fmt.Sprintf("[POST /libpod/volumes/create][%d] volumeCreateLibpodCreated  %+v", 201, o.Payload)
}

func (o *VolumeCreateLibpodCreated) GetPayload() *VolumeCreateLibpodCreatedBody {
	return o.Payload
}

func (o *VolumeCreateLibpodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeCreateLibpodCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateLibpodInternalServerError creates a VolumeCreateLibpodInternalServerError with default headers values
func NewVolumeCreateLibpodInternalServerError() *VolumeCreateLibpodInternalServerError {
	return &VolumeCreateLibpodInternalServerError{}
}

/*
VolumeCreateLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeCreateLibpodInternalServerError struct {
	Payload *VolumeCreateLibpodInternalServerErrorBody
}

// IsSuccess returns true when this volume create libpod internal server error response has a 2xx status code
func (o *VolumeCreateLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume create libpod internal server error response has a 3xx status code
func (o *VolumeCreateLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume create libpod internal server error response has a 4xx status code
func (o *VolumeCreateLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume create libpod internal server error response has a 5xx status code
func (o *VolumeCreateLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume create libpod internal server error response a status code equal to that given
func (o *VolumeCreateLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *VolumeCreateLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/volumes/create][%d] volumeCreateLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeCreateLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/volumes/create][%d] volumeCreateLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeCreateLibpodInternalServerError) GetPayload() *VolumeCreateLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeCreateLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeCreateLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VolumeCreateLibpodCreatedBody volume create libpod created body
swagger:model VolumeCreateLibpodCreatedBody
*/
type VolumeCreateLibpodCreatedBody struct {

	// Anonymous indicates that the volume was created as an anonymous
	// volume for a specific container, and will be be removed when any
	// container using it is removed.
	Anonymous bool `json:"Anonymous,omitempty"`

	// CreatedAt is the date and time the volume was created at. This is not
	// stored for older Libpod volumes; if so, it will be omitted.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"CreatedAt,omitempty"`

	// Driver is the driver used to create the volume.
	// If set to "local" or "", the Local driver (Podman built-in code) is
	// used to service the volume; otherwise, a volume plugin with the given
	// name is used to mount and manage the volume.
	Driver string `json:"Driver,omitempty"`

	// GID is the GID that the volume was created with.
	GID int64 `json:"GID,omitempty"`

	// Labels includes the volume's configured labels, key:value pairs that
	// can be passed during volume creation to provide information for third
	// party tools.
	Labels map[string]string `json:"Labels,omitempty"`

	// MountCount is the number of times this volume has been mounted.
	MountCount uint64 `json:"MountCount,omitempty"`

	// Mountpoint is the path on the host where the volume is mounted.
	Mountpoint string `json:"Mountpoint,omitempty"`

	// Name is the name of the volume.
	Name string `json:"Name,omitempty"`

	// NeedsChown indicates that the next time the volume is mounted into
	// a container, the container will chown the volume to the container process
	// UID/GID.
	NeedsChown bool `json:"NeedsChown,omitempty"`

	// NeedsCopyUp indicates that the next time the volume is mounted into
	NeedsCopyUp bool `json:"NeedsCopyUp,omitempty"`

	// Options is a set of options that were used when creating the volume.
	// For the Local driver, these are mount options that will be used to
	// determine how a local filesystem is mounted; they are handled as
	// parameters to Mount in a manner described in the volume create
	// manpage.
	// For non-local drivers, these are passed as-is to the volume plugin.
	Options map[string]string `json:"Options,omitempty"`

	// Scope is unused and provided solely for Docker compatibility. It is
	// unconditionally set to "local".
	Scope string `json:"Scope,omitempty"`

	// Status is used to return information on the volume's current state,
	// if the volume was created using a volume plugin (uses a Driver that
	// is not the local driver).
	// Status is provided to us by an external program, so no guarantees are
	// made about its format or contents. Further, it is an optional field,
	// so it may not be set even in cases where a volume plugin is in use.
	Status map[string]interface{} `json:"Status,omitempty"`

	// UID is the UID that the volume was created with.
	UID int64 `json:"UID,omitempty"`
}

// Validate validates this volume create libpod created body
func (o *VolumeCreateLibpodCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeCreateLibpodCreatedBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("volumeCreateLibpodCreated"+"."+"CreatedAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volume create libpod created body based on context it is used
func (o *VolumeCreateLibpodCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeCreateLibpodCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeCreateLibpodCreatedBody) UnmarshalBinary(b []byte) error {
	var res VolumeCreateLibpodCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VolumeCreateLibpodInternalServerErrorBody volume create libpod internal server error body
swagger:model VolumeCreateLibpodInternalServerErrorBody
*/
type VolumeCreateLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume create libpod internal server error body
func (o *VolumeCreateLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume create libpod internal server error body based on context it is used
func (o *VolumeCreateLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeCreateLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeCreateLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeCreateLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
