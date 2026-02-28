package stacktemplaterevisions_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/stacktemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/stacktemplates"
	"github.com/stretchr/testify/assert"
)

var (
	org          = "sg-provider-test"
	ownerOrg     = "sg-provider-test"
	templateId   = "sgsdkgo-stack-template"
	templateNmae = "sgsdkgo stack template"
)

var revisionAlias = "revision1"

func getStackTemplateRevisionClient() *stacktemplaterevisions.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := stacktemplaterevisions.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(header),
	)

	return client
}

func TestCreateStackTemplateRevision(t *testing.T) {
	client := getStackTemplateRevisionClient()

	notes := "initial stack template revision"
	alias := "v1.0"
	payload := stacktemplaterevisions.CreateStackTemplateRevisionRequest{
		Alias:            alias,
		Notes:            notes,
		TempalteName:     templateNmae,
		OwnerOrg:         ownerOrg,
		SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
		WorkflowsConfig: &sgsdkgo.WorkflowsConfig{
			Workflows: []*sgsdkgo.WorkflowsConfigWorkflow{
				{
					TemplateId: sgsdkgo.String("test-template"),
				},
			},
		},
	}

	resp, err := client.CreateStackTemplateRevision(context.TODO(), org, templateId, &payload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, alias, resp.Data.Revision.Alias)
}

func TestReadStackTemplateRevision(t *testing.T) {
	client := getStackTemplateRevisionClient()

	revisionId := templateId + ":1"
	resp, err := client.ReadStackTemplateRevision(context.TODO(), org, revisionId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, revisionId, *resp.Msg.Id)
}

func TestUpdateStackTemplateRevision(t *testing.T) {
	client := getStackTemplateRevisionClient()

	updatedNotes := "updated stack template revision notes"
	payload := stacktemplaterevisions.UpdateStackTemplateRevisionRequest{
		Alias:            sgsdkgo.Optional(revisionAlias),
		Notes:            sgsdkgo.Optional(updatedNotes),
		OwnerOrg:         sgsdkgo.Optional(ownerOrg),
		SourceConfigKind: sgsdkgo.Optional(stacktemplates.StackTemplateSourceConfigKindMixed),
	}
	revisionId := templateId + ":1"
	resp, err := client.UpdateStackTemplateRevision(context.TODO(), org, revisionId, &payload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, revisionAlias, resp.Data.Alias)
}

func TestDeleteStackTemplateRevision(t *testing.T) {
	client := getStackTemplateRevisionClient()
	revisionId := templateId + ":2"
	err := client.DeleteStackTemplateRevision(context.TODO(), org, revisionId, true)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
