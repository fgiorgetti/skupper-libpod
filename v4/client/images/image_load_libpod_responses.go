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

// ImageLoadLibpodReader is a Reader for the ImageLoadLibpod structure.
type ImageLoadLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageLoadLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageLoadLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageLoadLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageLoadLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageLoadLibpodOK creates a ImageLoadLibpodOK with default headers values
func NewImageLoadLibpodOK() *ImageLoadLibpodOK {
	return &ImageLoadLibpodOK{}
}

/*
ImageLoadLibpodOK describes a response with status code 200, with default header values.

Load response
*/
type ImageLoadLibpodOK struct {
	Payload *models.ImageLoadReport
}

// IsSuccess returns true when this image load libpod o k response has a 2xx status code
func (o *ImageLoadLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image load libpod o k response has a 3xx status code
func (o *ImageLoadLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image load libpod o k response has a 4xx status code
func (o *ImageLoadLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image load libpod o k response has a 5xx status code
func (o *ImageLoadLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image load libpod o k response a status code equal to that given
func (o *ImageLoadLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageLoadLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/load][%d] imageLoadLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageLoadLibpodOK) String() string {
	return fmt.Sprintf("[POST /libpod/images/load][%d] imageLoadLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageLoadLibpodOK) GetPayload() *models.ImageLoadReport {
	return o.Payload
}

func (o *ImageLoadLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageLoadReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageLoadLibpodBadRequest creates a ImageLoadLibpodBadRequest with default headers values
func NewImageLoadLibpodBadRequest() *ImageLoadLibpodBadRequest {
	return &ImageLoadLibpodBadRequest{}
}

/*
ImageLoadLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ImageLoadLibpodBadRequest struct {
	Payload *ImageLoadLibpodBadRequestBody
}

// IsSuccess returns true when this image load libpod bad request response has a 2xx status code
func (o *ImageLoadLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image load libpod bad request response has a 3xx status code
func (o *ImageLoadLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image load libpod bad request response has a 4xx status code
func (o *ImageLoadLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image load libpod bad request response has a 5xx status code
func (o *ImageLoadLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image load libpod bad request response a status code equal to that given
func (o *ImageLoadLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImageLoadLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/images/load][%d] imageLoadLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ImageLoadLibpodBadRequest) String() string {
	return fmt.Sprintf("[POST /libpod/images/load][%d] imageLoadLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ImageLoadLibpodBadRequest) GetPayload() *ImageLoadLibpodBadRequestBody {
	return o.Payload
}

func (o *ImageLoadLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageLoadLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageLoadLibpodInternalServerError creates a ImageLoadLibpodInternalServerError with default headers values
func NewImageLoadLibpodInternalServerError() *ImageLoadLibpodInternalServerError {
	return &ImageLoadLibpodInternalServerError{}
}

/*
ImageLoadLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageLoadLibpodInternalServerError struct {
	Payload *ImageLoadLibpodInternalServerErrorBody
}

// IsSuccess returns true when this image load libpod internal server error response has a 2xx status code
func (o *ImageLoadLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image load libpod internal server error response has a 3xx status code
func (o *ImageLoadLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image load libpod internal server error response has a 4xx status code
func (o *ImageLoadLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image load libpod internal server error response has a 5xx status code
func (o *ImageLoadLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image load libpod internal server error response a status code equal to that given
func (o *ImageLoadLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageLoadLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/load][%d] imageLoadLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageLoadLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/images/load][%d] imageLoadLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageLoadLibpodInternalServerError) GetPayload() *ImageLoadLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageLoadLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageLoadLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageLoadLibpodBadRequestBody image load libpod bad request body
swagger:model ImageLoadLibpodBadRequestBody
*/
type ImageLoadLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image load libpod bad request body
func (o *ImageLoadLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image load libpod bad request body based on context it is used
func (o *ImageLoadLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageLoadLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageLoadLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ImageLoadLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageLoadLibpodInternalServerErrorBody image load libpod internal server error body
swagger:model ImageLoadLibpodInternalServerErrorBody
*/
type ImageLoadLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image load libpod internal server error body
func (o *ImageLoadLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image load libpod internal server error body based on context it is used
func (o *ImageLoadLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageLoadLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageLoadLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageLoadLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
