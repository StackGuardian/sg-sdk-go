package main

import (
	"context"
	"fmt"
	"log"
	"os"

	api "github.com/StackGuardian/sg-sdk-go"
	sg "github.com/StackGuardian/sg-sdk-go/stackguardian"
)

func main() {
	// Create and configure the client
	config := sg.DefaultConfig()
	config.APIKey = "apikey " + os.Getenv("SG_API_TOKEN")

	client, err := sg.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Define your organization, workflow, and workflow group
	org := os.Getenv("SG_ORG")
	if org == "" {
		org = "demo-org"
	}
	wfGroup := "sg-sdk-go-test"
	workflow := "my-workflow"

	// Define the workflow run request
	workflowRun := &api.WorkflowRun{
		DeploymentPlatformConfig: []*api.DeploymentPlatformConfig{{
			Kind: api.DeploymentPlatformConfigKindEnumAwsRbac,
			Config: map[string]interface{}{
				"profileName":   "testAWSConnector",
				"integrationId": "/integrations/testAWSConnector",
			},
		}},
		WfType: api.WfTypeEnumTerraform.Ptr(),
		EnvironmentVariables: []*api.EnvVars{{
			Kind: api.EnvVarsKindEnumPlainText,
			Config: &api.EnvVarConfig{
				VarName:   "test",
				TextValue: api.String("testValue"),
			},
		}},
		VcsConfig: &api.VcsConfig{
			IacVcsConfig: &api.IacvcsConfig{
				IacTemplateId:          api.String("/stackguardian/aws-s3-demo-website:16"),
				UseMarketplaceTemplate: true,
			},
			IacInputData: &api.IacInputData{
				SchemaType: api.IacInputDataSchemaTypeEnumFormJsonschema,
				Data: map[string]interface{}{
					"bucket_region": "eu-central-1",
				},
			},
		},
		UserJobCpu:    api.Int(512),
		UserJobMemory: api.Int(1024),
		RunnerConstraints: &api.RunnerConstraints{
			Type: "shared",
		},
	}

	// Create the workflow run
	ctx := context.Background()
	fmt.Println("Creating workflow run...")
	response, err := client.WorkflowRuns.CreateWorkflowRun(
		ctx,
		org,
		workflow,
		wfGroup,
		workflowRun,
	)
	if err != nil {
		log.Fatalf("Failed to create workflow run: %v", err)
	}

	// Get the resource name from the response
	resourceName := response.Data.GetExtraProperties()["ResourceName"].(string)
	fmt.Printf("✓ Created workflow run: %s\n", resourceName)

	// Check the status
	fmt.Println("\nFetching workflow run status...")
	runStatus, err := client.WorkflowRuns.ReadWorkflowRun(
		ctx,
		org,
		workflow,
		wfGroup,
		resourceName,
	)
	if err != nil {
		log.Fatalf("Failed to read workflow run: %v", err)
	}

	fmt.Printf("✓ Workflow run status: %+v\n", runStatus.Msg.Statuses)
}
