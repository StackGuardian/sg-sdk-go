package client

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	testing "testing"

	sggosdk "github.com/StackGuardian/sg-sdk-go"
	client "github.com/StackGuardian/sg-sdk-go/client"
	option "github.com/StackGuardian/sg-sdk-go/option"
	assert "github.com/stretchr/testify/assert"
)

func TestSDK(t *testing.T) {

	API_KEY := "apikey " + os.Getenv("SG_API_TOKEN")
	SG_ORG := "demo-org"
	SG_WF_GROUP := "sg-sdk-go-test"
	SG_WF := "2aumphefkejtj3bv4q3wo"
	SG_WF_RUN := "3yzuf1izgfw3"
	SG_STACK := "Stack-test"
	SG_STACK_WF := "ansible-Rfde"
	SG_STACK_WF_RUN := "haoc1yepi6p5"
	SG_BASE_URL := os.Getenv("SG_BASE_URL")

	// Workflows

	t.Run("Create_and_delete_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		dummyResourceName := "sdk-test-workflow-create-delete"
		createWorkflowRequest := sggosdk.Workflow{
			ResourceName: sggosdk.Optional(dummyResourceName),
			DeploymentPlatformConfig: sggosdk.Optional([]*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "DummyConnectorForGoSDK",
					"integrationId": "/integrations/DummyConnectorForGoSDK"}}}),
			WfType: sggosdk.Optional(*sggosdk.WfTypeEnumCustom.Ptr()),
			EnvironmentVariables: sggosdk.Optional([]*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}}),
			VcsConfig: sggosdk.Optional(sggosdk.VcsConfig{
				IacVcsConfig: &sggosdk.IacvcsConfig{
					IacTemplateId:          sggosdk.String("/demo-org/ansible-dummy:3"),
					UseMarketplaceTemplate: true,
				},
				IacInputData: &sggosdk.IacInputData{
					SchemaType: sggosdk.IacInputDataSchemaTypeEnumFormJsonschema,
					Data: map[string]interface{}{
						"bucket_region": "eu-central-1",
					},
				},
			}),
			UserJobCpu:    sggosdk.Optional(512),
			UserJobMemory: sggosdk.Optional(1024),
			RunnerConstraints: sggosdk.Optional(sggosdk.RunnerConstraints{
				Type: "shared",
			}),
			Description: sggosdk.Optional("Dummy Workflow for GoSDK"),
		}
		createResponse, err := c.Workflows.CreateWorkflow(context.Background(), SG_ORG, SG_WF_GROUP, &createWorkflowRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createResponse.Data.ResourceName)
		assert.Equal(t, dummyResourceName, createResponse.Data.ResourceName)

		deleteResposnse, err := c.Workflows.DeleteWorkflow(context.Background(), SG_ORG, createResponse.Data.ResourceName, SG_WF_GROUP)
		assert.Equal(t, "Workflow "+createResponse.Data.ResourceName+" deleted", deleteResposnse.Msg)
		assert.Empty(t, err)
	})

	t.Run("Update_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		UpdateWorkflowRequest := sggosdk.PatchedWorkflow{
			DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "DummyConnectorForGoSDK",
					"integrationId": "/integrations/DummyConnectorForGoSDK"}}},
			WfType: sggosdk.WfTypeEnumCustom.Ptr(),
			EnvironmentVariables: []*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}},
			VcsConfig: &sggosdk.VcsConfig{
				IacVcsConfig: &sggosdk.IacvcsConfig{
					IacTemplateId:          sggosdk.String("/demo-org/ansible-dummy:3"),
					UseMarketplaceTemplate: true,
				},
				IacInputData: &sggosdk.IacInputData{
					SchemaType: sggosdk.IacInputDataSchemaTypeEnumFormJsonschema,
					Data: map[string]interface{}{
						"bucket_region": "eu-central-1",
					},
				},
			},
			UserJobCpu:    sggosdk.Int(512),
			UserJobMemory: sggosdk.Int(1024),
			RunnerConstraints: &sggosdk.RunnerConstraints{
				Type: "shared",
			},
			Description: sggosdk.String("Dummy Workflow for GoSDK"),
		}
		updateWorkflowResponse, err := c.Workflows.UpdateWorkflow(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, &UpdateWorkflowRequest)
		assert.Empty(t, err)
		assert.Equal(t, "Workflow "+SG_WF+" updated", updateWorkflowResponse.Msg)
	})

	t.Run("get_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.ReadWorkflow(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, SG_WF, response.Msg.ResourceName)
	})

	t.Run("ListAll_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		request := sggosdk.ListAllWorkflowsRequest{}
		response, err := c.Workflows.ListAllWorkflows(context.Background(), SG_ORG, SG_WF_GROUP, &request)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
		assert.Empty(t, err)
	})

	t.Run("List_all_artifacts_(workflow)", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.ListAllWorkflowArtifacts(context.Background(), SG_ORG, "CUSTOM-7OeX", "test-terragrunt")
		assert.Empty(t, err)
		assert.Equal(t, 15055, response.Data.Artifacts["orgs/demo-org/wfgrps/test-terragrunt/wfs/CUSTOM-7OeX/artifacts/tfstate.json"].Size)
	})

	t.Run("workflow_output", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.Outputs(context.Background(), SG_ORG, "CUSTOM-7OeX", "test-terragrunt")
		assert.Empty(t, err)
		assert.Equal(t, "Outputs retrived", response.Msg)
		assert.NotEmpty(t, response.Data.OutputsSignedUrl)
	})

	t.Run("Update_stack_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		UpdateWorkflowRequest := sggosdk.PatchedWorkflow{
			DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "DummyConnectorForGoSDK",
					"integrationId": "/integrations/DummyConnectorForGoSDK"}}},
			WfType: sggosdk.WfTypeEnumCustom.Ptr(),
			EnvironmentVariables: []*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}},
			VcsConfig: &sggosdk.VcsConfig{
				IacVcsConfig: &sggosdk.IacvcsConfig{
					IacTemplateId:          sggosdk.String("/demo-org/ansible-dummy:3"),
					UseMarketplaceTemplate: true,
				},
				IacInputData: &sggosdk.IacInputData{
					SchemaType: sggosdk.IacInputDataSchemaTypeEnumFormJsonschema,
					Data: map[string]interface{}{
						"bucket_region": "eu-central-1",
					},
				},
			},
			UserJobCpu:    sggosdk.Int(512),
			UserJobMemory: sggosdk.Int(1024),
			RunnerConstraints: &sggosdk.RunnerConstraints{
				Type: "shared",
			},
			Description: sggosdk.String("Dummy Workflow for GoSDK"),
		}
		updateWorkflowResponse, err := c.StackWorkflows.UpdateStackWorkflow(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, &UpdateWorkflowRequest)
		assert.Empty(t, err)
		assert.Equal(t, "Workflow "+SG_STACK_WF+" updated", updateWorkflowResponse.Msg)
	})

	t.Run("get_stackworkflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.ReadStackWorkflow(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, SG_STACK_WF, response.Msg.ResourceName)

	})

	t.Run("ListAll_stack_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		request := sggosdk.ListAllStackWorkflowsRequest{}
		response, err := c.StackWorkflows.ListAllStackWorkflows(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP, &request)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("List_all_artifacts_(stack_workflow)", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.ListAllStackWorkflowsArtifacts(context.Background(), SG_ORG, "stack1",
			"refeed2-null-resource-tf-JuNs", "refeed-test-nested-stackrunbug")
		assert.Empty(t, err)
		assert.Equal(t, "Outputs retrieved", response.Msg)
		assert.Equal(t, 817,
			response.Data.Artifacts["orgs/demo-org/wfgrps/refeed-test-nested-stackrunbug/stacks/stack1/wfs/refeed2-null-resource-tf-JuNs/artifacts/tfstate.json"].Size)
	})

	t.Run("stack_workflow_output", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.StackWorkflowOutputs(context.Background(), SG_ORG, "stack1",
			"refeed2-null-resource-tf-JuNs", "refeed-test-nested-stackrunbug")
		assert.Empty(t, err)
		assert.Equal(t, "Outputs retrived", response.Msg)
		assert.NotEmpty(t, response.Data.OutputsSignedUrl)
	})

	// Workflow Runs
	t.Run("ListAll_workflow_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		request := sggosdk.ListAllWorkflowRunsRequest{}
		response, err := c.WorkflowRuns.ListAllWorkflowRuns(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, &request)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("Get_workflow_runs_stack", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflowRuns.ReadStackWorkflowRun(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, SG_STACK_WF_RUN)
		assert.Equal(t, "QUEUED", response.Msg.Statuses["pre_0_step"][0].Name)
		assert.Empty(t, err)

	})

	t.Run("Get_workflow_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.WorkflowRuns.ReadWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN)
		assert.Equal(t, "QUEUED", response.Msg.Statuses["pre_0_step"][0].Name)
		assert.Empty(t, err)

	})

	t.Run("Create_workflow_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		createWorkflowRunRequest := sggosdk.WorkflowRun{
			DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "testAWSConnector",
					"integrationId": "/integrations/testAWSConnector"}}},
			WfType: sggosdk.WfTypeEnumTerraform.Ptr(),
			EnvironmentVariables: []*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}},
			VcsConfig: &sggosdk.VcsConfig{
				IacVcsConfig: &sggosdk.IacvcsConfig{
					IacTemplateId:          sggosdk.String("/stackguardian/aws-s3-demo-website:16"),
					UseMarketplaceTemplate: true,
				},
				IacInputData: &sggosdk.IacInputData{
					SchemaType: sggosdk.IacInputDataSchemaTypeEnumFormJsonschema,
					Data: map[string]interface{}{
						"bucket_region": "eu-central-1",
					},
				},
			},
			UserJobCpu:    sggosdk.Int(512),
			UserJobMemory: sggosdk.Int(1024),
			RunnerConstraints: &sggosdk.RunnerConstraints{
				Type: "shared",
			},
		}

		response, err := c.WorkflowRuns.CreateWorkflowRun(context.Background(),
			SG_ORG, SG_WF, SG_WF_GROUP, &createWorkflowRunRequest)
		assert.Empty(t, err)
		newWfRunName := response.Data.ResourceName
		assert.NotEmpty(t, newWfRunName)

	})

	t.Run("Approve_workflow_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		approveWfRunRequest := sggosdk.WorkflowRunApproval{
			Approve:                   sggosdk.Bool(true),
			Message:                   sggosdk.String("Approved"),
			ReasonForApprovalRequired: sggosdk.String("Approval reason"),
		}
		_, err := c.WorkflowRuns.ApproveWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN,
			&approveWfRunRequest)
		// We expect an error since the workflow run doesnt have any approvals pending
		assert.Contains(t, err.Error(), "No approval pending")

	})

	t.Run("Approve_workflow_runs_(stack)", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		approveWfRunRequest := sggosdk.WorkflowRunApproval{
			Approve:                   sggosdk.Bool(true),
			Message:                   sggosdk.String("Approved"),
			ReasonForApprovalRequired: sggosdk.String("Approval reason"),
		}
		err := c.StackWorkflowRuns.ApproveStackWorkflowRun(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, SG_STACK_WF_RUN,
			&approveWfRunRequest)
		// We expect an error since the workflow run doesnt have any approvals pending
		assert.Contains(t, err.Error(), "No approval pending")

	})

	t.Run("Get_workflow_runs_logs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		logs, err := c.WorkflowRuns.ReadWorkflowRunLogs(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(logs.Msg), 1)
	})

	t.Run("Get_workflow_runs_logs_(stack)", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		logs, err := c.StackWorkflowRuns.ReadStackWorkflowRunLogs(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, SG_STACK_WF_RUN)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(logs.Msg), 1)
	})

	t.Run("Cancel_workflow_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		_, err := c.WorkflowRuns.CancelWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN)
		// We expect an error since the workflow run is already failed
		if err != nil {
			assert.Contains(t, err.Error(), "Error cancelling Workflow Run "+SG_WF_RUN)
		}
	})

	t.Run("Update_workflow_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		updateWfRunRequest := sggosdk.PatchedWorkflowRun{
			DeploymentPlatformConfig: sggosdk.Optional([]*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "testAWSConnector",
					"integrationId": "/integrations/testAWSConnector"}}}),
			WfType: sggosdk.Optional(*sggosdk.WfTypeEnumTerraform.Ptr()),
			EnvironmentVariables: sggosdk.Optional([]*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("UpdatedValue")}}}),
			VcsConfig: sggosdk.Optional(sggosdk.VcsConfig{
				IacVcsConfig: &sggosdk.IacvcsConfig{
					IacTemplateId:          sggosdk.String("/stackguardian/aws-s3-demo-website:16"),
					UseMarketplaceTemplate: true,
				},
				IacInputData: &sggosdk.IacInputData{
					SchemaType: sggosdk.IacInputDataSchemaTypeEnumFormJsonschema,
					Data: map[string]interface{}{
						"bucket_region": "eu-central-1",
					},
				},
			}),
			UserJobCpu:    sggosdk.Optional(512),
			UserJobMemory: sggosdk.Optional(1024),
			RunnerConstraints: sggosdk.Optional(sggosdk.RunnerConstraints{
				Type: "shared",
			}),
		}
		updateWfRunResponse, err := c.WorkflowRuns.UpdateWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN, &updateWfRunRequest)
		assert.Empty(t, err)
		assert.Equal(t, "Workflow Run "+SG_WF_RUN+" updated", updateWfRunResponse.Msg)
	})

	// Stacks
	t.Run("Create_and_delete_stack", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		dummyResourceName := "sdk-test-stack-create-delete"
		createStackRequest := sggosdk.Stack{
			ResourceName: sggosdk.Optional(dummyResourceName),
			RunOnCreate:  sggosdk.Bool(false),
			DeploymentPlatformConfig: sggosdk.Optional([]*sggosdk.DeploymentPlatformConfig{
				{
					Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
					Config: map[string]interface{}{
						"profileName":   "DummyConnectorForGoSDK",
						"integrationId": "/integrations/DummyConnectorForGoSDK",
					},
				},
			}),
			EnvironmentVariables: sggosdk.Optional([]*sggosdk.EnvVars{
				{
					Kind: sggosdk.EnvVarsKindEnumPlainText,
					Config: &sggosdk.EnvVarConfig{
						VarName: "test", TextValue: sggosdk.String("testValue"),
					},
				},
			}),

			Description: sggosdk.Optional("Dummy Stack for GoSDK"),
			WorkflowsConfig: sggosdk.Optional(sggosdk.WorkflowsConfig{
				Workflows: []*sggosdk.WorkflowsConfigWorkflow{
					{
						NumberOfApprovalsRequired: sggosdk.Int(0),
						Description:               sggosdk.String("Dummy Workflow for GoSDK"),
						WfType:                    sggosdk.WfTypeEnumCustom.Ptr(),
						Id:                        sggosdk.String("cc0061e9-a75c-421b-a75b-ef918e9f4b28"),
						DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
							Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
							Config: map[string]interface{}{
								"profileName":   "DummyConnectorForGoSDK",
								"integrationId": "/integrations/DummyConnectorForGoSDK"}}},
					},
				}}),
		}
		createStackResponse, err := c.Stacks.CreateStack(context.Background(), SG_ORG, SG_WF_GROUP, &createStackRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createStackResponse.Data.Stack.ResourceName)
		assert.Equal(t, dummyResourceName, *createStackResponse.Data.Stack.ResourceName)
		assert.Equal(t, "Stack "+*createStackResponse.Data.Stack.ResourceName+" created", createStackResponse.Msg)

		err = c.StackWorkflows.DeleteStackWorkflow(context.Background(), SG_ORG, *createStackResponse.Data.Stack.ResourceName,
			strings.Split(createStackResponse.Data.Workflows[0].ResourceId, "/")[2], SG_WF_GROUP)
		assert.Empty(t, err)
		deleteResponse, err := c.Stacks.DeleteStack(context.Background(), SG_ORG, *createStackResponse.Data.Stack.ResourceName, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, "Stack "+*createStackResponse.Data.Stack.ResourceName+" deleted", *deleteResponse.Msg)
		assert.Empty(t, err)
	})

	t.Run("Read_stack", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Stacks.ReadStack(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, SG_STACK, *response.Msg.ResourceName)
	})

	t.Run("Run_stack", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		runStackRequest := sggosdk.StackAction{
			ActionType: string(sggosdk.ActionEnumApply),
		}
		response, err := c.StackRuns.CreateStackRun(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP, &runStackRequest)
		assert.Empty(t, err)
		assert.Equal(t, "Stack run scheduled", response.Msg)
	})

	t.Run("ListAll_stacks", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		request := sggosdk.ListAllStacksRequest{}
		response, err := c.Stacks.ListAllStacks(context.Background(), SG_ORG, SG_WF_GROUP, &request)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("Get_stack_outputs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Stacks.ReadStackOutputs(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("update_stack", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		createStackRequest := sggosdk.PatchedStack{
			DeploymentPlatformConfig: sggosdk.Optional([]*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "DummyConnectorForGoSDK",
					"integrationId": "/integrations/DummyConnectorForGoSDK"}}}),
			EnvironmentVariables: sggosdk.Optional([]*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}}),

			Description: sggosdk.Optional("Dummy Stack for GoSDK"),
			TemplatesConfig: sggosdk.Optional(sggosdk.TemplatesConfig{
				TemplateGroupId: sggosdk.String("/demo-org/ansible:4"),
				Templates: []*sggosdk.TemplateWorkflow{{
					NumberOfApprovalsRequired: sggosdk.Int(0),
					Description:               sggosdk.String("Dummy Workflow for GoSDK"),
					WfType:                    sggosdk.WfTypeEnumCustom.Ptr(),
					Id:                        sggosdk.String("cc0061e9-a75c-421b-a75b-ef918e9f4b28"),
				}},
			}),
		}
		updateStackResponse, err := c.Stacks.UpdateStack(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP, &createStackRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateStackResponse.Msg)
		assert.Equal(t, "Stack "+SG_STACK+" updated", updateStackResponse.Msg)
	})

	t.Run("list_all_stack_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		request := sggosdk.ListAllStackRunsRequest{}
		listResponse, err := c.StackRuns.ListAllStackRuns(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP, &request)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(listResponse.Msg), 1)
	})

	t.Run("get_stack_runs", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		getStackRunResponse, err := c.StackRuns.ReadStackRun(context.Background(), SG_ORG, SG_STACK, "5srghvu1y7nn", SG_WF_GROUP)
		assert.Empty(t, err)
		assert.NotEmpty(t, getStackRunResponse.Msg.ResourceName)
		assert.Equal(t, "/stackruns/5srghvu1y7nn", getStackRunResponse.Msg.ResourceName)
	})

	// Connectors and Cloud Connector Groups
	t.Run("create_and_delete_connector", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		connectorName := "test-connector-go-sdk-test"
		var settingsConfigArray []*sggosdk.SettingsConfig
		settingsConfigArray = append(settingsConfigArray, &sggosdk.SettingsConfig{
			ArmTenantId:       sggosdk.String("1"),
			ArmSubscriptionId: sggosdk.String("1"),
			ArmClientId:       sggosdk.String("1"),
			ArmClientSecret:   sggosdk.String("1"),
		})
		description := "test-connector description"
		createConnectorRequest := sggosdk.Integration{
			ResourceName: &connectorName,
			Description:  &description,
			Settings: &sggosdk.IntegrationsSettings{
				Kind:   sggosdk.IntegrationsSettingsKindEnumAzureStatic,
				Config: settingsConfigArray,
			},
		}
		createConnectorResponse, err := c.Connectors.CreateConnector(context.Background(), SG_ORG, &createConnectorRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createConnectorResponse.Msg)
		assert.Equal(t, "Connector "+connectorName+" created", *createConnectorResponse.Msg)
		//Check that the response contains the resource name
		assert.NotEmpty(t, createConnectorResponse.Data.ResourceName)
		assert.Equal(t, connectorName, createConnectorResponse.Data.ResourceName)

		deleteConnectorResponse, err := c.Connectors.DeleteConnector(context.Background(), connectorName, SG_ORG)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteConnectorResponse.Msg)
		assert.Equal(t, "Connector "+connectorName+" deleted", deleteConnectorResponse.Msg)
	})

	t.Run("read_connector", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		connectorName := "DummyConnectorForGoSDK"
		readConnectorResponse, err := c.Connectors.ReadConnector(context.Background(), connectorName, SG_ORG)
		assert.Empty(t, err)
		assert.NotEmpty(t, readConnectorResponse.Msg)
		assert.Equal(t, connectorName, readConnectorResponse.Msg.ResourceName)
	})

	t.Run("update_connector", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		connectorName := "DummyConnectorForGoSDK"
		var settingsConfigArray []*sggosdk.SettingsConfig
		settingsConfigArray = append(settingsConfigArray, &sggosdk.SettingsConfig{
			ArmTenantId:       sggosdk.String("1"),
			ArmSubscriptionId: sggosdk.String("1"),
			ArmClientId:       sggosdk.String("1"),
			ArmClientSecret:   sggosdk.String("1"),
		})
		updatedDescription := "updated description"
		updateConnectorRequest := sggosdk.PatchedIntegration{
			ResourceName: sggosdk.Optional(&connectorName),
			Description:  sggosdk.Optional(&updatedDescription),
			Settings: sggosdk.Optional(&sggosdk.IntegrationsSettings{
				Kind:   sggosdk.IntegrationsSettingsKindEnumAzureStatic,
				Config: settingsConfigArray,
			}),
		}
		updateConnectorResponse, err := c.Connectors.UpdateConnector(context.Background(), connectorName, SG_ORG, &updateConnectorRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateConnectorResponse.Msg)
		assert.Equal(t, "Integration "+connectorName+" updated", *updateConnectorResponse.Msg)
	})

	t.Run("listall_connector", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		request := sggosdk.ListAllConnectorsRequest{}
		listAllConnectorResponse, err := c.Connectors.ListAllConnectors(context.Background(), SG_ORG, &request)
		assert.Empty(t, err)
		assert.NotEmpty(t, listAllConnectorResponse.Msg)
		assert.GreaterOrEqual(t, len(listAllConnectorResponse.Msg), 1)
	})

	//Workflow Groups
	t.Run("create_and_delete_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "test-wfg-go-sdk-test"
		createWorkflowGroupRequest := sggosdk.WorkflowGroup{
			ResourceName: sggosdk.String(workflowGroupName),
			Description:  sggosdk.String("test-workflowGroup description"),
		}
		createWorkflowGroupResponse, err := c.WorkflowGroups.CreateWorkflowGroup(context.Background(), SG_ORG, &createWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createWorkflowGroupResponse)
		assert.Equal(t, "Workflow Group "+*createWorkflowGroupRequest.ResourceName+" created", *createWorkflowGroupResponse.Msg)

		deleteWorkflowGroupResponse, err := c.WorkflowGroups.DeleteWorkflowGroup(context.Background(), SG_ORG, workflowGroupName)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+workflowGroupName+" deleted", *deleteWorkflowGroupResponse.Msg)
	})

	t.Run("update_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "go-sdk-test-wfg"
		updateWorkflowGroupRequest := sggosdk.PatchedWorkflowGroup{
			ResourceName: sggosdk.Optional(workflowGroupName),
			Description:  sggosdk.Optional("updated description"),
		}
		updateWorkflowGroupResponse, err := c.WorkflowGroups.UpdateWorkflowGroup(context.Background(), SG_ORG, workflowGroupName, &updateWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+workflowGroupName+" updated", *updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "updated description", *updateWorkflowGroupResponse.Data.Description)
	})

	t.Run("update_nested_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := SG_WF_GROUP + "/1bger5ydab697a4jxe2gu"
		updateWorkflowGroupRequest := sggosdk.PatchedWorkflowGroup{
			// ResourceName: sggosdk.String(workflowGroupName),
			Description: sggosdk.Optional("updated description"),
		}
		updateWorkflowGroupResponse, err := c.WorkflowGroups.UpdateWorkflowGroup(context.Background(), SG_ORG, workflowGroupName, &updateWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group 1bger5ydab697a4jxe2gu updated", *updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "updated description", *updateWorkflowGroupResponse.Data.Description)
	})

	t.Run("read_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := SG_WF_GROUP
		readWorkflowGroupResponse, err := c.WorkflowGroups.ReadWorkflowGroup(context.Background(), SG_ORG, workflowGroupName)
		assert.Empty(t, err)
		assert.NotEmpty(t, readWorkflowGroupResponse.Msg)
		assert.Equal(t, workflowGroupName, *readWorkflowGroupResponse.Msg.ResourceName)
	})

	t.Run("read_nested_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := SG_WF_GROUP + "/1bger5ydab697a4jxe2gu"
		readWorkflowGroupResponse, err := c.WorkflowGroups.ReadWorkflowGroup(context.Background(), SG_ORG, workflowGroupName)
		assert.Empty(t, err)
		assert.NotEmpty(t, readWorkflowGroupResponse.Msg)
		assert.Equal(t, "1bger5ydab697a4jxe2gu", *readWorkflowGroupResponse.Msg.ResourceName)
	})

	t.Run("listall_workflow_groups", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		listAllWorkflowGroupResponse, err := c.WorkflowGroups.ListAllWorkflowGroups(context.Background(), SG_ORG, &sggosdk.ListAllWorkflowGroupsRequest{})
		assert.Empty(t, err)
		assert.NotEmpty(t, listAllWorkflowGroupResponse.Msg)
		assert.GreaterOrEqual(t, len(listAllWorkflowGroupResponse.Msg), 1)
	})

	//Nested Workflow Groups
	t.Run("listall_nested_workflow_groups", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		listAllWorkflowGroupResponse, err := c.WorkflowGroups.ListAllChildWorkflowGroups(context.Background(), SG_ORG, SG_WF_GROUP, &sggosdk.ListAllChildWorkflowGroupsRequest{})
		assert.Empty(t, err)
		assert.NotEmpty(t, listAllWorkflowGroupResponse.Msg)
		assert.GreaterOrEqual(t, len(listAllWorkflowGroupResponse.Msg), 1)
	})

	t.Run("create_and_delete_nested_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		parentWorkflowGroupName := SG_WF_GROUP
		dummyResourceName := "sg-sdk-child-workflow-group-create-delete"
		createWorkflowGroupRequest := sggosdk.WorkflowGroup{
			ResourceName: sggosdk.String(dummyResourceName),
			Description:  sggosdk.String("child workflowGroup description"),
		}
		createChildWorkflowGroupResponse, err := c.WorkflowGroups.CreateChildWorkflowGroup(
			context.Background(),
			SG_ORG,
			parentWorkflowGroupName,
			&createWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createChildWorkflowGroupResponse)
		assert.NotEmpty(t, createChildWorkflowGroupResponse.Msg)
		assert.Equal(t, dummyResourceName, *createChildWorkflowGroupResponse.Data.ResourceName)
		assert.Contains(t, *createChildWorkflowGroupResponse.Msg, "created")

		deleteWorkflowGroupResponse, err := c.WorkflowGroups.DeleteWorkflowGroup(context.Background(), SG_ORG,
			parentWorkflowGroupName+"/"+*createChildWorkflowGroupResponse.Data.ResourceName)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+*createChildWorkflowGroupResponse.Data.ResourceName+" deleted",
			*deleteWorkflowGroupResponse.Msg)
	})

	t.Run("create_and_delete_deep_nested_workflow_group", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		parentWorkflowGroupName := SG_WF_GROUP + "/1bger5ydab697a4jxe2gu"
		dummyResourceName := "sg-sdk-child-workflow-group-create-delete"
		createWorkflowGroupRequest := sggosdk.WorkflowGroup{
			ResourceName: sggosdk.String(dummyResourceName),
			Description:  sggosdk.String("child workflowGroup description"),
		}
		createChildWorkflowGroupResponse, err := c.WorkflowGroups.CreateChildWorkflowGroup(
			context.Background(),
			SG_ORG,
			parentWorkflowGroupName,
			&createWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createChildWorkflowGroupResponse)
		assert.NotEmpty(t, createChildWorkflowGroupResponse.Msg)
		assert.Equal(t, dummyResourceName, *createChildWorkflowGroupResponse.Data.ResourceName)
		assert.Contains(t, *createChildWorkflowGroupResponse.Msg, "created")

		deleteWorkflowGroupResponse, err := c.WorkflowGroups.DeleteWorkflowGroup(context.Background(), SG_ORG,
			parentWorkflowGroupName+"/"+*createChildWorkflowGroupResponse.Data.ResourceName)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+*createChildWorkflowGroupResponse.Data.ResourceName+" deleted",
			*deleteWorkflowGroupResponse.Msg)
	})

	//Roles
	t.Run("create_and_delete_roles", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		roleName := "Go-SDK-Test-Role"
		allowedPermissions := &sggosdk.AllowedPermissions{
			Name: "GET/api/v1/orgs/demo-org/policies/<policy>/",
			Paths: map[string][]string{
				"<wfGrps>": {"test"},
			},
		}
		createRoleRequest := sggosdk.Role{
			ResourceName: roleName,
			Description:  sggosdk.Optional("role description"),
			AllowedPermissions: sggosdk.Optional[map[string]*sggosdk.AllowedPermissions](map[string]*sggosdk.AllowedPermissions{
				"GET/api/v1/orgs/demo-org/policies/<policy>/": allowedPermissions,
			}),
		}
		createRoleResponse, err := c.AccessManagement.CreateRole(context.Background(), SG_ORG, &createRoleRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createRoleResponse)
		assert.Equal(t, "Role "+createRoleResponse.Data.ResourceName+" created", *createRoleResponse.Msg)

		err = c.AccessManagement.DeleteRole(context.Background(), SG_ORG, roleName)
		assert.Empty(t, err)
	})

	t.Run("update_role", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		roleName := "SDK-Test-Role"
		allowedPermissions := &sggosdk.AllowedPermissions{
			Name: "GET/api/v1/orgs/demo-org/policies/<policy>/",
			Paths: map[string][]string{
				"<wfGrps>": {"test"},
			},
		}
		updateRoleRequest := sggosdk.PatchedRole{
			ResourceName: sggosdk.Optional(roleName),
			Description:  sggosdk.Optional("updated description"),
			AllowedPermissions: sggosdk.Optional[map[string]*sggosdk.AllowedPermissions](map[string]*sggosdk.AllowedPermissions{
				"GET/api/v1/orgs/demo-org/policies/<policy>/": allowedPermissions,
			}),
		}
		updateRoleResponse, err := c.AccessManagement.UpdateRole(context.Background(), SG_ORG, roleName, &updateRoleRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateRoleResponse.Msg)
		assert.Equal(t, "Role /roles/"+roleName+" updated", *updateRoleResponse.Msg)
		assert.Equal(t, "updated description", *updateRoleResponse.Data.Description)
	})

	t.Run("read_role", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		roleName := "SDK-Test-Role"
		readRoleResponse, err := c.AccessManagement.ReadRole(context.Background(), SG_ORG, roleName)
		assert.Empty(t, err)
		assert.NotEmpty(t, readRoleResponse.Msg)
		assert.Equal(t, roleName, readRoleResponse.Msg.ResourceName)
	})

	t.Run("listall_role", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		err := c.AccessManagement.ListAllRoles(context.Background(), SG_ORG, &sggosdk.ListAllRolesRequest{})
		assert.Empty(t, err)
	})

	// Users/Role assignment
	t.Run("add_and_remove_users", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		userName := "deleteme@stackguardian.io"
		createUserRequest := sggosdk.AddUserToOrganization{
			Role:         sggosdk.String("SDK-Test-Role"),
			UserId:       userName,
			ResendInvite: sggosdk.Bool(false),
		}
		createUserResponse, err := c.AccessManagement.CreateUser(context.Background(), SG_ORG, &createUserRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createUserResponse)
		assert.Equal(t, userName+" invited.", *createUserResponse.Msg)

		userId := "eu-central-1_C6bwuggLI/local/" + userName
		removeUserRequest := sggosdk.GetorRemoveUserFromOrganization{
			UserId: &userId,
		}
		deleteUserResponse, err := c.AccessManagement.DeleteUser(context.Background(), SG_ORG, &removeUserRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteUserResponse.Msg)
		assert.Equal(t, userName+" removed from /orgs/demo-org", *deleteUserResponse.Msg)
	})

	t.Run("read_users", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		userName := "test@dummy.com"

		removeUserRequest := sggosdk.GetorRemoveUserFromOrganization{
			UserId: &userName,
		}
		getUserResponse, err := c.AccessManagement.ReadUser(context.Background(), SG_ORG, &removeUserRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, getUserResponse.Msg)
	})

	t.Run("update_users", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		userName := "test@dummy.com"

		updateUserRequest := sggosdk.AddUserToOrganization{
			Role:   sggosdk.String("SDK-Test-Role"),
			UserId: userName,
		}
		updateUserResponse, err := c.AccessManagement.UpdateUser(context.Background(), SG_ORG, &updateUserRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateUserResponse.Msg)
	})

	t.Run("create_and_delete_policy", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		createPolicyRequest := sggosdk.PolymorphicPolicy{
			PolicyType: "GENERAL",
			General: &sggosdk.PolicyGeneral{
				ResourceName:              sggosdk.String("GoSDKTestPolicyCreate"),
				Description:               sggosdk.String("SDK Test Policy Description"),
				NumberOfApprovalsRequired: sggosdk.Int(1),
			},
		}
		createPolicyResponse, err := c.Policies.CreatePolicy(context.Background(), SG_ORG, &createPolicyRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createPolicyResponse.Msg)
		assert.Equal(t, "Policy "+*createPolicyRequest.General.ResourceName+" created", *createPolicyResponse.Msg)

		//TODO: Add response
		err = c.Policies.DeletePolicy(context.Background(), SG_ORG, *createPolicyRequest.General.ResourceName)
		assert.Empty(t, err)
	})

	t.Run("read_policies", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		policyName := "SDKTestPolicyForUpdate"
		readPolicyResponse, err := c.Policies.ReadPolicy(context.Background(), SG_ORG, policyName)
		assert.Empty(t, err)
		assert.NotEmpty(t, readPolicyResponse.Msg)
		// assert.Equal(t, policyName, *readPolicyResponse.Msg.ResourceName)
	})

	t.Run("listAll_policies", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		//TODO: Add response
		request := sggosdk.ListAllPoliciesRequest{}
		err := c.Policies.ListAllPolicies(context.Background(), SG_ORG, &request)
		assert.Empty(t, err)
	})

	t.Run("update_policy_insights", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		policyName := "Exclusion-policy-n9kv2"
		updatePolicyRequest := sggosdk.PatchedPolymorphicPolicy{
			PolicyType: "FILTER.INSIGHT",
			FilterInsight: &sggosdk.PatchedPolicyFilterInsight{
				ResourceName:   sggosdk.Optional(policyName),
				Description:    sggosdk.Null[string](),
				PoliciesConfig: sggosdk.Optional([]*sggosdk.PoliciesFilterInsightConfig{}),
			},
		}
		createPolicyResponse, err := c.Policies.UpdatePolicy(context.Background(), SG_ORG, policyName, &updatePolicyRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createPolicyResponse.Msg)
		assert.Equal(t, "Policy "+policyName+" updated", *createPolicyResponse.Msg)
	})
	t.Run("update_policy", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		policyName := "SDKTestPolicyForUpdate"
		updatePolicyRequest := sggosdk.PatchedPolymorphicPolicy{
			PolicyType: "GENERAL",
			General: &sggosdk.PatchedPolicyGeneral{
				ResourceName:              sggosdk.Optional("test-to-delete"),
				Description:               sggosdk.Optional("SDK Test Policy Description"),
				NumberOfApprovalsRequired: sggosdk.Optional(1),
				Tags:                      sggosdk.Optional([]string{}),
				PoliciesConfig:            sggosdk.Optional([]*sggosdk.PoliciesConfig{}),
			},
		}
		createPolicyResponse, err := c.Policies.UpdatePolicy(context.Background(), SG_ORG, policyName, &updatePolicyRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createPolicyResponse.Msg)
		assert.Equal(t, "Policy "+policyName+" updated", *createPolicyResponse.Msg)
	})

	t.Run("patch_policy_with_nil", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		policyName := "SDKTestPolicyForPatch"
		createPolicyRequest := sggosdk.PolymorphicPolicy{
			PolicyType: "GENERAL",
			General: &sggosdk.PolicyGeneral{
				ResourceName:              sggosdk.String(policyName),
				Description:               sggosdk.String("SDK Test Policy Description"),
				NumberOfApprovalsRequired: sggosdk.Int(1),
				PoliciesConfig: []*sggosdk.PoliciesConfig{
					{
						Name:   "PolicyConfig-1",
						OnFail: sggosdk.OnFailEnumFail,
						OnPass: sggosdk.OnPassEnumPass,
						PolicyVcsConfig: &sggosdk.PolicyVcsConfig{
							PolicyTemplateId:       sggosdk.String("/demo-org/sg_workflow_check_userschedules-Copy:1"),
							UseMarketplaceTemplate: true,
						},
					},
				},
			},
		}
		createPolicyResponse, err := c.Policies.CreatePolicy(context.Background(), SG_ORG, &createPolicyRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createPolicyResponse.Msg)
		assert.Equal(t, "Policy "+*createPolicyRequest.General.ResourceName+" created", *createPolicyResponse.Msg)

		// now lets update the policy to unset the policiesConfig and description
		updatePolicyRequest := sggosdk.PatchedPolymorphicPolicy{
			PolicyType: "GENERAL",
			General: &sggosdk.PatchedPolicyGeneral{
				ResourceName:              sggosdk.Optional(policyName),
				Description:               sggosdk.Null[string](),
				NumberOfApprovalsRequired: sggosdk.Optional(1),
				PoliciesConfig:            sggosdk.Null[[]*sggosdk.PoliciesConfig](),
			},
		}
		updatePolicyResponse, err := c.Policies.UpdatePolicy(context.Background(), SG_ORG, policyName, &updatePolicyRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updatePolicyResponse.Msg)
		// make sure the description and policiesConfig are unset
		assert.Equal(t, "", *updatePolicyResponse.Data.General.Description)
		assert.Equal(t, 0, len(updatePolicyResponse.Data.General.PoliciesConfig))
		assert.Equal(t, "Policy "+policyName+" updated", *updatePolicyResponse.Msg)

		// now lets delete the policy
		err = c.Policies.DeletePolicy(context.Background(), SG_ORG, *createPolicyRequest.General.ResourceName)
		assert.Empty(t, err)
	})

	t.Run("umarshal_create_workflow", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		var createWorkflowRequest *sggosdk.Workflow
		payload, err := os.ReadFile("create_workflow.json")
		assert.Empty(t, err)
		err = json.Unmarshal(
			payload,
			&createWorkflowRequest)
		assert.Empty(t, err)
		createResponse, err := c.Workflows.CreateWorkflow(context.Background(), SG_ORG, SG_WF_GROUP, createWorkflowRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createResponse.Data.ResourceName)

		deleteResposnse, err := c.Workflows.DeleteWorkflow(context.Background(), SG_ORG, createResponse.Data.ResourceName, SG_WF_GROUP)
		assert.Equal(t, "Workflow "+createResponse.Data.ResourceName+" deleted", deleteResposnse.Msg)
		assert.Empty(t, err)
	})

	t.Run("umarshal_create_stack", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		var createStackRequest *sggosdk.Stack
		payload, err := os.ReadFile("create_stack.json")
		assert.Empty(t, err)
		err = json.Unmarshal(
			payload,
			&createStackRequest)
		assert.Empty(t, err)

		createStackRequest.RunOnCreate = sggosdk.Bool(false)

		createStackResponse, err := c.Stacks.CreateStack(context.Background(), SG_ORG, SG_WF_GROUP, createStackRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createStackResponse.Data.Stack.ResourceName)
		assert.Equal(t, "Stack "+*createStackResponse.Data.Stack.ResourceName+" created", createStackResponse.Msg)

		err = c.StackWorkflows.DeleteStackWorkflow(context.Background(), SG_ORG, *createStackResponse.Data.Stack.ResourceName,
			strings.Split(createStackResponse.Data.Workflows[0].ResourceId, "/")[2], SG_WF_GROUP)
		assert.Empty(t, err)
		deleteResponse, err := c.Stacks.DeleteStack(context.Background(), SG_ORG, *createStackResponse.Data.Stack.ResourceName, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, "Stack "+*createStackResponse.Data.Stack.ResourceName+" deleted", *deleteResponse.Msg)
		assert.Empty(t, err)
	})

	t.Run("create_and_delete_runnergroups", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		runnerGroupName := "GoSDKTestRunnerGroup"
		createRunnerGroupRequest := sggosdk.RunnerGroup{
			ResourceName: sggosdk.Optional(runnerGroupName),
			Description:  sggosdk.Optional("SDK Test Runner Group Description"),
			StorageBackendConfig: &sggosdk.StorageBackendConfig{
				Type:         sggosdk.StorageBackendConfigTypeEnumAwsS3,
				AwsRegion:    sggosdk.String("eu-central-1"),
				S3BucketName: sggosdk.String("test-bucket"),
				Auth: &sggosdk.StorageBackendConfigAuth{
					IntegrationId: "/integrations/test_test",
				},
			},
		}
		createRunnerGroupResponse, err := c.RunnerGroups.CreateNewRunnerGroup(context.Background(), SG_ORG, &createRunnerGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createRunnerGroupResponse.Msg)
		assert.Equal(t, "Runner Group "+runnerGroupName+" created", *createRunnerGroupResponse.Msg)

		// Now delete it
		deleteRunnerGroupResponse, err := c.RunnerGroups.DeleteRunnerGroup(context.Background(), SG_ORG, runnerGroupName)
		assert.Empty(t, err)
		assert.Equal(t, "Runner group deleted succesfully", *deleteRunnerGroupResponse.Msg)
	})

	t.Run("read_runnergroup", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		policyName := "shared-dind-2"
		readRunnerGroupRequest := sggosdk.ReadRunnerGroupRequest{}
		readPolicyResponse, err := c.RunnerGroups.ReadRunnerGroup(context.Background(), SG_ORG, policyName, &readRunnerGroupRequest)
		assert.Empty(t, err)
		assert.Equal(t, policyName, *readPolicyResponse.Msg.ResourceName)
	})

	t.Run("update_runnergroups", func(t *testing.T) {
		c := client.NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		runnerGroupName := "SDK-test"
		updateRunnerGroupRequest := sggosdk.PatchedRunnerGroup{
			Description: sggosdk.Optional("SDK Test Runner Group Description"),
		}
		updateRunnerGroupResponse, err := c.RunnerGroups.UpdateRunnerGroup(context.Background(), SG_ORG, runnerGroupName, &updateRunnerGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateRunnerGroupResponse.Msg)
		assert.Equal(t, "Runner Group /runnergroups/"+runnerGroupName+" updated", *updateRunnerGroupResponse.Msg)
	})
}
