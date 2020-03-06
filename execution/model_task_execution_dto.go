/*
 * EXECUTION
 *
 * Retrieves and changes execution data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TaskExecutionDto struct {
	AllowedStatuses []string `json:"allowedStatuses,omitempty"`
	DenyRerunReason string `json:"denyRerunReason,omitempty"`
	DeployedApplicationVersions []ApplicationVersionDto `json:"deployedApplicationVersions,omitempty"`
	DetailedStatusDescription string `json:"detailedStatusDescription,omitempty"`
	DetailedStatusDescriptionJson *interface{} `json:"detailedStatusDescriptionJson,omitempty"`
	DetailedStatusSeverity string `json:"detailedStatusSeverity,omitempty"`
	EndDate int64 `json:"endDate,omitempty"`
	Id int64 `json:"id,omitempty"`
	PercentCompleted float32 `json:"percentCompleted,omitempty"`
	PhaseId int64 `json:"phaseId,omitempty"`
	StartDate int64 `json:"startDate,omitempty"`
	Status string `json:"status,omitempty"`
	StatusDescription string `json:"statusDescription,omitempty"`
	TaskExecutionDisabled bool `json:"taskExecutionDisabled,omitempty"`
	TaskId int64 `json:"taskId,omitempty"`
}
