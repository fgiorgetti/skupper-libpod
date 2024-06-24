// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// ContainerCreateReader is a Reader for the ContainerCreate structure.
type ContainerCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewContainerCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerCreateCreated creates a ContainerCreateCreated with default headers values
func NewContainerCreateCreated() *ContainerCreateCreated {
	return &ContainerCreateCreated{}
}

/*
ContainerCreateCreated describes a response with status code 201, with default header values.

Create container
*/
type ContainerCreateCreated struct {
	Payload *ContainerCreateCreatedBody
}

// IsSuccess returns true when this container create created response has a 2xx status code
func (o *ContainerCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container create created response has a 3xx status code
func (o *ContainerCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container create created response has a 4xx status code
func (o *ContainerCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this container create created response has a 5xx status code
func (o *ContainerCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this container create created response a status code equal to that given
func (o *ContainerCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ContainerCreateCreated) Error() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateCreated  %+v", 201, o.Payload)
}

func (o *ContainerCreateCreated) String() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateCreated  %+v", 201, o.Payload)
}

func (o *ContainerCreateCreated) GetPayload() *ContainerCreateCreatedBody {
	return o.Payload
}

func (o *ContainerCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCreateCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerCreateBadRequest creates a ContainerCreateBadRequest with default headers values
func NewContainerCreateBadRequest() *ContainerCreateBadRequest {
	return &ContainerCreateBadRequest{}
}

/*
ContainerCreateBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ContainerCreateBadRequest struct {
	Payload *ContainerCreateBadRequestBody
}

// IsSuccess returns true when this container create bad request response has a 2xx status code
func (o *ContainerCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container create bad request response has a 3xx status code
func (o *ContainerCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container create bad request response has a 4xx status code
func (o *ContainerCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this container create bad request response has a 5xx status code
func (o *ContainerCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this container create bad request response a status code equal to that given
func (o *ContainerCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ContainerCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerCreateBadRequest) GetPayload() *ContainerCreateBadRequestBody {
	return o.Payload
}

func (o *ContainerCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCreateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerCreateNotFound creates a ContainerCreateNotFound with default headers values
func NewContainerCreateNotFound() *ContainerCreateNotFound {
	return &ContainerCreateNotFound{}
}

/*
ContainerCreateNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerCreateNotFound struct {
	Payload *ContainerCreateNotFoundBody
}

// IsSuccess returns true when this container create not found response has a 2xx status code
func (o *ContainerCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container create not found response has a 3xx status code
func (o *ContainerCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container create not found response has a 4xx status code
func (o *ContainerCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container create not found response has a 5xx status code
func (o *ContainerCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container create not found response a status code equal to that given
func (o *ContainerCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateNotFound  %+v", 404, o.Payload)
}

func (o *ContainerCreateNotFound) String() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateNotFound  %+v", 404, o.Payload)
}

func (o *ContainerCreateNotFound) GetPayload() *ContainerCreateNotFoundBody {
	return o.Payload
}

func (o *ContainerCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCreateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerCreateConflict creates a ContainerCreateConflict with default headers values
func NewContainerCreateConflict() *ContainerCreateConflict {
	return &ContainerCreateConflict{}
}

/*
ContainerCreateConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ContainerCreateConflict struct {
	Payload *ContainerCreateConflictBody
}

// IsSuccess returns true when this container create conflict response has a 2xx status code
func (o *ContainerCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container create conflict response has a 3xx status code
func (o *ContainerCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container create conflict response has a 4xx status code
func (o *ContainerCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this container create conflict response has a 5xx status code
func (o *ContainerCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this container create conflict response a status code equal to that given
func (o *ContainerCreateConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ContainerCreateConflict) Error() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateConflict  %+v", 409, o.Payload)
}

func (o *ContainerCreateConflict) String() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateConflict  %+v", 409, o.Payload)
}

func (o *ContainerCreateConflict) GetPayload() *ContainerCreateConflictBody {
	return o.Payload
}

func (o *ContainerCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCreateConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerCreateInternalServerError creates a ContainerCreateInternalServerError with default headers values
func NewContainerCreateInternalServerError() *ContainerCreateInternalServerError {
	return &ContainerCreateInternalServerError{}
}

/*
ContainerCreateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerCreateInternalServerError struct {
	Payload *ContainerCreateInternalServerErrorBody
}

// IsSuccess returns true when this container create internal server error response has a 2xx status code
func (o *ContainerCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container create internal server error response has a 3xx status code
func (o *ContainerCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container create internal server error response has a 4xx status code
func (o *ContainerCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container create internal server error response has a 5xx status code
func (o *ContainerCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container create internal server error response a status code equal to that given
func (o *ContainerCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/create][%d] containerCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerCreateInternalServerError) GetPayload() *ContainerCreateInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCreateInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerCreateBadRequestBody container create bad request body
swagger:model ContainerCreateBadRequestBody
*/
type ContainerCreateBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container create bad request body
func (o *ContainerCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container create bad request body based on context it is used
func (o *ContainerCreateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ContainerCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerCreateConflictBody container create conflict body
swagger:model ContainerCreateConflictBody
*/
type ContainerCreateConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container create conflict body
func (o *ContainerCreateConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container create conflict body based on context it is used
func (o *ContainerCreateConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCreateConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCreateConflictBody) UnmarshalBinary(b []byte) error {
	var res ContainerCreateConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerCreateCreatedBody container create created body
swagger:model ContainerCreateCreatedBody
*/
type ContainerCreateCreatedBody struct {

	// ID of the container created
	ID string `json:"Id,omitempty"`

	// Warnings during container creation
	Warnings []string `json:"Warnings"`
}

// Validate validates this container create created body
func (o *ContainerCreateCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container create created body based on context it is used
func (o *ContainerCreateCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res ContainerCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerCreateInternalServerErrorBody container create internal server error body
swagger:model ContainerCreateInternalServerErrorBody
*/
type ContainerCreateInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container create internal server error body
func (o *ContainerCreateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container create internal server error body based on context it is used
func (o *ContainerCreateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCreateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCreateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerCreateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerCreateNotFoundBody container create not found body
swagger:model ContainerCreateNotFoundBody
*/
type ContainerCreateNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container create not found body
func (o *ContainerCreateNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container create not found body based on context it is used
func (o *ContainerCreateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCreateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCreateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerCreateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}