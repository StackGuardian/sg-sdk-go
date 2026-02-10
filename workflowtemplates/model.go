package workflowtemplates

import (
	"fmt"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/core"
)

type WorkflowTemplateSourceConfigKindEnum string

const (
	WorkflowTemplateSourceConfigKindTerraform       WorkflowTemplateSourceConfigKindEnum = "TERRAFORM"
	WorkflowTemplateSourceConfigKindOpentofu        WorkflowTemplateSourceConfigKindEnum = "OPENTOFU"
	WorkflowTemplateSourceConfigKindAnsiblePlaybook WorkflowTemplateSourceConfigKindEnum = "ANSIBLE_PLAYBOOK"
	WorkflowTemplateSourceConfigKindHelm            WorkflowTemplateSourceConfigKindEnum = "HELM"
	WorkflowTemplateSourceConfigKindKubectl         WorkflowTemplateSourceConfigKindEnum = "KUBECTL"
	WorkflowTemplateSourceConfigKindCloudformation  WorkflowTemplateSourceConfigKindEnum = "CLOUDFORMATION"
	WorkflowTemplateSourceConfigKindCustom          WorkflowTemplateSourceConfigKindEnum = "CUSTOM"
)

func NewWorkflowTemplateSourceConfigKindEnumFromString(s string) (WorkflowTemplateSourceConfigKindEnum, error) {
	switch s {
	case "TERRAFORM":
		return WorkflowTemplateSourceConfigKindTerraform, nil
	case "OPENTOFU":
		return WorkflowTemplateSourceConfigKindOpentofu, nil
	case "ANSIBLE_PLAYBOOK":
		return WorkflowTemplateSourceConfigKindAnsiblePlaybook, nil
	case "HELM":
		return WorkflowTemplateSourceConfigKindHelm, nil
	case "KUBECTL":
		return WorkflowTemplateSourceConfigKindKubectl, nil
	case "CLOUDFORMATION":
		return WorkflowTemplateSourceConfigKindCloudformation, nil
	case "CUSTOM":
		return WorkflowTemplateSourceConfigKindCustom, nil
	}
	var t WorkflowTemplateSourceConfigKindEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t WorkflowTemplateSourceConfigKindEnum) Ptr() *WorkflowTemplateSourceConfigKindEnum {
	return &t
}

type SourceConfigDestKindEnum string

const (
	SourceConfigDestKindEnumGithubCom       SourceConfigDestKindEnum = "GITHUB_COM"
	SourceConfigDestKindEnumGithubAppCustom SourceConfigDestKindEnum = "GITHUB_APP_CUSTOM"
	SourceConfigDestKindEnumGitOther        SourceConfigDestKindEnum = "GIT_OTHER"
	SourceConfigDestKindEnumBitbucketOrg    SourceConfigDestKindEnum = "BITBUCKET_ORG"
	SourceConfigDestKindEnumGitlabCom       SourceConfigDestKindEnum = "GITLAB_COM"
	SourceConfigDestKindEnumAzureDevops     SourceConfigDestKindEnum = "AZURE_DEVOPS"
	SourceConfigDestKindEnumAzureDevopsSp   SourceConfigDestKindEnum = "AZURE_DEVOPS_SP"
)

func (t SourceConfigDestKindEnum) Ptr() *SourceConfigDestKindEnum {
	return &t
}

type RuntimeSourceConfig struct {
	Auth                    *string `json:"auth,omitempty" url:"auth"`
	GitCoreAutoCRLF         *bool   `json:"gitCoreAutoCRLF,omitempty" url:"gitCoreAutoCRLF,omitempty"`
	GitSparseCheckoutConfig *string `json:"gitSparseCheckoutConfig,omitempty" url:"gitSparseCheckoutConfig,omitempty"`
	IncludeSubModule        *bool   `json:"includeSubModule,omitempty" url:"includeSubModule,omitempty"`
	IsPrivate               *bool   `json:"isPrivate,omitempty" url:"isPrivate,omitempty"`
	Ref                     *string `json:"ref,omitempty" url:"ref,omitempty"`
	Repo                    string  `json:"repo,omitempty" url:"repo,omitempty"`
	WorkingDir              *string `json:"workingDir,omitempty" url:"workingDir,omitempty"`
}

type RuntimeSource struct {
	SourceConfigDestKind *SourceConfigDestKindEnum `json:"sourceConfigDestKind,omitempty" url:"sourceConfigDestKind,omitempty"`
	Config               *RuntimeSourceConfig      `json:"config,omitempty" url:"config,omitempty"`
}

type VCSTriggersCreateTagCreateRevision struct {
	Enabled *bool `json:"enabled,omitempty" url:"enabled,omitempty"`
}

type VCSTriggersCreateTag struct {
	CreateRevision *VCSTriggersCreateTagCreateRevision `json:"createRevision,omitempty" url:"createRevision,omitempty"`
}

type VCSTriggersTypeEnum string

const (
	VCSTriggersTypeEnumGithubCom       VCSTriggersTypeEnum = "GITHUB_COM"
	VCSTriggersTypeEnumGithubAppCustom VCSTriggersTypeEnum = "GITHUB_APP_CUSTOM"
	VCSTriggersTypeEnumGitlabOauthSsh  VCSTriggersTypeEnum = "GITLAB_OAUTH_SSH"
	VCSTriggersTypeEnumGitlabCom       VCSTriggersTypeEnum = "GITLAB_COM"
)

func (t VCSTriggersTypeEnum) Ptr() *VCSTriggersTypeEnum {
	return &t
}

