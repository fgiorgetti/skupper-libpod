// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NetworkListLibpodReader is a Reader for the NetworkListLibpod structure.
type NetworkListLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkListLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkListLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewNetworkListLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkListLibpodOK creates a NetworkListLibpodOK with default headers values
func NewNetworkListLibpodOK() *NetworkListLibpodOK {
	return &NetworkListLibpodOK{}
}

/*
NetworkListLibpodOK describes a response with status code 200, with default header values.

Network list
*/
type NetworkListLibpodOK struct {
	Payload []*models.Network
}

// IsSuccess returns true when this network list libpod o k response has a 2xx status code
func (o *NetworkListLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network list libpod o k response has a 3xx status code
func (o *NetworkListLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network list libpod o k response has a 4xx status code
func (o *NetworkListLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network list libpod o k response has a 5xx status code
func (o *NetworkListLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network list libpod o k response a status code equal to that given
func (o *NetworkListLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkListLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/networks/json][%d] networkListLibpodOK  %+v", 200, o.Payload)
}

func (o *NetworkListLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/networks/json][%d] networkListLibpodOK  %+v", 200, o.Payload)
}

func (o *NetworkListLibpodOK) GetPayload() []*models.Network {
	return o.Payload
}

func (o *NetworkListLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkListLibpodInternalServerError creates a NetworkListLibpodInternalServerError with default headers values
func NewNetworkListLibpodInternalServerError() *NetworkListLibpodInternalServerError {
	return &NetworkListLibpodInternalServerError{}
}

/*
NetworkListLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkListLibpodInternalServerError struct {
	Payload *NetworkListLibpodInternalServerErrorBody
}

// IsSuccess returns true when this network list libpod internal server error response has a 2xx status code
func (o *NetworkListLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network list libpod internal server error response has a 3xx status code
func (o *NetworkListLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network list libpod internal server error response has a 4xx status code
func (o *NetworkListLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network list libpod internal server error response has a 5xx status code
func (o *NetworkListLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network list libpod internal server error response a status code equal to that given
func (o *NetworkListLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkListLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/networks/json][%d] networkListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkListLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/networks/json][%d] networkListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkListLibpodInternalServerError) GetPayload() *NetworkListLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkListLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkListLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkListLibpodInternalServerErrorBody network list libpod internal server error body
swagger:model NetworkListLibpodInternalServerErrorBody
*/
type NetworkListLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network list libpod internal server error body
func (o *NetworkListLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network list libpod internal server error body based on context it is used
func (o *NetworkListLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkListLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkListLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkListLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}