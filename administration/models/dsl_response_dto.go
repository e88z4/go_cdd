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

// DslResponseDto dsl response dto
// swagger:model DslResponseDto
type DslResponseDto struct {

	// entities
	Entities []*DslIdentifiableDto `json:"entities"`

	// warnings
	Warnings []*DslWarningDto `json:"warnings"`
}

// Validate validates this dsl response dto
func (m *DslResponseDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DslResponseDto) validateEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DslResponseDto) validateWarnings(formats strfmt.Registry) error {

	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	for i := 0; i < len(m.Warnings); i++ {
		if swag.IsZero(m.Warnings[i]) { // not required
			continue
		}

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DslResponseDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DslResponseDto) UnmarshalBinary(b []byte) error {
	var res DslResponseDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
