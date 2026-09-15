package stackworkflowruns

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Stack Workflow Runs API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Stack Workflow Runs service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateStackWorkflowRun initiates a new workflow run for a specific stack.
func (s *Service) CreateStackWorkflowRun(ctx context.Context, org, stack, wf, wfGrp string, request *api.WorkflowRun) (*api.GeneratedWorkflowRunsStackCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/", org, wfGrp, stack, wf)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowRunsStackCreateResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodPost,
		Path:     path,
		Body:     request,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReadStackWorkflowRun retrieves detailed information about a specific stack workflow run.
func (s *Service) ReadStackWorkflowRun(ctx context.Context, org, stack, wf, wfGrp, wfRun string) (*api.GeneratedWorkflowRunStackGet, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/", org, wfGrp, stack, wf, wfRun)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowRunStackGet
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodGet,
		Path:     path,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReadStackWorkflowRunLogs retrieves execution logs for a stack workflow run.
//
// This endpoint returns a signed URL that can be used to fetch the logs in `text/plain` format.
// This signed URL is valid for 60 minutes. After expiration, you can request a new signed URL by calling this endpoint again.
func (s *Service) ReadStackWorkflowRunLogs(ctx context.Context, org, stack, wf, wfGrp, wfRun string) (*api.GeneratedWorkflowRunLogs, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/logs/", org, wfGrp, stack, wf, wfRun)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowRunLogs
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodGet,
		Path:     path,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ApproveStackWorkflowRun provides approval for a Stack Workflow run.
func (s *Service) ApproveStackWorkflowRun(ctx context.Context, org, stack, wf, wfGrp, wfRun string, request *api.WorkflowRunApproval) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/resume/", org, wfGrp, stack, wf, wfRun)
	if err != nil {
		return err
	}

	err = s.client.Do(ctx, &internal.RequestOptions{
		Method: http.MethodPost,
		Path:   path,
		Body:   request,
	})
	if err != nil {
		return err
	}

	return nil
}
