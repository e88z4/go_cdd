/*
 * DESIGN
 *
 * Retrieves, updates, and deletes design data
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ReleaseTokenDto struct {
	Description string `json:"description,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsProduction bool `json:"isProduction,omitempty"`
	IsSystem bool `json:"isSystem,omitempty"`
	Name string `json:"name,omitempty"`
	Release *IdentifiableDto `json:"release,omitempty"`
	ReleaseTokenTemplate *IdentifiableDto `json:"releaseTokenTemplate,omitempty"`
	Scope string `json:"scope,omitempty"`
	Value string `json:"value,omitempty"`
}
