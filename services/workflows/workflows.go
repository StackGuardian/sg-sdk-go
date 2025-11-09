package workflows

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Workflows API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Workflows service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateWorkflow creates a new workflow in the Workflow Group.
//
// To create a workflow with a state file:
// 1. Create a workflow using this `Create Workflow` endpoint.
// 2. Use the 'GetSignedUrlToUploadTfstateFile' method to get a signed upload URL for this Workflow.
// 3. Upload the state file to the returned signed URL.
func (s *Service) CreateWorkflow(ctx context.Context, org, wfGrp string, request *api.Workflow) (*api.GeneratedWorkflowCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowCreateResponse
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

// ReadWorkflow retrieves the details of an existing Workflow.
func (s *Service) ReadWorkflow(ctx context.Context, org, wf, wfGrp string) (*api.WorkflowGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGetResponse
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

// DeleteWorkflow deletes an existing workflow.
func (s *Service) DeleteWorkflow(ctx context.Context, org, wf, wfGrp string) (*api.GeneratedWorkflowDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowDeleteResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodDelete,
		Path:     path,
		Response: &response,
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateWorkflow updates an existing workflow's configuration.
func (s *Service) UpdateWorkflow(ctx context.Context, org, wf, wfGrp string, request *api.PatchedWorkflow) (*api.GeneratedWorkflowUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowUpdateResponse
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

// ListAllWorkflowArtifacts retrieves a list of all artifacts for a workflow.
// This List All endpoint does not support pagination at the moment.
func (s *Service) ListAllWorkflowArtifacts(ctx context.Context, org, wf, wfGrp string) (*api.GeneratedWorkflowListAllArtifactsResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/listall_artifacts/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowListAllArtifactsResponse
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

// Outputs retrieves the outputs for a workflow.
func (s *Service) Outputs(ctx context.Context, org, wf, wfGrp string) (*api.GeneratedWorkflowOutputsResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/outputs/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedWorkflowOutputsResponse
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

// GetSignedUrlToUploadTfstateFile returns a signed URL to upload a tfstate file for a workflow.
// The state file can be uploaded by performing a PUT operation on the returned URL.
// This URL is valid for 5 minutes.
func (s *Service) GetSignedUrlToUploadTfstateFile(ctx context.Context, org, wf, wfGrp string, request *api.GetSignedUrlToUploadTfstateFileRequest) (*api.GeneratedWorkflowUploadUrlResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/%v/tfstate_upload_url/", org, wfGrp, wf)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"fileName": request.FileName,
	})

	var response api.GeneratedWorkflowUploadUrlResponse
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

// ListAllWorkflows retrieves a list of all workflows in a workflow group.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllWorkflows(ctx context.Context, org, wfGrp string, request *api.ListAllWorkflowsRequest) (*api.WorkflowsListAll, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfs/listall/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.WorkflowsListAll
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
