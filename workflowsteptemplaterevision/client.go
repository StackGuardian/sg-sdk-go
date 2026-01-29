package workflowsteptemplaterevision

import (
	"context"
	"fmt"
	"net/http"

	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/internal"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/workflowsteptemplate"
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

func (c *Client) CreateWorkflowStepTemplateRevision(
	ctx context.Context,
	org string,
	templateId string,
	request *CreateWorkflowStepTemplateRevisionModel,
	opts ...option.RequestOption,
) (*CreateWorkflowStepTemplateRevisionResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/WORKFLOW_STEP/%v/%v/revisions/",
		org,
		templateId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	headers.Add("x-sg-orgid", org)

	var response *CreateWorkflowStepTemplateRevisionResponseModel
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

func (c *Client) UpdateWorkflowStepTemplateRevision(
	ctx context.Context,
	org string,
	revisionId string,
	request *UpdateWorkflowStepTemplateRevisionModel,
	opts ...option.RequestOption,
) (*UpdateWorkflowStepTemplateRevisionResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		workflowsteptemplate.TemplateType,
		org,
		revisionId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	var response *UpdateWorkflowStepTemplateRevisionResponseModel
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

func (c *Client) ReadWorkflowStepTemplateRevision(
	ctx context.Context,
	org string,
	revisionId string,
	opts ...option.RequestOption,
) (*ReadWorkflowStepTemplateRevisionResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		workflowsteptemplate.TemplateType,
		org,
		revisionId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	var response *ReadWorkflowStepTemplateRevisionResponseModel
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteWorkflowStepTemplateRevision(
	ctx context.Context,
	org string,
	revisionId string,
	keepParentTemplate bool,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		workflowsteptemplate.TemplateType,
		org,
		revisionId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	if keepParentTemplate {
		options.QueryParameters.Set("keepParentTemplate", "true")
	}

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
		},
	); err != nil {
		return err
	}
	return nil
}
