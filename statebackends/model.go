package statebackends

import (
	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/core"
)

// StateBackendConfig describes where a state backend stores state files.
type StateBackendConfig struct {
	// One of aws_s3, azure_blob_storage.
	Type                          string  `json:"type" url:"-"`
	S3BucketName                  *string `json:"s3BucketName,omitempty" url:"-"`
	AwsRegion                     *string `json:"awsRegion,omitempty" url:"-"`
	AzureBlobStorageAccountName   *string `json:"azureBlobStorageAccountName,omitempty" url:"-"`
	AzureBlobStorageContainerName *string `json:"azureBlobStorageContainerName,omitempty" url:"-"`
	AzureBlobStorageResourceId    *string `json:"azureBlobStorageResourceId,omitempty" url:"-"`
	AzureBlobStorageAccessKey     *string `json:"azureBlobStorageAccessKey,omitempty" url:"-"`
	// Required for aws_s3; for azure_blob_storage either Auth or AzureBlobStorageAccessKey
	Auth *StateBackendAuth `json:"auth,omitempty" url:"-"`
}

// StateBackendAuth references the connector whose credentials reach the storage backend.
type StateBackendAuth struct {
	IntegrationId string `json:"integrationId" url:"-"`
}

// StateBackend is the create request body.
type StateBackend struct {
	ResourceName       *string               `json:"ResourceName,omitempty" url:"-"`
	Description        *string               `json:"Description,omitempty" url:"-"`
	Tags               []string              `json:"Tags,omitempty" url:"-"`
	IsActive           *sgsdkgo.IsPublicEnum `json:"IsActive,omitempty" url:"-"`
	StateBackendConfig *StateBackendConfig   `json:"StateBackendConfig,omitempty" url:"-"`
	// Required by the API; pass an empty list to start without state files
	Statefiles []string `json:"Statefiles" url:"-"`
}

// PatchedStateBackend is the update request body.
type PatchedStateBackend struct {
	ResourceName       *core.Optional[string]               `json:"ResourceName,omitempty" url:"-"`
	Description        *core.Optional[string]               `json:"Description,omitempty" url:"-"`
	Tags               *core.Optional[[]string]             `json:"Tags,omitempty" url:"-"`
	IsActive           *core.Optional[sgsdkgo.IsPublicEnum] `json:"IsActive,omitempty" url:"-"`
	StateBackendConfig *core.Optional[StateBackendConfig]   `json:"StateBackendConfig,omitempty" url:"-"`
	Statefiles         *core.Optional[[]string]             `json:"Statefiles,omitempty" url:"-"`
}

// StateBackendData is a state backend as returned by the API.
type StateBackendData struct {
	Id                 *string               `json:"Id,omitempty" url:"-"`
	ResourceName       *string               `json:"ResourceName,omitempty" url:"-"`
	Description        *string               `json:"Description,omitempty" url:"-"`
	Tags               []string              `json:"Tags,omitempty" url:"-"`
	IsActive           *sgsdkgo.IsPublicEnum `json:"IsActive,omitempty" url:"-"`
	IsArchive          *string               `json:"IsArchive,omitempty" url:"-"`
	StateBackendConfig *StateBackendConfig   `json:"StateBackendConfig,omitempty" url:"-"`
	CreatedAt          *int64                `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt         *int64                `json:"ModifiedAt,omitempty" url:"-"`
	ParentId           *string               `json:"ParentId,omitempty" url:"-"`
	ResourceType       *string               `json:"ResourceType,omitempty" url:"-"`
	ResourceId         *string               `json:"ResourceId,omitempty" url:"-"`
	SubResourceId      *string               `json:"SubResourceId,omitempty" url:"-"`
	OrgId              *string               `json:"OrgId,omitempty" url:"-"`
	DocVersion         *string               `json:"DocVersion,omitempty" url:"-"`
	Authors            []string              `json:"Authors,omitempty" url:"-"`
	Statefiles         []interface{}         `json:"Statefiles,omitempty" url:"-"`
	StatefileCount     *int                  `json:"StatefileCount,omitempty" url:"-"`
}

type StateBackendResponse struct {
	Msg  *string           `json:"msg,omitempty" url:"-"`
	Data *StateBackendData `json:"data,omitempty" url:"-"`
}

type DeleteStateBackendResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

type ListAllStateBackendsRequest struct {
	// Pagination token to retrieve the next set of results
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	// Limit the number of results returned. Default is 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Filter state backends by resource name, tags or description
	SearchQuery *string `json:"-" url:"searchQuery,omitempty"`
}

type StateBackendListAllResponse struct {
	Msg              *string             `json:"msg,omitempty" url:"-"`
	Data             []*StateBackendData `json:"data,omitempty" url:"-"`
	Lastevaluatedkey *string             `json:"lastevaluatedkey,omitempty" url:"-"`
}

type ListStatefilesRequest struct {
	// Base64-encoded JSON of a StateBackendConfig
	Config string `json:"-" url:"config"`
}

type ListStatefilesResponse struct {
	Msg  *string                  `json:"msg,omitempty" url:"-"`
	Data []map[string]interface{} `json:"data,omitempty" url:"-"`
}
