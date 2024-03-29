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

// ListOryAccessControlPoliciesReader is a Reader for the ListOryAccessControlPolicies structure.
type ListOryAccessControlPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOryAccessControlPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOryAccessControlPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewListOryAccessControlPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListOryAccessControlPoliciesOK creates a ListOryAccessControlPoliciesOK with default headers values
func NewListOryAccessControlPoliciesOK() *ListOryAccessControlPoliciesOK {
	return &ListOryAccessControlPoliciesOK{}
}

/*ListOryAccessControlPoliciesOK handles this case with default header values.

Policies is an array of policies.
*/
type ListOryAccessControlPoliciesOK struct {
	Payload []*models.OryAccessControlPolicy
}

func (o *ListOryAccessControlPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/policies][%d] listOryAccessControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListOryAccessControlPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOryAccessControlPoliciesInternalServerError creates a ListOryAccessControlPoliciesInternalServerError with default headers values
func NewListOryAccessControlPoliciesInternalServerError() *ListOryAccessControlPoliciesInternalServerError {
	return &ListOryAccessControlPoliciesInternalServerError{}
}

/*ListOryAccessControlPoliciesInternalServerError handles this case with default header values.

The standard error format
*/
type ListOryAccessControlPoliciesInternalServerError struct {
	Payload *ListOryAccessControlPoliciesInternalServerErrorBody
}

func (o *ListOryAccessControlPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /engines/acp/ory/{flavor}/policies][%d] listOryAccessControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListOryAccessControlPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListOryAccessControlPoliciesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListOryAccessControlPoliciesInternalServerErrorBody list ory access control policies internal server error body
swagger:model ListOryAccessControlPoliciesInternalServerErrorBody
*/
type ListOryAccessControlPoliciesInternalServerErrorBody struct {

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

// Validate validates this list ory access control policies internal server error body
func (o *ListOryAccessControlPoliciesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOryAccessControlPoliciesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOryAccessControlPoliciesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ListOryAccessControlPoliciesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
