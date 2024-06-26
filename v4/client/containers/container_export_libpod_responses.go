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

// ContainerExportLibpodReader is a Reader for the ContainerExportLibpod structure.
type ContainerExportLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerExportLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerExportLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerExportLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerExportLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerExportLibpodOK creates a ContainerExportLibpodOK with default headers values
func NewContainerExportLibpodOK() *ContainerExportLibpodOK {
	return &ContainerExportLibpodOK{}
}

/*
ContainerExportLibpodOK describes a response with status code 200, with default header values.

tarball is returned in body
*/
type ContainerExportLibpodOK struct {
}

// IsSuccess returns true when this container export libpod o k response has a 2xx status code
func (o *ContainerExportLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container export libpod o k response has a 3xx status code
func (o *ContainerExportLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container export libpod o k response has a 4xx status code
func (o *ContainerExportLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container export libpod o k response has a 5xx status code
func (o *ContainerExportLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container export libpod o k response a status code equal to that given
func (o *ContainerExportLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerExportLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/export][%d] containerExportLibpodOK ", 200)
}

func (o *ContainerExportLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/export][%d] containerExportLibpodOK ", 200)
}

func (o *ContainerExportLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerExportLibpodNotFound creates a ContainerExportLibpodNotFound with default headers values
func NewContainerExportLibpodNotFound() *ContainerExportLibpodNotFound {
	return &ContainerExportLibpodNotFound{}
}

/*
ContainerExportLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerExportLibpodNotFound struct {
	Payload *ContainerExportLibpodNotFoundBody
}

// IsSuccess returns true when this container export libpod not found response has a 2xx status code
func (o *ContainerExportLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container export libpod not found response has a 3xx status code
func (o *ContainerExportLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container export libpod not found response has a 4xx status code
func (o *ContainerExportLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container export libpod not found response has a 5xx status code
func (o *ContainerExportLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container export libpod not found response a status code equal to that given
func (o *ContainerExportLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerExportLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/export][%d] containerExportLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerExportLibpodNotFound) String() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/export][%d] containerExportLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerExportLibpodNotFound) GetPayload() *ContainerExportLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerExportLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerExportLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExportLibpodInternalServerError creates a ContainerExportLibpodInternalServerError with default headers values
func NewContainerExportLibpodInternalServerError() *ContainerExportLibpodInternalServerError {
	return &ContainerExportLibpodInternalServerError{}
}

/*
ContainerExportLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerExportLibpodInternalServerError struct {
	Payload *ContainerExportLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container export libpod internal server error response has a 2xx status code
func (o *ContainerExportLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container export libpod internal server error response has a 3xx status code
func (o *ContainerExportLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container export libpod internal server error response has a 4xx status code
func (o *ContainerExportLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container export libpod internal server error response has a 5xx status code
func (o *ContainerExportLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container export libpod internal server error response a status code equal to that given
func (o *ContainerExportLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerExportLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/export][%d] containerExportLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerExportLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/export][%d] containerExportLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerExportLibpodInternalServerError) GetPayload() *ContainerExportLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerExportLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerExportLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerExportLibpodInternalServerErrorBody container export libpod internal server error body
swagger:model ContainerExportLibpodInternalServerErrorBody
*/
type ContainerExportLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container export libpod internal server error body
func (o *ContainerExportLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container export libpod internal server error body based on context it is used
func (o *ContainerExportLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExportLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExportLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerExportLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerExportLibpodNotFoundBody container export libpod not found body
swagger:model ContainerExportLibpodNotFoundBody
*/
type ContainerExportLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container export libpod not found body
func (o *ContainerExportLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container export libpod not found body based on context it is used
func (o *ContainerExportLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExportLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExportLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerExportLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
