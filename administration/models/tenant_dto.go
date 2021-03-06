// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantDto tenant dto
// swagger:model TenantDto
type TenantDto struct {

	// country
	Country string `json:"country,omitempty"`

	// creation date
	CreationDate int64 `json:"creationDate,omitempty"`

	// creator username
	CreatorUsername string `json:"creatorUsername,omitempty"`

	// enterprise site Id
	EnterpriseSiteID int64 `json:"enterpriseSiteId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is enabled
	IsEnabled bool `json:"isEnabled,omitempty"`

	// last login date
	LastLoginDate int64 `json:"lastLoginDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization
	Organization string `json:"organization,omitempty"`

	// payment plan creator email
	PaymentPlanCreatorEmail string `json:"paymentPlanCreatorEmail,omitempty"`

	// payment plan creator first name
	PaymentPlanCreatorFirstName string `json:"paymentPlanCreatorFirstName,omitempty"`

	// payment plan creator last name
	PaymentPlanCreatorLastName string `json:"paymentPlanCreatorLastName,omitempty"`

	// payment plan Id
	PaymentPlanID string `json:"paymentPlanId,omitempty"`

	// schema name
	SchemaName string `json:"schemaName,omitempty"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this tenant dto
func (m *TenantDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TenantDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantDto) UnmarshalBinary(b []byte) error {
	var res TenantDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
