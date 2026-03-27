package organizations

import (
	"context"
	"net/http"

	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
)

// Service provides access to the Organizations API.
type Service struct {
	client *internal.HTTPClient
}

// NewService creates a new Organizations service.
func NewServiceWithHTTPClient(httpClient *internal.HTTPClient) *Service {
	return &Service{
		client: httpClient,
	}
}

// ReadOrganization retrieves the details of an Organization.
func (s *Service) ReadOrganization(ctx context.Context, org string) (*api.OrgGetResponse, error) {
	path, err := internal.BuildURL("", "/api/v1/orgs/%v/", org)
	if err != nil {
		return nil, err
	}

	var response api.OrgGetResponse
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
