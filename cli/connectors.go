package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/connectors"
)

func (a *app) connectorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connectors",
		Short: "Manage connectors (cloud and VCS integrations)",
	}

	create := a.command("create", "Create a connector", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.Integration
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Connectors.CreateConnector(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get CONNECTOR", "Get a connector", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Connectors.ReadConnector(ctx, args[0], a.org)
			return err
		})

	del := a.command("delete CONNECTOR", "Delete a connector", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Connectors.DeleteConnector(ctx, args[0], a.org)
			return err
		})

	update := a.command("update CONNECTOR", "Update a connector", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedIntegration
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Connectors.UpdateConnector(ctx, args[0], a.org, &req)
			return err
		})
	addBodyFlags(update)

	list := a.command("list", "List connectors", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllConnectorsRequest{
				ResourceNames:    flagString(cmd, "resource-names"),
				Kind:             flagString(cmd, "kind"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
			}
			_, err := c.Connectors.ListAllConnectors(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("resource-names", "", "filter by resource name")
	list.Flags().String("kind", "", "filter by kind (comma-separated), e.g. AWS_RBAC,AZURE_OIDC")

	authenticate := a.command("authenticate CONNECTOR", "Verify that a connector can authenticate against its provider", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Connectors.AuthenticateConnector(ctx, args[0], a.org)
			return err
		})

	listAccounts := a.command("list-accounts", "List the cloud accounts or subscriptions reachable with a connector's credentials", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req connectors.ListAccountsRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Connectors.ListAccounts(ctx, a.org, &req)
			return err
		})
	addBodyFlags(listAccounts)

	githubRepos := a.command("github-repos CONNECTOR", "List the repositories accessible to a GITHUB_COM connector", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Connectors.GetGithubComRepos(ctx, args[0], a.org)
			return err
		})

	repos := a.command("repos CONNECTOR", "List the repositories reachable through a VCS connector", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := connectors.ListRepositoriesRequest{
				Login:          flagString(cmd, "login"),
				InstallationId: flagString(cmd, "installation-id"),
				Page:           flagString(cmd, "page"),
				Limit:          flagInt(cmd, "limit"),
				SearchQuery:    flagString(cmd, "search-query"),
			}
			_, err := c.Connectors.ListRepositories(ctx, args[0], a.org, &req)
			return err
		})
	repos.Flags().String("login", "", "GitHub account or organization login (GITHUB_COM)")
	repos.Flags().String("installation-id", "", "GitHub App installation id (GITHUB_COM)")
	repos.Flags().String("page", "", "page number or provider cursor")
	repos.Flags().Int("limit", 0, "page size")
	repos.Flags().String("search-query", "", "filter by name or namespace")

	cmd.AddCommand(create, get, del, update, list, authenticate, listAccounts, githubRepos, repos)
	return cmd
}
