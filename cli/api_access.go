package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) apiAccessCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api-access",
		Short: "Manage API accesses (API keys and OIDC identities)",
	}

	create := a.command("create", "Create an API access", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.ApiAccess
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.AccessManagement.CreateApiAccess(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get ACCESS_ID", "Get an API access", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.AccessManagement.ReadApiAccess(ctx, args[0], a.org)
			return err
		})

	del := a.command("delete ACCESS_ID", "Delete an API access", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.AccessManagement.DeleteApiAccess(ctx, args[0], a.org)
			return err
		})

	update := a.command("update ACCESS_ID", "Update an API access", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedApiAccessPatch
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.AccessManagement.UpdateApiAccess(ctx, args[0], a.org, &req)
			return err
		})
	addBodyFlags(update)

	regenerate := a.command("regenerate-key ACCESS_ID", "Regenerate the API key of an API access", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.ApiAccessRegenerate
			if v := flagInt(cmd, "expires-at"); v != nil {
				req.ExpiresAt = sgsdkgo.Optional(*v)
			}
			if getBool(cmd, "no-expiration") {
				req.ExpiresAt = sgsdkgo.Null[int]()
			}
			_, err := c.AccessManagement.RegenerateApiKey(ctx, args[0], a.org, &req)
			return err
		})
	regenerate.Flags().Int("expires-at", 0, "new expiration as a Unix timestamp in milliseconds")
	regenerate.Flags().Bool("no-expiration", false, "remove the expiration (sends an explicit null)")
	regenerate.MarkFlagsMutuallyExclusive("expires-at", "no-expiration")

	list := a.command("list", "List API accesses", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllApiAccessesRequest{
				LastEvaluatedKey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				ResourceName:     flagString(cmd, "resource-name"),
				Roles:            flagString(cmd, "roles"),
				Status:           flagString(cmd, "status"),
			}
			if v := flagString(cmd, "access-type"); v != nil {
				if err := checkEnum("access-type", *v, "APIKEY", "OIDC"); err != nil {
					return err
				}
				t := sgsdkgo.ListAllApiAccessesRequestAccessType(*v)
				req.AccessType = &t
			}
			_, err := c.AccessManagement.ListAllApiAccesses(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("access-type", "", "filter by access type (APIKEY, OIDC)")
	list.Flags().String("resource-name", "", "filter by resource name (contains match)")
	list.Flags().String("roles", "", "filter by roles (comma-separated)")
	list.Flags().String("status", "", "filter by status (comma-separated: active, expired)")

	cmd.AddCommand(create, get, del, update, regenerate, list)
	return cmd
}
