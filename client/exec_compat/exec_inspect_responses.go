// Code generated by go-swagger; DO NOT EDIT.

package exec_compat

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

// ExecInspectReader is a Reader for the ExecInspect structure.
type ExecInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewExecInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecInspectOK creates a ExecInspectOK with default headers values
func NewExecInspectOK() *ExecInspectOK {
	return &ExecInspectOK{}
}

/*
ExecInspectOK describes a response with status code 200, with default header values.

no error
*/
type ExecInspectOK struct {
}

// IsSuccess returns true when this exec inspect o k response has a 2xx status code
func (o *ExecInspectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this exec inspect o k response has a 3xx status code
func (o *ExecInspectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec inspect o k response has a 4xx status code
func (o *ExecInspectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this exec inspect o k response has a 5xx status code
func (o *ExecInspectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this exec inspect o k response a status code equal to that given
func (o *ExecInspectOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExecInspectOK) Error() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectOK ", 200)
}

func (o *ExecInspectOK) String() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectOK ", 200)
}

func (o *ExecInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecInspectNotFound creates a ExecInspectNotFound with default headers values
func NewExecInspectNotFound() *ExecInspectNotFound {
	return &ExecInspectNotFound{}
}

/*
ExecInspectNotFound describes a response with status code 404, with default header values.

No such exec instance
*/
type ExecInspectNotFound struct {
	Payload *ExecInspectNotFoundBody
}

// IsSuccess returns true when this exec inspect not found response has a 2xx status code
func (o *ExecInspectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exec inspect not found response has a 3xx status code
func (o *ExecInspectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec inspect not found response has a 4xx status code
func (o *ExecInspectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this exec inspect not found response has a 5xx status code
func (o *ExecInspectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this exec inspect not found response a status code equal to that given
func (o *ExecInspectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ExecInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectNotFound  %+v", 404, o.Payload)
}

func (o *ExecInspectNotFound) String() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectNotFound  %+v", 404, o.Payload)
}

func (o *ExecInspectNotFound) GetPayload() *ExecInspectNotFoundBody {
	return o.Payload
}

func (o *ExecInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExecInspectNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecInspectInternalServerError creates a ExecInspectInternalServerError with default headers values
func NewExecInspectInternalServerError() *ExecInspectInternalServerError {
	return &ExecInspectInternalServerError{}
}

/*
ExecInspectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ExecInspectInternalServerError struct {
	Payload *ExecInspectInternalServerErrorBody
}

// IsSuccess returns true when this exec inspect internal server error response has a 2xx status code
func (o *ExecInspectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exec inspect internal server error response has a 3xx status code
func (o *ExecInspectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec inspect internal server error response has a 4xx status code
func (o *ExecInspectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this exec inspect internal server error response has a 5xx status code
func (o *ExecInspectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this exec inspect internal server error response a status code equal to that given
func (o *ExecInspectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ExecInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecInspectInternalServerError) String() string {
	return fmt.Sprintf("[GET /exec/{id}/json][%d] execInspectInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecInspectInternalServerError) GetPayload() *ExecInspectInternalServerErrorBody {
	return o.Payload
}

func (o *ExecInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExecInspectInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExecInspectInternalServerErrorBody exec inspect internal server error body
swagger:model ExecInspectInternalServerErrorBody
*/
type ExecInspectInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this exec inspect internal server error body
func (o *ExecInspectInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exec inspect internal server error body based on context it is used
func (o *ExecInspectInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecInspectInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecInspectInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ExecInspectInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExecInspectNotFoundBody exec inspect not found body
swagger:model ExecInspectNotFoundBody
*/
type ExecInspectNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this exec inspect not found body
func (o *ExecInspectNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exec inspect not found body based on context it is used
func (o *ExecInspectNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecInspectNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecInspectNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ExecInspectNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}