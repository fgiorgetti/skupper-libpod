// Code generated by go-swagger; DO NOT EDIT.

package secrets_compat

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

// SecretDeleteReader is a Reader for the SecretDelete structure.
type SecretDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSecretDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSecretDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretDeleteNoContent creates a SecretDeleteNoContent with default headers values
func NewSecretDeleteNoContent() *SecretDeleteNoContent {
	return &SecretDeleteNoContent{}
}

/*
SecretDeleteNoContent describes a response with status code 204, with default header values.

no error
*/
type SecretDeleteNoContent struct {
}

// IsSuccess returns true when this secret delete no content response has a 2xx status code
func (o *SecretDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secret delete no content response has a 3xx status code
func (o *SecretDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret delete no content response has a 4xx status code
func (o *SecretDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret delete no content response has a 5xx status code
func (o *SecretDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this secret delete no content response a status code equal to that given
func (o *SecretDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *SecretDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{name}][%d] secretDeleteNoContent ", 204)
}

func (o *SecretDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /secrets/{name}][%d] secretDeleteNoContent ", 204)
}

func (o *SecretDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecretDeleteNotFound creates a SecretDeleteNotFound with default headers values
func NewSecretDeleteNotFound() *SecretDeleteNotFound {
	return &SecretDeleteNotFound{}
}

/*
SecretDeleteNotFound describes a response with status code 404, with default header values.

No such secret
*/
type SecretDeleteNotFound struct {
	Payload *SecretDeleteNotFoundBody
}

// IsSuccess returns true when this secret delete not found response has a 2xx status code
func (o *SecretDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret delete not found response has a 3xx status code
func (o *SecretDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret delete not found response has a 4xx status code
func (o *SecretDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this secret delete not found response has a 5xx status code
func (o *SecretDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this secret delete not found response a status code equal to that given
func (o *SecretDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SecretDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{name}][%d] secretDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SecretDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /secrets/{name}][%d] secretDeleteNotFound  %+v", 404, o.Payload)
}

func (o *SecretDeleteNotFound) GetPayload() *SecretDeleteNotFoundBody {
	return o.Payload
}

func (o *SecretDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretDeleteNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretDeleteInternalServerError creates a SecretDeleteInternalServerError with default headers values
func NewSecretDeleteInternalServerError() *SecretDeleteInternalServerError {
	return &SecretDeleteInternalServerError{}
}

/*
SecretDeleteInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretDeleteInternalServerError struct {
	Payload *SecretDeleteInternalServerErrorBody
}

// IsSuccess returns true when this secret delete internal server error response has a 2xx status code
func (o *SecretDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret delete internal server error response has a 3xx status code
func (o *SecretDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret delete internal server error response has a 4xx status code
func (o *SecretDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret delete internal server error response has a 5xx status code
func (o *SecretDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secret delete internal server error response a status code equal to that given
func (o *SecretDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SecretDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /secrets/{name}][%d] secretDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /secrets/{name}][%d] secretDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretDeleteInternalServerError) GetPayload() *SecretDeleteInternalServerErrorBody {
	return o.Payload
}

func (o *SecretDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretDeleteInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecretDeleteInternalServerErrorBody secret delete internal server error body
swagger:model SecretDeleteInternalServerErrorBody
*/
type SecretDeleteInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret delete internal server error body
func (o *SecretDeleteInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret delete internal server error body based on context it is used
func (o *SecretDeleteInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretDeleteInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretDeleteInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SecretDeleteInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecretDeleteNotFoundBody secret delete not found body
swagger:model SecretDeleteNotFoundBody
*/
type SecretDeleteNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret delete not found body
func (o *SecretDeleteNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret delete not found body based on context it is used
func (o *SecretDeleteNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretDeleteNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretDeleteNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SecretDeleteNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
