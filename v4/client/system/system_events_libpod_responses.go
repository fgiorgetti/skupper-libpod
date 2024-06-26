// Code generated by go-swagger; DO NOT EDIT.

package system

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

// SystemEventsLibpodReader is a Reader for the SystemEventsLibpod structure.
type SystemEventsLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemEventsLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemEventsLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemEventsLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemEventsLibpodOK creates a SystemEventsLibpodOK with default headers values
func NewSystemEventsLibpodOK() *SystemEventsLibpodOK {
	return &SystemEventsLibpodOK{}
}

/*
SystemEventsLibpodOK describes a response with status code 200, with default header values.

returns a string of json data describing an event
*/
type SystemEventsLibpodOK struct {
}

// IsSuccess returns true when this system events libpod o k response has a 2xx status code
func (o *SystemEventsLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system events libpod o k response has a 3xx status code
func (o *SystemEventsLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system events libpod o k response has a 4xx status code
func (o *SystemEventsLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system events libpod o k response has a 5xx status code
func (o *SystemEventsLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system events libpod o k response a status code equal to that given
func (o *SystemEventsLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemEventsLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/events][%d] systemEventsLibpodOK ", 200)
}

func (o *SystemEventsLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/events][%d] systemEventsLibpodOK ", 200)
}

func (o *SystemEventsLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemEventsLibpodInternalServerError creates a SystemEventsLibpodInternalServerError with default headers values
func NewSystemEventsLibpodInternalServerError() *SystemEventsLibpodInternalServerError {
	return &SystemEventsLibpodInternalServerError{}
}

/*
SystemEventsLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SystemEventsLibpodInternalServerError struct {
	Payload *SystemEventsLibpodInternalServerErrorBody
}

// IsSuccess returns true when this system events libpod internal server error response has a 2xx status code
func (o *SystemEventsLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this system events libpod internal server error response has a 3xx status code
func (o *SystemEventsLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system events libpod internal server error response has a 4xx status code
func (o *SystemEventsLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this system events libpod internal server error response has a 5xx status code
func (o *SystemEventsLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this system events libpod internal server error response a status code equal to that given
func (o *SystemEventsLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SystemEventsLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/events][%d] systemEventsLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemEventsLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/events][%d] systemEventsLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemEventsLibpodInternalServerError) GetPayload() *SystemEventsLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *SystemEventsLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SystemEventsLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SystemEventsLibpodInternalServerErrorBody system events libpod internal server error body
swagger:model SystemEventsLibpodInternalServerErrorBody
*/
type SystemEventsLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this system events libpod internal server error body
func (o *SystemEventsLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system events libpod internal server error body based on context it is used
func (o *SystemEventsLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SystemEventsLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemEventsLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SystemEventsLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}