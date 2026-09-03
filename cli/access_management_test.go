package cli

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganizations(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "get", args: []string{"organizations", "get"}, method: "GET", path: "/api/v1/orgs/demo-org/"},
		{name: "alias", args: []string{"org", "get"}, method: "GET", path: "/api/v1/orgs/demo-org/"},
	})
}

func TestApiAccess(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "create", args: []string{"api-access", "create", "-b", `{"ResourceName":"k1"}`},
			method: "POST", path: "/api/v1/orgs/demo-org/apiaccesses/", body: `{"ResourceName":"k1"}`},
		{name: "get", args: []string{"api-access", "get", "id1"}, method: "GET", path: "/api/v1/orgs/demo-org/apiaccesses/id1/"},
		{name: "delete", args: []string{"api-access", "rm", "id1"}, method: "DELETE", path: "/api/v1/orgs/demo-org/apiaccesses/id1/"},
		{name: "update", args: []string{"api-access", "update", "id1", "-b", `{"Description":"d"}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/apiaccesses/id1/", body: `{"Description":"d"}`},
		{name: "regenerate expires", args: []string{"api-access", "regenerate-key", "id1", "--expires-at", "123"},
			method: "POST", path: "/api/v1/orgs/demo-org/apiaccesses/id1/regenerate/", body: `{"ExpiresAt":123}`},
		{name: "regenerate null", args: []string{"api-access", "regenerate-key", "id1", "--no-expiration"},
			method: "POST", path: "/api/v1/orgs/demo-org/apiaccesses/id1/regenerate/", body: `{"ExpiresAt":null}`},
		{name: "regenerate empty", args: []string{"api-access", "regenerate-key", "id1"},
			method: "POST", path: "/api/v1/orgs/demo-org/apiaccesses/id1/regenerate/", body: `{}`},
		{name: "list", args: []string{"api-access", "list", "--access-type", "APIKEY", "--limit", "10", "--last-evaluated-key", "abc",
			"--resource-name", "r", "--roles", "x,y", "--status", "active"},
			method: "GET", path: "/api/v1/orgs/demo-org/apiaccesses/listall/",
			query: url.Values{"accessType": {"APIKEY"}, "limit": {"10"}, "lastEvaluatedKey": {"abc"}, "resourceName": {"r"}, "roles": {"x,y"}, "status": {"active"}}},
		{name: "list plain", args: []string{"api-access", "ls"}, method: "GET", path: "/api/v1/orgs/demo-org/apiaccesses/listall/"},
	})
	t.Run("invalid access type", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "api-access", "list", "--access-type", "BAD")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--access-type")
	})
}

func TestAuditLogs(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "list", args: []string{"audit-logs", "list", "--effect", "Allow", "--start-time", "1", "--end-time", "2", "--limit", "5",
			"--last-evaluated-key", "k", "--principal-email", "e@x", "--request-method", "GET", "--resource", "SG_SIGN_IN", "--source-ip", "1.2.3.4"},
			method: "GET", path: "/api/v1/orgs/demo-org/audit_logs/",
			query: url.Values{"effect": {"Allow"}, "startTime": {"1"}, "endTime": {"2"}, "limit": {"5"}, "lastevaluatedkey": {"k"},
				"principalEmail": {"e@x"}, "request_method": {"GET"}, "resource": {"SG_SIGN_IN"}, "sourceIp": {"1.2.3.4"}}},
	})
}

func TestUsers(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "get by id", args: []string{"users", "get", "--user-id", "u@x"},
			method: "POST", path: "/api/v1/orgs/demo-org/get_user/", body: `{"userId":"u@x"}`},
		{name: "get by alias", args: []string{"users", "get", "--alias", "grp"},
			method: "POST", path: "/api/v1/orgs/demo-org/get_user/", body: `{"alias":"grp"}`},
		{name: "create", args: []string{"users", "create", "-b", `{"userId":"u@x","role":"ADMIN"}`},
			method: "POST", path: "/api/v1/orgs/demo-org/invite_user/", body: `{"userId":"u@x","role":"ADMIN"}`},
		{name: "delete", args: []string{"users", "delete", "--user-id", "u@x"},
			method: "POST", path: "/api/v1/orgs/demo-org/remove_user/", body: `{"userId":"u@x"}`},
		{name: "update", args: []string{"users", "update", "-b", `{"userId":"u@x","role":"USER"}`},
			method: "POST", path: "/api/v1/orgs/demo-org/update_user/", body: `{"userId":"u@x","role":"USER"}`},
		{name: "list", args: []string{"users", "list", "--alias", "a", "--entity-type", "EMAIL", "--login-method", "m", "--roles", "ADMIN", "--user-id", "u"},
			method: "GET", path: "/api/v1/orgs/demo-org/users/listall/",
			query: url.Values{"alias": {"a"}, "entityType": {"EMAIL"}, "loginMethod": {"m"}, "roles": {"ADMIN"}, "userId": {"u"}}},
	})
	t.Run("get requires selector", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "users", "get")
		require.Error(t, err)
	})
}

func TestRoles(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "create", args: []string{"roles", "create", "-b", `{"ResourceName":"r1"}`},
			method: "POST", path: "/api/v1/orgs/demo-org/roles/", body: `{"ResourceName":"r1"}`},
		{name: "get", args: []string{"roles", "get", "r1"}, method: "GET", path: "/api/v1/orgs/demo-org/roles/r1/"},
		{name: "delete", args: []string{"roles", "delete", "r1"}, method: "DELETE", path: "/api/v1/orgs/demo-org/roles/r1/"},
		{name: "update", args: []string{"roles", "update", "r1", "-b", `{"Description":"d"}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/roles/r1/", body: `{"ResourceName":"r1","Description":"d"}`},
		{name: "list", args: []string{"roles", "list", "--include-predefined-roles"},
			method: "GET", path: "/api/v1/orgs/demo-org/roles/listall/", query: url.Values{"includePredefinedRoles": {"true"}},
			response: `{"msg":[{"ResourceName":"ADMIN"}]}`},
	})
}