type VCSTriggers struct {
	CreateTag *VCSTriggersCreateTag `json:"create_tag,omitempty" url:"create_tag,omitempty"`
	Type      *VCSTriggersTypeEnum  `json:"type,omitempty" url:"type,omitempty"`
}

type ReadWorkflowTemplateResponse struct {
	Id               *string                               `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateName     *string                               `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	TemplateType     *sgsdkgo.TemplateTypeEnum             `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	OwnerOrg         *string                               `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	ShortDescription *string                               `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	SourceConfigKind *WorkflowTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	IsActive         *sgsdkgo.IsPublicEnum                 `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *sgsdkgo.IsPublicEnum                 `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	RuntimeSource    *RuntimeSource                        `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SharedOrgsList   []string                              `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	Tags             []string                              `json:"Tags,omitempty" url:"Tags,omitempty"`
	VCSTriggers      *VCSTriggers                          `json:"VCSTriggers,omitempty" url:"VCSTriggers,omitempty"`
	ContextTags      map[string]string                     `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
}

type ReadWorkflowTemplateResponseModel struct {
	Msg ReadWorkflowTemplateResponse `json:"msg,omitempty" url:"msg,omitempty"`
}

type CreateWorkflowTemplateRequest struct {
	Id               *string                               `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateName     string                                `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	TemplateType     sgsdkgo.TemplateTypeEnum              `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	OwnerOrg         string                                `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	SourceConfigKind *WorkflowTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	IsActive         *sgsdkgo.IsPublicEnum                 `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *sgsdkgo.IsPublicEnum                 `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	RuntimeSource    *RuntimeSource                        `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SharedOrgsList   []string                              `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	Tags             []string                              `json:"Tags,omitempty" url:"Tags,omitempty"`
	VCSTriggers      *VCSTriggers                          `json:"VCSTriggers,omitempty" url:"VCSTriggers,omitempty"`
	ContextTags      map[string]string                     `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	ShortDescription *string                               `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
}

type CreateWorkflowTemplateResponseData struct {
	Parent CreateWorkflowTemplateRequest `json:"parent,omitempty" url:"parent,omitempty"`
}
type CreateWorkflowTemplateResponseModel struct {
	Msg  string                             `json:"msg,omitempty" url:"msg,omitempty"`
	Data CreateWorkflowTemplateResponseData `json:"data,omitempty" url:"data,omitempty"`
}

type UpdateWorkflowTemplateResponse struct {
	Id               *core.Optional[string]                               `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateName     *core.Optional[string]                               `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	TemplateType     *core.Optional[sgsdkgo.TemplateTypeEnum]             `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	OwnerOrg         *core.Optional[string]                               `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	SourceConfigKind *core.Optional[WorkflowTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	IsActive         *core.Optional[sgsdkgo.IsPublicEnum]                 `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *core.Optional[sgsdkgo.IsPublicEnum]                 `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	RuntimeSource    *core.Optional[RuntimeSource]                        `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SharedOrgsList   *core.Optional[[]string]                             `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	Tags             *core.Optional[[]string]                             `json:"Tags,omitempty" url:"Tags,omitempty"`
	VCSTriggers      *core.Optional[VCSTriggers]                          `json:"VCSTriggers,omitempty" url:"VCSTriggers,omitempty"`
	ContextTags      *core.Optional[map[string]string]                    `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	ShortDescription *core.Optional[string]                               `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
}

type UpdateWorkflowTemplateResponseModel struct {
	Msg  string                         `json:"msg,omitempty" url:"msg,omitempty"`
	Data UpdateWorkflowTemplateResponse `json:"data,omitempty" url:"data,omitempty"`
}

type RuntimeSourceConfigUpdate struct {
	GitCoreAutoCRLF         *bool   `json:"gitCoreAutoCRLF,omitempty" url:"gitCoreAutoCRLF,omitempty"`
	GitSparseCheckoutConfig *string `json:"gitSparseCheckoutConfig,omitempty" url:"gitSparseCheckoutConfig,omitempty"`
	IncludeSubModule        *bool   `json:"includeSubModule,omitempty" url:"includeSubModule,omitempty"`
	IsPrivate               *bool   `json:"isPrivate,omitempty" url:"isPrivate,omitempty"`
	Ref                     *string `json:"ref,omitempty" url:"ref,omitempty"`
	WorkingDir              *string `json:"workingDir,omitempty" url:"workingDir,omitempty"`
}

type RuntimeSourceUpdate struct {
	SourceConfigDestKind *SourceConfigDestKindEnum  `json:"sourceConfigDestKind,omitempty" url:"sourceConfigDestKind,omitempty"`
	Config               *RuntimeSourceConfigUpdate `json:"config,omitempty" url:"config,omitempty"`
}

type UpdateWorkflowTemplateRequest struct {
	TemplateName     *core.Optional[string]                               `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	OwnerOrg         *core.Optional[string]                               `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	IsActive         *core.Optional[sgsdkgo.IsPublicEnum]                 `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *core.Optional[sgsdkgo.IsPublicEnum]                 `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	ShortDescription *core.Optional[string]                               `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	SourceConfigKind *core.Optional[WorkflowTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	RuntimeSource    *core.Optional[RuntimeSourceUpdate]                  `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SharedOrgsList   *core.Optional[[]string]                             `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	Tags             *core.Optional[[]string]                             `json:"Tags,omitempty" url:"Tags,omitempty"`
	VCSTriggers      *core.Optional[VCSTriggers]                          `json:"VCSTriggers,omitempty" url:"VCSTriggers,omitempty"`
	ContextTags      *core.Optional[map[string]string]                    `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
}
