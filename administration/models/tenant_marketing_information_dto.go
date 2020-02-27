// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantMarketingInformationDto tenant marketing information dto
// swagger:model TenantMarketingInformationDto
type TenantMarketingInformationDto struct {

	// authorized to update tenant information
	AuthorizedToUpdateTenantInformation bool `json:"authorizedToUpdateTenantInformation,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this tenant marketing information dto
func (m *TenantMarketingInformationDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TenantMarketingInformationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantMarketingInformationDto) UnmarshalBinary(b []byte) error {
	var res TenantMarketingInformationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}