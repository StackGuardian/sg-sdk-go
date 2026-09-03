package cli

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkflows(t *testing.T) {
	base := "/api/v1/orgs/demo-org/wfgrps/g/wfs/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"workflows", "create", "--wfgrp", "g", "-b", `{"ResourceName":"w"}`},
			method: "POST", path: base, body: `{"ResourceName":"w"}`},
		{name: "get", args: []string{"workflows", "get", "w", "--wfgrp", "g"}, method: "GET", path: base + "w/"},
		{name: "get nested group", args: []string{"wf", "get", "w", "--wfgrp", "parent/child"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfgrps/parent/child/wfs/w/"},
		{name: "delete", args: []string{"workflows", "delete", "w", "--wfgrp", "g"}, method: "DELETE", path: base + "w/"},
		{name: "update", args: []string{"workflows", "update", "w", "--wfgrp", "g", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "w/", body: `{"Description":"d"}`},
		{name: "update upgrade mode", args: []string{"workflows", "update", "w", "--wfgrp", "g", "--upgrade-mode", "MANUAL", "-b", `{"Description":"d"}`},
			method: "PATCH", path: base + "w/", query: url.Values{"upgradeMode": {"MANUAL"}}, body: `{"Description":"d"}`},
		{name: "artifacts", args: []string{"workflows", "artifacts", "list", "--wf", "w", "--wfgrp", "g"}, method: "GET", path: base + "w/listall_artifacts/"},
		{name: "outputs", args: []string{"workflows", "outputs", "w", "--wfgrp", "g"}, method: "GET", path: base + "w/outputs/"},
		{name: "tfstate url", args: []string{"workflows", "tfstate-upload-url", "w", "--wfgrp", "g", "--filename", "s.json"},
			method: "GET", path: base + "w/tfstate_upload_url/", query: url.Values{"filename": {"s.json"}}},
		{name: "vcs triggers", args: []string{"workflows", "vcs-triggers", "w", "--wfgrp", "g", "-b", `{}`},
			method: "POST", path: base + "w/webhooks/vcs_triggers/", body: `{"VCSConfig":null,"VCSTriggers":null}`},
		{name: "list", args: []string{"workflows", "list", "--wfgrp", "g", "--description", "d", "--iac-template-id", "i", "--latest-wf-run-statuses", "3",
			"--resource-names", "n", "--runner-names", "r", "--tags", "t", "--last-evaluated-key", "k", "--limit", "5"},
			method: "GET", path: base + "listall/",
			query: url.Values{"Description": {"d"}, "IACTemplateId": {"i"}, "LatestWfRunStatuses": {"3"}, "ResourceNames": {"n"},
				"RunnerNames": {"r"}, "Tags": {"t"}, "lastevaluatedkey": {"k"}, "limit": {"5"}}},
	})
	t.Run("invalid upgrade mode", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "workflows", "update", "w", "--wfgrp", "g", "--upgrade-mode", "NOPE", "-b", `{}`)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--upgrade-mode")
	})
}

func TestWorkflowRuns(t *testing.T) {
	scope := []string{"--wfgrp", "g", "--wf", "w"}
	base := "/api/v1/orgs/demo-org/wfgrps/g/wfs/w/wfruns/"
	with := func(args ...string) []string { return append(args, scope...) }
	runCases(t, []cmdCase{
		{name: "create empty", args: with("workflow-runs", "create"), method: "POST", path: base, body: `{}`},
		{name: "create body", args: with("wfruns", "create", "-b", `{"WfType":"TERRAFORM"}`),
			method: "POST", path: base, body: `{"WfType":"TERRAFORM"}`},
		{name: "get", args: with("workflow-runs", "get", "r1"), method: "GET", path: base + "r1/"},
		{name: "update", args: with("workflow-runs", "update", "r1", "-b", `{"WfType":"CUSTOM"}`),
			method: "PATCH", path: base + "r1/", body: `{"WfType":"CUSTOM"}`},
		{name: "cancel", args: with("workflow-runs", "cancel", "r1"), method: "PATCH", path: base + "r1/cancel/"},
		{name: "logs", args: with("workflow-runs", "logs", "r1"), method: "GET", path: base + "r1/logs/"},
		{name: "approve", args: with("workflow-runs", "approve", "r1", "--message", "ok"),
			method: "POST", path: base + "r1/resume/", body: `{"Approve":true,"Message":"ok"}`},
		{name: "reject", args: with("workflow-runs", "approve", "r1", "--reject", "--approval-step", "plan"),
			method: "POST", path: base + "r1/resume/", body: `{"Approve":false,"ApprovalStep":"plan"}`},
		{name: "list", args: with("workflow-runs", "list", "--context-tags", "c", "--limit", "5", "--last-evaluated-key", "k"),
			method: "GET", path: base + "listall/", query: url.Values{"contextTags": {"c"}, "limit": {"5"}, "lastevaluatedkey": {"k"}}},
	})
}

func TestWorkflowRunFacts(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "get", args: []string{"workflow-run-facts", "get", "f1", "--wfgrp", "g", "--wf", "w", "--wf-run", "r1"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfgrps/g/wfs/w/wfruns/r1/wfrunfacts/f1/", response: `{"facts":1}`},
	})
}

