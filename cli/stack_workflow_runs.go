package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/stackworkflowruns"
)

func (a *app) stackWorkflowRunsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack-workflow-runs",
		Short: "Trigger, inspect and approve runs of a stack workflow",
	}
	scope := func(cmd *cobra.Command) {
		addWfGrpFlag(cmd)
		addStackFlag(cmd)
		addWfFlag(cmd)
	}

	create := a.command("create", "Trigger a run of a stack workflow (body optional)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.WorkflowRun
			if err := readOptionalBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflowRuns.CreateStackWorkflowRun(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(create)
	addBodyFlags(create)

	get := a.command("get WORKFLOW_RUN", "Get a stack workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflowRuns.ReadStackWorkflowRun(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(get)

	logs := a.command("logs WORKFLOW_RUN", "Get the logs of a stack workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflowRuns.ReadStackWorkflowRunLogs(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(logs)

	approve := a.command("approve WORKFLOW_RUN", "Approve (or, with --reject, reject) a stack workflow run waiting for approval", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			return c.StackWorkflowRuns.ApproveStackWorkflowRun(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], approvalFromFlags(cmd))
		})
	scope(approve)
	addApprovalFlags(approve)

	update := a.command("update WORKFLOW_RUN", "Update a stack workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedWorkflowRun
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflowRuns.UpdateStackWorkflowRun(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], &req)
			return err
		})
	scope(update)
	addBodyFlags(update)

	del := a.command("delete WORKFLOW_RUN", "Cancel a stack workflow run through the DELETE endpoint", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflowRuns.DeleteStackWorkflowRun(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(del)

	list := a.command("list", "List the runs of a stack workflow", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := stackworkflowruns.ListAllStackWorkflowRunsRequest{
				Limit:                  flagInt(cmd, "limit"),
				Lastevaluatedkey:       flagString(cmd, "last-evaluated-key"),
				CreatedAfter:           flagString(cmd, "created-after"),
				ParentKSUID:            flagString(cmd, "parent-ksuid"),
				ResourceNames:          flagString(cmd, "resource-names"),
				TemplateId:             flagString(cmd, "template-id"),
				UseMarketplaceTemplate: flagString(cmd, "use-marketplace-template"),
				Author:                 flagString(cmd, "author"),
				Approver:               flagString(cmd, "approver"),
				LatestStatus:           flagString(cmd, "latest-status"),
				RepoId:                 flagString(cmd, "repo-id"),
				Ref:                    flagString(cmd, "ref"),
				TerraformAction:        flagString(cmd, "terraform-action"),
				RunnerType:             flagString(cmd, "runner-type"),
				ContextTags:            flagString(cmd, "context-tags"),
			}
			_, err := c.StackWorkflowRuns.ListAllStackWorkflowRuns(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(list)
	addPaginationFlags(list)
	list.Flags().String("created-after", "", "only runs created after this Unix timestamp in milliseconds")
	list.Flags().String("parent-ksuid", "", "KSUID of the parent workflow")
	list.Flags().String("resource-names", "", "filter by resource names (comma-separated)")
	list.Flags().String("template-id", "", "filter by IaC template id")
	list.Flags().String("use-marketplace-template", "", "filter by marketplace template usage (true/false)")
	list.Flags().String("author", "", "filter by authors (comma-separated)")
	list.Flags().String("approver", "", "filter by approvers (comma-separated)")
	list.Flags().String("latest-status", "", "filter by latest status (comma-separated)")
	list.Flags().String("repo-id", "", "filter by VCS repository")
	list.Flags().String("ref", "", "filter by VCS ref")
	list.Flags().String("terraform-action", "", "filter by Terraform action (comma-separated)")
	list.Flags().String("runner-type", "", "filter by runner type")
	list.Flags().String("context-tags", "", "filter by context tags as key or key:value (comma-separated)")

	cmd.AddCommand(create, get, logs, approve, update, del, list)
	return cmd
}
