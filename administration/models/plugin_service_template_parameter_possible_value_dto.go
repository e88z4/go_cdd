// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PluginServiceTemplateParameterPossibleValueDto plugin service template parameter possible value dto
// swagger:model PluginServiceTemplateParameterPossibleValueDto
type PluginServiceTemplateParameterPossibleValueDto struct {

	// parameters
	Parameters []*PluginServiceTemplateParameterDto `json:"parameters"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this plugin service template parameter possible value dto
func (m *PluginServiceTemplateParameterPossibleValueDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginServiceTemplateParameterPossibleValueDto) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginServiceTemplateParameterPossibleValueDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginServiceTemplateParameterPossibleValueDto) UnmarshalBinary(b []byte) error {
	var res PluginServiceTemplateParameterPossibleValueDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}