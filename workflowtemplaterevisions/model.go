package workflowtemplaterevisions

import (
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplates"
)

type DeploymentPlatformConfigKindEnum string

const (
	DeploymentPlatformConfigKindEnumAwsStatic   DeploymentPlatformConfigKindEnum = "AWS_STATIC"
	DeploymentPlatformConfigKindEnumAwsRbac     DeploymentPlatformConfigKindEnum = "AWS_RBAC"
	DeploymentPlatformConfigKindEnumAwsOidc     DeploymentPlatformConfigKindEnum = "AWS_OIDC"
	DeploymentPlatformConfigKindEnumAzureStatic DeploymentPlatformConfigKindEnum = "AZURE_STATIC"
	DeploymentPlatformConfigKindEnumAzureOidc   DeploymentPlatformConfigKindEnum = "AZURE_OIDC"
	DeploymentPlatformConfigKindEnumGcpStatic   DeploymentPlatformConfigKindEnum = "GCP_STATIC"
	DeploymentPlatformConfigKindEnumGcpOidc     DeploymentPlatformConfigKindEnum = "GCP_OIDC"
)

func (t DeploymentPlatformConfigKindEnum) Ptr() *DeploymentPlatformConfigKindEnum {
	return &t
}

type DeploymentPlatformConfigConfig struct {
	IntegrationId string  `json:"integrationId,omitempty" url:"integrationId,omitempty"`
	ProfileName   *string `json:"profileName,omitempty" url:"profileName,omitempty"`
}

type DeploymentPlatformConfig struct {
	Kind   DeploymentPlatformConfigKindEnum `json:"kind,omitempty" url:"kind,omitempty"`
	Config DeploymentPlatformConfigConfig   `json:"config,omitempty" url:"config,omitempty"`
}

type Deprecation struct {
	EffectiveDate *string `json:"effectiveDate,omitempty" url:"effectiveDate,omitempty"`
	Message       *string `json:"message,omitempty" url:"message,omitempty"`
}

type MinistepsNotificationRecepients struct {
	Recipients []string `json:"recipients,omitempty" url:"recipients,omitempty"`
}

type MinistepsWebhooksSchema struct {
	WebhookName   string  `json:"webhookName,omitempty" url:"webhookName,omitempty"`
	WebhookUrl    string  `json:"webhookUrl,omitempty" url:"webhookUrl,omitempty"`
	WebhookSecret *string `json:"webhookSecret,omitempty" url:"webhookSecret,omitempty"`
}

type MinistepsWfChainingSchema struct {
	WorkflowGroupId    string  `json:"workflowGroupId,omitempty" url:"workflowGroupId,omitempty"`
	StackId            *string `json:"stackId,omitempty" url:"stackId,omitempty"`
	WorkflowId         *string `json:"workflowId,omitempty" url:"workflowId,omitempty"`
	WorkflowRunPayload *string `json:"workflowRunPayload,omitempty" url:"workflowRunPayload,omitempty"`
	StackRunPayload    *string `json:"stackRunPayload,omitempty" url:"stackRunPayload,omitempty"`
}

type MinistepsNotificationsEmail struct {
	APPROVAL_REQUIRED []MinistepsNotificationRecepients `json:"APPROVAL_REQUIRED,omitempty" url:"APPROVAL_REQUIRED,omitempty"`
	CANCELLED         []MinistepsNotificationRecepients `json:"CANCELLED,omitempty" url:"CANCELLED,omitempty"`
	COMPLETED         []MinistepsNotificationRecepients `json:"COMPLETED,omitempty" url:"COMPLETED,omitempty"`
	DRIFT_DETECTED    []MinistepsNotificationRecepients `json:"DRIFT_DETECTED,omitempty" url:"DRIFT_DETECTED,omitempty"`
	ERRORED           []MinistepsNotificationRecepients `json:"ERRORED,omitempty" url:"ERRORED,omitempty"`
}

type MinistepsNotifications struct {
	Email *MinistepsNotificationsEmail `json:"email,omitempty" url:"email,omitempty"`
}

type MinistepsWebhooks struct {
	APPROVAL_REQUIRED []MinistepsWebhooksSchema `json:"APPROVAL_REQUIRED,omitempty" url:"APPROVAL_REQUIRED,omitempty"`
	CANCELLED         []MinistepsWebhooksSchema `json:"CANCELLED,omitempty" url:"CANCELLED,omitempty"`
	COMPLETED         []MinistepsWebhooksSchema `json:"COMPLETED,omitempty" url:"COMPLETED,omitempty"`
	DRIFT_DETECTED    []MinistepsWebhooksSchema `json:"DRIFT_DETECTED,omitempty" url:"DRIFT_DETECTED,omitempty"`
	ERRORED           []MinistepsWebhooksSchema `json:"ERRORED,omitempty" url:"ERRORED,omitempty"`
}

type MinistepsWorkflowChaining struct {
	COMPLETED []MinistepsWfChainingSchema `json:"COMPLETED,omitempty" url:"COMPLETED,omitempty"`
	ERRORED   []MinistepsWfChainingSchema `json:"ERRORED,omitempty" url:"ERRORED,omitempty"`
}

