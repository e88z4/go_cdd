// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RedirectToLandingPageDto redirect to landing page dto
// swagger:model RedirectToLandingPageDto
type RedirectToLandingPageDto struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this redirect to landing page dto
func (m *RedirectToLandingPageDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RedirectToLandingPageDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedirectToLandingPageDto) UnmarshalBinary(b []byte) error {
	var res RedirectToLandingPageDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
