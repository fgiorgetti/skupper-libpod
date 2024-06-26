// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// ContainerAttachLibpodReader is a Reader for the ContainerAttachLibpod structure.
type ContainerAttachLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerAttachLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 101:
		result := NewContainerAttachLibpodSwitchingProtocols()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewContainerAttachLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerAttachLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerAttachLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerAttachLibpodSwitchingProtocols creates a ContainerAttachLibpodSwitchingProtocols with default headers values
func NewContainerAttachLibpodSwitchingProtocols() *ContainerAttachLibpodSwitchingProtocols {
	return &ContainerAttachLibpodSwitchingProtocols{}
}

/*
ContainerAttachLibpodSwitchingProtocols describes a response with status code 101, with default header values.

No error, connection has been hijacked for transporting streams.
*/
type ContainerAttachLibpodSwitchingProtocols struct {
}

// IsSuccess returns true when this container attach libpod switching protocols response has a 2xx status code
func (o *ContainerAttachLibpodSwitchingProtocols) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach libpod switching protocols response has a 3xx status code
func (o *ContainerAttachLibpodSwitchingProtocols) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach libpod switching protocols response has a 4xx status code
func (o *ContainerAttachLibpodSwitchingProtocols) IsClientError() bool {
	return false
}

// IsServerError returns true when this container attach libpod switching protocols response has a 5xx status code
func (o *ContainerAttachLibpodSwitchingProtocols) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach libpod switching protocols response a status code equal to that given
func (o *ContainerAttachLibpodSwitchingProtocols) IsCode(code int) bool {
	return code == 101
}

func (o *ContainerAttachLibpodSwitchingProtocols) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodSwitchingProtocols ", 101)
}

func (o *ContainerAttachLibpodSwitchingProtocols) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodSwitchingProtocols ", 101)
}

func (o *ContainerAttachLibpodSwitchingProtocols) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachLibpodBadRequest creates a ContainerAttachLibpodBadRequest with default headers values
func NewContainerAttachLibpodBadRequest() *ContainerAttachLibpodBadRequest {
	return &ContainerAttachLibpodBadRequest{}
}

/*
ContainerAttachLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ContainerAttachLibpodBadRequest struct {
	Payload *ContainerAttachLibpodBadRequestBody
}

// IsSuccess returns true when this container attach libpod bad request response has a 2xx status code
func (o *ContainerAttachLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach libpod bad request response has a 3xx status code
func (o *ContainerAttachLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach libpod bad request response has a 4xx status code
func (o *ContainerAttachLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this container attach libpod bad request response has a 5xx status code
func (o *ContainerAttachLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach libpod bad request response a status code equal to that given
func (o *ContainerAttachLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ContainerAttachLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerAttachLibpodBadRequest) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerAttachLibpodBadRequest) GetPayload() *ContainerAttachLibpodBadRequestBody {
	return o.Payload
}

func (o *ContainerAttachLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerAttachLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachLibpodNotFound creates a ContainerAttachLibpodNotFound with default headers values
func NewContainerAttachLibpodNotFound() *ContainerAttachLibpodNotFound {
	return &ContainerAttachLibpodNotFound{}
}

/*
ContainerAttachLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerAttachLibpodNotFound struct {
	Payload *ContainerAttachLibpodNotFoundBody
}

// IsSuccess returns true when this container attach libpod not found response has a 2xx status code
func (o *ContainerAttachLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach libpod not found response has a 3xx status code
func (o *ContainerAttachLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach libpod not found response has a 4xx status code
func (o *ContainerAttachLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container attach libpod not found response has a 5xx status code
func (o *ContainerAttachLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container attach libpod not found response a status code equal to that given
func (o *ContainerAttachLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerAttachLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerAttachLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerAttachLibpodNotFound) GetPayload() *ContainerAttachLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerAttachLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerAttachLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachLibpodInternalServerError creates a ContainerAttachLibpodInternalServerError with default headers values
func NewContainerAttachLibpodInternalServerError() *ContainerAttachLibpodInternalServerError {
	return &ContainerAttachLibpodInternalServerError{}
}

/*
ContainerAttachLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerAttachLibpodInternalServerError struct {
	Payload *ContainerAttachLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container attach libpod internal server error response has a 2xx status code
func (o *ContainerAttachLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container attach libpod internal server error response has a 3xx status code
func (o *ContainerAttachLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container attach libpod internal server error response has a 4xx status code
func (o *ContainerAttachLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container attach libpod internal server error response has a 5xx status code
func (o *ContainerAttachLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container attach libpod internal server error response a status code equal to that given
func (o *ContainerAttachLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerAttachLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerAttachLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/attach][%d] containerAttachLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerAttachLibpodInternalServerError) GetPayload() *ContainerAttachLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerAttachLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerAttachLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerAttachLibpodBadRequestBody container attach libpod bad request body
swagger:model ContainerAttachLibpodBadRequestBody
*/
type ContainerAttachLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container attach libpod bad request body
func (o *ContainerAttachLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container attach libpod bad request body based on context it is used
func (o *ContainerAttachLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerAttachLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerAttachLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ContainerAttachLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerAttachLibpodInternalServerErrorBody container attach libpod internal server error body
swagger:model ContainerAttachLibpodInternalServerErrorBody
*/
type ContainerAttachLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container attach libpod internal server error body
func (o *ContainerAttachLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container attach libpod internal server error body based on context it is used
func (o *ContainerAttachLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerAttachLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerAttachLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerAttachLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerAttachLibpodNotFoundBody container attach libpod not found body
swagger:model ContainerAttachLibpodNotFoundBody
*/
type ContainerAttachLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container attach libpod not found body
func (o *ContainerAttachLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container attach libpod not found body based on context it is used
func (o *ContainerAttachLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerAttachLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerAttachLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerAttachLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
