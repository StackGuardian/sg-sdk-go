package stacktemplaterevisions

import (
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/stacktemplates"
)

type Deprecation struct {
	EffectiveDate *string `json:"effectiveDate,omitempty" url:"effectiveDate,omitempty"`
	Message       *string `json:"message,omitempty" url:"message,omitempty"`
}

// CreateStackTemplateRevisionRequest corresponds to the StackTemplateRevision schema.
type CreateStackTemplateRevisionRequest struct {
	TempalteName     string                                            `json:"TemplateName,omitempty" url:"TemplateName,omitempty"`
	TemplateType     string                                            `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	OwnerOrg         string                                            `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	SourceConfigKind *stacktemplates.StackTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	Tags             []string                                          `json:"Tags,omitempty" url:"Tags,omitempty"`
	ContextTags      map[string]string                                 `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	IsActive         *sgsdkgo.IsPublicEnum                             `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *sgsdkgo.IsPublicEnum                             `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Alias            string                                            `json:"Alias,omitempty" url:"Alias,omitempty"`
	Notes            string                                            `json:"Notes,omitempty" url:"Notes,omitempty"`
	Deprecation      *Deprecation                                      `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	LongDescription  *string                                           `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	WorkflowsConfig  *sgsdkgo.WorkflowsConfig                          `json:"WorkflowsConfig,omitempty" url:"WorkflowsConfig,omitempty"`
	Actions          map[string]*sgsdkgo.Actions                       `json:"Actions,omitempty" url:"Actions,omitempty"`
}

type CreateStackTemplateRevisionResponse struct {
	CreateStackTemplateRevisionRequest
	Id         string `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateId string `json:"TemplateId,omitempty" url:"TemplateId,omitempty"`
}

type CreateStackTemplateRevisionResponseModel struct {
	Msg  string `json:"msg,omitempty" url:"msg,omitempty"`
	Data struct {
		Revision CreateStackTemplateRevisionResponse `json:"revision,omitempty" url:"revision,omitempty"`
	} `json:"data,omitempty" url:"data,omitempty"`
}

type ReadStackTemplateRevisionModel struct {
	CreateStackTemplateRevisionRequest
	Id         *string `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateId string  `json:"TemplateId,omitempty" url:"TemplateId,omitempty"`
}

type ReadStackTemplateRevisionResponseModel struct {
	Msg ReadStackTemplateRevisionModel `json:"msg,omitempty" url:"msg,omitempty"`
}

// UpdateStackTemplateRevisionRequest corresponds to the PatchedStackTemplateRevisionUpdate schema.
type UpdateStackTemplateRevisionRequest struct {
	OwnerOrg         *core.Optional[string]                                           `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	Tags             *core.Optional[[]string]                                         `json:"Tags,omitempty" url:"Tags,omitempty"`
	ContextTags      *core.Optional[map[string]string]                                `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	IsActive         *core.Optional[sgsdkgo.IsPublicEnum]                             `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *core.Optional[sgsdkgo.IsPublicEnum]                             `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Alias            *core.Optional[string]                                           `json:"Alias,omitempty" url:"Alias,omitempty"`
	Notes            *core.Optional[string]                                           `json:"Notes,omitempty" url:"Notes,omitempty"`
	Deprecation      *core.Optional[Deprecation]                                      `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	WorkflowsConfig  *core.Optional[sgsdkgo.WorkflowsConfig]                          `json:"WorkflowsConfig,omitempty" url:"WorkflowsConfig,omitempty"`
	Actions          *core.Optional[map[string]*sgsdkgo.Actions]                      `json:"Actions,omitempty" url:"Actions,omitempty"`
	SourceConfigKind *core.Optional[stacktemplates.StackTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
}

type UpdateStackTemplateRevisionResponseModel struct {
	Msg  string                             `json:"msg,omitempty" url:"msg,omitempty"`
	Data UpdateStackTemplateRevisionRequest `json:"data,omitempty" url:"data,omitempty"`
}
