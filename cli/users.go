package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) usersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Manage users and SSO groups in the organization",
	}

	userSelector := func(cmd *cobra.Command) *sgsdkgo.GetorRemoveUserFromOrganization {
		return &sgsdkgo.GetorRemoveUserFromOrganization{
			UserId: flagString(cmd, "user-id"),
			Alias:  flagString(cmd, "alias"),
		}
	}
	addSelectorFlags := func(cmd *cobra.Command) {
		cmd.Flags().String("user-id", "", "user id (email) or SSO group id")
		cmd.Flags().String("alias", "", "alias of an SSO group")
		cmd.MarkFlagsOneRequired("user-id", "alias")
	}

	get := a.command("get", "Get a user or SSO group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.AccessManagement.ReadUser(ctx, a.org, userSelector(cmd))
			return err
		})
	addSelectorFlags(get)

	create := a.command("create", "Invite a user or SSO group to the organization", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.AddUserToOrganization
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.AccessManagement.CreateUser(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	del := a.command("delete", "Remove a user or SSO group from the organization", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.AccessManagement.DeleteUser(ctx, a.org, userSelector(cmd))
			return err
		})
	addSelectorFlags(del)

	update := a.command("update", "Update a user or SSO group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.AddUserToOrganization
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.AccessManagement.UpdateUser(ctx, a.org, &req)
			return err
		})
	addBodyFlags(update)

	list := a.command("list", "List users and SSO groups", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllUsersRequest{
				Alias:       flagString(cmd, "alias"),
				EntityType:  flagString(cmd, "entity-type"),
				LoginMethod: flagString(cmd, "login-method"),
				Roles:       flagString(cmd, "roles"),
				UserId:      flagString(cmd, "user-id"),
			}
			_, err := c.AccessManagement.ListAllUsers(ctx, a.org, &req)
			return err
		})
	list.Flags().String("alias", "", "filter by alias")
	list.Flags().String("entity-type", "", "filter by entity type (EMAIL, GROUP)")
	list.Flags().String("login-method", "", "filter by login method")
	list.Flags().String("roles", "", "filter by roles (comma-separated)")
	list.Flags().String("user-id", "", "filter by user id")

	cmd.AddCommand(get, create, del, update, list)
	return cmd
}
