// Code generated by go-swagger; DO NOT EDIT.

package volumes_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fgiorgetti/skupper-libpod/pkg/libpod/models"
)

// VolumeListReader is a Reader for the VolumeList structure.
type VolumeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVolumeListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeListOK creates a VolumeListOK with default headers values
func NewVolumeListOK() *VolumeListOK {
	return &VolumeListOK{}
}

/*
VolumeListOK describes a response with status code 200, with default header values.

Volume list response
*/
type VolumeListOK struct {
	Payload *VolumeListOKBody
}

// IsSuccess returns true when this volume list o k response has a 2xx status code
func (o *VolumeListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume list o k response has a 3xx status code
func (o *VolumeListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume list o k response has a 4xx status code
func (o *VolumeListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume list o k response has a 5xx status code
func (o *VolumeListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume list o k response a status code equal to that given
func (o *VolumeListOK) IsCode(code int) bool {
	return code == 200
}

func (o *VolumeListOK) Error() string {
	return fmt.Sprintf("[GET /volumes][%d] volumeListOK  %+v", 200, o.Payload)
}

func (o *VolumeListOK) String() string {
	return fmt.Sprintf("[GET /volumes][%d] volumeListOK  %+v", 200, o.Payload)
}

func (o *VolumeListOK) GetPayload() *VolumeListOKBody {
	return o.Payload
}

func (o *VolumeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeListInternalServerError creates a VolumeListInternalServerError with default headers values
func NewVolumeListInternalServerError() *VolumeListInternalServerError {
	return &VolumeListInternalServerError{}
}

/*
VolumeListInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeListInternalServerError struct {
	Payload *VolumeListInternalServerErrorBody
}

// IsSuccess returns true when this volume list internal server error response has a 2xx status code
func (o *VolumeListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume list internal server error response has a 3xx status code
func (o *VolumeListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume list internal server error response has a 4xx status code
func (o *VolumeListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume list internal server error response has a 5xx status code
func (o *VolumeListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume list internal server error response a status code equal to that given
func (o *VolumeListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *VolumeListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumes][%d] volumeListInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeListInternalServerError) String() string {
	return fmt.Sprintf("[GET /volumes][%d] volumeListInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeListInternalServerError) GetPayload() *VolumeListInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeListInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VolumeListInternalServerErrorBody volume list internal server error body
swagger:model VolumeListInternalServerErrorBody
*/
type VolumeListInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume list internal server error body
func (o *VolumeListInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume list internal server error body based on context it is used
func (o *VolumeListInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeListInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeListInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeListInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VolumeListOKBody volume list o k body
swagger:model VolumeListOKBody
*/
type VolumeListOKBody struct {

	// volumes
	Volumes []*models.VolumeListOKBody `json:"Volumes"`
}

// Validate validates this volume list o k body
func (o *VolumeListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeListOKBody) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(o.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(o.Volumes); i++ {
		if swag.IsZero(o.Volumes[i]) { // not required
			continue
		}

		if o.Volumes[i] != nil {
			if err := o.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeListOK" + "." + "Volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeListOK" + "." + "Volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this volume list o k body based on the context it is used
func (o *VolumeListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeListOKBody) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Volumes); i++ {

		if o.Volumes[i] != nil {
			if err := o.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeListOK" + "." + "Volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumeListOK" + "." + "Volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *VolumeListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeListOKBody) UnmarshalBinary(b []byte) error {
	var res VolumeListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}