type Ministeps struct {
	Notifications *MinistepsNotifications    `json:"notifications,omitempty" url:"notifications,omitempty"`
	Webhooks      *MinistepsWebhooks         `json:"webhooks,omitempty" url:"webhooks,omitempty"`
	WfChaining    *MinistepsWorkflowChaining `json:"wfChaining,omitempty" url:"wfChaining,omitempty"`
}

type UserSchedulesStateEnum string

const (
	UserSchedulesStateEnumEnabled  UserSchedulesStateEnum = "ENABLED"
	UserSchedulesStateEnumDisabled UserSchedulesStateEnum = "DISABLED"
)

func (t UserSchedulesStateEnum) Ptr() *UserSchedulesStateEnum {
	return &t
}

type UserSchedules struct {
	Cron   string                 `json:"cron,omitempty" url:"cron,omitempty"`
	State  UserSchedulesStateEnum `json:"state,omitempty" url:"state,omitempty"`
	Desc   *string                `json:"desc,omitempty" url:"desc,omitempty"`
	Name   *string                `json:"name,omitempty" url:"name,omitempty"`
	Inputs struct {
		ContextTags          map[string]string        `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
		EnableChaining       *bool                    `json:"EnableChaining,omitempty" url:"EnableChaining,omitempty"`
		EnvironmentVariables []sgsdkgo.EnvVars        `json:"EnvironmentVariables,omitempty" url:"EnvironmentVariables,omitempty"`
		Ministeps            Ministeps                `json:"Ministeps,omitempty" url:"Ministeps,omitempty"`
		ScheduledAt          *string                  `json:"ScheduledAt,omitempty" url:"ScheduledAt,omitempty"`
		TerraformAction      *sgsdkgo.TerraformAction `json:"TerraformAction,omitempty" url:"TerraformAction,omitempty"`
		TerraformConfig      *sgsdkgo.TerraformConfig `json:"TerraformConfig,omitempty" url:"TerraformConfig,omitempty"`
		UserJobCPU           *int                     `json:"UserJobCPU,omitempty" url:"UserJobCPU,omitempty"`
		UserJobMemory        *int                     `json:"UserJobMemory,omitempty" url:"UserJobMemory,omitempty"`
		VCSConfig            *sgsdkgo.VcsConfig       `json:"VCSConfig,omitempty" url:"VCSConfig,omitempty"`
	} `json:"inputs,omitempty" url:"inputs,omitempty"`
}

type CreateWorkflowTemplateRevisionsRequest struct {
	TemplateType              string                                                  `json:"TemplateType,omitempty" url:"TemplateType,omitempty"`
	OwnerOrg                  string                                                  `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	LongDescription           *string                                                 `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	SourceConfigKind          *workflowtemplates.WorkflowTemplateSourceConfigKindEnum `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	Alias                     string                                                  `json:"Alias,omitempty" url:"Alias,omitempty"`
	Approvers                 []string                                                `json:"Approvers,omitempty" url:"Approvers,omitempty"`
	ContextTags               map[string]string                                       `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	DeploymentPlatformConfig  *DeploymentPlatformConfig                               `json:"DeploymentPlatformConfig,omitempty" url:"DeploymentPlatformConfig,omitempty"`
	Deprecation               *Deprecation                                            `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	EnvironmentVariables      []sgsdkgo.EnvVars                                       `json:"EnvironmentVariables,omitempty" url:"EnvironmentVariables,omitempty"`
	InputSchemas              []sgsdkgo.InputSchemas                                  `json:"InputSchemas,omitempty" url:"InputSchemas,omitempty"`
	IsActive                  *sgsdkgo.IsPublicEnum                                   `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic                  *sgsdkgo.IsPublicEnum                                   `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Ministeps                 *Ministeps                                              `json:"Ministeps,omitempty" url:"Ministeps,omitempty"`
	Notes                     string                                                  `json:"Notes,omitempty" url:"Notes,omitempty"`
	NumberOfApprovalsRequired *int                                                    `json:"NumberOfApprovalsRequired,omitempty" url:"NumberOfApprovalsRequired,omitempty"`
	RunnerConstraints         *sgsdkgo.RunnerConstraints                              `json:"RunnerConstraints,omitempty" url:"RunnerConstraints,omitempty"`
	RuntimeSource             *workflowtemplates.RuntimeSource                        `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	Tags                      []string                                                `json:"Tags,omitempty" url:"Tags,omitempty"`
	TerraformConfig           *sgsdkgo.TerraformConfig                                `json:"TerraformConfig,omitempty" url:"TerraformConfig,omitempty"`
	UserJobCPU                *int                                                    `json:"UserJobCPU,omitempty" url:"UserJobCPU,omitempty"`
	UserJobMemory             *int                                                    `json:"UserJobMemory,omitempty" url:"UserJobMemory,omitempty"`
	UserSchedules             []UserSchedules                                         `json:"UserSchedules,omitempty" url:"UserSchedules,omitempty"`
	WfStepsConfig             []sgsdkgo.WfStepsConfig                                 `json:"WfStepsConfig,omitempty" url:"WfStepsConfig,omitempty"`
}

type CreateWorkflowTemplateRevisionResponse struct {
	CreateWorkflowTemplateRevisionsRequest
	Id         string `json:"Id,omitempty" url:"Id,omitempty"`
	TempalteId string `json:"TemplateId,omitempty" url:"TemplateId,omitempty"`
}

type CreateWorkflowTemplateRevisionResponseModel struct {
	Msg  string `json:"msg,omitempty" url:"msg,omitempty"`
	Data struct {
		Revision CreateWorkflowTemplateRevisionResponse `json:"revision,omitempty" url:"revision,omitempty"`
	} `json:"data,omitempty" url:"data,omitempty"`
}

type ReadWorkflowTemplateRevisionModel struct {
	CreateWorkflowTemplateRevisionsRequest
	Id         *string `json:"Id,omitempty" url:"Id,omitempty"`
	TemplateId string  `json:"TemplateId,omitempty" url:"TemplateId,omitempty"`
}

type ReadWorkflowTemplateRevisionResponseModel struct {
	Msg ReadWorkflowTemplateRevisionModel `json:"msg,omitempty" url:"msg,omitempty"`
}

type UpdateWorkflowTemplateRevisionRequest struct {
	OwnerOrg                  *core.Optional[string]                                                 `json:"OwnerOrg,omitempty" url:"OwnerOrg,omitempty"`
	SourceConfigKind          *core.Optional[workflowtemplates.WorkflowTemplateSourceConfigKindEnum] `json:"SourceConfigKind,omitempty" url:"SourceConfigKind,omitempty"`
	Alias                     *core.Optional[string]                                                 `json:"Alias,omitempty" url:"Alias,omitempty"`
	Approvers                 *core.Optional[[]string]                                               `json:"Approvers,omitempty" url:"Approvers,omitempty"`
	LongDescription           *core.Optional[string]                                                 `json:"LongDescription,omitempty" url:"LongDescription,omitempty"`
	ContextTags               *core.Optional[map[string]string]                                      `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	DeploymentPlatformConfig  *core.Optional[DeploymentPlatformConfig]                               `json:"DeploymentPlatformConfig,omitempty" url:"DeploymentPlatformConfig,omitempty"`
	Deprecation               *core.Optional[Deprecation]                                            `json:"Deprecation,omitempty" url:"Deprecation,omitempty"`
	EnvironmentVariables      *core.Optional[[]sgsdkgo.EnvVars]                                      `json:"EnvironmentVariables,omitempty" url:"EnvironmentVariables,omitempty"`
	InputSchemas              *core.Optional[[]sgsdkgo.InputSchemas]                                 `json:"InputSchemas,omitempty" url:"InputSchemas,omitempty"`
	IsActive                  *core.Optional[sgsdkgo.IsPublicEnum]                                   `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsPublic                  *core.Optional[sgsdkgo.IsPublicEnum]                                   `json:"IsPublic,omitempty" url:"IsPublic,omitempty"`
	Ministeps                 *core.Optional[Ministeps]                                              `json:"Ministeps,omitempty" url:"Ministeps,omitempty"`
	Notes                     *core.Optional[string]                                                 `json:"Notes,omitempty" url:"Notes,omitempty"`
	NumberOfApprovalsRequired *core.Optional[int]                                                    `json:"NumberOfApprovalsRequired,omitempty" url:"NumberOfApprovalsRequired,omitempty"`
	RunnerConstraints         *core.Optional[sgsdkgo.RunnerConstraints]                              `json:"RunnerConstraints,omitempty" url:"RunnerConstraints,omitempty"`
	RuntimeSource             *core.Optional[workflowtemplates.RuntimeSourceUpdate]                  `json:"RuntimeSource,omitempty" url:"RuntimeSource,omitempty"`
	Tags                      *core.Optional[[]string]                                               `json:"Tags,omitempty" url:"Tags,omitempty"`
	TerraformConfig           *core.Optional[sgsdkgo.TerraformConfig]                                `json:"TerraformConfig,omitempty" url:"TerraformConfig,omitempty"`
	UserJobCPU                *core.Optional[int]                                                    `json:"UserJobCPU,omitempty" url:"UserJobCPU,omitempty"`
	UserJobMemory             *core.Optional[int]                                                    `json:"UserJobMemory,omitempty" url:"UserJobMemory,omitempty"`
	UserSchedules             *core.Optional[[]UserSchedules]                                        `json:"UserSchedules,omitempty" url:"UserSchedules,omitempty"`
	WfStepsConfig             *core.Optional[[]sgsdkgo.WfStepsConfig]                                `json:"WfStepsConfig,omitempty" url:"WfStepsConfig,omitempty"`
}

type UpdateWorkflowTemplateRevisionResponseModel struct {
	Msg  string                                `json:"msg,omitempty" url:"msg,omitempty"`
	Data UpdateWorkflowTemplateRevisionRequest `json:"data,omitempty" url:"data,omitempty"`
}
