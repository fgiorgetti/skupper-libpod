// Code generated by go-swagger; DO NOT EDIT.

package images

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

// ImageUntagLibpodReader is a Reader for the ImageUntagLibpod structure.
type ImageUntagLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageUntagLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImageUntagLibpodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageUntagLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageUntagLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageUntagLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageUntagLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageUntagLibpodCreated creates a ImageUntagLibpodCreated with default headers values
func NewImageUntagLibpodCreated() *ImageUntagLibpodCreated {
	return &ImageUntagLibpodCreated{}
}

/*
ImageUntagLibpodCreated describes a response with status code 201, with default header values.

no error
*/
type ImageUntagLibpodCreated struct {
}

// IsSuccess returns true when this image untag libpod created response has a 2xx status code
func (o *ImageUntagLibpodCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image untag libpod created response has a 3xx status code
func (o *ImageUntagLibpodCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image untag libpod created response has a 4xx status code
func (o *ImageUntagLibpodCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this image untag libpod created response has a 5xx status code
func (o *ImageUntagLibpodCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this image untag libpod created response a status code equal to that given
func (o *ImageUntagLibpodCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ImageUntagLibpodCreated) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodCreated ", 201)
}

func (o *ImageUntagLibpodCreated) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodCreated ", 201)
}

func (o *ImageUntagLibpodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageUntagLibpodBadRequest creates a ImageUntagLibpodBadRequest with default headers values
func NewImageUntagLibpodBadRequest() *ImageUntagLibpodBadRequest {
	return &ImageUntagLibpodBadRequest{}
}

/*
ImageUntagLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ImageUntagLibpodBadRequest struct {
	Payload *ImageUntagLibpodBadRequestBody
}

// IsSuccess returns true when this image untag libpod bad request response has a 2xx status code
func (o *ImageUntagLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image untag libpod bad request response has a 3xx status code
func (o *ImageUntagLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image untag libpod bad request response has a 4xx status code
func (o *ImageUntagLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image untag libpod bad request response has a 5xx status code
func (o *ImageUntagLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image untag libpod bad request response a status code equal to that given
func (o *ImageUntagLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImageUntagLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ImageUntagLibpodBadRequest) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ImageUntagLibpodBadRequest) GetPayload() *ImageUntagLibpodBadRequestBody {
	return o.Payload
}

func (o *ImageUntagLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageUntagLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageUntagLibpodNotFound creates a ImageUntagLibpodNotFound with default headers values
func NewImageUntagLibpodNotFound() *ImageUntagLibpodNotFound {
	return &ImageUntagLibpodNotFound{}
}

/*
ImageUntagLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageUntagLibpodNotFound struct {
	Payload *ImageUntagLibpodNotFoundBody
}

// IsSuccess returns true when this image untag libpod not found response has a 2xx status code
func (o *ImageUntagLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image untag libpod not found response has a 3xx status code
func (o *ImageUntagLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image untag libpod not found response has a 4xx status code
func (o *ImageUntagLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image untag libpod not found response has a 5xx status code
func (o *ImageUntagLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image untag libpod not found response a status code equal to that given
func (o *ImageUntagLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImageUntagLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ImageUntagLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ImageUntagLibpodNotFound) GetPayload() *ImageUntagLibpodNotFoundBody {
	return o.Payload
}

func (o *ImageUntagLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageUntagLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageUntagLibpodConflict creates a ImageUntagLibpodConflict with default headers values
func NewImageUntagLibpodConflict() *ImageUntagLibpodConflict {
	return &ImageUntagLibpodConflict{}
}

/*
ImageUntagLibpodConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ImageUntagLibpodConflict struct {
	Payload *ImageUntagLibpodConflictBody
}

// IsSuccess returns true when this image untag libpod conflict response has a 2xx status code
func (o *ImageUntagLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image untag libpod conflict response has a 3xx status code
func (o *ImageUntagLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image untag libpod conflict response has a 4xx status code
func (o *ImageUntagLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image untag libpod conflict response has a 5xx status code
func (o *ImageUntagLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image untag libpod conflict response a status code equal to that given
func (o *ImageUntagLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ImageUntagLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodConflict  %+v", 409, o.Payload)
}

func (o *ImageUntagLibpodConflict) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodConflict  %+v", 409, o.Payload)
}

func (o *ImageUntagLibpodConflict) GetPayload() *ImageUntagLibpodConflictBody {
	return o.Payload
}

func (o *ImageUntagLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageUntagLibpodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageUntagLibpodInternalServerError creates a ImageUntagLibpodInternalServerError with default headers values
func NewImageUntagLibpodInternalServerError() *ImageUntagLibpodInternalServerError {
	return &ImageUntagLibpodInternalServerError{}
}

/*
ImageUntagLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageUntagLibpodInternalServerError struct {
	Payload *ImageUntagLibpodInternalServerErrorBody
}

// IsSuccess returns true when this image untag libpod internal server error response has a 2xx status code
func (o *ImageUntagLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image untag libpod internal server error response has a 3xx status code
func (o *ImageUntagLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image untag libpod internal server error response has a 4xx status code
func (o *ImageUntagLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image untag libpod internal server error response has a 5xx status code
func (o *ImageUntagLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image untag libpod internal server error response a status code equal to that given
func (o *ImageUntagLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageUntagLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageUntagLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/untag][%d] imageUntagLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageUntagLibpodInternalServerError) GetPayload() *ImageUntagLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageUntagLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageUntagLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageUntagLibpodBadRequestBody image untag libpod bad request body
swagger:model ImageUntagLibpodBadRequestBody
*/
type ImageUntagLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image untag libpod bad request body
func (o *ImageUntagLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image untag libpod bad request body based on context it is used
func (o *ImageUntagLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageUntagLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageUntagLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ImageUntagLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageUntagLibpodConflictBody image untag libpod conflict body
swagger:model ImageUntagLibpodConflictBody
*/
type ImageUntagLibpodConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image untag libpod conflict body
func (o *ImageUntagLibpodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image untag libpod conflict body based on context it is used
func (o *ImageUntagLibpodConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageUntagLibpodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageUntagLibpodConflictBody) UnmarshalBinary(b []byte) error {
	var res ImageUntagLibpodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageUntagLibpodInternalServerErrorBody image untag libpod internal server error body
swagger:model ImageUntagLibpodInternalServerErrorBody
*/
type ImageUntagLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image untag libpod internal server error body
func (o *ImageUntagLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image untag libpod internal server error body based on context it is used
func (o *ImageUntagLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageUntagLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageUntagLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageUntagLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageUntagLibpodNotFoundBody image untag libpod not found body
swagger:model ImageUntagLibpodNotFoundBody
*/
type ImageUntagLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image untag libpod not found body
func (o *ImageUntagLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image untag libpod not found body based on context it is used
func (o *ImageUntagLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageUntagLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageUntagLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImageUntagLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
