// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

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

// NetworkDeleteReader is a Reader for the NetworkDelete structure.
type NetworkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewNetworkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewNetworkDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkDeleteNoContent creates a NetworkDeleteNoContent with default headers values
func NewNetworkDeleteNoContent() *NetworkDeleteNoContent {
	return &NetworkDeleteNoContent{}
}

/*
NetworkDeleteNoContent describes a response with status code 204, with default header values.

no error
*/
type NetworkDeleteNoContent struct {
}

// IsSuccess returns true when this network delete no content response has a 2xx status code
func (o *NetworkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network delete no content response has a 3xx status code
func (o *NetworkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network delete no content response has a 4xx status code
func (o *NetworkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this network delete no content response has a 5xx status code
func (o *NetworkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this network delete no content response a status code equal to that given
func (o *NetworkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *NetworkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{name}][%d] networkDeleteNoContent ", 204)
}

func (o *NetworkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{name}][%d] networkDeleteNoContent ", 204)
}

func (o *NetworkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkDeleteNotFound creates a NetworkDeleteNotFound with default headers values
func NewNetworkDeleteNotFound() *NetworkDeleteNotFound {
	return &NetworkDeleteNotFound{}
}

/*
NetworkDeleteNotFound describes a response with status code 404, with default header values.

No such network
*/
type NetworkDeleteNotFound struct {
	Payload *NetworkDeleteNotFoundBody
}

// IsSuccess returns true when this network delete not found response has a 2xx status code
func (o *NetworkDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network delete not found response has a 3xx status code
func (o *NetworkDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network delete not found response has a 4xx status code
func (o *NetworkDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this network delete not found response has a 5xx status code
func (o *NetworkDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this network delete not found response a status code equal to that given
func (o *NetworkDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NetworkDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /networks/{name}][%d] networkDeleteNotFound  %+v", 404, o.Payload)
}

func (o *NetworkDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /networks/{name}][%d] networkDeleteNotFound  %+v", 404, o.Payload)
}

func (o *NetworkDeleteNotFound) GetPayload() *NetworkDeleteNotFoundBody {
	return o.Payload
}

func (o *NetworkDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkDeleteNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkDeleteInternalServerError creates a NetworkDeleteInternalServerError with default headers values
func NewNetworkDeleteInternalServerError() *NetworkDeleteInternalServerError {
	return &NetworkDeleteInternalServerError{}
}

/*
NetworkDeleteInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkDeleteInternalServerError struct {
	Payload *NetworkDeleteInternalServerErrorBody
}

// IsSuccess returns true when this network delete internal server error response has a 2xx status code
func (o *NetworkDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network delete internal server error response has a 3xx status code
func (o *NetworkDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network delete internal server error response has a 4xx status code
func (o *NetworkDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network delete internal server error response has a 5xx status code
func (o *NetworkDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network delete internal server error response a status code equal to that given
func (o *NetworkDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /networks/{name}][%d] networkDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /networks/{name}][%d] networkDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkDeleteInternalServerError) GetPayload() *NetworkDeleteInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkDeleteInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkDeleteInternalServerErrorBody network delete internal server error body
swagger:model NetworkDeleteInternalServerErrorBody
*/
type NetworkDeleteInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network delete internal server error body
func (o *NetworkDeleteInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network delete internal server error body based on context it is used
func (o *NetworkDeleteInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDeleteInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDeleteInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkDeleteInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkDeleteNotFoundBody network delete not found body
swagger:model NetworkDeleteNotFoundBody
*/
type NetworkDeleteNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network delete not found body
func (o *NetworkDeleteNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network delete not found body based on context it is used
func (o *NetworkDeleteNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDeleteNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDeleteNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NetworkDeleteNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
