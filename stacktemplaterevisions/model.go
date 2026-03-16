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

// StackTemplateRevisionWorkflowsConfig corresponds to the WorkflowsConfig schema
// for stack template revisions.
type StackTemplateRevisionWorkflowsConfig struct {
	Workflows []*StackTemplateRevisionWorkflow `json:"workflows,omitempty" url:"workflows,omitempty"`
}

// StackTemplateRevisionWorkflow corresponds to the WorkflowsConfigWorkflow schema
// for stack template revisions — only the fields accepted by the API are included.
type StackTemplateRevisionWorkflow struct {
	Id                        *string                             `json:"id,omitempty" url:"id,omitempty"`
	TemplateId                *string                             `json:"templateId,omitempty" url:"templateId,omitempty"`
	ResourceName              *string                             `json:"ResourceName,omitempty" url:"ResourceName,omitempty"`
	WfStepsConfig             []*sgsdkgo.WfStepsConfig            `json:"WfStepsConfig,omitempty" url:"WfStepsConfig,omitempty"`
	TerraformConfig           *sgsdkgo.TerraformConfig            `json:"TerraformConfig,omitempty" url:"TerraformConfig,omitempty"`
	EnvironmentVariables      []*sgsdkgo.EnvVars                  `json:"EnvironmentVariables,omitempty" url:"EnvironmentVariables,omitempty"`
	DeploymentPlatformConfig  []*sgsdkgo.DeploymentPlatformConfig `json:"DeploymentPlatformConfig,omitempty" url:"DeploymentPlatformConfig,omitempty"`
	UserSchedules             []*sgsdkgo.UserSchedules            `json:"UserSchedules,omitempty" url:"UserSchedules,omitempty"`
	MiniSteps                 *sgsdkgo.MiniStepsSchema            `json:"MiniSteps,omitempty" url:"MiniSteps,omitempty"`
	Approvers                 []string                            `json:"Approvers,omitempty" url:"Approvers,omitempty"`
	NumberOfApprovalsRequired *int                                `json:"NumberOfApprovalsRequired,omitempty" url:"NumberOfApprovalsRequired,omitempty"`
	RunnerConstraints         *sgsdkgo.RunnerConstraints          `json:"RunnerConstraints,omitempty" url:"RunnerConstraints,omitempty"`
	UserJobCpu                *int                                `json:"UserJobCPU,omitempty" url:"UserJobCPU,omitempty"`
	UserJobMemory             *int                                `json:"UserJobMemory,omitempty" url:"UserJobMemory,omitempty"`
	VcsConfig                 *sgsdkgo.VcsConfig                  `json:"VCSConfig,omitempty" url:"VCSConfig,omitempty"`
	IacInputData              *sgsdkgo.TemplatesIacInputData      `json:"iacInputData,omitempty" url:"iacInputData,omitempty"`
	InputSchemas              []*sgsdkgo.InputSchemas             `json:"inputSchemas,omitempty" url:"inputSchemas,omitempty"`
}

// CreateStackTemplateRevisionRequest corresponds to the StackTemplateRevision schema.
type CreateStackTemplateRevisionRequest struct {
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
	WorkflowsConfig  *StackTemplateRevisionWorkflowsConfig             `json:"WorkflowsConfig,omitempty" url:"WorkflowsConfig,omitempty"`
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
	LongDescription  *string                                                          `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	IsActive         *core.Optional[sgsdkgo.IsPublicEnum]                             `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic         *core.Optional[sgsdkgo.IsPublicEnum]                             `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Alias            *core.Optional[string]                                           `json:"Alias,omitempty" url:"Alias,omitempty"`
	Notes            *core.Optional[string]                                           `json:"Notes,omitempty" url:"Notes,omitempty"`
	Deprecation      *core.Optional[Deprecation]                                      `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	WorkflowsConfig  *core.Optional[StackTemplateRevisionWorkflowsConfig]             `json:"WorkflowsConfig,omitempty" url:"WorkflowsConfig,omitempty"`
	Actions          *core.Optional[map[string]*sgsdkgo.Actions]                      `json:"Actions,omitempty" url:"Actions,omitempty"`
	SourceConfigKind *core.Optional[stacktemplates.StackTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
}

type UpdateStackTemplateRevisionResponseModel struct {
	Msg  string                             `json:"msg,omitempty" url:"msg,omitempty"`
	Data UpdateStackTemplateRevisionRequest `json:"data,omitempty" url:"data,omitempty"`
}
