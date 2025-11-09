package workflowrunfacts

import (
	"context"
	"net/http"

	"github.com/StackGuardian/sg-sdk-go/internal"
	api "github.com/StackGuardian/sg-sdk-go"
)

// Service provides access to the Workflow Run Facts API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Workflow Run Facts service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// ReadWorkflowRunFacts gets workflow run facts details inside a workflow.
//
// This endpoint returns a signed URL which can be used to get the full contents of the workflow run facts.
// This signed URL is valid for 60 minutes. After expiration, you can request a new signed URL by calling this endpoint again.
//
// For more information, please refer to https://github.com/StackGuardian/feedback/discussions/109
func (s *Service) ReadWorkflowRunFacts(ctx context.Context, org, wf, wfGrp, wfRun, wfRunFacts string) (map[string]interface{}, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/wfrunfacts/%v/", org, wfGrp, wf, wfRun, wfRunFacts)
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodGet,
		Path:     path,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
