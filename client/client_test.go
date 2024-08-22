package client

import (
	"context"
	http "net/http"
	"os"
	testing "testing"
	time "time"

	sggosdk "github.com/StackGuardian/sg-sdk-go"
	option "github.com/StackGuardian/sg-sdk-go/option"
	assert "github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	API_KEY := "apikey " + os.Getenv("SG_API_TOKEN")
	SG_ORG := "demo-org"
	SG_WF_GROUP := "sg-sdk-go-test"
	SG_WF := "2aumphefkejtj3bv4q3wo"
	SG_WF_RUN := "3yzuf1izgfw3"
	SG_STACK := "Stack-test"
	SG_STACK_WF := "ansible-Rfde"
	SG_STACK_WF_RUN := "haoc1yepi6p5"
	SG_BASE_URL := os.Getenv("SG_BASE_URL")
	t.Run("default", func(t *testing.T) {
		c := NewClient()
		assert.Empty(t, c.baseURL)
	})

	t.Run("base url", func(t *testing.T) {
		c := NewClient(
			option.WithBaseURL("test.co"),
		)
		assert.Equal(t, "test.co", c.baseURL)
	})

	t.Run("http client", func(t *testing.T) {
		httpClient := &http.Client{
			Timeout: 5 * time.Second,
		}
		c := NewClient(
			option.WithHTTPClient(httpClient),
		)
		assert.Empty(t, c.baseURL)
	})

	t.Run("http header", func(t *testing.T) {
		header := make(http.Header)
		header.Set("X-API-Tenancy", "test")
		c := NewClient(
			option.WithHTTPHeader(header),
		)
		assert.Empty(t, c.baseURL)
		assert.Equal(t, "test", c.header.Get("X-API-Tenancy"))
	})
	// Workflows

	t.Run("Create and delete workflow", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		createWorkflowRequest := sggosdk.Workflow{
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
		createResponse, err := c.Workflows.CreateWorkflow(context.Background(), SG_ORG, SG_WF_GROUP, &createWorkflowRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createResponse.Data.ResourceName)

		deleteResposnse, err := c.Workflows.DeleteWorkflow(context.Background(), SG_ORG, createResponse.Data.ResourceName, SG_WF_GROUP)
		assert.Equal(t, "Workflow "+createResponse.Data.ResourceName+" deleted", deleteResposnse.Msg)
		assert.Empty(t, err)
	})

	t.Run("Update workflow", func(t *testing.T) {
		c := NewClient(
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

	t.Run("get workflow", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.ReadWorkflow(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, SG_WF, response.Msg.ResourceName)
	})

	t.Run("ListAll workflow", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.ListAllWorkflows(context.Background(), SG_ORG, SG_WF_GROUP)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
		assert.Empty(t, err)
	})

	t.Run("List all artifacts (workflow)", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.ListAllWorkflowArtifacts(context.Background(), SG_ORG, "CUSTOM-7OeX", "test-terragrunt")
		assert.Empty(t, err)
		assert.Equal(t, 15055, response.Data.Artifacts["orgs/demo-org/wfgrps/test-terragrunt/wfs/CUSTOM-7OeX/artifacts/tfstate.json"].Size)
	})

	t.Run("workflow output", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Workflows.Outputs(context.Background(), SG_ORG, "CUSTOM-7OeX", "test-terragrunt")
		assert.Empty(t, err)
		assert.Equal(t, "Outputs retrived", response.Msg)
		assert.Equal(t, "stackguardian-proper-escargot", response.Data.Outputs["id"]["value"].(string))
	})

	t.Run("Update stack workflow", func(t *testing.T) {
		c := NewClient(
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

	t.Run("get stackworkflow", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.ReadStackWorkflow(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, SG_STACK_WF, response.Msg.ResourceName)

	})

	t.Run("ListAll stack workflow", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.ListAllStackWorkflows(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("List all artifacts (stack workflow)", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.ListAllStackWorkflowsArtifacts(context.Background(), SG_ORG, "stack1",
			"refeed2-null-resource-tf-JuNs", "refeed-test-nested-stackrunbug")
		assert.Empty(t, err)
		assert.Equal(t, "Outputs retrieved", response.Msg)
		assert.Equal(t, 654,
			response.Data.Artifacts["orgs/demo-org/wfgrps/refeed-test-nested-stackrunbug/stacks/stack1/wfs/refeed2-null-resource-tf-JuNs/artifacts/tfstate.json"].Size)
	})

	t.Run("stack workflow output", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflows.StackWorkflowOutputs(context.Background(), SG_ORG, "stack1",
			"refeed2-null-resource-tf-JuNs", "refeed-test-nested-stackrunbug")
		assert.Empty(t, err)
		assert.Equal(t, float64(13), response.Data.Outputs["message_length"]["value"].(float64))
	})

	// Workflow Runs
	t.Run("ListAll workflow runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.WorkflowRuns.ListAllWorkflowRuns(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("Get workflow runs stack", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.StackWorkflowRuns.ReadStackWorkflowRun(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, SG_STACK_WF_RUN)
		assert.Equal(t, "QUEUED", response.Msg.Statuses["pre_0_step"][0].Name)
		assert.Empty(t, err)

	})

	t.Run("Get workflow runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.WorkflowRuns.ReadWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN)
		assert.Equal(t, "QUEUED", response.Msg.Statuses["pre_0_step"][0].Name)
		assert.Empty(t, err)

	})

	t.Run("Create workflow runs", func(t *testing.T) {
		c := NewClient(
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
		newWfRunName := response.Data.GetExtraProperties()["ResourceName"].(string)
		assert.NotEmpty(t, newWfRunName)

	})

	t.Run("Approve workflow runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		approveWfRunRequest := sggosdk.WorkflowRunApproval{
			Approve: true,
			Message: sggosdk.String("Approved"),
		}
		err := c.WorkflowRuns.ApproveWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN,
			&approveWfRunRequest)
		// We expect an error since the workflow run doesnt have any approvals pending
		assert.Contains(t, err.Error(), "No approval pending")

	})

	t.Run("Approve workflow runs (stack)", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		approveWfRunRequest := sggosdk.WorkflowRunApproval{
			Approve: true,
			Message: sggosdk.String("Approved"),
		}
		err := c.StackWorkflowRuns.ApproveStackWorkflowRun(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, SG_STACK_WF_RUN,
			&approveWfRunRequest)
		// We expect an error since the workflow run doesnt have any approvals pending
		assert.Contains(t, err.Error(), "No approval pending")

	})

	t.Run("Get workflow runs logs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		logs, err := c.WorkflowRuns.ReadWorkflowRunLogs(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(logs.Msg), 1)
	})

	t.Run("Get workflow runs logs (stack)", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		logs, err := c.StackWorkflowRuns.ReadStackWorkflowRunLogs(context.Background(), SG_ORG, SG_STACK, SG_STACK_WF, SG_WF_GROUP, SG_STACK_WF_RUN)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(logs.Msg), 1)
	})

	t.Run("Cancel workflow runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		err := c.WorkflowRuns.CancelWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN)
		// We expect an error since the workflow run is already failed
		if err != nil {
			assert.Contains(t, err.Error(), "Error cancelling Workflow Run "+SG_WF_RUN)
		}
	})

	t.Run("Update workflow runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)

		updateWfRunRequest := sggosdk.PatchedWorkflowRun{
			DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "testAWSConnector",
					"integrationId": "/integrations/testAWSConnector"}}},
			WfType: sggosdk.WfTypeEnumTerraform.Ptr(),
			EnvironmentVariables: []*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("UpdatedValue")}}},
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
		updateWfRunResponse, err := c.WorkflowRuns.UpdateWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, SG_WF_RUN, &updateWfRunRequest)
		assert.Empty(t, err)
		assert.Equal(t, "Workflow Run "+SG_WF_RUN+" updated", updateWfRunResponse.Msg)
	})

	// Stacks
	// t.Run("Create and delete stack", func(t *testing.T) {
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL(SG_BASE_URL),
	// 	)
	// 	createStackRequest := sggosdk.Stack{
	// 		RunOnCreate: sggosdk.Bool(false),
	// 		DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
	// 			Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
	// 			Config: map[string]interface{}{
	// 				"profileName":   "DummyConnectorForGoSDK",
	// 				"integrationId": "/integrations/DummyConnectorForGoSDK"}}},
	// 		EnvironmentVariables: []*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
	// 			Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}},

	// 		Description: sggosdk.String("Dummy Stack for GoSDK"),
	// 		TemplatesConfig: &sggosdk.TemplatesConfig{
	// 			TemplateGroupId: sggosdk.String("/demo-org/ansible:4"),
	// 			Templates: []*sggosdk.TemplateWorkflow{{
	// 				NumberOfApprovalsRequired: sggosdk.Int(0),
	// 				Description:               sggosdk.String("Dummy Workflow for GoSDK"),
	// 				WfType:                    sggosdk.WfTypeEnumCustom.Ptr(),
	// 				Id:                        "cc0061e9-a75c-421b-a75b-ef918e9f4b28",
	// 			}},
	// 		},
	// 	}
	// 	createStackResponse, err := c.Stacks.CreateStack(context.Background(), SG_ORG, SG_WF_GROUP, &createStackRequest)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, createStackResponse.Data.Stack.ResourceName)
	// 	assert.Equal(t, "Stack "+createStackResponse.Data.Stack.ResourceName+" created", createStackResponse.Msg)

	// 	err = c.StackWorkflows.DeleteStackWorkflow(context.Background(), SG_ORG, createStackResponse.Data.Stack.ResourceName,
	// 		createStackResponse.Data.Workflows[0].ResourceName, SG_WF_GROUP)
	// 	assert.Empty(t, err)
	// 	deleteResponse, err := c.Stacks.DeleteStack(context.Background(), SG_ORG, createStackResponse.Data.Stack.ResourceName, SG_WF_GROUP)
	// 	assert.Empty(t, err)
	// 	assert.Equal(t, "Stack "+createStackResponse.Data.Stack.ResourceName+" deleted", deleteResponse.Msg)
	// 	assert.Empty(t, err)
	// })

	t.Run("Read stack", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Stacks.ReadStack(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.Equal(t, SG_STACK, response.Msg.ResourceName)
	})

	t.Run("Run stack", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		runStackRequest := sggosdk.StackAction{
			ActionType:   sggosdk.ActionTypeEnumApply,
			ResourceName: sggosdk.String("5srghvu1y7nn"),
		}
		response, err := c.StackWorkflowRuns.CreateStackRun(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP, &runStackRequest)
		assert.Empty(t, err)
		assert.Equal(t, "Stack run scheduled", response.Msg)
	})

	t.Run("ListAll stacks", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Stacks.ListAllStacks(context.Background(), SG_ORG, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("Get stack outputs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		response, err := c.Stacks.ReadStackOutputs(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(response.Msg), 1)
	})

	t.Run("update stack", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		createStackRequest := sggosdk.PatchedStack{
			DeploymentPlatformConfig: []*sggosdk.DeploymentPlatformConfig{{
				Kind: sggosdk.DeploymentPlatformConfigKindEnumAwsRbac,
				Config: map[string]interface{}{
					"profileName":   "DummyConnectorForGoSDK",
					"integrationId": "/integrations/DummyConnectorForGoSDK"}}},
			EnvironmentVariables: []*sggosdk.EnvVars{{Kind: sggosdk.EnvVarsKindEnumPlainText,
				Config: &sggosdk.EnvVarConfig{VarName: "test", TextValue: sggosdk.String("testValue")}}},

			Description: sggosdk.String("Dummy Stack for GoSDK"),
			TemplatesConfig: &sggosdk.TemplatesConfig{
				TemplateGroupId: sggosdk.String("/demo-org/ansible:4"),
				Templates: []*sggosdk.TemplateWorkflow{{
					NumberOfApprovalsRequired: sggosdk.Int(0),
					Description:               sggosdk.String("Dummy Workflow for GoSDK"),
					WfType:                    sggosdk.WfTypeEnumCustom.Ptr(),
					Id:                        sggosdk.String("cc0061e9-a75c-421b-a75b-ef918e9f4b28"),
				}},
			},
		}
		updateStackResponse, err := c.Stacks.UpdateStack(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP, &createStackRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateStackResponse.Msg)
		assert.Equal(t, "Stack "+SG_STACK+" updated", updateStackResponse.Msg)
	})

	t.Run("list all stack runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		listResponse, err := c.StackWorkflowRuns.ListAllStackRuns(context.Background(), SG_ORG, SG_STACK, SG_WF_GROUP)
		assert.Empty(t, err)
		assert.GreaterOrEqual(t, len(listResponse.Msg), 1)
	})

	t.Run("get stack runs", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		getStackRunResponse, err := c.StackWorkflowRuns.ReadStackRun(context.Background(), SG_ORG, SG_STACK, "5srghvu1y7nn", SG_WF_GROUP)
		assert.Empty(t, err)
		assert.NotEmpty(t, getStackRunResponse.Msg.ResourceName)
		assert.Equal(t, "/stackruns/5srghvu1y7nn", getStackRunResponse.Msg.ResourceName)
	})

	// Connectors and Cloud Connector Groups
	t.Run("create_and_delete_connector", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		connectorName := "test-connector-go-sdk-test"
		createConnectorRequest := sggosdk.Integration{
			ResourceName: sggosdk.String(connectorName),
			Description:  sggosdk.String("test-connector description"),
			Settings: &sggosdk.Settings{
				Kind: sggosdk.SettingsKindEnumAzureStatic,
				Config: []map[string]interface{}{
					{
						"armTenantId":       "1",
						"armSubscriptionId": "1",
						"armClientId":       "1",
						"armClientSecret":   "1",
					},
				},
			},
		}
		createConnectorResponse, err := c.Connectors.CreateConnector(context.Background(), SG_ORG, &createConnectorRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createConnectorResponse.Msg)
		assert.Equal(t, "Connector "+*createConnectorRequest.ResourceName+" created", *createConnectorResponse.Msg)
		//Check that the response contains the resource name
		assert.NotEmpty(t, createConnectorResponse.Data.ResourceName)
		assert.Equal(t, connectorName, createConnectorResponse.Data.ResourceName)

		deleteConnectorResponse, err := c.Connectors.DeleteConnector(context.Background(), connectorName, SG_ORG)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteConnectorResponse.Msg)
		assert.Equal(t, "Connector "+connectorName+" deleted", deleteConnectorResponse.Msg)
	})

	t.Run("read_connector", func(t *testing.T) {
		c := NewClient(
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
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		connectorName := "DummyConnectorForGoSDK"
		updateConnectorRequest := sggosdk.PatchedIntegration{
			ResourceName: sggosdk.String(connectorName),
			Description:  sggosdk.String("updated description"),
			Settings: &sggosdk.Settings{
				Kind: sggosdk.SettingsKindEnumAzureStatic,
				Config: []map[string]interface{}{
					{
						"armTenantId":       "1",
						"armSubscriptionId": "1",
						"armClientId":       "1",
						"armClientSecret":   "1",
					},
				},
			},
		}
		updateConnectorResponse, err := c.Connectors.UpdateConnector(context.Background(), *updateConnectorRequest.ResourceName, SG_ORG, &updateConnectorRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateConnectorResponse.Msg)
		assert.Equal(t, "Integration "+connectorName+" updated", *updateConnectorResponse.Msg)
	})

	t.Run("listall_connector", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		listAllConnectorResponse, err := c.Connectors.ListAllConnector(context.Background(), SG_ORG)
		assert.Empty(t, err)
		assert.NotEmpty(t, listAllConnectorResponse.Msg)
		assert.GreaterOrEqual(t, len(listAllConnectorResponse.Msg), 1)
	})

	//Workflow Groups
	t.Run("create_and_delete_workflow_group", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "test-wfg-go-sdk-test"
		createWorkflowGroupRequest := sggosdk.WorkflowGroup{
			ResourceName: sggosdk.String(workflowGroupName),
			Description:  sggosdk.String("test-workflowGroup description"),
			IsActive:     sggosdk.IsArchiveEnumZero.Ptr(),
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
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "go-sdk-test-wfg"
		updateWorkflowGroupRequest := sggosdk.PatchedWorkflowGroup{
			ResourceName: sggosdk.String(workflowGroupName),
			Description:  sggosdk.String("updated description"),
			IsActive:     sggosdk.IsArchiveEnumZero.Ptr(),
		}
		updateWorkflowGroupResponse, err := c.WorkflowGroups.UpdateWorkflowGroup(context.Background(), SG_ORG, *updateWorkflowGroupRequest.ResourceName, &updateWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+workflowGroupName+" updated", *updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "updated description", *updateWorkflowGroupResponse.Data.Description)
	})

	t.Run("update_nested_workflow_group", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "sg-sdk-go-test/1bger5ydab697a4jxe2gu"
		updateWorkflowGroupRequest := sggosdk.PatchedWorkflowGroup{
			// ResourceName: sggosdk.String(workflowGroupName),
			Description: sggosdk.String("updated description"),
			IsActive:    sggosdk.IsArchiveEnumZero.Ptr(),
		}
		updateWorkflowGroupResponse, err := c.WorkflowGroups.UpdateWorkflowGroup(context.Background(), SG_ORG, workflowGroupName, &updateWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group 1bger5ydab697a4jxe2gu updated", *updateWorkflowGroupResponse.Msg)
		assert.Equal(t, "updated description", *updateWorkflowGroupResponse.Data.Description)
	})

	t.Run("read_workflow_group", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "sg-sdk-go-test"
		readWorkflowGroupResponse, err := c.WorkflowGroups.ReadWorkflowGroup(context.Background(), SG_ORG, workflowGroupName)
		assert.Empty(t, err)
		assert.NotEmpty(t, readWorkflowGroupResponse.Msg)
		assert.Equal(t, workflowGroupName, *readWorkflowGroupResponse.Msg.ResourceName)
	})

	t.Run("read_nested_workflow_group", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		workflowGroupName := "sg-sdk-go-test/1bger5ydab697a4jxe2gu"
		readWorkflowGroupResponse, err := c.WorkflowGroups.ReadWorkflowGroup(context.Background(), SG_ORG, workflowGroupName)
		assert.Empty(t, err)
		assert.NotEmpty(t, readWorkflowGroupResponse.Msg)
		assert.Equal(t, "1bger5ydab697a4jxe2gu", *readWorkflowGroupResponse.Msg.ResourceName)
	})

	t.Run("listall_workflow_groups", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		listAllWorkflowGroupResponse, err := c.WorkflowGroups.ListAllWorkflowGroups(context.Background(), SG_ORG)
		assert.Empty(t, err)
		assert.NotEmpty(t, listAllWorkflowGroupResponse.Msg)
		assert.GreaterOrEqual(t, len(listAllWorkflowGroupResponse.Msg), 1)
	})

	//Nested Workflow Groups
	t.Run("listall_nested_workflow_groups", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		listAllWorkflowGroupResponse, err := c.WorkflowGroups.ListAllChildWorkflowGroups(context.Background(), SG_ORG, "sg-sdk-go-test")
		assert.Empty(t, err)
		assert.NotEmpty(t, listAllWorkflowGroupResponse.Msg)
		assert.GreaterOrEqual(t, len(listAllWorkflowGroupResponse.Msg), 1)
	})

	t.Run("create_and_delete_nested_workflow_group", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		parentWorkflowGroupName := "sg-sdk-go-test"
		createWorkflowGroupRequest := sggosdk.WorkflowGroup{
			Description: sggosdk.String("child workflowGroup description"),
			IsActive:    sggosdk.IsArchiveEnumZero.Ptr(),
		}
		createChildWorkflowGroupResponse, err := c.WorkflowGroups.CreateChildWorkflowGroup(
			context.Background(),
			SG_ORG,
			parentWorkflowGroupName,
			&createWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createChildWorkflowGroupResponse)
		assert.NotEmpty(t, createChildWorkflowGroupResponse.Msg)
		assert.Contains(t, *createChildWorkflowGroupResponse.Msg, "created")

		deleteWorkflowGroupResponse, err := c.WorkflowGroups.DeleteWorkflowGroup(context.Background(), SG_ORG,
			parentWorkflowGroupName+"/"+*createChildWorkflowGroupResponse.Data.ResourceName)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+*createChildWorkflowGroupResponse.Data.ResourceName+" deleted",
			*deleteWorkflowGroupResponse.Msg)
	})

	t.Run("create_and_delete_deep_nested_workflow_group", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		parentWorkflowGroupName := "sg-sdk-go-test/1bger5ydab697a4jxe2gu"
		createWorkflowGroupRequest := sggosdk.WorkflowGroup{
			Description: sggosdk.String("child workflowGroup description"),
			IsActive:    sggosdk.IsArchiveEnumZero.Ptr(),
		}
		createChildWorkflowGroupResponse, err := c.WorkflowGroups.CreateChildWorkflowGroup(
			context.Background(),
			SG_ORG,
			parentWorkflowGroupName,
			&createWorkflowGroupRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createChildWorkflowGroupResponse)
		assert.NotEmpty(t, createChildWorkflowGroupResponse.Msg)
		assert.Contains(t, *createChildWorkflowGroupResponse.Msg, "created")

		deleteWorkflowGroupResponse, err := c.WorkflowGroups.DeleteWorkflowGroup(context.Background(), SG_ORG,
			parentWorkflowGroupName+"/"+*createChildWorkflowGroupResponse.Data.ResourceName)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteWorkflowGroupResponse.Msg)
		assert.Equal(t, "Workflow Group "+*createChildWorkflowGroupResponse.Data.ResourceName+" deleted",
			*deleteWorkflowGroupResponse.Msg)
	})

	//Roles
	// t.Run("create_and_delete_roles", func(t *testing.T) {
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL(SG_BASE_URL),
	// 	)
	// 	roleName := "Go-SDK-Test-Role"
	// 	createRoleRequest := sggosdk.Role{
	// 		ResourceName: roleName,
	// 		Description:  sggosdk.String("role description"),
	// 		AllowedPermissions: map[string]interface{}{
	// 			"GET/api/v1/orgs/demo-org/policies/<policy>/": "",
	// 		},
	// 		Actions: []string{"demo-org"},
	// 	}
	// 	createRoleResponse, err := c.UsersRoles.CreateRole(context.Background(), SG_ORG, &createRoleRequest)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, createRoleResponse)
	// 	assert.Equal(t, "Role "+createRoleResponse.Data.ResourceName+" created", *createRoleResponse.Msg)

	// 	deleteRoleResponse, err := c.UsersRoles.DeleteRole(context.Background(), SG_ORG, roleName)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, deleteRoleResponse.Msg)
	// 	assert.Equal(t, "Role "+roleName+" deleted", *deleteRoleResponse.Msg)
	// })

	// t.Run("update_role", func(t *testing.T) {
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL(SG_BASE_URL),
	// 	)
	// 	roleName := "SDK-Test-Role"
	// 	updateRoleRequest := sggosdk.PatchedRole{
	// 		ResourceName: sggosdk.String(roleName),
	// 		Description:  sggosdk.String("updated description"),
	// 		AllowedPermissions: map[string]interface{}{
	// 			"GET/api/v1/orgs/demo-org/policies/<policy>/": "",
	// 		},
	// 		Actions: []string{"demo-org"},
	// 	}
	// 	updateRoleResponse, err := c.UsersRoles.UpdateRole(context.Background(), SG_ORG, *updateRoleRequest.ResourceName, &updateRoleRequest)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, updateRoleResponse.Msg)
	// 	assert.Equal(t, "Role /roles/"+roleName+" updated", *updateRoleResponse.Msg)
	// 	assert.Equal(t, "updated description", *updateRoleResponse.Data.Description)
	// })

	// t.Run("read_role", func(t *testing.T) {
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL(SG_BASE_URL),
	// 	)
	// 	roleName := "SDK-Test-Role"
	// 	readRoleResponse, err := c.UsersRoles.ReadRole(context.Background(), SG_ORG, roleName)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, readRoleResponse.Msg)
	// 	assert.Equal(t, roleName, readRoleResponse.Msg.ResourceName)
	// })

	// t.Run("listall_role", func(t *testing.T) {
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL(SG_BASE_URL),
	// 	)
	// 	listAllRoleResponse, err := c.UsersRoles.ListAllRoles(context.Background(), SG_ORG)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, listAllRoleResponse.Msg)
	// 	assert.GreaterOrEqual(t, len(listAllRoleResponse.Msg), 1)
	// })

	// Users
	t.Run("add_and_remove_users", func(t *testing.T) {
		c := NewClient(
			option.WithApiKey(API_KEY),
			option.WithBaseURL(SG_BASE_URL),
		)
		userName := "Dummy@dummy.dummy"
		createUserRequest := sggosdk.AddUserToOrganization{
			Role:   "Demo-role",
			UserId: userName,
		}
		createUserResponse, err := c.UsersRoles.AddUser(context.Background(), SG_ORG, &createUserRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, createUserResponse)
		assert.Equal(t, userName+" invited.", *createUserResponse.Msg)

		removeUserRequest := sggosdk.GetorRemoveUserFromOrganization{
			UserId: "eu-central-1_C6bwuggLI/local/" + userName,
		}
		deleteUserResponse, err := c.UsersRoles.RemoveUser(context.Background(), SG_ORG, &removeUserRequest)
		assert.Empty(t, err)
		assert.NotEmpty(t, deleteUserResponse.Msg)
		assert.Equal(t, userName+" removed from /orgs/demo-org", *deleteUserResponse.Msg)
	})

	// t.Run("read_users", func(t *testing.T) {
	// 	principalHeader := http.Header{}
	// 	principalHeader.Add("PrincipalId", "test")
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL("http://localhost:8000"),
	// 		option.WithHTTPHeader(principalHeader),
	// 	)
	// 	// c := NewClient(
	// 	// 	option.WithApiKey(API_KEY),
	// 	// 	option.WithBaseURL(SG_BASE_URL),
	// 	// )
	// 	userName := "test@dummy.com"

	// 	removeUserRequest := sggosdk.GetorRemoveUserFromOrganization{
	// 		UserId: userName,
	// 	}
	// 	getUserResponse, err := c.UsersRoles.GetUser(context.Background(), SG_ORG, &removeUserRequest)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, getUserResponse.Msg)
	// })

	// t.Run("update_users", func(t *testing.T) {
	// 	principalHeader := http.Header{}
	// 	principalHeader.Add("PrincipalId", "test")
	// 	c := NewClient(
	// 		option.WithApiKey(API_KEY),
	// 		option.WithBaseURL("http://localhost:8000"),
	// 		option.WithHTTPHeader(principalHeader),
	// 	)
	// 	// c := NewClient(
	// 	// 	option.WithApiKey(API_KEY),
	// 	// 	option.WithBaseURL(SG_BASE_URL),
	// 	// )
	// 	userName := "test@dummy.com"

	// 	updateUserRequest := sggosdk.AddUserToOrganization{
	// 		Role:   "SDK-Test-Role",
	// 		UserId: userName,
	// 	}
	// 	updateUserResponse, err := c.UsersRoles.UpdateUser(context.Background(), SG_ORG, &updateUserRequest)
	// 	assert.Empty(t, err)
	// 	assert.NotEmpty(t, updateUserResponse.Msg)
	// })

}
