package stackworkflows

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Stack Workflows API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Stack Workflows service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// ReadStackWorkflow retrieves the details of an existing Workflow in a Stack.
func (s *Service) ReadStackWorkflow(ctx context.Context, org, stack, wf, wfGrp string) (*api.WorkflowGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/", org, wfGrp, stack, wf)
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

// DeleteStackWorkflow deletes an existing workflow in a stack.
func (s *Service) DeleteStackWorkflow(ctx context.Context, org, stack, wf, wfGrp string) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/", org, wfGrp, stack, wf)
	if err != nil {
		return err
	}

	err = s.client.Do(ctx, &internal.RequestOptions{
		Method: http.MethodDelete,
		Path:   path,
	})
	if err != nil {
		return err
	}

	return nil
}

// UpdateStackWorkflow updates an existing workflow in a stack.
func (s *Service) UpdateStackWorkflow(ctx context.Context, org, stack, wf, wfGrp string, request *api.PatchedWorkflow) (*api.GeneratedWorkflowUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/", org, wfGrp, stack, wf)
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

// ListAllStackWorkflowsArtifacts retrieves a list of all artifacts for a workflow in a stack.
// This List All endpoint does not support pagination at the moment.
func (s *Service) ListAllStackWorkflowsArtifacts(ctx context.Context, org, stack, wf, wfGrp string) (*api.GeneratedWorkflowListAllArtifactsResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/listall_artifacts/", org, wfGrp, stack, wf)
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

// StackWorkflowOutputs retrieves the outputs for a workflow in a stack.
func (s *Service) StackWorkflowOutputs(ctx context.Context, org, stack, wf, wfGrp string) (*api.GeneratedWorkflowOutputsResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/outputs/", org, wfGrp, stack, wf)
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

// GetSignedUrlToUploadTfstateFileForStackWorkflow returns a signed URL to upload a tfstate file for a Stack Workflow.
// The state file can be uploaded by performing a PUT operation on the returned URL.
// The URL is valid for 5 minutes.
func (s *Service) GetSignedUrlToUploadTfstateFileForStackWorkflow(ctx context.Context, org, stack, wf, wfGrp string, request *api.GetSignedUrlToUploadTfstateFileForStackWorkflowRequest) (*api.GeneratedWorkflowUploadUrlResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/tfstate_upload_url/", org, wfGrp, stack, wf)
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

// ListAllStackWorkflows retrieves a list of all workflows in a Stack.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllStackWorkflows(ctx context.Context, org, stack, wfGrp string, request *api.ListAllStackWorkflowsRequest) (*api.WorkflowsListAll, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/listall/", org, wfGrp, stack)
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
