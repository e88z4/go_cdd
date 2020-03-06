/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TrackDto struct {
	CreationDate int64 `json:"creationDate,omitempty"`
	DeliveredTogether bool `json:"deliveredTogether,omitempty"`
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	Milestones []MilestoneDto `json:"milestones,omitempty"`
	Name string `json:"name,omitempty"`
	Owners []PartyDto `json:"owners,omitempty"`
	Permissions *TrackPermissionsDto `json:"permissions,omitempty"`
	ProductionEndDate int64 `json:"productionEndDate,omitempty"`
	ProductionStartDate int64 `json:"productionStartDate,omitempty"`
	Releases []NamedIdentifiableDto `json:"releases,omitempty"`
	Status string `json:"status,omitempty"`
}