package stackworkflowrunfacts

import (
	"context"
	"net/http"

	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Stack Workflow Run Facts API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Stack Workflow Run Facts service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// GetStackWorkflowRunFacts gets the workflow run facts of a Stack workflow.
//
// This endpoint returns a signed URL which can be used to get the full contents of the Stack Workflow run facts.
// This signed URL is valid for 60 minutes. After expiration, you can request a new signed URL by calling this endpoint again.
//
// For more information, please refer to https://github.com/StackGuardian/feedback/discussions/109
func (s *Service) GetStackWorkflowRunFacts(ctx context.Context, org, stack, wf, wfGrp, wfRun, wfRunFacts string) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/wfrunfacts/%v/", org, wfGrp, stack, wf, wfRun, wfRunFacts)
	if err != nil {
		return err
	}

	err = s.client.Do(ctx, &internal.RequestOptions{
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		return err
	}

	return nil
}
