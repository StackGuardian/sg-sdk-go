package cli

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganizationsManagement(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "create", args: []string{"organizations", "create", "-b", `{"ResourceName":"neworg","selectedOnboardingGoal":"gitops"}`},
			method: "POST", path: "/api/v1/orgs/", body: `{"ResourceName":"neworg","selectedOnboardingGoal":"gitops"}`},
		{name: "list", args: []string{"organizations", "list"}, method: "GET", path: "/api/v1/orgs/listall/", response: `{"msg":["a","b"]}`},
		{name: "update", args: []string{"organizations", "update", "-b", `{"ExtraSgCreditsLimit":5}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/", body: `{"ExtraSgCreditsLimit":5}`},
		{name: "delete", args: []string{"organizations", "delete"}, method: "DELETE", path: "/api/v1/orgs/demo-org/"},
		{name: "workflows", args: []string{"organizations", "workflows", "--limit", "5", "--fetch-sg-owned-wfs=false", "--resource-names", "n", "--is-active", "1"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfs/listall/",
			query: url.Values{"limit": {"5"}, "fetchSgOwnedWfs": {"false"}, "resourceNames": {"n"}, "IsActive": {"1"}}, response: `{"msg":[]}`},
		{name: "count workflows", args: []string{"organizations", "count-workflows", "--tags", "t"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfs/listall/", query: url.Values{"onlyCount": {"1"}, "tags": {"t"}}, response: `{"msg":3}`},
		{name: "bulk action", args: []string{"organizations", "bulk-action", "--action", "activate", "--resource-ids", "/wfgrps/g/wfs/w,/wfgrps/g/stacks/s"},
			method: "POST", path: "/api/v1/orgs/demo-org/wfs/bulk-action/", body: `{"action":"activate","resourceIds":["/wfgrps/g/wfs/w","/wfgrps/g/stacks/s"]}`},
	})
	t.Run("create works without org", func(t *testing.T) {
		f := newFakeAPI(t)
		t.Setenv(envOrg, "")
		root := NewRootCmd()
		var out bytes.Buffer
		root.SetOut(&out)
		root.SetArgs([]string{"--base-url", f.srv.URL, "--api-key", "k", "organizations", "list"})
		require.NoError(t, root.Execute())
		assert.Equal(t, "/api/v1/orgs/listall/", f.last(t).Path)
	})
	t.Run("invalid bulk action", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "organizations", "bulk-action", "--action", "nuke", "--resource-ids", "/a")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--action")
	})
}

func TestAccessExtras(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "role binding", args: []string{"role-bindings", "get", "default"}, method: "GET", path: "/api/v1/orgs/demo-org/rolebindings/default/"},
		{name: "api token create", args: []string{"api-tokens", "create", "--regenerate"},
			method: "POST", path: "/api/v1/orgs/demo-org/api_token/", body: `{"regenerate":true}`, response: `{"msg":"API Key exists","data":"key"}`},
		{name: "api token runner group", args: []string{"api-tokens", "create", "--runner-group-id", "/runnergroups/rg"},
			method: "POST", path: "/api/v1/orgs/demo-org/api_token/", body: `{"runnerGroupId":"/runnergroups/rg"}`},
		{name: "api token delete", args: []string{"api-tokens", "delete", "pool/idp/u@x"},
			method: "DELETE", path: "/api/v1/orgs/demo-org/api_token/", body: `"pool/idp/u@x"`},
	})
}

func TestRunnerGroupExtras(t *testing.T) {
	base := "/api/v1/orgs/demo-org/runnergroups/"
	runCases(t, []cmdCase{
		{name: "list", args: []string{"runner-groups", "list", "--limit", "5", "--resource-names", "n"},
			method: "GET", path: base + "listall/", query: url.Values{"limit": {"5"}, "filterresourcenames": {"n"}}, response: `{"msg":[]}`},
		{name: "register", args: []string{"runner-groups", "register", "rg"}, method: "POST", path: base + "rg/register/"},
		{name: "storage auth", args: []string{"runner-groups", "storage-backend-auth", "rg"}, method: "POST", path: base + "rg/storage-backend-auth/"},
	})
}

func TestConnectorExtras(t *testing.T) {
	base := "/api/v1/orgs/demo-org/integrations/"
	runCases(t, []cmdCase{
		{name: "authenticate", args: []string{"connectors", "authenticate", "c"}, method: "GET", path: base + "c/authenticate/"},
		{name: "list accounts", args: []string{"connectors", "list-accounts", "-b", `{"ResourceId":"g"}`},
			method: "POST", path: base + "list_accounts/", body: `{"ResourceId":"g"}`, response: `{"msg":[{"Id":"1"}]}`},
		{name: "github repos", args: []string{"connectors", "github-repos", "c"}, method: "GET", path: base + "c/get_githubcom_repos/", response: `{"msg":["o/r"]}`},
		{name: "repos", args: []string{"connectors", "repos", "c", "--login", "l", "--installation-id", "i", "--page", "2", "--limit", "5", "--search-query", "q"},
			method: "GET", path: base + "c/repos/", query: url.Values{"login": {"l"}, "installationId": {"i"}, "page": {"2"}, "limit": {"5"}, "searchQuery": {"q"}}},
		{name: "child authenticate", args: []string{"connector-groups", "connectors", "authenticate", "c", "--group", "g"},
			method: "GET", path: "/api/v1/orgs/demo-org/integrationgroups/g/integrations/c/authenticate/"},
		{name: "discovery scan", args: []string{"connector-groups", "discovery-scan", "g", "--benchmark", "inventory,cis"},
			method: "POST", path: "/api/v1/orgs/demo-org/integrationgroups/g/discovery_scan/", query: url.Values{"benchmark": {"inventory,cis"}}},
	})
}

func TestSecretsBulk(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "read bulk", args: []string{"secrets", "read-bulk", "--names", "a,b"},
			method: "POST", path: "/api/v1/orgs/demo-org/secrets/listall/", body: `{"ResourceNames":["a","b"]}`, response: `{"msg":[]}`},
	})
}

func TestTemplateExtras(t *testing.T) {
	orgHeader := map[string]string{"X-Sg-Orgid": "demo-org"}
	runCases(t, []cmdCase{
		{name: "subscription create", args: []string{"templates", "subscriptions", "create", "-b", `{"IACSubscriptions":{},"WfStepSubscriptions":{},"PolicySubscriptions":{},"IACGroupSubscriptions":{}}`},
			method: "POST", path: "/api/v1/orgs/demo-org/subscriptions/", body: `{"IACSubscriptions":{},"WfStepSubscriptions":{},"PolicySubscriptions":{},"IACGroupSubscriptions":{}}`},
		{name: "subscription update", args: []string{"templates", "subscriptions", "update", "default", "--action", "subscribe", "-b", `{"IACSubscriptions":{"/o/t":{}}}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/subscriptions/default/subscribe/", body: `{"IACSubscriptions":{"/o/t":{}}}`},
		{name: "artifacts list", args: []string{"templates", "artifacts", "list", "t", "--type", "IAC", "--prefix", "p", "--max-keys", "5", "--start-after-key", "k"},
			method: "GET", path: "/api/v1/templatetypes/IAC/demo-org/t/listall_artifacts/", headers: orgHeader,
			query: url.Values{"artifactPrefix": {"p"}, "maxKeys": {"5"}, "startAfterKey": {"k"}}},
		{name: "artifact download url", args: []string{"templates", "artifacts", "download-url", "t", "a/b.zip", "--type", "IAC", "--owner-org", "stackguardian"},
			method: "GET", path: "/api/v1/templatetypes/IAC/stackguardian/t/artifacts/a/b.zip/download/presigned-url/", headers: orgHeader},
		{name: "artifact upload url", args: []string{"templates", "artifacts", "upload-url", "t", "a", "--type", "IAC_GROUP"},
			method: "GET", path: "/api/v1/templatetypes/IAC_GROUP/demo-org/t/artifacts/a/upload/presigned-url/", headers: orgHeader},
		{name: "artifact delete", args: []string{"templates", "artifacts", "delete", "t", "a", "--type", "IAC"},
			method: "DELETE", path: "/api/v1/templatetypes/IAC/demo-org/t/artifacts/a/", headers: orgHeader},
		{name: "input schema", args: []string{"templates", "input-schema", "t:2", "--type", "IAC", "--schema-type", "FORM_JSONSCHEMA"},
			method: "GET", path: "/api/v1/templatetypes/IAC/demo-org/t:2/get_input_schema/", query: url.Values{"SchemaType": {"FORM_JSONSCHEMA"}}},
		{name: "vcs triggers", args: []string{"templates", "vcs-triggers", "t", "--type", "IAC", "-b", `{}`},
			method: "POST", path: "/api/v1/templatetypes/IAC/demo-org/t/webhooks/vcs_triggers/", body: `{"VCSConfig":null,"VCSTriggers":null}`},
		{name: "public list", args: []string{"templates", "public", "list", "--type", "IAC", "--owner-orgs", "stackguardian", "--search-query", "q", "--limit", "5"},
			method: "GET", path: "/api/v1/public/templatetypes/IAC/templates/listall/",
			query: url.Values{"OwnerOrgs": {"stackguardian"}, "SearchQuery": {"q"}, "limit": {"5"}}},
		{name: "public list all", args: []string{"templates", "public", "list-all", "--template-types", "IAC,IAC_GROUP", "--tags", "a"},
			method: "GET", path: "/api/v1/public/templates/listall/", query: url.Values{"templateTypes": {"IAC,IAC_GROUP"}, "tags": {"a"}}},
		{name: "public get", args: []string{"templates", "public", "get", "t:1", "--type", "IAC", "--owner-org", "stackguardian"},
			method: "GET", path: "/api/v1/public/templatetypes/IAC/stackguardian/t:1/"},
	})
}

