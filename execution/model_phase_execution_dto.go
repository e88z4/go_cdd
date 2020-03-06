/*
 * EXECUTION
 *
 * Retrieves and changes execution data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PhaseExecutionDto struct {
	AllowedStatuses []string `json:"allowedStatuses,omitempty"`
	ApprovedBy *UserDto `json:"approvedBy,omitempty"`
	ApprovedDate int64 `json:"approvedDate,omitempty"`
	DenyApproveReason string `json:"denyApproveReason,omitempty"`
	DenyStartReason string `json:"denyStartReason,omitempty"`
	EndDate int64 `json:"endDate,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsApproved bool `json:"isApproved,omitempty"`
	PercentCompleted float32 `json:"percentCompleted,omitempty"`
	PhaseId int64 `json:"phaseId,omitempty"`
	StartDate int64 `json:"startDate,omitempty"`
	Status string `json:"status,omitempty"`
}