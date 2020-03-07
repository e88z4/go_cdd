// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestSuiteResultDto test suite result dto
// swagger:model TestSuiteResultDto
type TestSuiteResultDto struct {

	// application
	Application *NamedIdentifiableDto `json:"application,omitempty"`

	// application version
	ApplicationVersion *NamedIdentifiableDto `json:"applicationVersion,omitempty"`

	// build number
	BuildNumber string `json:"buildNumber,omitempty"`

	// change Id
	ChangeID string `json:"changeId,omitempty"`

	// commit Id
	CommitID string `json:"commitId,omitempty"`

	// endpoint
	Endpoint *NamedIdentifiableDto `json:"endpoint,omitempty"`

	// environments
	Environments []*NamedIdentifiableDto `json:"environments"`

	// execution detailed status description
	ExecutionDetailedStatusDescription string `json:"executionDetailedStatusDescription,omitempty"`

	// execution duration
	ExecutionDuration int64 `json:"executionDuration,omitempty"`

	// execution end date
	ExecutionEndDate int64 `json:"executionEndDate,omitempty"`

	// execution start date
	ExecutionStartDate int64 `json:"executionStartDate,omitempty"`

	// execution status description
	ExecutionStatusDescription string `json:"executionStatusDescription,omitempty"`

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number of failed test cases
	NumberOfFailedTestCases int32 `json:"numberOfFailedTestCases,omitempty"`

	// number of successful test cases
	NumberOfSuccessfulTestCases int32 `json:"numberOfSuccessfulTestCases,omitempty"`

	// phase
	Phase *NamedIdentifiableDto `json:"phase,omitempty"`

	// phase execution Id
	PhaseExecutionID string `json:"phaseExecutionId,omitempty"`

	// phase execution start date
	PhaseExecutionStartDate int64 `json:"phaseExecutionStartDate,omitempty"`

	// plugin
	Plugin *NamedIdentifiableDto `json:"plugin,omitempty"`

	// release
	Release *NamedVersionedIdentifiableDto `json:"release,omitempty"`

	// result URL
	ResultURL string `json:"resultURL,omitempty"`

	// strategies
	Strategies []string `json:"strategies"`

	// tags
	Tags []string `json:"tags"`

	// task
	Task *NamedIdentifiableDto `json:"task,omitempty"`

	// test case results
	TestCaseResults []*TestCaseResultDto `json:"testCaseResults"`

	// test suite
	TestSuite *StringIdentifiableDto `json:"testSuite,omitempty"`

	// test suite execution end date
	TestSuiteExecutionEndDate int64 `json:"testSuiteExecutionEndDate,omitempty"`

	// test suite execution order
	TestSuiteExecutionOrder int64 `json:"testSuiteExecutionOrder,omitempty"`

	// test suite execution status
	// Enum: [PASSED FAILED SKIPPED ERROR DISABLED]
	TestSuiteExecutionStatus string `json:"testSuiteExecutionStatus,omitempty"`

	// test suites execution Id
	TestSuitesExecutionID string `json:"testSuitesExecutionId,omitempty"`
}

// Validate validates this test suite result dto
func (m *TestSuiteResultDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestCaseResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestSuite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestSuiteExecutionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestSuiteResultDto) validateApplication(formats strfmt.Registry) error {

	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *TestSuiteResultDto) validateApplicationVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationVersion) { // not required
		return nil
	}

	if m.ApplicationVersion != nil {
		if err := m.ApplicationVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationVersion")
			}
			return err
		}
	}

	return nil
}

func (m *TestSuiteResultDto) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *TestSuiteResultDto) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestSuiteResultDto) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *TestSuiteResultDto) validatePlugin(formats strfmt.Registry) error {

	if swag.IsZero(m.Plugin) { // not required
		return nil
	}

	if m.Plugin != nil {
		if err := m.Plugin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugin")
			}
			return err
		}
	}

	return nil
}

func (m *TestSuiteResultDto) validateRelease(formats strfmt.Registry) error {

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

func (m *TestSuiteResultDto) validateTask(formats strfmt.Registry) error {

	if swag.IsZero(m.Task) { // not required
		return nil
	}

	if m.Task != nil {
		if err := m.Task.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task")
			}
			return err
		}
	}

	return nil
}

func (m *TestSuiteResultDto) validateTestCaseResults(formats strfmt.Registry) error {

	if swag.IsZero(m.TestCaseResults) { // not required
		return nil
	}

	for i := 0; i < len(m.TestCaseResults); i++ {
		if swag.IsZero(m.TestCaseResults[i]) { // not required
			continue
		}

		if m.TestCaseResults[i] != nil {
			if err := m.TestCaseResults[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testCaseResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TestSuiteResultDto) validateTestSuite(formats strfmt.Registry) error {

	if swag.IsZero(m.TestSuite) { // not required
		return nil
	}

	if m.TestSuite != nil {
		if err := m.TestSuite.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("testSuite")
			}
			return err
		}
	}

	return nil
}

var testSuiteResultDtoTypeTestSuiteExecutionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PASSED","FAILED","SKIPPED","ERROR","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		testSuiteResultDtoTypeTestSuiteExecutionStatusPropEnum = append(testSuiteResultDtoTypeTestSuiteExecutionStatusPropEnum, v)
	}
}

const (

	// TestSuiteResultDtoTestSuiteExecutionStatusPASSED captures enum value "PASSED"
	TestSuiteResultDtoTestSuiteExecutionStatusPASSED string = "PASSED"

	// TestSuiteResultDtoTestSuiteExecutionStatusFAILED captures enum value "FAILED"
	TestSuiteResultDtoTestSuiteExecutionStatusFAILED string = "FAILED"

	// TestSuiteResultDtoTestSuiteExecutionStatusSKIPPED captures enum value "SKIPPED"
	TestSuiteResultDtoTestSuiteExecutionStatusSKIPPED string = "SKIPPED"

	// TestSuiteResultDtoTestSuiteExecutionStatusERROR captures enum value "ERROR"
	TestSuiteResultDtoTestSuiteExecutionStatusERROR string = "ERROR"

	// TestSuiteResultDtoTestSuiteExecutionStatusDISABLED captures enum value "DISABLED"
	TestSuiteResultDtoTestSuiteExecutionStatusDISABLED string = "DISABLED"
)

// prop value enum
func (m *TestSuiteResultDto) validateTestSuiteExecutionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, testSuiteResultDtoTypeTestSuiteExecutionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TestSuiteResultDto) validateTestSuiteExecutionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TestSuiteExecutionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateTestSuiteExecutionStatusEnum("testSuiteExecutionStatus", "body", m.TestSuiteExecutionStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestSuiteResultDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestSuiteResultDto) UnmarshalBinary(b []byte) error {
	var res TestSuiteResultDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}