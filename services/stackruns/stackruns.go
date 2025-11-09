package stackruns

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Stack Runs API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Stack Runs service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateStackRun initiates a new run of an existing stack.
func (s *Service) CreateStackRun(ctx context.Context, org, stack, wfGrp string, request *api.StackAction) (*api.GeneratedStackRunsResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/stackruns/", org, wfGrp, stack)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedStackRunsResponse
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

// ReadStackRun retrieves details of all workflow runs within a specific stack run.
func (s *Service) ReadStackRun(ctx context.Context, org, stack, stackRun, wfGrp string) (*api.GeneratedStackRunsGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/stackruns/%v/", org, wfGrp, stack, stackRun)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedStackRunsGetResponse
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

// ListAllStackRuns retrieves a list of all stack runs for a specific stack.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllStackRuns(ctx context.Context, org, stack, wfGrp string, request *api.ListAllStackRunsRequest) (*api.GeneratedStackRunsListAllResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/stacks/%v/stackruns/listall/", org, wfGrp, stack)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.GeneratedStackRunsListAllResponse
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
