package connectors

import (
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
)

// ListAccountsRequest carries the connector settings to list accounts with. ResourceId may name a
// connector group whose stored config is merged under Settings.config[0].
type ListAccountsRequest struct {
	ResourceId   *string                       `json:"ResourceId,omitempty" url:"-"`
	ResourceName *string                       `json:"ResourceName,omitempty" url:"-"`
	Settings     *sgsdkgo.IntegrationsSettings `json:"Settings,omitempty" url:"-"`
}

// ListAccountsResponse holds the cloud accounts or subscriptions reachable through a connector's
// credentials: AWS Organizations account objects for AWS kinds, Azure subscription objects for Azure kinds.
type ListAccountsResponse struct {
	Msg []map[string]interface{} `json:"msg,omitempty" url:"-"`
}

// GithubComReposResponse lists the "owner/repo" full names accessible to a GITHUB_COM connector.
type GithubComReposResponse struct {
	Msg []string `json:"msg,omitempty" url:"-"`
}

type ListRepositoriesRequest struct {
	// GitHub account or organization login (required for GITHUB_COM connectors)
	Login *string `json:"-" url:"login,omitempty"`
	// GitHub App installation id (required for GITHUB_COM connectors)
	InstallationId *string `json:"-" url:"installationId,omitempty"`
	// 1-based page number, or an opaque cursor for providers that use one
	Page *string `json:"-" url:"page,omitempty"`
	// Page size. Default is 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Filter repositories by name or namespace
	SearchQuery *string `json:"-" url:"searchQuery,omitempty"`
}

type RepositoryOwner struct {
	Id   interface{} `json:"id,omitempty" url:"-"`
	Name *string     `json:"name,omitempty" url:"-"`
}

type Repository struct {
	Id            interface{}      `json:"id,omitempty" url:"-"`
	Url           *string          `json:"url,omitempty" url:"-"`
	Provider      *string          `json:"provider,omitempty" url:"-"`
	Name          *string          `json:"name,omitempty" url:"-"`
	Slug          *string          `json:"slug,omitempty" url:"-"`
	DefaultBranch *string          `json:"defaultBranch,omitempty" url:"-"`
	Namespace     *string          `json:"namespace,omitempty" url:"-"`
	Owner         *RepositoryOwner `json:"owner,omitempty" url:"-"`
	Visibility    *string          `json:"visibility,omitempty" url:"-"`
	UpdatedAt     *string          `json:"updatedAt,omitempty" url:"-"`
	OwnerType     *string          `json:"ownerType,omitempty" url:"-"`
}

// RepositoriesResponse lists repositories; the API answers 204 with no body when there are none.
type RepositoriesResponse struct {
	Msg      *string       `json:"msg,omitempty" url:"-"`
	Data     []*Repository `json:"data,omitempty" url:"-"`
	NextPage interface{}   `json:"next_page,omitempty" url:"-"`
}
