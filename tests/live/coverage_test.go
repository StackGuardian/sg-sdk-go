// Package live holds integration tests that run against a real StackGuardian
// organization. They are skipped unless SG_API_TOKEN and SG_ORG are set:
//
//	SG_API_TOKEN=... SG_ORG=my-org go test ./tests/live/
//
// TestReadEndpointsDecode calls every read-only endpoint the organization has
// data for, fails when a live response does not decode into the SDK's typed
// response, and fails when the response carries fields the SDK type does not
// declare (which would otherwise be silently dropped by typed consumers).
package live

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"testing"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/billing"
	"github.com/StackGuardian/sg-sdk-go/chats"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/option"
	"github.com/StackGuardian/sg-sdk-go/organizations"
	"github.com/StackGuardian/sg-sdk-go/resources"
	"github.com/StackGuardian/sg-sdk-go/runnergroups"
	"github.com/StackGuardian/sg-sdk-go/statebackends"
	"github.com/StackGuardian/sg-sdk-go/templates"
	"github.com/StackGuardian/sg-sdk-go/workflowgroups"
)

// capture records the last raw response body so it can be compared with the typed decode.
type capture struct {
	body   []byte
	status int
}

func (c *capture) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	resp.Body = io.NopCloser(bytes.NewReader(b))
	c.body, c.status = b, resp.StatusCode
	return resp, nil
}

// undeclared collects keys present (with a non-empty value) in raw but absent from typed,
// descending into objects and the first element of arrays.
func undeclared(path string, raw, typed interface{}, out *[]string) {
	switch r := raw.(type) {
	case map[string]interface{}:
		t, _ := typed.(map[string]interface{})
		keys := make([]string, 0, len(r))
		for k := range r {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			tv, ok := t[k]
			if !ok {
				if !emptyValue(r[k]) {
					*out = append(*out, path+k)
				}
				continue
			}
			undeclared(path+k+".", r[k], tv, out)
		}
	case []interface{}:
		t, _ := typed.([]interface{})
		if len(r) > 0 && len(t) > 0 {
			undeclared(path+"[].", r[0], t[0], out)
		}
	}
}

func emptyValue(v interface{}) bool {
	switch x := v.(type) {
	case nil:
		return true
	case string:
		return x == ""
	case float64:
		return x == 0
	case bool:
		return !x
	case []interface{}:
		return len(x) == 0
	case map[string]interface{}:
		return len(x) == 0
	}
	return false
}

func trimPrefix(s *string, prefix string) string {
	if s == nil {
		return ""
	}
	return strings.TrimPrefix(*s, prefix)
}

