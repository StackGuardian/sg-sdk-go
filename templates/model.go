package templates

import (
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
)

// TemplateArtifactRequest carries the caller's organization for template artifact endpoints.
type TemplateArtifactRequest struct {
	// Current organization name of the user, e.g. my-sg-org
	SgOrgid string `json:"-" url:"-"`
}

type TemplateArtifactUrl struct {
	SignedUrl    string `json:"signedUrl" url:"-"`
	ArtifactPath string `json:"artifactPath" url:"-"`
}

type TemplateArtifactUrlResponse struct {
	Msg  *string              `json:"msg,omitempty" url:"-"`
	Data *TemplateArtifactUrl `json:"data,omitempty" url:"-"`
}

type DeleteTemplateArtifactResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

type ListAllTemplateArtifactsRequest struct {
	// Current organization name of the user, e.g. my-sg-org
	SgOrgid string `json:"-" url:"-"`
	// Only list artifacts whose path starts with this prefix
	ArtifactPrefix *string `json:"-" url:"artifactPrefix,omitempty"`
	// Maximum number of keys to return
	MaxKeys *int `json:"-" url:"maxKeys,omitempty"`
	// Pagination: start after this key
	StartAfterKey *string `json:"-" url:"startAfterKey,omitempty"`
}

type ListAllTemplateArtifactsResponse struct {
	Msg  *string                `json:"msg,omitempty" url:"-"`
	Data map[string]interface{} `json:"data,omitempty" url:"-"`
}

// UpdateSubscriptionRequest lists the templates to subscribe to or unsubscribe from,
// keyed by template id, per subscription type.
type UpdateSubscriptionRequest struct {
	WfStepSubscriptions   map[string]interface{} `json:"WfStepSubscriptions,omitempty" url:"-"`
	IacSubscriptions      map[string]interface{} `json:"IACSubscriptions,omitempty" url:"-"`
	PolicySubscriptions   map[string]interface{} `json:"PolicySubscriptions,omitempty" url:"-"`
	IacGroupSubscriptions map[string]interface{} `json:"IACGroupSubscriptions,omitempty" url:"-"`
}

type UpdateSubscriptionResponse struct {
	Msg *sgsdkgo.Subscription `json:"msg,omitempty" url:"-"`
}

// CreateSubscriptionRequest creates a subscription set. All four maps are required by the API;
// send empty objects for unused types. Values are keyed by template id.
type CreateSubscriptionRequest struct {
	ResourceName          *string                `json:"ResourceName,omitempty" url:"-"`
	IacSubscriptions      map[string]interface{} `json:"IACSubscriptions" url:"-"`
	WfStepSubscriptions   map[string]interface{} `json:"WfStepSubscriptions" url:"-"`
	PolicySubscriptions   map[string]interface{} `json:"PolicySubscriptions" url:"-"`
	IacGroupSubscriptions map[string]interface{} `json:"IACGroupSubscriptions" url:"-"`
}

type CreateSubscriptionResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

type ListAllPublicTemplatesRequest struct {
	Limit            *int    `json:"-" url:"limit,omitempty"`
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	// Parent template id including the owner org (e.g. /stackguardian/aws-s3-demo-website) to list its revisions
	TemplateId *string `json:"-" url:"TemplateId,omitempty"`
	// Comma-separated owner organizations
	OwnerOrgs   *string `json:"-" url:"OwnerOrgs,omitempty"`
	SearchQuery *string `json:"-" url:"SearchQuery,omitempty"`
	// Comma-separated key or key:value pairs
	ContextTags       *string `json:"-" url:"ContextTags,omitempty"`
	SourceConfigKinds *string `json:"-" url:"SourceConfigKinds,omitempty"`
}

type ListAllPublicTemplatesV2Request struct {
	Limit            *int    `json:"-" url:"limit,omitempty"`
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	IsActive         *string `json:"-" url:"isActive,omitempty"`
	TemplateName     *string `json:"-" url:"templateName,omitempty"`
	GitHubComRepoID  *string `json:"-" url:"gitHubComRepoID,omitempty"`
	SearchQuery      *string `json:"-" url:"searchQuery,omitempty"`
	// Comma-separated values
	Tags              *string `json:"-" url:"tags,omitempty"`
	TemplateTypes     *string `json:"-" url:"templateTypes,omitempty"`
	SourceConfigKinds *string `json:"-" url:"sourceConfigKinds,omitempty"`
	ContextTags       *string `json:"-" url:"contextTags,omitempty"`
	OwnerOrgs         *string `json:"-" url:"ownerOrgs,omitempty"`
}

type GetTemplateInputSchemaRequest struct {
	// Which InputSchemas entry to decode: FORM_JSONSCHEMA, RAW_JSON, NO_CODE_JSON or TIRITH_JSON
	SchemaType string `json:"-" url:"SchemaType"`
}

type TemplateInputSchemaResponse struct {
	Msg map[string]interface{} `json:"msg,omitempty" url:"-"`
}
