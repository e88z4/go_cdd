// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrackPermissionsDto track permissions dto
// swagger:model TrackPermissionsDto
type TrackPermissionsDto struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// is add release allowed
	IsAddReleaseAllowed bool `json:"isAddReleaseAllowed,omitempty"`

	// is approve allowed
	IsApproveAllowed bool `json:"isApproveAllowed,omitempty"`

	// is create allowed
	IsCreateAllowed bool `json:"isCreateAllowed,omitempty"`

	// is delete allowed
	IsDeleteAllowed bool `json:"isDeleteAllowed,omitempty"`

	// is edit allowed
	IsEditAllowed bool `json:"isEditAllowed,omitempty"`

	// is remove release allowed
	IsRemoveReleaseAllowed bool `json:"isRemoveReleaseAllowed,omitempty"`
}

// Validate validates this track permissions dto
func (m *TrackPermissionsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrackPermissionsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrackPermissionsDto) UnmarshalBinary(b []byte) error {
	var res TrackPermissionsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
