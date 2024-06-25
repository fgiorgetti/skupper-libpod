// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fgiorgetti/skupper-libpod/models"
)

// NetworkCreateReader is a Reader for the NetworkCreate structure.
type NetworkCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNetworkCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkCreateOK creates a NetworkCreateOK with default headers values
func NewNetworkCreateOK() *NetworkCreateOK {
	return &NetworkCreateOK{}
}

/*
NetworkCreateOK describes a response with status code 200, with default header values.

Network create
*/
type NetworkCreateOK struct {
	Payload *NetworkCreateOKBody
}

// IsSuccess returns true when this network create o k response has a 2xx status code
func (o *NetworkCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network create o k response has a 3xx status code
func (o *NetworkCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create o k response has a 4xx status code
func (o *NetworkCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create o k response has a 5xx status code
func (o *NetworkCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network create o k response a status code equal to that given
func (o *NetworkCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkCreateOK) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateOK  %+v", 200, o.Payload)
}

func (o *NetworkCreateOK) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateOK  %+v", 200, o.Payload)
}

func (o *NetworkCreateOK) GetPayload() *NetworkCreateOKBody {
	return o.Payload
}

func (o *NetworkCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateBadRequest creates a NetworkCreateBadRequest with default headers values
func NewNetworkCreateBadRequest() *NetworkCreateBadRequest {
	return &NetworkCreateBadRequest{}
}

/*
NetworkCreateBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type NetworkCreateBadRequest struct {
	Payload *NetworkCreateBadRequestBody
}

// IsSuccess returns true when this network create bad request response has a 2xx status code
func (o *NetworkCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create bad request response has a 3xx status code
func (o *NetworkCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create bad request response has a 4xx status code
func (o *NetworkCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this network create bad request response has a 5xx status code
func (o *NetworkCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this network create bad request response a status code equal to that given
func (o *NetworkCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NetworkCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateBadRequest  %+v", 400, o.Payload)
}

func (o *NetworkCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateBadRequest  %+v", 400, o.Payload)
}

func (o *NetworkCreateBadRequest) GetPayload() *NetworkCreateBadRequestBody {
	return o.Payload
}

func (o *NetworkCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkCreateInternalServerError creates a NetworkCreateInternalServerError with default headers values
func NewNetworkCreateInternalServerError() *NetworkCreateInternalServerError {
	return &NetworkCreateInternalServerError{}
}

/*
NetworkCreateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkCreateInternalServerError struct {
	Payload *NetworkCreateInternalServerErrorBody
}

// IsSuccess returns true when this network create internal server error response has a 2xx status code
func (o *NetworkCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this network create internal server error response has a 3xx status code
func (o *NetworkCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network create internal server error response has a 4xx status code
func (o *NetworkCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this network create internal server error response has a 5xx status code
func (o *NetworkCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this network create internal server error response a status code equal to that given
func (o *NetworkCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NetworkCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /networks/create][%d] networkCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *NetworkCreateInternalServerError) GetPayload() *NetworkCreateInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkCreateInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
NetworkCreateBadRequestBody network create bad request body
swagger:model NetworkCreateBadRequestBody
*/
type NetworkCreateBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network create bad request body
func (o *NetworkCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network create bad request body based on context it is used
func (o *NetworkCreateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkCreateInternalServerErrorBody network create internal server error body
swagger:model NetworkCreateInternalServerErrorBody
*/
type NetworkCreateInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network create internal server error body
func (o *NetworkCreateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network create internal server error body based on context it is used
func (o *NetworkCreateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
NetworkCreateOKBody network create o k body
swagger:model NetworkCreateOKBody
*/
type NetworkCreateOKBody struct {

	// attachable
	Attachable bool `json:"Attachable,omitempty"`

	// Check for networks with duplicate names.
	// Network is primarily keyed based on a random ID and not on the name.
	// Network name is strictly a user-friendly alias to the network
	// which is uniquely identified using ID.
	// And there is no guaranteed way to check for duplicates.
	// Option CheckDuplicate is there to provide a best effort checking of any networks
	// which has the same name but it is not guaranteed to catch all name collisions.
	CheckDuplicate bool `json:"CheckDuplicate,omitempty"`

	// config from
	ConfigFrom *models.ConfigReference `json:"ConfigFrom,omitempty"`

	// config only
	ConfigOnly bool `json:"ConfigOnly,omitempty"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// enable IPv6
	EnableIPV6 bool `json:"EnableIPv6,omitempty"`

	// IP a m
	IPAM *models.IPAM `json:"IPAM,omitempty"`

	// ingress
	Ingress bool `json:"Ingress,omitempty"`

	// internal
	Internal bool `json:"Internal,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// options
	Options map[string]string `json:"Options,omitempty"`

	// scope
	Scope string `json:"Scope,omitempty"`
}

// Validate validates this network create o k body
func (o *NetworkCreateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfigFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIPAM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkCreateOKBody) validateConfigFrom(formats strfmt.Registry) error {
	if swag.IsZero(o.ConfigFrom) { // not required
		return nil
	}

	if o.ConfigFrom != nil {
		if err := o.ConfigFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkCreateOK" + "." + "ConfigFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkCreateOK" + "." + "ConfigFrom")
			}
			return err
		}
	}

	return nil
}

func (o *NetworkCreateOKBody) validateIPAM(formats strfmt.Registry) error {
	if swag.IsZero(o.IPAM) { // not required
		return nil
	}

	if o.IPAM != nil {
		if err := o.IPAM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkCreateOK" + "." + "IPAM")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkCreateOK" + "." + "IPAM")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network create o k body based on the context it is used
func (o *NetworkCreateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfigFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIPAM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkCreateOKBody) contextValidateConfigFrom(ctx context.Context, formats strfmt.Registry) error {

	if o.ConfigFrom != nil {
		if err := o.ConfigFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkCreateOK" + "." + "ConfigFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkCreateOK" + "." + "ConfigFrom")
			}
			return err
		}
	}

	return nil
}

func (o *NetworkCreateOKBody) contextValidateIPAM(ctx context.Context, formats strfmt.Registry) error {

	if o.IPAM != nil {
		if err := o.IPAM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkCreateOK" + "." + "IPAM")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkCreateOK" + "." + "IPAM")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkCreateOKBody) UnmarshalBinary(b []byte) error {
	var res NetworkCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}