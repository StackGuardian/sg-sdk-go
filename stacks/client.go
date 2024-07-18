// This file was auto-generated by Fern from our API Definition.

package stacks

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

// Creates a Stack
func (c *Client) Create(
	ctx context.Context,
	org string,
	wfGrp string,
	request *sgsdkgo.StacksCreateRequest,
	opts ...option.RequestOption,
) (*sgsdkgo.StackCreatePatchResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/",
		org,
		wfGrp,
	)

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.StackCreatePatchResponse
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

// Get workflow details inside a workflow group
func (c *Client) ReadStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// Delete a Stack
func (c *Client) DeleteStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodDelete,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// Update workflow attributes
func (c *Client) UpdateStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	request *sgsdkgo.PatchedStack,
	opts ...option.RequestOption,
) (*sgsdkgo.StackCreatePatchResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.StackCreatePatchResponse
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

// Get outputs for a Stack
func (c *Client) GetStackOutputs(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/outputs/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// Run Stack
func (c *Client) RunStack(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	request *sgsdkgo.StackAction,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/stackruns/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
		},
	); err != nil {
		return err
	}
	return nil
}

// List all Workflow Runs in a Stack
func (c *Client) GetStackRuns(
	ctx context.Context,
	org string,
	stack string,
	stackRun string,
	wfGrp string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/stackruns/%v",
		org,
		wfGrp,
		stack,
		stackRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// List all Workflow Runs in a Stack
func (c *Client) ListAllStackRuns(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/stackruns/listall/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// List all Stacks in a Workflow Group
func (c *Client) ListAllStacks(
	ctx context.Context,
	org string,
	wfGrp string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/listall/",
		org,
		wfGrp,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}
