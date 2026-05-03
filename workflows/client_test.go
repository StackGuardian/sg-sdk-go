package workflows_test

import (
	"context"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/workflows"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testOrg = "sg-provider-test"
)

// getNewClient creates a workflows client with test configuration
func getNewClient() *workflows.Client {
	workflowClient := workflows.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
	)
	return workflowClient
}

// getMainClient creates a main SDK client for accessing WorkflowGroups and other services
func getMainClient() *client.Client {
	mainClient := client.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
	)
	return mainClient
}

// createAndCleanupWorkflowGroup creates a workflow group for a test and defers its deletion
// Returns the workflow group name to use in the test
func createAndCleanupWorkflowGroup(t *testing.T, mainClient *client.Client, baseName string) string {
	wfgName := baseName
	wfgRequest := &sgsdkgo.WorkflowGroup{
		ResourceName: sgsdkgo.String(wfgName),
		Description:  sgsdkgo.String("Test workflow group"),
	}

	resp, err := mainClient.WorkflowGroups.CreateWorkflowGroup(context.Background(), testOrg, wfgRequest)
	if err != nil {
		t.Fatalf("Failed to create test workflow group %s: %v", wfgName, err)
	}

	createdWfgName := *resp.Data.ResourceName

	// Defer cleanup of the workflow group after test completes
	t.Cleanup(func() {
		_, err := mainClient.WorkflowGroups.DeleteWorkflowGroup(context.Background(), testOrg, createdWfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow group %s: %v", createdWfgName, err)
		}
	})

	return createdWfgName
}

// TestCreateWorkflowWithId tests creating a new workflow with explicit ID
func TestCreateWorkflowWithId(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-create-with-id")

	workflowId := "test-workflow-001"
	resourceName := "test-workflow"
	description := "Test workflow description"
	tags := []string{"test", "automation"}

	request := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		Description:  &description,
		Tags:         tags,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	resp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, request)
	if err != nil {
		t.Fatalf("CreateWorkflow failed: %v", err)
	}

	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Msg)
	assert.NotNil(t, resp.Data)
	assert.Equal(t, resourceName, resp.Data.ResourceName)
	assert.Equal(t, description, resp.Data.Description)

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, resp.Data.Id, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", resp.Data.Id, err)
		}
	}()
}

// TestCreateWorkflow tests creating a new workflow without explicit ID
func TestCreateWorkflow(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-create-auto")

	resourceName := "test-workflow-auto"
	description := "Test workflow description"
	tags := []string{"test", "automation"}

	request := &workflows.Workflow{
		ResourceName: &resourceName,
		Description:  &description,
		Tags:         tags,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	resp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, request)
	if err != nil {
		t.Fatalf("CreateWorkflow failed: %v", err)
	}

	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Msg)
	assert.NotNil(t, resp.Data)
	assert.Equal(t, resourceName, resp.Data.ResourceName)
	assert.Equal(t, description, resp.Data.Description)

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, resp.Data.Id, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", resp.Data.Id, err)
		}
	}()
}

// TestCreateWorkflowWithParallelExecution tests creating a workflow with parallel execution enabled
func TestCreateWorkflowWithParallelExecution(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-create-parallel")

	workflowId := "test-workflow-parallel"
	resourceName := "test-workflow-parallel"
	description := "Workflow with parallel execution"

	request := &workflows.Workflow{
		Id:                &workflowId,
		ResourceName:      &resourceName,
		Description:       &description,
		IsActive:          sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:            sgsdkgo.WfTypeEnumCustom.Ptr(),
		ParallelExecution: sgsdkgo.ParallelExecutionEnumEnabled.Ptr(),
	}

	resp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, request)
	if err != nil {
		t.Fatalf("CreateWorkflow with parallel execution failed: %v", err)
	}

	require.NotNil(t, resp)
	assert.NotNil(t, resp.Data)
	assert.Equal(t, resourceName, resp.Data.ResourceName)

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, resp.Data.Id, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", resp.Data.Id, err)
		}
	}()
}

// TestCreateWorkflowMinimalConfig tests creating a workflow with minimal configuration
func TestCreateWorkflowMinimalConfig(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-create-minimal")

	workflowId := "test-workflow-minimal"
	resourceName := "test-workflow-minimal"

	request := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	resp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, request)
	if err != nil {
		t.Fatalf("CreateWorkflow with minimal config failed: %v", err)
	}

	require.NotNil(t, resp)
	assert.NotEmpty(t, resp.Msg)

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, resp.Data.Id, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", resp.Data.Id, err)
		}
	}()
}

// TestReadWorkflow tests reading an existing workflow
func TestReadWorkflow(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-read")

	// First, create a workflow
	workflowId := "test-read-workflow"
	resourceName := "test-read-workflow"
	description := "Workflow for testing read operation"
	createRequest := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		Description:  &description,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	createResp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, createRequest)
	if err != nil {
		t.Fatalf("Failed to create workflow for read test: %v", err)
	}

	createdWorkflowId := createResp.Data.Id

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", createdWorkflowId, err)
		}
	}()

	// Now read the workflow
	readResp, err := client.ReadWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
	if err != nil {
		t.Fatalf("ReadWorkflow failed: %v", err)
	}

	require.NotNil(t, readResp)
	require.NotNil(t, readResp.Msg)
	assert.NotEmpty(t, readResp.Msg.Id)
	assert.Equal(t, resourceName, *readResp.Msg.ResourceName)
}

