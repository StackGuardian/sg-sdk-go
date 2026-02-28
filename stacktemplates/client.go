package stacktemplates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/StackGuardian/sg-sdk-go/core"
	"github.com/StackGuardian/sg-sdk-go/internal"
	"github.com/StackGuardian/sg-sdk-go/option"
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

func (c *Client) ReadStackTemplate(
	ctx context.Context,
	org string,
	templateId string,
	opts ...option.RequestOption,
) (*ReadStackTemplateResponseModel, error) {
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

	var response *ReadStackTemplateResponseModel
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

func (c *Client) CreateStackTemplate(
	ctx context.Context,
	org string,
	createFirstRevision bool,
	request *CreateStackTemplateRequest,
	opts ...option.RequestOption,
) (*CreateStackTemplateResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL + "/api/v1/templates/",
	)

	request.TemplateType = "IAC_GROUP"

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	if !createFirstRevision {
		options.QueryParameters.Set("createFirstRevision", "false")
	}

	var response *CreateStackTemplateResponseModel
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

func (c *Client) UpdateStackTemplate(
	ctx context.Context,
	org string,
	templateId string,
	request *UpdateStackTemplateRequest,
	opts ...option.RequestOption,
) (*UpdateStackTemplateResponseModel, error) {
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

	var response *UpdateStackTemplateResponseModel
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

func (c *Client) DeleteStackTemplate(
	ctx context.Context,
	org string,
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
