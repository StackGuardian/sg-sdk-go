package stacktemplates

import (
	"fmt"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/core"
)

const TemplateType = "IAC_GROUP"

type StackTemplateSourceConfigKindEnum string

const (
	StackTemplateSourceConfigKindTerraform         StackTemplateSourceConfigKindEnum = "TERRAFORM"
	StackTemplateSourceConfigKindOpentofu          StackTemplateSourceConfigKindEnum = "OPENTOFU"
	StackTemplateSourceConfigKindAnsiblePlaybook   StackTemplateSourceConfigKindEnum = "ANSIBLE_PLAYBOOK"
	StackTemplateSourceConfigKindHelm              StackTemplateSourceConfigKindEnum = "HELM"
	StackTemplateSourceConfigKindKubectl           StackTemplateSourceConfigKindEnum = "KUBECTL"
	StackTemplateSourceConfigKindCloudformation    StackTemplateSourceConfigKindEnum = "CLOUDFORMATION"
	StackTemplateSourceConfigKindDockerImage       StackTemplateSourceConfigKindEnum = "DOCKER_IMAGE"
	StackTemplateSourceConfigKindOpaRego           StackTemplateSourceConfigKindEnum = "OPA_REGO"
	StackTemplateSourceConfigKindSgPolicyFramework StackTemplateSourceConfigKindEnum = "SG_POLICY_FRAMEWORK"
	StackTemplateSourceConfigKindCheckov           StackTemplateSourceConfigKindEnum = "CHECKOV"
	StackTemplateSourceConfigKindSteampipe         StackTemplateSourceConfigKindEnum = "STEAMPIPE"
	StackTemplateSourceConfigKindMixed             StackTemplateSourceConfigKindEnum = "MIXED"
	StackTemplateSourceConfigKindCustom            StackTemplateSourceConfigKindEnum = "CUSTOM"
)

func NewStackTemplateSourceConfigKindEnumFromString(s string) (StackTemplateSourceConfigKindEnum, error) {
	switch s {
	case "TERRAFORM":
		return StackTemplateSourceConfigKindTerraform, nil
	case "OPENTOFU":
		return StackTemplateSourceConfigKindOpentofu, nil
	case "ANSIBLE_PLAYBOOK":
		return StackTemplateSourceConfigKindAnsiblePlaybook, nil
	case "HELM":
		return StackTemplateSourceConfigKindHelm, nil
	case "KUBECTL":
		return StackTemplateSourceConfigKindKubectl, nil
	case "CLOUDFORMATION":
		return StackTemplateSourceConfigKindCloudformation, nil
	case "DOCKER_IMAGE":
		return StackTemplateSourceConfigKindDockerImage, nil
	case "OPA_REGO":
		return StackTemplateSourceConfigKindOpaRego, nil
	case "SG_POLICY_FRAMEWORK":
		return StackTemplateSourceConfigKindSgPolicyFramework, nil
	case "CHECKOV":
		return StackTemplateSourceConfigKindCheckov, nil
	case "STEAMPIPE":
		return StackTemplateSourceConfigKindSteampipe, nil
	case "MIXED":
		return StackTemplateSourceConfigKindMixed, nil
	case "CUSTOM":
		return StackTemplateSourceConfigKindCustom, nil
	}
	var t StackTemplateSourceConfigKindEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t StackTemplateSourceConfigKindEnum) Ptr() *StackTemplateSourceConfigKindEnum {
	return &t
}

// CreateStackTemplateRequest corresponds to the StackTemplateCreate schema.
type CreateStackTemplateRequest struct {
	Id               *string                            `json:"Id,omitempty" url:"Id,omitempty"`
	OwnerOrg         string                             `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	TemplateType     sgsdkgo.TemplateTypeEnum           `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	SourceConfigKind *StackTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	Tags             []string                           `json:"Tags,omitempty" url:"Tags,omitempty"`
	ContextTags      map[string]string                  `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	IsActive         *sgsdkgo.IsPublicEnum              `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *sgsdkgo.IsPublicEnum              `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	TemplateName     string                             `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	SharedOrgsList   []string                           `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	ShortDescription *string                            `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
}

type CreateStackTemplateResponseData struct {
	Parent CreateStackTemplateRequest `json:"parent,omitempty" url:"parent,omitempty"`
}

type CreateStackTemplateResponseModel struct {
	Msg  string                          `json:"msg,omitempty" url:"msg,omitempty"`
	Data CreateStackTemplateResponseData `json:"data,omitempty" url:"data,omitempty"`
}

// ReadStackTemplateResponse holds the data returned when reading a stack template.
type ReadStackTemplateResponse struct {
	Id               *string                            `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateName     *string                            `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	TemplateType     *sgsdkgo.TemplateTypeEnum          `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	OwnerOrg         *string                            `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	ShortDescription *string                            `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	SourceConfigKind *StackTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	IsActive         *sgsdkgo.IsPublicEnum              `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *sgsdkgo.IsPublicEnum              `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	SharedOrgsList   []string                           `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	Tags             []string                           `json:"Tags,omitempty" url:"Tags,omitempty"`
	ContextTags      map[string]string                  `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
}

type ReadStackTemplateResponseModel struct {
	Msg ReadStackTemplateResponse `json:"msg,omitempty" url:"msg,omitempty"`
}

// UpdateStackTemplateRequest corresponds to the PatchedStackTemplateUpdate schema.
type UpdateStackTemplateRequest struct {
	OwnerOrg         *core.Optional[string]                            `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	Tags             *core.Optional[[]string]                          `json:"Tags,omitempty" url:"Tags,omitempty"`
	ContextTags      *core.Optional[map[string]string]                 `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	IsActive         *core.Optional[sgsdkgo.IsPublicEnum]              `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *core.Optional[sgsdkgo.IsPublicEnum]              `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	TemplateName     *core.Optional[string]                            `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	SharedOrgsList   *core.Optional[[]string]                          `json:"SharedOrgsList,omitempty" url:"SharedOrgsList,omitempty"`
	ShortDescription *core.Optional[string]                            `json:"ShortDescription,omitempty" url:"ShortDescription,omitempty"`
	SourceConfigKind *core.Optional[StackTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
}

type UpdateStackTemplateResponseModel struct {
	Msg  string                    `json:"msg,omitempty" url:"msg,omitempty"`
	Data ReadStackTemplateResponse `json:"data,omitempty" url:"data,omitempty"`
}
