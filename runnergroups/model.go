package runnergroups

import (
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
)

type ListAllRunnerGroupsRequest struct {
	// Limit the number of results returned. Default is 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Pagination token to retrieve the next set of results
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	// Filter by resource names (accepted by the API but currently not applied)
	FilterResourceNames *string `json:"-" url:"filterresourcenames,omitempty"`
}

// RunnerGroupSummary is the projection returned by the runner group list endpoint.
type RunnerGroupSummary struct {
	ResourceName  *string               `json:"ResourceName,omitempty" url:"-"`
	SubResourceId *string               `json:"SubResourceId,omitempty" url:"-"`
	ResourceId    *string               `json:"ResourceId,omitempty" url:"-"`
	IsActive      *sgsdkgo.IsPublicEnum `json:"IsActive,omitempty" url:"-"`
	Description   *string               `json:"Description,omitempty" url:"-"`
	Tags          []string              `json:"Tags,omitempty" url:"-"`
	Authors       []string              `json:"Authors,omitempty" url:"-"`
	CreatedAt     *int64                `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt    *int64                `json:"ModifiedAt,omitempty" url:"-"`
}

// RunnerGroupListAllResponse lists runner groups; the API answers 204 with no body when there are none.
type RunnerGroupListAllResponse struct {
	Msg              []*RunnerGroupSummary `json:"msg,omitempty" url:"-"`
	Lastevaluatedkey *string               `json:"lastevaluatedkey,omitempty" url:"-"`
}

type RunnerRegistrationMetadata struct {
	AwsDefaultRegion  *string `json:"AWSDefaultRegion,omitempty" url:"-"`
	EcsCluster        *string `json:"ECSCluster,omitempty" url:"-"`
	SsmActivationId   *string `json:"SSMActivationId,omitempty" url:"-"`
	SsmActivationCode *string `json:"SSMActivationCode,omitempty" url:"-"`
}

// RunnerRegistration is returned when a runner node registers with its runner group.
type RunnerRegistration struct {
	RegistrationMetadata []*RunnerRegistrationMetadata `json:"RegistrationMetadata,omitempty" url:"-"`
	OrgName              *string                       `json:"OrgName,omitempty" url:"-"`
	OrgId                *string                       `json:"OrgId,omitempty" url:"-"`
	RunnerId             *string                       `json:"RunnerId,omitempty" url:"-"`
	RunnerGroupId        *string                       `json:"RunnerGroupId,omitempty" url:"-"`
	Tags                 []string                      `json:"Tags,omitempty" url:"-"`
	// The full runner group item, including StorageBackendConfig and RunnerGroupSignature
	RunnerGroup map[string]interface{} `json:"RunnerGroup,omitempty" url:"-"`
}

type RegisterRunnerResponse struct {
	Msg  *string             `json:"msg,omitempty" url:"-"`
	Data *RunnerRegistration `json:"data,omitempty" url:"-"`
}

type AwsCredentials struct {
	AccessKeyId     string  `json:"AccessKeyId" url:"-"`
	SecretAccessKey string  `json:"SecretAccessKey" url:"-"`
	SessionToken    *string `json:"SessionToken,omitempty" url:"-"`
}

// StorageBackendAuth holds short-lived credentials for the runner group's log/artifact storage backend.
type StorageBackendAuth struct {
	// aws_s3 or azure_blob_storage
	StorageBackendType string          `json:"StorageBackendType" url:"-"`
	Expiration         string          `json:"Expiration" url:"-"`
	AwsCreds           *AwsCredentials `json:"AwsCreds,omitempty" url:"-"`
	AzureSasToken      *string         `json:"AzureSasToken,omitempty" url:"-"`
}

type StorageBackendAuthResponse struct {
	Msg  *string             `json:"msg,omitempty" url:"-"`
	Data *StorageBackendAuth `json:"data,omitempty" url:"-"`
}
