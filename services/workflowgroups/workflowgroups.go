package workflowgroups

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Workflow Groups API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Workflow Groups service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateWorkflowGroup creates a new Workflow Group.
func (s *Service) CreateWorkflowGroup(ctx context.Context, org string, request *api.WorkflowGroup) (*api.WorkflowGroupCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/", org)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupCreateResponse
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

// ReadWorkflowGroup reads an existing Workflow Group.
// Note: wfGrp can be a nested path like "parent/child" for nested workflow groups.
func (s *Service) ReadWorkflowGroup(ctx context.Context, org, wfGrp string) (*api.WorkflowGroupGetResponse, error) {
	// BuildURL handles nested workflow groups (containing "/") automatically
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupGetResponse
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

// UpdateWorkflowGroup updates an existing Workflow Group.
// Note: wfGrp can be a nested path like "parent/child" for nested workflow groups.
func (s *Service) UpdateWorkflowGroup(ctx context.Context, org, wfGrp string, request *api.PatchedWorkflowGroup) (*api.WorkflowGroupPatch, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupPatch
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

// DeleteWorkflowGroup deletes an existing Workflow Group.
// Note: wfGrp can be a nested path like "parent/child" for nested workflow groups.
func (s *Service) DeleteWorkflowGroup(ctx context.Context, org, wfGrp string) (*api.WorkflowGroupDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupDeleteResponse
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

// CreateChildWorkflowGroup creates a new Child Workflow Group within an existing Workflow Group.
// Note: wfGrp can be a nested path like "parent/child" for nested workflow groups.
func (s *Service) CreateChildWorkflowGroup(ctx context.Context, org, wfGrp string, request *api.WorkflowGroup) (*api.WorkflowGroupCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfgrps/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupCreateResponse
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

// ListAllChildWorkflowGroups lists all Child Workflow Groups in an existing Workflow Group.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllChildWorkflowGroups(ctx context.Context, org, wfGrp string, request *api.ListAllChildWorkflowGroupsRequest) (*api.WorkflowGroupListAllResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/wfgrps/listall/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupListAllResponse
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

// ListAllWorkflowGroups lists all Workflow Groups in an Organization.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllWorkflowGroups(ctx context.Context, org string, request *api.ListAllWorkflowGroupsRequest) (*api.WorkflowGroupListAllResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/listall/", org)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.WorkflowGroupListAllResponse
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
