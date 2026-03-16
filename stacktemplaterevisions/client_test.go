package stacktemplaterevisions_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/stacktemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/stacktemplates"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplates"
	"github.com/stretchr/testify/assert"
)

const (
	org               = "sg-provider-test"
	ownerOrg          = "sg-provider-test"
	stackTemplateId   = "sgsdkgo-stack-template"
	stackTemplateName = "sgsdkgo stack template"
	wfTemplateId      = "sgsdkgo-workflow-template-for-stack"
	wfTemplateName    = "sgsdkgo workflow template for stack"
)

var revisionAlias = "revision1"

func newRevisionClient() *stacktemplaterevisions.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return stacktemplaterevisions.NewClient(
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

func newWorkflowTemplateClient() *workflowtemplates.Client {
	h := http.Header{}
	h.Set("x-sg-internal-auth-orgid", org)
	return workflowtemplates.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(h),
	)
}

func createWorkflowTemplate() error {
	id := wfTemplateId
	description := "workflow template fixture for stack template revision tests"
	request := workflowtemplates.CreateWorkflowTemplateRequest{
		Id:               &id,
		TemplateName:     wfTemplateName,
		OwnerOrg:         fmt.Sprintf("/orgs/%v", org),
		SourceConfigKind: workflowtemplates.WorkflowTemplateSourceConfigKindTerraform.Ptr(),
		ShortDescription: &description,
		IsActive:         sgsdkgo.IsPublicEnumZero.Ptr(),
	}
	_, err := newWorkflowTemplateClient().CreateWorkflowTemplate(context.TODO(), org, false, &request)
	return err
}

