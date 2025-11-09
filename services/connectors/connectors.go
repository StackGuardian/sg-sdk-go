package connectors

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Connectors API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Connectors service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateConnector creates a new Connector inside an Organization.
func (s *Service) CreateConnector(ctx context.Context, org string, request *api.Integration) (*api.IntegrationCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrations/", org)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationCreateResponse
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

// ReadConnector reads an existing Connector inside an Organization.
func (s *Service) ReadConnector(ctx context.Context, integration, org string) (*api.GeneratedConnectorReadResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrations/%v/", org, integration)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedConnectorReadResponse
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

// DeleteConnector deletes an existing Connector inside an Organization.
func (s *Service) DeleteConnector(ctx context.Context, integration, org string) (*api.GeneratedConnectorDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrations/%v/", org, integration)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedConnectorDeleteResponse
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

// UpdateConnector updates an existing Connector inside an Organization.
func (s *Service) UpdateConnector(ctx context.Context, integration, org string, request *api.PatchedIntegration) (*api.IntegrationUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrations/%v/", org, integration)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationUpdateResponse
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

// ListAllConnectors lists all Connectors inside an Organization.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllConnectors(ctx context.Context, org string, request *api.ListAllConnectorsRequest) (*api.GeneratedConnectorListAllResponseMsg, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrations/listall/", org)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.GeneratedConnectorListAllResponseMsg
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
