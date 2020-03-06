/*
 * ADMINISTRATION
 *
 * Retrieves, syncs, and deletes administration data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RoleDto struct {
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsProjectSpecific bool `json:"isProjectSpecific,omitempty"`
	ModificationDate int64 `json:"modificationDate,omitempty"`
	Modifier *UserDto `json:"modifier,omitempty"`
	Name string `json:"name,omitempty"`
	NumberOfRoleParties int64 `json:"numberOfRoleParties,omitempty"`
	Permissions []PermissionDto `json:"permissions,omitempty"`
	System bool `json:"system,omitempty"`
}
