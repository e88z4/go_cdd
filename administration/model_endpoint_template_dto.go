/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EndpointTemplateDto struct {
	ConnectivityTest bool `json:"connectivityTest,omitempty"`
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsPluginEndpointExist bool `json:"isPluginEndpointExist,omitempty"`
	IsTaggable bool `json:"isTaggable,omitempty"`
	Name string `json:"name,omitempty"`
	PluginEndpointsOptional bool `json:"pluginEndpointsOptional,omitempty"`
	PluginId int64 `json:"pluginId,omitempty"`
	PluginName string `json:"pluginName,omitempty"`
	ServiceType string `json:"serviceType,omitempty"`
	TemplateParameters []PluginServiceTemplateParameterDto `json:"templateParameters,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
}