func TestWorkflowGroupResources(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "list resources", args: []string{"workflow-groups", "list-resources", "parent/child", "--resource-types", "WORKFLOW,STACK", "--limit", "5", "--context-tags", "env:prod"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfgrps/parent/child/listall/",
			query: url.Values{"resourceTypes": {"WORKFLOW,STACK"}, "limit": {"5"}, "contextTags": {"env:prod"}}, response: `{"msg":[]}`},
	})
}

func TestStackCompare(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "compare flags", args: []string{"stacks", "compare", "s", "--wfgrp", "g", "--target-template-group-id", "/o/t", "--upgrade-mode", "MANUAL"},
			method: "POST", path: "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/compare/", body: `{"targetTemplateGroupId":"/o/t","upgradeMode":"MANUAL"}`},
		{name: "compare body", args: []string{"stacks", "compare", "s", "--wfgrp", "g", "-b", `{"targetTemplateGroupId":"/o/t","patchData":{"a":1}}`},
			method: "POST", path: "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/compare/", body: `{"targetTemplateGroupId":"/o/t","patchData":{"a":1}}`},
	})
}

func TestWorkflowExtras(t *testing.T) {
	base := "/api/v1/orgs/demo-org/wfgrps/g/wfs/w/"
	scope := []string{"--wfgrp", "g", "--wf", "w"}
	with := func(args ...string) []string { return append(args, scope...) }
	runCases(t, []cmdCase{
		{name: "file upload url", args: []string{"workflows", "file-upload-url", "w", "--wfgrp", "g", "--filename", "f", "--folder", "d"},
			method: "GET", path: base + "file_upload_url/", query: url.Values{"filename": {"f"}, "folder": {"d"}}},
		{name: "compare", args: []string{"workflows", "compare", "w", "--wfgrp", "g", "-b", `{"targetTemplateId":"/o/t"}`},
			method: "POST", path: base + "compare/", body: `{"targetTemplateId":"/o/t"}`},
		{name: "compare flags", args: []string{"workflows", "compare", "w", "--wfgrp", "g", "--target-template-id", "/o/t"},
			method: "POST", path: base + "compare/", body: `{"targetTemplateId":"/o/t"}`},
		{name: "artifact url", args: with("workflows", "artifacts", "get-url", "--artifact-path", "p", "--version-id", "v"),
			method: "GET", path: base + "get_artifact/", query: url.Values{"artifactPath": {"p"}, "versionId": {"v"}}},
		{name: "artifact get", args: with("workflows", "artifacts", "get", "a/b.json"),
			method: "GET", path: base + "artifacts/a/b.json/", response: `{"version":4}`},
		{name: "artifact create", args: with("workflows", "artifacts", "create", "tf", "-b", `{"version":4}`),
			method: "POST", path: base + "artifacts/tf/", body: `{"version":4}`},
		{name: "artifact delete", args: with("workflows", "artifacts", "delete", "tf"), method: "DELETE", path: base + "artifacts/tf/"},
		{name: "artifact versions", args: with("workflows", "artifacts", "versions", "tf", "--limit", "5"),
			method: "GET", path: base + "artifacts/tf/versions/", query: url.Values{"limit": {"5"}}},
		{name: "artifact rollback", args: with("workflows", "artifacts", "rollback", "tf", "--version-id", "v"),
			method: "POST", path: base + "artifacts/tf/rollback/", body: `{"VersionId":"v"}`},
		{name: "artifact lock", args: with("workflows", "artifacts", "lock", "tf"), method: "POST", path: base + "artifacts/tf/lock/", response: `{"ID":"x"}`},
		{name: "artifact unlock", args: with("workflows", "artifacts", "unlock", "tf"), method: "DELETE", path: base + "artifacts/tf/lock/", response: `{"ID":"x"}`},
	})
}

