// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegisteredPluginServerDto registered plugin server dto
// swagger:model RegisteredPluginServerDto
type RegisteredPluginServerDto struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// plugin server address
	PluginServerAddress string `json:"pluginServerAddress,omitempty"`

	// plugin server port
	PluginServerPort int32 `json:"pluginServerPort,omitempty"`

	// plugin server schema
	PluginServerSchema string `json:"pluginServerSchema,omitempty"`
}

// Validate validates this registered plugin server dto
func (m *RegisteredPluginServerDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisteredPluginServerDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisteredPluginServerDto) UnmarshalBinary(b []byte) error {
	var res RegisteredPluginServerDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}