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

// ContainerKillLibpodReader is a Reader for the ContainerKillLibpod structure.
type ContainerKillLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerKillLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerKillLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerKillLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerKillLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerKillLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerKillLibpodNoContent creates a ContainerKillLibpodNoContent with default headers values
func NewContainerKillLibpodNoContent() *ContainerKillLibpodNoContent {
	return &ContainerKillLibpodNoContent{}
}

/*
ContainerKillLibpodNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerKillLibpodNoContent struct {
}

// IsSuccess returns true when this container kill libpod no content response has a 2xx status code
func (o *ContainerKillLibpodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container kill libpod no content response has a 3xx status code
func (o *ContainerKillLibpodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container kill libpod no content response has a 4xx status code
func (o *ContainerKillLibpodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container kill libpod no content response has a 5xx status code
func (o *ContainerKillLibpodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container kill libpod no content response a status code equal to that given
func (o *ContainerKillLibpodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerKillLibpodNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodNoContent ", 204)
}

func (o *ContainerKillLibpodNoContent) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodNoContent ", 204)
}

func (o *ContainerKillLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerKillLibpodNotFound creates a ContainerKillLibpodNotFound with default headers values
func NewContainerKillLibpodNotFound() *ContainerKillLibpodNotFound {
	return &ContainerKillLibpodNotFound{}
}

/*
ContainerKillLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerKillLibpodNotFound struct {
	Payload *ContainerKillLibpodNotFoundBody
}

// IsSuccess returns true when this container kill libpod not found response has a 2xx status code
func (o *ContainerKillLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container kill libpod not found response has a 3xx status code
func (o *ContainerKillLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container kill libpod not found response has a 4xx status code
func (o *ContainerKillLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container kill libpod not found response has a 5xx status code
func (o *ContainerKillLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container kill libpod not found response a status code equal to that given
func (o *ContainerKillLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerKillLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerKillLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerKillLibpodNotFound) GetPayload() *ContainerKillLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerKillLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerKillLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerKillLibpodConflict creates a ContainerKillLibpodConflict with default headers values
func NewContainerKillLibpodConflict() *ContainerKillLibpodConflict {
	return &ContainerKillLibpodConflict{}
}

/*
ContainerKillLibpodConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ContainerKillLibpodConflict struct {
	Payload *ContainerKillLibpodConflictBody
}

// IsSuccess returns true when this container kill libpod conflict response has a 2xx status code
func (o *ContainerKillLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container kill libpod conflict response has a 3xx status code
func (o *ContainerKillLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container kill libpod conflict response has a 4xx status code
func (o *ContainerKillLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this container kill libpod conflict response has a 5xx status code
func (o *ContainerKillLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this container kill libpod conflict response a status code equal to that given
func (o *ContainerKillLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ContainerKillLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodConflict  %+v", 409, o.Payload)
}

func (o *ContainerKillLibpodConflict) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodConflict  %+v", 409, o.Payload)
}

func (o *ContainerKillLibpodConflict) GetPayload() *ContainerKillLibpodConflictBody {
	return o.Payload
}

func (o *ContainerKillLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerKillLibpodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerKillLibpodInternalServerError creates a ContainerKillLibpodInternalServerError with default headers values
func NewContainerKillLibpodInternalServerError() *ContainerKillLibpodInternalServerError {
	return &ContainerKillLibpodInternalServerError{}
}

/*
ContainerKillLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerKillLibpodInternalServerError struct {
	Payload *ContainerKillLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container kill libpod internal server error response has a 2xx status code
func (o *ContainerKillLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container kill libpod internal server error response has a 3xx status code
func (o *ContainerKillLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container kill libpod internal server error response has a 4xx status code
func (o *ContainerKillLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container kill libpod internal server error response has a 5xx status code
func (o *ContainerKillLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container kill libpod internal server error response a status code equal to that given
func (o *ContainerKillLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerKillLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerKillLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/kill][%d] containerKillLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerKillLibpodInternalServerError) GetPayload() *ContainerKillLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerKillLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerKillLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerKillLibpodConflictBody container kill libpod conflict body
swagger:model ContainerKillLibpodConflictBody
*/
type ContainerKillLibpodConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container kill libpod conflict body
func (o *ContainerKillLibpodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container kill libpod conflict body based on context it is used
func (o *ContainerKillLibpodConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerKillLibpodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerKillLibpodConflictBody) UnmarshalBinary(b []byte) error {
	var res ContainerKillLibpodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerKillLibpodInternalServerErrorBody container kill libpod internal server error body
swagger:model ContainerKillLibpodInternalServerErrorBody
*/
type ContainerKillLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container kill libpod internal server error body
func (o *ContainerKillLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container kill libpod internal server error body based on context it is used
func (o *ContainerKillLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerKillLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerKillLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerKillLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerKillLibpodNotFoundBody container kill libpod not found body
swagger:model ContainerKillLibpodNotFoundBody
*/
type ContainerKillLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container kill libpod not found body
func (o *ContainerKillLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container kill libpod not found body based on context it is used
func (o *ContainerKillLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerKillLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerKillLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerKillLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
