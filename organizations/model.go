package organizations

// CreateOrganizationRequest creates an organization. Settings follows the organization
// settings schema (infraCostSettings, aiSettings, auditLogSettings, workflowDefaults).
type CreateOrganizationRequest struct {
	// Slug of 4 to 50 characters; generated when omitted
	ResourceName *string                `json:"ResourceName,omitempty" url:"-"`
	Settings     map[string]interface{} `json:"Settings,omitempty" url:"-"`
	// gitops, self-service or codify
	SelectedOnboardingGoal *string `json:"selectedOnboardingGoal,omitempty" url:"-"`
}

type OrganizationCreateResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

// OrganizationListAllResponse lists the names of the organizations the caller belongs to.
// The API answers 204 with no body when there are none.
type OrganizationListAllResponse struct {
	Msg []string `json:"msg,omitempty" url:"-"`
}

// PatchedOrganization updates an organization. EmailRecipients maps an email address to
// {"filters": {"type": ["ALL", "COST", "COMPLIANCE", "SECURITY"]}}.
type PatchedOrganization struct {
	EmailRecipients        map[string]interface{} `json:"EmailRecipients,omitempty" url:"-"`
	Settings               map[string]interface{} `json:"Settings,omitempty" url:"-"`
	SelectedOnboardingGoal *string                `json:"selectedOnboardingGoal,omitempty" url:"-"`
	ExtraSgCreditsLimit    *int                   `json:"ExtraSgCreditsLimit,omitempty" url:"-"`
}

type OrganizationUpdateResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
	// The attributes that were changed
	Data map[string]interface{} `json:"data,omitempty" url:"-"`
}

type OrganizationDeleteResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

// ListAllOrganizationWorkflowsRequest filters the organization-wide workflow listing.
// Comma-separated values are accepted where the API documents them.
type ListAllOrganizationWorkflowsRequest struct {
	Limit            *int    `json:"-" url:"limit,omitempty"`
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	// Set to false to exclude workflows owned by StackGuardian
	FetchSgOwnedWfs                 *bool   `json:"-" url:"fetchSgOwnedWfs,omitempty"`
	LatestWfRunStatuses             *string `json:"-" url:"latestWfRunStatuses,omitempty"`
	LatestDriftRunStatuses          *string `json:"-" url:"latestDriftRunStatuses,omitempty"`
	RepoId                          *string `json:"-" url:"repoId,omitempty"`
	UseMarketplaceTemplate          *string `json:"-" url:"useMarketplaceTemplate,omitempty"`
	IacTemplateId                   *string `json:"-" url:"IACTemplateId,omitempty"`
	ParentIacTemplateId             *string `json:"-" url:"ParentIACTemplateId,omitempty"`
	TerraformConfigTerraformVersion *string `json:"-" url:"terraformConfigTerraformVersion,omitempty"`
	RunnerNames                     *string `json:"-" url:"RunnerNames,omitempty"`
	// Any value restricts the listing to workflows with drift information
	TerraformConfigDriftCheck *string `json:"-" url:"terraformConfigDriftCheck,omitempty"`
	ResourceNames             *string `json:"-" url:"resourceNames,omitempty"`
	Description               *string `json:"-" url:"description,omitempty"`
	Tags                      *string `json:"-" url:"tags,omitempty"`
	// 0 or 1
	IsActive *string `json:"-" url:"IsActive,omitempty"`
	// Internal: set by CountWorkflows
	OnlyCount *string `json:"-" url:"onlyCount,omitempty"`
}

