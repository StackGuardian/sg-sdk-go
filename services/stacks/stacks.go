package stacks

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Stacks API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Stacks service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateStack creates a new Stack.
func (s *Service) CreateStack(ctx context.Context, org, wfGrp string, request *api.Stack) (*api.GeneratedStackCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedStackCreateResponse
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

// ReadStack retrieves details of an existing stack.
func (s *Service) ReadStack(ctx context.Context, org, stack, wfGrp string) (*api.GeneratedStackGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/", org, wfGrp, stack)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedStackGetResponse
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

// UpdateStack updates an existing stack.
func (s *Service) UpdateStack(ctx context.Context, org, stack, wfGrp string, request *api.PatchedStack) (*api.GeneratedStackCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/", org, wfGrp, stack)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedStackCreateResponse
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

// DeleteStack deletes an existing stack.
func (s *Service) DeleteStack(ctx context.Context, org, stack, wfGrp string) (*api.StackDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/", org, wfGrp, stack)
	if err != nil {
		return nil, err
	}

	var response api.StackDeleteResponse
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

// ReadStackOutputs reads outputs of an existing Stack.
func (s *Service) ReadStackOutputs(ctx context.Context, org, stack, wfGrp string) (*api.GeneratedStackOutputsResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/outputs/", org, wfGrp, stack)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedStackOutputsResponse
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

// ListAllStacks lists all the Stacks inside a Workflow Group.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllStacks(ctx context.Context, org, wfGrp string, request *api.ListAllStacksRequest) (*api.GeneratedStackListAllResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/listall/", org, wfGrp)
	if err != nil {
		return nil, err
	}

	// Convert request to query params if needed
	// For now, we'll pass it as query parameters
	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.GeneratedStackListAllResponse
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
