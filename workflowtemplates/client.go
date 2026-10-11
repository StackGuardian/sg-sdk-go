package workflowtemplates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/internal"
	"github.com/StackGuardian/sg-sdk-go/option"
)

const TemplateType = "IAC"

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

func (c *Client) ReadWorkflowTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	templateId string,
	opts ...option.RequestOption,
) (*ReadWorkflowTemplateResponseModel, error) {
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

	var response *ReadWorkflowTemplateResponseModel
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

func (c *Client) CreateWorkflowTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	createFirstRevision bool,
	request *CreateWorkflowTemplateRequest,
	opts ...option.RequestOption,
) (*CreateWorkflowTemplateResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL + "/api/v1/templates/",
	)

	request.TemplateType = "IAC"

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	if !createFirstRevision {
		options.QueryParameters.Set("createFirstRevision", "false")
	}

	var response *CreateWorkflowTemplateResponseModel
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

func (c *Client) UpdateWorkflowTemplate(
	ctx context.Context,
	org string, // org of the user making the request
	templateId string,
	request *UpdateWorkflowTemplateRequest,
	opts ...option.RequestOption,
) (*UpdateWorkflowTemplateResponseModel, error) {
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

	var response *UpdateWorkflowTemplateResponseModel
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

func (c *Client) DeleteWorkflowTemplate(
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
