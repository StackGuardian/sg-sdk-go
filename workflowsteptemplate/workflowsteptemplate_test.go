package workflowsteptemplate_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	option "github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/workflowsteptemplate"
	"github.com/stretchr/testify/assert"
)

const (
	org                      = "sg-provider-test"
	workflowStepTemplateType = "WORKFLOW_STEP"
	workflowStepTemplateName = "sg-sdk-go-workflow-step-template"
	dockerImage              = "https://hub.docker.com/layers/library/redis/latest/images/sha256-8895092a9e19016d0e094d9bb45dfe0ddd4dfd08a2fc44d35cfb81c852f5d4d3"
	templateId               = "sg-sdk-go-workflow-step-template"
)

func getNewClient() *workflowsteptemplate.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := workflowsteptemplate.NewClient(option.WithApiKey(os.Getenv("API_KEY")), option.WithBaseURL(os.Getenv("API_URI")), option.WithHTTPHeader(header))

	return client
}

func TestCreateWorkflowStepTemplate(t *testing.T) {
	isActive := workflowsteptemplate.IsPublicEnum("0")
	description := "test description"
	var workflowStepTemplate = workflowsteptemplate.CreateWorkflowStepTemplate{
		TemplateType:     workflowStepTemplateType,
		SourceConfigKind: workflowsteptemplate.WorkflowStepTemplateSourceConfigKindDockerImageEnum,
		ShortDescription: &description,
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

	client := getNewClient()
	resp, err := client.CreateWorkflowStepTemplate(context.TODO(), org, false, &workflowStepTemplate)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, workflowStepTemplateName, resp.Data.Parent.TemplateName)
}

func TestReadWorkflowStepTemplate(t *testing.T) {
	client := getNewClient()

	resp, err := client.ReadWorkflowStepTemplate(context.TODO(), org, templateId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, templateId, resp.Msg.TemplateName)
}

func TestUpdateWorkflowStepTemplate(t *testing.T) {
	client := getNewClient()

	updatedTemplateName := "sg-sdk-go-updated-workflow-step-template"
	requestPayload := workflowsteptemplate.UpdateWorkflowStepTemplateRequestModel{
		TemplateType:     sgsdkgo.Optional(workflowsteptemplate.TemplateTypeEnum("IAC")),
		SourceConfigKind: sgsdkgo.Optional(workflowsteptemplate.WorkflowStepTemplateSourceConfigKindDockerImageEnum),
		TemplateName:     sgsdkgo.Optional(updatedTemplateName),
		ContextTags:      sgsdkgo.Null[map[string]string](),
		OwnerOrg:         sgsdkgo.Optional(fmt.Sprintf("/orgs/%v", org)),
		RuntimeSource: sgsdkgo.Optional(workflowsteptemplate.WorkflowStepRuntimeSource{
			Config: &workflowsteptemplate.WorkflowStepRuntimeSourceConfig{
				DockerImage: dockerImage,
			},
			SourceConfigDestKind: workflowsteptemplate.SourceConfigDestKindContainerRegistryEnum,
		}),
	}

	resp, err := client.UpdateWorkflowStepTemplate(context.TODO(), org, templateId, &requestPayload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, updatedTemplateName, resp.Data.TemplateName)
}

func TestDeleteWorkflowStepTemplate(t *testing.T) {
	client := getNewClient()
	err := client.DeleteWorkflowStepTemplate(context.TODO(), org, templateId)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
