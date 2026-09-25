package stacks

// CompareStackRequest is the body for a dry-run comparison of a stack against a stack template.
type CompareStackRequest struct {
	TargetTemplateGroupId string                 `json:"targetTemplateGroupId" url:"-"`
	UpgradeMode           *string                `json:"upgradeMode,omitempty" url:"-"`
	PatchData             map[string]interface{} `json:"patchData,omitempty" url:"-"`
}

type CompareStackResponse struct {
	Msg  *string                `json:"msg,omitempty" url:"-"`
	Data map[string]interface{} `json:"data,omitempty" url:"-"`
}
