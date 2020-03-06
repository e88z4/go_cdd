/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PluginUrlDto struct {
	DiscoveryUrl string `json:"discoveryUrl,omitempty"`
	IsUseProxy bool `json:"isUseProxy,omitempty"`
}
