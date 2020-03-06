/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EndpointDto struct {
	Applications []ApplicationDto `json:"applications,omitempty"`
	ConnectivityArguments []ArgumentDto `json:"connectivityArguments,omitempty"`
	Description string `json:"description,omitempty"`
	Environments []EnvironmentDto `json:"environments,omitempty"`
	Id int64 `json:"id,omitempty"`
	InUse bool `json:"inUse,omitempty"`
	LastConnectivityStatus string `json:"lastConnectivityStatus,omitempty"`
	LastConnectivityTestDate int64 `json:"lastConnectivityTestDate,omitempty"`
	LastConnectivityTestMessage string `json:"lastConnectivityTestMessage,omitempty"`
	Name string `json:"name,omitempty"`
	PluginService *PluginServiceDto `json:"pluginService,omitempty"`
	Type_ string `json:"type,omitempty"`
}
