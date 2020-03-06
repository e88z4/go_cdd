/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PagedResultDto struct {
	Data []interface{} `json:"data,omitempty"`
	PageNumber int32 `json:"pageNumber,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
	TotalResultsCount int64 `json:"totalResultsCount,omitempty"`
}
