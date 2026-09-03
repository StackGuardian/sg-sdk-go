package stackworkflowruns

// ListAllStackWorkflowRunsRequest filters the runs of a stack workflow. Comma-separated
// values are accepted where the API documents them.
type ListAllStackWorkflowRunsRequest struct {
	Limit            *int    `json:"-" url:"limit,omitempty"`
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	// Unix timestamp in milliseconds
	CreatedAfter           *string `json:"-" url:"createdAfter,omitempty"`
	ParentKSUID            *string `json:"-" url:"parentKSUID,omitempty"`
	ResourceNames          *string `json:"-" url:"resourceNames,omitempty"`
	TemplateId             *string `json:"-" url:"templateId,omitempty"`
	UseMarketplaceTemplate *string `json:"-" url:"useMarketplaceTemplate,omitempty"`
	Author                 *string `json:"-" url:"author,omitempty"`
	Approver               *string `json:"-" url:"approver,omitempty"`
	LatestStatus           *string `json:"-" url:"latestStatus,omitempty"`
	RepoId                 *string `json:"-" url:"repoId,omitempty"`
	Ref                    *string `json:"-" url:"ref,omitempty"`
	TerraformAction        *string `json:"-" url:"terraformAction,omitempty"`
	RunnerType             *string `json:"-" url:"runnerType,omitempty"`
	// Comma-separated key or key:value pairs
	ContextTags *string `json:"-" url:"contextTags,omitempty"`
}

type StackWorkflowRunDeleteResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}
