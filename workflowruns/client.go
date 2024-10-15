// This file was auto-generated by Fern from our API Definition.

package workflowruns

import (
	context "context"
	http "net/http"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	core "github.com/StackGuardian/sg-sdk-go/core"
	option "github.com/StackGuardian/sg-sdk-go/option"
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

// Run workflow
func (c *Client) CreateWorkflowRun(
	ctx context.Context,
	org string,
	wf string,
	wfGrp string,
	request *sgsdkgo.WorkflowRun,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowRunCreatePatchResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/",
		org,
		wfGrp,
		wf,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowRunCreatePatchResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Read workflow run details
func (c *Client) ReadWorkflowRun(
	ctx context.Context,
	org string,
	wf string,
	wfGrp string,
	wfRun string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowRunsGet, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/",
		org,
		wfGrp,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowRunsGet
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Patch Workflow Run inside a workflow
func (c *Client) UpdateWorkflowRun(
	ctx context.Context,
	org string,
	wf string,
	wfGrp string,
	wfRun string,
	request *sgsdkgo.PatchedWorkflowRun,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkfkowRunsUpdateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/",
		org,
		wfGrp,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkfkowRunsUpdateResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Patch Workflow Run inside a workflow
func (c *Client) CancelWorkflowRun(
	ctx context.Context,
	org string,
	wf string,
	wfGrp string,
	wfRun string,
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
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/cancel/",
		org,
		wfGrp,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}

// Read execution logs for a Workflow Run inside a workflow
func (c *Client) ReadWorkflowRunLogs(
	ctx context.Context,
	org string,
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
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/logs/",
		org,
		wfGrp,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowRunLogs
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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

// Provide approval for a Workflow Run
func (c *Client) ApproveWorkflowRun(
	ctx context.Context,
	org string,
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
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/%v/resume/",
		org,
		wfGrp,
		wf,
		wfRun,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
		},
	); err != nil {
		return err
	}
	return nil
}

// List all Workflow Runs in a Workflow
func (c *Client) ListAllWorkflowRuns(
	ctx context.Context,
	org string,
	wf string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.GeneratedWorkflowRunListAll, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfs/%v/wfruns/listall/",
		org,
		wfGrp,
		wf,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.GeneratedWorkflowRunListAll
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
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
