package secrets

type ReadSecretsInBulkRequest struct {
	// Bare secret names
	ResourceNames []string `json:"ResourceNames" url:"-"`
}

type SecretDetail struct {
	// Full parameter name, e.g. /orgs/<org>/secrets/<name>
	ResourceName  *string                  `json:"ResourceName,omitempty" url:"-"`
	ResourceValue *string                  `json:"ResourceValue,omitempty" url:"-"`
	Attributes    []map[string]interface{} `json:"Attributes,omitempty" url:"-"`
}

// SecretsBulkResponse lists the requested secrets; the API answers 204 with no body when none match.
type SecretsBulkResponse struct {
	Msg []*SecretDetail `json:"msg,omitempty" url:"-"`
}
