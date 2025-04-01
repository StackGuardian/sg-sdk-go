// This file was auto-generated by Fern from our API Definition.

package stackworkflowrunfacts

import (
	context "context"
	http "net/http"

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

// Get the workflow run facts of a Stack workflow.
//
// This endpoint returns a signed URL which can be used to get the full contents of the Stack Workflow run facts. This signed URL is valid for 60 minutes. After expiration, you can request a new signed URL by calling this endpoint again.
//
// For more information, please refer to [this discussion](https://github.com/StackGuardian/feedback/discussions/109).
func (c *Client) GetStackWorkflowRunFacts(
	ctx context.Context,
	org string,
	stack string,
	wf string,
	wfGrp string,
	wfRun string,
	wfRunFacts string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/stacks/%v/wfs/%v/wfruns/%v/wfrunfacts/%v/",
		org,
		wfGrp,
		stack,
		wf,
		wfRun,
		wfRunFacts,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

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
		},
	); err != nil {
		return err
	}
	return nil
}
