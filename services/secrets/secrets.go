package secrets

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Secrets API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Secrets service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// CreateSecret creates a new secret within the organization.
func (s *Service) CreateSecret(ctx context.Context, org string, request *api.Secret) (*api.SecretResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/secrets/", org)
	if err != nil {
		return nil, err
	}

	var response api.SecretResponse
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

// DeleteSecret deletes an existing secret from the organization.
func (s *Service) DeleteSecret(ctx context.Context, org, secret string) (*api.SecretResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/secrets/%v/", org, secret)
	if err != nil {
		return nil, err
	}

	var response api.SecretResponse
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

// UpdateSecret updates an existing secret's configuration or value.
func (s *Service) UpdateSecret(ctx context.Context, org, secret string, request *api.PatchedSecret) (*api.SecretResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/secrets/%v/", org, secret)
	if err != nil {
		return nil, err
	}

	var response api.SecretResponse
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

// ListAllSecrets lists all the secrets in the organization.
// This List All endpoint does not support pagination at the moment.
func (s *Service) ListAllSecrets(ctx context.Context, org string) (*api.SecretListAllResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/secrets/listall/", org)
	if err != nil {
		return nil, err
	}

	var response api.SecretListAllResponse
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
