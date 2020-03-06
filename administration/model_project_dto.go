/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProjectDto struct {
	ActiveReleasesNumber int64 `json:"activeReleasesNumber,omitempty"`
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
