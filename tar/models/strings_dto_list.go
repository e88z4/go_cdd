// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StringsDtoList strings dto list
// swagger:model StringsDtoList
type StringsDtoList struct {

	// data
	Data []string `json:"data"`
}

// Validate validates this strings dto list
func (m *StringsDtoList) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StringsDtoList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StringsDtoList) UnmarshalBinary(b []byte) error {
	var res StringsDtoList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
