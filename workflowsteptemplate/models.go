package workflowsteptemplate

import (
	"encoding/json"

	"github.com/StackGuardian/sg-sdk-go/core"
)

type WorkflowStepTemplateSourceConfigKindEnum string

const WorkflowStepTemplateSourceConfigKindDockerImageEnum WorkflowStepTemplateSourceConfigKindEnum = "DOCKER_IMAGE"

type IsPublicEnum string

const (
	IsPublicEnumZero IsPublicEnum = "0"
	IsPublicEnumOne  IsPublicEnum = "1"
)

type SourceConfigDestKindEnum string

const SourceConfigDestKindContainerRegistryEnum SourceConfigDestKindEnum = "CONTAINER_REGISTRY"

type TemplateTypeEnum string

const (
	TemplateTypeIACEnum          = "IAC"
	TemplateTypeIACGroupEnum     = "IAC_GROUP"
	TemplateTypeIACPolicyEnum    = "IAC_POLICY"
	TemplateTypeWorkflowStepEnum = "WORKFLOW_STEP"
)

type WorkflowStepRuntimeSourceConfig struct {
	IsPrivate              *bool   `json:"isPrivate,omitempty" url:"isPrivate,omitempty"`
	Auth                   *string `json:"auth,omitempty" url:"auth,omitempty"`
	DockerImage            string  `json:"dockerImage,omitempty" url:"dockerImage,omitempty"`
	DockerRegistryUsername *string `json:"dockerRegistryUsername,omitempty" url:"dockerRegistryUsername,omitempty"`
	LocalWorkspaceDir      *string `json:"localWorkspaceDir,omitempty" url:"localWorkspaceDir,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

type WorkflowStepRuntimeSource struct {
	Config               *WorkflowStepRuntimeSourceConfig `json:"config,omitempty" url:"config,omitempty"`
	SourceConfigDestKind SourceConfigDestKindEnum         `json:"sourceConfigDestKind,omitempty" url:"sourceConfigDestKind,omitempty"`
	AdditionalConfig     map[string]interface{}           `json:"additionalConfig,omitempty" url:"additionalConfig,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

type UpdateWorkflowStepTemplateRequestModel struct {
	SourceConfigKind *core.Optional[WorkflowStepTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	ShortDescription *core.Optional[string]                                   `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	RuntimeSource    *core.Optional[WorkflowStepRuntimeSource]                `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	OwnerOrg         *core.Optional[string]                                   `json:"OwnerOrg" url:"OwnerOrg"`
	Tags             *core.Optional[[]string]                                 `json:"Tags,omitempty" url:"Tags,omitempty"`
	// Contextual tags to give context to your tags
	ContextTags    *core.Optional[map[string]string] `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	IsActive       *core.Optional[IsPublicEnum]      `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic       *core.Optional[IsPublicEnum]      `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	TemplateName   *core.Optional[string]            `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	TemplateType   *core.Optional[TemplateTypeEnum]  `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	LatestRevision *core.Optional[int]               `json:"LatestRevision,omitempty" url:"LatestRevision,omitempty"`
	NextRevision   *core.Optional[int]               `json:"NextRevision,omitempty" url:"NextRevision,omitempty"`
	SharedOrgsList *core.Optional[[]string]          `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
}

type UpdateWorkflowStepTemplateResponse struct {
	Id               string                                   `json:"Id,omitempty" url:"Id,omitempty"`
	ShortDescription *string                                  `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	SourceConfigKind WorkflowStepTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	RuntimeSource    *WorkflowStepRuntimeSource               `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	TemplateId       *string                                  `json:"TemplateId,omitempty" url:"TemplateId,omitempty"`
	OwnerOrg         string                                   `json:"OwnerOrg" url:"OwnerOrg"`
	Tags             []string                                 `json:"Tags,omitempty" url:"Tags,omitempty"`
	// Contextual tags to give context to your tags
	ContextTags    map[string]string `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	IsActive       *IsPublicEnum     `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic       *IsPublicEnum     `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	TemplateName   string            `json:"TemplateName,omitempty" url:"TemplateName"`
	TemplateType   TemplateTypeEnum  `json:"TemplateType,omitempty" url:"TemplateType"`
	LatestRevision *int              `json:"LatestRevision,omitempty" url:"LatestRevision,omitempty"`
	NextRevision   *int              `json:"NextRevision,omitempty" url:"NextRevision,omitempty"`
	SharedOrgsList []string          `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}
type UpdateWorkflowStepTemplateResponseModel struct {
	Msg  string                             `json:"msg,omitempty" url:"msg,omitempty"`
	Data UpdateWorkflowStepTemplateResponse `json:"data,omitempty" url:"data,omitempty"`
}

type ReadWorkflowStepTemplateResponseModel struct {
	Msg UpdateWorkflowStepTemplateResponse `json:"msg,omitempty" url:"msg,omitempty"`
}

type CreateWorkflowStepTemplate struct {
	TemplateType     TemplateTypeEnum                         `json:"TemplateType,omitempty" url:"TemplateType"`
	SourceConfigKind WorkflowStepTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	TemplateName     string                                   `json:"TemplateName,omitempty" url:"TemplateName"`
	ContextTags      map[string]string                        `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	ShortDescription *string                                  `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	Id               *string                                  `json:"Id,omitempty" url:"Id,omitempty"`
	IsActive         *IsPublicEnum                            `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *IsPublicEnum                            `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	OwnerOrg         string                                   `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	RuntimeSource    *WorkflowStepRuntimeSource               `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SharedOrgsList   []string                                 `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	Tags             []string                                 `json:"Tags,omitempty" url:"Tags,omitempty"`
}

type CreateWorkflowStepTemplateResponseData struct {
	Parent UpdateWorkflowStepTemplateResponse `json:"parent,omitempty" url:"parent,omitempty"`
}
type CreateWorkflowStepTemplateResponseModel struct {
	Msg  string                                 `json:"msg,omitempty" url:"msg,omitempty"`
	Data CreateWorkflowStepTemplateResponseData `json:"data,omitempty" url:"data,omitempty"`
}
