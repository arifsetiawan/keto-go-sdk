// Code generated by go-swagger; DO NOT EDIT.

package engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/arifsetiawan/keto-go-sdk/models"
)

// GetOryAccessControlPolicyRoleReader is a Reader for the GetOryAccessControlPolicyRole structure.
type GetOryAccessControlPolicyRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOryAccessControlPolicyRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOryAccessControlPolicyRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetOryAccessControlPolicyRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOryAccessControlPolicyRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOryAccessControlPolicyRoleOK creates a GetOryAccessControlPolicyRoleOK with default headers values
func NewGetOryAccessControlPolicyRoleOK() *GetOryAccessControlPolicyRoleOK {
	return &GetOryAccessControlPolicyRoleOK{}
}

/*GetOryAccessControlPolicyRoleOK handles this case with default header values.

oryAccessControlPolicyRole
*/
type GetOryAccessControlPolicyRoleOK struct {
	Payload *models.OryAccessControlPolicyRole
}

func (o *GetOryAccessControlPolicyRoleOK) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/roles/{id}][%d] getOryAccessControlPolicyRoleOK  %+v", 200, o.Payload)
}

func (o *GetOryAccessControlPolicyRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OryAccessControlPolicyRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOryAccessControlPolicyRoleNotFound creates a GetOryAccessControlPolicyRoleNotFound with default headers values
func NewGetOryAccessControlPolicyRoleNotFound() *GetOryAccessControlPolicyRoleNotFound {
	return &GetOryAccessControlPolicyRoleNotFound{}
}

/*GetOryAccessControlPolicyRoleNotFound handles this case with default header values.

The standard error format
*/
type GetOryAccessControlPolicyRoleNotFound struct {
	Payload *GetOryAccessControlPolicyRoleNotFoundBody
}

func (o *GetOryAccessControlPolicyRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/roles/{id}][%d] getOryAccessControlPolicyRoleNotFound  %+v", 404, o.Payload)
}

func (o *GetOryAccessControlPolicyRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOryAccessControlPolicyRoleNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOryAccessControlPolicyRoleInternalServerError creates a GetOryAccessControlPolicyRoleInternalServerError with default headers values
func NewGetOryAccessControlPolicyRoleInternalServerError() *GetOryAccessControlPolicyRoleInternalServerError {
	return &GetOryAccessControlPolicyRoleInternalServerError{}
}

/*GetOryAccessControlPolicyRoleInternalServerError handles this case with default header values.

The standard error format
*/
type GetOryAccessControlPolicyRoleInternalServerError struct {
	Payload *GetOryAccessControlPolicyRoleInternalServerErrorBody
}

func (o *GetOryAccessControlPolicyRoleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/roles/{id}][%d] getOryAccessControlPolicyRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOryAccessControlPolicyRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOryAccessControlPolicyRoleInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOryAccessControlPolicyRoleInternalServerErrorBody get ory access control policy role internal server error body
swagger:model GetOryAccessControlPolicyRoleInternalServerErrorBody
*/
type GetOryAccessControlPolicyRoleInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get ory access control policy role internal server error body
func (o *GetOryAccessControlPolicyRoleInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOryAccessControlPolicyRoleInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOryAccessControlPolicyRoleInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetOryAccessControlPolicyRoleInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOryAccessControlPolicyRoleNotFoundBody get ory access control policy role not found body
swagger:model GetOryAccessControlPolicyRoleNotFoundBody
*/
type GetOryAccessControlPolicyRoleNotFoundBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get ory access control policy role not found body
func (o *GetOryAccessControlPolicyRoleNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOryAccessControlPolicyRoleNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOryAccessControlPolicyRoleNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetOryAccessControlPolicyRoleNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