func TestStackWorkflowExtras(t *testing.T) {
	base := "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/"
	scope := []string{"--wfgrp", "g", "--stack", "s"}
	with := func(args ...string) []string { return append(args, scope...) }
	art := func(args ...string) []string { return append(append(args, scope...), "--wf", "w") }
	runCases(t, []cmdCase{
		{name: "create", args: with("stack-workflows", "create", "-b", `{"ResourceName":"w"}`), method: "POST", path: base, body: `{"ResourceName":"w"}`},
		{name: "file upload url", args: with("stack-workflows", "file-upload-url", "w", "--filename", "f"),
			method: "GET", path: base + "w/file_upload_url/", query: url.Values{"filename": {"f"}}},
		{name: "compare", args: with("stack-workflows", "compare", "w", "--target-template-id", "/o/t"),
			method: "POST", path: base + "w/compare/", body: `{"targetTemplateId":"/o/t"}`},
		{name: "vcs triggers", args: with("stack-workflows", "vcs-triggers", "w", "-b", `{}`),
			method: "POST", path: base + "w/webhooks/vcs_triggers/", body: `{"VCSConfig":null,"VCSTriggers":null}`},
		{name: "artifact url", args: art("stack-workflows", "artifacts", "get-url", "--artifact-path", "p"),
			method: "GET", path: base + "w/get_artifact/", query: url.Values{"artifactPath": {"p"}}},
		{name: "artifact get", args: art("stack-workflows", "artifacts", "get", "tf"), method: "GET", path: base + "w/artifacts/tf/", response: `{"v":1}`},
		{name: "artifact create", args: art("stack-workflows", "artifacts", "create", "tf", "-b", `[1,2]`),
			method: "POST", path: base + "w/artifacts/tf/", body: `[1,2]`},
		{name: "artifact delete", args: art("stack-workflows", "artifacts", "delete", "tf"), method: "DELETE", path: base + "w/artifacts/tf/"},
		{name: "artifact lock", args: art("stack-workflows", "artifacts", "lock", "tf"), method: "POST", path: base + "w/artifacts/tf/lock/", response: `{"ID":"x"}`},
		{name: "artifact unlock", args: art("stack-workflows", "artifacts", "unlock", "tf"), method: "DELETE", path: base + "w/artifacts/tf/lock/", response: `{"ID":"x"}`},
	})
}

