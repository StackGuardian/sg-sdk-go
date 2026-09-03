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


### Command-line interface

The module also ships `sg`, a command-line interface that exposes every SDK operation (one subcommand per API method).

```
go install github.com/StackGuardian/sg-sdk-go/cmd/sg@latest
# or, from a checkout:
make cli        # builds bin/sg
```

Configure it with flags or environment variables:

| Flag         | Environment variable | Description                                                   |
|--------------|----------------------|---------------------------------------------------------------|
| `--api-key`  | `SG_API_TOKEN`       | API token (the `apikey ` prefix is added if missing)          |
| `--base-url` | `SG_BASE_URL`        | API base URL (default `https://api.app.stackguardian.io`)      |
| `--org`      | `SG_ORG`             | Organization name used by every command                       |

```
export SG_API_TOKEN=... SG_ORG=demo-org

sg workflow-groups list --limit 10
sg workflows get my-wf --wfgrp my-group
sg workflow-runs create --wfgrp my-group --wf my-wf
sg workflow-runs logs <run-id> --wfgrp my-group --wf my-wf
sg workflow-runs approve <run-id> --wfgrp my-group --wf my-wf --reject --message "not now"
sg workflows create --wfgrp my-group --body-file workflow.json
echo '{"Description":"updated"}' | sg stacks update my-stack --wfgrp my-group -f -
sg templates get my-template:5 --type IAC --owner-org stackguardian
sg stack-workflow-runs get <run-id> --wfgrp my-group --stack my-stack --wf my-wf
```

- Create and update commands take the request body as JSON (`--body`/`-b` inline, or `--body-file`/`-f` with a path or `-` for stdin). Field names match the API schema; a JSON `null` is sent as an explicit null, and omitted fields are omitted from the request.
- Resources nested under workflow groups, stacks and workflows are scoped with `--wfgrp`, `--stack` and `--wf`; nested workflow groups are written `parent/child`.
- Responses are printed as the API returned them (pretty-printed JSON; `--compact` for one line). Errors go to stderr with exit code 1.
- `--verbose` logs each request and response status, `--timeout` and `--max-attempts` control the HTTP behaviour, and `sg completion <shell>` generates shell completions.
- Run `sg --help` or `sg <resource> --help` for the full command list. The response types were validated against a live organization; `SG_API_TOKEN=... SG_ORG=... go test ./tests/live/` repeats that check. Every SDK method has exactly one subcommand, so the CLI covers the whole API surface: organizations, users, roles, role bindings, API accesses and legacy API tokens, audit logs, benchmark reports, connectors and connector groups, policies, runner groups, secrets, state backends, templates (including artifacts, subscriptions and the public marketplace), workflow groups, stacks and stack runs, workflows and stack workflows (including artifacts, locks and comparisons), workflow runs and run facts, resource search and move, billing, and AI chats.
- Some endpoints only accept a user (Cognito) token rather than an API key, for example creating or listing organizations and the public template listing; the command help says so.

### Reporting bugs
If you encounter a bug with the SG SDK for Go we would like to hear about it. Please search the [existing issues](https://github.com/StackGuardian/sg-sdk-go/issues) and see if others are experiencing the same issue before opening a new one. 

Please include the version of the SG SDK for Go, the Go version and the OS you are using along with steps to replicate the issue when appropriate.

