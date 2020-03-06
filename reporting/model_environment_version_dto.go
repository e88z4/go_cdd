/*
 * REPORTING
 *
 * Retrieves, updates, and deletes report and widget information
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EnvironmentVersionDto struct {
	Application *NamedIdentifiableDto `json:"application,omitempty"`
	ApplicationVersion *NamedIdentifiableDto `json:"applicationVersion,omitempty"`
	BuildNumber string `json:"buildNumber,omitempty"`
	CommitId string `json:"commitId,omitempty"`
	Environment *NamedIdentifiableDto `json:"environment,omitempty"`
	ExternalStatus string `json:"externalStatus,omitempty"`
	Phase *NamedIdentifiableDto `json:"phase,omitempty"`
	PromotionDate int64 `json:"promotionDate,omitempty"`
	Release *NamedIdentifiableDto `json:"release,omitempty"`
	Status string `json:"status,omitempty"`
	Task *NamedIdentifiableDto `json:"task,omitempty"`
}