func TestRunExtras(t *testing.T) {
	stackBase := "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/w/wfruns/"
	stackScope := []string{"--wfgrp", "g", "--stack", "s", "--wf", "w"}
	ss := func(args ...string) []string { return append(args, stackScope...) }
	runCases(t, []cmdCase{
		{name: "stack list", args: ss("stack-workflow-runs", "list", "--latest-status", "RUNNING", "--limit", "5", "--context-tags", "env:prod"),
			method: "GET", path: stackBase + "listall/", query: url.Values{"latestStatus": {"RUNNING"}, "limit": {"5"}, "contextTags": {"env:prod"}}, response: `{"msg":[]}`},
		{name: "stack update", args: ss("stack-workflow-runs", "update", "r1", "-b", `{"WfType":"CUSTOM"}`),
			method: "PATCH", path: stackBase + "r1/", body: `{"WfType":"CUSTOM"}`},
		{name: "stack delete", args: ss("stack-workflow-runs", "delete", "r1"), method: "DELETE", path: stackBase + "r1/"},
		{name: "delete", args: []string{"workflow-runs", "delete", "r1", "--wfgrp", "g", "--wf", "w"},
			method: "DELETE", path: "/api/v1/orgs/demo-org/wfgrps/g/wfs/w/wfruns/r1/"},
		{name: "get by ksuid", args: []string{"workflow-runs", "get-by-ksuid", "rk", "--parent-ksuid", "pk"},
			method: "GET", path: "/api/v1/orgs/demo-org/wfs/pk/wfruns/rk/"},
		{name: "update by ksuid", args: []string{"workflow-runs", "update-by-ksuid", "rk", "--parent-ksuid", "pk", "-b", `{"WfType":"CUSTOM"}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/wfs/pk/wfruns/rk/", body: `{"WfType":"CUSTOM"}`},
		{name: "facts create", args: []string{"workflow-run-facts", "create", "f1", "--wfgrp", "g", "--wf", "w", "--wf-run", "r1", "-b", `{"TfDrift":{"a":1}}`},
			method: "POST", path: "/api/v1/orgs/demo-org/wfgrps/g/wfs/w/wfruns/r1/wfrunfacts/f1/", body: `{"TfDrift":{"a":1}}`, response: `{"msg":"ok"}`},
		{name: "facts update", args: []string{"workflow-run-facts", "update", "f1", "--wfgrp", "g", "--wf", "w", "--wf-run", "r1", "-b", `{"Tags":["a"]}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/wfgrps/g/wfs/w/wfruns/r1/wfrunfacts/f1/", body: `{"Tags":["a"]}`, response: `{"msg":"ok"}`},
		{name: "stack facts create", args: []string{"stack-workflow-run-facts", "create", "f1", "--wfgrp", "g", "--stack", "s", "--wf", "w", "--wf-run", "r1", "-b", `{"TfDrift":{}}`},
			method: "POST", path: "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/w/wfruns/r1/wfrunfacts/f1/", body: `{"TfDrift":{}}`, response: `{"msg":"ok"}`},
		{name: "stack facts update", args: []string{"stack-workflow-run-facts", "update", "f1", "--wfgrp", "g", "--stack", "s", "--wf", "w", "--wf-run", "r1", "-b", `{"IsActive":"0"}`},
			method: "PATCH", path: "/api/v1/orgs/demo-org/wfgrps/g/stacks/s/wfs/w/wfruns/r1/wfrunfacts/f1/", body: `{"IsActive":"0"}`, response: `{"msg":"ok"}`},
	})
}

func TestStateBackends(t *testing.T) {
	base := "/api/v1/orgs/demo-org/statebackends/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"state-backends", "create", "-b", `{"ResourceName":"sb","Statefiles":[],"StateBackendConfig":{"type":"aws_s3","s3BucketName":"b","auth":{"integrationId":"/integrations/aws"}}}`},
			method: "POST", path: base, body: `{"ResourceName":"sb","Statefiles":[],"StateBackendConfig":{"type":"aws_s3","s3BucketName":"b","auth":{"integrationId":"/integrations/aws"}}}`},
		{name: "get", args: []string{"state-backends", "get", "sb"}, method: "GET", path: base + "sb/"},
		{name: "update", args: []string{"state-backends", "update", "sb", "-b", `{"Description":"d","Tags":null}`},
			method: "PATCH", path: base + "sb/", body: `{"Description":"d","Tags":null}`},
		{name: "delete", args: []string{"state-backends", "delete", "sb"}, method: "DELETE", path: base + "sb/"},
		{name: "list", args: []string{"state-backends", "list", "--search-query", "q", "--limit", "5"},
			method: "GET", path: base + "listall/", query: url.Values{"searchQuery": {"q"}, "limit": {"5"}}},
		{name: "statefiles", args: []string{"state-backends", "list-statefiles", "--config", "abc"},
			method: "GET", path: base + "list_statefiles/", query: url.Values{"config": {"abc"}}},
	})
}

