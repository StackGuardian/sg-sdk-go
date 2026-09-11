package chats

// Chat is the create request body. Config is polymorphic on ChatType: infra2code chats carry
// wfGrpId/wfId/latestWfRunId/connectorId/iacTools/resources/providers/regions/Pr/codeGeneration/
// pendingReview/runPlan, kyro chats carry activeLeafMessageId.
type Chat struct {
	// infra2code (default) or kyro
	ChatType     *string                `json:"ChatType,omitempty" url:"-"`
	Config       map[string]interface{} `json:"Config" url:"-"`
	ResourceName *string                `json:"ResourceName,omitempty" url:"-"`
	Description  *string                `json:"Description,omitempty" url:"-"`
	Tags         []string               `json:"Tags,omitempty" url:"-"`
	// template, workflow or git; seeds the chat's git bundle from that source
	SourceType        *string `json:"SourceType,omitempty" url:"-"`
	SourceTemplateId  *string `json:"SourceTemplateId,omitempty" url:"-"`
	SourceWorkflowId  *string `json:"SourceWorkflowId,omitempty" url:"-"`
	SourceRepoUrl     *string `json:"SourceRepoUrl,omitempty" url:"-"`
	SourceRef         *string `json:"SourceRef,omitempty" url:"-"`
	SourceWorkingDir  *string `json:"SourceWorkingDir,omitempty" url:"-"`
	SourceConnectorId *string `json:"SourceConnectorId,omitempty" url:"-"`
}

type ChatCreated struct {
	ResourceId    *string `json:"ResourceId,omitempty" url:"-"`
	S3Prefix      *string `json:"S3Prefix,omitempty" url:"-"`
	ResourceName  *string `json:"ResourceName,omitempty" url:"-"`
	BundleVersion *int    `json:"BundleVersion,omitempty" url:"-"`
}

type ChatCreateResponse struct {
	Msg  *string      `json:"msg,omitempty" url:"-"`
	Data *ChatCreated `json:"data,omitempty" url:"-"`
}

