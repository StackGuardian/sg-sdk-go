package cli

import (
	"net/url"
	"testing"
)

func TestConnectorGroups(t *testing.T) {
	base := "/api/v1/orgs/demo-org/integrationgroups/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"connector-groups", "create", "-b", `{"ResourceName":"g"}`},
			method: "POST", path: base, body: `{"ResourceName":"g"}`},
		{name: "get", args: []string{"connector-groups", "get", "g"}, method: "GET", path: base + "g/"},
		{name: "delete", args: []string{"connector-groups", "delete", "g"}, method: "DELETE", path: base + "g/"},
		{name: "update", args: []string{"connector-groups", "update", "g", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "g/", body: `{"Description":"d"}`},
		{name: "authenticate", args: []string{"connector-groups", "authenticate", "g"}, method: "GET", path: base + "g/authenticate/"},
		{name: "list", args: []string{"connector-groups", "list", "--resource-names", "n", "--kind", "k", "--last-evaluated-key", "l", "--limit", "5", "--scope", "s"},
			method: "GET", path: base + "listall/",
			query: url.Values{"ResourceNames": {"n"}, "kind": {"k"}, "lastevaluatedkey": {"l"}, "limit": {"5"}, "scope": {"s"}}},
		{name: "child get", args: []string{"connector-groups", "connectors", "get", "c", "--group", "g"},
			method: "GET", path: base + "g/integrations/c/"},
		{name: "child delete", args: []string{"connector-groups", "child", "delete", "c", "--group", "g"},
			method: "DELETE", path: base + "g/integrations/c/"},
		{name: "child update", args: []string{"connector-groups", "connectors", "update", "c", "--group", "g", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "g/integrations/c/", body: `{"Description":"d"}`},
		{name: "child list", args: []string{"connector-groups", "connectors", "list", "--group", "g", "--kind", "k"},
			method: "GET", path: base + "g/integrations/listall/", query: url.Values{"kind": {"k"}}},
	})
}

func TestConnectors(t *testing.T) {
	base := "/api/v1/orgs/demo-org/integrations/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"connectors", "create", "-b", `{"ResourceName":"c"}`},
			method: "POST", path: base, body: `{"ResourceName":"c"}`},
		{name: "get", args: []string{"connectors", "get", "c"}, method: "GET", path: base + "c/"},
		{name: "delete", args: []string{"connectors", "delete", "c"}, method: "DELETE", path: base + "c/"},
		{name: "update", args: []string{"connectors", "update", "c", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "c/", body: `{"Description":"d"}`},
		{name: "list", args: []string{"connectors", "list", "--resource-names", "n", "--kind", "AWS_RBAC", "--last-evaluated-key", "l", "--limit", "5"},
			method: "GET", path: base + "listall/",
			query: url.Values{"ResourceNames": {"n"}, "kind": {"AWS_RBAC"}, "lastevaluatedkey": {"l"}, "limit": {"5"}}},
	})
}
