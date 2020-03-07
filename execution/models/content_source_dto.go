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

// ContentSourceDto content source dto
// swagger:model ContentSourceDto
type ContentSourceDto struct {

	// content items
	ContentItems []*ContentItemDto `json:"contentItems"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// plugin service
	PluginService *PluginServiceDto `json:"pluginService,omitempty"`
}

// Validate validates this content source dto
func (m *ContentSourceDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentSourceDto) validateContentItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ContentItems); i++ {
		if swag.IsZero(m.ContentItems[i]) { // not required
			continue
		}

		if m.ContentItems[i] != nil {
			if err := m.ContentItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contentItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentSourceDto) validatePluginService(formats strfmt.Registry) error {

	if swag.IsZero(m.PluginService) { // not required
		return nil
	}

	if m.PluginService != nil {
		if err := m.PluginService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pluginService")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentSourceDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentSourceDto) UnmarshalBinary(b []byte) error {
	var res ContentSourceDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}