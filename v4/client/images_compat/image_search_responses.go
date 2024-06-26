// Code generated by go-swagger; DO NOT EDIT.

package images_compat

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

// ImageSearchReader is a Reader for the ImageSearch structure.
type ImageSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageSearchOK creates a ImageSearchOK with default headers values
func NewImageSearchOK() *ImageSearchOK {
	return &ImageSearchOK{}
}

/*
ImageSearchOK describes a response with status code 200, with default header values.

Search results
*/
type ImageSearchOK struct {
	Payload *ImageSearchOKBody
}

// IsSuccess returns true when this image search o k response has a 2xx status code
func (o *ImageSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image search o k response has a 3xx status code
func (o *ImageSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image search o k response has a 4xx status code
func (o *ImageSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image search o k response has a 5xx status code
func (o *ImageSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image search o k response a status code equal to that given
func (o *ImageSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageSearchOK) Error() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchOK  %+v", 200, o.Payload)
}

func (o *ImageSearchOK) String() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchOK  %+v", 200, o.Payload)
}

func (o *ImageSearchOK) GetPayload() *ImageSearchOKBody {
	return o.Payload
}

func (o *ImageSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageSearchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageSearchBadRequest creates a ImageSearchBadRequest with default headers values
func NewImageSearchBadRequest() *ImageSearchBadRequest {
	return &ImageSearchBadRequest{}
}

/*
ImageSearchBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ImageSearchBadRequest struct {
	Payload *ImageSearchBadRequestBody
}

// IsSuccess returns true when this image search bad request response has a 2xx status code
func (o *ImageSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image search bad request response has a 3xx status code
func (o *ImageSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image search bad request response has a 4xx status code
func (o *ImageSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image search bad request response has a 5xx status code
func (o *ImageSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image search bad request response a status code equal to that given
func (o *ImageSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImageSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchBadRequest  %+v", 400, o.Payload)
}

func (o *ImageSearchBadRequest) String() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchBadRequest  %+v", 400, o.Payload)
}

func (o *ImageSearchBadRequest) GetPayload() *ImageSearchBadRequestBody {
	return o.Payload
}

func (o *ImageSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageSearchBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageSearchInternalServerError creates a ImageSearchInternalServerError with default headers values
func NewImageSearchInternalServerError() *ImageSearchInternalServerError {
	return &ImageSearchInternalServerError{}
}

/*
ImageSearchInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageSearchInternalServerError struct {
	Payload *ImageSearchInternalServerErrorBody
}

// IsSuccess returns true when this image search internal server error response has a 2xx status code
func (o *ImageSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image search internal server error response has a 3xx status code
func (o *ImageSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image search internal server error response has a 4xx status code
func (o *ImageSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image search internal server error response has a 5xx status code
func (o *ImageSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image search internal server error response a status code equal to that given
func (o *ImageSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageSearchInternalServerError) String() string {
	return fmt.Sprintf("[GET /images/search][%d] imageSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageSearchInternalServerError) GetPayload() *ImageSearchInternalServerErrorBody {
	return o.Payload
}

func (o *ImageSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageSearchInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageSearchBadRequestBody image search bad request body
swagger:model ImageSearchBadRequestBody
*/
type ImageSearchBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image search bad request body
func (o *ImageSearchBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image search bad request body based on context it is used
func (o *ImageSearchBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageSearchBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageSearchBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ImageSearchBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageSearchInternalServerErrorBody image search internal server error body
swagger:model ImageSearchInternalServerErrorBody
*/
type ImageSearchInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image search internal server error body
func (o *ImageSearchInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image search internal server error body based on context it is used
func (o *ImageSearchInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageSearchInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageSearchInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageSearchInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageSearchOKBody image search o k body
swagger:model ImageSearchOKBody
*/
type ImageSearchOKBody struct {

	// Automated indicates if the image was created by an automated build.
	Automated string `json:"Automated,omitempty"`

	// Description of the image.
	Description string `json:"Description,omitempty"`

	// Index is the image index (e.g., "docker.io" or "quay.io")
	Index string `json:"Index,omitempty"`

	// Name is the canonical name of the image (e.g., "docker.io/library/alpine").
	Name string `json:"Name,omitempty"`

	// Official indicates if it's an official image.
	Official string `json:"Official,omitempty"`

	// Stars is the number of stars of the image.
	Stars int64 `json:"Stars,omitempty"`

	// Tag is the image tag
	Tag string `json:"Tag,omitempty"`
}

// Validate validates this image search o k body
func (o *ImageSearchOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image search o k body based on context it is used
func (o *ImageSearchOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageSearchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageSearchOKBody) UnmarshalBinary(b []byte) error {
	var res ImageSearchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