func TestResources(t *testing.T) {
	runCases(t, []cmdCase{
		{name: "move", args: []string{"resources", "move", "--source-ids", "/a,/b", "--destination", "/wfgrps/g"},
			method: "POST", path: "/api/v1/orgs/demo-org/resources/move/", body: `{"SourceResourceIds":["/a","/b"],"DestinationParentId":"/wfgrps/g"}`},
		{name: "search", args: []string{"resources", "search", "--search-query", "q", "--limit", "5", "--resource-types", "WORKFLOW,STACK"},
			method: "POST", path: "/api/v1/orgs/demo-org/search/", body: `{"SearchQuery":"q","Limit":5,"ResourceTypes":["WORKFLOW","STACK"]}`},
		{name: "search body", args: []string{"resources", "search", "-b", `{"ResourceTypes":["WORKFLOW"],"Tags":["a"]}`},
			method: "POST", path: "/api/v1/orgs/demo-org/search/", body: `{"ResourceTypes":["WORKFLOW"],"Tags":["a"]}`},
		{name: "tags", args: []string{"resources", "tags", "--context-tag-key", "env", "--limit", "5"},
			method: "GET", path: "/api/v1/orgs/demo-org/search/tags/", query: url.Values{"ContextTagKey": {"env"}, "Limit": {"5"}}},
	})
}