func deleteWorkflowTemplate() error {
	err := newWorkflowTemplateClient().DeleteWorkflowTemplate(context.TODO(), org, wfTemplateId)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func createStackTemplate() error {
	id := stackTemplateId
	description := "stack template fixture for revision tests"
	request := stacktemplates.CreateStackTemplateRequest{
		Id:               &id,
		TemplateName:     stackTemplateName,
		OwnerOrg:         fmt.Sprintf("/orgs/%v", org),
		SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
		ShortDescription: &description,
		IsActive:         sgsdkgo.IsPublicEnumZero.Ptr(),
	}
	_, err := newStackTemplateClient().CreateStackTemplate(context.TODO(), org, false, &request)
	return err
}

func deleteStackTemplate() error {
	return newStackTemplateClient().DeleteStackTemplate(context.TODO(), org, stackTemplateId)
}

func TestCreateStackTemplateRevision(t *testing.T) {
	client := newRevisionClient()
	defer deleteWorkflowTemplate()
	defer deleteStackTemplate()
	defer client.DeleteStackTemplateRevision(context.TODO(), org, stackTemplateId+":1", true)

	if err := createWorkflowTemplate(); err != nil {
		t.Fatalf("setup: createWorkflowTemplate: %v", err)
	}

	if err := createStackTemplate(); err != nil {
		t.Fatalf("setup: createStackTemplate: %v", err)
	}

	alias := "v1.0"
	resp, err := client.CreateStackTemplateRevision(context.TODO(), org, stackTemplateId,
		&stacktemplaterevisions.CreateStackTemplateRevisionRequest{
			Alias:            alias,
			Notes:            "initial stack template revision",
			OwnerOrg:         ownerOrg,
			SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
			WorkflowsConfig: &stacktemplaterevisions.StackTemplateRevisionWorkflowsConfig{
				Workflows: []*stacktemplaterevisions.StackTemplateRevisionWorkflow{
					{TemplateId: sgsdkgo.String("/" + ownerOrg + "/" + wfTemplateId)},
				},
			},
		},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, alias, resp.Data.Revision.Alias)
}

func TestReadStackTemplateRevision(t *testing.T) {
	client := newRevisionClient()

	defer deleteWorkflowTemplate()
	defer deleteStackTemplate()
	defer client.DeleteStackTemplateRevision(context.TODO(), org, stackTemplateId+":1", true)

	if err := createWorkflowTemplate(); err != nil {
		t.Fatalf("setup: createWorkflowTemplate: %v", err)
	}

	if err := createStackTemplate(); err != nil {
		t.Fatalf("setup: createStackTemplate: %v", err)
	}

	if _, err := client.CreateStackTemplateRevision(context.TODO(), org, stackTemplateId,
		&stacktemplaterevisions.CreateStackTemplateRevisionRequest{
			Alias:            "v1.0",
			Notes:            "initial stack template revision",
			OwnerOrg:         ownerOrg,
			SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
			WorkflowsConfig: &stacktemplaterevisions.StackTemplateRevisionWorkflowsConfig{
				Workflows: []*stacktemplaterevisions.StackTemplateRevisionWorkflow{
					{TemplateId: sgsdkgo.String("/" + ownerOrg + "/" + wfTemplateId)},
				},
			},
		},
	); err != nil {
		t.Fatalf("setup: CreateStackTemplateRevision: %v", err)
	}

	revisionId := stackTemplateId + ":1"
	resp, err := client.ReadStackTemplateRevision(context.TODO(), org, revisionId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, revisionId, *resp.Msg.Id)
}

func TestUpdateStackTemplateRevision(t *testing.T) {
	if err := createWorkflowTemplate(); err != nil {
		t.Fatalf("setup: createWorkflowTemplate: %v", err)
	}
	defer deleteWorkflowTemplate()

	if err := createStackTemplate(); err != nil {
		t.Fatalf("setup: createStackTemplate: %v", err)
	}
	defer deleteStackTemplate()

	client := newRevisionClient()
	if _, err := client.CreateStackTemplateRevision(context.TODO(), org, stackTemplateId,
		&stacktemplaterevisions.CreateStackTemplateRevisionRequest{
			Alias:            "v1.0",
			Notes:            "initial stack template revision",
			OwnerOrg:         ownerOrg,
			SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
			WorkflowsConfig: &stacktemplaterevisions.StackTemplateRevisionWorkflowsConfig{
				Workflows: []*stacktemplaterevisions.StackTemplateRevisionWorkflow{
					{TemplateId: sgsdkgo.String("/" + ownerOrg + "/" + wfTemplateId)},
				},
			},
		},
	); err != nil {
		t.Fatalf("setup: CreateStackTemplateRevision: %v", err)
	}
	defer client.DeleteStackTemplateRevision(context.TODO(), org, stackTemplateId+":1", true)

	revisionId := stackTemplateId + ":1"
	resp, err := client.UpdateStackTemplateRevision(context.TODO(), org, revisionId,
		&stacktemplaterevisions.UpdateStackTemplateRevisionRequest{
			Alias:    sgsdkgo.Optional(revisionAlias),
			Notes:    sgsdkgo.Optional("updated stack template revision notes"),
			OwnerOrg: sgsdkgo.Optional(ownerOrg),
		},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, revisionAlias, resp.Data.Alias.Value)
}

func TestDeleteStackTemplateRevision(t *testing.T) {
	if err := createWorkflowTemplate(); err != nil {
		t.Fatalf("setup: createWorkflowTemplate: %v", err)
	}
	defer deleteWorkflowTemplate()

	if err := createStackTemplate(); err != nil {
		t.Fatalf("setup: createStackTemplate: %v", err)
	}
	defer deleteStackTemplate()

	client := newRevisionClient()
	if _, err := client.CreateStackTemplateRevision(context.TODO(), org, stackTemplateId,
		&stacktemplaterevisions.CreateStackTemplateRevisionRequest{
			Alias:            "v1.0",
			Notes:            "initial stack template revision",
			OwnerOrg:         ownerOrg,
			SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
			WorkflowsConfig: &stacktemplaterevisions.StackTemplateRevisionWorkflowsConfig{
				Workflows: []*stacktemplaterevisions.StackTemplateRevisionWorkflow{
					{TemplateId: sgsdkgo.String("/" + ownerOrg + "/" + wfTemplateId)},
				},
			},
		},
	); err != nil {
		t.Fatalf("setup: CreateStackTemplateRevision: %v", err)
	}

	revisionId := stackTemplateId + ":1"
	if err := client.DeleteStackTemplateRevision(context.TODO(), org, revisionId, true); err != nil {
		t.Fatalf(err.Error())
	}
}
