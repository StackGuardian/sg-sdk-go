package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) workflowRunFactsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workflow-run-facts",
		Short: "Read the facts of a workflow run",
	}

	get := a.command("get RUN_FACTS", "Get the facts of a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.WorkflowRunFacts.ReadWorkflowRunFacts(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), getString(cmd, "wf-run"), args[0])
			return err
		})
	addWfGrpFlag(get)
	addWfFlag(get)
	requiredString(get, "wf-run", "workflow run id")

	create := a.command("create RUN_FACTS", "Create or replace the facts of a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req map[string]interface{}
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowRunFacts.CreateWorkflowRunFacts(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), getString(cmd, "wf-run"), args[0], req)
			return err
		})
	addWfGrpFlag(create)
	addWfFlag(create)
	requiredString(create, "wf-run", "workflow run id")
	addBodyFlags(create)

	update := a.command("update RUN_FACTS", "Update the facts of a workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req map[string]interface{}
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowRunFacts.UpdateWorkflowRunFacts(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), getString(cmd, "wf-run"), args[0], req)
			return err
		})
	addWfGrpFlag(update)
	addWfFlag(update)
	requiredString(update, "wf-run", "workflow run id")
	addBodyFlags(update)

	cmd.AddCommand(get, create, update)
	return cmd
}