// TestUpdateWorkflow tests updating an existing workflow
func TestUpdateWorkflow(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-update")

	// First, create a workflow
	workflowId := "test-update-workflow"
	resourceName := "test-update-workflow"
	originalDesc := "Original description"
	createRequest := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		Description:  &originalDesc,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	createResp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, createRequest)
	if err != nil {
		t.Fatalf("Failed to create workflow for update test: %v", err)
	}

	createdWorkflowId := createResp.Data.Id

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", createdWorkflowId, err)
		}
	}()

	// Update the workflow
	updatedDescription := "Updated description"
	updatedTags := []string{"updated", "test"}

	updateRequest := &workflows.PatchedWorkflow{
		Description: sgsdkgo.Optional(updatedDescription),
		Tags:        sgsdkgo.Optional(updatedTags),
	}

	updateResp, err := client.UpdateWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName, updateRequest)
	if err != nil {
		t.Fatalf("UpdateWorkflow failed: %v", err)
	}

	require.NotNil(t, updateResp)
	assert.NotEmpty(t, updateResp.Msg)
	assert.Equal(t, updatedDescription, updateResp.Data.Description)
}

// TestUpdateWorkflowParallelExecution tests updating parallel execution setting
func TestUpdateWorkflowParallelExecution(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-update-parallel")

	// Create a workflow
	workflowId := "test-update-parallel-workflow"
	resourceName := "test-update-parallel-workflow"
	createRequest := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	createResp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, createRequest)
	if err != nil {
		t.Fatalf("Failed to create workflow: %v", err)
	}

	createdWorkflowId := createResp.Data.Id

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", createdWorkflowId, err)
		}
	}()

	// Update parallel execution
	updateRequest := &workflows.PatchedWorkflow{
		ParallelExecution: sgsdkgo.Optional(sgsdkgo.ParallelExecutionEnumEnabled),
	}

	updateResp, err := client.UpdateWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName, updateRequest)
	if err != nil {
		t.Fatalf("UpdateWorkflow with parallel execution failed: %v", err)
	}

	require.NotNil(t, updateResp)
	assert.NotEmpty(t, updateResp.Msg)
}

// TestDeleteWorkflow tests deleting a workflow
func TestDeleteWorkflow(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-delete")

	// First, create a workflow to delete
	workflowId := "test-delete-workflow"
	resourceName := "test-delete-workflow"
	deleteDesc := "Workflow to be deleted"
	createRequest := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		Description:  &deleteDesc,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
	}

	createResp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, createRequest)
	if err != nil {
		t.Fatalf("Failed to create workflow for delete test: %v", err)
	}

	createdWorkflowId := createResp.Data.Id

	// Delete the workflow
	deleteResp, err := client.DeleteWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
	if err != nil {
		t.Fatalf("DeleteWorkflow failed: %v", err)
	}

	require.NotNil(t, deleteResp)
	assert.NotEmpty(t, deleteResp.Msg)

	// Cleanup: attempt to delete again (in case first delete failed), log error if already deleted
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
		if err != nil {
			t.Logf("Workflow already deleted or cleanup failed: %v", err)
		}
	}()
}

// TestWorkflowCRUDCycle tests a complete CRUD cycle
func TestWorkflowCRUDCycle(t *testing.T) {
	client := getNewClient()
	mainClient := getMainClient()

	// Create a test-specific workflow group
	wfgName := createAndCleanupWorkflowGroup(t, mainClient, "test-wfg-crud-cycle")

	// CREATE
	workflowId := "test-crud-workflow"
	resourceName := "test-crud-workflow"
	description := "Original description"
	crudTags := []string{"crud-test"}
	createRequest := &workflows.Workflow{
		Id:           &workflowId,
		ResourceName: &resourceName,
		Description:  &description,
		IsActive:     sgsdkgo.IsPublicEnumOne.Ptr(),
		WfType:       sgsdkgo.WfTypeEnumCustom.Ptr(),
		Tags:         crudTags,
	}

	createResp, err := client.CreateWorkflow(context.Background(), testOrg, wfgName, createRequest)
	require.NoError(t, err)
	require.NotNil(t, createResp)

	createdWorkflowId := createResp.Data.Id
	assert.Equal(t, resourceName, createResp.Data.ResourceName)
	assert.Equal(t, description, createResp.Data.Description)

	// Cleanup: delete the workflow after test
	defer func() {
		_, err := client.DeleteWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
		if err != nil {
			t.Logf("Failed to cleanup workflow %s: %v", createdWorkflowId, err)
		}
	}()

	// READ
	readResp, err := client.ReadWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName)
	require.NoError(t, err)
	require.NotNil(t, readResp)
	assert.NotEmpty(t, readResp.Msg.Id)

	// UPDATE
	updatedDescription := "Updated description in CRUD test"
	updateRequest := &workflows.PatchedWorkflow{
		Description: sgsdkgo.Optional(updatedDescription),
	}

	updateResp, err := client.UpdateWorkflow(context.Background(), testOrg, createdWorkflowId, wfgName, updateRequest)
	require.NoError(t, err)
	require.NotNil(t, updateResp)
	assert.Equal(t, updatedDescription, updateResp.Data.Description)
}
