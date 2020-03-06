/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MaintenanceWindowDto struct {
	Description string `json:"description,omitempty"`
	EndDate int64 `json:"endDate,omitempty"`
	Environments []EnvironmentDto `json:"environments,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsDisabled bool `json:"isDisabled,omitempty"`
	Name string `json:"name,omitempty"`
	StartDate int64 `json:"startDate,omitempty"`
}