package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/connectorgroups"
)

func (a *app) connectorGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connector-groups",
		Short: "Manage connector groups (cloud integration groups)",
	}

	create := a.command("create", "Create a connector group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.IntegrationGroups
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.ConnectorGroups.CreateConnectorGroup(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get GROUP", "Get a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.ReadConnectorGroup(ctx, args[0], a.org)
			return err
		})

	del := a.command("delete GROUP", "Delete a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.DeleteConnectorGroup(ctx, args[0], a.org)
			return err
		})

	update := a.command("update GROUP", "Update a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedIntegrationGroups
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.ConnectorGroups.UpdateConnectorGroup(ctx, args[0], a.org, &req)
			return err
		})
	addBodyFlags(update)

	authenticate := a.command("authenticate GROUP", "Authenticate a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.AuthenticateConnectorGroup(ctx, args[0], a.org)
			return err
		})

	list := a.command("list", "List connector groups", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllConnectorGroupsRequest{
				ResourceNames:    flagString(cmd, "resource-names"),
				Kind:             flagString(cmd, "kind"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				Scope:            flagString(cmd, "scope"),
			}
			_, err := c.ConnectorGroups.ListAllConnectorGroups(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("resource-names", "", "filter by resource name")
	list.Flags().String("kind", "", "filter by kind")
	list.Flags().String("scope", "", "filter by scope")

	scan := a.command("discovery-scan GROUP", "Trigger discovery scans across the group's child connectors", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.TriggerDiscoveryScan(ctx, args[0], a.org, &connectorgroups.DiscoveryScanRequest{Benchmark: flagString(cmd, "benchmark")})
			return err
		})
	scan.Flags().String("benchmark", "", "benchmarks to scan (comma-separated, default inventory)")

	cmd.AddCommand(create, get, del, update, authenticate, list, scan, a.connectorGroupChildrenCmd())
	return cmd
}

// connectorGroupChildrenCmd manages the connectors generated inside a cloud connector group.
func (a *app) connectorGroupChildrenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "connectors",
		Aliases: []string{"child", "children"},
		Short:   "Manage the connectors inside a connector group",
	}

	get := a.command("get CONNECTOR", "Get a connector in a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.ReadChildInACloudConnectorGroup(ctx, args[0], getString(cmd, "group"), a.org)
			return err
		})
	requiredString(get, "group", "connector group name")

	del := a.command("delete CONNECTOR", "Delete a connector from a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.DeleteChildInACloudConnectorGroup(ctx, args[0], getString(cmd, "group"), a.org)
			return err
		})
	requiredString(del, "group", "connector group name")

	update := a.command("update CONNECTOR", "Update a connector in a connector group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedIntegration
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.ConnectorGroups.UpdateChildInACloudConnectorGroup(ctx, args[0], getString(cmd, "group"), a.org, &req)
			return err
		})
	requiredString(update, "group", "connector group name")
	addBodyFlags(update)

	list := a.command("list", "List the connectors in a connector group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllConnectorsInAGroupRequest{
				ResourceNames:    flagString(cmd, "resource-names"),
				Kind:             flagString(cmd, "kind"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				Scope:            flagString(cmd, "scope"),
			}
			_, err := c.ConnectorGroups.ListAllConnectorsInAGroup(ctx, getString(cmd, "group"), a.org, &req)
			return err
		})
	requiredString(list, "group", "connector group name")
	addPaginationFlags(list)
	list.Flags().String("resource-names", "", "filter by resource name")
	list.Flags().String("kind", "", "filter by kind")
	list.Flags().String("scope", "", "filter by scope")

	authenticate := a.command("authenticate CONNECTOR", "Verify that a connector in the group can authenticate against its provider", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.ConnectorGroups.AuthenticateChildInACloudConnectorGroup(ctx, args[0], getString(cmd, "group"), a.org)
			return err
		})
	requiredString(authenticate, "group", "connector group name")

	cmd.AddCommand(get, del, update, list, authenticate)
	return cmd
}
