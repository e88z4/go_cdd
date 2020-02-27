// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserUpdatePolicyDto user update policy dto
// swagger:model UserUpdatePolicyDto
type UserUpdatePolicyDto struct {

	// allowed user email change
	AllowedUserEmailChange bool `json:"allowedUserEmailChange,omitempty"`

	// allowed user password change
	AllowedUserPasswordChange bool `json:"allowedUserPasswordChange,omitempty"`

	// allowed user rename
	AllowedUserRename bool `json:"allowedUserRename,omitempty"`

	// allowed user role change
	AllowedUserRoleChange bool `json:"allowedUserRoleChange,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this user update policy dto
func (m *UserUpdatePolicyDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdatePolicyDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdatePolicyDto) UnmarshalBinary(b []byte) error {
	var res UserUpdatePolicyDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}