package cli

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplates(t *testing.T) {
	orgHeader := map[string]string{"X-Sg-Orgid": "demo-org"}
	runCases(t, []cmdCase{
		{name: "subscriptions", args: []string{"templates", "subscriptions", "get", "--type", "IACSubscriptions"},
			method: "GET", path: "/api/v1/orgs/demo-org/subscriptions/default/", query: url.Values{"subscriptionType": {"IACSubscriptions"}}},
		{name: "list", args: []string{"templates", "list", "--type", "IAC", "--is-public", "1", "--owner-orgs", "o1,o2", "--search-query", "q",
			"--template-id", "tid", "--last-evaluated-key", "k"},
			method: "GET", path: "/api/v1/templatetypes/IAC/templates/listall/", headers: orgHeader,
			query: url.Values{"IsPublic": {"1"}, "OwnerOrgs": {"o1,o2"}, "SearchQuery": {"q"}, "TemplateId": {"tid"}, "lastevaluatedkey": {"k"}}},
		{name: "list by owner", args: []string{"templates", "list-by-owner", "owner-org", "--is-public", "0", "--search-query", "q", "--shared", "--last-evaluated-key", "k"},
			method: "GET", path: "/api/v1/orgs/owner-org/templates/listall/", headers: orgHeader,
			query: url.Values{"IsPublic": {"0"}, "SearchQuery": {"q"}, "isSharedTemplate": {"true"}, "lastevaluatedkey": {"k"}}},
		{name: "create", args: []string{"templates", "create", "-b", `{"TemplateType":"IAC","TemplateName":"t1"}`},
			method: "POST", path: "/api/v1/templates/", headers: orgHeader,
			body:     `{"TemplateType":"IAC","TemplateName":"t1","OwnerOrg":"","SourceConfigKind":""}`,
			response: `{"TemplateType":"IAC","TemplateName":"t1"}`},
		{name: "get", args: []string{"templates", "get", "t1:5", "--type", "IAC"},
			method: "GET", path: "/api/v1/templatetypes/IAC/demo-org/t1:5/", headers: orgHeader},
		{name: "get other owner", args: []string{"templates", "get", "t1", "--type", "IAC_GROUP", "--owner-org", "stackguardian"},
			method: "GET", path: "/api/v1/templatetypes/IAC_GROUP/stackguardian/t1/", headers: orgHeader},
		{name: "delete", args: []string{"templates", "delete", "t1", "--type", "IAC_POLICY"},
			method: "DELETE", path: "/api/v1/templatetypes/IAC_POLICY/demo-org/t1/", headers: orgHeader},
		{name: "update", args: []string{"templates", "update", "t1:2", "--type", "WORKFLOW_STEP", "-b", `{"TemplateName":"t2"}`},
			method: "PATCH", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1:2/", headers: orgHeader, body: `{"TemplateName":"t2"}`},
		{name: "group iac", args: []string{"templates", "get-group-iac", "grp", "sub:1", "--owner-org", "stackguardian"},
			method: "GET", path: "/api/v1/templatetypes/IAC_GROUP/stackguardian/grp/IAC/sub:1/", headers: orgHeader},
	})
	t.Run("invalid template type", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "templates", "get", "t1", "--type", "NOPE")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--type")
	})
	t.Run("invalid subscription type", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "templates", "subscriptions", "get", "--type", "NOPE")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--type")
	})
}
