package stacktemplaterevisions

import (
	"context"
	"fmt"
	"net/http"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
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

func (c *Client) CreateStackTemplateRevision(
	ctx context.Context,
	org string,
	templateId string,
	request *CreateStackTemplateRevisionRequest,
	opts ...option.RequestOption,
) (*CreateStackTemplateRevisionResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/IAC_GROUP/%v/%v/revisions/",
		org,
		templateId,
	)

	if request != nil {
		request.TemplateType = string(sgsdkgo.TemplateTypeEnumIacGroup)
	}

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	headers.Add("x-sg-orgid", org)

	var response *CreateStackTemplateRevisionResponseModel
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

func (c *Client) ReadStackTemplateRevision(
	ctx context.Context,
	org string,
	revisionId string,
	opts ...option.RequestOption,
) (*ReadStackTemplateRevisionResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		"IAC_GROUP",
		org,
		revisionId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	var response *ReadStackTemplateRevisionResponseModel
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

func (c *Client) UpdateStackTemplateRevision(
	ctx context.Context,
	org string,
	revisionId string,
	request *UpdateStackTemplateRevisionRequest,
	opts ...option.RequestOption,
) (*UpdateStackTemplateRevisionResponseModel, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)

	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/templatetypes/%v/%v/%v/",
		"IAC_GROUP",
		org,
		revisionId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", fmt.Sprintf("%v", org))

	var response *UpdateStackTemplateRevisionResponseModel
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

func (c *Client) DeleteStackTemplateRevision(
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
		baseURL+"/api/v1/templatetypes/IAC_GROUP/%v/%v/",
		org,
		revisionId,
	)

	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("x-sg-orgid", org)

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
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
		},
	); err != nil {
		return err
	}

	return nil
}
