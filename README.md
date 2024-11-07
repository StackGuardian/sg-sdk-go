<a href="https://www.stackguardian.io/">
    <img src=".github/stackguardian_logo.svg" alt="StackGuardian logo" title="StackGuardian" align="right" height="40" />
</a>

# StackGuardian SDK For Go (sg-sdk-go)
`sg-sdk-go` is the StackGuardian SDK for the Go Programming language.

The SG SDK requires a minimum version of `Go 1.19`.

Check out the notes in the release for information about the latest bug fixes, updates and features added to the SDK.

### Getting started

It's recommended to store your API token and base URL in environment variables:
```
SG_BASE_URL (default: https://api.app.stackguardian.io)
SG_API_TOKEN
```

Install the SDK:
To get started working with the SDK, setup your project for Go modules and retrieve the SDK dependencies using `go get`.
```
go get github.com/StackGuardian/sg-sdk-go@v1.0.0
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

	// Create a new WorkflowRun using the client and request from above
	response, err := c.WorkflowRuns.CreateWorkflowRun(context.Background(),
		SG_ORG, SG_WF, SG_WF_GROUP, &createWorkflowRunRequest)
	if err != nil {
		fmt.Println(err)
	}
	// Get the resource name of the newly created WF run from the response
	var wfRunResourceName string = response.Data.GetExtraProperties()["ResourceName"].(string)

	// Get the status of the newly created WF run
	wfRunResponse, err := c.WorkflowRuns.ReadWorkflowRun(context.Background(), SG_ORG, SG_WF, SG_WF_GROUP, wfRunResourceName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wfRunResponse.Msg.Statuses["pre_0_step"][0].Name)

}
```


### Reporting bugs
If you encounter a bug with the SG SDK for Go we would like to hear about it. Please search the [existing issues](https://github.com/StackGuardian/sg-sdk-go/issues) and see if others are experiencing the same issue before opening a new one. 

Please include the version of the SG SDK for Go, the Go version and the OS you are using along with steps to replicate the issue when appropriate.