// OrganizationWorkflow is the projection returned by the organization-wide workflow listing.
// Fact fields (drift, cost, policy results) are opaque objects.
type OrganizationWorkflow struct {
	ParentId                    *string     `json:"ParentId,omitempty" url:"-"`
	ResourceId                  *string     `json:"ResourceId,omitempty" url:"-"`
	ResourceName                *string     `json:"ResourceName,omitempty" url:"-"`
	SubResourceId               *string     `json:"SubResourceId,omitempty" url:"-"`
	IsActive                    *string     `json:"IsActive,omitempty" url:"-"`
	Description                 *string     `json:"Description,omitempty" url:"-"`
	Tags                        []string    `json:"Tags,omitempty" url:"-"`
	Authors                     []string    `json:"Authors,omitempty" url:"-"`
	CreatedAt                   *int64      `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt                  *int64      `json:"ModifiedAt,omitempty" url:"-"`
	LatestTerraformAction       *string     `json:"LatestTerraformAction,omitempty" url:"-"`
	LatestWfrunStatus           *string     `json:"LatestWfrunStatus,omitempty" url:"-"`
	LatestWfrunStatusKey        *string     `json:"LatestWfrunStatusKey,omitempty" url:"-"`
	LatestDriftRunStatus        *string     `json:"LatestDriftRunStatus,omitempty" url:"-"`
	WfType                      *string     `json:"WfType,omitempty" url:"-"`
	WorkflowRunFactsProvider    *string     `json:"WorkflowRunFactsProvider,omitempty" url:"-"`
	GitHubComRepoID             *string     `json:"GitHubComRepoID,omitempty" url:"-"`
	TfStateCleaned              interface{} `json:"TfStateCleaned,omitempty" url:"-"`
	PolicyEvalResults           interface{} `json:"PolicyEvalResults,omitempty" url:"-"`
	InfracostBreakdown          interface{} `json:"InfracostBreakdown,omitempty" url:"-"`
	InfracostBreakdownPreApply  interface{} `json:"InfracostBreakdownPreApply,omitempty" url:"-"`
	InfracostBreakdownPostApply interface{} `json:"InfracostBreakdownPostApply,omitempty" url:"-"`
	TfDrift                     interface{} `json:"TfDrift,omitempty" url:"-"`
	GitHubComSync               interface{} `json:"GitHubComSync,omitempty" url:"-"`
	VcsTriggers                 interface{} `json:"VCSTriggers,omitempty" url:"-"`
	VcsConfig                   interface{} `json:"VCSConfig,omitempty" url:"-"`
	RunnerConstraints           interface{} `json:"RunnerConstraints,omitempty" url:"-"`
	CfStateCleaned              interface{} `json:"CfStateCleaned,omitempty" url:"-"`
	CfDrift                     interface{} `json:"CfDrift,omitempty" url:"-"`
	K8sResources                interface{} `json:"K8sResources,omitempty" url:"-"`
	K8sDrift                    interface{} `json:"K8sDrift,omitempty" url:"-"`
	AnsibleOutputs              interface{} `json:"AnsibleOutputs,omitempty" url:"-"`
	AnsiblePlan                 interface{} `json:"AnsiblePlan,omitempty" url:"-"`
	AnsibleDrift                interface{} `json:"AnsibleDrift,omitempty" url:"-"`
	BicepResources              interface{} `json:"BicepResources,omitempty" url:"-"`
	SgCustomWorkflowRunFacts    interface{} `json:"SGCustomWorkflowRunFacts,omitempty" url:"-"`
	TerragruntDrift             interface{} `json:"TerragruntDrift,omitempty" url:"-"`
}

// OrganizationWorkflowsListAllResponse lists workflows across the organization; the API answers
// 204 with no body when there are none.
type OrganizationWorkflowsListAllResponse struct {
	Msg              []*OrganizationWorkflow `json:"msg,omitempty" url:"-"`
	Lastevaluatedkey *string                 `json:"lastevaluatedkey,omitempty" url:"-"`
}

type OrganizationWorkflowsCountResponse struct {
	Msg *int `json:"msg,omitempty" url:"-"`
}

// WorkflowsBulkActionRequest activates or deactivates workflows and stacks. ResourceIds are
// paths relative to the organization such as /wfgrps/group/wfs/workflow or /wfgrps/group/stacks/stack.
type WorkflowsBulkActionRequest struct {
	// activate or deactivate
	Action      string   `json:"action" url:"-"`
	ResourceIds []string `json:"resourceIds" url:"-"`
}

type BulkActionResult struct {
	Updated *int     `json:"updated,omitempty" url:"-"`
	Failed  *int     `json:"failed,omitempty" url:"-"`
	Errors  []string `json:"errors,omitempty" url:"-"`
}

// WorkflowsBulkActionResponse reports the outcome; the API answers 207 when some batches failed.
type WorkflowsBulkActionResponse struct {
	Msg *BulkActionResult `json:"msg,omitempty" url:"-"`
}
