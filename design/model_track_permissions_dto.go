/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TrackPermissionsDto struct {
	Enabled bool `json:"enabled,omitempty"`
	IsAddReleaseAllowed bool `json:"isAddReleaseAllowed,omitempty"`
	IsApproveAllowed bool `json:"isApproveAllowed,omitempty"`
	IsCreateAllowed bool `json:"isCreateAllowed,omitempty"`
	IsDeleteAllowed bool `json:"isDeleteAllowed,omitempty"`
	IsEditAllowed bool `json:"isEditAllowed,omitempty"`
	IsRemoveReleaseAllowed bool `json:"isRemoveReleaseAllowed,omitempty"`
}
