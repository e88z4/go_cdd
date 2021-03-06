// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DslWarningDto dsl warning dto
// swagger:model DslWarningDto
type DslWarningDto struct {

	// entity name
	EntityName string `json:"entityName,omitempty"`

	// message code
	MessageCode string `json:"messageCode,omitempty"`

	// names
	Names []string `json:"names"`
}

// Validate validates this dsl warning dto
func (m *DslWarningDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DslWarningDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DslWarningDto) UnmarshalBinary(b []byte) error {
	var res DslWarningDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
