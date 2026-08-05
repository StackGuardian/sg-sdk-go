package policies

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Policies API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Policies service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreatePolicy creates a new Policy inside an Organization.
//
// To create a new Insight Filter Policy, please have a look at https://github.com/StackGuardian/feedback/discussions/147
func (s *Service) CreatePolicy(ctx context.Context, org string, request *api.PolymorphicPolicy) (*api.PolicyCreateUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/policies/", org)
	if err != nil {
		return nil, err
	}

	var response api.PolicyCreateUpdateResponse
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

// ReadPolicy reads an existing policy's details.
func (s *Service) ReadPolicy(ctx context.Context, org, policy string) (*api.PolicyGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/policies/%v/", org, policy)
	if err != nil {
		return nil, err
	}

	var response api.PolicyGetResponse
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

// DeletePolicy permanently removes an existing policy from the Organization.
func (s *Service) DeletePolicy(ctx context.Context, org, policy string) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/policies/%v/", org, policy)
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

// UpdatePolicy updates an existing policy's details.
func (s *Service) UpdatePolicy(ctx context.Context, org, policy string, request *api.PatchedPolymorphicPolicy) (*api.PolicyCreateUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/policies/%v/", org, policy)
	if err != nil {
		return nil, err
	}

	var response api.PolicyCreateUpdateResponse
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

// ListAllPolicies lists all the policies inside an Organization.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllPolicies(ctx context.Context, org string, request *api.ListAllPoliciesRequest) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/policies/listall/", org)
	if err != nil {
		return err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return err
	}

	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:      http.MethodGet,
		Path:        path,
		QueryParams: queryParams,
	})
	if err != nil {
		return err
	}

	return nil
}
