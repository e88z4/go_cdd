/*
 * CDS
 *
 * Retrieves, updates, and deletes cds data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ChangedFileDto struct {
	Application *NamedIdentifiableDto `json:"application,omitempty"`
	ApplicationVersion *NamedIdentifiableDto `json:"applicationVersion,omitempty"`
	FileId string `json:"fileId,omitempty"`
	Filename string `json:"filename,omitempty"`
	Id string `json:"id,omitempty"`
	LastAccessDate int64 `json:"lastAccessDate,omitempty"`
	LastChangeId string `json:"lastChangeId,omitempty"`
	LastCommitDate int64 `json:"lastCommitDate,omitempty"`
	Status string `json:"status,omitempty"`
}