func TestReadEndpointsDecode(t *testing.T) {
	token, org := os.Getenv("SG_API_TOKEN"), os.Getenv("SG_ORG")
	if token == "" || org == "" {
		t.Skip("SG_API_TOKEN and SG_ORG are required for live tests")
	}
	baseURL := os.Getenv("SG_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.app.stackguardian.io"
	}
	cap := &capture{}
	c := client.NewClient(
		option.WithApiKey("apikey "+strings.TrimPrefix(token, "apikey ")),
		option.WithBaseURL(baseURL),
		option.WithHTTPClient(&http.Client{Transport: cap}),
	)
	ctx := context.Background()
	limit := 2
	s := func(v string) *string { return &v }

	// Names discovered from list calls; dependent calls are skipped when empty.
	var wfGrp, wf, wfRun, conn, policy, role, access, chat, iacTmpl, stackTmpl string

	type call struct {
		name string
		need *string // skip when the discovered name is empty
		fn   func() (interface{}, error)
	}
	calls := []call{
		{"Organizations.ReadOrganization", nil, func() (interface{}, error) { return c.Organizations.ReadOrganization(ctx, org) }},
		{"Organizations.ListAllWorkflows", nil, func() (interface{}, error) {
			return c.Organizations.ListAllWorkflows(ctx, org, &organizations.ListAllOrganizationWorkflowsRequest{Limit: &limit})
		}},
		{"Organizations.CountWorkflows", nil, func() (interface{}, error) { return c.Organizations.CountWorkflows(ctx, org, nil) }},
		{"AccessManagement.ReadRoleBinding", nil, func() (interface{}, error) { return c.AccessManagement.ReadRoleBinding(ctx, org, "default") }},
		{"AccessManagement.ListAllUsers", nil, func() (interface{}, error) {
			return c.AccessManagement.ListAllUsers(ctx, org, &sgsdkgo.ListAllUsersRequest{})
		}},
		{"AccessManagement.ListAllRoles", nil, func() (interface{}, error) {
			r, err := c.AccessManagement.ListAllRoles(ctx, org, &sgsdkgo.ListAllRolesRequest{})
			if items, ok := r["msg"].([]interface{}); ok && len(items) > 0 {
				if item, ok := items[0].(map[string]interface{}); ok {
					if id, ok := item["ResourceId"].(string); ok {
						role = strings.TrimPrefix(id, "/roles/")
					}
				}
			}
			return r, err
		}},
		{"AccessManagement.ReadRole", &role, func() (interface{}, error) { return c.AccessManagement.ReadRole(ctx, org, role) }},
		{"AccessManagement.ListAllApiAccesses", nil, func() (interface{}, error) {
			r, err := c.AccessManagement.ListAllApiAccesses(ctx, org, &sgsdkgo.ListAllApiAccessesRequest{Limit: &limit})
			if err == nil && r != nil && len(r.Data) > 0 {
				access = strings.TrimPrefix(r.Data[0].ResourceId, "/apiaccesses/")
			}
			return r, err
		}},
		{"AccessManagement.ReadApiAccess", &access, func() (interface{}, error) { return c.AccessManagement.ReadApiAccess(ctx, access, org) }},
		{"AccessManagement.ReadAuditLogs", nil, func() (interface{}, error) {
			return c.AccessManagement.ReadAuditLogs(ctx, org, &sgsdkgo.ReadAuditLogsRequest{Limit: &limit})
		}},
		{"Policies.ListAllPolicies", nil, func() (interface{}, error) {
			r, err := c.Policies.ListAllPolicies(ctx, org, &sgsdkgo.ListAllPoliciesRequest{Limit: &limit})
			if items, ok := r["msg"].([]interface{}); ok && len(items) > 0 {
				if item, ok := items[0].(map[string]interface{}); ok {
					if id, ok := item["ResourceId"].(string); ok {
						policy = strings.TrimPrefix(id, "/policies/")
					}
				}
			}
			return r, err
		}},
		{"Policies.ReadPolicy", &policy, func() (interface{}, error) { return c.Policies.ReadPolicy(ctx, org, policy) }},
		{"Secrets.ListAllSecrets", nil, func() (interface{}, error) { return c.Secrets.ListAllSecrets(ctx, org) }},
		{"RunnerGroups.ListAllRunnerGroups", nil, func() (interface{}, error) {
			return c.RunnerGroups.ListAllRunnerGroups(ctx, org, &runnergroups.ListAllRunnerGroupsRequest{})
		}},
		{"StateBackends.ListAllStateBackends", nil, func() (interface{}, error) {
			return c.StateBackends.ListAllStateBackends(ctx, org, &statebackends.ListAllStateBackendsRequest{})
		}},
		{"Connectors.ListAllConnectors", nil, func() (interface{}, error) {
			r, err := c.Connectors.ListAllConnectors(ctx, org, &sgsdkgo.ListAllConnectorsRequest{Limit: &limit})
			if err == nil && r != nil && len(r.Msg) > 0 {
				conn = strings.TrimPrefix(r.Msg[0].ResourceId, "/integrations/")
			}
			return r, err
		}},
		{"Connectors.ReadConnector", &conn, func() (interface{}, error) { return c.Connectors.ReadConnector(ctx, conn, org) }},
		{"ConnectorGroups.ListAllConnectorGroups", nil, func() (interface{}, error) {
			return c.ConnectorGroups.ListAllConnectorGroups(ctx, org, &sgsdkgo.ListAllConnectorGroupsRequest{})
		}},
		{"WorkflowGroups.ListAllWorkflowGroups", nil, func() (interface{}, error) {
			r, err := c.WorkflowGroups.ListAllWorkflowGroups(ctx, org, &sgsdkgo.ListAllWorkflowGroupsRequest{Limit: &limit})
			if err == nil && r != nil && len(r.Msg) > 0 {
				wfGrp = trimPrefix(r.Msg[0].ResourceName, "")
			}
			return r, err
		}},
		{"WorkflowGroups.ReadWorkflowGroup", &wfGrp, func() (interface{}, error) { return c.WorkflowGroups.ReadWorkflowGroup(ctx, org, wfGrp) }},
		{"WorkflowGroups.ListAllChildWorkflowGroups", &wfGrp, func() (interface{}, error) {
			return c.WorkflowGroups.ListAllChildWorkflowGroups(ctx, org, wfGrp, &sgsdkgo.ListAllChildWorkflowGroupsRequest{})
		}},
		{"WorkflowGroups.ListAllResourcesInWorkflowGroup", &wfGrp, func() (interface{}, error) {
			return c.WorkflowGroups.ListAllResourcesInWorkflowGroup(ctx, org, wfGrp, &workflowgroups.ListAllResourcesRequest{Limit: &limit})
		}},
		{"Workflows.ListAllWorkflows", &wfGrp, func() (interface{}, error) {
			r, err := c.Workflows.ListAllWorkflows(ctx, org, wfGrp, &sgsdkgo.ListAllWorkflowsRequest{Limit: &limit})
			if err == nil && r != nil && len(r.Msg) > 0 {
				wf = r.Msg[0].ResourceName
			}
			return r, err
		}},
		{"Workflows.ReadWorkflow", &wf, func() (interface{}, error) { return c.Workflows.ReadWorkflow(ctx, org, wf, wfGrp) }},
		{"Workflows.ListAllWorkflowArtifacts", &wf, func() (interface{}, error) { return c.Workflows.ListAllWorkflowArtifacts(ctx, org, wf, wfGrp) }},
		{"WorkflowRuns.ListAllWorkflowRuns", &wf, func() (interface{}, error) {
			r, err := c.WorkflowRuns.ListAllWorkflowRuns(ctx, org, wf, wfGrp, &sgsdkgo.ListAllWorkflowRunsRequest{Limit: &limit})
			if err == nil && r != nil && len(r.Msg) > 0 {
				wfRun = r.Msg[0].ResourceName
			}
			return r, err
		}},
		{"WorkflowRuns.ReadWorkflowRun", &wfRun, func() (interface{}, error) { return c.WorkflowRuns.ReadWorkflowRun(ctx, org, wf, wfGrp, wfRun) }},
		{"WorkflowRuns.ReadWorkflowRunLogs", &wfRun, func() (interface{}, error) { return c.WorkflowRuns.ReadWorkflowRunLogs(ctx, org, wf, wfGrp, wfRun) }},
		{"Stacks.ListAllStacks", &wfGrp, func() (interface{}, error) {
			return c.Stacks.ListAllStacks(ctx, org, wfGrp, &sgsdkgo.ListAllStacksRequest{})
		}},
		{"Templates.ListAllTemplatesBasedOnOwnerOrg", nil, func() (interface{}, error) {
			r, err := c.Templates.ListAllTemplatesBasedOnOwnerOrg(ctx, org, &sgsdkgo.ListAllTemplatesBasedOnOwnerOrgRequest{SgOrgid: org})
			if err == nil && r != nil {
				for _, tpl := range r.Msg {
					if tpl.TemplateType == nil || tpl.TemplateId == nil {
						continue
					}
					name := strings.TrimPrefix(*tpl.TemplateId, "/"+org+"/")
					switch *tpl.TemplateType {
					case "IAC":
						if iacTmpl == "" {
							iacTmpl = name
						}
					case "IAC_GROUP":
						if stackTmpl == "" {
							stackTmpl = name
						}
					}
				}
			}
			return r, err
		}},
		{"Templates.ListAllTemplates", nil, func() (interface{}, error) {
			return c.Templates.ListAllTemplates(ctx, sgsdkgo.ListAllTemplatesRequestTemplateTypeIac, &sgsdkgo.ListAllTemplatesRequest{SgOrgid: org, OwnerOrgs: s(org)})
		}},
		{"Templates.ReadTemplateRevision", &iacTmpl, func() (interface{}, error) {
			return c.Templates.ReadTemplateRevision(ctx, org, iacTmpl, sgsdkgo.ReadTemplateRevisionRequestTemplateTypeIac, &sgsdkgo.ReadTemplateRevisionRequest{SgOrgid: org})
		}},
		{"Templates.ListAllTemplateArtifacts", &iacTmpl, func() (interface{}, error) {
			return c.Templates.ListAllTemplateArtifacts(ctx, org, "IAC", iacTmpl+":1", &templates.ListAllTemplateArtifactsRequest{SgOrgid: org})
		}},
		{"Templates.ReadSubscription", nil, func() (interface{}, error) {
			return c.Templates.ReadSubscription(ctx, org, &sgsdkgo.ReadSubscriptionRequest{SubscriptionType: sgsdkgo.ReadSubscriptionRequestSubscriptionTypeIacSubscriptions})
		}},
		{"WorkflowTemplates.ReadWorkflowTemplate", &iacTmpl, func() (interface{}, error) { return c.WorkflowTemplates.ReadWorkflowTemplate(ctx, org, iacTmpl) }},
		{"StackTemplates.ReadStackTemplate", &stackTmpl, func() (interface{}, error) { return c.StackTemplates.ReadStackTemplate(ctx, org, stackTmpl) }},
		{"Chats.ListAllChats", nil, func() (interface{}, error) {
			r, err := c.Chats.ListAllChats(ctx, org, &chats.ListAllChatsRequest{Limit: &limit})
			if err == nil && r != nil && len(r.Msg) > 0 {
				chat = trimPrefix(r.Msg[0].ResourceId, "/chats/")
			}
			return r, err
		}},
		{"Chats.ReadChat", &chat, func() (interface{}, error) { return c.Chats.ReadChat(ctx, org, chat, &chats.ReadChatRequest{}) }},
		{"Chats.GetChatPrStatus", &chat, func() (interface{}, error) { return c.Chats.GetChatPrStatus(ctx, org, chat) }},
		{"Chats.GetChatPrSyncStatus", &chat, func() (interface{}, error) { return c.Chats.GetChatPrSyncStatus(ctx, org, chat) }},
		{"Chats.ListAllChatArtifacts", &chat, func() (interface{}, error) {
			return c.Chats.ListAllChatArtifacts(ctx, org, chat, &chats.ListAllChatArtifactsRequest{})
		}},
		{"Billing.ListBalance", nil, func() (interface{}, error) { return c.Billing.ListBalance(ctx, org) }},
		{"Billing.GetBillingDetails", nil, func() (interface{}, error) { return c.Billing.GetBillingDetails(ctx, org) }},
		{"Billing.ListInvoices", nil, func() (interface{}, error) {
			return c.Billing.ListInvoices(ctx, org, &billing.ListInvoicesRequest{Limit: &limit})
		}},
		{"Resources.ListResourceTags", nil, func() (interface{}, error) {
			return c.Resources.ListResourceTags(ctx, org, &resources.ListResourceTagsRequest{Limit: &limit})
		}},
		{"Resources.SearchResources", nil, func() (interface{}, error) {
			return c.Resources.SearchResources(ctx, org, &resources.SearchResourcesRequest{ResourceTypes: []string{"WORKFLOW"}, Limit: &limit})
		}},
		{"BenchmarkReports.GetBenchmarkReports", nil, func() (interface{}, error) {
			return c.BenchmarkReports.GetBenchmarkReports(ctx, org, &sgsdkgo.GetBenchmarkReportsRequest{Limit: &limit, GroupBy: s("CSP")})
		}},
	}

	for _, cl := range calls {
		cl := cl
		t.Run(cl.name, func(t *testing.T) {
			if cl.need != nil && *cl.need == "" {
				t.Skip("no resource of this kind in the organization")
			}
			cap.body, cap.status = nil, 0
			res, err := cl.fn()
			if err != nil {
				t.Fatalf("status=%d: %v", cap.status, err)
			}
			var raw, typed interface{}
			if err := json.Unmarshal(cap.body, &raw); err != nil {
				return // empty (204) or non-JSON body: nothing to compare
			}
			tb, err := json.Marshal(res)
			if err != nil {
				t.Fatalf("re-encode typed response: %v", err)
			}
			_ = json.Unmarshal(tb, &typed)
			var miss []string
			undeclared("", raw, typed, &miss)
			if len(miss) > 0 {
				t.Errorf("live response carries fields the SDK type does not declare: %s", strings.Join(miss, ", "))
			}
		})
	}
}
