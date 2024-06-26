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

// ContainerLogsLibpodReader is a Reader for the ContainerLogsLibpod structure.
type ContainerLogsLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerLogsLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerLogsLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerLogsLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerLogsLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerLogsLibpodOK creates a ContainerLogsLibpodOK with default headers values
func NewContainerLogsLibpodOK() *ContainerLogsLibpodOK {
	return &ContainerLogsLibpodOK{}
}

/*
ContainerLogsLibpodOK describes a response with status code 200, with default header values.

logs returned as a stream in response body.
*/
type ContainerLogsLibpodOK struct {
}

// IsSuccess returns true when this container logs libpod o k response has a 2xx status code
func (o *ContainerLogsLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container logs libpod o k response has a 3xx status code
func (o *ContainerLogsLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs libpod o k response has a 4xx status code
func (o *ContainerLogsLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container logs libpod o k response has a 5xx status code
func (o *ContainerLogsLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container logs libpod o k response a status code equal to that given
func (o *ContainerLogsLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerLogsLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] containerLogsLibpodOK ", 200)
}

func (o *ContainerLogsLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] containerLogsLibpodOK ", 200)
}

func (o *ContainerLogsLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerLogsLibpodNotFound creates a ContainerLogsLibpodNotFound with default headers values
func NewContainerLogsLibpodNotFound() *ContainerLogsLibpodNotFound {
	return &ContainerLogsLibpodNotFound{}
}

/*
ContainerLogsLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerLogsLibpodNotFound struct {
	Payload *ContainerLogsLibpodNotFoundBody
}

// IsSuccess returns true when this container logs libpod not found response has a 2xx status code
func (o *ContainerLogsLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container logs libpod not found response has a 3xx status code
func (o *ContainerLogsLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs libpod not found response has a 4xx status code
func (o *ContainerLogsLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container logs libpod not found response has a 5xx status code
func (o *ContainerLogsLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container logs libpod not found response a status code equal to that given
func (o *ContainerLogsLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerLogsLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] containerLogsLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerLogsLibpodNotFound) String() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] containerLogsLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerLogsLibpodNotFound) GetPayload() *ContainerLogsLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerLogsLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerLogsLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerLogsLibpodInternalServerError creates a ContainerLogsLibpodInternalServerError with default headers values
func NewContainerLogsLibpodInternalServerError() *ContainerLogsLibpodInternalServerError {
	return &ContainerLogsLibpodInternalServerError{}
}

/*
ContainerLogsLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerLogsLibpodInternalServerError struct {
	Payload *ContainerLogsLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container logs libpod internal server error response has a 2xx status code
func (o *ContainerLogsLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container logs libpod internal server error response has a 3xx status code
func (o *ContainerLogsLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs libpod internal server error response has a 4xx status code
func (o *ContainerLogsLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container logs libpod internal server error response has a 5xx status code
func (o *ContainerLogsLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container logs libpod internal server error response a status code equal to that given
func (o *ContainerLogsLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerLogsLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] containerLogsLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerLogsLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/logs][%d] containerLogsLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerLogsLibpodInternalServerError) GetPayload() *ContainerLogsLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerLogsLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerLogsLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerLogsLibpodInternalServerErrorBody container logs libpod internal server error body
swagger:model ContainerLogsLibpodInternalServerErrorBody
*/
type ContainerLogsLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container logs libpod internal server error body
func (o *ContainerLogsLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container logs libpod internal server error body based on context it is used
func (o *ContainerLogsLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerLogsLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerLogsLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerLogsLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerLogsLibpodNotFoundBody container logs libpod not found body
swagger:model ContainerLogsLibpodNotFoundBody
*/
type ContainerLogsLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container logs libpod not found body
func (o *ContainerLogsLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container logs libpod not found body based on context it is used
func (o *ContainerLogsLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerLogsLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerLogsLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerLogsLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
