package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/accessmanagement"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) roleBindingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "role-bindings",
		Short: "Inspect role bindings",
	}
	cmd.AddCommand(
		a.command("get ROLE_BINDING", "Get a role binding (for example \"default\")", cobra.ExactArgs(1),
			func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
				_, err := c.AccessManagement.ReadRoleBinding(ctx, a.org, args[0])
				return err
			}),
	)
	return cmd
}

func (a *app) apiTokensCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api-tokens",
		Short: "Manage legacy user API keys",
	}

	create := a.command("create", "Create or retrieve the caller's API key, or look up a runner group's key", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := accessmanagement.ApiTokenRequest{
				Regenerate:    flagBool(cmd, "regenerate"),
				RunnerGroupId: flagString(cmd, "runner-group-id"),
			}
			_, err := c.AccessManagement.CreateApiToken(ctx, a.org, &req)
			return err
		})
	create.Flags().Bool("regenerate", false, "replace the existing key with a new one")
	create.Flags().String("runner-group-id", "", "look up the key of this runner group (/runnergroups/<name>)")

	del := a.command("delete PRINCIPAL_ID", "Revoke the API key of a principal (\"<userPoolId>/<idp>/<email>\")", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.AccessManagement.DeleteApiToken(ctx, a.org, args[0])
			return err
		})

	cmd.AddCommand(create, del)
	return cmd
}
