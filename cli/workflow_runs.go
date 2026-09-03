package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) workflowRunsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workflow-runs",
		Aliases: []string{"wfruns", "wfrun"},
		Short:   "Trigger, inspect, cancel and approve workflow runs",
	}
	scope := func(cmd *cobra.Command) {
		addWfGrpFlag(cmd)
		addWfFlag(cmd)
	}

	create := a.command("create", "Trigger a workflow run (body optional)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.WorkflowRun
			if err := readOptionalBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowRuns.CreateWorkflowRun(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(create)
	addBodyFlags(create)

	get := a.command("get WORKFLOW_RUN", "Get a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRuns.ReadWorkflowRun(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(get)

	update := a.command("update WORKFLOW_RUN", "Update a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedWorkflowRun
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowRuns.UpdateWorkflowRun(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], &req)
			return err
		})
	scope(update)
	addBodyFlags(update)

	cancel := a.command("cancel WORKFLOW_RUN", "Cancel a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRuns.CancelWorkflowRun(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(cancel)

	logs := a.command("logs WORKFLOW_RUN", "Get the logs of a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRuns.ReadWorkflowRunLogs(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(logs)

	approve := a.command("approve WORKFLOW_RUN", "Approve (or, with --reject, reject) a workflow run waiting for approval", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRuns.ApproveWorkflowRun(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], approvalFromFlags(cmd))
			return err
		})
	scope(approve)
	addApprovalFlags(approve)

	list := a.command("list", "List the runs of a workflow", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllWorkflowRunsRequest{
				ContextTags:      flagString(cmd, "context-tags"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
			}
			_, err := c.WorkflowRuns.ListAllWorkflowRuns(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(list)
	addPaginationFlags(list)
	list.Flags().String("context-tags", "", "filter by context tags")

	del := a.command("delete WORKFLOW_RUN", "Cancel a workflow run through the DELETE endpoint", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRuns.DeleteWorkflowRun(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(del)

	getByKsuid := a.command("get-by-ksuid RESOURCE_KSUID", "Get a workflow run by KSUIDs (runner and machine tokens only)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRuns.ReadWorkflowRunByKsuid(ctx, a.org, getString(cmd, "parent-ksuid"), args[0])
			return err
		})
	requiredString(getByKsuid, "parent-ksuid", "KSUID of the parent workflow")

	updateByKsuid := a.command("update-by-ksuid RESOURCE_KSUID", "Update a workflow run by KSUIDs (runner and machine tokens only)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedWorkflowRun
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowRuns.UpdateWorkflowRunByKsuid(ctx, a.org, getString(cmd, "parent-ksuid"), args[0], &req)
			return err
		})
	requiredString(updateByKsuid, "parent-ksuid", "KSUID of the parent workflow")
	addBodyFlags(updateByKsuid)

	cmd.AddCommand(create, get, update, cancel, logs, approve, list, del, getByKsuid, updateByKsuid)
	return cmd
}
