package workflowtemplates_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplates"
	"github.com/stretchr/testify/assert"
)

const (
	org                  = "sg-provider-test"
	workflowTemplateType = "IAC"
	workflowTemplateName = "sgsdkgo workflow template"
	ownerOrg             = "sg-provider-test"
	templateId           = "sgsdkgo-workflow-template"
	repo                 = "https://github.com/taherkk/taher-null-resource.git"
)

func getNewClient() *workflowtemplates.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := workflowtemplates.NewClient(option.WithApiKey(os.Getenv("API_KEY")), option.WithBaseURL(os.Getenv("API_URI")), option.WithHTTPHeader(header))

	return client
}

func TestReadWorkflowTemplate(t *testing.T) {
	client := getNewClient()

	resp, err := client.ReadWorkflowTemplate(context.TODO(), org, templateId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, templateId, *resp.Msg.Id)
}

func TestCreateWorkflowTemplate(t *testing.T) {
	description := "test description"
	var templateId = templateId
	var isPrivate = false
	var workflowStepTemplate = workflowtemplates.CreateWorkflowTemplateRequest{
		Id:               &templateId,
		TemplateType:     sgsdkgo.TemplateTypeEnumIac,
		SourceConfigKind: workflowtemplates.WorkflowTemplateSourceConfigKindTerraform.Ptr(),
		ShortDescription: &description,
		TemplateName:     workflowTemplateName,
		ContextTags: map[string]string{
			"test": "tag-testing",
		},
		OwnerOrg: fmt.Sprintf("/orgs/%v", org),
		RuntimeSource: &workflowtemplates.RuntimeSource{
			Config: &workflowtemplates.RuntimeSourceConfig{
				IsPrivate: &isPrivate,
				Repo:      repo,
			},
			SourceConfigDestKind: workflowtemplates.SourceConfigDestKindEnumGithubCom.Ptr(),
		},
		IsActive: sgsdkgo.IsPublicEnumZero.Ptr(),
	}

	client := getNewClient()
	resp, err := client.CreateWorkflowTemplate(context.TODO(), org, false, &workflowStepTemplate)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, workflowTemplateName, resp.Data.Parent.TemplateName)
}

func TestUpdateWorkflowStepTemplate(t *testing.T) {
	client := getNewClient()

	updatedTemplateName := "sgsdkgo updated workflow template"
	updatedRef := "main"
	isPrivate := false
	requestPayload := workflowtemplates.UpdateWorkflowTemplateRequest{
		SourceConfigKind: sgsdkgo.Optional(workflowtemplates.WorkflowTemplateSourceConfigKindTerraform),
		TemplateName:     sgsdkgo.Optional(updatedTemplateName),
		ContextTags:      sgsdkgo.Null[map[string]string](),
		OwnerOrg:         sgsdkgo.Optional(fmt.Sprintf("/orgs/%v", org)),
		RuntimeSource: sgsdkgo.Optional(workflowtemplates.RuntimeSourceUpdate{
			Config: &workflowtemplates.RuntimeSourceConfigUpdate{
				IsPrivate: &isPrivate,
				Ref:       &updatedRef,
			},
			SourceConfigDestKind: workflowtemplates.SourceConfigDestKindEnumGithubCom.Ptr(),
		}),
	}

	resp, err := client.UpdateWorkflowTemplate(context.TODO(), org, templateId, &requestPayload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, updatedTemplateName, resp.Data.TemplateName)
}

func TestDeleteWorkflowTemplate(t *testing.T) {
	client := getNewClient()
	err := client.DeleteWorkflowTemplate(context.TODO(), org, templateId)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
