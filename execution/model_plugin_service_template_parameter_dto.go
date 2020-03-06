/*
 * EXECUTION
 *
 * Retrieves and changes execution data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PluginServiceTemplateParameterDto struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	Description string `json:"description,omitempty"`
	Direction string `json:"direction,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	DisplayType string `json:"displayType,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsOptional bool `json:"isOptional,omitempty"`
	MultipleSelection bool `json:"multipleSelection,omitempty"`
	Name string `json:"name,omitempty"`
	ParameterDependencies []string `json:"parameterDependencies,omitempty"`
	ParameterMethod string `json:"parameterMethod,omitempty"`
	ParameterOrder int32 `json:"parameterOrder,omitempty"`
	PossibleValues []PluginServiceTemplateParameterPossibleValueDto `json:"possibleValues,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
}