/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VersionDto struct {
	Build string `json:"build,omitempty"`
	ConnectivityArguments []ArgumentDto `json:"connectivityArguments,omitempty"`
	GitVersion string `json:"gitVersion,omitempty"`
	HomeFolder string `json:"homeFolder,omitempty"`
	LastConnectivityStatus string `json:"lastConnectivityStatus,omitempty"`
	LastConnectivityTestDate int64 `json:"lastConnectivityTestDate,omitempty"`
	LastConnectivityTestMessage string `json:"lastConnectivityTestMessage,omitempty"`
	LastConnectivityTestTime string `json:"lastConnectivityTestTime,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Services []VersionDto `json:"services,omitempty"`
	Version string `json:"version,omitempty"`
	VersionDate int64 `json:"versionDate,omitempty"`
	VersionTime string `json:"versionTime,omitempty"`
}
