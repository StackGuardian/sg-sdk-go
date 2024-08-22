// This file was auto-generated by Fern from our API Definition.

package connectors

import (
	context "context"
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	core "github.com/StackGuardian/sg-sdk-go/core"
	option "github.com/StackGuardian/sg-sdk-go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Create Connector inside an Organization
func (c *Client) CreateConnector(
	ctx context.Context,
	org string,
	request *sgsdkgo.Integration,
	opts ...option.RequestOption,
) (*sgsdkgo.IntegrationCreateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/api/v1/orgs/%v/integrations/", org)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.IntegrationCreateResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Read Connector
func (c *Client) ReadConnector(
	ctx context.Context,
	integration string,
	org string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedConnectorReadResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/integrations/%v/",
		org,
		integration,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedConnectorReadResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete Connector
func (c *Client) DeleteConnector(
	ctx context.Context,
	integration string,
	org string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedConnectorDeleteResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/integrations/%v/",
		org,
		integration,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedConnectorDeleteResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodDelete,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update Connector
func (c *Client) UpdateConnector(
	ctx context.Context,
	integration string,
	org string,
	request *sgsdkgo.PatchedIntegration,
	opts ...option.RequestOption,
) (*sgsdkgo.IntegrationUpdateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/integrations/%v/",
		org,
		integration,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.IntegrationUpdateResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPatch,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// List all Connector
func (c *Client) ListAllConnector(
	ctx context.Context,
	org string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedConnectorListAllResponseMsg, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/api/v1/orgs/%v/integrations/listall/", org)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedConnectorListAllResponseMsg
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}