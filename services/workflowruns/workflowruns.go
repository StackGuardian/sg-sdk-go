package workflowruns

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Workflow Runs API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Workflow Runs service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateWorkflowRun initiates a new workflow run.
func (s *Service) CreateWorkflowRun(ctx context.Context, org, wf, wfGrp string, request *api.WorkflowRun) (*api.WorkflowRunCreatePatchResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowRunCreatePatchResponse
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

// ReadWorkflowRun retrieves the details of an existing workflow run.
func (s *Service) ReadWorkflowRun(ctx context.Context, org, wf, wfGrp, wfRun string) (*api.GeneratedWorkflowRunsGet, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/", org, wfGrp, wf, wfRun)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowRunsGet
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

// UpdateWorkflowRun updates the details of a workflow run.
func (s *Service) UpdateWorkflowRun(ctx context.Context, org, wf, wfGrp, wfRun string, request *api.PatchedWorkflowRun) (*api.GeneratedWorkfkowRunsUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/", org, wfGrp, wf, wfRun)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkfkowRunsUpdateResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodPatch,
		Path:     path,
		Body:     request,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CancelWorkflowRun cancels a running or queued workflow run.
func (s *Service) CancelWorkflowRun(ctx context.Context, org, wf, wfGrp, wfRun string) (*api.WorkflowRunsCancelResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/cancel/", org, wfGrp, wf, wfRun)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowRunsCancelResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodPatch,
		Path:     path,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReadWorkflowRunLogs retrieves execution logs for a workflow run.
//
// This endpoint returns a signed URL that can be used to fetch the logs in `text/plain` format.
// This signed URL is valid for 60 minutes. After expiration, you can request a new signed URL by calling this endpoint again.
func (s *Service) ReadWorkflowRunLogs(ctx context.Context, org, wf, wfGrp, wfRun string) (*api.GeneratedWorkflowRunLogs, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/logs/", org, wfGrp, wf, wfRun)
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

// ApproveWorkflowRun provides approval for a Workflow Run.
func (s *Service) ApproveWorkflowRun(ctx context.Context, org, wf, wfGrp, wfRun string, request *api.WorkflowRunApproval) (*api.WorkflowRunApprovalResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/resume/", org, wfGrp, wf, wfRun)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowRunApprovalResponse
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

// ListAllWorkflowRuns retrieves a list of all workflow runs.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllWorkflowRuns(ctx context.Context, org, wf, wfGrp string, request *api.ListAllWorkflowRunsRequest) (*api.GeneratedWorkflowRunListAll, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/listall/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowRunListAll
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:      http.MethodGet,
		Path:        path,
		QueryParams: queryParams,
		Response:    &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}
