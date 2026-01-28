package workflowsteptemplaterevision_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/core/workflowsteptemplate"
	workflowsteptemplaterevision "github.com/StackGuardian/sg-sdk-go/core/workflowstetemplaterevision"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/stretchr/testify/assert"
)

const (
	org                      = "sg-provider-test"
	workflowStepTemplateType = "WORKFLOW_STEP"
	workflowStepTemplateName = "test-workflow-step-template-revision"
	ownerOrg                 = "sg-provider-test"
	dockerImage              = "https://hub.docker.com/layers/library/redis/latest/images/sha256-8895092a9e19016d0e094d9bb45dfe0ddd4dfd08a2fc44d35cfb81c852f5d4d3"
	templateId               = "test-workflow-step-template-revision"
)

var revisionAlias = "revision1"

func GetWorkflowStepTemplateClient() *workflowsteptemplate.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := workflowsteptemplate.NewClient(option.WithApiKey(os.Getenv("API_KEY")), option.WithBaseURL(os.Getenv("API_URI")), option.WithHTTPHeader(header))

	return client
}

func workflowStempTemplateCreateFixture() error {
	isActive := workflowsteptemplate.IsPublicEnum("0")
	var workflowStepTemplate = workflowsteptemplate.CreateWorkflowStepTemplate{
		TemplateType:     workflowStepTemplateType,
		SourceConfigKind: workflowsteptemplate.WorkflowStepTemplateSourceConfigKindDockerImageEnum,
		TemplateName:     workflowStepTemplateName,
		ContextTags: map[string]string{
			"test": "tag-testing",
		},
		OwnerOrg: fmt.Sprintf("/orgs/%v", org),
		RuntimeSource: &workflowsteptemplate.WorkflowStepRuntimeSource{
			Config: &workflowsteptemplate.WorkflowStepRuntimeSourceConfig{
				DockerImage: dockerImage,
			},
			SourceConfigDestKind: workflowsteptemplate.SourceConfigDestKindContainerRegistryEnum,
		},
		IsActive: &isActive,
	}

	client := GetWorkflowStepTemplateClient()
	_, err := client.CreateWorkflowStepTemplate(context.TODO(), org, false, &workflowStepTemplate)
	if err != nil {
		return err
	}
	return nil
}

func workflowStepTemplateDeleteFixture() error {
	client := GetWorkflowStepTemplateClient()
	err := client.DeleteWorkflowStepTemplate(context.TODO(), org, templateId)
	if err != nil {
		return err
	}
	return nil
}

func getNewClient() *workflowsteptemplaterevision.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := workflowsteptemplaterevision.NewClient(option.WithApiKey(os.Getenv("API_KEY")), option.WithBaseURL(os.Getenv("API_URI")), option.WithHTTPHeader(header))

	return client
}

func TestCreateWorkflowStepTemplateRevision(t *testing.T) {
	err := workflowStempTemplateCreateFixture()
	if err != nil {
		t.Fatalf("failed to create parent template")
	}

	client := getNewClient()

	var createRevisionPayload = workflowsteptemplaterevision.CreateWorkflowStepTemplateRevisionModel{
		TemplateType:     workflowsteptemplate.TemplateTypeEnum("WORKFLOW_STEP"),
		Alias:            &revisionAlias,
		SourceConfigKind: workflowsteptemplate.WorkflowStepTemplateSourceConfigKindDockerImageEnum,
		RuntimeSource: &workflowsteptemplate.WorkflowStepRuntimeSource{
			Config: &workflowsteptemplate.WorkflowStepRuntimeSourceConfig{
				DockerImage: dockerImage,
			},
			SourceConfigDestKind: workflowsteptemplate.SourceConfigDestKindContainerRegistryEnum,
		},
		ContextTags: map[string]string{
			"test": "tag-testing",
		},
		OwnerOrg: fmt.Sprintf("/orgs/%v", org),
	}

	resp, err := client.CreateWorkflowStepTemplateRevision(context.TODO(), org, templateId, &createRevisionPayload)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, revisionAlias, *resp.Data.Revision.Alias)
}

func TestUpdateWorkflowStepTemplateRevision(t *testing.T) {
	client := getNewClient()

	newNotes := "changing long description"
	updatedTemplateName := workflowsteptemplaterevision.UpdateWorkflowStepTemplateRevisionModel{
		Notes: sgsdkgo.Optional(newNotes),
	}

	resp, err := client.UpdateWorkflowStepTemplateRevision(context.TODO(), org, fmt.Sprintf("%v:1", templateId), &updatedTemplateName)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, newNotes, *resp.Data.Notes)
}

func TestReadWorkflowStepTemplateRevision(t *testing.T) {
	client := getNewClient()

	resp, err := client.ReadWorkflowStepTemplateRevision(context.TODO(), org, fmt.Sprintf("%v:1", templateId))
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, revisionAlias, *resp.Msg.Alias)
}

func TestDeleteWorkflowStepTemplateRevision(t *testing.T) {
	client := getNewClient()

	err := client.DeleteWorkflowStepTemplateRevision(context.TODO(), org, fmt.Sprintf("%v:1", templateId), true)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
