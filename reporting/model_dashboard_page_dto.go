/*
 * REPORTING
 *
 * Retrieves, updates, and deletes report and widget information
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DashboardPageDto struct {
	Content string `json:"content,omitempty"`
	Id int64 `json:"id,omitempty"`
	LayoutType string `json:"layoutType,omitempty"`
	Name string `json:"name,omitempty"`
}
