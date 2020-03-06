/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApplicationVersionDto struct {
	Application *NamedIdentifiableDto `json:"application,omitempty"`
	BasedOn *NamedIdentifiableDto `json:"basedOn,omitempty"`
	BuildNumber string `json:"buildNumber,omitempty"`
	CommitId string `json:"commitId,omitempty"`
	CommitSource *CommitSourceDto `json:"commitSource,omitempty"`
	ContentSource *ContentSourceDto `json:"contentSource,omitempty"`
	CreationDate int64 `json:"creationDate,omitempty"`
	DependentBy []ApplicationVersionDto `json:"dependentBy,omitempty"`
	DependsOn []ApplicationVersionDto `json:"dependsOn,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
