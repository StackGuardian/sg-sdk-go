// This file was auto-generated by Fern from our API Definition.

package stackworkflows

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

// Read Workflow (Stack)
func (c *Client) ReadStackWorkflow(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGetResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v",
		org,
		wfGrp,
		stack,
		wf,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGetResponse
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

// Delete Workflow (Stack)
func (c *Client) DeleteStackWorkflow(
	ctx context.Context,
	org string,
	stack string,
	wf string,
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
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v",
		org,
		wfGrp,
		stack,
		wf,
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

// Update Workflow (Stack)
func (c *Client) UpdateStackWorkflow(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	request *sgsdkgo.PatchedWorkflow,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowUpdateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v",
		org,
		wfGrp,
		stack,
		wf,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowUpdateResponse
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

// List all artifacts (Stack)
func (c *Client) ListAllStackWorkflowsArtifacts(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowListAllArtifactsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/listall_artifacts/",
		org,
		wfGrp,
		stack,
		wf,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowListAllArtifactsResponse
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

// Workflow Outputs (Stack)
func (c *Client) StackWorkflowOutputs(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowOutputsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/outputs/",
		org,
		wfGrp,
		stack,
		wf,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowOutputsResponse
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

// List all Workflows (Stack)
func (c *Client) ListAllStackWorkflows(
	ctx context.Context,
	org string,
	stack string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowsListAll, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/listall/",
		org,
		wfGrp,
		stack,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowsListAll
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
