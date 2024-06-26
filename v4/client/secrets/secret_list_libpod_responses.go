// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fgiorgetti/skupper-libpod/v4/models"
)

// SecretListLibpodReader is a Reader for the SecretListLibpod structure.
type SecretListLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretListLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretListLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSecretListLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretListLibpodOK creates a SecretListLibpodOK with default headers values
func NewSecretListLibpodOK() *SecretListLibpodOK {
	return &SecretListLibpodOK{}
}

/*
SecretListLibpodOK describes a response with status code 200, with default header values.

Secret list response
*/
type SecretListLibpodOK struct {
	Payload []*models.SecretInfoReport
}

// IsSuccess returns true when this secret list libpod o k response has a 2xx status code
func (o *SecretListLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secret list libpod o k response has a 3xx status code
func (o *SecretListLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret list libpod o k response has a 4xx status code
func (o *SecretListLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret list libpod o k response has a 5xx status code
func (o *SecretListLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secret list libpod o k response a status code equal to that given
func (o *SecretListLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecretListLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/secrets/json][%d] secretListLibpodOK  %+v", 200, o.Payload)
}

func (o *SecretListLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/secrets/json][%d] secretListLibpodOK  %+v", 200, o.Payload)
}

func (o *SecretListLibpodOK) GetPayload() []*models.SecretInfoReport {
	return o.Payload
}

func (o *SecretListLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretListLibpodInternalServerError creates a SecretListLibpodInternalServerError with default headers values
func NewSecretListLibpodInternalServerError() *SecretListLibpodInternalServerError {
	return &SecretListLibpodInternalServerError{}
}

/*
SecretListLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretListLibpodInternalServerError struct {
	Payload *SecretListLibpodInternalServerErrorBody
}

// IsSuccess returns true when this secret list libpod internal server error response has a 2xx status code
func (o *SecretListLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret list libpod internal server error response has a 3xx status code
func (o *SecretListLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret list libpod internal server error response has a 4xx status code
func (o *SecretListLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret list libpod internal server error response has a 5xx status code
func (o *SecretListLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secret list libpod internal server error response a status code equal to that given
func (o *SecretListLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SecretListLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/secrets/json][%d] secretListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretListLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/secrets/json][%d] secretListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretListLibpodInternalServerError) GetPayload() *SecretListLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *SecretListLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretListLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecretListLibpodInternalServerErrorBody secret list libpod internal server error body
swagger:model SecretListLibpodInternalServerErrorBody
*/
type SecretListLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret list libpod internal server error body
func (o *SecretListLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret list libpod internal server error body based on context it is used
func (o *SecretListLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretListLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretListLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SecretListLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
