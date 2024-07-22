# StackGuardian Go SDK (sg-sdk-go)
Go SDK for StackGuardian APIs. This SDK is currently in alpha.

Currently, the SDK supports the following SG APIs:
- WorkflowRuns.GetWorkflowRun
- WorkflowRuns.GetWorkflowRuns(Stack)
- WorkflowRuns.ApproveWorkflowRun
- WorkflowRuns.ApproveWorkflowRun(Stack)
- WorkflowRuns.CreateWorkflowRun
- WorkflowRuns.CreateWorkflowRun(Stack)
- WorkflowRuns.ListWorkflowRuns
- WorkflowRuns.ListWorkflowRuns(Stack)
- WorkflowRuns.GetWorkflowRunLogs
- WorkflowRuns.GetWorkflowRunLogs(Stack)
- WorkflowRuns.CancelWorkflowRun
- WorkflowRuns.DeleteWorkflowRun
- WorkflowRuns.DeleteWorkflowRun(Stack)
- WorkflowRuns.UpdateWorkflowRun
- Workflows.Create
- Workflows.Delete
- Workflows.Update
- Workflow.Read
- Workflow.List
- Workflows.Output
- Workflows.ListAllArtifacts

### Setup

It's recommended to store your API token and base URL in environment variables:
```
SG_BASE_URL (default: https://api.app.stackguardian.io)
SG_API_TOKEN
```

Install the SDK:
```
go get github.com/stackguardian/sg-sdk-go
```

### Sample Usage

```go
import (
	"context"
	"fmt"
	"os"

	sggosdk "github.com/StackGuardian/sg-sdk-go"
	client "github.com/StackGuardian/sg-sdk-go/client"
	option "github.com/StackGuardian/sg-sdk-go/option"
)

func main() {

	// Define the API key, base URL, org and workflow details
	API_KEY := "apikey " + os.Getenv("SG_API_TOKEN")
	SG_ORG := "demo-org"
	SG_WF_GROUP := "sg-sdk-go-test"
	SG_WF := "2aumphefkejtj3bv4q3wo"
	SG_BASE_URL := os.Getenv("SG_BASE_URL")

	// Create a new client using the API key and base URL
	c := client.NewClient(
		option.WithApiKey(API_KEY),
		option.WithBaseURL(SG_BASE_URL),
	)

	// Create a new WorkflowRun request
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
			IacVcsConfig: &sggosdk.IacVcsConfig{
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

	// Create a new WorkflowRun using the client and request from above
	response, err := c.WorkflowRuns.CreateWorkflowRun(context.Background(),
		SG_ORG, SG_WF, SG_WF_GROUP, &createWorkflowRunRequest)
	if err != nil {
		fmt.Println(err)
	}
	// Get the resource name of the newly created WF run from the response
	var wfRunResourceName string = response.Data.GetExtraProperties()["ResourceName"].(string)

	// Get the status of the newly created WF run
	wfRunResponse, err := c.WorkflowRuns.GetWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, wfRunResourceName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wfRunResponse.Msg.Statuses["pre_0_step"][0].Name)

}
```

