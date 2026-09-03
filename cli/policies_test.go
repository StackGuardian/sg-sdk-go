package cli

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPolicies(t *testing.T) {
	base := "/api/v1/orgs/demo-org/policies/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"policies", "create", "-b", `{"PolicyType":"GENERAL","ResourceName":"p"}`},
			method: "POST", path: base, body: `{"PolicyType":"GENERAL","ResourceName":"p"}`},
		{name: "get", args: []string{"policies", "get", "p"}, method: "GET", path: base + "p/"},
		{name: "delete", args: []string{"policies", "delete", "p"}, method: "DELETE", path: base + "p/"},
		{name: "update", args: []string{"policies", "update", "p", "-b", `{"PolicyType":"GENERAL","Description":"d"}`},
			method: "PATCH", path: base + "p/", body: `{"PolicyType":"GENERAL","Description":"d"}`},
		{name: "list", args: []string{"policies", "list", "--description", "d", "--policy-types", "GENERAL", "--resource-names", "n",
			"--tags", "t", "--last-evaluated-key", "l", "--limit", "5", "--search-query", "q"},
			method: "GET", path: base + "listall/",
			query: url.Values{"Description": {"d"}, "PolicyTypes": {"GENERAL"}, "ResourceNames": {"n"}, "Tags": {"t"},
				"lastevaluatedkey": {"l"}, "limit": {"5"}, "searchQuery": {"q"}},
			response: `{"msg":[]}`},
	})
	t.Run("create without discriminator", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "policies", "create", "-b", `{"ResourceName":"p"}`)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "PolicyType")
	})
}

func TestRunnerGroups(t *testing.T) {
	base := "/api/v1/orgs/demo-org/runnergroups/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"runner-groups", "create", "-b", `{"ResourceName":"rg"}`},
			method: "POST", path: base, body: `{"ResourceName":"rg"}`},
		{name: "get", args: []string{"runner-groups", "get", "rg", "--get-active-workflows", "--get-active-workflows-details"},
			method: "GET", path: base + "rg/", query: url.Values{"getActiveWorkflows": {"true"}, "getActiveWorkflowsDetails": {"true"}}},
		{name: "get plain", args: []string{"runner-groups", "get", "rg"}, method: "GET", path: base + "rg/"},
		{name: "delete", args: []string{"runner-groups", "delete", "rg"}, method: "DELETE", path: base + "rg/"},
		{name: "update", args: []string{"runner-groups", "update", "rg", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "rg/", body: `{"Description":"d"}`},
		{name: "deregister", args: []string{"runner-groups", "deregister-runner", "rg", "--runner-id", "r1", "--container-instance-ids", "a,b", "--force"},
			method: "POST", path: base + "rg/deregister/", body: `{"RunnerId":"r1","ContainerInstanceIds":["a","b"],"ForceDeregister":true}`},
		{name: "update state", args: []string{"runner-groups", "update-runner-state", "rg", "--status", "DRAINING", "--runner-id", "r1"},
			method: "POST", path: base + "rg/runner_status/", body: `{"RunnerId":"r1","Status":"DRAINING"}`},
	})
	t.Run("invalid status", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "runner-groups", "update-runner-state", "rg", "--status", "BAD")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--status")
	})
}

func TestSecrets(t *testing.T) {
	base := "/api/v1/orgs/demo-org/secrets/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"secrets", "create", "-b", `{"ResourceName":"s","ResourceValue":"v"}`},
			method: "POST", path: base, body: `{"ResourceName":"s","ResourceValue":"v"}`},
		{name: "delete", args: []string{"secrets", "delete", "s"}, method: "DELETE", path: base + "s/"},
		{name: "update", args: []string{"secrets", "update", "s", "-b", `{"ResourceValue":"v2"}`},
			method: "PATCH", path: base + "s/", body: `{"ResourceName":"s","ResourceValue":"v2"}`},
		{name: "list", args: []string{"secrets", "list"}, method: "GET", path: base + "listall/"},
	})
}

func TestBenchmarkReports(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "get", args: []string{"benchmark-reports", "get", "--date", "01_01_2025", "--detailed", "--filter-csp", "AWS",
			"--filter-account-id", "1", "--filter-benchmark", "cis_v200", "--filter-control-id", "c", "--filter-control-title", "t",
			"--filter-region", "global", "--filter-severity", "High", "--filter-status", "fails", "--group-by", "g",
			"--limit", "10", "--page", "2", "--required-columns", "CSP,status"},
			method: "GET", path: "/api/v1/orgs/demo-org/reports/benchmark/",
			query: url.Values{"date": {"01_01_2025"}, "detailed": {"true"}, "filter:CSP": {"AWS"}, "filter:accountId": {"1"},
				"filter:benchmark": {"cis_v200"}, "filter:controlId": {"c"}, "filter:controlTitle": {"t"}, "filter:region": {"global"},
				"filter:severity": {"High"}, "filter:status": {"fails"}, "groupBy": {"g"}, "limit": {"10"}, "page": {"2"},
				"requiredColumns": {"CSP,status"}},
			response: `[{"status":"fails"}]`},
	})
}
