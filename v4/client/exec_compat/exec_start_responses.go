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

// ExecStartReader is a Reader for the ExecStart structure.
type ExecStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewExecStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewExecStartConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecStartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecStartOK creates a ExecStartOK with default headers values
func NewExecStartOK() *ExecStartOK {
	return &ExecStartOK{}
}

/*
ExecStartOK describes a response with status code 200, with default header values.

no error
*/
type ExecStartOK struct {
}

// IsSuccess returns true when this exec start o k response has a 2xx status code
func (o *ExecStartOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this exec start o k response has a 3xx status code
func (o *ExecStartOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec start o k response has a 4xx status code
func (o *ExecStartOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this exec start o k response has a 5xx status code
func (o *ExecStartOK) IsServerError() bool {
	return false
}

// IsCode returns true when this exec start o k response a status code equal to that given
func (o *ExecStartOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExecStartOK) Error() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartOK ", 200)
}

func (o *ExecStartOK) String() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartOK ", 200)
}

func (o *ExecStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecStartNotFound creates a ExecStartNotFound with default headers values
func NewExecStartNotFound() *ExecStartNotFound {
	return &ExecStartNotFound{}
}

/*
ExecStartNotFound describes a response with status code 404, with default header values.

No such exec instance
*/
type ExecStartNotFound struct {
	Payload *ExecStartNotFoundBody
}

// IsSuccess returns true when this exec start not found response has a 2xx status code
func (o *ExecStartNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exec start not found response has a 3xx status code
func (o *ExecStartNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec start not found response has a 4xx status code
func (o *ExecStartNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this exec start not found response has a 5xx status code
func (o *ExecStartNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this exec start not found response a status code equal to that given
func (o *ExecStartNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ExecStartNotFound) Error() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartNotFound  %+v", 404, o.Payload)
}

func (o *ExecStartNotFound) String() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartNotFound  %+v", 404, o.Payload)
}

func (o *ExecStartNotFound) GetPayload() *ExecStartNotFoundBody {
	return o.Payload
}

func (o *ExecStartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExecStartNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecStartConflict creates a ExecStartConflict with default headers values
func NewExecStartConflict() *ExecStartConflict {
	return &ExecStartConflict{}
}

/*
ExecStartConflict describes a response with status code 409, with default header values.

container is not running
*/
type ExecStartConflict struct {
}

// IsSuccess returns true when this exec start conflict response has a 2xx status code
func (o *ExecStartConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exec start conflict response has a 3xx status code
func (o *ExecStartConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec start conflict response has a 4xx status code
func (o *ExecStartConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this exec start conflict response has a 5xx status code
func (o *ExecStartConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this exec start conflict response a status code equal to that given
func (o *ExecStartConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ExecStartConflict) Error() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartConflict ", 409)
}

func (o *ExecStartConflict) String() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartConflict ", 409)
}

func (o *ExecStartConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecStartInternalServerError creates a ExecStartInternalServerError with default headers values
func NewExecStartInternalServerError() *ExecStartInternalServerError {
	return &ExecStartInternalServerError{}
}

/*
ExecStartInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ExecStartInternalServerError struct {
	Payload *ExecStartInternalServerErrorBody
}

// IsSuccess returns true when this exec start internal server error response has a 2xx status code
func (o *ExecStartInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this exec start internal server error response has a 3xx status code
func (o *ExecStartInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this exec start internal server error response has a 4xx status code
func (o *ExecStartInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this exec start internal server error response has a 5xx status code
func (o *ExecStartInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this exec start internal server error response a status code equal to that given
func (o *ExecStartInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ExecStartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecStartInternalServerError) String() string {
	return fmt.Sprintf("[POST /exec/{id}/start][%d] execStartInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecStartInternalServerError) GetPayload() *ExecStartInternalServerErrorBody {
	return o.Payload
}

func (o *ExecStartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExecStartInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ExecStartBody exec start body
swagger:model ExecStartBody
*/
type ExecStartBody struct {

	// Detach from the command. Not presently supported.
	Detach bool `json:"Detach,omitempty"`

	// Allocate a pseudo-TTY. Presently ignored.
	Tty bool `json:"Tty,omitempty"`
}

// Validate validates this exec start body
func (o *ExecStartBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exec start body based on context it is used
func (o *ExecStartBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecStartBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecStartBody) UnmarshalBinary(b []byte) error {
	var res ExecStartBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExecStartInternalServerErrorBody exec start internal server error body
swagger:model ExecStartInternalServerErrorBody
*/
type ExecStartInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this exec start internal server error body
func (o *ExecStartInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exec start internal server error body based on context it is used
func (o *ExecStartInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecStartInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecStartInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ExecStartInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ExecStartNotFoundBody exec start not found body
swagger:model ExecStartNotFoundBody
*/
type ExecStartNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this exec start not found body
func (o *ExecStartNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exec start not found body based on context it is used
func (o *ExecStartNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ExecStartNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExecStartNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ExecStartNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
