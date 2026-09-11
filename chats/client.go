package chats

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

// Create a chat, optionally seeded from a template, workflow or git repository.
func (c *Client) CreateChat(
	ctx context.Context,
	org string,
	request *Chat,
	opts ...option.RequestOption,
) (*ChatCreateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/",
		org,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *ChatCreateResponse
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

// List the chats in an organization.
func (c *Client) ListAllChats(
	ctx context.Context,
	org string,
	request *ListAllChatsRequest,
	opts ...option.RequestOption,
) (*ChatListAllResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/listall/",
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

	var response *ChatListAllResponse
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

// Read a chat.
func (c *Client) ReadChat(
	ctx context.Context,
	org string,
	chat string,
	request *ReadChatRequest,
	opts ...option.RequestOption,
) (*ChatReadResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/",
		org,
		chat,
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

	var response *ChatReadResponse
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

// Update a chat. Nested Config values are deep-merged.
func (c *Client) UpdateChat(
	ctx context.Context,
	org string,
	chat string,
	request *PatchedChat,
	opts ...option.RequestOption,
) (*ChatUpdateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *ChatUpdateResponse
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

// Archive a chat.
func (c *Client) DeleteChat(
	ctx context.Context,
	org string,
	chat string,
	opts ...option.RequestOption,
) (*ChatDeleteResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *ChatDeleteResponse
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
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Get a presigned URL that lists a chat's artifacts.
func (c *Client) ListAllChatArtifacts(
	ctx context.Context,
	org string,
	chat string,
	request *ListAllChatArtifactsRequest,
	opts ...option.RequestOption,
) (*ChatArtifactUrlResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/listall_artifacts/",
		org,
		chat,
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

	var response *ChatArtifactUrlResponse
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

// Get a presigned URL to download a chat artifact.
func (c *Client) GetChatArtifact(
	ctx context.Context,
	org string,
	chat string,
	request *GetChatArtifactRequest,
	opts ...option.RequestOption,
) (*ChatArtifactUrlResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/get_artifact/",
		org,
		chat,
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

	var response *ChatArtifactUrlResponse
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

// Get a presigned URL to upload a chat artifact. The parameters are sent as query parameters; the upload itself goes to the returned URL.
func (c *Client) SaveChatArtifact(
	ctx context.Context,
	org string,
	chat string,
	request *SaveChatArtifactRequest,
	opts ...option.RequestOption,
) (*ChatArtifactUrlResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/save_artifact/",
		org,
		chat,
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

	var response *ChatArtifactUrlResponse
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

// Commit an uploaded git bundle, bumping the chat's BundleVersion with optimistic concurrency.
func (c *Client) CommitChatBundle(
	ctx context.Context,
	org string,
	chat string,
	request *CommitChatBundleRequest,
	opts ...option.RequestOption,
) (*CommitChatBundleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/commit_bundle/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *CommitChatBundleResponse
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

// Push the chat's generated code to a repository and open (or update) a pull request.
func (c *Client) CreateChatPr(
	ctx context.Context,
	org string,
	chat string,
	request *CreateChatPrRequest,
	opts ...option.RequestOption,
) (*ChatPrResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/create_pr/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *ChatPrResponse
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

// Read the live status of the chat's pull request.
func (c *Client) GetChatPrStatus(
	ctx context.Context,
	org string,
	chat string,
	opts ...option.RequestOption,
) (*ChatPrStatusResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/pr_status/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *ChatPrStatusResponse
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

// Compare the chat's local bundle with the state of its pull request.
func (c *Client) GetChatPrSyncStatus(
	ctx context.Context,
	org string,
	chat string,
	opts ...option.RequestOption,
) (*ChatPrSyncStatusResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/pr_sync_status/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *ChatPrSyncStatusResponse
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

// Send a message to a kyro chat. Returns the ids of the stored user message and the pending assistant reply.
func (c *Client) CreateMessage(
	ctx context.Context,
	org string,
	chat string,
	request *Message,
	opts ...option.RequestOption,
) (*MessageCreateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/messages/",
		org,
		chat,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *MessageCreateResponse
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

// List the messages of a chat, oldest first.
func (c *Client) ListAllMessages(
	ctx context.Context,
	org string,
	chat string,
	request *ListAllMessagesRequest,
	opts ...option.RequestOption,
) (*MessageListAllResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/messages/listall/",
		org,
		chat,
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

	var response *MessageListAllResponse
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

// Read a message.
func (c *Client) ReadMessage(
	ctx context.Context,
	org string,
	chat string,
	message string,
	opts ...option.RequestOption,
) (*MessageReadResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/messages/%v/",
		org,
		chat,
		message,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *MessageReadResponse
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

// Stop a pending or streaming assistant message.
func (c *Client) CancelMessage(
	ctx context.Context,
	org string,
	chat string,
	message string,
	opts ...option.RequestOption,
) (*MessageActionResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/messages/%v/cancel/",
		org,
		chat,
		message,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *MessageActionResponse
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

// Retry an errored or stopped assistant message; returns the id of the new pending reply.
func (c *Client) RetryMessage(
	ctx context.Context,
	org string,
	chat string,
	message string,
	opts ...option.RequestOption,
) (*MessageRetryResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.app.stackguardian.io",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/orgs/%v/chats/%v/messages/%v/retry/",
		org,
		chat,
		message,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *MessageRetryResponse
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
