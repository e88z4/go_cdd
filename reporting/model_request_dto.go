/*
 * REPORTING
 *
 * Retrieves, updates, and deletes report and widget information
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RequestDto struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Queries []QueryDto `json:"queries,omitempty"`
}
