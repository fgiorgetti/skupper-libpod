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

// ImageHistoryLibpodReader is a Reader for the ImageHistoryLibpod structure.
type ImageHistoryLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageHistoryLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageHistoryLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageHistoryLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageHistoryLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageHistoryLibpodOK creates a ImageHistoryLibpodOK with default headers values
func NewImageHistoryLibpodOK() *ImageHistoryLibpodOK {
	return &ImageHistoryLibpodOK{}
}

/*
ImageHistoryLibpodOK describes a response with status code 200, with default header values.

History response
*/
type ImageHistoryLibpodOK struct {
	Payload *ImageHistoryLibpodOKBody
}

// IsSuccess returns true when this image history libpod o k response has a 2xx status code
func (o *ImageHistoryLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image history libpod o k response has a 3xx status code
func (o *ImageHistoryLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image history libpod o k response has a 4xx status code
func (o *ImageHistoryLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image history libpod o k response has a 5xx status code
func (o *ImageHistoryLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image history libpod o k response a status code equal to that given
func (o *ImageHistoryLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageHistoryLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/history][%d] imageHistoryLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageHistoryLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/history][%d] imageHistoryLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageHistoryLibpodOK) GetPayload() *ImageHistoryLibpodOKBody {
	return o.Payload
}

func (o *ImageHistoryLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageHistoryLibpodOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageHistoryLibpodNotFound creates a ImageHistoryLibpodNotFound with default headers values
func NewImageHistoryLibpodNotFound() *ImageHistoryLibpodNotFound {
	return &ImageHistoryLibpodNotFound{}
}

/*
ImageHistoryLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageHistoryLibpodNotFound struct {
	Payload *ImageHistoryLibpodNotFoundBody
}

// IsSuccess returns true when this image history libpod not found response has a 2xx status code
func (o *ImageHistoryLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image history libpod not found response has a 3xx status code
func (o *ImageHistoryLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image history libpod not found response has a 4xx status code
func (o *ImageHistoryLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image history libpod not found response has a 5xx status code
func (o *ImageHistoryLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image history libpod not found response a status code equal to that given
func (o *ImageHistoryLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImageHistoryLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/history][%d] imageHistoryLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ImageHistoryLibpodNotFound) String() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/history][%d] imageHistoryLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ImageHistoryLibpodNotFound) GetPayload() *ImageHistoryLibpodNotFoundBody {
	return o.Payload
}

func (o *ImageHistoryLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageHistoryLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageHistoryLibpodInternalServerError creates a ImageHistoryLibpodInternalServerError with default headers values
func NewImageHistoryLibpodInternalServerError() *ImageHistoryLibpodInternalServerError {
	return &ImageHistoryLibpodInternalServerError{}
}

/*
ImageHistoryLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageHistoryLibpodInternalServerError struct {
	Payload *ImageHistoryLibpodInternalServerErrorBody
}

// IsSuccess returns true when this image history libpod internal server error response has a 2xx status code
func (o *ImageHistoryLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image history libpod internal server error response has a 3xx status code
func (o *ImageHistoryLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image history libpod internal server error response has a 4xx status code
func (o *ImageHistoryLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image history libpod internal server error response has a 5xx status code
func (o *ImageHistoryLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image history libpod internal server error response a status code equal to that given
func (o *ImageHistoryLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageHistoryLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/history][%d] imageHistoryLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageHistoryLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/history][%d] imageHistoryLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageHistoryLibpodInternalServerError) GetPayload() *ImageHistoryLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageHistoryLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageHistoryLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageHistoryLibpodInternalServerErrorBody image history libpod internal server error body
swagger:model ImageHistoryLibpodInternalServerErrorBody
*/
type ImageHistoryLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image history libpod internal server error body
func (o *ImageHistoryLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image history libpod internal server error body based on context it is used
func (o *ImageHistoryLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageHistoryLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageHistoryLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageHistoryLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageHistoryLibpodNotFoundBody image history libpod not found body
swagger:model ImageHistoryLibpodNotFoundBody
*/
type ImageHistoryLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image history libpod not found body
func (o *ImageHistoryLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image history libpod not found body based on context it is used
func (o *ImageHistoryLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageHistoryLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageHistoryLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImageHistoryLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageHistoryLibpodOKBody image history libpod o k body
swagger:model ImageHistoryLibpodOKBody
*/
type ImageHistoryLibpodOKBody struct {

	// comment
	Comment string `json:"Comment,omitempty"`

	// created
	Created int64 `json:"Created,omitempty"`

	// created by
	CreatedBy string `json:"CreatedBy,omitempty"`

	// ID
	ID string `json:"Id,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// tags
	Tags []string `json:"Tags"`
}

// Validate validates this image history libpod o k body
func (o *ImageHistoryLibpodOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image history libpod o k body based on context it is used
func (o *ImageHistoryLibpodOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageHistoryLibpodOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageHistoryLibpodOKBody) UnmarshalBinary(b []byte) error {
	var res ImageHistoryLibpodOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
