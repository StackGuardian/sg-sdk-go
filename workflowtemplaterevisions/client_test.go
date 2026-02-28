package workflowtemplaterevisions_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplates"
	"github.com/stretchr/testify/assert"
)

var (
	org                  = "sg-provider-test"
	workflowTemplateType = "IAC"
	workflowTemplateName = "test workflow template revision"
	ownerOrg             = "sg-provider-test"
	templateId           = "sgsdkgo-workflow-template"
	alias                = "revision-alias"
)

var revisionAlias = "revision1"

func GetWorkflowTemplateRevisionClient() *workflowtemplaterevisions.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := workflowtemplaterevisions.NewClient(option.WithApiKey(os.Getenv("API_KEY")), option.WithBaseURL(os.Getenv("API_URI")), option.WithHTTPHeader(header))

	return client
}

func TestCreateWorkflowTemplateRevision(t *testing.T) {
	client := GetWorkflowTemplateRevisionClient()

	notest := "revision notes"
	revisionAlias := "revision2"
	payload := workflowtemplaterevisions.CreateWorkflowTemplateRevisionsRequest{
		Alias:            revisionAlias,
		Notes:            notest,
		OwnerOrg:         org,
		SourceConfigKind: workflowtemplates.WorkflowTemplateSourceConfigKindTerraform.Ptr(),
		TemplateType:     workflowTemplateType,
	}

	resp, err := client.CreateWorkflowTemplateRevision(context.TODO(), org, templateId, &payload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, revisionAlias, resp.Data.Revision.Alias)
}

func TestDeleteWorkflowTemplateRevision(t *testing.T) {
	client := GetWorkflowTemplateRevisionClient()
	revisionId := "sgsdkgo-workflow-template:3"
	err := client.DeleteWorkflowTemplateRevision(context.TODO(), org, revisionId, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetWokflowTemplateRevision(t *testing.T) {
	client := GetWorkflowTemplateRevisionClient()

	revisionId := templateId + ":1"
	resp, err := client.ReadWorkflowTemplateRevision(context.TODO(), org, revisionId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, revisionId, *resp.Msg.Id)
}

func TestUpdateWorkflowTemplateRevision(t *testing.T) {
	client := GetWorkflowTemplateRevisionClient()

	updatedNotes := "revision notes again"
	userJobMemory := 2048
	payload := workflowtemplaterevisions.UpdateWorkflowTemplateRevisionRequest{
		Alias:            sgsdkgo.Optional(revisionAlias),
		Notes:            sgsdkgo.Optional(updatedNotes),
		OwnerOrg:         sgsdkgo.Optional(org),
		SourceConfigKind: sgsdkgo.Optional(workflowtemplates.WorkflowTemplateSourceConfigKindTerraform),
		UserJobMemory:    sgsdkgo.Optional(userJobMemory),
	}
	revisionId := templateId + ":1"
	resp, err := client.UpdateWorkflowTemplateRevision(context.TODO(), org, revisionId, &payload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, revisionAlias, resp.Data.Alias)
}
