/*
 * REPORTING
 *
 * Retrieves, updates, and deletes report and widget information
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PhasePluginTestExecutionDto struct {
	Application *NamedIdentifiableDto `json:"application,omitempty"`
	ApplicationVersion *NamedIdentifiableDto `json:"applicationVersion,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Phase *NamedIdentifiableDto `json:"phase,omitempty"`
	PluginTestExecutions []PluginTestExecutionDto `json:"pluginTestExecutions,omitempty"`
}