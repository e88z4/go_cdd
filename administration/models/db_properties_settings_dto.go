// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DbPropertiesSettingsDto db properties settings dto
// swagger:model DbPropertiesSettingsDto
type DbPropertiesSettingsDto struct {

	// host
	Host string `json:"host,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this db properties settings dto
func (m *DbPropertiesSettingsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DbPropertiesSettingsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DbPropertiesSettingsDto) UnmarshalBinary(b []byte) error {
	var res DbPropertiesSettingsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
