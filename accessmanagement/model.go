package accessmanagement

// RoleBinding is an organization role binding. Bindings maps a principal id
// ("<userPoolId>/<idp>/<email>") to {"roles": [...]}; key bindings are masked by the API.
type RoleBinding struct {
	OrgId                     *string                `json:"OrgId,omitempty" url:"-"`
	ParentId                  *string                `json:"ParentId,omitempty" url:"-"`
	SubResourceId             *string                `json:"SubResourceId,omitempty" url:"-"`
	ResourceId                *string                `json:"ResourceId,omitempty" url:"-"`
	ResourceName              *string                `json:"ResourceName,omitempty" url:"-"`
	ResourceType              *string                `json:"ResourceType,omitempty" url:"-"`
	Description               *string                `json:"Description,omitempty" url:"-"`
	DocVersion                *string                `json:"DocVersion,omitempty" url:"-"`
	IsActive                  *string                `json:"IsActive,omitempty" url:"-"`
	IsArchive                 *string                `json:"IsArchive,omitempty" url:"-"`
	Tags                      []string               `json:"Tags,omitempty" url:"-"`
	Authors                   []string               `json:"Authors,omitempty" url:"-"`
	CreatedAt                 *int64                 `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt                *int64                 `json:"ModifiedAt,omitempty" url:"-"`
	Bindings                  map[string]interface{} `json:"Bindings,omitempty" url:"-"`
	RunnerGroupAPIKeyBindings map[string]interface{} `json:"RunnerGroupAPIKeyBindings,omitempty" url:"-"`
	OrgRoleBindingConfig      map[string]interface{} `json:"OrgRoleBindingConfig,omitempty" url:"-"`
	AliasToUserIdMap          map[string]interface{} `json:"AliasToUserIdMap,omitempty" url:"-"`
}

type RoleBindingResponse struct {
	Msg *RoleBinding `json:"msg,omitempty" url:"-"`
}

// ApiTokenRequest creates or retrieves the caller's legacy API key. Set Regenerate to replace an
// existing key, or RunnerGroupId ("/runnergroups/<name>") to look up a runner group's key.
type ApiTokenRequest struct {
	Regenerate    *bool   `json:"regenerate,omitempty" url:"-"`
	RunnerGroupId *string `json:"runnerGroupId,omitempty" url:"-"`
}

// ApiTokenResponse carries the API key itself in Data.
type ApiTokenResponse struct {
	Msg  *string `json:"msg,omitempty" url:"-"`
	Data *string `json:"data,omitempty" url:"-"`
}

type ApiTokenDeleteResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}
