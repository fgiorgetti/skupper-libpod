// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// PodListLibpodReader is a Reader for the PodListLibpod structure.
type PodListLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodListLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodListLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPodListLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodListLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPodListLibpodOK creates a PodListLibpodOK with default headers values
func NewPodListLibpodOK() *PodListLibpodOK {
	return &PodListLibpodOK{}
}

/*
PodListLibpodOK describes a response with status code 200, with default header values.

List pods
*/
type PodListLibpodOK struct {
	Payload []*models.ListPodsReport
}

// IsSuccess returns true when this pod list libpod o k response has a 2xx status code
func (o *PodListLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pod list libpod o k response has a 3xx status code
func (o *PodListLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod list libpod o k response has a 4xx status code
func (o *PodListLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod list libpod o k response has a 5xx status code
func (o *PodListLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pod list libpod o k response a status code equal to that given
func (o *PodListLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *PodListLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/pods/json][%d] podListLibpodOK  %+v", 200, o.Payload)
}

func (o *PodListLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/pods/json][%d] podListLibpodOK  %+v", 200, o.Payload)
}

func (o *PodListLibpodOK) GetPayload() []*models.ListPodsReport {
	return o.Payload
}

func (o *PodListLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodListLibpodBadRequest creates a PodListLibpodBadRequest with default headers values
func NewPodListLibpodBadRequest() *PodListLibpodBadRequest {
	return &PodListLibpodBadRequest{}
}

/*
PodListLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type PodListLibpodBadRequest struct {
	Payload *PodListLibpodBadRequestBody
}

// IsSuccess returns true when this pod list libpod bad request response has a 2xx status code
func (o *PodListLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod list libpod bad request response has a 3xx status code
func (o *PodListLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod list libpod bad request response has a 4xx status code
func (o *PodListLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pod list libpod bad request response has a 5xx status code
func (o *PodListLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pod list libpod bad request response a status code equal to that given
func (o *PodListLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PodListLibpodBadRequest) Error() string {
	return fmt.Sprintf("[GET /libpod/pods/json][%d] podListLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *PodListLibpodBadRequest) String() string {
	return fmt.Sprintf("[GET /libpod/pods/json][%d] podListLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *PodListLibpodBadRequest) GetPayload() *PodListLibpodBadRequestBody {
	return o.Payload
}

func (o *PodListLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodListLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodListLibpodInternalServerError creates a PodListLibpodInternalServerError with default headers values
func NewPodListLibpodInternalServerError() *PodListLibpodInternalServerError {
	return &PodListLibpodInternalServerError{}
}

/*
PodListLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PodListLibpodInternalServerError struct {
	Payload *PodListLibpodInternalServerErrorBody
}

// IsSuccess returns true when this pod list libpod internal server error response has a 2xx status code
func (o *PodListLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod list libpod internal server error response has a 3xx status code
func (o *PodListLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod list libpod internal server error response has a 4xx status code
func (o *PodListLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod list libpod internal server error response has a 5xx status code
func (o *PodListLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pod list libpod internal server error response a status code equal to that given
func (o *PodListLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PodListLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/pods/json][%d] podListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *PodListLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/pods/json][%d] podListLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *PodListLibpodInternalServerError) GetPayload() *PodListLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PodListLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodListLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PodListLibpodBadRequestBody pod list libpod bad request body
swagger:model PodListLibpodBadRequestBody
*/
type PodListLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod list libpod bad request body
func (o *PodListLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod list libpod bad request body based on context it is used
func (o *PodListLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodListLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodListLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PodListLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PodListLibpodInternalServerErrorBody pod list libpod internal server error body
swagger:model PodListLibpodInternalServerErrorBody
*/
type PodListLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod list libpod internal server error body
func (o *PodListLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod list libpod internal server error body based on context it is used
func (o *PodListLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodListLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodListLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PodListLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
