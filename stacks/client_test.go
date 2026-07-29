package stacks_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/stacks"
	"github.com/StackGuardian/sg-sdk-go/stacktemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/stacktemplates"
	"github.com/StackGuardian/sg-sdk-go/workflowgroups"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplates"
	"github.com/stretchr/testify/assert"
)

const (
	org               = "sg-provider-test"
	ownerOrg          = "sg-provider-test"
	wfGrpId           = "sgsdkgo-stack-crud-wfgrp"
	wfTemplateId      = "sgsdkgo-workflow-template-for-stack-crud"
	wfTemplateName    = "sgsdkgo workflow template for stack CRUD tests"
	stackTemplateId   = "sgsdkgo-stack-template-for-stack-crud"
	stackTemplateName = "sgsdkgo stack template for stack CRUD tests"
	revisionAlias     = "v1"
	stackResourceName = "sgsdkgo-stack-crud-test"
)

func newStacksClient() *stacks.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return stacks.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func newWorkflowGroupClient() *workflowgroups.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return workflowgroups.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func newWorkflowTemplateClient() *workflowtemplates.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return workflowtemplates.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func newWorkflowTemplateRevisionClient() *workflowtemplaterevisions.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return workflowtemplaterevisions.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func newStackTemplateClient() *stacktemplates.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return stacktemplates.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func newStackTemplateRevisionClient() *stacktemplaterevisions.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return stacktemplaterevisions.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func createWorkflowGroup() error {
	id := wfGrpId
	_, err := newWorkflowGroupClient().CreateWorkflowGroup(context.TODO(), org, &sgsdkgo.WorkflowGroup{
		Id:           &id,
		ResourceName: &id,
	})
	return err
}

func deleteWorkflowGroup() error {
	_, err := newWorkflowGroupClient().DeleteWorkflowGroup(context.TODO(), org, wfGrpId)
	return err
}

func createWorkflowTemplate() error {
	id := wfTemplateId
	description := "workflow template fixture for stack CRUD tests"
	_, err := newWorkflowTemplateClient().CreateWorkflowTemplate(context.TODO(), org, false,
		&workflowtemplates.CreateWorkflowTemplateRequest{
			Id:               &id,
			TemplateName:     wfTemplateName,
			OwnerOrg:         fmt.Sprintf("/orgs/%v", org),
			SourceConfigKind: workflowtemplates.WorkflowTemplateSourceConfigKindTerraform.Ptr(),
			ShortDescription: &description,
			IsActive:         sgsdkgo.IsPublicEnumZero.Ptr(),
		},
	)
	return err
}

func deleteWorkflowTemplate() error {
	return newWorkflowTemplateClient().DeleteWorkflowTemplate(context.TODO(), org, wfTemplateId)
}

func createWorkflowTemplateRevision() error {
	_, err := newWorkflowTemplateRevisionClient().CreateWorkflowTemplateRevision(context.TODO(), org, wfTemplateId,
		&workflowtemplaterevisions.CreateWorkflowTemplateRevisionsRequest{
			Alias:            revisionAlias,
			Notes:            "initial revision for stack CRUD tests",
			OwnerOrg:         ownerOrg,
			SourceConfigKind: workflowtemplates.WorkflowTemplateSourceConfigKindTerraform.Ptr(),
		},
	)
	return err
}

func deleteWorkflowTemplateRevision() error {
	return newWorkflowTemplateRevisionClient().DeleteWorkflowTemplateRevision(context.TODO(), org, wfTemplateId+":1", true)
}

func createStackTemplate() error {
	id := stackTemplateId
	description := "stack template fixture for stack CRUD tests"
	_, err := newStackTemplateClient().CreateStackTemplate(context.TODO(), org, false,
		&stacktemplates.CreateStackTemplateRequest{
			Id:               &id,
			TemplateName:     stackTemplateName,
			OwnerOrg:         fmt.Sprintf("/orgs/%v", org),
			SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
			ShortDescription: &description,
			IsActive:         sgsdkgo.IsPublicEnumZero.Ptr(),
		},
	)
	return err
}

func deleteStackTemplate() error {
	return newStackTemplateClient().DeleteStackTemplate(context.TODO(), org, stackTemplateId)
}

func createStackTemplateRevision() error {
	_, err := newStackTemplateRevisionClient().CreateStackTemplateRevision(context.TODO(), org, stackTemplateId,
		&stacktemplaterevisions.CreateStackTemplateRevisionRequest{
			Alias:            revisionAlias,
			Notes:            "initial revision for stack CRUD tests",
			OwnerOrg:         ownerOrg,
			SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
			WorkflowsConfig: &stacktemplaterevisions.StackTemplateRevisionWorkflowsConfig{
				Workflows: []*stacktemplaterevisions.StackTemplateRevisionWorkflow{
					{TemplateId: sgsdkgo.String("/" + ownerOrg + "/" + wfTemplateId)},
				},
			},
		},
	)
	return err
}

