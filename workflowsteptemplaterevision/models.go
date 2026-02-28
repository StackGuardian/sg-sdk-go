package workflowsteptemplaterevision

import (
	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/workflowsteptemplate"
)

type Deprecation struct {
	EffectiveDate *string `json:"effectiveDate,omitempty" url:"effectiveDate,omitempty"`
	Message       *string `json:"message,omitempty" url:"message,omitempty"`
}

type CreateWorkflowStepTemplateRevisionModel struct {
	TemplateType     workflowsteptemplate.TemplateTypeEnum                         `json:"TemplateType,omitempty" url:"TemplateType"`
	SourceConfigKind workflowsteptemplate.WorkflowStepTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	LongDescription  *string                                                       `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	ContextTags      map[string]string                                             `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	OwnerOrg         string                                                        `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	Tags             []string                                                      `json:"Tags,omitempty" url:"Tags,omitempty"`
	Alias            *string                                                       `json:"Alias,omitempty" url:"Alias,omitempty"`
	Notes            *string                                                       `json:"Notes,omitempty" url:"Notes,omitempty"`
	IsActive         *workflowsteptemplate.IsPublicEnum                            `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *workflowsteptemplate.IsPublicEnum                            `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	RuntimeSource    *workflowsteptemplate.WorkflowStepRuntimeSource               `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	Deprecation      *Deprecation                                                  `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
}

type WorkflowStepTemplateRevisionResponseData struct {
	Id               string
	Alias            *string                                                       `json:"Alias,omitempty" url:"Alias,omitempty"`
	ContextTags      map[string]string                                             `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	Deprecation      *Deprecation                                                  `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	LongDescription  *string                                                       `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	IsActive         *workflowsteptemplate.IsPublicEnum                            `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *workflowsteptemplate.IsPublicEnum                            `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Notes            *string                                                       `json:"Notes,omitempty" url:"Notes,omitempty"`
	OwnerOrg         string                                                        `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	RuntimeSource    *workflowsteptemplate.WorkflowStepRuntimeSource               `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SourceConfigKind workflowsteptemplate.WorkflowStepTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	Tags             []string                                                      `json:"Tags,omitempty" url:"Tags,omitempty"`
	TemplateType     workflowsteptemplate.TemplateTypeEnum                         `json:"TemplateType,omitempty" url:"TemplateType"`
}

type CreateWorkflowStepTemplateRevisionResponseModel struct {
	Msg  *string `json:"msg,omitempty" url:"msg,omitempty"`
	Data struct {
		Revision WorkflowStepTemplateRevisionResponseData `json:"revision,omitempty" url:"revision,omitempty"`
	} `json:"data,omitempty" url:"data,omitempty"`
}

type UpdateWorkflowStepTemplateRevisionModel struct {
	TemplateType     *core.Optional[workflowsteptemplate.TemplateTypeEnum]                         `json:"TemplateType,omitempty" url:"TemplateType"`
	ContextTags      *core.Optional[map[string]string]                                             `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	LongDescription  *core.Optional[string]                                                        `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	IsActive         *core.Optional[workflowsteptemplate.IsPublicEnum]                             `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *core.Optional[workflowsteptemplate.IsPublicEnum]                             `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Alias            *core.Optional[string]                                                        `json:"Alias,omitempty" url:"Alias,omitempty"`
	Deprecation      *core.Optional[Deprecation]                                                   `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	Notes            *core.Optional[string]                                                        `json:"Notes,omitempty" url:"Notes,omitempty"`
	OwnerOrg         string                                                                        `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	RuntimeSource    *core.Optional[workflowsteptemplate.WorkflowStepRuntimeSource]                `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	SourceConfigKind *core.Optional[workflowsteptemplate.WorkflowStepTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	Tags             *core.Optional[[]string]                                                      `json:"Tags,omitempty" url:"Tags,omitempty"`
}

type UpdateWorkflowStepTemplateRevisionResponseModel struct {
	Msg  *string                                  `json:"msg,omitempty" url:"msg,omitempty"`
	Data WorkflowStepTemplateRevisionResponseData `json:"data,omitempty" url:"data,omitempty"`
}

type ReadWorkflowStepTemplateRevisionResponseModel struct {
	Msg *WorkflowStepTemplateRevisionResponseData `json:"msg,omitempty" url:"msg,omitempty"`
}
