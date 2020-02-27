// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CRUDPermissionsDto c r u d permissions dto
// swagger:model CRUDPermissionsDto
type CRUDPermissionsDto struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// is create allowed
	IsCreateAllowed bool `json:"isCreateAllowed,omitempty"`

	// is delete allowed
	IsDeleteAllowed bool `json:"isDeleteAllowed,omitempty"`

	// is edit allowed
	IsEditAllowed bool `json:"isEditAllowed,omitempty"`
}

// Validate validates this c r u d permissions dto
func (m *CRUDPermissionsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CRUDPermissionsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CRUDPermissionsDto) UnmarshalBinary(b []byte) error {
	var res CRUDPermissionsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
