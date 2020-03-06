/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SystemPreferenceDto struct {
	Category string `json:"category,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Id int64 `json:"id,omitempty"`
	MaximumValue int64 `json:"maximumValue,omitempty"`
	MinimumValue int64 `json:"minimumValue,omitempty"`
	Name string `json:"name,omitempty"`
	Optional bool `json:"optional,omitempty"`
	PossibleValues []string `json:"possibleValues,omitempty"`
	Type_ string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
