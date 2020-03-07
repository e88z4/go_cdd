// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DashboardPermissionsDto dashboard permissions dto
// swagger:model DashboardPermissionsDto
type DashboardPermissionsDto struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// is upload widget allowed
	IsUploadWidgetAllowed bool `json:"isUploadWidgetAllowed,omitempty"`
}

// Validate validates this dashboard permissions dto
func (m *DashboardPermissionsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DashboardPermissionsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardPermissionsDto) UnmarshalBinary(b []byte) error {
	var res DashboardPermissionsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
