package accessmanagement

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Access Management API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Access Management service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateApiAccess creates a new API access (API Key or OIDC) inside an Organization.
func (s *Service) CreateApiAccess(ctx context.Context, org string, request *api.ApiAccess) (*api.ApiAccessCreateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/apiaccesses/", org)
	if err != nil {
		return nil, err
	}

	var response api.ApiAccessCreateResponse
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

// ReadApiAccess retrieves the details of an existing API access.
func (s *Service) ReadApiAccess(ctx context.Context, accessId, org string) (*api.ApiAccessGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/apiaccesses/%v/", org, accessId)
	if err != nil {
		return nil, err
	}

	var response api.ApiAccessGetResponse
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

// DeleteApiAccess deletes the specified API access from the Organization.
func (s *Service) DeleteApiAccess(ctx context.Context, accessId, org string) (*api.ApiAccessDeleteResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/apiaccesses/%v/", org, accessId)
	if err != nil {
		return nil, err
	}

	var response api.ApiAccessDeleteResponse
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

// UpdateApiAccess updates an existing API access. Note: ResourceName and AccessType cannot be changed.
func (s *Service) UpdateApiAccess(ctx context.Context, accessId, org string, request *api.PatchedApiAccessPatch) (*api.ApiAccessUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/apiaccesses/%v/", org, accessId)
	if err != nil {
		return nil, err
	}

	var response api.ApiAccessUpdateResponse
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

// RegenerateApiKey regenerates the API key for an existing API access.
// Only works for APIKEY access type. A new expiration date must be provided.
func (s *Service) RegenerateApiKey(ctx context.Context, accessId, org string, request *api.ApiAccessRegenerate) (*api.ApiAccessRegenerateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/apiaccesses/%v/regenerate/", org, accessId)
	if err != nil {
		return nil, err
	}

	var response api.ApiAccessRegenerateResponse
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

// ListAllApiAccesses lists all the API accesses inside an Organization. Supports pagination and filtering.
func (s *Service) ListAllApiAccesses(ctx context.Context, org string, request *api.ListAllApiAccessesRequest) (*api.ApiAccessListResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/apiaccesses/listall/", org)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.ApiAccessListResponse
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

// ReadAuditLogs retrieves the audit logs of an Organization.
// To filter logs via query parameters, start and end time must be provided in Unix timestamp format (milliseconds).
func (s *Service) ReadAuditLogs(ctx context.Context, org string, request *api.ReadAuditLogsRequest) (*api.ReadAuditLogResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/audit_logs/", org)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.ReadAuditLogResponse
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

// ReadUser retrieves the details of an user or SSO group within an Organization.
func (s *Service) ReadUser(ctx context.Context, org string, request *api.GetorRemoveUserFromOrganization) (*api.RemoveUserFromOrganizationResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/get_user/", org)
	if err != nil {
		return nil, err
	}

	var response api.RemoveUserFromOrganizationResponse
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

// CreateUser invites users or SSO groups to an Organization.
func (s *Service) CreateUser(ctx context.Context, org string, request *api.AddUserToOrganization) (*api.AddUserToOrganizationResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/invite_user/", org)
	if err != nil {
		return nil, err
	}

	var response api.AddUserToOrganizationResponse
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

// DeleteUser deletes an existing user or SSO group from an Organization.
func (s *Service) DeleteUser(ctx context.Context, org string, request *api.GetorRemoveUserFromOrganization) (*api.RemoveUserFromOrganizationResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/remove_user/", org)
	if err != nil {
		return nil, err
	}

	var response api.RemoveUserFromOrganizationResponse
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

// CreateRole creates a new Role inside an Organization.
func (s *Service) CreateRole(ctx context.Context, org string, request *api.Role) (*api.RoleCreateUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/roles/", org)
	if err != nil {
		return nil, err
	}

	var response api.RoleCreateUpdateResponse
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

// ReadRole retrieves the details of an existing Role.
func (s *Service) ReadRole(ctx context.Context, org, role string) (*api.RoleGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/roles/%v/", org, role)
	if err != nil {
		return nil, err
	}

	var response api.RoleGetResponse
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

// DeleteRole deletes the specified Role from the Organization.
func (s *Service) DeleteRole(ctx context.Context, org, role string) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/roles/%v/", org, role)
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

// UpdateRole updates an existing Role.
func (s *Service) UpdateRole(ctx context.Context, org, role string, request *api.PatchedRole) (*api.RoleCreateUpdateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/roles/%v/", org, role)
	if err != nil {
		return nil, err
	}

	var response api.RoleCreateUpdateResponse
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

// ListAllRoles lists all the Roles inside an Organization.
// This List All endpoint does not support pagination at the moment.
func (s *Service) ListAllRoles(ctx context.Context, org string, request *api.ListAllRolesRequest) error {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/roles/listall/", org)
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

// UpdateUser updates an existing user or SSO group within an Organization.
func (s *Service) UpdateUser(ctx context.Context, org string, request *api.AddUserToOrganization) (*api.AddUserToOrganizationResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/update_user/", org)
	if err != nil {
		return nil, err
	}

	var response api.AddUserToOrganizationResponse
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

// ListAllUsers lists all users within an Organization.
func (s *Service) ListAllUsers(ctx context.Context, org string, request *api.ListAllUsersRequest) (*api.ListAllUsersInOrganizationResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/users/listall/", org)
	if err != nil {
		return nil, err
	}

	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}

	var response api.ListAllUsersInOrganizationResponse
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
