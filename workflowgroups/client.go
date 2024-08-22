// This file was auto-generated by Fern from our API Definition.

package workflowgroups

import (
	context "context"
	http "net/http"
	"strings"

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

// Create Workflow Group in an Organization
func (c *Client) CreateWorkflowGroup(
	ctx context.Context,
	org string,
	request *sgsdkgo.WorkflowGroup,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupCreateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/api/v1/orgs/%v/wfgrps/", org)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupCreateResponse
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

// Read Workflow Group
func (c *Client) ReadWorkflowGroup(
	ctx context.Context,
	org string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupGetResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}

	//If wfGrp contains "/" then it might be a nested workflow group
	//In this case we need to pass through the / without encoding it
	var endpointURL string
	if strings.Contains(wfGrp, "/") {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/",
			org,
		)
		endpointURL += wfGrp
	} else {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/%v",
			org,
			wfGrp,
		)
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupGetResponse
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

// Delete Workflow Group
func (c *Client) DeleteWorkflowGroup(
	ctx context.Context,
	org string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupDeleteResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}

	//If wfGrp contains "/" then it might be a nested workflow group
	//In this case we need to pass through the / without encoding it
	var endpointURL string
	if strings.Contains(wfGrp, "/") {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/",
			org,
		)
		endpointURL += wfGrp
	} else {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/%v",
			org,
			wfGrp,
		)
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupDeleteResponse
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

// Update Workflow Group
func (c *Client) UpdateWorkflowGroup(
	ctx context.Context,
	org string,
	wfGrp string,
	request *sgsdkgo.PatchedWorkflowGroup,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupPatch, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}

	//If wfGrp contains "/" then it might be a nested workflow group
	//In this case we need to pass through the / without encoding it
	var endpointURL string
	if strings.Contains(wfGrp, "/") {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/",
			org,
		)
		endpointURL += wfGrp
	} else {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/%v",
			org,
			wfGrp,
		)
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupPatch
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

// Create Child Workflow Group
func (c *Client) CreateChildWorkflowGroup(
	ctx context.Context,
	org string,
	wfGrp string,
	request *sgsdkgo.WorkflowGroup,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupCreateResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}

	//If wfGrp contains "/" then it might be a nested workflow group
	//In this case we need to pass through the / without encoding it
	var endpointURL string
	if strings.Contains(wfGrp, "/") {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/",
			org,
		)
		endpointURL += wfGrp + "/wfgrps/"
	} else {
		endpointURL = core.EncodeURL(
			baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfgrps/",
			org,
			wfGrp,
		)
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupCreateResponse
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

// List all Child Workflow Groups
func (c *Client) ListAllChildWorkflowGroups(
	ctx context.Context,
	org string,
	wfGrp string,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupListAllResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(
		baseURL+"/api/v1/orgs/%v/wfgrps/%v/wfgrps/listall/",
		org,
		wfGrp,
	)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupListAllResponse
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

// List all Workflow Groups
func (c *Client) ListAllWorkflowGroups(
	ctx context.Context,
	org string,
	opts ...option.RequestOption,
) (*sgsdkgo.WorkflowGroupListAllResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.app.stackguardian.io"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/api/v1/orgs/%v/wfgrps/listall/", org)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *sgsdkgo.WorkflowGroupListAllResponse
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