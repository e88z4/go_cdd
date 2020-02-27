// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReleaseTokenValueDto release token value dto
// swagger:model ReleaseTokenValueDto
type ReleaseTokenValueDto struct {

	// id
	ID int64 `json:"id,omitempty"`

	// release manifest
	ReleaseManifest *NamedIdentifiableDto `json:"releaseManifest,omitempty"`

	// release token
	ReleaseToken *NamedIdentifiableDto `json:"releaseToken,omitempty"`

	// token display name
	TokenDisplayName string `json:"tokenDisplayName,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this release token value dto
func (m *ReleaseTokenValueDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReleaseManifest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseTokenValueDto) validateReleaseManifest(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseManifest) { // not required
		return nil
	}

	if m.ReleaseManifest != nil {
		if err := m.ReleaseManifest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseManifest")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseTokenValueDto) validateReleaseToken(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseToken) { // not required
		return nil
	}

	if m.ReleaseToken != nil {
		if err := m.ReleaseToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseToken")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseTokenValueDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseTokenValueDto) UnmarshalBinary(b []byte) error {
	var res ReleaseTokenValueDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}