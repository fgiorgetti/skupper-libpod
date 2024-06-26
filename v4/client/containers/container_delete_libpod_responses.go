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

	"github.com/fgiorgetti/skupper-libpod/v4/models"
)

// ContainerDeleteLibpodReader is a Reader for the ContainerDeleteLibpod structure.
type ContainerDeleteLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerDeleteLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerDeleteLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewContainerDeleteLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerDeleteLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerDeleteLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerDeleteLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerDeleteLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerDeleteLibpodOK creates a ContainerDeleteLibpodOK with default headers values
func NewContainerDeleteLibpodOK() *ContainerDeleteLibpodOK {
	return &ContainerDeleteLibpodOK{}
}

/*
ContainerDeleteLibpodOK describes a response with status code 200, with default header values.

Rm containers
*/
type ContainerDeleteLibpodOK struct {
	Payload []*models.LibpodContainersRmReport
}

// IsSuccess returns true when this container delete libpod o k response has a 2xx status code
func (o *ContainerDeleteLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container delete libpod o k response has a 3xx status code
func (o *ContainerDeleteLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete libpod o k response has a 4xx status code
func (o *ContainerDeleteLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container delete libpod o k response has a 5xx status code
func (o *ContainerDeleteLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete libpod o k response a status code equal to that given
func (o *ContainerDeleteLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ContainerDeleteLibpodOK) Error() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodOK  %+v", 200, o.Payload)
}

func (o *ContainerDeleteLibpodOK) String() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodOK  %+v", 200, o.Payload)
}

func (o *ContainerDeleteLibpodOK) GetPayload() []*models.LibpodContainersRmReport {
	return o.Payload
}

func (o *ContainerDeleteLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteLibpodNoContent creates a ContainerDeleteLibpodNoContent with default headers values
func NewContainerDeleteLibpodNoContent() *ContainerDeleteLibpodNoContent {
	return &ContainerDeleteLibpodNoContent{}
}

/*
ContainerDeleteLibpodNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerDeleteLibpodNoContent struct {
}

// IsSuccess returns true when this container delete libpod no content response has a 2xx status code
func (o *ContainerDeleteLibpodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container delete libpod no content response has a 3xx status code
func (o *ContainerDeleteLibpodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete libpod no content response has a 4xx status code
func (o *ContainerDeleteLibpodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this container delete libpod no content response has a 5xx status code
func (o *ContainerDeleteLibpodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete libpod no content response a status code equal to that given
func (o *ContainerDeleteLibpodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ContainerDeleteLibpodNoContent) Error() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodNoContent ", 204)
}

func (o *ContainerDeleteLibpodNoContent) String() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodNoContent ", 204)
}

func (o *ContainerDeleteLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerDeleteLibpodBadRequest creates a ContainerDeleteLibpodBadRequest with default headers values
func NewContainerDeleteLibpodBadRequest() *ContainerDeleteLibpodBadRequest {
	return &ContainerDeleteLibpodBadRequest{}
}

/*
ContainerDeleteLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ContainerDeleteLibpodBadRequest struct {
	Payload *ContainerDeleteLibpodBadRequestBody
}

// IsSuccess returns true when this container delete libpod bad request response has a 2xx status code
func (o *ContainerDeleteLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete libpod bad request response has a 3xx status code
func (o *ContainerDeleteLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete libpod bad request response has a 4xx status code
func (o *ContainerDeleteLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this container delete libpod bad request response has a 5xx status code
func (o *ContainerDeleteLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete libpod bad request response a status code equal to that given
func (o *ContainerDeleteLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ContainerDeleteLibpodBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerDeleteLibpodBadRequest) String() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ContainerDeleteLibpodBadRequest) GetPayload() *ContainerDeleteLibpodBadRequestBody {
	return o.Payload
}

func (o *ContainerDeleteLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerDeleteLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteLibpodNotFound creates a ContainerDeleteLibpodNotFound with default headers values
func NewContainerDeleteLibpodNotFound() *ContainerDeleteLibpodNotFound {
	return &ContainerDeleteLibpodNotFound{}
}

/*
ContainerDeleteLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerDeleteLibpodNotFound struct {
	Payload *ContainerDeleteLibpodNotFoundBody
}

// IsSuccess returns true when this container delete libpod not found response has a 2xx status code
func (o *ContainerDeleteLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete libpod not found response has a 3xx status code
func (o *ContainerDeleteLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete libpod not found response has a 4xx status code
func (o *ContainerDeleteLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container delete libpod not found response has a 5xx status code
func (o *ContainerDeleteLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete libpod not found response a status code equal to that given
func (o *ContainerDeleteLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerDeleteLibpodNotFound) Error() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerDeleteLibpodNotFound) String() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ContainerDeleteLibpodNotFound) GetPayload() *ContainerDeleteLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerDeleteLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerDeleteLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteLibpodConflict creates a ContainerDeleteLibpodConflict with default headers values
func NewContainerDeleteLibpodConflict() *ContainerDeleteLibpodConflict {
	return &ContainerDeleteLibpodConflict{}
}

/*
ContainerDeleteLibpodConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ContainerDeleteLibpodConflict struct {
	Payload *ContainerDeleteLibpodConflictBody
}

// IsSuccess returns true when this container delete libpod conflict response has a 2xx status code
func (o *ContainerDeleteLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete libpod conflict response has a 3xx status code
func (o *ContainerDeleteLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete libpod conflict response has a 4xx status code
func (o *ContainerDeleteLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this container delete libpod conflict response has a 5xx status code
func (o *ContainerDeleteLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this container delete libpod conflict response a status code equal to that given
func (o *ContainerDeleteLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ContainerDeleteLibpodConflict) Error() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodConflict  %+v", 409, o.Payload)
}

func (o *ContainerDeleteLibpodConflict) String() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodConflict  %+v", 409, o.Payload)
}

func (o *ContainerDeleteLibpodConflict) GetPayload() *ContainerDeleteLibpodConflictBody {
	return o.Payload
}

func (o *ContainerDeleteLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerDeleteLibpodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerDeleteLibpodInternalServerError creates a ContainerDeleteLibpodInternalServerError with default headers values
func NewContainerDeleteLibpodInternalServerError() *ContainerDeleteLibpodInternalServerError {
	return &ContainerDeleteLibpodInternalServerError{}
}

/*
ContainerDeleteLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerDeleteLibpodInternalServerError struct {
	Payload *ContainerDeleteLibpodInternalServerErrorBody
}

// IsSuccess returns true when this container delete libpod internal server error response has a 2xx status code
func (o *ContainerDeleteLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container delete libpod internal server error response has a 3xx status code
func (o *ContainerDeleteLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container delete libpod internal server error response has a 4xx status code
func (o *ContainerDeleteLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container delete libpod internal server error response has a 5xx status code
func (o *ContainerDeleteLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container delete libpod internal server error response a status code equal to that given
func (o *ContainerDeleteLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerDeleteLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerDeleteLibpodInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /libpod/containers/{name}][%d] containerDeleteLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerDeleteLibpodInternalServerError) GetPayload() *ContainerDeleteLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerDeleteLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerDeleteLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerDeleteLibpodBadRequestBody container delete libpod bad request body
swagger:model ContainerDeleteLibpodBadRequestBody
*/
type ContainerDeleteLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container delete libpod bad request body
func (o *ContainerDeleteLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container delete libpod bad request body based on context it is used
func (o *ContainerDeleteLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerDeleteLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerDeleteLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ContainerDeleteLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerDeleteLibpodConflictBody container delete libpod conflict body
swagger:model ContainerDeleteLibpodConflictBody
*/
type ContainerDeleteLibpodConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container delete libpod conflict body
func (o *ContainerDeleteLibpodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container delete libpod conflict body based on context it is used
func (o *ContainerDeleteLibpodConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerDeleteLibpodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerDeleteLibpodConflictBody) UnmarshalBinary(b []byte) error {
	var res ContainerDeleteLibpodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerDeleteLibpodInternalServerErrorBody container delete libpod internal server error body
swagger:model ContainerDeleteLibpodInternalServerErrorBody
*/
type ContainerDeleteLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container delete libpod internal server error body
func (o *ContainerDeleteLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container delete libpod internal server error body based on context it is used
func (o *ContainerDeleteLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerDeleteLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerDeleteLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerDeleteLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerDeleteLibpodNotFoundBody container delete libpod not found body
swagger:model ContainerDeleteLibpodNotFoundBody
*/
type ContainerDeleteLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container delete libpod not found body
func (o *ContainerDeleteLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container delete libpod not found body based on context it is used
func (o *ContainerDeleteLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerDeleteLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerDeleteLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerDeleteLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
