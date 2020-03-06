/*
 * REPORTING
 *
 * Retrieves, updates, and deletes report and widget information
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ExecutedParameterReportDto struct {
	Direction string `json:"direction,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ParameterName string `json:"parameterName,omitempty"`
	ParameterOrder int32 `json:"parameterOrder,omitempty"`
	Value string `json:"value,omitempty"`
}