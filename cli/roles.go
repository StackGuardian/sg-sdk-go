package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) rolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roles",
		Short: "Manage roles",
	}

	create := a.command("create", "Create a role", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.Role
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.AccessManagement.CreateRole(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get ROLE", "Get a role", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.AccessManagement.ReadRole(ctx, a.org, args[0])
			return err
		})

	del := a.command("delete ROLE", "Delete a role", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			return c.AccessManagement.DeleteRole(ctx, a.org, args[0])
		})

	update := a.command("update ROLE", "Update a role", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedRole
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			if req.ResourceName == nil {
				req.ResourceName = sgsdkgo.Optional(args[0]) // the API requires it on updates
			}
			_, err := c.AccessManagement.UpdateRole(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	list := a.command("list", "List roles", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllRolesRequest{
				IncludePredefinedRoles: flagBool(cmd, "include-predefined-roles"),
			}
			_, err := c.AccessManagement.ListAllRoles(ctx, a.org, &req)
			return err
		})
	list.Flags().Bool("include-predefined-roles", false, "include predefined system roles")

	cmd.AddCommand(create, get, del, update, list)
	return cmd
}
