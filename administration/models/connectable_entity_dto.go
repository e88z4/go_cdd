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

// ConnectableEntityDto connectable entity dto
// swagger:model ConnectableEntityDto
type ConnectableEntityDto struct {

	// connectivity arguments
	ConnectivityArguments []*ArgumentDto `json:"connectivityArguments"`

	// last connectivity status
	LastConnectivityStatus string `json:"lastConnectivityStatus,omitempty"`

	// last connectivity test date
	LastConnectivityTestDate int64 `json:"lastConnectivityTestDate,omitempty"`

	// last connectivity test message
	LastConnectivityTestMessage string `json:"lastConnectivityTestMessage,omitempty"`
}

// Validate validates this connectable entity dto
func (m *ConnectableEntityDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectivityArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectableEntityDto) validateConnectivityArguments(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectivityArguments) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectivityArguments); i++ {
		if swag.IsZero(m.ConnectivityArguments[i]) { // not required
			continue
		}

		if m.ConnectivityArguments[i] != nil {
			if err := m.ConnectivityArguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectivityArguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectableEntityDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectableEntityDto) UnmarshalBinary(b []byte) error {
	var res ConnectableEntityDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
