package cli

import (
	"sort"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// expectedLeaves is the complete command surface: one leaf per SDK method.
// Update it deliberately when the SDK gains or loses an operation.
const expectedLeaves = `
api-access create
api-access delete
api-access get
api-access list
api-access regenerate-key
api-access update
api-tokens create
api-tokens delete
audit-logs list
benchmark-reports get
billing balance
billing change-plan
billing dashboard-url
billing details
billing invoices
billing payment-methods detach
billing payment-methods list
billing payment-methods set-default
billing payment-methods setup-intent
billing profile setup
billing profile update
chats artifacts get-url
chats artifacts list
chats artifacts upload-url
chats commit-bundle
chats create
chats create-pr
chats delete
chats get
chats list
chats messages cancel
chats messages create
chats messages get
chats messages list
chats messages retry
chats pr-status
chats pr-sync-status
chats update
connector-groups authenticate
connector-groups connectors authenticate
connector-groups connectors delete
connector-groups connectors get
connector-groups connectors list
connector-groups connectors update
connector-groups create
connector-groups delete
connector-groups discovery-scan
connector-groups get
connector-groups list
connector-groups update
connectors authenticate
connectors create
connectors delete
connectors get
connectors github-repos
connectors list
connectors list-accounts
connectors repos
connectors update
organizations bulk-action
organizations count-workflows
organizations create
organizations delete
organizations get
organizations list
organizations update
organizations workflows
policies create
policies delete
policies get
policies list
policies update
resources move
resources search
resources tags
role-bindings get
roles create
roles delete
roles get
roles list
roles update
runner-groups create
runner-groups delete
runner-groups deregister-runner
runner-groups get
runner-groups list
runner-groups register
runner-groups storage-backend-auth
runner-groups update
runner-groups update-runner-state
secrets create
secrets delete
secrets list
secrets read-bulk
secrets update
stack-runs create
stack-runs get
stack-runs list
stack-template-revisions create
stack-template-revisions delete
stack-template-revisions get
stack-template-revisions update
stack-templates create
stack-templates delete
stack-templates get
stack-templates update
stack-workflow-run-facts create
stack-workflow-run-facts get
stack-workflow-run-facts update
stack-workflow-runs approve
stack-workflow-runs create
stack-workflow-runs delete
stack-workflow-runs get
stack-workflow-runs list
stack-workflow-runs logs
stack-workflow-runs update
stack-workflows artifacts create
stack-workflows artifacts delete
stack-workflows artifacts get
stack-workflows artifacts get-url
stack-workflows artifacts list
stack-workflows artifacts lock
stack-workflows artifacts unlock
stack-workflows compare
stack-workflows create
stack-workflows delete
stack-workflows file-upload-url
stack-workflows get
stack-workflows list
stack-workflows outputs
stack-workflows tfstate-upload-url
stack-workflows update
stack-workflows vcs-triggers
stacks compare
stacks create
stacks delete
stacks get
stacks list
stacks outputs
stacks update
state-backends create
state-backends delete
state-backends get
state-backends list
state-backends list-statefiles
state-backends update
templates artifacts delete
templates artifacts download-url
templates artifacts list
templates artifacts upload-url
templates create
templates delete
templates get
templates get-group-iac
templates input-schema
templates list
templates list-by-owner
templates public get
templates public list
templates public list-all
templates subscriptions create
templates subscriptions get
templates subscriptions update
templates update
templates vcs-triggers
users create
users delete
users get
users list
users update
workflow-groups create
workflow-groups create-child
workflow-groups delete
workflow-groups get
workflow-groups list
workflow-groups list-children
workflow-groups list-resources
workflow-groups update
workflow-run-facts create
workflow-run-facts get
workflow-run-facts update
workflow-runs approve
workflow-runs cancel
workflow-runs create
workflow-runs delete
workflow-runs get
workflow-runs get-by-ksuid
workflow-runs list
workflow-runs logs
workflow-runs update
workflow-runs update-by-ksuid
workflow-step-template-revisions create
workflow-step-template-revisions delete
workflow-step-template-revisions get
workflow-step-template-revisions update
workflow-step-templates create
workflow-step-templates delete
workflow-step-templates get
workflow-step-templates update
workflow-template-revisions create
workflow-template-revisions delete
workflow-template-revisions get
workflow-template-revisions update
workflow-templates create
workflow-templates delete
workflow-templates get
workflow-templates update
workflows artifacts create
workflows artifacts delete
workflows artifacts get
workflows artifacts get-url
workflows artifacts list
workflows artifacts lock
workflows artifacts rollback
workflows artifacts unlock
workflows artifacts versions
workflows compare
workflows create
workflows delete
workflows file-upload-url
workflows get
workflows list
workflows outputs
workflows tfstate-upload-url
workflows update
workflows vcs-triggers
`

func leafCommands(root *cobra.Command) []string {
	var leaves []string
	var walk func(c *cobra.Command, prefix string)
	walk = func(c *cobra.Command, prefix string) {
		if c.Name() == "help" || c.Name() == "completion" {
			return
		}
		if !c.HasSubCommands() {
			leaves = append(leaves, prefix+c.Name())
			return
		}
		for _, sub := range c.Commands() {
			walk(sub, prefix+c.Name()+" ")
		}
	}
	for _, sub := range root.Commands() {
		walk(sub, "")
	}
	sort.Strings(leaves)
	return leaves
}

func TestCommandTree(t *testing.T) {
	want := strings.Split(strings.TrimSpace(expectedLeaves), "\n")
	assert.Equal(t, want, leafCommands(NewRootCmd()))
	assert.Len(t, want, 226, "one CLI leaf command per SDK method")
}

func TestEveryLeafHasRunAndShort(t *testing.T) {
	var walk func(c *cobra.Command)
	walk = func(c *cobra.Command) {
		if c.Name() == "help" || c.Name() == "completion" {
			return
		}
		assert.NotEmpty(t, c.Short, "command %q needs a Short description", c.CommandPath())
		if !c.HasSubCommands() {
			assert.NotNil(t, c.RunE, "leaf %q needs RunE", c.CommandPath())
		}
		for _, sub := range c.Commands() {
			walk(sub)
		}
	}
	walk(NewRootCmd())
}
