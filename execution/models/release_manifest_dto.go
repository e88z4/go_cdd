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

// ReleaseManifestDto release manifest dto
// swagger:model ReleaseManifestDto
type ReleaseManifestDto struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// release
	Release *IdentifiableDto `json:"release,omitempty"`

	// token values
	TokenValues []*ReleaseTokenValueDto `json:"tokenValues"`
}

// Validate validates this release manifest dto
func (m *ReleaseManifestDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseManifestDto) validateRelease(formats strfmt.Registry) error {

	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseManifestDto) validateTokenValues(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenValues) { // not required
		return nil
	}

	for i := 0; i < len(m.TokenValues); i++ {
		if swag.IsZero(m.TokenValues[i]) { // not required
			continue
		}

		if m.TokenValues[i] != nil {
			if err := m.TokenValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokenValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseManifestDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseManifestDto) UnmarshalBinary(b []byte) error {
	var res ReleaseManifestDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
