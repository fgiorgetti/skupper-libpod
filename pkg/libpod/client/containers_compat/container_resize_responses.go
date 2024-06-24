// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// ContainerResizeReader is a Reader for the ContainerResize structure.
type ContainerResizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerResizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerResizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerResizeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerResizeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerResizeOK creates a ContainerResizeOK with default headers values
func NewContainerResizeOK() *ContainerResizeOK {
	return &ContainerResizeOK{}
}

/*
ContainerResizeOK describes a response with status code 200, with default header values.

Success
*/
type ContainerResizeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this container resize o k response has a 2xx status code
func (o *ContainerResizeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container resize o k response has a 3xx status code
func (o *ContainerResizeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container resize o k response has a 4xx status code
func (o *ContainerResizeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container resize o k response has a 5xx status code
func (o *ContainerResizeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container resize o k response a status code equal to that given
func (o *ContainerResizeOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerResizeOK) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/resize][%d] containerResizeOK  %+v", 200, o.Payload)
}

func (o *ContainerResizeOK) String() string {
	return fmt.Sprintf("[POST /containers/{name}/resize][%d] containerResizeOK  %+v", 200, o.Payload)
}

func (o *ContainerResizeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ContainerResizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerResizeNotFound creates a ContainerResizeNotFound with default headers values
func NewContainerResizeNotFound() *ContainerResizeNotFound {
	return &ContainerResizeNotFound{}
}

/*
ContainerResizeNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerResizeNotFound struct {
	Payload *ContainerResizeNotFoundBody
}

// IsSuccess returns true when this container resize not found response has a 2xx status code
func (o *ContainerResizeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container resize not found response has a 3xx status code
func (o *ContainerResizeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container resize not found response has a 4xx status code
func (o *ContainerResizeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container resize not found response has a 5xx status code
func (o *ContainerResizeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container resize not found response a status code equal to that given
func (o *ContainerResizeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerResizeNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/resize][%d] containerResizeNotFound  %+v", 404, o.Payload)
}

func (o *ContainerResizeNotFound) String() string {
	return fmt.Sprintf("[POST /containers/{name}/resize][%d] containerResizeNotFound  %+v", 404, o.Payload)
}

func (o *ContainerResizeNotFound) GetPayload() *ContainerResizeNotFoundBody {
	return o.Payload
}

func (o *ContainerResizeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerResizeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerResizeInternalServerError creates a ContainerResizeInternalServerError with default headers values
func NewContainerResizeInternalServerError() *ContainerResizeInternalServerError {
	return &ContainerResizeInternalServerError{}
}

/*
ContainerResizeInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerResizeInternalServerError struct {
	Payload *ContainerResizeInternalServerErrorBody
}

// IsSuccess returns true when this container resize internal server error response has a 2xx status code
func (o *ContainerResizeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container resize internal server error response has a 3xx status code
func (o *ContainerResizeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container resize internal server error response has a 4xx status code
func (o *ContainerResizeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container resize internal server error response has a 5xx status code
func (o *ContainerResizeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container resize internal server error response a status code equal to that given
func (o *ContainerResizeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerResizeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/resize][%d] containerResizeInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerResizeInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/{name}/resize][%d] containerResizeInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerResizeInternalServerError) GetPayload() *ContainerResizeInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerResizeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerResizeInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerResizeInternalServerErrorBody container resize internal server error body
swagger:model ContainerResizeInternalServerErrorBody
*/
type ContainerResizeInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container resize internal server error body
func (o *ContainerResizeInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container resize internal server error body based on context it is used
func (o *ContainerResizeInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerResizeInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerResizeInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerResizeInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerResizeNotFoundBody container resize not found body
swagger:model ContainerResizeNotFoundBody
*/
type ContainerResizeNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container resize not found body
func (o *ContainerResizeNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container resize not found body based on context it is used
func (o *ContainerResizeNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerResizeNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerResizeNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerResizeNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
