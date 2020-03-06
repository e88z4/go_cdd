/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TestSourceDto struct {
	ApplicationVersion *ApplicationVersionDto `json:"applicationVersion,omitempty"`
	CreationDate int64 `json:"creationDate,omitempty"`
	Description string `json:"description,omitempty"`
	DetailedStatusDescription string `json:"detailedStatusDescription,omitempty"`
	Id int64 `json:"id,omitempty"`
	LastCommitId string `json:"lastCommitId,omitempty"`
	LastExecutionDate int64 `json:"lastExecutionDate,omitempty"`
	LastImportDate int64 `json:"lastImportDate,omitempty"`
	LimitNumberOfTestSuite int32 `json:"limitNumberOfTestSuite,omitempty"`
	Name string `json:"name,omitempty"`
	NumberOfTestSuite int32 `json:"numberOfTestSuite,omitempty"`
	PluginService *PluginServiceDto `json:"pluginService,omitempty"`
	Status string `json:"status,omitempty"`
	StatusDescription string `json:"statusDescription,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TestSuites []StringIdentifiableDto `json:"testSuites,omitempty"`
	TotalNumberOfTestSuite int32 `json:"totalNumberOfTestSuite,omitempty"`
}