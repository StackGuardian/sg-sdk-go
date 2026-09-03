package cli

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkflowGroups(t *testing.T) {
	base := "/api/v1/orgs/demo-org/wfgrps/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"workflow-groups", "create", "-b", `{"ResourceName":"g"}`},
			method: "POST", path: base, body: `{"ResourceName":"g"}`},
		{name: "get", args: []string{"workflow-groups", "get", "g"}, method: "GET", path: base + "g/"},
		{name: "get nested", args: []string{"wfgrps", "get", "parent/child"}, method: "GET", path: base + "parent/child/"},
		{name: "delete", args: []string{"workflow-groups", "delete", "g"}, method: "DELETE", path: base + "g/"},
		{name: "update", args: []string{"workflow-groups", "update", "g", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "g/", body: `{"Description":"d"}`},
		{name: "create child", args: []string{"workflow-groups", "create-child", "g", "-b", `{"ResourceName":"c"}`},
			method: "POST", path: base + "g/wfgrps/", body: `{"ResourceName":"c"}`},
		{name: "list children", args: []string{"workflow-groups", "list-children", "g", "--description", "d", "--resource-names", "n", "--tags", "t", "--last-evaluated-key", "k", "--limit", "5"},
			method: "GET", path: base + "g/wfgrps/listall/",
			query: url.Values{"Description": {"d"}, "ResourceNames": {"n"}, "Tags": {"t"}, "lastevaluatedkey": {"k"}, "limit": {"5"}}},
		{name: "list", args: []string{"workflow-groups", "list", "--search-query", "q", "--limit", "5"},
			method: "GET", path: base + "listall/", query: url.Values{"searchQuery": {"q"}, "limit": {"5"}}},
	})
}

func TestStacks(t *testing.T) {
	base := "/api/v1/orgs/demo-org/wfgrps/g/stacks/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"stacks", "create", "--wfgrp", "g", "-b", `{"ResourceName":"s"}`},
			method: "POST", path: base, body: `{"ResourceName":"s"}`},
		{name: "create run on create", args: []string{"stacks", "create", "--wfgrp", "g", "--run-on-create", "-b", `{"ResourceName":"s"}`},
			method: "POST", path: base, query: url.Values{"runOnCreate": {"true"}}, body: `{"ResourceName":"s"}`},
		{name: "get", args: []string{"stacks", "get", "s", "--wfgrp", "g"}, method: "GET", path: base + "s/"},
		{name: "delete", args: []string{"stacks", "delete", "s", "--wfgrp", "g"}, method: "DELETE", path: base + "s/"},
		{name: "update", args: []string{"stacks", "update", "s", "--wfgrp", "g", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "s/", body: `{"Description":"d"}`},
		{name: "outputs", args: []string{"stacks", "outputs", "s", "--wfgrp", "g"}, method: "GET", path: base + "s/outputs/"},
		{name: "list", args: []string{"stacks", "list", "--wfgrp", "g", "--description", "d", "--resource-names", "n", "--tags", "t", "--last-evaluated-key", "k", "--limit", "5"},
			method: "GET", path: base + "listall/",
			query: url.Values{"Description": {"d"}, "ResourceNames": {"n"}, "Tags": {"t"}, "lastevaluatedkey": {"k"}, "limit": {"5"}}},
	})
	t.Run("wfgrp required", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "stacks", "get", "s")
		require.Error(t, err)
	})
}

func TestStackRuns(t *testing.T) {
	base := "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/stackruns/"
	runCases(t, []cmdCase{
		{name: "create action", args: []string{"stack-runs", "create", "--wfgrp", "g", "--stack", "s", "--action-type", "APPLY"},
			method: "POST", path: base, body: `{"ActionType":"APPLY"}`},
		{name: "create body", args: []string{"stack-runs", "create", "--wfgrp", "g", "--stack", "s", "-b", `{"ActionType":"DESTROY"}`},
			method: "POST", path: base, body: `{"ActionType":"DESTROY"}`},
		{name: "get", args: []string{"stack-runs", "get", "r1", "--wfgrp", "g", "--stack", "s"}, method: "GET", path: base + "r1/"},
		{name: "list", args: []string{"stack-runs", "list", "--wfgrp", "g", "--stack", "s", "--limit", "5", "--last-evaluated-key", "k"},
			method: "GET", path: base + "listall/", query: url.Values{"limit": {"5"}, "lastevaluatedkey": {"k"}}},
	})
	t.Run("action required", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "stack-runs", "create", "--wfgrp", "g", "--stack", "s")
		require.Error(t, err)
	})
}

