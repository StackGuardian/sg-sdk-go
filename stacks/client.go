// This file was auto-generated by Fern from our API Definition.

package stacks

import (
	context "context"
	http "net/http"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	core "github.com/StackGuardian/sg-sdk-go/core"
	internal "github.com/StackGuardian/sg-sdk-go/internal"
	option "github.com/StackGuardian/sg-sdk-go/option"
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

// Creates a new Stack.
func (c *Client) CreateStack(
	ctx context.Context,
	org string,
	wfGrp string,
	request *sgsdkgo.Stack,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackCreateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/",
		org,
		wfGrp,
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
	headers.Set("Content-Type", "application/json")

	var response *sgsdkgo.GeneratedStackCreateResponse
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

// Retrieves details of an existing stack.
func (c *Client) ReadStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackGetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/",
		org,
		wfGrp,
		stack,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *sgsdkgo.GeneratedStackGetResponse
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

// Deletes an existing stack.
func (c *Client) DeleteStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.StackDeleteResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/",
		org,
		wfGrp,
		stack,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *sgsdkgo.StackDeleteResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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

// Updates an existing stack.
func (c *Client) UpdateStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	request *sgsdkgo.PatchedStack,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackCreateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/",
		org,
		wfGrp,
		stack,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *sgsdkgo.GeneratedStackCreateResponse
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

// Read outputs of an existing Stack.
func (c *Client) ReadStackOutputs(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackOutputsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/outputs/",
		org,
		wfGrp,
		stack,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *sgsdkgo.GeneratedStackOutputsResponse
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

// Lists all the Stacks inside a Workflow Group.
func (c *Client) ListAllStacks(
	ctx context.Context,
	org string,
	wfGrp string,
	request *sgsdkgo.ListAllStacksRequest,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackListAllResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/listall/",
		org,
		wfGrp,
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

	var response *sgsdkgo.GeneratedStackListAllResponse
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
