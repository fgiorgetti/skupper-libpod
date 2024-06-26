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
)

// PutContainerArchiveLibpodReader is a Reader for the PutContainerArchiveLibpod structure.
type PutContainerArchiveLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContainerArchiveLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContainerArchiveLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutContainerArchiveLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutContainerArchiveLibpodForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutContainerArchiveLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutContainerArchiveLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutContainerArchiveLibpodOK creates a PutContainerArchiveLibpodOK with default headers values
func NewPutContainerArchiveLibpodOK() *PutContainerArchiveLibpodOK {
	return &PutContainerArchiveLibpodOK{}
}

/*
PutContainerArchiveLibpodOK describes a response with status code 200, with default header values.

no error
*/
type PutContainerArchiveLibpodOK struct {
}

// IsSuccess returns true when this put container archive libpod o k response has a 2xx status code
func (o *PutContainerArchiveLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put container archive libpod o k response has a 3xx status code
func (o *PutContainerArchiveLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive libpod o k response has a 4xx status code
func (o *PutContainerArchiveLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put container archive libpod o k response has a 5xx status code
func (o *PutContainerArchiveLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive libpod o k response a status code equal to that given
func (o *PutContainerArchiveLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutContainerArchiveLibpodOK) Error() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodOK ", 200)
}

func (o *PutContainerArchiveLibpodOK) String() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodOK ", 200)
}

func (o *PutContainerArchiveLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContainerArchiveLibpodBadRequest creates a PutContainerArchiveLibpodBadRequest with default headers values
func NewPutContainerArchiveLibpodBadRequest() *PutContainerArchiveLibpodBadRequest {
	return &PutContainerArchiveLibpodBadRequest{}
}

/*
PutContainerArchiveLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type PutContainerArchiveLibpodBadRequest struct {
	Payload *PutContainerArchiveLibpodBadRequestBody
}

// IsSuccess returns true when this put container archive libpod bad request response has a 2xx status code
func (o *PutContainerArchiveLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive libpod bad request response has a 3xx status code
func (o *PutContainerArchiveLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive libpod bad request response has a 4xx status code
func (o *PutContainerArchiveLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put container archive libpod bad request response has a 5xx status code
func (o *PutContainerArchiveLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive libpod bad request response a status code equal to that given
func (o *PutContainerArchiveLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutContainerArchiveLibpodBadRequest) Error() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *PutContainerArchiveLibpodBadRequest) String() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *PutContainerArchiveLibpodBadRequest) GetPayload() *PutContainerArchiveLibpodBadRequestBody {
	return o.Payload
}

func (o *PutContainerArchiveLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutContainerArchiveLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveLibpodForbidden creates a PutContainerArchiveLibpodForbidden with default headers values
func NewPutContainerArchiveLibpodForbidden() *PutContainerArchiveLibpodForbidden {
	return &PutContainerArchiveLibpodForbidden{}
}

/*
PutContainerArchiveLibpodForbidden describes a response with status code 403, with default header values.

the container rootfs is read-only
*/
type PutContainerArchiveLibpodForbidden struct {
}

// IsSuccess returns true when this put container archive libpod forbidden response has a 2xx status code
func (o *PutContainerArchiveLibpodForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive libpod forbidden response has a 3xx status code
func (o *PutContainerArchiveLibpodForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive libpod forbidden response has a 4xx status code
func (o *PutContainerArchiveLibpodForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put container archive libpod forbidden response has a 5xx status code
func (o *PutContainerArchiveLibpodForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive libpod forbidden response a status code equal to that given
func (o *PutContainerArchiveLibpodForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutContainerArchiveLibpodForbidden) Error() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodForbidden ", 403)
}

func (o *PutContainerArchiveLibpodForbidden) String() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodForbidden ", 403)
}

func (o *PutContainerArchiveLibpodForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContainerArchiveLibpodNotFound creates a PutContainerArchiveLibpodNotFound with default headers values
func NewPutContainerArchiveLibpodNotFound() *PutContainerArchiveLibpodNotFound {
	return &PutContainerArchiveLibpodNotFound{}
}

/*
PutContainerArchiveLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type PutContainerArchiveLibpodNotFound struct {
	Payload *PutContainerArchiveLibpodNotFoundBody
}

// IsSuccess returns true when this put container archive libpod not found response has a 2xx status code
func (o *PutContainerArchiveLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive libpod not found response has a 3xx status code
func (o *PutContainerArchiveLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive libpod not found response has a 4xx status code
func (o *PutContainerArchiveLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put container archive libpod not found response has a 5xx status code
func (o *PutContainerArchiveLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive libpod not found response a status code equal to that given
func (o *PutContainerArchiveLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutContainerArchiveLibpodNotFound) Error() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodNotFound  %+v", 404, o.Payload)
}

func (o *PutContainerArchiveLibpodNotFound) String() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodNotFound  %+v", 404, o.Payload)
}

func (o *PutContainerArchiveLibpodNotFound) GetPayload() *PutContainerArchiveLibpodNotFoundBody {
	return o.Payload
}

func (o *PutContainerArchiveLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutContainerArchiveLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveLibpodInternalServerError creates a PutContainerArchiveLibpodInternalServerError with default headers values
func NewPutContainerArchiveLibpodInternalServerError() *PutContainerArchiveLibpodInternalServerError {
	return &PutContainerArchiveLibpodInternalServerError{}
}

/*
PutContainerArchiveLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutContainerArchiveLibpodInternalServerError struct {
	Payload *PutContainerArchiveLibpodInternalServerErrorBody
}

// IsSuccess returns true when this put container archive libpod internal server error response has a 2xx status code
func (o *PutContainerArchiveLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive libpod internal server error response has a 3xx status code
func (o *PutContainerArchiveLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive libpod internal server error response has a 4xx status code
func (o *PutContainerArchiveLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put container archive libpod internal server error response has a 5xx status code
func (o *PutContainerArchiveLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put container archive libpod internal server error response a status code equal to that given
func (o *PutContainerArchiveLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutContainerArchiveLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContainerArchiveLibpodInternalServerError) String() string {
	return fmt.Sprintf("[PUT /libpod/containers/{name}/archive][%d] putContainerArchiveLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContainerArchiveLibpodInternalServerError) GetPayload() *PutContainerArchiveLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PutContainerArchiveLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutContainerArchiveLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutContainerArchiveLibpodBadRequestBody put container archive libpod bad request body
swagger:model PutContainerArchiveLibpodBadRequestBody
*/
type PutContainerArchiveLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive libpod bad request body
func (o *PutContainerArchiveLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive libpod bad request body based on context it is used
func (o *PutContainerArchiveLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutContainerArchiveLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutContainerArchiveLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutContainerArchiveLibpodInternalServerErrorBody put container archive libpod internal server error body
swagger:model PutContainerArchiveLibpodInternalServerErrorBody
*/
type PutContainerArchiveLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive libpod internal server error body
func (o *PutContainerArchiveLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive libpod internal server error body based on context it is used
func (o *PutContainerArchiveLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutContainerArchiveLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutContainerArchiveLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutContainerArchiveLibpodNotFoundBody put container archive libpod not found body
swagger:model PutContainerArchiveLibpodNotFoundBody
*/
type PutContainerArchiveLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive libpod not found body
func (o *PutContainerArchiveLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive libpod not found body based on context it is used
func (o *PutContainerArchiveLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutContainerArchiveLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutContainerArchiveLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
