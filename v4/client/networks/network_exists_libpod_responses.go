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
)

// NetworkExistsLibpodReader is a Reader for the NetworkExistsLibpod structure.
type NetworkExistsLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkExistsLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNetworkExistsLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewNetworkExistsLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkExistsLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkExistsLibpodNoContent creates a NetworkExistsLibpodNoContent with default headers values
func NewNetworkExistsLibpodNoContent() *NetworkExistsLibpodNoContent {
	return &NetworkExistsLibpodNoContent{}
}

/*
NetworkExistsLibpodNoContent describes a response with status code 204, with default header values.

network exists
*/
type NetworkExistsLibpodNoContent struct {
}

// IsSuccess returns true when this network exists libpod no content response has a 2xx status code
func (o *NetworkExistsLibpodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network exists libpod no content response has a 3xx status code
func (o *NetworkExistsLibpodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network exists libpod no content response has a 4xx status code
func (o *NetworkExistsLibpodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this network exists libpod no content response has a 5xx status code
func (o *NetworkExistsLibpodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this network exists libpod no content response a status code equal to that given
func (o *NetworkExistsLibpodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *NetworkExistsLibpodNoContent) Error() string {
	return fmt.Sprintf("[GET /libpod/networks/{name}/exists][%d] networkExistsLibpodNoContent ", 204)
}

func (o *NetworkExistsLibpodNoContent) String() string {
	return fmt.Sprintf("[GET /libpod/networks/{name}/exists][%d] networkExistsLibpodNoContent ", 204)
}

func (o *NetworkExistsLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkExistsLibpodNotFound creates a NetworkExistsLibpodNotFound with default headers values
func NewNetworkExistsLibpodNotFound() *NetworkExistsLibpodNotFound {
	return &NetworkExistsLibpodNotFound{}
}

/*
NetworkExistsLibpodNotFound describes a response with status code 404, with default header values.

No such network
*/
type NetworkExistsLibpodNotFound struct {
	Payload *NetworkExistsLibpodNotFoundBody
}

// IsSuccess returns true when this network exists libpod not found response has a 2xx status code
func (o *NetworkExistsLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network exists libpod not found response has a 3xx status code
func (o *NetworkExistsLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network exists libpod not found response has a 4xx status code
func (o *NetworkExistsLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this network exists libpod not found response has a 5xx status code
func (o *NetworkExistsLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this network exists libpod not found response a status code equal to that given
func (o *NetworkExistsLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NetworkExistsLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/networks/{name}/exists][%d] networkExistsLibpodNotFound  %+v", 404, o.Payload)
}

func (o *NetworkExistsLibpodNotFound) String() string {
	return fmt.Sprintf("[GET /libpod/networks/{name}/exists][%d] networkExistsLibpodNotFound  %+v", 404, o.Payload)
}

func (o *NetworkExistsLibpodNotFound) GetPayload() *NetworkExistsLibpodNotFoundBody {
	return o.Payload
}

func (o *NetworkExistsLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkExistsLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkExistsLibpodInternalServerError creates a NetworkExistsLibpodInternalServerError with default headers values
func NewNetworkExistsLibpodInternalServerError() *NetworkExistsLibpodInternalServerError {
	return &NetworkExistsLibpodInternalServerError{}
}

/*
NetworkExistsLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkExistsLibpodInternalServerError struct {
	Payload *NetworkExistsLibpodInternalServerErrorBody
}

// IsSuccess returns true when this network exists libpod internal server error response has a 2xx status code
func (o *NetworkExistsLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network exists libpod internal server error response has a 3xx status code
func (o *NetworkExistsLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network exists libpod internal server error response has a 4xx status code
func (o *NetworkExistsLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network exists libpod internal server error response has a 5xx status code
func (o *NetworkExistsLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network exists libpod internal server error response a status code equal to that given
func (o *NetworkExistsLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkExistsLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/networks/{name}/exists][%d] networkExistsLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkExistsLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/networks/{name}/exists][%d] networkExistsLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkExistsLibpodInternalServerError) GetPayload() *NetworkExistsLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkExistsLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkExistsLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkExistsLibpodInternalServerErrorBody network exists libpod internal server error body
swagger:model NetworkExistsLibpodInternalServerErrorBody
*/
type NetworkExistsLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network exists libpod internal server error body
func (o *NetworkExistsLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network exists libpod internal server error body based on context it is used
func (o *NetworkExistsLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkExistsLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkExistsLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkExistsLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkExistsLibpodNotFoundBody network exists libpod not found body
swagger:model NetworkExistsLibpodNotFoundBody
*/
type NetworkExistsLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network exists libpod not found body
func (o *NetworkExistsLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network exists libpod not found body based on context it is used
func (o *NetworkExistsLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkExistsLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkExistsLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NetworkExistsLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
