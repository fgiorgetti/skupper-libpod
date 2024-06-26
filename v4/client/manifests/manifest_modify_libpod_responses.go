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

	"github.com/fgiorgetti/skupper-libpod/v4/models"
)

// ManifestModifyLibpodReader is a Reader for the ManifestModifyLibpod structure.
type ManifestModifyLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ManifestModifyLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewManifestModifyLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewManifestModifyLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewManifestModifyLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewManifestModifyLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewManifestModifyLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewManifestModifyLibpodOK creates a ManifestModifyLibpodOK with default headers values
func NewManifestModifyLibpodOK() *ManifestModifyLibpodOK {
	return &ManifestModifyLibpodOK{}
}

/*
ManifestModifyLibpodOK describes a response with status code 200, with default header values.

ManifestModifyLibpodOK manifest modify libpod o k
*/
type ManifestModifyLibpodOK struct {
	Payload *models.ManifestModifyReport
}

// IsSuccess returns true when this manifest modify libpod o k response has a 2xx status code
func (o *ManifestModifyLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this manifest modify libpod o k response has a 3xx status code
func (o *ManifestModifyLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest modify libpod o k response has a 4xx status code
func (o *ManifestModifyLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this manifest modify libpod o k response has a 5xx status code
func (o *ManifestModifyLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this manifest modify libpod o k response a status code equal to that given
func (o *ManifestModifyLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ManifestModifyLibpodOK) Error() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodOK  %+v", 200, o.Payload)
}

func (o *ManifestModifyLibpodOK) String() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodOK  %+v", 200, o.Payload)
}

func (o *ManifestModifyLibpodOK) GetPayload() *models.ManifestModifyReport {
	return o.Payload
}

func (o *ManifestModifyLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManifestModifyReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestModifyLibpodBadRequest creates a ManifestModifyLibpodBadRequest with default headers values
func NewManifestModifyLibpodBadRequest() *ManifestModifyLibpodBadRequest {
	return &ManifestModifyLibpodBadRequest{}
}

/*
ManifestModifyLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ManifestModifyLibpodBadRequest struct {
	Payload *ManifestModifyLibpodBadRequestBody
}

// IsSuccess returns true when this manifest modify libpod bad request response has a 2xx status code
func (o *ManifestModifyLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this manifest modify libpod bad request response has a 3xx status code
func (o *ManifestModifyLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest modify libpod bad request response has a 4xx status code
func (o *ManifestModifyLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this manifest modify libpod bad request response has a 5xx status code
func (o *ManifestModifyLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this manifest modify libpod bad request response a status code equal to that given
func (o *ManifestModifyLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ManifestModifyLibpodBadRequest) Error() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ManifestModifyLibpodBadRequest) String() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ManifestModifyLibpodBadRequest) GetPayload() *ManifestModifyLibpodBadRequestBody {
	return o.Payload
}

func (o *ManifestModifyLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ManifestModifyLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestModifyLibpodNotFound creates a ManifestModifyLibpodNotFound with default headers values
func NewManifestModifyLibpodNotFound() *ManifestModifyLibpodNotFound {
	return &ManifestModifyLibpodNotFound{}
}

/*
ManifestModifyLibpodNotFound describes a response with status code 404, with default header values.

No such manifest
*/
type ManifestModifyLibpodNotFound struct {
	Payload *ManifestModifyLibpodNotFoundBody
}

// IsSuccess returns true when this manifest modify libpod not found response has a 2xx status code
func (o *ManifestModifyLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this manifest modify libpod not found response has a 3xx status code
func (o *ManifestModifyLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest modify libpod not found response has a 4xx status code
func (o *ManifestModifyLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this manifest modify libpod not found response has a 5xx status code
func (o *ManifestModifyLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this manifest modify libpod not found response a status code equal to that given
func (o *ManifestModifyLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ManifestModifyLibpodNotFound) Error() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ManifestModifyLibpodNotFound) String() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodNotFound  %+v", 404, o.Payload)
}

func (o *ManifestModifyLibpodNotFound) GetPayload() *ManifestModifyLibpodNotFoundBody {
	return o.Payload
}

func (o *ManifestModifyLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ManifestModifyLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestModifyLibpodConflict creates a ManifestModifyLibpodConflict with default headers values
func NewManifestModifyLibpodConflict() *ManifestModifyLibpodConflict {
	return &ManifestModifyLibpodConflict{}
}

/*
ManifestModifyLibpodConflict describes a response with status code 409, with default header values.

Operation had partial success, both Images and Errors may have members
*/
type ManifestModifyLibpodConflict struct {
	Payload *models.ManifestModifyReport
}

// IsSuccess returns true when this manifest modify libpod conflict response has a 2xx status code
func (o *ManifestModifyLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this manifest modify libpod conflict response has a 3xx status code
func (o *ManifestModifyLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest modify libpod conflict response has a 4xx status code
func (o *ManifestModifyLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this manifest modify libpod conflict response has a 5xx status code
func (o *ManifestModifyLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this manifest modify libpod conflict response a status code equal to that given
func (o *ManifestModifyLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ManifestModifyLibpodConflict) Error() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodConflict  %+v", 409, o.Payload)
}

func (o *ManifestModifyLibpodConflict) String() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodConflict  %+v", 409, o.Payload)
}

func (o *ManifestModifyLibpodConflict) GetPayload() *models.ManifestModifyReport {
	return o.Payload
}

func (o *ManifestModifyLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManifestModifyReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewManifestModifyLibpodInternalServerError creates a ManifestModifyLibpodInternalServerError with default headers values
func NewManifestModifyLibpodInternalServerError() *ManifestModifyLibpodInternalServerError {
	return &ManifestModifyLibpodInternalServerError{}
}

/*
ManifestModifyLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ManifestModifyLibpodInternalServerError struct {
	Payload *ManifestModifyLibpodInternalServerErrorBody
}

// IsSuccess returns true when this manifest modify libpod internal server error response has a 2xx status code
func (o *ManifestModifyLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this manifest modify libpod internal server error response has a 3xx status code
func (o *ManifestModifyLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this manifest modify libpod internal server error response has a 4xx status code
func (o *ManifestModifyLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this manifest modify libpod internal server error response has a 5xx status code
func (o *ManifestModifyLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this manifest modify libpod internal server error response a status code equal to that given
func (o *ManifestModifyLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ManifestModifyLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ManifestModifyLibpodInternalServerError) String() string {
	return fmt.Sprintf("[PUT /libpod/manifests/{name}][%d] manifestModifyLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ManifestModifyLibpodInternalServerError) GetPayload() *ManifestModifyLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ManifestModifyLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ManifestModifyLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ManifestModifyLibpodBadRequestBody manifest modify libpod bad request body
swagger:model ManifestModifyLibpodBadRequestBody
*/
type ManifestModifyLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this manifest modify libpod bad request body
func (o *ManifestModifyLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manifest modify libpod bad request body based on context it is used
func (o *ManifestModifyLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ManifestModifyLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ManifestModifyLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ManifestModifyLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ManifestModifyLibpodInternalServerErrorBody manifest modify libpod internal server error body
swagger:model ManifestModifyLibpodInternalServerErrorBody
*/
type ManifestModifyLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this manifest modify libpod internal server error body
func (o *ManifestModifyLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manifest modify libpod internal server error body based on context it is used
func (o *ManifestModifyLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ManifestModifyLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ManifestModifyLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ManifestModifyLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ManifestModifyLibpodNotFoundBody manifest modify libpod not found body
swagger:model ManifestModifyLibpodNotFoundBody
*/
type ManifestModifyLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this manifest modify libpod not found body
func (o *ManifestModifyLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this manifest modify libpod not found body based on context it is used
func (o *ManifestModifyLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ManifestModifyLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ManifestModifyLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ManifestModifyLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}