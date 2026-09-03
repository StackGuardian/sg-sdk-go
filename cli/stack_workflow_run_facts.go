package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) stackWorkflowRunFactsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack-workflow-run-facts",
		Short: "Read the facts of a stack workflow run",
	}

	get := a.command("get RUN_FACTS", "Get the facts of a stack workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflowRunFacts.GetStackWorkflowRunFacts(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), getString(cmd, "wf-run"), args[0])
			return err
		})
	addWfGrpFlag(get)
	addStackFlag(get)
	addWfFlag(get)
	requiredString(get, "wf-run", "workflow run id")

	scope := func(cmd *cobra.Command) {
		addWfGrpFlag(cmd)
		addStackFlag(cmd)
		addWfFlag(cmd)
		requiredString(cmd, "wf-run", "workflow run id")
	}

	create := a.command("create RUN_FACTS", "Create or replace the facts of a stack workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req map[string]interface{}
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflowRunFacts.CreateStackWorkflowRunFacts(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), getString(cmd, "wf-run"), args[0], req)
			return err
		})
	scope(create)
	addBodyFlags(create)

	update := a.command("update RUN_FACTS", "Update the facts of a stack workflow run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req map[string]interface{}
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflowRunFacts.UpdateStackWorkflowRunFacts(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), getString(cmd, "wf-run"), args[0], req)
			return err
		})
	scope(update)
	addBodyFlags(update)

	cmd.AddCommand(get, create, update)
	return cmd
}
