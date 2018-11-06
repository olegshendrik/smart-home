// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AddNodeOKBodyData add node o k body data
// swagger:model addNodeOKBodyData
type AddNodeOKBodyData struct {

	// node id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this add node o k body data
func (m *AddNodeOKBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AddNodeOKBodyData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddNodeOKBodyData) UnmarshalBinary(b []byte) error {
	var res AddNodeOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
