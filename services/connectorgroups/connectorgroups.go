package connectorgroups

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Connector Groups API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Connector Groups service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateConnectorGroup creates a new Connector Group.
// A connector group can be created with multiple child connectors by including them in the `childIntegrations` field.
func (s *Service) CreateConnectorGroup(ctx context.Context, org string, request *api.IntegrationGroups) (*api.IntegrationGroupsCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/", org)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationGroupsCreateResponse
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

// ReadConnectorGroup reads the attributes of a specific Connector Group.
func (s *Service) ReadConnectorGroup(ctx context.Context, integrationgroup, org string) (*api.IntegrationGroupGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/", org, integrationgroup)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationGroupGetResponse
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

// DeleteConnectorGroup deletes a specific Connector Group.
func (s *Service) DeleteConnectorGroup(ctx context.Context, integrationgroup, org string) (*api.IntegrationGroupsDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/", org, integrationgroup)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationGroupsDeleteResponse
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

// UpdateConnectorGroup updates the attributes of a specific Connector Group.
// Use this endpoint to add a new child connector to a connector group by including it in the `childIntegrations` field.
func (s *Service) UpdateConnectorGroup(ctx context.Context, integrationgroup, org string, request *api.PatchedIntegrationGroups) (*api.IntegrationGroupsCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/", org, integrationgroup)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationGroupsCreateResponse
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

// AuthenticateConnectorGroup authenticates the attributes of a specific Connector Group.
func (s *Service) AuthenticateConnectorGroup(ctx context.Context, integrationgroup, org string) (*api.IntegrationGroupsAuthenticationResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/authenticate/", org, integrationgroup)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationGroupsAuthenticationResponse
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

// ReadChildInACloudConnectorGroup reads an existing Child Connector in a Cloud Connector Group.
func (s *Service) ReadChildInACloudConnectorGroup(ctx context.Context, integration, integrationgroup, org string) (*api.IntegrationGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/integrations/%v/", org, integrationgroup, integration)
	if err != nil {
		return nil, err
	}

	var response api.IntegrationGetResponse
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

// DeleteChildInACloudConnectorGroup deletes a specific Child Connector in a Cloud Connector Group.
func (s *Service) DeleteChildInACloudConnectorGroup(ctx context.Context, integration, integrationgroup, org string) (*api.GeneratedCloudConnectorGroupDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/integrations/%v/", org, integrationgroup, integration)
	if err != nil {
		return nil, err
	}

	var response api.GeneratedCloudConnectorGroupDeleteResponse
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

// UpdateChildInACloudConnectorGroup updates a specific Child Connector in a Cloud Connector Group.
// To add a new child connector to a connector group, use the `UpdateConnectorGroup` method.
func (s *Service) UpdateChildInACloudConnectorGroup(ctx context.Context, integration, integrationgroup, org string, request *api.PatchedIntegration) (*api.IntegrationUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/integrations/%v/", org, integrationgroup, integration)
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

// ListAllConnectorsInAGroup lists all the Connectors within a specified Connector Group.
func (s *Service) ListAllConnectorsInAGroup(ctx context.Context, integrationgroup, org string, request *api.ListAllConnectorsInAGroupRequest) (*api.IntegrationGroupsListAllIntegrations, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/%v/integrations/listall/", org, integrationgroup)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.IntegrationGroupsListAllIntegrations
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

// ListAllConnectorGroups lists all the Connector Groups in an Organization.
func (s *Service) ListAllConnectorGroups(ctx context.Context, org string, request *api.ListAllConnectorGroupsRequest) (*api.IntegrationGroupsListAllResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/integrationgroups/listall/", org)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.IntegrationGroupsListAllResponse
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
