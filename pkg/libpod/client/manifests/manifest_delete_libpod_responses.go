// Code generated by go-swagger; DO NOT EDIT.

package manifests

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

// ManifestDeleteLibpodReader is a Reader for the ManifestDeleteLibpod structure.
type ManifestDeleteLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManifestDeleteLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManifestDeleteLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewManifestDeleteLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewManifestDeleteLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewManifestDeleteLibpodOK creates a ManifestDeleteLibpodOK with default headers values
func NewManifestDeleteLibpodOK() *ManifestDeleteLibpodOK {
	return &ManifestDeleteLibpodOK{}
}

/*
ManifestDeleteLibpodOK describes a response with status code 200, with default header values.

Remove response
*/
type ManifestDeleteLibpodOK struct {
	Payload *models.LibpodImagesRemoveReport
}

// IsSuccess returns true when this manifest delete libpod o k response has a 2xx status code
func (o *ManifestDeleteLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this manifest delete libpod o k response has a 3xx status code
func (o *ManifestDeleteLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest delete libpod o k response has a 4xx status code
func (o *ManifestDeleteLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this manifest delete libpod o k response has a 5xx status code
func (o *ManifestDeleteLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this manifest delete libpod o k response a status code equal to that given
func (o *ManifestDeleteLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ManifestDeleteLibpodOK) Error() string {
	return fmt.Sprintf("[DELETE /libpod/manifests/{name}][%d] manifestDeleteLibpodOK  %+v", 200, o.Payload)
}

func (o *ManifestDeleteLibpodOK) String() string {
	return fmt.Sprintf("[DELETE /libpod/manifests/{name}][%d] manifestDeleteLibpodOK  %+v", 200, o.Payload)
}

func (o *ManifestDeleteLibpodOK) GetPayload() *models.LibpodImagesRemoveReport {
	return o.Payload
}

func (o *ManifestDeleteLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LibpodImagesRemoveReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestDeleteLibpodNotFound creates a ManifestDeleteLibpodNotFound with default headers values
func NewManifestDeleteLibpodNotFound() *ManifestDeleteLibpodNotFound {
	return &ManifestDeleteLibpodNotFound{}
}

/*
ManifestDeleteLibpodNotFound describes a response with status code 404, with default header values.

No such manifest
*/
type ManifestDeleteLibpodNotFound struct {
	Payload *ManifestDeleteLibpodNotFoundBody
}

// IsSuccess returns true when this manifest delete libpod not found response has a 2xx status code
func (o *ManifestDeleteLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this manifest delete libpod not found response has a 3xx status code
func (o *ManifestDeleteLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest delete libpod not found response has a 4xx status code
func (o *ManifestDeleteLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this manifest delete libpod not found response has a 5xx status code
func (o *ManifestDeleteLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this manifest delete libpod not found response a status code equal to that given
func (o *ManifestDeleteLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ManifestDeleteLibpodNotFound) Error() string {
	return fmt.Sprintf("[DELETE /libpod/manifests/{name}][%d] manifestDeleteLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ManifestDeleteLibpodNotFound) String() string {
	return fmt.Sprintf("[DELETE /libpod/manifests/{name}][%d] manifestDeleteLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ManifestDeleteLibpodNotFound) GetPayload() *ManifestDeleteLibpodNotFoundBody {
	return o.Payload
}

func (o *ManifestDeleteLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ManifestDeleteLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestDeleteLibpodInternalServerError creates a ManifestDeleteLibpodInternalServerError with default headers values
func NewManifestDeleteLibpodInternalServerError() *ManifestDeleteLibpodInternalServerError {
	return &ManifestDeleteLibpodInternalServerError{}
}

/*
ManifestDeleteLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ManifestDeleteLibpodInternalServerError struct {
	Payload *ManifestDeleteLibpodInternalServerErrorBody
}

// IsSuccess returns true when this manifest delete libpod internal server error response has a 2xx status code
func (o *ManifestDeleteLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this manifest delete libpod internal server error response has a 3xx status code
func (o *ManifestDeleteLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest delete libpod internal server error response has a 4xx status code
func (o *ManifestDeleteLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this manifest delete libpod internal server error response has a 5xx status code
func (o *ManifestDeleteLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this manifest delete libpod internal server error response a status code equal to that given
func (o *ManifestDeleteLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ManifestDeleteLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /libpod/manifests/{name}][%d] manifestDeleteLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ManifestDeleteLibpodInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /libpod/manifests/{name}][%d] manifestDeleteLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ManifestDeleteLibpodInternalServerError) GetPayload() *ManifestDeleteLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ManifestDeleteLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ManifestDeleteLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ManifestDeleteLibpodInternalServerErrorBody manifest delete libpod internal server error body
swagger:model ManifestDeleteLibpodInternalServerErrorBody
*/
type ManifestDeleteLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this manifest delete libpod internal server error body
func (o *ManifestDeleteLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manifest delete libpod internal server error body based on context it is used
func (o *ManifestDeleteLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ManifestDeleteLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ManifestDeleteLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ManifestDeleteLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ManifestDeleteLibpodNotFoundBody manifest delete libpod not found body
swagger:model ManifestDeleteLibpodNotFoundBody
*/
type ManifestDeleteLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this manifest delete libpod not found body
func (o *ManifestDeleteLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manifest delete libpod not found body based on context it is used
func (o *ManifestDeleteLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ManifestDeleteLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ManifestDeleteLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ManifestDeleteLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