// ChatData is a chat as stored by the API.
type ChatData struct {
	OrgId         *string                `json:"OrgId,omitempty" url:"-"`
	ResourceId    *string                `json:"ResourceId,omitempty" url:"-"`
	ResourceName  *string                `json:"ResourceName,omitempty" url:"-"`
	ResourceType  *string                `json:"ResourceType,omitempty" url:"-"`
	ResourceKSUID *string                `json:"ResourceKSUID,omitempty" url:"-"`
	ChatType      *string                `json:"ChatType,omitempty" url:"-"`
	Status        *string                `json:"Status,omitempty" url:"-"`
	S3Prefix      *string                `json:"S3Prefix,omitempty" url:"-"`
	BundleVersion *int                   `json:"BundleVersion,omitempty" url:"-"`
	Description   *string                `json:"Description,omitempty" url:"-"`
	DocVersion    *string                `json:"DocVersion,omitempty" url:"-"`
	Authors       []string               `json:"Authors,omitempty" url:"-"`
	Tags          []string               `json:"Tags,omitempty" url:"-"`
	CreatedAt     *int64                 `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt    *int64                 `json:"ModifiedAt,omitempty" url:"-"`
	Config        map[string]interface{} `json:"Config,omitempty" url:"-"`
}

type ListAllChatsRequest struct {
	// Page size. Default is 20.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Pagination token to retrieve the next set of results
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
	// Comma-separated terms matched against name, description and tags
	SearchQuery *string `json:"-" url:"searchQuery,omitempty"`
	// Refresh the pull request status of listed chats from the VCS provider
	RefreshPrStatus *bool `json:"-" url:"refreshPrStatus,omitempty"`
	// Filter by chat type (infra2code or kyro)
	ChatType *string `json:"-" url:"chatType,omitempty"`
}

type ChatListAllResponse struct {
	Msg              []*ChatData `json:"msg,omitempty" url:"-"`
	Lastevaluatedkey *string     `json:"lastevaluatedkey,omitempty" url:"-"`
}

type ReadChatRequest struct {
	// Refresh the pull request status from the VCS provider
	RefreshPrStatus *bool `json:"-" url:"refreshPrStatus,omitempty"`
}

type ChatReadResponse struct {
	Msg *ChatData `json:"msg,omitempty" url:"-"`
}

// PatchedChat is the update request body. Config is deep-merged by the API.
type PatchedChat struct {
	ResourceName *string                `json:"ResourceName,omitempty" url:"-"`
	Status       *string                `json:"Status,omitempty" url:"-"`
	Description  *string                `json:"Description,omitempty" url:"-"`
	Tags         []string               `json:"Tags,omitempty" url:"-"`
	Config       map[string]interface{} `json:"Config,omitempty" url:"-"`
}

type ChatUpdateResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
	// The attributes that were changed
	Data map[string]interface{} `json:"data,omitempty" url:"-"`
}

type ChatDeleteResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

type ListAllChatArtifactsRequest struct {
	Prefix        *string `json:"-" url:"prefix,omitempty"`
	MaxKeys       *int    `json:"-" url:"maxKeys,omitempty"`
	StartAfterKey *string `json:"-" url:"startAfterKey,omitempty"`
}

type GetChatArtifactRequest struct {
	// Path relative to the chat's artifact store
	ArtifactPath string `json:"-" url:"artifactPath"`
}

type SaveChatArtifactRequest struct {
	// Path relative to the chat's artifact store
	ArtifactPath string `json:"-" url:"artifactPath"`
	// Content type the upload must use. Default is application/octet-stream.
	ContentType *string `json:"-" url:"contentType,omitempty"`
}

// ChatArtifactUrl is a presigned S3 URL; the artifact bytes are transferred with it, not through the API.
type ChatArtifactUrl struct {
	SignedUrl    string  `json:"signedUrl" url:"-"`
	ArtifactPath *string `json:"artifactPath,omitempty" url:"-"`
	ExpiresIn    *int    `json:"expiresIn,omitempty" url:"-"`
}

type ChatArtifactUrlResponse struct {
	Msg *ChatArtifactUrl `json:"msg,omitempty" url:"-"`
}

type CommitChatBundleRequest struct {
	// The BundleVersion the caller last saw; the commit fails with 409 if it changed
	ExpectedVersion int `json:"expected_version" url:"-"`
}

type ChatBundleVersion struct {
	BundleVersion *int   `json:"BundleVersion,omitempty" url:"-"`
	ModifiedAt    *int64 `json:"ModifiedAt,omitempty" url:"-"`
}

type CommitChatBundleResponse struct {
	Msg *ChatBundleVersion `json:"msg,omitempty" url:"-"`
}

type CreateChatPrRequest struct {
	VcsConnector  string  `json:"vcsConnector" url:"-"`
	RepoUrl       string  `json:"repoUrl" url:"-"`
	TargetBranch  *string `json:"targetBranch,omitempty" url:"-"`
	SourceBranch  *string `json:"sourceBranch,omitempty" url:"-"`
	WorkingDir    *string `json:"workingDir,omitempty" url:"-"`
	CommitMessage *string `json:"commitMessage,omitempty" url:"-"`
	PrTitle       *string `json:"prTitle,omitempty" url:"-"`
	PrDescription *string `json:"prDescription,omitempty" url:"-"`
	ReuseBranch   *bool   `json:"reuseBranch,omitempty" url:"-"`
}

type ChatPr struct {
	PrNumber     *int    `json:"prNumber,omitempty" url:"-"`
	PrUrl        *string `json:"prUrl,omitempty" url:"-"`
	SourceBranch *string `json:"sourceBranch,omitempty" url:"-"`
	TargetBranch *string `json:"targetBranch,omitempty" url:"-"`
}

type ChatPrResponse struct {
	Msg  *string `json:"msg,omitempty" url:"-"`
	Data *ChatPr `json:"data,omitempty" url:"-"`
}

// ChatPrStatusResponse carries the pull request state; Data is absent when the chat has no PR.
type ChatPrStatusResponse struct {
	Msg  *string                `json:"msg,omitempty" url:"-"`
	Data map[string]interface{} `json:"data,omitempty" url:"-"`
}

type ChatPrSyncStatus struct {
	InSync                  *bool   `json:"inSync,omitempty" url:"-"`
	LocalCodeChanged        *bool   `json:"localCodeChanged,omitempty" url:"-"`
	RemoteChanged           *bool   `json:"remoteChanged,omitempty" url:"-"`
	CurrentBundleVersion    *int    `json:"currentBundleVersion,omitempty" url:"-"`
	LastPushedBundleVersion *int    `json:"lastPushedBundleVersion,omitempty" url:"-"`
	HeadCommitSha           *string `json:"headCommitSha,omitempty" url:"-"`
	LastPushedCommitSha     *string `json:"lastPushedCommitSha,omitempty" url:"-"`
	PrStatus                *string `json:"prStatus,omitempty" url:"-"`
	PrUrl                   *string `json:"prUrl,omitempty" url:"-"`
}

type ChatPrSyncStatusResponse struct {
	Msg  *string           `json:"msg,omitempty" url:"-"`
	Data *ChatPrSyncStatus `json:"data,omitempty" url:"-"`
}

// ContentBlock is one block of a message. Type is text, tool_use, tool_result or image.
type ContentBlock struct {
	Type      string                 `json:"type" url:"-"`
	Text      *string                `json:"text,omitempty" url:"-"`
	ToolUseId *string                `json:"toolUseId,omitempty" url:"-"`
	Name      *string                `json:"name,omitempty" url:"-"`
	Input     map[string]interface{} `json:"input,omitempty" url:"-"`
	Content   []interface{}          `json:"content,omitempty" url:"-"`
	IsError   *bool                  `json:"isError,omitempty" url:"-"`
	Source    map[string]interface{} `json:"source,omitempty" url:"-"`
}

// Message is the create request body. Only kyro chats accept messages.
type Message struct {
	Content []*ContentBlock `json:"content" url:"-"`
	// Idempotency key: a repeated id returns the existing message instead of creating one
	ClientRequestId *string `json:"clientRequestId,omitempty" url:"-"`
}

// MessageCreated holds the ids of the user message and the pending assistant reply. When the
// request was deduplicated by ClientRequestId the API returns the existing message item instead.
type MessageCreated struct {
	UserMessageId      *string `json:"userMessageId,omitempty" url:"-"`
	AssistantMessageId *string `json:"assistantMessageId,omitempty" url:"-"`
}

type MessageCreateResponse struct {
	Msg *MessageCreated `json:"msg,omitempty" url:"-"`
}

type MessageData struct {
	OrgId           *string         `json:"OrgId,omitempty" url:"-"`
	ResourceId      *string         `json:"ResourceId,omitempty" url:"-"`
	ResourceType    *string         `json:"ResourceType,omitempty" url:"-"`
	ResourceKSUID   *string         `json:"ResourceKSUID,omitempty" url:"-"`
	ChatId          *string         `json:"ChatId,omitempty" url:"-"`
	Role            *string         `json:"Role,omitempty" url:"-"`
	Status          *string         `json:"Status,omitempty" url:"-"`
	Content         []*ContentBlock `json:"Content,omitempty" url:"-"`
	ContentText     *string         `json:"ContentText,omitempty" url:"-"`
	ParentMessageId *string         `json:"ParentMessageId,omitempty" url:"-"`
	Author          *string         `json:"Author,omitempty" url:"-"`
	ClientRequestId *string         `json:"ClientRequestId,omitempty" url:"-"`
	CreatedAt       *int64          `json:"CreatedAt,omitempty" url:"-"`
	ModifiedAt      *int64          `json:"ModifiedAt,omitempty" url:"-"`
}

type ListAllMessagesRequest struct {
	// Page size. Default is 50.
	Limit *int `json:"-" url:"limit,omitempty"`
	// Pagination token from a previous response
	Lastevaluatedkey *string `json:"-" url:"lastevaluatedkey,omitempty"`
}

type MessageListAllResponse struct {
	Msg              []*MessageData `json:"msg,omitempty" url:"-"`
	Lastevaluatedkey interface{}    `json:"lastevaluatedkey,omitempty" url:"-"`
}

type MessageReadResponse struct {
	Msg *MessageData `json:"msg,omitempty" url:"-"`
}

type MessageActionResponse struct {
	Msg *string `json:"msg,omitempty" url:"-"`
}

type MessageRetried struct {
	AssistantMessageId *string `json:"assistantMessageId,omitempty" url:"-"`
}

type MessageRetryResponse struct {
	Msg *MessageRetried `json:"msg,omitempty" url:"-"`
}
