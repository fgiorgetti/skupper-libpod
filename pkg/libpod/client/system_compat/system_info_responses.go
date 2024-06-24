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

// SystemInfoReader is a Reader for the SystemInfo structure.
type SystemInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSystemInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemInfoOK creates a SystemInfoOK with default headers values
func NewSystemInfoOK() *SystemInfoOK {
	return &SystemInfoOK{}
}

/*
SystemInfoOK describes a response with status code 200, with default header values.

to be determined
*/
type SystemInfoOK struct {
}

// IsSuccess returns true when this system info o k response has a 2xx status code
func (o *SystemInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system info o k response has a 3xx status code
func (o *SystemInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system info o k response has a 4xx status code
func (o *SystemInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system info o k response has a 5xx status code
func (o *SystemInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this system info o k response a status code equal to that given
func (o *SystemInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *SystemInfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] systemInfoOK ", 200)
}

func (o *SystemInfoOK) String() string {
	return fmt.Sprintf("[GET /info][%d] systemInfoOK ", 200)
}

func (o *SystemInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemInfoInternalServerError creates a SystemInfoInternalServerError with default headers values
func NewSystemInfoInternalServerError() *SystemInfoInternalServerError {
	return &SystemInfoInternalServerError{}
}

/*
SystemInfoInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SystemInfoInternalServerError struct {
	Payload *SystemInfoInternalServerErrorBody
}

// IsSuccess returns true when this system info internal server error response has a 2xx status code
func (o *SystemInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this system info internal server error response has a 3xx status code
func (o *SystemInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system info internal server error response has a 4xx status code
func (o *SystemInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this system info internal server error response has a 5xx status code
func (o *SystemInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this system info internal server error response a status code equal to that given
func (o *SystemInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SystemInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /info][%d] systemInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /info][%d] systemInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemInfoInternalServerError) GetPayload() *SystemInfoInternalServerErrorBody {
	return o.Payload
}

func (o *SystemInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SystemInfoInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SystemInfoInternalServerErrorBody system info internal server error body
swagger:model SystemInfoInternalServerErrorBody
*/
type SystemInfoInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this system info internal server error body
func (o *SystemInfoInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system info internal server error body based on context it is used
func (o *SystemInfoInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SystemInfoInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SystemInfoInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SystemInfoInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}