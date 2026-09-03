// Package cli implements the sg command-line interface on top of the SDK.
package cli

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/option"
)

// Version is reported by `sg --version`. Override it at build time with
// -ldflags "-X github.com/StackGuardian/sg-sdk-go/cli.Version=v1.2.3".
var Version = "dev"

const (
	envAPIToken    = "SG_API_TOKEN"
	envBaseURL     = "SG_BASE_URL"
	envOrg         = "SG_ORG"
	defaultBaseURL = "https://api.app.stackguardian.io"
)

// app carries the global flag values and the response capture shared by every command.
type app struct {
	apiKey      string
	baseURL     string
	org         string
	maxAttempts uint
	timeout     time.Duration
	compact     bool
	verbose     bool

	capture *capture
}

// action performs one SDK call. The response body is printed by app.run afterwards.
type action func(ctx context.Context, c *client.Client) error

// NewRootCmd builds the full command tree. Every call returns an independent
// tree so tests stay isolated from each other.
func NewRootCmd() *cobra.Command {
	a := &app{}
	root := &cobra.Command{
		Use:   "sg",
		Short: "StackGuardian command-line interface",
		Long: `sg is a command-line interface for the StackGuardian API.

Authentication and scoping come from flags or environment variables:
  --api-key   SG_API_TOKEN   API token (the "apikey " prefix is added when missing)
  --base-url  SG_BASE_URL    API base URL (default ` + defaultBaseURL + `)
  --org       SG_ORG         Organization name

Create and update commands take the request body as JSON via --body or
--body-file (use "-" to read from stdin). Responses are printed as JSON exactly
as returned by the API.`,
		Version:       Version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	pf := root.PersistentFlags()
	pf.StringVar(&a.apiKey, "api-key", "", "API token (env "+envAPIToken+")")
	pf.StringVar(&a.baseURL, "base-url", "", "API base URL (env "+envBaseURL+", default "+defaultBaseURL+")")
	pf.StringVar(&a.org, "org", "", "organization name (env "+envOrg+")")
	pf.UintVar(&a.maxAttempts, "max-attempts", 0, "maximum HTTP attempts per request (0 uses the SDK default)")
	pf.DurationVar(&a.timeout, "timeout", 0, "request timeout such as 30s (0 disables the timeout)")
	pf.BoolVar(&a.compact, "compact", false, "print compact single-line JSON")
	pf.BoolVar(&a.verbose, "verbose", false, "log HTTP requests and response status to stderr")

	root.AddCommand(
		a.organizationsCmd(),
		a.apiAccessCmd(),
		a.auditLogsCmd(),
		a.usersCmd(),
		a.rolesCmd(),
		a.benchmarkReportsCmd(),
		a.connectorGroupsCmd(),
		a.connectorsCmd(),
		a.policiesCmd(),
		a.runnerGroupsCmd(),
		a.secretsCmd(),
		a.templatesCmd(),
		a.workflowGroupsCmd(),
		a.stacksCmd(),
		a.stackRunsCmd(),
		a.stackWorkflowsCmd(),
		a.stackWorkflowRunsCmd(),
		a.stackWorkflowRunFactsCmd(),
		a.workflowsCmd(),
		a.workflowRunsCmd(),
		a.workflowRunFactsCmd(),
		a.workflowStepTemplatesCmd(),
		a.workflowStepTemplateRevisionsCmd(),
		a.workflowTemplatesCmd(),
		a.workflowTemplateRevisionsCmd(),
		a.stackTemplatesCmd(),
		a.stackTemplateRevisionsCmd(),
		a.roleBindingsCmd(),
		a.apiTokensCmd(),
		a.stateBackendsCmd(),
		a.resourcesCmd(),
		a.billingCmd(),
		a.chatsCmd(),
	)
	return root
}

// Execute runs the CLI with os.Args and returns the process exit code.
func Execute() int {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return 1
	}
	return 0
}

// run resolves configuration, builds a client, executes fn and prints the
// captured response body.
func (a *app) run(cmd *cobra.Command, fn action) error {
	if err := a.resolve(cmd.Annotations[annotationOrg] != orgOptional); err != nil {
		return err
	}
	c := a.newClient(cmd)
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	if a.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, a.timeout)
		defer cancel()
	}
	if err := fn(ctx, c); err != nil {
		return err
	}
	return a.printResponse(cmd)
}

const (
	// annotationOrg marks commands that work without an organization (for example
	// creating or listing organizations).
	annotationOrg = "sg.org"
	orgOptional   = "optional"
)

// resolve fills unset global values from the environment and validates them.
func (a *app) resolve(needOrg bool) error {
	if a.apiKey == "" {
		a.apiKey = os.Getenv(envAPIToken)
	}
	if a.apiKey == "" {
		return errors.New("an API token is required: set --api-key or " + envAPIToken)
	}
	if !strings.HasPrefix(strings.ToLower(a.apiKey), "apikey ") {
		a.apiKey = "apikey " + a.apiKey
	}
	if a.baseURL == "" {
		a.baseURL = os.Getenv(envBaseURL)
	}
	if a.baseURL == "" {
		a.baseURL = defaultBaseURL
	}
	a.baseURL = strings.TrimRight(a.baseURL, "/")
	if a.org == "" {
		a.org = os.Getenv(envOrg)
	}
	if a.org == "" && needOrg {
		return errors.New("an organization is required: set --org or " + envOrg)
	}
	return nil
}

func (a *app) newClient(cmd *cobra.Command) *client.Client {
	a.capture = &capture{next: http.DefaultTransport, verbose: a.verbose, log: cmd.ErrOrStderr()}
	opts := []option.RequestOption{
		option.WithApiKey(a.apiKey),
		option.WithBaseURL(a.baseURL),
		option.WithHTTPClient(&http.Client{Transport: a.capture}),
	}
	if a.maxAttempts > 0 {
		opts = append(opts, option.WithMaxAttempts(a.maxAttempts))
	}
	return client.NewClient(opts...)
}

// handler is the body of a leaf command: it performs the SDK call for the
// parsed flags and positional args.
type handler func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error

// command builds a leaf command that runs h through app.run. Common verbs get
// conventional aliases (get/read, list/ls, delete/rm).
func (a *app) command(use, short string, args cobra.PositionalArgs, h handler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  args,
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.run(cmd, func(ctx context.Context, c *client.Client) error {
				return h(ctx, c, cmd, args)
			})
		},
	}
	switch strings.Fields(use)[0] {
	case "get":
		cmd.Aliases = []string{"read"}
	case "list":
		cmd.Aliases = []string{"ls"}
	case "delete":
		cmd.Aliases = []string{"rm"}
	}
	return cmd
}
