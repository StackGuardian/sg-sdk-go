package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/stacks"
)

func (a *app) stacksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stacks",
		Short: "Manage stacks inside a workflow group",
	}

	create := a.command("create", "Create a stack", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.Stack
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			req.RunOnCreate = flagBool(cmd, "run-on-create")
			_, err := c.Stacks.CreateStack(ctx, a.org, getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(create)
	addBodyFlags(create)
	create.Flags().Bool("run-on-create", false, "trigger the stack run immediately after creation")

	get := a.command("get STACK", "Get a stack", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Stacks.ReadStack(ctx, a.org, args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(get)

	del := a.command("delete STACK", "Delete a stack", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Stacks.DeleteStack(ctx, a.org, args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(del)

	update := a.command("update STACK", "Update a stack", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedStack
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Stacks.UpdateStack(ctx, a.org, args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(update)
	addBodyFlags(update)

	outputs := a.command("outputs STACK", "Get the outputs of a stack", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Stacks.ReadStackOutputs(ctx, a.org, args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(outputs)

	list := a.command("list", "List the stacks in a workflow group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllStacksRequest{
				Description:      flagString(cmd, "description"),
				ResourceNames:    flagString(cmd, "resource-names"),
				Tags:             flagString(cmd, "tags"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
			}
			_, err := c.Stacks.ListAllStacks(ctx, a.org, getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(list)
	addPaginationFlags(list)
	list.Flags().String("description", "", "filter by description")
	list.Flags().String("resource-names", "", "filter by resource names")
	list.Flags().String("tags", "", "filter by tags")

	compare := a.command("compare STACK", "Compare a stack against a stack template without applying changes", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req stacks.CompareStackRequest
			if hasBody(cmd) {
				if err := readBody(cmd, &req); err != nil {
					return err
				}
			} else {
				req.TargetTemplateGroupId = getString(cmd, "target-template-group-id")
				req.UpgradeMode = flagString(cmd, "upgrade-mode")
			}
			_, err := c.Stacks.CompareStack(ctx, a.org, args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(compare)
	addBodyFlags(compare)
	compare.Flags().String("target-template-group-id", "", "stack template to compare against")
	compare.Flags().String("upgrade-mode", "", "upgrade mode to simulate")
	compare.MarkFlagsOneRequired("target-template-group-id", "body", "body-file")

	cmd.AddCommand(create, get, del, update, outputs, list, compare)
	return cmd
}
