package workflowsteptemplate

import (
	"context"
	"fmt"
	"net/http"

	core "github.com/StackGuardian/sg-sdk-go/core"
	internal "github.com/StackGuardian/sg-sdk-go/internal"
	option "github.com/StackGuardian/sg-sdk-go/option"
)

const TemplateType = "WORKFLOW_STEP"

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

func (c *Client) ReadWorkflowStepTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	templateId string,
	opts ...option.RequestOption,
) (*ReadWorkflowStepTemplateResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		TemplateType,
		org,
		templateId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	var response *ReadWorkflowStepTemplateResponseModel
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

func (c *Client) CreateWorkflowStepTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	createFirstRevision bool,
	request *CreateWorkflowStepTemplate,
	opts ...option.RequestOption,
) (*CreateWorkflowStepTemplateResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL + "/api/v1/templates/",
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	if !createFirstRevision {
		options.QueryParameters.Set("createFirstRevision", "false")
	}

	var response *CreateWorkflowStepTemplateResponseModel
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

func (c *Client) UpdateWorkflowStepTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	templateId string,
	request *UpdateWorkflowStepTemplateRequestModel,
	opts ...option.RequestOption,
) (*UpdateWorkflowStepTemplateResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		TemplateType,
		org,
		templateId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	var response *UpdateWorkflowStepTemplateResponseModel
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

func (c *Client) DeleteWorkflowStepTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	templateId string,
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
		TemplateType,
		org,
		templateId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

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
