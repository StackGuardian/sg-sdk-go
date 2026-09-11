package billing

import (
	"context"
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

// Read the organization's credit balances.
func (c *Client) ListBalance(
	ctx context.Context,
	org string,
	opts ...option.RequestOption,
) (*BalanceResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/list-balance/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *BalanceResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:                endpointURL,
			Method:             http.MethodGet,
			Headers:            headers,
			MaxAttempts:        options.MaxAttempts,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get the URL of the embeddable billing dashboard (invoices or usage).
func (c *Client) GetDashboardUrl(
	ctx context.Context,
	org string,
	request *DashboardRequest,
	opts ...option.RequestOption,
) (*DashboardResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/dashboard/",
		org,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *DashboardResponse
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

// List the organization's invoices.
func (c *Client) ListInvoices(
	ctx context.Context,
	org string,
	request *ListInvoicesRequest,
	opts ...option.RequestOption,
) (*InvoicesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/invoices/",
		org,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *InvoicesResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:                endpointURL,
			Method:             http.MethodGet,
			Headers:            headers,
			MaxAttempts:        options.MaxAttempts,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Read the organization's current package, trial state and billing information.
func (c *Client) GetBillingDetails(
	ctx context.Context,
	org string,
	opts ...option.RequestOption,
) (*BillingDetailsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/details/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *BillingDetailsResponse
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

// Move the organization to another billing package.
func (c *Client) ChangePlan(
	ctx context.Context,
	org string,
	request *ChangePlanRequest,
	opts ...option.RequestOption,
) (*ChangePlanResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/change-plan/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *ChangePlanResponse
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

// Create the organization's Stripe billing profile. Call it once before adding payment methods.
func (c *Client) SetupBillingProfile(
	ctx context.Context,
	org string,
	request *BillingProfileRequest,
	opts ...option.RequestOption,
) (*SetupBillingProfileResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/profile/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *SetupBillingProfileResponse
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

// Update the email, name or address of the organization's billing profile.
func (c *Client) UpdateBillingProfile(
	ctx context.Context,
	org string,
	request *BillingProfileRequest,
	opts ...option.RequestOption,
) (*UpdateBillingProfileResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/profile/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *UpdateBillingProfileResponse
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

// List the cards on the organization's billing profile.
func (c *Client) ListPaymentMethods(
	ctx context.Context,
	org string,
	opts ...option.RequestOption,
) (*PaymentMethodsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/payment-methods/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *PaymentMethodsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:                endpointURL,
			Method:             http.MethodGet,
			Headers:            headers,
			MaxAttempts:        options.MaxAttempts,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create a Stripe SetupIntent for adding a card; the returned client secret is consumed by Stripe.js.
func (c *Client) CreatePaymentSetupIntent(
	ctx context.Context,
	org string,
	opts ...option.RequestOption,
) (*SetupIntentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/payment-methods/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *SetupIntentResponse
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
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Make a card the organization's default payment method.
func (c *Client) SetDefaultPaymentMethod(
	ctx context.Context,
	org string,
	request *PaymentMethodRequest,
	opts ...option.RequestOption,
) (*BillingMessageResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/set-default-payment-method/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *BillingMessageResponse
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

// Remove a card from the organization's billing profile.
func (c *Client) DetachCard(
	ctx context.Context,
	org string,
	request *PaymentMethodRequest,
	opts ...option.RequestOption,
) (*BillingMessageResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/billing/detach-card/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *BillingMessageResponse
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
