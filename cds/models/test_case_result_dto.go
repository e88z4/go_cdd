// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestCaseResultDto test case result dto
// swagger:model TestCaseResultDto
type TestCaseResultDto struct {

	// change Id
	ChangeID string `json:"changeId,omitempty"`

	// execution duration
	ExecutionDuration int64 `json:"executionDuration,omitempty"`

	// execution end date
	ExecutionEndDate int64 `json:"executionEndDate,omitempty"`

	// execution start date
	ExecutionStartDate int64 `json:"executionStartDate,omitempty"`

	// execution status
	// Enum: [PASSED FAILED SKIPPED ERROR DISABLED]
	ExecutionStatus string `json:"executionStatus,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// result message
	ResultMessage string `json:"resultMessage,omitempty"`

	// result URL
	ResultURL string `json:"resultURL,omitempty"`

	// test suite result Id
	TestSuiteResultID string `json:"testSuiteResultId,omitempty"`
}

// Validate validates this test case result dto
func (m *TestCaseResultDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var testCaseResultDtoTypeExecutionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASSED","FAILED","SKIPPED","ERROR","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		testCaseResultDtoTypeExecutionStatusPropEnum = append(testCaseResultDtoTypeExecutionStatusPropEnum, v)
	}
}

const (

	// TestCaseResultDtoExecutionStatusPASSED captures enum value "PASSED"
	TestCaseResultDtoExecutionStatusPASSED string = "PASSED"

	// TestCaseResultDtoExecutionStatusFAILED captures enum value "FAILED"
	TestCaseResultDtoExecutionStatusFAILED string = "FAILED"

	// TestCaseResultDtoExecutionStatusSKIPPED captures enum value "SKIPPED"
	TestCaseResultDtoExecutionStatusSKIPPED string = "SKIPPED"

	// TestCaseResultDtoExecutionStatusERROR captures enum value "ERROR"
	TestCaseResultDtoExecutionStatusERROR string = "ERROR"

	// TestCaseResultDtoExecutionStatusDISABLED captures enum value "DISABLED"
	TestCaseResultDtoExecutionStatusDISABLED string = "DISABLED"
)

// prop value enum
func (m *TestCaseResultDto) validateExecutionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, testCaseResultDtoTypeExecutionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TestCaseResultDto) validateExecutionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutionStatusEnum("executionStatus", "body", m.ExecutionStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestCaseResultDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestCaseResultDto) UnmarshalBinary(b []byte) error {
	var res TestCaseResultDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
