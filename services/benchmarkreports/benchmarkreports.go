package benchmarkreports

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Benchmark Reports API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Benchmark Reports service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// GetBenchmarkReports retrieves benchmark data for your organization.
// You can group, filter, and fetch detailed information about various compliance controls,
// resources, and benchmarks across different cloud service providers.
func (s *Service) GetBenchmarkReports(ctx context.Context, org string, request *api.GetBenchmarkReportsRequest) ([]interface{}, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/reports/benchmark/", org)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"groupBy":              request.GroupBy,
		"filter":               request.Filter,
		"page":                 request.Page,
		"pageSize":             request.PageSize,
		"resourceType":         request.ResourceType,
		"controlId":            request.ControlId,
		"benchmarkName":        request.BenchmarkName,
		"status":               request.Status,
		"integrationType":      request.IntegrationType,
		"integrationName":      request.IntegrationName,
		"integrationGroupName": request.IntegrationGroupName,
		"regions":              request.Regions,
	})

	var response []interface{}
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:      http.MethodGet,
		Path:        path,
		QueryParams: queryParams,
		Response:    &response,
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
