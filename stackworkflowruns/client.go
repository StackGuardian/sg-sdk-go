// This file was auto-generated by Fern from our API Definition.

package stackworkflowruns

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

// Run Stack
func (c *Client) CreateStackRun(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	request *sgsdkgo.StackAction,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackRunsResponse, error) {
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

	var response *sgsdkgo.GeneratedStackRunsResponse
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

// Read all workflow runs of a stack run
func (c *Client) ReadStackRun(
	ctx context.Context,
	org string,
	stack string,
	stackRun string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackRunsGetResponse, error) {
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

	var response *sgsdkgo.GeneratedStackRunsGetResponse
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

// List all stack runs of a stack
func (c *Client) ListAllStackRuns(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedStackRunsListAllResponse, error) {
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

	var response *sgsdkgo.GeneratedStackRunsListAllResponse
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

// Read a single workflow run inside a stack's workflow
func (c *Client) ReadStackWorkflowRun(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	wfRun string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowRunStackGet, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/",
		org,
		wfGrp,
		stack,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowRunStackGet
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

// Read execution logs for a single workflow run inside a stack's workflow
func (c *Client) ReadStackWorkflowRunLogs(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	wfRun string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowRunLogs, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/logs/",
		org,
		wfGrp,
		stack,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowRunLogs
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

// Provide approval for a workflow run inside a stack's workflow
func (c *Client) ApproveStackWorkflowRun(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	wfRun string,
	request *sgsdkgo.WorkflowRunApproval,
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
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/resume/",
		org,
		wfGrp,
		stack,
		wf,
		wfRun,
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