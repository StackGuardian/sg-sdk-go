package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) policiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policies",
		Short: "Manage policies",
	}

	create := a.command("create", "Create a policy (body must include PolicyType: GENERAL or FILTER.INSIGHT)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.PolymorphicPolicy
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Policies.CreatePolicy(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get POLICY", "Get a policy", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Policies.ReadPolicy(ctx, a.org, args[0])
			return err
		})

	del := a.command("delete POLICY", "Delete a policy", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			return c.Policies.DeletePolicy(ctx, a.org, args[0])
		})

	update := a.command("update POLICY", "Update a policy (body must include PolicyType)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedPolymorphicPolicy
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Policies.UpdatePolicy(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	list := a.command("list", "List policies", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllPoliciesRequest{
				Description:      flagString(cmd, "description"),
				PolicyTypes:      flagString(cmd, "policy-types"),
				ResourceNames:    flagString(cmd, "resource-names"),
				Tags:             flagString(cmd, "tags"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				SearchQuery:      flagString(cmd, "search-query"),
			}
			_, err := c.Policies.ListAllPolicies(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("description", "", "filter by description")
	list.Flags().String("policy-types", "", "filter by policy types (default GENERAL)")
	list.Flags().String("resource-names", "", "filter by resource names")
	list.Flags().String("tags", "", "filter by tags")
	list.Flags().String("search-query", "", "search by resource name, tags or description")

	cmd.AddCommand(create, get, del, update, list)
	return cmd
}
