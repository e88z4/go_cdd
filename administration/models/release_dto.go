// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReleaseDto release dto
// swagger:model ReleaseDto
type ReleaseDto struct {

	// application versions
	ApplicationVersions []*ApplicationVersionDto `json:"applicationVersions"`

	// applications
	Applications []*ApplicationDto `json:"applications"`

	// conflicts
	Conflicts []ConflictDto `json:"conflicts"`

	// creation date
	CreationDate int64 `json:"creationDate,omitempty"`

	// creator first name
	CreatorFirstName string `json:"creatorFirstName,omitempty"`

	// creator last name
	CreatorLastName string `json:"creatorLastName,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end date
	EndDate int64 `json:"endDate,omitempty"`

	// execution data
	ExecutionData *ReleaseExecutionDto `json:"executionData,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// member parties
	MemberParties []*PartyDto `json:"memberParties"`

	// name
	Name string `json:"name,omitempty"`

	// owner parties
	OwnerParties []*PartyDto `json:"ownerParties"`

	// phases
	Phases []*PhaseDto `json:"phases"`

	// planned end date
	PlannedEndDate int64 `json:"plannedEndDate,omitempty"`

	// planned start date
	PlannedStartDate int64 `json:"plannedStartDate,omitempty"`

	// production phase
	ProductionPhase *NamedIdentifiableDto `json:"productionPhase,omitempty"`

	// project
	Project *NamedIdentifiableDto `json:"project,omitempty"`

	// score
	Score *ScoreDto `json:"score,omitempty"`

	// start date
	StartDate int64 `json:"startDate,omitempty"`

	// track
	Track *ApprovedDateNamedIdentifiableDto `json:"track,omitempty"`

	// version of the release
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this release dto
func (m *ReleaseDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberParties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerParties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductionPhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseDto) validateApplicationVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplicationVersions); i++ {
		if swag.IsZero(m.ApplicationVersions[i]) { // not required
			continue
		}

		if m.ApplicationVersions[i] != nil {
			if err := m.ApplicationVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applicationVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseDto) validateApplications(formats strfmt.Registry) error {

	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {
		if swag.IsZero(m.Applications[i]) { // not required
			continue
		}

		if m.Applications[i] != nil {
			if err := m.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseDto) validateExecutionData(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionData) { // not required
		return nil
	}

	if m.ExecutionData != nil {
		if err := m.ExecutionData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("executionData")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseDto) validateMemberParties(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberParties) { // not required
		return nil
	}

	for i := 0; i < len(m.MemberParties); i++ {
		if swag.IsZero(m.MemberParties[i]) { // not required
			continue
		}

		if m.MemberParties[i] != nil {
			if err := m.MemberParties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberParties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseDto) validateOwnerParties(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerParties) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnerParties); i++ {
		if swag.IsZero(m.OwnerParties[i]) { // not required
			continue
		}

		if m.OwnerParties[i] != nil {
			if err := m.OwnerParties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ownerParties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseDto) validatePhases(formats strfmt.Registry) error {

	if swag.IsZero(m.Phases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phases); i++ {
		if swag.IsZero(m.Phases[i]) { // not required
			continue
		}

		if m.Phases[i] != nil {
			if err := m.Phases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReleaseDto) validateProductionPhase(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductionPhase) { // not required
		return nil
	}

	if m.ProductionPhase != nil {
		if err := m.ProductionPhase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productionPhase")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseDto) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseDto) validateScore(formats strfmt.Registry) error {

	if swag.IsZero(m.Score) { // not required
		return nil
	}

	if m.Score != nil {
		if err := m.Score.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("score")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseDto) validateTrack(formats strfmt.Registry) error {

	if swag.IsZero(m.Track) { // not required
		return nil
	}

	if m.Track != nil {
		if err := m.Track.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("track")
			}
			return err
		}
	}

	return nil
}

func (m *ReleaseDto) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseDto) UnmarshalBinary(b []byte) error {
	var res ReleaseDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}