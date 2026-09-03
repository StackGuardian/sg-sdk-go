package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/statebackends"
)

func (a *app) stateBackendsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state-backends",
		Short: "Manage state backends (S3 or Azure Blob storage for state files)",
	}

	create := a.command("create", "Create a state backend", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req statebackends.StateBackend
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StateBackends.CreateStateBackend(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get STATE_BACKEND", "Get a state backend", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.StateBackends.ReadStateBackend(ctx, a.org, args[0])
			return err
		})

	update := a.command("update STATE_BACKEND", "Update a state backend", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req statebackends.PatchedStateBackend
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StateBackends.UpdateStateBackend(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete STATE_BACKEND", "Delete a state backend", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.StateBackends.DeleteStateBackend(ctx, a.org, args[0])
			return err
		})

	list := a.command("list", "List state backends", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := statebackends.ListAllStateBackendsRequest{
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				SearchQuery:      flagString(cmd, "search-query"),
			}
			_, err := c.StateBackends.ListAllStateBackends(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("search-query", "", "search by resource name, tags or description")

	statefiles := a.command("list-statefiles", "List the state files in a storage backend", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.StateBackends.ListStatefiles(ctx, a.org, &statebackends.ListStatefilesRequest{Config: getString(cmd, "config")})
			return err
		})
	requiredString(statefiles, "config", "base64-encoded JSON StateBackendConfig")

	cmd.AddCommand(create, get, update, del, list, statefiles)
	return cmd
}
