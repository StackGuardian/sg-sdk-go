package cli

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) stackRunsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack-runs",
		Short: "Trigger and inspect stack runs",
	}

	create := a.command("create", "Trigger a stack run", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.StackAction
			switch {
			case hasBody(cmd):
				if err := readBody(cmd, &req); err != nil {
					return err
				}
			case cmd.Flags().Changed("action-type"):
				req.ActionType = getString(cmd, "action-type")
			default:
				return errors.New("an action is required: use --action-type or --body")
			}
			_, err := c.StackRuns.CreateStackRun(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(create)
	addStackFlag(create)
	addBodyFlags(create)
	create.Flags().String("action-type", "", "action to run, e.g. APPLY (shortcut for a body of {\"ActionType\": ...})")
	create.MarkFlagsMutuallyExclusive("action-type", "body")
	create.MarkFlagsMutuallyExclusive("action-type", "body-file")

	get := a.command("get STACK_RUN", "Get a stack run", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackRuns.ReadStackRun(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(get)
	addStackFlag(get)

	list := a.command("list", "List the runs of a stack", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllStackRunsRequest{
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
			}
			_, err := c.StackRuns.ListAllStackRuns(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(list)
	addStackFlag(list)
	addPaginationFlags(list)

	cmd.AddCommand(create, get, list)
	return cmd
}
