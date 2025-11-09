package runnergroups

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Runner Groups API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Runner Groups service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateNewRunnerGroup creates a new Runner Group in the specified organization.
func (s *Service) CreateNewRunnerGroup(ctx context.Context, org string, request *api.RunnerGroup) (*api.RunnerGroupCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/runnergroups/", org)
	if err != nil {
		return nil, err
	}

	var response api.RunnerGroupCreateResponse
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

// ReadRunnerGroup retrieves details of an existing runner group.
func (s *Service) ReadRunnerGroup(ctx context.Context, org, runnerGroup string, request *api.ReadRunnerGroupRequest) (*api.RunnerGroupSerializerResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/runnergroups/%v/", org, runnerGroup)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.RunnerGroupSerializerResponse
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

// DeleteRunnerGroup deletes an existing Runner Group.
func (s *Service) DeleteRunnerGroup(ctx context.Context, org, runnerGroup string) (*api.RunnerGroupDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/runnergroups/%v/", org, runnerGroup)
	if err != nil {
		return nil, err
	}

	var response api.RunnerGroupDeleteResponse
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

// UpdateRunnerGroup updates an existing Runner Group.
func (s *Service) UpdateRunnerGroup(ctx context.Context, org, runnerGroup string, request *api.PatchedRunnerGroup) (*api.RunnerGroupCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/runnergroups/%v/", org, runnerGroup)
	if err != nil {
		return nil, err
	}

	var response api.RunnerGroupCreateResponse
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

// DeregisterRunner deregisters the runner from the specified runner group.
func (s *Service) DeregisterRunner(ctx context.Context, org, runnerGroup string, request *api.RunnerDeregister) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/runnergroups/%v/deregister/", org, runnerGroup)
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

// UpdateRunnerState updates the state of a runner in the specified runner group.
func (s *Service) UpdateRunnerState(ctx context.Context, org, runnerGroup string, request *api.RunnerStatus) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/runnergroups/%v/runner_status/", org, runnerGroup)
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
