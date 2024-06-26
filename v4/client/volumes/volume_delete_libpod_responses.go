// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeDeleteLibpodReader is a Reader for the VolumeDeleteLibpod structure.
type VolumeDeleteLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeDeleteLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVolumeDeleteLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVolumeDeleteLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewVolumeDeleteLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeDeleteLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeDeleteLibpodNoContent creates a VolumeDeleteLibpodNoContent with default headers values
func NewVolumeDeleteLibpodNoContent() *VolumeDeleteLibpodNoContent {
	return &VolumeDeleteLibpodNoContent{}
}

/*
VolumeDeleteLibpodNoContent describes a response with status code 204, with default header values.

no error
*/
type VolumeDeleteLibpodNoContent struct {
}

// IsSuccess returns true when this volume delete libpod no content response has a 2xx status code
func (o *VolumeDeleteLibpodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume delete libpod no content response has a 3xx status code
func (o *VolumeDeleteLibpodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume delete libpod no content response has a 4xx status code
func (o *VolumeDeleteLibpodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume delete libpod no content response has a 5xx status code
func (o *VolumeDeleteLibpodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this volume delete libpod no content response a status code equal to that given
func (o *VolumeDeleteLibpodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *VolumeDeleteLibpodNoContent) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodNoContent ", 204)
}

func (o *VolumeDeleteLibpodNoContent) String() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodNoContent ", 204)
}

func (o *VolumeDeleteLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVolumeDeleteLibpodNotFound creates a VolumeDeleteLibpodNotFound with default headers values
func NewVolumeDeleteLibpodNotFound() *VolumeDeleteLibpodNotFound {
	return &VolumeDeleteLibpodNotFound{}
}

/*
VolumeDeleteLibpodNotFound describes a response with status code 404, with default header values.

No such volume
*/
type VolumeDeleteLibpodNotFound struct {
	Payload *VolumeDeleteLibpodNotFoundBody
}

// IsSuccess returns true when this volume delete libpod not found response has a 2xx status code
func (o *VolumeDeleteLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume delete libpod not found response has a 3xx status code
func (o *VolumeDeleteLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume delete libpod not found response has a 4xx status code
func (o *VolumeDeleteLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume delete libpod not found response has a 5xx status code
func (o *VolumeDeleteLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume delete libpod not found response a status code equal to that given
func (o *VolumeDeleteLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *VolumeDeleteLibpodNotFound) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodNotFound  %+v", 404, o.Payload)
}

func (o *VolumeDeleteLibpodNotFound) String() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodNotFound  %+v", 404, o.Payload)
}

func (o *VolumeDeleteLibpodNotFound) GetPayload() *VolumeDeleteLibpodNotFoundBody {
	return o.Payload
}

func (o *VolumeDeleteLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeDeleteLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeDeleteLibpodConflict creates a VolumeDeleteLibpodConflict with default headers values
func NewVolumeDeleteLibpodConflict() *VolumeDeleteLibpodConflict {
	return &VolumeDeleteLibpodConflict{}
}

/*
VolumeDeleteLibpodConflict describes a response with status code 409, with default header values.

Volume is in use and cannot be removed
*/
type VolumeDeleteLibpodConflict struct {
}

// IsSuccess returns true when this volume delete libpod conflict response has a 2xx status code
func (o *VolumeDeleteLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume delete libpod conflict response has a 3xx status code
func (o *VolumeDeleteLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume delete libpod conflict response has a 4xx status code
func (o *VolumeDeleteLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume delete libpod conflict response has a 5xx status code
func (o *VolumeDeleteLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this volume delete libpod conflict response a status code equal to that given
func (o *VolumeDeleteLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *VolumeDeleteLibpodConflict) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodConflict ", 409)
}

func (o *VolumeDeleteLibpodConflict) String() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodConflict ", 409)
}

func (o *VolumeDeleteLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVolumeDeleteLibpodInternalServerError creates a VolumeDeleteLibpodInternalServerError with default headers values
func NewVolumeDeleteLibpodInternalServerError() *VolumeDeleteLibpodInternalServerError {
	return &VolumeDeleteLibpodInternalServerError{}
}

/*
VolumeDeleteLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeDeleteLibpodInternalServerError struct {
	Payload *VolumeDeleteLibpodInternalServerErrorBody
}

// IsSuccess returns true when this volume delete libpod internal server error response has a 2xx status code
func (o *VolumeDeleteLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume delete libpod internal server error response has a 3xx status code
func (o *VolumeDeleteLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume delete libpod internal server error response has a 4xx status code
func (o *VolumeDeleteLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume delete libpod internal server error response has a 5xx status code
func (o *VolumeDeleteLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume delete libpod internal server error response a status code equal to that given
func (o *VolumeDeleteLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *VolumeDeleteLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeDeleteLibpodInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /libpod/volumes/{name}][%d] volumeDeleteLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeDeleteLibpodInternalServerError) GetPayload() *VolumeDeleteLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeDeleteLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeDeleteLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VolumeDeleteLibpodInternalServerErrorBody volume delete libpod internal server error body
swagger:model VolumeDeleteLibpodInternalServerErrorBody
*/
type VolumeDeleteLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume delete libpod internal server error body
func (o *VolumeDeleteLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume delete libpod internal server error body based on context it is used
func (o *VolumeDeleteLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeDeleteLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeDeleteLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeDeleteLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VolumeDeleteLibpodNotFoundBody volume delete libpod not found body
swagger:model VolumeDeleteLibpodNotFoundBody
*/
type VolumeDeleteLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume delete libpod not found body
func (o *VolumeDeleteLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume delete libpod not found body based on context it is used
func (o *VolumeDeleteLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeDeleteLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeDeleteLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res VolumeDeleteLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
