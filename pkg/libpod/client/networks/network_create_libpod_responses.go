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

	"github.com/fgiorgetti/skupper-libpod/pkg/libpod/models"
)

// NetworkCreateLibpodReader is a Reader for the NetworkCreateLibpod structure.
type NetworkCreateLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkCreateLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkCreateLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNetworkCreateLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewNetworkCreateLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkCreateLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkCreateLibpodOK creates a NetworkCreateLibpodOK with default headers values
func NewNetworkCreateLibpodOK() *NetworkCreateLibpodOK {
	return &NetworkCreateLibpodOK{}
}

/*
NetworkCreateLibpodOK describes a response with status code 200, with default header values.

Network create
*/
type NetworkCreateLibpodOK struct {
	Payload *models.Network
}

// IsSuccess returns true when this network create libpod o k response has a 2xx status code
func (o *NetworkCreateLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network create libpod o k response has a 3xx status code
func (o *NetworkCreateLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create libpod o k response has a 4xx status code
func (o *NetworkCreateLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create libpod o k response has a 5xx status code
func (o *NetworkCreateLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network create libpod o k response a status code equal to that given
func (o *NetworkCreateLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkCreateLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodOK  %+v", 200, o.Payload)
}

func (o *NetworkCreateLibpodOK) String() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodOK  %+v", 200, o.Payload)
}

func (o *NetworkCreateLibpodOK) GetPayload() *models.Network {
	return o.Payload
}

func (o *NetworkCreateLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Network)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateLibpodBadRequest creates a NetworkCreateLibpodBadRequest with default headers values
func NewNetworkCreateLibpodBadRequest() *NetworkCreateLibpodBadRequest {
	return &NetworkCreateLibpodBadRequest{}
}

/*
NetworkCreateLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type NetworkCreateLibpodBadRequest struct {
	Payload *NetworkCreateLibpodBadRequestBody
}

// IsSuccess returns true when this network create libpod bad request response has a 2xx status code
func (o *NetworkCreateLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create libpod bad request response has a 3xx status code
func (o *NetworkCreateLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create libpod bad request response has a 4xx status code
func (o *NetworkCreateLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this network create libpod bad request response has a 5xx status code
func (o *NetworkCreateLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this network create libpod bad request response a status code equal to that given
func (o *NetworkCreateLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NetworkCreateLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *NetworkCreateLibpodBadRequest) String() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *NetworkCreateLibpodBadRequest) GetPayload() *NetworkCreateLibpodBadRequestBody {
	return o.Payload
}

func (o *NetworkCreateLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateLibpodConflict creates a NetworkCreateLibpodConflict with default headers values
func NewNetworkCreateLibpodConflict() *NetworkCreateLibpodConflict {
	return &NetworkCreateLibpodConflict{}
}

/*
NetworkCreateLibpodConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type NetworkCreateLibpodConflict struct {
	Payload *NetworkCreateLibpodConflictBody
}

// IsSuccess returns true when this network create libpod conflict response has a 2xx status code
func (o *NetworkCreateLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create libpod conflict response has a 3xx status code
func (o *NetworkCreateLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create libpod conflict response has a 4xx status code
func (o *NetworkCreateLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this network create libpod conflict response has a 5xx status code
func (o *NetworkCreateLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this network create libpod conflict response a status code equal to that given
func (o *NetworkCreateLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *NetworkCreateLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodConflict  %+v", 409, o.Payload)
}

func (o *NetworkCreateLibpodConflict) String() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodConflict  %+v", 409, o.Payload)
}

func (o *NetworkCreateLibpodConflict) GetPayload() *NetworkCreateLibpodConflictBody {
	return o.Payload
}

func (o *NetworkCreateLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateLibpodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateLibpodInternalServerError creates a NetworkCreateLibpodInternalServerError with default headers values
func NewNetworkCreateLibpodInternalServerError() *NetworkCreateLibpodInternalServerError {
	return &NetworkCreateLibpodInternalServerError{}
}

/*
NetworkCreateLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkCreateLibpodInternalServerError struct {
	Payload *NetworkCreateLibpodInternalServerErrorBody
}

// IsSuccess returns true when this network create libpod internal server error response has a 2xx status code
func (o *NetworkCreateLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create libpod internal server error response has a 3xx status code
func (o *NetworkCreateLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create libpod internal server error response has a 4xx status code
func (o *NetworkCreateLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create libpod internal server error response has a 5xx status code
func (o *NetworkCreateLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network create libpod internal server error response a status code equal to that given
func (o *NetworkCreateLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkCreateLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkCreateLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/networks/create][%d] networkCreateLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkCreateLibpodInternalServerError) GetPayload() *NetworkCreateLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkCreateLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkCreateLibpodBadRequestBody network create libpod bad request body
swagger:model NetworkCreateLibpodBadRequestBody
*/
type NetworkCreateLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network create libpod bad request body
func (o *NetworkCreateLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network create libpod bad request body based on context it is used
func (o *NetworkCreateLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkCreateLibpodConflictBody network create libpod conflict body
swagger:model NetworkCreateLibpodConflictBody
*/
type NetworkCreateLibpodConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network create libpod conflict body
func (o *NetworkCreateLibpodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network create libpod conflict body based on context it is used
func (o *NetworkCreateLibpodConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateLibpodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateLibpodConflictBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateLibpodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkCreateLibpodInternalServerErrorBody network create libpod internal server error body
swagger:model NetworkCreateLibpodInternalServerErrorBody
*/
type NetworkCreateLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network create libpod internal server error body
func (o *NetworkCreateLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network create libpod internal server error body based on context it is used
func (o *NetworkCreateLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}