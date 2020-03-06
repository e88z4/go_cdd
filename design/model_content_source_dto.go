/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ContentSourceDto struct {
	ContentItems []ContentItemDto `json:"contentItems,omitempty"`
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	PluginService *PluginServiceDto `json:"pluginService,omitempty"`
}
