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

// PagedResultDtoTestSuiteResultDto paged result dto test suite result dto
// swagger:model PagedResultDto«TestSuiteResultDto»
type PagedResultDtoTestSuiteResultDto struct {

	// data
	Data []*TestSuiteResultDto `json:"data"`

	// page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// total results count
	TotalResultsCount int64 `json:"totalResultsCount,omitempty"`
}

// Validate validates this paged result dto test suite result dto
func (m *PagedResultDtoTestSuiteResultDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagedResultDtoTestSuiteResultDto) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PagedResultDtoTestSuiteResultDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagedResultDtoTestSuiteResultDto) UnmarshalBinary(b []byte) error {
	var res PagedResultDtoTestSuiteResultDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
