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

// ImagePushLibpodReader is a Reader for the ImagePushLibpod structure.
type ImagePushLibpodReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ImagePushLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagePushLibpodOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImagePushLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagePushLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagePushLibpodOK creates a ImagePushLibpodOK with default headers values
func NewImagePushLibpodOK(writer io.Writer) *ImagePushLibpodOK {
	return &ImagePushLibpodOK{

		Payload: writer,
	}
}

/*
ImagePushLibpodOK describes a response with status code 200, with default header values.

no error
*/
type ImagePushLibpodOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this image push libpod o k response has a 2xx status code
func (o *ImagePushLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image push libpod o k response has a 3xx status code
func (o *ImagePushLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image push libpod o k response has a 4xx status code
func (o *ImagePushLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image push libpod o k response has a 5xx status code
func (o *ImagePushLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image push libpod o k response a status code equal to that given
func (o *ImagePushLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImagePushLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodOK  %+v", 200, o.Payload)
}

func (o *ImagePushLibpodOK) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodOK  %+v", 200, o.Payload)
}

func (o *ImagePushLibpodOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ImagePushLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushLibpodNotFound creates a ImagePushLibpodNotFound with default headers values
func NewImagePushLibpodNotFound() *ImagePushLibpodNotFound {
	return &ImagePushLibpodNotFound{}
}

/*
ImagePushLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImagePushLibpodNotFound struct {
	Payload *ImagePushLibpodNotFoundBody
}

// IsSuccess returns true when this image push libpod not found response has a 2xx status code
func (o *ImagePushLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image push libpod not found response has a 3xx status code
func (o *ImagePushLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image push libpod not found response has a 4xx status code
func (o *ImagePushLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image push libpod not found response has a 5xx status code
func (o *ImagePushLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image push libpod not found response a status code equal to that given
func (o *ImagePushLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ImagePushLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ImagePushLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ImagePushLibpodNotFound) GetPayload() *ImagePushLibpodNotFoundBody {
	return o.Payload
}

func (o *ImagePushLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagePushLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushLibpodInternalServerError creates a ImagePushLibpodInternalServerError with default headers values
func NewImagePushLibpodInternalServerError() *ImagePushLibpodInternalServerError {
	return &ImagePushLibpodInternalServerError{}
}

/*
ImagePushLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImagePushLibpodInternalServerError struct {
	Payload *ImagePushLibpodInternalServerErrorBody
}

// IsSuccess returns true when this image push libpod internal server error response has a 2xx status code
func (o *ImagePushLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image push libpod internal server error response has a 3xx status code
func (o *ImagePushLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image push libpod internal server error response has a 4xx status code
func (o *ImagePushLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image push libpod internal server error response has a 5xx status code
func (o *ImagePushLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image push libpod internal server error response a status code equal to that given
func (o *ImagePushLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImagePushLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImagePushLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/push][%d] imagePushLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImagePushLibpodInternalServerError) GetPayload() *ImagePushLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImagePushLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagePushLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImagePushLibpodInternalServerErrorBody image push libpod internal server error body
swagger:model ImagePushLibpodInternalServerErrorBody
*/
type ImagePushLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image push libpod internal server error body
func (o *ImagePushLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image push libpod internal server error body based on context it is used
func (o *ImagePushLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagePushLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagePushLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImagePushLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImagePushLibpodNotFoundBody image push libpod not found body
swagger:model ImagePushLibpodNotFoundBody
*/
type ImagePushLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image push libpod not found body
func (o *ImagePushLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image push libpod not found body based on context it is used
func (o *ImagePushLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagePushLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagePushLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImagePushLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
