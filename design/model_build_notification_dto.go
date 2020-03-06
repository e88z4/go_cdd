/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BuildNotificationDto struct {
	ApplicationName string `json:"applicationName,omitempty"`
	ApplicationVersionBuildNumber string `json:"applicationVersionBuildNumber,omitempty"`
	ApplicationVersionName string `json:"applicationVersionName,omitempty"`
	BuildStartTime int64 `json:"buildStartTime,omitempty"`
	Commits []CommitIdDto `json:"commits,omitempty"`
	ReleaseTokens []ArgumentDto `json:"releaseTokens,omitempty"`
}
