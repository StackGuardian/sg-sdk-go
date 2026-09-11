package workflowgroups

// ListAllResourcesRequest filters the resources (workflows, stacks and child groups) of a workflow group.
type ListAllResourcesRequest struct {
	Limit            *int    `json:"-" url:"limit,omitempty"`
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	ResourceNames    *string `json:"-" url:"resourceNames,omitempty"`
	Description      *string `json:"-" url:"description,omitempty"`
	Tags             *string `json:"-" url:"tags,omitempty"`
	SearchQuery      *string `json:"-" url:"searchQuery,omitempty"`
	LatestStatuses   *string `json:"-" url:"latestStatuses,omitempty"`
	// Comma-separated resource types: WORKFLOW, STACK, WORKFLOW_GROUP
	ResourceTypes *string `json:"-" url:"resourceTypes,omitempty"`
	// Comma-separated key:value pairs
	ContextTags *string `json:"-" url:"contextTags,omitempty"`
	// 0 or 1
	IsActive *string `json:"-" url:"IsActive,omitempty"`
}

// WorkflowGroupResource is one entry of a workflow group listing; discriminate on ResourceType.
type WorkflowGroupResource struct {
	ParentId              *string                `json:"ParentId,omitempty" url:"-"`
	ResourceId            *string                `json:"ResourceId,omitempty" url:"-"`
	ResourceName          *string                `json:"ResourceName,omitempty" url:"-"`
	SubResourceId         *string                `json:"SubResourceId,omitempty" url:"-"`
	ResourceType          *string                `json:"ResourceType,omitempty" url:"-"`
	IsActive              *string                `json:"IsActive,omitempty" url:"-"`
	Description           *string                `json:"Description,omitempty" url:"-"`
	Tags                  []string               `json:"Tags,omitempty" url:"-"`
	Authors               []string               `json:"Authors,omitempty" url:"-"`
	CreatedAt             *int64                 `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt            *int64                 `json:"ModifiedAt,omitempty" url:"-"`
	LatestWfStatus        *string                `json:"LatestWfStatus,omitempty" url:"-"`
	LatestTerraformAction *string                `json:"LatestTerraformAction,omitempty" url:"-"`
	LatestWfrunStatus     *string                `json:"LatestWfrunStatus,omitempty" url:"-"`
	LatestWfrunStatusKey  *string                `json:"LatestWfrunStatusKey,omitempty" url:"-"`
	WfType                *string                `json:"WfType,omitempty" url:"-"`
	TemplateGroupId       *string                `json:"TemplateGroupId,omitempty" url:"-"`
	WorkflowsConfig       interface{}            `json:"WorkflowsConfig,omitempty" url:"-"`
	TemplatesConfig       map[string]interface{} `json:"TemplatesConfig,omitempty" url:"-"`
	ContextTags           map[string]interface{} `json:"ContextTags,omitempty" url:"-"`
	WfStepsConfig         interface{}            `json:"WfStepsConfig,omitempty" url:"-"`
}

// WorkflowGroupResourcesListAllResponse lists a group's resources; the API answers 204 with no body when there are none.
type WorkflowGroupResourcesListAllResponse struct {
	Msg              []*WorkflowGroupResource `json:"msg,omitempty" url:"-"`
	Lastevaluatedkey *string                  `json:"lastevaluatedkey,omitempty" url:"-"`
}
