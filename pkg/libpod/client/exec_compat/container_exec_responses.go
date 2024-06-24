// Code generated by go-swagger; DO NOT EDIT.

package exec_compat

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

// ContainerExecReader is a Reader for the ContainerExec structure.
type ContainerExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewContainerExecCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerExecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerExecConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerExecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerExecCreated creates a ContainerExecCreated with default headers values
func NewContainerExecCreated() *ContainerExecCreated {
	return &ContainerExecCreated{}
}

/*
ContainerExecCreated describes a response with status code 201, with default header values.

no error
*/
type ContainerExecCreated struct {
}

// IsSuccess returns true when this container exec created response has a 2xx status code
func (o *ContainerExecCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container exec created response has a 3xx status code
func (o *ContainerExecCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container exec created response has a 4xx status code
func (o *ContainerExecCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this container exec created response has a 5xx status code
func (o *ContainerExecCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this container exec created response a status code equal to that given
func (o *ContainerExecCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ContainerExecCreated) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecCreated ", 201)
}

func (o *ContainerExecCreated) String() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecCreated ", 201)
}

func (o *ContainerExecCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerExecNotFound creates a ContainerExecNotFound with default headers values
func NewContainerExecNotFound() *ContainerExecNotFound {
	return &ContainerExecNotFound{}
}

/*
ContainerExecNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerExecNotFound struct {
	Payload *ContainerExecNotFoundBody
}

// IsSuccess returns true when this container exec not found response has a 2xx status code
func (o *ContainerExecNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container exec not found response has a 3xx status code
func (o *ContainerExecNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container exec not found response has a 4xx status code
func (o *ContainerExecNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container exec not found response has a 5xx status code
func (o *ContainerExecNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container exec not found response a status code equal to that given
func (o *ContainerExecNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ContainerExecNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecNotFound  %+v", 404, o.Payload)
}

func (o *ContainerExecNotFound) String() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecNotFound  %+v", 404, o.Payload)
}

func (o *ContainerExecNotFound) GetPayload() *ContainerExecNotFoundBody {
	return o.Payload
}

func (o *ContainerExecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerExecNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExecConflict creates a ContainerExecConflict with default headers values
func NewContainerExecConflict() *ContainerExecConflict {
	return &ContainerExecConflict{}
}

/*
ContainerExecConflict describes a response with status code 409, with default header values.

container is paused
*/
type ContainerExecConflict struct {
}

// IsSuccess returns true when this container exec conflict response has a 2xx status code
func (o *ContainerExecConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container exec conflict response has a 3xx status code
func (o *ContainerExecConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container exec conflict response has a 4xx status code
func (o *ContainerExecConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this container exec conflict response has a 5xx status code
func (o *ContainerExecConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this container exec conflict response a status code equal to that given
func (o *ContainerExecConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ContainerExecConflict) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecConflict ", 409)
}

func (o *ContainerExecConflict) String() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecConflict ", 409)
}

func (o *ContainerExecConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerExecInternalServerError creates a ContainerExecInternalServerError with default headers values
func NewContainerExecInternalServerError() *ContainerExecInternalServerError {
	return &ContainerExecInternalServerError{}
}

/*
ContainerExecInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerExecInternalServerError struct {
	Payload *ContainerExecInternalServerErrorBody
}

// IsSuccess returns true when this container exec internal server error response has a 2xx status code
func (o *ContainerExecInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container exec internal server error response has a 3xx status code
func (o *ContainerExecInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container exec internal server error response has a 4xx status code
func (o *ContainerExecInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this container exec internal server error response has a 5xx status code
func (o *ContainerExecInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this container exec internal server error response a status code equal to that given
func (o *ContainerExecInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ContainerExecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerExecInternalServerError) String() string {
	return fmt.Sprintf("[POST /containers/{name}/exec][%d] containerExecInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerExecInternalServerError) GetPayload() *ContainerExecInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerExecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerExecInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ContainerExecBody container exec body
swagger:model ContainerExecBody
*/
type ContainerExecBody struct {

	// Attach to stderr of the exec command
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// Attach to stdin of the exec command
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// Attach to stdout of the exec command
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// Command to run, as a string or array of strings.
	Cmd []string `json:"Cmd"`

	// "Override the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _."
	//
	DetachKeys string `json:"DetachKeys,omitempty"`

	// A list of environment variables in the form ["VAR=value", ...]
	Env []string `json:"Env"`

	// Runs the exec process with extended privileges
	Privileged *bool `json:"Privileged,omitempty"`

	// Allocate a pseudo-TTY
	Tty bool `json:"Tty,omitempty"`

	// "The user, and optionally, group to run the exec process inside the container. Format is one of: user, user:group, uid, or uid:gid."
	//
	User string `json:"User,omitempty"`

	// The working directory for the exec process inside the container.
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this container exec body
func (o *ContainerExecBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container exec body based on context it is used
func (o *ContainerExecBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExecBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExecBody) UnmarshalBinary(b []byte) error {
	var res ContainerExecBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerExecInternalServerErrorBody container exec internal server error body
swagger:model ContainerExecInternalServerErrorBody
*/
type ContainerExecInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container exec internal server error body
func (o *ContainerExecInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container exec internal server error body based on context it is used
func (o *ContainerExecInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExecInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExecInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerExecInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ContainerExecNotFoundBody container exec not found body
swagger:model ContainerExecNotFoundBody
*/
type ContainerExecNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container exec not found body
func (o *ContainerExecNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container exec not found body based on context it is used
func (o *ContainerExecNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExecNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExecNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerExecNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
