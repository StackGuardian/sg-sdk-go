package stackguardian

import (
	api "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/internal"
	"github.com/StackGuardian/sg-sdk-go/services/accessmanagement"
	"github.com/StackGuardian/sg-sdk-go/services/benchmarkreports"
	"github.com/StackGuardian/sg-sdk-go/services/connectorgroups"
	"github.com/StackGuardian/sg-sdk-go/services/connectors"
	"github.com/StackGuardian/sg-sdk-go/services/organizations"
	"github.com/StackGuardian/sg-sdk-go/services/policies"
	"github.com/StackGuardian/sg-sdk-go/services/runnergroups"
	"github.com/StackGuardian/sg-sdk-go/services/secrets"
	"github.com/StackGuardian/sg-sdk-go/services/stackruns"
	"github.com/StackGuardian/sg-sdk-go/services/stacks"
	"github.com/StackGuardian/sg-sdk-go/services/stackworkflowrunfacts"
	"github.com/StackGuardian/sg-sdk-go/services/stackworkflowruns"
	"github.com/StackGuardian/sg-sdk-go/services/stackworkflows"
	"github.com/StackGuardian/sg-sdk-go/services/templates"
	"github.com/StackGuardian/sg-sdk-go/services/workflowgroups"
	"github.com/StackGuardian/sg-sdk-go/services/workflowrunfacts"
	"github.com/StackGuardian/sg-sdk-go/services/workflowruns"
	"github.com/StackGuardian/sg-sdk-go/services/workflows"
)

// Re-export error helper functions for convenience
var (
	IsNotFoundError      = api.IsNotFoundError
	IsUnauthorizedError  = api.IsUnauthorizedError
)

// Client is the main StackGuardian SDK client.
type Client struct {
	config     *Config
	httpClient *internal.HTTPClient

	// Service clients
	Organizations         *organizations.Service
	AccessManagement      *accessmanagement.Service
	ConnectorGroups       *connectorgroups.Service
	Connectors            *connectors.Service
	Policies              *policies.Service
	BenchmarkReports      *benchmarkreports.Service
	RunnerGroups          *runnergroups.Service
	Secrets               *secrets.Service
	Templates             *templates.Service
	WorkflowGroups        *workflowgroups.Service
	Stacks                *stacks.Service
	StackRuns             *stackruns.Service
	StackWorkflows        *stackworkflows.Service
	StackWorkflowRuns     *stackworkflowruns.Service
	StackWorkflowRunFacts *stackworkflowrunfacts.Service
	Workflows             *workflows.Service
	WorkflowRuns          *workflowruns.Service
	WorkflowRunFacts      *workflowrunfacts.Service
}

// NewClient creates a new StackGuardian client with the given configuration.
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		config = DefaultConfig()
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Create shared HTTP client
	httpClient := internal.NewHTTPClient(
		config.HTTPClient,
		config.BaseURL,
		config.APIKey,
		config.UserAgent,
		config.MaxRetries,
		config.RetryWaitMin,
		config.RetryWaitMax,
	)

	client := &Client{
		config:     config,
		httpClient: httpClient,
	}

	// Initialize service clients
	client.Organizations = organizations.NewServiceWithHTTPClient(httpClient)
	client.AccessManagement = accessmanagement.NewServiceWithHTTPClient(httpClient)
	client.ConnectorGroups = connectorgroups.NewServiceWithHTTPClient(httpClient)
	client.Connectors = connectors.NewServiceWithHTTPClient(httpClient)
	client.Policies = policies.NewServiceWithHTTPClient(httpClient)
	client.BenchmarkReports = benchmarkreports.NewServiceWithHTTPClient(httpClient)
	client.RunnerGroups = runnergroups.NewServiceWithHTTPClient(httpClient)
	client.Secrets = secrets.NewServiceWithHTTPClient(httpClient)
	client.Templates = templates.NewServiceWithHTTPClient(httpClient)
	client.WorkflowGroups = workflowgroups.NewServiceWithHTTPClient(httpClient)
	client.Stacks = stacks.NewServiceWithHTTPClient(httpClient)
	client.StackRuns = stackruns.NewServiceWithHTTPClient(httpClient)
	client.StackWorkflows = stackworkflows.NewServiceWithHTTPClient(httpClient)
	client.StackWorkflowRuns = stackworkflowruns.NewServiceWithHTTPClient(httpClient)
	client.StackWorkflowRunFacts = stackworkflowrunfacts.NewServiceWithHTTPClient(httpClient)
	client.Workflows = workflows.NewServiceWithHTTPClient(httpClient)
	client.WorkflowRuns = workflowruns.NewServiceWithHTTPClient(httpClient)
	client.WorkflowRunFacts = workflowrunfacts.NewServiceWithHTTPClient(httpClient)

	return client, nil
}

// Config returns the client configuration.
func (c *Client) Config() *Config {
	return c.config
}
