package templates

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Templates API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Templates service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// ReadSubscription reads all the templates that are subscribed by an organization.
func (s *Service) ReadSubscription(ctx context.Context, org string, request *api.ReadSubscriptionRequest) (*api.GetSubscriptionResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/subscriptions/default/", org)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.GetSubscriptionResponse
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

// ListAllTemplatesBasedOnOwnerOrg lists all Templates and its revisions created by the Organization.
func (s *Service) ListAllTemplatesBasedOnOwnerOrg(ctx context.Context, org string, request *api.ListAllTemplatesBasedOnOwnerOrgRequest) (*api.ListallTemplatesResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/templates/listall/", org)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.ListallTemplatesResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:      http.MethodGet,
		Path:        path,
		QueryParams: queryParams,
		Response:    &response,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateTemplateRevision creates a new revision of a template or creates the initial template if it doesn't exist.
func (s *Service) CreateTemplateRevision(ctx context.Context, request *api.CreateTemplateRevisionRequest) (*api.CreateTemplateResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/templates/")
	if err != nil {
		return nil, err
	}

	var response api.CreateTemplateResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodPost,
		Path:     path,
		Body:     request,
		Response: &response,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReadTemplateRevision retrieves a specific template revision or the base template.
// Use format `template-name:revision` to fetch a specific revision (e.g., `my-template:5`).
// If revision is not passed (e.g., `my-template`), it will fetch the base template details.
func (s *Service) ReadTemplateRevision(ctx context.Context, org, templateRevision string, templateType api.ReadTemplateRevisionRequestTemplateType, request *api.ReadTemplateRevisionRequest) (*api.TemplateGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/templatetypes/%v/%v/%v/", templateType, org, templateRevision)
	if err != nil {
		return nil, err
	}

	var response api.TemplateGetResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodGet,
		Path:     path,
		Response: &response,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteTemplateRevision deletes a specific template revision.
// If all revisions are deleted, the parent template is also removed.
func (s *Service) DeleteTemplateRevision(ctx context.Context, org, templateRevision, templateType string, request *api.DeleteTemplateRevisionRequest) error {
	path, err := internal.BuildURL("", "/api/v1/templatetypes/%v/%v/%v/", templateType, org, templateRevision)
	if err != nil {
		return err
	}

	err = s.client.Do(ctx, &internal.RequestOptions{
		Method: http.MethodDelete,
		Path:   path,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// UpdateTemplateRevision updates an existing parent template or its revision with new configuration.
func (s *Service) UpdateTemplateRevision(ctx context.Context, org, templateRevision, templateType string, request *api.PatchedTemplateUpdate) (*api.TemplateCreatePatchResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/templatetypes/%v/%v/%v/", templateType, org, templateRevision)
	if err != nil {
		return nil, err
	}

	var response api.TemplateCreatePatchResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodPatch,
		Path:     path,
		Body:     request,
		Response: &response,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ListAllTemplates lists all Templates and its revisions created or subscribed by the Organization.
// Supports Pagination and Filtering using query parameters.
func (s *Service) ListAllTemplates(ctx context.Context, templateType api.ListAllTemplatesRequestTemplateType, request *api.ListAllTemplatesRequest) (*api.ListallTemplatesResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/templatetypes/%v/templates/listall/", templateType)
	if err != nil {
		return nil, err
	}

	queryParams := internal.EncodeQueryParams(map[string]interface{}{
		"filter":   request.Filter,
		"page":     request.Page,
		"pageSize": request.PageSize,
	})

	var response api.ListallTemplatesResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:      http.MethodGet,
		Path:        path,
		QueryParams: queryParams,
		Response:    &response,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReadIacGroupsIacTemplate retrieves an IAC Group's IAC Template configuration.
func (s *Service) ReadIacGroupsIacTemplate(ctx context.Context, org, subTemplateId, template string, request *api.ReadIacGroupsIacTemplateRequest) (*api.TemplateGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/templatetypes/IAC_GROUP/%v/%v/IAC/%v/", org, template, subTemplateId)
	if err != nil {
		return nil, err
	}

	var response api.TemplateGetResponse
	err = s.client.Do(ctx, &internal.RequestOptions{
		Method:   http.MethodGet,
		Path:     path,
		Response: &response,
		Headers: map[string]string{
			"x-sg-orgid": request.SgOrgid,
		},
	})
	if err != nil {
		return nil, err
	}

	return &response, nil
}
