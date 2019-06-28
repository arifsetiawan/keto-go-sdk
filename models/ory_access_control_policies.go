// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OryAccessControlPolicies Policies is an array of policies.
// swagger:model oryAccessControlPolicies
type OryAccessControlPolicies struct {

	// The request body.
	//
	// in: body
	// type: array
	Body []*OryAccessControlPolicy `json:"Body"`
}

// Validate validates this ory access control policies
func (m *OryAccessControlPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OryAccessControlPolicies) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	for i := 0; i < len(m.Body); i++ {
		if swag.IsZero(m.Body[i]) { // not required
			continue
		}

		if m.Body[i] != nil {
			if err := m.Body[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Body" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OryAccessControlPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OryAccessControlPolicies) UnmarshalBinary(b []byte) error {
	var res OryAccessControlPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
