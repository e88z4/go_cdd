/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BasicPluginServiceParameterDtoobject struct {
	Id int64 `json:"id,omitempty"`
	TemplateParameterId int64 `json:"templateParameterId,omitempty"`
	Value *interface{} `json:"value,omitempty"`
	ValueId string `json:"valueId,omitempty"`
}
