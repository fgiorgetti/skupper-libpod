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

// ContainerStopLibpodReader is a Reader for the ContainerStopLibpod structure.
type ContainerStopLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerStopLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerStopLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewContainerStopLibpodNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerStopLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerStopLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerStopLibpodNoContent creates a ContainerStopLibpodNoContent with default headers values
func NewContainerStopLibpodNoContent() *ContainerStopLibpodNoContent {
	return &ContainerStopLibpodNoContent{}
}

/*
ContainerStopLibpodNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerStopLibpodNoContent struct {
}

// IsSuccess returns true when this container stop libpod no content response has a 2xx status code
func (o *ContainerStopLibpodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container stop libpod no content response has a 3xx status code
func (o *ContainerStopLibpodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container stop libpod no content response has a 4xx status code
func (o *ContainerStopLibpodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container stop libpod no content response has a 5xx status code
func (o *ContainerStopLibpodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container stop libpod no content response a status code equal to that given
func (o *ContainerStopLibpodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerStopLibpodNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodNoContent ", 204)
}

func (o *ContainerStopLibpodNoContent) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodNoContent ", 204)
}

func (o *ContainerStopLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerStopLibpodNotModified creates a ContainerStopLibpodNotModified with default headers values
func NewContainerStopLibpodNotModified() *ContainerStopLibpodNotModified {
	return &ContainerStopLibpodNotModified{}
}

/*
ContainerStopLibpodNotModified describes a response with status code 304, with default header values.

Container already stopped
*/
type ContainerStopLibpodNotModified struct {
	Payload *ContainerStopLibpodNotModifiedBody
}

// IsSuccess returns true when this container stop libpod not modified response has a 2xx status code
func (o *ContainerStopLibpodNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container stop libpod not modified response has a 3xx status code
func (o *ContainerStopLibpodNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this container stop libpod not modified response has a 4xx status code
func (o *ContainerStopLibpodNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this container stop libpod not modified response has a 5xx status code
func (o *ContainerStopLibpodNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this container stop libpod not modified response a status code equal to that given
func (o *ContainerStopLibpodNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *ContainerStopLibpodNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodNotModified  %+v", 304, o.Payload)
}

func (o *ContainerStopLibpodNotModified) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodNotModified  %+v", 304, o.Payload)
}

func (o *ContainerStopLibpodNotModified) GetPayload() *ContainerStopLibpodNotModifiedBody {
	return o.Payload
}

func (o *ContainerStopLibpodNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStopLibpodNotModifiedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStopLibpodNotFound creates a ContainerStopLibpodNotFound with default headers values
func NewContainerStopLibpodNotFound() *ContainerStopLibpodNotFound {
	return &ContainerStopLibpodNotFound{}
}

/*
ContainerStopLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerStopLibpodNotFound struct {
	Payload *ContainerStopLibpodNotFoundBody
}

// IsSuccess returns true when this container stop libpod not found response has a 2xx status code
func (o *ContainerStopLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container stop libpod not found response has a 3xx status code
func (o *ContainerStopLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container stop libpod not found response has a 4xx status code
func (o *ContainerStopLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container stop libpod not found response has a 5xx status code
func (o *ContainerStopLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container stop libpod not found response a status code equal to that given
func (o *ContainerStopLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerStopLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerStopLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerStopLibpodNotFound) GetPayload() *ContainerStopLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerStopLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStopLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStopLibpodInternalServerError creates a ContainerStopLibpodInternalServerError with default headers values
func NewContainerStopLibpodInternalServerError() *ContainerStopLibpodInternalServerError {
	return &ContainerStopLibpodInternalServerError{}
}

/*
ContainerStopLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerStopLibpodInternalServerError struct {
	Payload *ContainerStopLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container stop libpod internal server error response has a 2xx status code
func (o *ContainerStopLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container stop libpod internal server error response has a 3xx status code
func (o *ContainerStopLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container stop libpod internal server error response has a 4xx status code
func (o *ContainerStopLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container stop libpod internal server error response has a 5xx status code
func (o *ContainerStopLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container stop libpod internal server error response a status code equal to that given
func (o *ContainerStopLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerStopLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerStopLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/stop][%d] containerStopLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerStopLibpodInternalServerError) GetPayload() *ContainerStopLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerStopLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStopLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerStopLibpodInternalServerErrorBody container stop libpod internal server error body
swagger:model ContainerStopLibpodInternalServerErrorBody
*/
type ContainerStopLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container stop libpod internal server error body
func (o *ContainerStopLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container stop libpod internal server error body based on context it is used
func (o *ContainerStopLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStopLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStopLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerStopLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerStopLibpodNotFoundBody container stop libpod not found body
swagger:model ContainerStopLibpodNotFoundBody
*/
type ContainerStopLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container stop libpod not found body
func (o *ContainerStopLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container stop libpod not found body based on context it is used
func (o *ContainerStopLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStopLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStopLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerStopLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerStopLibpodNotModifiedBody container stop libpod not modified body
swagger:model ContainerStopLibpodNotModifiedBody
*/
type ContainerStopLibpodNotModifiedBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container stop libpod not modified body
func (o *ContainerStopLibpodNotModifiedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container stop libpod not modified body based on context it is used
func (o *ContainerStopLibpodNotModifiedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStopLibpodNotModifiedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStopLibpodNotModifiedBody) UnmarshalBinary(b []byte) error {
	var res ContainerStopLibpodNotModifiedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
