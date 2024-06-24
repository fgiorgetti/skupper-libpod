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

// ContainerStartLibpodReader is a Reader for the ContainerStartLibpod structure.
type ContainerStartLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerStartLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerStartLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewContainerStartLibpodNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerStartLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerStartLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerStartLibpodNoContent creates a ContainerStartLibpodNoContent with default headers values
func NewContainerStartLibpodNoContent() *ContainerStartLibpodNoContent {
	return &ContainerStartLibpodNoContent{}
}

/*
ContainerStartLibpodNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerStartLibpodNoContent struct {
}

// IsSuccess returns true when this container start libpod no content response has a 2xx status code
func (o *ContainerStartLibpodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container start libpod no content response has a 3xx status code
func (o *ContainerStartLibpodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container start libpod no content response has a 4xx status code
func (o *ContainerStartLibpodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container start libpod no content response has a 5xx status code
func (o *ContainerStartLibpodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container start libpod no content response a status code equal to that given
func (o *ContainerStartLibpodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerStartLibpodNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodNoContent ", 204)
}

func (o *ContainerStartLibpodNoContent) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodNoContent ", 204)
}

func (o *ContainerStartLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerStartLibpodNotModified creates a ContainerStartLibpodNotModified with default headers values
func NewContainerStartLibpodNotModified() *ContainerStartLibpodNotModified {
	return &ContainerStartLibpodNotModified{}
}

/*
ContainerStartLibpodNotModified describes a response with status code 304, with default header values.

Container already started
*/
type ContainerStartLibpodNotModified struct {
	Payload *ContainerStartLibpodNotModifiedBody
}

// IsSuccess returns true when this container start libpod not modified response has a 2xx status code
func (o *ContainerStartLibpodNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container start libpod not modified response has a 3xx status code
func (o *ContainerStartLibpodNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this container start libpod not modified response has a 4xx status code
func (o *ContainerStartLibpodNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this container start libpod not modified response has a 5xx status code
func (o *ContainerStartLibpodNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this container start libpod not modified response a status code equal to that given
func (o *ContainerStartLibpodNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *ContainerStartLibpodNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodNotModified  %+v", 304, o.Payload)
}

func (o *ContainerStartLibpodNotModified) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodNotModified  %+v", 304, o.Payload)
}

func (o *ContainerStartLibpodNotModified) GetPayload() *ContainerStartLibpodNotModifiedBody {
	return o.Payload
}

func (o *ContainerStartLibpodNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStartLibpodNotModifiedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStartLibpodNotFound creates a ContainerStartLibpodNotFound with default headers values
func NewContainerStartLibpodNotFound() *ContainerStartLibpodNotFound {
	return &ContainerStartLibpodNotFound{}
}

/*
ContainerStartLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerStartLibpodNotFound struct {
	Payload *ContainerStartLibpodNotFoundBody
}

// IsSuccess returns true when this container start libpod not found response has a 2xx status code
func (o *ContainerStartLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container start libpod not found response has a 3xx status code
func (o *ContainerStartLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container start libpod not found response has a 4xx status code
func (o *ContainerStartLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container start libpod not found response has a 5xx status code
func (o *ContainerStartLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container start libpod not found response a status code equal to that given
func (o *ContainerStartLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerStartLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerStartLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerStartLibpodNotFound) GetPayload() *ContainerStartLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerStartLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStartLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStartLibpodInternalServerError creates a ContainerStartLibpodInternalServerError with default headers values
func NewContainerStartLibpodInternalServerError() *ContainerStartLibpodInternalServerError {
	return &ContainerStartLibpodInternalServerError{}
}

/*
ContainerStartLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerStartLibpodInternalServerError struct {
	Payload *ContainerStartLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container start libpod internal server error response has a 2xx status code
func (o *ContainerStartLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container start libpod internal server error response has a 3xx status code
func (o *ContainerStartLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container start libpod internal server error response has a 4xx status code
func (o *ContainerStartLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container start libpod internal server error response has a 5xx status code
func (o *ContainerStartLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container start libpod internal server error response a status code equal to that given
func (o *ContainerStartLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerStartLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerStartLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/start][%d] containerStartLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerStartLibpodInternalServerError) GetPayload() *ContainerStartLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerStartLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStartLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerStartLibpodInternalServerErrorBody container start libpod internal server error body
swagger:model ContainerStartLibpodInternalServerErrorBody
*/
type ContainerStartLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container start libpod internal server error body
func (o *ContainerStartLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container start libpod internal server error body based on context it is used
func (o *ContainerStartLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStartLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStartLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerStartLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerStartLibpodNotFoundBody container start libpod not found body
swagger:model ContainerStartLibpodNotFoundBody
*/
type ContainerStartLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container start libpod not found body
func (o *ContainerStartLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container start libpod not found body based on context it is used
func (o *ContainerStartLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStartLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStartLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerStartLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerStartLibpodNotModifiedBody container start libpod not modified body
swagger:model ContainerStartLibpodNotModifiedBody
*/
type ContainerStartLibpodNotModifiedBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container start libpod not modified body
func (o *ContainerStartLibpodNotModifiedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container start libpod not modified body based on context it is used
func (o *ContainerStartLibpodNotModifiedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStartLibpodNotModifiedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStartLibpodNotModifiedBody) UnmarshalBinary(b []byte) error {
	var res ContainerStartLibpodNotModifiedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}