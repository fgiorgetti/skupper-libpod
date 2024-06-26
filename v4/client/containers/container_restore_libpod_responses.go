// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// ContainerRestoreLibpodReader is a Reader for the ContainerRestoreLibpod structure.
type ContainerRestoreLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRestoreLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerRestoreLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerRestoreLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerRestoreLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerRestoreLibpodOK creates a ContainerRestoreLibpodOK with default headers values
func NewContainerRestoreLibpodOK() *ContainerRestoreLibpodOK {
	return &ContainerRestoreLibpodOK{}
}

/*
ContainerRestoreLibpodOK describes a response with status code 200, with default header values.

tarball is returned in body if exported
*/
type ContainerRestoreLibpodOK struct {
}

// IsSuccess returns true when this container restore libpod o k response has a 2xx status code
func (o *ContainerRestoreLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container restore libpod o k response has a 3xx status code
func (o *ContainerRestoreLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container restore libpod o k response has a 4xx status code
func (o *ContainerRestoreLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container restore libpod o k response has a 5xx status code
func (o *ContainerRestoreLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container restore libpod o k response a status code equal to that given
func (o *ContainerRestoreLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerRestoreLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/restore][%d] containerRestoreLibpodOK ", 200)
}

func (o *ContainerRestoreLibpodOK) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/restore][%d] containerRestoreLibpodOK ", 200)
}

func (o *ContainerRestoreLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRestoreLibpodNotFound creates a ContainerRestoreLibpodNotFound with default headers values
func NewContainerRestoreLibpodNotFound() *ContainerRestoreLibpodNotFound {
	return &ContainerRestoreLibpodNotFound{}
}

/*
ContainerRestoreLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerRestoreLibpodNotFound struct {
	Payload *ContainerRestoreLibpodNotFoundBody
}

// IsSuccess returns true when this container restore libpod not found response has a 2xx status code
func (o *ContainerRestoreLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container restore libpod not found response has a 3xx status code
func (o *ContainerRestoreLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container restore libpod not found response has a 4xx status code
func (o *ContainerRestoreLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container restore libpod not found response has a 5xx status code
func (o *ContainerRestoreLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container restore libpod not found response a status code equal to that given
func (o *ContainerRestoreLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerRestoreLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/restore][%d] containerRestoreLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerRestoreLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/restore][%d] containerRestoreLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerRestoreLibpodNotFound) GetPayload() *ContainerRestoreLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerRestoreLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRestoreLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRestoreLibpodInternalServerError creates a ContainerRestoreLibpodInternalServerError with default headers values
func NewContainerRestoreLibpodInternalServerError() *ContainerRestoreLibpodInternalServerError {
	return &ContainerRestoreLibpodInternalServerError{}
}

/*
ContainerRestoreLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerRestoreLibpodInternalServerError struct {
	Payload *ContainerRestoreLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container restore libpod internal server error response has a 2xx status code
func (o *ContainerRestoreLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container restore libpod internal server error response has a 3xx status code
func (o *ContainerRestoreLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container restore libpod internal server error response has a 4xx status code
func (o *ContainerRestoreLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container restore libpod internal server error response has a 5xx status code
func (o *ContainerRestoreLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container restore libpod internal server error response a status code equal to that given
func (o *ContainerRestoreLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerRestoreLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/restore][%d] containerRestoreLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerRestoreLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/restore][%d] containerRestoreLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerRestoreLibpodInternalServerError) GetPayload() *ContainerRestoreLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerRestoreLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRestoreLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerRestoreLibpodInternalServerErrorBody container restore libpod internal server error body
swagger:model ContainerRestoreLibpodInternalServerErrorBody
*/
type ContainerRestoreLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container restore libpod internal server error body
func (o *ContainerRestoreLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container restore libpod internal server error body based on context it is used
func (o *ContainerRestoreLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRestoreLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRestoreLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerRestoreLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerRestoreLibpodNotFoundBody container restore libpod not found body
swagger:model ContainerRestoreLibpodNotFoundBody
*/
type ContainerRestoreLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container restore libpod not found body
func (o *ContainerRestoreLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container restore libpod not found body based on context it is used
func (o *ContainerRestoreLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRestoreLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRestoreLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerRestoreLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
