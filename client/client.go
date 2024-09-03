// This file was auto-generated by Fern from our API Definition.

package client

import (
	connectors "github.com/StackGuardian/sg-sdk-go/connectors"
	core "github.com/StackGuardian/sg-sdk-go/core"
	option "github.com/StackGuardian/sg-sdk-go/option"
	organizations "github.com/StackGuardian/sg-sdk-go/organizations"
	policies "github.com/StackGuardian/sg-sdk-go/policies"
	runnergroups "github.com/StackGuardian/sg-sdk-go/runnergroups"
	stacks "github.com/StackGuardian/sg-sdk-go/stacks"
	stackworkflowrunfacts "github.com/StackGuardian/sg-sdk-go/stackworkflowrunfacts"
	stackworkflowruns "github.com/StackGuardian/sg-sdk-go/stackworkflowruns"
	stackworkflows "github.com/StackGuardian/sg-sdk-go/stackworkflows"
	templates "github.com/StackGuardian/sg-sdk-go/templates"
	usersroles "github.com/StackGuardian/sg-sdk-go/usersroles"
	workflowgroups "github.com/StackGuardian/sg-sdk-go/workflowgroups"
	workflowrunfacts "github.com/StackGuardian/sg-sdk-go/workflowrunfacts"
	workflowruns "github.com/StackGuardian/sg-sdk-go/workflowruns"
	workflows "github.com/StackGuardian/sg-sdk-go/workflows"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Organizations         *organizations.Client
	Connectors            *connectors.Client
	UsersRoles            *usersroles.Client
	Policies              *policies.Client
	RunnerGroups          *runnergroups.Client
	Templates             *templates.Client
	WorkflowGroups        *workflowgroups.Client
	Stacks                *stacks.Client
	StackWorkflowRuns     *stackworkflowruns.Client
	StackWorkflows        *stackworkflows.Client
	StackWorkflowRunFacts *stackworkflowrunfacts.Client
	Workflows             *workflows.Client
	WorkflowRuns          *workflowruns.Client
	WorkflowRunFacts      *workflowrunfacts.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:                options.ToHeader(),
		Organizations:         organizations.NewClient(opts...),
		Connectors:            connectors.NewClient(opts...),
		UsersRoles:            usersroles.NewClient(opts...),
		Policies:              policies.NewClient(opts...),
		RunnerGroups:          runnergroups.NewClient(opts...),
		Templates:             templates.NewClient(opts...),
		WorkflowGroups:        workflowgroups.NewClient(opts...),
		Stacks:                stacks.NewClient(opts...),
		StackWorkflowRuns:     stackworkflowruns.NewClient(opts...),
		StackWorkflows:        stackworkflows.NewClient(opts...),
		StackWorkflowRunFacts: stackworkflowrunfacts.NewClient(opts...),
		Workflows:             workflows.NewClient(opts...),
		WorkflowRuns:          workflowruns.NewClient(opts...),
		WorkflowRunFacts:      workflowrunfacts.NewClient(opts...),
	}
}
