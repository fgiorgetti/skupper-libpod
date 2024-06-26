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

	"github.com/fgiorgetti/skupper-libpod/v4/models"
)

// ImageListLibpodReader is a Reader for the ImageListLibpod structure.
type ImageListLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageListLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageListLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewImageListLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageListLibpodOK creates a ImageListLibpodOK with default headers values
func NewImageListLibpodOK() *ImageListLibpodOK {
	return &ImageListLibpodOK{}
}

/*
ImageListLibpodOK describes a response with status code 200, with default header values.

Image summary for libpod API
*/
type ImageListLibpodOK struct {
	Payload []*models.LibpodImageSummary
}

// IsSuccess returns true when this image list libpod o k response has a 2xx status code
func (o *ImageListLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image list libpod o k response has a 3xx status code
func (o *ImageListLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image list libpod o k response has a 4xx status code
func (o *ImageListLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image list libpod o k response has a 5xx status code
func (o *ImageListLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image list libpod o k response a status code equal to that given
func (o *ImageListLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageListLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/images/json][%d] imageListLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageListLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/images/json][%d] imageListLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageListLibpodOK) GetPayload() []*models.LibpodImageSummary {
	return o.Payload
}

func (o *ImageListLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListLibpodInternalServerError creates a ImageListLibpodInternalServerError with default headers values
func NewImageListLibpodInternalServerError() *ImageListLibpodInternalServerError {
	return &ImageListLibpodInternalServerError{}
}

/*
ImageListLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageListLibpodInternalServerError struct {
	Payload *ImageListLibpodInternalServerErrorBody
}

// IsSuccess returns true when this image list libpod internal server error response has a 2xx status code
func (o *ImageListLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image list libpod internal server error response has a 3xx status code
func (o *ImageListLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image list libpod internal server error response has a 4xx status code
func (o *ImageListLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image list libpod internal server error response has a 5xx status code
func (o *ImageListLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image list libpod internal server error response a status code equal to that given
func (o *ImageListLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageListLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/json][%d] imageListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageListLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/images/json][%d] imageListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageListLibpodInternalServerError) GetPayload() *ImageListLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageListLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageListLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageListLibpodInternalServerErrorBody image list libpod internal server error body
swagger:model ImageListLibpodInternalServerErrorBody
*/
type ImageListLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image list libpod internal server error body
func (o *ImageListLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image list libpod internal server error body based on context it is used
func (o *ImageListLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageListLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageListLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageListLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}