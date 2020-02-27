// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BasicPluginServiceParameterDtoObject basic plugin service parameter dto object
// swagger:model BasicPluginServiceParameterDto«object»
type BasicPluginServiceParameterDtoObject struct {

	// class name
	ClassName string `json:"className,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// template parameter Id
	TemplateParameterID int64 `json:"templateParameterId,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`

	// value Id
	ValueID string `json:"valueId,omitempty"`
}

// Validate validates this basic plugin service parameter dto object
func (m *BasicPluginServiceParameterDtoObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BasicPluginServiceParameterDtoObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasicPluginServiceParameterDtoObject) UnmarshalBinary(b []byte) error {
	var res BasicPluginServiceParameterDtoObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}