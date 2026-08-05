package tests

import (
	"context"
	"os"
	"testing"

	"github.com/StackGuardian/sg-sdk-go/stackguardian"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration tests require SG_API_TOKEN environment variable
// Run with: SG_API_TOKEN=your-token go test ./tests/

func getTestClient(t *testing.T) *stackguardian.Client {
	t.Helper()

	apiKey := os.Getenv("SG_API_TOKEN")
	if apiKey == "" {
		t.Skip("SG_API_TOKEN environment variable not set")
	}

	config := stackguardian.DefaultConfig()
	config.APIKey = "apikey " + apiKey

	client, err := stackguardian.NewClient(config)
	require.NoError(t, err, "Failed to create client")

	return client
}

func TestIntegration_Organizations_ReadOrganization(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getTestClient(t)
	ctx := context.Background()

	// This assumes you have an org name in SG_ORG env var
	orgName := os.Getenv("SG_ORG")
	if orgName == "" {
		t.Skip("SG_ORG environment variable not set")
	}

	org, err := client.Organizations.ReadOrganization(ctx, orgName)

	require.NoError(t, err)
	assert.NotNil(t, org)
}

func TestIntegration_WorkflowGroups_ListAll(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := getTestClient(t)
	ctx := context.Background()

	orgName := os.Getenv("SG_ORG")
	if orgName == "" {
		t.Skip("SG_ORG environment variable not set")
	}

	// Test listing workflow groups
	// Note: request parameter is nil for default pagination
	groups, err := client.WorkflowGroups.ListAllWorkflowGroups(ctx, orgName, nil)

	require.NoError(t, err)
	assert.NotNil(t, groups)
}

// Add more integration tests as needed for your specific use cases
// Remember: Integration tests require valid API credentials
// They are automatically skipped in short mode: go test -short ./...