func TestBilling(t *testing.T) {
	base := "/api/v1/orgs/demo-org/billing/"
	runCases(t, []cmdCase{
		{name: "balance", args: []string{"billing", "balance"}, method: "GET", path: base + "list-balance/"},
		{name: "dashboard", args: []string{"billing", "dashboard-url", "--type", "usage"}, method: "GET", path: base + "dashboard/", query: url.Values{"type": {"usage"}}},
		{name: "invoices", args: []string{"billing", "invoices", "--limit", "5", "--next-page", "c"},
			method: "GET", path: base + "invoices/", query: url.Values{"limit": {"5"}, "next_page": {"c"}}},
		{name: "details", args: []string{"billing", "details"}, method: "GET", path: base + "details/"},
		{name: "change plan", args: []string{"billing", "change-plan", "--target-package", "SG | Standard"},
			method: "POST", path: base + "change-plan/", body: `{"TargetPackage":"SG | Standard"}`},
		{name: "profile setup", args: []string{"billing", "profile", "setup", "-b", `{"Email":"e@x"}`},
			method: "POST", path: base + "profile/", body: `{"Email":"e@x"}`},
		{name: "profile update", args: []string{"billing", "profile", "update", "-b", `{"Name":"n"}`},
			method: "PATCH", path: base + "profile/", body: `{"Name":"n"}`},
		{name: "payment list", args: []string{"billing", "payment-methods", "list"}, method: "GET", path: base + "payment-methods/"},
		{name: "payment intent", args: []string{"billing", "payment-methods", "setup-intent"}, method: "POST", path: base + "payment-methods/"},
		{name: "payment default", args: []string{"billing", "payment-methods", "set-default", "pm_1"},
			method: "POST", path: base + "set-default-payment-method/", body: `{"PaymentMethodId":"pm_1"}`},
		{name: "payment detach", args: []string{"billing", "payment-methods", "detach", "pm_1"},
			method: "POST", path: base + "detach-card/", body: `{"PaymentMethodId":"pm_1"}`},
	})
	t.Run("invalid dashboard type", func(t *testing.T) {
		_, _, err := execCLI(t, newFakeAPI(t), "", "billing", "dashboard-url", "--type", "nope")
		require.Error(t, err)
	})
}

