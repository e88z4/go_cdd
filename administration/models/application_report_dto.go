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

// ApplicationReportDto application report dto
// swagger:model ApplicationReportDto
type ApplicationReportDto struct {

	// application
	Application *ApplicationDto `json:"application,omitempty"`

	// environment versions
	EnvironmentVersions []*EnvironmentVersionDto `json:"environmentVersions"`
}

// Validate validates this application report dto
func (m *ApplicationReportDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationReportDto) validateApplication(formats strfmt.Registry) error {

	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationReportDto) validateEnvironmentVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvironmentVersions); i++ {
		if swag.IsZero(m.EnvironmentVersions[i]) { // not required
			continue
		}

		if m.EnvironmentVersions[i] != nil {
			if err := m.EnvironmentVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environmentVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationReportDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationReportDto) UnmarshalBinary(b []byte) error {
	var res ApplicationReportDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}