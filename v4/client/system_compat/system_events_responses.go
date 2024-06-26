// Code generated by go-swagger; DO NOT EDIT.

package system_compat

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

// SystemEventsReader is a Reader for the SystemEvents structure.
type SystemEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemEventsOK creates a SystemEventsOK with default headers values
func NewSystemEventsOK() *SystemEventsOK {
	return &SystemEventsOK{}
}

/*
SystemEventsOK describes a response with status code 200, with default header values.

returns a string of json data describing an event
*/
type SystemEventsOK struct {
}

// IsSuccess returns true when this system events o k response has a 2xx status code
func (o *SystemEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system events o k response has a 3xx status code
func (o *SystemEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system events o k response has a 4xx status code
func (o *SystemEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system events o k response has a 5xx status code
func (o *SystemEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system events o k response a status code equal to that given
func (o *SystemEventsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] systemEventsOK ", 200)
}

func (o *SystemEventsOK) String() string {
	return fmt.Sprintf("[GET /events][%d] systemEventsOK ", 200)
}

func (o *SystemEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemEventsInternalServerError creates a SystemEventsInternalServerError with default headers values
func NewSystemEventsInternalServerError() *SystemEventsInternalServerError {
	return &SystemEventsInternalServerError{}
}

/*
SystemEventsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SystemEventsInternalServerError struct {
	Payload *SystemEventsInternalServerErrorBody
}

// IsSuccess returns true when this system events internal server error response has a 2xx status code
func (o *SystemEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this system events internal server error response has a 3xx status code
func (o *SystemEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system events internal server error response has a 4xx status code
func (o *SystemEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this system events internal server error response has a 5xx status code
func (o *SystemEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this system events internal server error response a status code equal to that given
func (o *SystemEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SystemEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /events][%d] systemEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemEventsInternalServerError) String() string {
	return fmt.Sprintf("[GET /events][%d] systemEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemEventsInternalServerError) GetPayload() *SystemEventsInternalServerErrorBody {
	return o.Payload
}

func (o *SystemEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SystemEventsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SystemEventsInternalServerErrorBody system events internal server error body
swagger:model SystemEventsInternalServerErrorBody
*/
type SystemEventsInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this system events internal server error body
func (o *SystemEventsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system events internal server error body based on context it is used
func (o *SystemEventsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SystemEventsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemEventsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SystemEventsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