func TestStackWorkflows(t *testing.T) {
	scope := []string{"--wfgrp", "g", "--stack", "s"}
	base := "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/"
	with := func(args ...string) []string { return append(args, scope...) }
	runCases(t, []cmdCase{
		{name: "get", args: with("stack-workflows", "get", "w"), method: "GET", path: base + "w/"},
		{name: "delete", args: with("stack-workflows", "delete", "w"), method: "DELETE", path: base + "w/"},
		{name: "update", args: with("stack-workflows", "update", "w", "-b", `{"Description":"d"}`),
			method: "PATCH", path: base + "w/", body: `{"Description":"d"}`},
		{name: "artifacts", args: with("stack-workflows", "artifacts", "list", "--wf", "w"), method: "GET", path: base + "w/listall_artifacts/"},
		{name: "outputs", args: with("stack-workflows", "outputs", "w"), method: "GET", path: base + "w/outputs/"},
		{name: "tfstate url", args: with("stack-workflows", "tfstate-upload-url", "w", "--filename", "state.json"),
			method: "GET", path: base + "w/tfstate_upload_url/", query: url.Values{"filename": {"state.json"}}},
		{name: "list", args: with("stack-workflows", "list", "--description", "d", "--iac-template-id", "i", "--latest-wf-run-statuses", "3",
			"--resource-names", "n", "--runner-names", "r", "--tags", "t", "--last-evaluated-key", "k", "--limit", "5"),
			method: "GET", path: base + "listall/",
			query: url.Values{"Description": {"d"}, "IACTemplateId": {"i"}, "LatestWfRunStatuses": {"3"}, "ResourceNames": {"n"},
				"RunnerNames": {"r"}, "Tags": {"t"}, "lastevaluatedkey": {"k"}, "limit": {"5"}}},
	})
}

func TestStackWorkflowRuns(t *testing.T) {
	scope := []string{"--wfgrp", "g", "--stack", "s", "--wf", "w"}
	base := "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/w/wfruns/"
	with := func(args ...string) []string { return append(args, scope...) }
	runCases(t, []cmdCase{
		{name: "create empty", args: with("stack-workflow-runs", "create"), method: "POST", path: base, body: `{}`},
		{name: "create body", args: with("stack-workflow-runs", "create", "-b", `{"WfType":"TERRAFORM"}`),
			method: "POST", path: base, body: `{"WfType":"TERRAFORM"}`},
		{name: "get", args: with("stack-workflow-runs", "get", "r1"), method: "GET", path: base + "r1/"},
		{name: "logs", args: with("stack-workflow-runs", "logs", "r1"), method: "GET", path: base + "r1/logs/"},
		{name: "approve", args: with("stack-workflow-runs", "approve", "r1", "--message", "ok", "--approval-step", "plan", "--reason", "policy"),
			method: "POST", path: base + "r1/resume/", body: `{"Approve":true,"Message":"ok","ApprovalStep":"plan","ReasonForApprovalRequired":"policy"}`},
		{name: "reject", args: with("stack-workflow-runs", "approve", "r1", "--reject"),
			method: "POST", path: base + "r1/resume/", body: `{"Approve":false}`},
	})
}

func TestStackWorkflowRunFacts(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "get", args: []string{"stack-workflow-run-facts", "get", "f1", "--wfgrp", "g", "--stack", "s", "--wf", "w", "--wf-run", "r1"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/w/wfruns/r1/wfrunfacts/f1/"},
	})
}
