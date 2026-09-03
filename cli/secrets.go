package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/secrets"
)

func (a *app) secretsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secrets",
		Short: "Manage organization secrets",
	}

	create := a.command("create", "Create a secret", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.Secret
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Secrets.CreateSecret(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	del := a.command("delete SECRET", "Delete a secret", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Secrets.DeleteSecret(ctx, a.org, args[0])
			return err
		})

	update := a.command("update SECRET", "Update a secret", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedSecret
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			if req.ResourceName == nil {
				req.ResourceName = sgsdkgo.Optional(args[0]) // the API requires it on updates
			}
			_, err := c.Secrets.UpdateSecret(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	list := a.command("list", "List secrets", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Secrets.ListAllSecrets(ctx, a.org)
			return err
		})

	readBulk := a.command("read-bulk", "Read several secrets by name", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.Secrets.ReadSecretsInBulk(ctx, a.org, &secrets.ReadSecretsInBulkRequest{ResourceNames: flagStringSlice(cmd, "names")})
			return err
		})
	readBulk.Flags().StringSlice("names", nil, "secret names (comma-separated)")
	_ = readBulk.MarkFlagRequired("names")

	cmd.AddCommand(create, del, update, list, readBulk)
	return cmd
}