func TestChats(t *testing.T) {
	base := "/api/v1/orgs/demo-org/chats/"
	runCases(t, []cmdCase{
		{name: "create", args: []string{"chats", "create", "-b", `{"ChatType":"kyro","Config":{}}`},
			method: "POST", path: base, body: `{"ChatType":"kyro","Config":{}}`},
		{name: "list", args: []string{"chats", "list", "--chat-type", "kyro", "--refresh-pr-status", "--limit", "5"},
			method: "GET", path: base + "listall/", query: url.Values{"chatType": {"kyro"}, "refreshPrStatus": {"true"}, "limit": {"5"}}, response: `{"msg":[]}`},
		{name: "get", args: []string{"chats", "get", "c1", "--refresh-pr-status"}, method: "GET", path: base + "c1/", query: url.Values{"refreshPrStatus": {"true"}}},
		{name: "update", args: []string{"chats", "update", "c1", "-b", `{"Status":"archived"}`}, method: "PATCH", path: base + "c1/", body: `{"Status":"archived"}`},
		{name: "delete", args: []string{"chats", "delete", "c1"}, method: "DELETE", path: base + "c1/"},
		{name: "artifacts list", args: []string{"chats", "artifacts", "list", "c1", "--prefix", "p"}, method: "GET", path: base + "c1/listall_artifacts/", query: url.Values{"prefix": {"p"}}},
		{name: "artifact url", args: []string{"chats", "artifacts", "get-url", "c1", "--artifact-path", "a"}, method: "GET", path: base + "c1/get_artifact/", query: url.Values{"artifactPath": {"a"}}},
		{name: "artifact upload url", args: []string{"chats", "artifacts", "upload-url", "c1", "--artifact-path", "a", "--content-type", "text/plain"},
			method: "POST", path: base + "c1/save_artifact/", query: url.Values{"artifactPath": {"a"}, "contentType": {"text/plain"}}},
		{name: "commit bundle", args: []string{"chats", "commit-bundle", "c1", "--expected-version", "3"},
			method: "POST", path: base + "c1/commit_bundle/", body: `{"expected_version":3}`},
		{name: "create pr", args: []string{"chats", "create-pr", "c1", "-b", `{"vcsConnector":"v","repoUrl":"https://x","reuseBranch":true}`},
			method: "POST", path: base + "c1/create_pr/", body: `{"vcsConnector":"v","repoUrl":"https://x","reuseBranch":true}`},
		{name: "pr status", args: []string{"chats", "pr-status", "c1"}, method: "GET", path: base + "c1/pr_status/"},
		{name: "pr sync", args: []string{"chats", "pr-sync-status", "c1"}, method: "GET", path: base + "c1/pr_sync_status/"},
		{name: "message create text", args: []string{"chats", "messages", "create", "--chat", "c1", "--text", "hi", "--client-request-id", "r1"},
			method: "POST", path: base + "c1/messages/", body: `{"content":[{"type":"text","text":"hi"}],"clientRequestId":"r1"}`},
		{name: "message create body", args: []string{"chats", "messages", "create", "--chat", "c1", "-b", `{"content":[{"type":"tool_result","toolUseId":"t","isError":true}]}`},
			method: "POST", path: base + "c1/messages/", body: `{"content":[{"type":"tool_result","toolUseId":"t","isError":true}]}`},
		{name: "messages list", args: []string{"chats", "messages", "list", "--chat", "c1", "--limit", "5"},
			method: "GET", path: base + "c1/messages/listall/", query: url.Values{"limit": {"5"}}, response: `{"msg":[]}`},
		{name: "message get", args: []string{"chats", "messages", "get", "m1", "--chat", "c1"}, method: "GET", path: base + "c1/messages/m1/"},
		{name: "message cancel", args: []string{"chats", "messages", "cancel", "m1", "--chat", "c1"}, method: "POST", path: base + "c1/messages/m1/cancel/"},
		{name: "message retry", args: []string{"chats", "messages", "retry", "m1", "--chat", "c1"}, method: "POST", path: base + "c1/messages/m1/retry/"},
	})
}
