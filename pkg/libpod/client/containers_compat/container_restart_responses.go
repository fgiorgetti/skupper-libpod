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

// ContainerRestartReader is a Reader for the ContainerRestart structure.
type ContainerRestartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRestartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerRestartNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerRestartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerRestartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerRestartNoContent creates a ContainerRestartNoContent with default headers values
func NewContainerRestartNoContent() *ContainerRestartNoContent {
	return &ContainerRestartNoContent{}
}

/*
ContainerRestartNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerRestartNoContent struct {
}

// IsSuccess returns true when this container restart no content response has a 2xx status code
func (o *ContainerRestartNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container restart no content response has a 3xx status code
func (o *ContainerRestartNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container restart no content response has a 4xx status code
func (o *ContainerRestartNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container restart no content response has a 5xx status code
func (o *ContainerRestartNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container restart no content response a status code equal to that given
func (o *ContainerRestartNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerRestartNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/restart][%d] containerRestartNoContent ", 204)
}

func (o *ContainerRestartNoContent) String() string {
	return fmt.Sprintf("[POST /containers/{name}/restart][%d] containerRestartNoContent ", 204)
}

func (o *ContainerRestartNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRestartNotFound creates a ContainerRestartNotFound with default headers values
func NewContainerRestartNotFound() *ContainerRestartNotFound {
	return &ContainerRestartNotFound{}
}

/*
ContainerRestartNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerRestartNotFound struct {
	Payload *ContainerRestartNotFoundBody
}

// IsSuccess returns true when this container restart not found response has a 2xx status code
func (o *ContainerRestartNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container restart not found response has a 3xx status code
func (o *ContainerRestartNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container restart not found response has a 4xx status code
func (o *ContainerRestartNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container restart not found response has a 5xx status code
func (o *ContainerRestartNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container restart not found response a status code equal to that given
func (o *ContainerRestartNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerRestartNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/restart][%d] containerRestartNotFound  %+v", 404, o.Payload)
}

func (o *ContainerRestartNotFound) String() string {
	return fmt.Sprintf("[POST /containers/{name}/restart][%d] containerRestartNotFound  %+v", 404, o.Payload)
}

func (o *ContainerRestartNotFound) GetPayload() *ContainerRestartNotFoundBody {
	return o.Payload
}

func (o *ContainerRestartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRestartNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRestartInternalServerError creates a ContainerRestartInternalServerError with default headers values
func NewContainerRestartInternalServerError() *ContainerRestartInternalServerError {
	return &ContainerRestartInternalServerError{}
}

/*
ContainerRestartInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerRestartInternalServerError struct {
	Payload *ContainerRestartInternalServerErrorBody
}

// IsSuccess returns true when this container restart internal server error response has a 2xx status code
func (o *ContainerRestartInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container restart internal server error response has a 3xx status code
func (o *ContainerRestartInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container restart internal server error response has a 4xx status code
func (o *ContainerRestartInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container restart internal server error response has a 5xx status code
func (o *ContainerRestartInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container restart internal server error response a status code equal to that given
func (o *ContainerRestartInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerRestartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/restart][%d] containerRestartInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerRestartInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/{name}/restart][%d] containerRestartInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerRestartInternalServerError) GetPayload() *ContainerRestartInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerRestartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRestartInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerRestartInternalServerErrorBody container restart internal server error body
swagger:model ContainerRestartInternalServerErrorBody
*/
type ContainerRestartInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container restart internal server error body
func (o *ContainerRestartInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container restart internal server error body based on context it is used
func (o *ContainerRestartInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRestartInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRestartInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerRestartInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerRestartNotFoundBody container restart not found body
swagger:model ContainerRestartNotFoundBody
*/
type ContainerRestartNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container restart not found body
func (o *ContainerRestartNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container restart not found body based on context it is used
func (o *ContainerRestartNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRestartNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRestartNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerRestartNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}