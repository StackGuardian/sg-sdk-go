package resources

type MoveResourcesRequest struct {
	SourceResourceIds   []string `json:"SourceResourceIds" url:"-"`
	DestinationParentId string   `json:"DestinationParentId" url:"-"`
}

type MoveResourcesResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

// SearchResourcesRequest filters an organization-wide resource search. Enum-valued
// lists (ResourceTypes, WorkflowTypes, statuses) take the API's string constants.
type SearchResourcesRequest struct {
	SearchQuery         *string     `json:"SearchQuery,omitempty" url:"-"`
	Limit               *int        `json:"Limit,omitempty" url:"-"`
	LastEvaluatedKey    *string     `json:"LastEvaluatedKey,omitempty" url:"-"`
	ResourceTypes       []string    `json:"ResourceTypes,omitempty" url:"-"`
	SortBy              interface{} `json:"SortBy,omitempty" url:"-"`
	WorkflowTypes       []string    `json:"WorkflowTypes,omitempty" url:"-"`
	IacTemplate         *string     `json:"IacTemplate,omitempty" url:"-"`
	PolicyTemplate      *string     `json:"PolicyTemplate,omitempty" url:"-"`
	WfStepTemplate      *string     `json:"WfStepTemplate,omitempty" url:"-"`
	TemplateGroupId     *string     `json:"TemplateGroupId,omitempty" url:"-"`
	CloudIntegrationIds []string    `json:"CloudIntegrationIds,omitempty" url:"-"`
	GitIntegrationIds   []string    `json:"GitIntegrationIds,omitempty" url:"-"`
	RunnerGroupIds      []string    `json:"RunnerGroupIds,omitempty" url:"-"`
	Tags                []string    `json:"Tags,omitempty" url:"-"`
	Authors             []string    `json:"Authors,omitempty" url:"-"`
	WorkflowStatuses    []string    `json:"WorkflowStatuses,omitempty" url:"-"`
	StackStatuses       []string    `json:"StackStatuses,omitempty" url:"-"`
	WfrunStatuses       []string    `json:"WfrunStatuses,omitempty" url:"-"`
	RunnerGroupStatuses []string    `json:"RunnerGroupStatuses,omitempty" url:"-"`
	Description         *string     `json:"Description,omitempty" url:"-"`
	ContextTags         []string    `json:"ContextTags,omitempty" url:"-"`
	WorkflowRunId       *string     `json:"WorkflowRunId,omitempty" url:"-"`
	ModifiedAtFrom      *int64      `json:"ModifiedAtFrom,omitempty" url:"-"`
	ModifiedAtTo        *int64      `json:"ModifiedAtTo,omitempty" url:"-"`
	CreatedAtFrom       *int64      `json:"CreatedAtFrom,omitempty" url:"-"`
	CreatedAtTo         *int64      `json:"CreatedAtTo,omitempty" url:"-"`
}

type SearchResourcesResponse struct {
	Msg              interface{}            `json:"msg,omitempty" url:"-"`
	Data             map[string]interface{} `json:"data,omitempty" url:"-"`
	Lastevaluatedkey *string                `json:"lastevaluatedkey,omitempty" url:"-"`
}

type ListResourceTagsRequest struct {
	ContextTagKey *string `json:"-" url:"ContextTagKey,omitempty"`
	Limit         *int    `json:"-" url:"Limit,omitempty"`
	Query         *string `json:"-" url:"Query,omitempty"`
	ResourceType  *string `json:"-" url:"ResourceType,omitempty"`
	TagKind       *string `json:"-" url:"TagKind,omitempty"`
}

type ListResourceTagsResponse struct {
	Msg  *string                  `json:"msg,omitempty" url:"-"`
	Data []map[string]interface{} `json:"data,omitempty" url:"-"`
}
