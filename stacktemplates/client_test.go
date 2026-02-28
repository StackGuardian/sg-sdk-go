package stacktemplates_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/stacktemplates"
	"github.com/stretchr/testify/assert"
)

const (
	org               = "sg-provider-test"
	stackTemplateName = "sgsdkgo stack template"
	templateId        = "sgsdkgo-stack-template"
)

func getNewClient() *stacktemplates.Client {
	header := http.Header{}
	header.Set("x-sg-internal-auth-orgid", org)

	client := stacktemplates.NewClient(
		option.WithApiKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("API_URI")),
		option.WithHTTPHeader(header),
	)

	return client
}

func TestReadStackTemplate(t *testing.T) {
	client := getNewClient()

	resp, err := client.ReadStackTemplate(context.TODO(), org, templateId)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	assert.Equal(t, templateId, *resp.Msg.Id)
}

func TestCreateStackTemplate(t *testing.T) {
	description := "test stack template description"
	var id = templateId
	request := stacktemplates.CreateStackTemplateRequest{
		Id:               &id,
		TemplateName:     stackTemplateName,
		SourceConfigKind: stacktemplates.StackTemplateSourceConfigKindMixed.Ptr(),
		ShortDescription: &description,
		OwnerOrg:         fmt.Sprintf("/orgs/%v", org),
		ContextTags: map[string]string{
			"test": "tag-testing",
		},
		IsActive: sgsdkgo.IsPublicEnumZero.Ptr(),
	}

	client := getNewClient()
	resp, err := client.CreateStackTemplate(context.TODO(), org, false, &request)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, stackTemplateName, resp.Data.Parent.TemplateName)
}

func TestUpdateStackTemplate(t *testing.T) {
	client := getNewClient()

	updatedTemplateName := "sgsdkgo updated stack template"
	requestPayload := stacktemplates.UpdateStackTemplateRequest{
		TemplateName:     sgsdkgo.Optional(updatedTemplateName),
		OwnerOrg:         sgsdkgo.Optional(fmt.Sprintf("/orgs/%v", org)),
		ContextTags:      sgsdkgo.Null[map[string]string](),
		SourceConfigKind: sgsdkgo.Optional(stacktemplates.StackTemplateSourceConfigKindMixed),
	}

	resp, err := client.UpdateStackTemplate(context.TODO(), org, templateId, &requestPayload)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, updatedTemplateName, *resp.Data.TemplateName)
}

func TestDeleteStackTemplate(t *testing.T) {
	client := getNewClient()
	err := client.DeleteStackTemplate(context.TODO(), org, templateId)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