func TestTemplatePackages(t *testing.T) {
	runCases(t, []cmdCase{
		// workflow step templates
		{name: "step get", args: []string{"workflow-step-templates", "get", "t1"},
			method: "GET", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1/"},
		{name: "step create", args: []string{"workflow-step-templates", "create", "-b", `{"TemplateName":"t1"}`},
			method: "POST", path: "/api/v1/templates/", body: `{"TemplateType":"WORKFLOW_STEP","TemplateName":"t1"}`},
		{name: "step create no revision", args: []string{"workflow-step-templates", "create", "--create-first-revision=false", "-b", `{"TemplateName":"t1"}`},
			method: "POST", path: "/api/v1/templates/", query: url.Values{"createFirstRevision": {"false"}}, body: `{"TemplateType":"WORKFLOW_STEP","TemplateName":"t1"}`},
		{name: "step update", args: []string{"workflow-step-templates", "update", "t1", "-b", `{"ShortDescription":"d"}`},
			method: "PATCH", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1/", body: `{"ShortDescription":"d"}`},
		{name: "step delete", args: []string{"workflow-step-templates", "delete", "t1"},
			method: "DELETE", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1/"},
		// workflow step template revisions
		{name: "step rev create", args: []string{"workflow-step-template-revisions", "create", "t1", "-b", `{"LongDescription":"d"}`},
			method: "POST", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1/revisions/", body: `{"TemplateType":"WORKFLOW_STEP","LongDescription":"d"}`},
		{name: "step rev update", args: []string{"workflow-step-template-revisions", "update", "t1:1", "-b", `{"LongDescription":"d"}`},
			method: "PATCH", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1:1/", body: `{"LongDescription":"d"}`},
		{name: "step rev get", args: []string{"workflow-step-template-revisions", "get", "t1:1"},
			method: "GET", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1:1/"},
		{name: "step rev delete", args: []string{"workflow-step-template-revisions", "delete", "t1:1", "--keep-parent-template"},
			method: "DELETE", path: "/api/v1/templatetypes/WORKFLOW_STEP/demo-org/t1:1/", query: url.Values{"keepParentTemplate": {"true"}}},
		// workflow templates
		{name: "wf tpl get", args: []string{"workflow-templates", "get", "t1"},
			method: "GET", path: "/api/v1/templatetypes/IAC/demo-org/t1/"},
		{name: "wf tpl create", args: []string{"workflow-templates", "create", "-b", `{"TemplateName":"t1"}`},
			method: "POST", path: "/api/v1/templates/", body: `{"TemplateType":"IAC","TemplateName":"t1"}`},
		{name: "wf tpl update", args: []string{"workflow-templates", "update", "t1", "-b", `{"TemplateName":"t2"}`},
			method: "PATCH", path: "/api/v1/templatetypes/IAC/demo-org/t1/", body: `{"TemplateName":"t2"}`},
		{name: "wf tpl delete", args: []string{"workflow-templates", "delete", "t1"},
			method: "DELETE", path: "/api/v1/templatetypes/IAC/demo-org/t1/"},
		// workflow template revisions
		{name: "wf rev create", args: []string{"workflow-template-revisions", "create", "t1", "-b", `{"LongDescription":"d"}`},
			method: "POST", path: "/api/v1/templatetypes/IAC/demo-org/t1/revisions/", body: `{"TemplateType":"IAC","LongDescription":"d"}`},
		{name: "wf rev delete", args: []string{"workflow-template-revisions", "delete", "t1:1"},
			method: "DELETE", path: "/api/v1/templatetypes/IAC/demo-org/t1:1/"},
		{name: "wf rev get", args: []string{"workflow-template-revisions", "get", "t1:1"},
			method: "GET", path: "/api/v1/templatetypes/IAC/demo-org/t1:1/"},
		{name: "wf rev update", args: []string{"workflow-template-revisions", "update", "t1:1", "-b", `{"Alias":"a"}`},
			method: "PATCH", path: "/api/v1/templatetypes/IAC/demo-org/t1:1/", body: `{"Alias":"a"}`},
		// stack templates
		{name: "stack tpl get", args: []string{"stack-templates", "get", "t1"},
			method: "GET", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1/"},
		{name: "stack tpl create", args: []string{"stack-templates", "create", "-b", `{"OwnerOrg":"demo-org"}`},
			method: "POST", path: "/api/v1/templates/", body: `{"TemplateType":"IAC_GROUP","OwnerOrg":"demo-org"}`},
		{name: "stack tpl update", args: []string{"stack-templates", "update", "t1", "-b", `{"Tags":["a"]}`},
			method: "PATCH", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1/", body: `{"Tags":["a"]}`},
		{name: "stack tpl delete", args: []string{"stack-templates", "delete", "t1"},
			method: "DELETE", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1/"},
		// stack template revisions
		{name: "stack rev create", args: []string{"stack-template-revisions", "create", "t1", "-b", `{"Tags":["a"]}`},
			method: "POST", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1/revisions/", body: `{"TemplateType":"IAC_GROUP","Tags":["a"]}`},
		{name: "stack rev get", args: []string{"stack-template-revisions", "get", "t1:1"},
			method: "GET", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1:1/"},
		{name: "stack rev update", args: []string{"stack-template-revisions", "update", "t1:1", "-b", `{"LongDescription":"d"}`},
			method: "PATCH", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1:1/", body: `{"LongDescription":"d"}`},
		{name: "stack rev delete", args: []string{"stack-template-revisions", "delete", "t1:1", "--keep-parent-template"},
			method: "DELETE", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t1:1/", query: url.Values{"keepParentTemplate": {"true"}}},
	})
}
