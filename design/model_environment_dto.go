/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EnvironmentDto struct {
	Applications []ShallowApplicationDto `json:"applications,omitempty"`
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsProduction bool `json:"isProduction,omitempty"`
	Name string `json:"name,omitempty"`
	Project *NamedIdentifiableDto `json:"project,omitempty"`
}