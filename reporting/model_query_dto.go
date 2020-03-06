/*
 * REPORTING
 *
 * Retrieves, updates, and deletes report and widget information
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QueryDto struct {
	Category string `json:"category,omitempty"`
	Filters []FilterDto `json:"filters,omitempty"`
	Name string `json:"name,omitempty"`
	QueryArguments []ArgumentDto `json:"queryArguments,omitempty"`
	Version string `json:"version,omitempty"`
}