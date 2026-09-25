package statebackends

import (
	"context"
	"net/http"

	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/internal"
	"github.com/StackGuardian/sg-sdk-go/option"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Create a state backend in an organization.
func (c *Client) CreateStateBackend(
	ctx context.Context,
	org string,
	request *StateBackend,
	opts ...option.RequestOption,
) (*StateBackendResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/statebackends/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *StateBackendResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Read a state backend.
func (c *Client) ReadStateBackend(
	ctx context.Context,
	org string,
	stateBackend string,
	opts ...option.RequestOption,
) (*StateBackendResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/statebackends/%v/",
		org,
		stateBackend,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *StateBackendResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update a state backend.
func (c *Client) UpdateStateBackend(
	ctx context.Context,
	org string,
	stateBackend string,
	request *PatchedStateBackend,
	opts ...option.RequestOption,
) (*StateBackendResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/statebackends/%v/",
		org,
		stateBackend,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *StateBackendResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete a state backend.
func (c *Client) DeleteStateBackend(
	ctx context.Context,
	org string,
	stateBackend string,
	opts ...option.RequestOption,
) (*DeleteStateBackendResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/statebackends/%v/",
		org,
		stateBackend,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *DeleteStateBackendResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:                endpointURL,
			Method:             http.MethodDelete,
			Headers:            headers,
			MaxAttempts:        options.MaxAttempts,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// List all state backends in an organization.
func (c *Client) ListAllStateBackends(
	ctx context.Context,
	org string,
	request *ListAllStateBackendsRequest,
	opts ...option.RequestOption,
) (*StateBackendListAllResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/statebackends/listall/",
		org,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *StateBackendListAllResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:                endpointURL,
			Method:             http.MethodGet,
			Headers:            headers,
			MaxAttempts:        options.MaxAttempts,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// List the .json and .tfstate files in a storage backend described by a base64-encoded StateBackendConfig.
func (c *Client) ListStatefiles(
	ctx context.Context,
	org string,
	request *ListStatefilesRequest,
	opts ...option.RequestOption,
) (*ListStatefilesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/statebackends/list_statefiles/",
		org,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *ListStatefilesResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:                endpointURL,
			Method:             http.MethodGet,
			Headers:            headers,
			MaxAttempts:        options.MaxAttempts,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