func deleteStackTemplateRevision() error {
	return newStackTemplateRevisionClient().DeleteStackTemplateRevision(context.TODO(), org, stackTemplateId+":1", true)
}

// setupStackFixtures creates the workflow group, workflow template revision,
// and stack template revision a Stack needs to be instantiated from, failing
// the test on error.
//
// Each fixture's cleanup is registered via t.Cleanup *before* the create call
// that it undoes, not after: the API can create a resource but still return
// an error for the call that created it, so registering cleanup only on the
// success path would leave it behind. Registering first means it always
// runs, whether the create succeeds or fails - and deleting something that
// was never created is a harmless no-op we ignore.
//
// It also deletes any leftover Stack and fixtures from a previous run up
// front, leaf-to-root starting with the Stack itself, ignoring errors -
// there's usually nothing to delete. That covers the case a prior run's own
// cleanup never got to run at all, e.g. the test process was killed.
func setupStackFixtures(t *testing.T, client *stacks.Client) {
	t.Helper()

	deleteTestStack(client)
	deleteStackTemplateRevision()
	deleteStackTemplate()
	deleteWorkflowTemplateRevision()
	deleteWorkflowTemplate()
	deleteWorkflowGroup()

	t.Cleanup(func() { deleteWorkflowGroup() })
	if err := createWorkflowGroup(); err != nil {
		t.Fatalf("setup: createWorkflowGroup: %v", err)
	}

	t.Cleanup(func() { deleteWorkflowTemplate() })
	if err := createWorkflowTemplate(); err != nil {
		t.Fatalf("setup: createWorkflowTemplate: %v", err)
	}

	t.Cleanup(func() { deleteWorkflowTemplateRevision() })
	if err := createWorkflowTemplateRevision(); err != nil {
		t.Fatalf("setup: createWorkflowTemplateRevision: %v", err)
	}

	t.Cleanup(func() { deleteStackTemplate() })
	if err := createStackTemplate(); err != nil {
		t.Fatalf("setup: createStackTemplate: %v", err)
	}

	t.Cleanup(func() { deleteStackTemplateRevision() })
	if err := createStackTemplateRevision(); err != nil {
		t.Fatalf("setup: createStackTemplateRevision: %v", err)
	}
}

// newCreateStackRequest builds a Stack out of the stack template revision
// fixture: TemplateGroupId points at it, and the backend materializes the
// Stack's workflows from the template's WorkflowsConfig.
func newCreateStackRequest() *sgsdkgo.Stack {
	return &sgsdkgo.Stack{
		ResourceName:    sgsdkgo.String(stackResourceName),
		Description:     sgsdkgo.String("stack CRUD test fixture"),
		TemplateGroupId: sgsdkgo.String(fmt.Sprintf("/%v/%v:1", ownerOrg, stackTemplateId)),
	}
}

func deleteTestStack(client *stacks.Client) (*sgsdkgo.StackDeleteResponse, error) {
	return client.DeleteStack(context.TODO(), org, stackResourceName, wfGrpId,
		&sgsdkgo.DeleteStackRequest{ForceDelete: sgsdkgo.Bool(true)})
}

// createTestStack deletes any stack left behind under stackResourceName,
// registers cleanup before creating a fresh one (see setupStackFixtures for
// why), and fails the test on error.
func createTestStack(t *testing.T, client *stacks.Client) *sgsdkgo.GeneratedStackCreateResponse {
	t.Helper()
	deleteTestStack(client)

	t.Cleanup(func() { deleteTestStack(client) })
	resp, err := client.CreateStack(context.TODO(), org, wfGrpId, newCreateStackRequest())
	if err != nil {
		t.Fatalf("setup: CreateStack: %v", err)
	}

	return resp
}

func TestCreateStack(t *testing.T) {
	client := newStacksClient()
	setupStackFixtures(t, client)

	resp := createTestStack(t, client)

	assert.Equal(t, stackResourceName, *resp.Data.Stack.ResourceName)
}

func TestReadStack(t *testing.T) {
	client := newStacksClient()
	setupStackFixtures(t, client)

	createTestStack(t, client)

	resp, err := client.ReadStack(context.TODO(), org, stackResourceName, wfGrpId)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, stackResourceName, *resp.Msg.ResourceName)
}

func TestUpdateStack(t *testing.T) {
	client := newStacksClient()
	setupStackFixtures(t, client)

	createTestStack(t, client)

	updatedDescription := "updated stack description"
	resp, err := client.UpdateStack(context.TODO(), org, stackResourceName, wfGrpId,
		&sgsdkgo.PatchedStack{Description: sgsdkgo.Optional(updatedDescription)},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, updatedDescription, *resp.Data.Description)
}

func TestDeleteStack(t *testing.T) {
	client := newStacksClient()
	setupStackFixtures(t, client)

	createTestStack(t, client)

	resp, err := deleteTestStack(client)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Contains(t, *resp.Msg, "deleted")
}
