<a href="https://www.stackguardian.io/">
    <img src=".github/stackguardian_logo.svg" alt="StackGuardian logo" title="StackGuardian" align="right" height="40" />
</a>

# StackGuardian SDK For Go (sg-sdk-go)

`sg-sdk-go` is the official StackGuardian SDK for the Go programming language.

[![Go Reference](https://pkg.go.dev/badge/github.com/StackGuardian/sg-sdk-go.svg)](https://pkg.go.dev/github.com/StackGuardian/sg-sdk-go)

The SDK requires a minimum version of **Go 1.19**.

## Features

- **Industry-standard structure** - Clean, maintainable architecture following Go best practices
- **Easy to use** - Simple, intuitive API with clear service organization
- **Built-in retry logic** - Automatic retries with exponential backoff for failed requests
- **Context support** - All methods accept `context.Context` for cancellation and timeouts
- **Nested workflow groups** - Full support for nested workflow groups (e.g., `parent/child`)
- **Comprehensive error handling** - Typed errors with detailed information
- **No external dependencies** - Uses only the Go standard library (except for testing)

## Installation

```bash
go get github.com/StackGuardian/sg-sdk-go@latest
```

## Quick Start

### Configuration

Set your API token and base URL as environment variables (recommended):

```bash
export SG_BASE_URL="https://api.app.stackguardian.io"  # Optional, this is the default
export SG_API_TOKEN="your-api-token-here"
```

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sg "github.com/StackGuardian/sg-sdk-go"
)

func main() {
	// Create a configuration
	config := sg.DefaultConfig()
	config.APIKey = "apikey " + os.Getenv("SG_API_TOKEN")
	// config.BaseURL is already set to the default

	// Create a new client
	client, err := sg.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Use the client to interact with StackGuardian
	ctx := context.Background()
	org, err := client.Organizations.ReadOrganization(ctx, "my-org")
	if err != nil {
		log.Fatalf("Failed to read organization: %v", err)
	}

	fmt.Printf("Organization: %+v\n", org)
}
```

## Advanced Example: Creating a Workflow Run

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sg "github.com/StackGuardian/sg-sdk-go"
)

func main() {
	// Create and configure the client
	config := sg.DefaultConfig()
	config.APIKey = "apikey " + os.Getenv("SG_API_TOKEN")

	client, err := sg.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Define the workflow run request
	workflowRun := &sg.WorkflowRun{
		DeploymentPlatformConfig: []*sg.DeploymentPlatformConfig{{
			Kind: sg.DeploymentPlatformConfigKindEnumAwsRbac,
			Config: map[string]interface{}{
				"profileName":   "testAWSConnector",
				"integrationId": "/integrations/testAWSConnector",
			},
		}},
		WfType: sg.WfTypeEnumTerraform.Ptr(),
		EnvironmentVariables: []*sg.EnvVars{{
			Kind: sg.EnvVarsKindEnumPlainText,
			Config: &sg.EnvVarConfig{
				VarName:   "test",
				TextValue: sg.String("testValue"),
			},
		}},
		VcsConfig: &sg.VcsConfig{
			IacVcsConfig: &sg.IacvcsConfig{
				IacTemplateId:          sg.String("/stackguardian/aws-s3-demo-website:16"),
				UseMarketplaceTemplate: true,
			},
			IacInputData: &sg.IacInputData{
				SchemaType: sg.IacInputDataSchemaTypeEnumFormJsonschema,
				Data: map[string]interface{}{
					"bucket_region": "eu-central-1",
				},
			},
		},
		UserJobCpu:    sg.Int(512),
		UserJobMemory: sg.Int(1024),
		RunnerConstraints: &sg.RunnerConstraints{
			Type: "shared",
		},
	}

	// Create the workflow run
	ctx := context.Background()
	response, err := client.WorkflowRuns.CreateWorkflowRun(
		ctx,
		"my-org",
		"my-workflow",
		"my-workflow-group",
		workflowRun,
	)
	if err != nil {
		log.Fatalf("Failed to create workflow run: %v", err)
	}

	// Get the resource name from the response
	resourceName := response.Data.GetExtraProperties()["ResourceName"].(string)
	fmt.Printf("Created workflow run: %s\n", resourceName)

	// Check the status
	runStatus, err := client.WorkflowRuns.ReadWorkflowRun(
		ctx,
		"my-org",
		"my-workflow",
		"my-workflow-group",
		resourceName,
	)
	if err != nil {
		log.Fatalf("Failed to read workflow run: %v", err)
	}

	fmt.Printf("Workflow run status: %+v\n", runStatus)
}
```

## Working with Stacks

```go
// Create a stack
stack := &sg.Stack{
	ResourceName: sg.String("my-stack"),
	Description:  sg.String("My infrastructure stack"),
	// ... other stack configuration
}

response, err := client.Stacks.CreateStack(ctx, "my-org", "my-wf-group", stack)
if err != nil {
	log.Fatalf("Failed to create stack: %v", err)
}

// Read a stack
stackDetails, err := client.Stacks.ReadStack(ctx, "my-org", "my-stack", "my-wf-group")
if err != nil {
	log.Fatalf("Failed to read stack: %v", err)
}

// List all stacks
listRequest := &sg.ListAllStacksRequest{
	Page:     sg.Int(1),
	PageSize: sg.Int(10),
}
stacks, err := client.Stacks.ListAllStacks(ctx, "my-org", "my-wf-group", listRequest)
if err != nil {
	log.Fatalf("Failed to list stacks: %v", err)
}
```

## Working with Nested Workflow Groups

The SDK fully supports nested workflow groups. Simply use the path notation:

```go
// Read a nested workflow group
wfGroup, err := client.WorkflowGroups.ReadWorkflowGroup(ctx, "my-org", "parent/child")
if err != nil {
	log.Fatalf("Failed to read workflow group: %v", err)
}

// Create a child workflow group
childGroup := &sg.WorkflowGroup{
	ResourceName: sg.String("grandchild"),
	Description:  sg.String("A nested workflow group"),
}
response, err := client.WorkflowGroups.CreateChildWorkflowGroup(
	ctx,
	"my-org",
	"parent/child",
	childGroup,
)
```

## Configuration Options

The SDK provides flexible configuration:

```go
config := &sg.Config{
	APIKey:       "apikey your-token",
	BaseURL:      "https://api.app.stackguardian.io",
	HTTPClient:   customHTTPClient,        // Optional: Use your own HTTP client
	MaxRetries:   5,                        // Optional: Default is 3
	RetryWaitMin: 2 * time.Second,         // Optional: Default is 1 second
	RetryWaitMax: 60 * time.Second,        // Optional: Default is 30 seconds
	UserAgent:    "my-app/1.0.0",          // Optional: Custom user agent
}

client, err := sg.NewClient(config)
```

Or use the default configuration:

```go
config := sg.DefaultConfig()
config.APIKey = "apikey " + os.Getenv("SG_API_TOKEN")
client, err := sg.NewClient(config)
```

## Error Handling

The SDK provides typed errors with detailed information:

```go
org, err := client.Organizations.ReadOrganization(ctx, "my-org")
if err != nil {
	// Check for specific error types
	if sg.IsNotFoundError(err) {
		fmt.Println("Organization not found")
	} else if sg.IsUnauthorizedError(err) {
		fmt.Println("Invalid API key or insufficient permissions")
	} else {
		fmt.Printf("Error: %v\n", err)
	}
	return
}
```

## Context and Cancellation

All API methods accept a `context.Context` for cancellation and timeouts:

```go
// Create a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

// Use the context
org, err := client.Organizations.ReadOrganization(ctx, "my-org")
if err != nil {
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Request timed out")
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
```

## Available Services

The SDK provides access to all StackGuardian services:

- **Organizations** - Manage organizations
- **AccessManagement** - Manage API keys, users, and roles
- **WorkflowGroups** - Manage workflow groups (including nested groups)
- **Workflows** - Manage workflows
- **WorkflowRuns** - Manage workflow runs
- **WorkflowRunFacts** - Get workflow run facts
- **Stacks** - Manage stacks
- **StackRuns** - Manage stack runs
- **StackWorkflows** - Manage stack workflows
- **StackWorkflowRuns** - Manage stack workflow runs
- **StackWorkflowRunFacts** - Get stack workflow run facts
- **Policies** - Manage policies
- **ConnectorGroups** - Manage connector groups
- **Connectors** - Manage connectors
- **RunnerGroups** - Manage runner groups
- **Secrets** - Manage secrets
- **Templates** - Manage templates
- **BenchmarkReports** - Get benchmark reports

## Migration from v1.x

If you're migrating from v1.x (Fern-generated SDK), here are the key changes:

### Client Creation

**Old:**
```go
import "github.com/StackGuardian/sg-sdk-go/client"
import "github.com/StackGuardian/sg-sdk-go/option"

c := client.NewClient(
	option.WithApiKey(apiKey),
	option.WithBaseURL(baseURL),
)
```

**New:**
```go
import sg "github.com/StackGuardian/sg-sdk-go"

config := sg.DefaultConfig()
config.APIKey = apiKey
config.BaseURL = baseURL
client, err := sg.NewClient(config)
```

### API Calls

**Old:**
```go
response, err := c.WorkflowRuns.CreateWorkflowRun(
	context.Background(),
	org, wf, wfGrp,
	request,
)
```

**New:**
```go
response, err := client.WorkflowRuns.CreateWorkflowRun(
	context.Background(),
	org, wf, wfGrp,
	request,
)
```

The API surface is largely the same, but the configuration is more straightforward and maintainable.

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## Reporting Issues

If you encounter bugs or have feature requests, please [open an issue](https://github.com/StackGuardian/sg-sdk-go/issues) on GitHub.

When reporting bugs, please include:
- SDK version
- Go version
- Operating system
- Steps to reproduce the issue
- Example code (if applicable)

## License

This SDK is distributed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for more information.

## Resources

- [StackGuardian Documentation](https://docs.stackguardian.io/)
- [API Documentation](https://docs.stackguardian.io/docs/api/overview)
- [Go Package Documentation](https://pkg.go.dev/github.com/StackGuardian/sg-sdk-go)
