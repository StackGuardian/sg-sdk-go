package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/runnergroups"
)

func (a *app) runnerGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runner-groups",
		Short: "Manage private runner groups",
	}

	create := a.command("create", "Create a runner group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.RunnerGroup
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.RunnerGroups.CreateNewRunnerGroup(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get RUNNER_GROUP", "Get a runner group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := sgsdkgo.ReadRunnerGroupRequest{
				GetActiveWorkflows:        flagBool(cmd, "get-active-workflows"),
				GetActiveWorkflowsDetails: flagBool(cmd, "get-active-workflows-details"),
			}
			_, err := c.RunnerGroups.ReadRunnerGroup(ctx, a.org, args[0], &req)
			return err
		})
	get.Flags().Bool("get-active-workflows", false, "include the active workflows")
	get.Flags().Bool("get-active-workflows-details", false, "include details of the active workflows")

	del := a.command("delete RUNNER_GROUP", "Delete a runner group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.RunnerGroups.DeleteRunnerGroup(ctx, a.org, args[0])
			return err
		})

	update := a.command("update RUNNER_GROUP", "Update a runner group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedRunnerGroup
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.RunnerGroups.UpdateRunnerGroup(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	deregister := a.command("deregister-runner RUNNER_GROUP", "Deregister a runner from a runner group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.RunnerDeregister
			if v := flagString(cmd, "runner-id"); v != nil {
				req.RunnerId = sgsdkgo.Optional(*v)
			}
			if v := flagStringSlice(cmd, "container-instance-ids"); v != nil {
				req.ContainerInstanceIds = sgsdkgo.Optional(v)
			}
			if v := flagBool(cmd, "force"); v != nil {
				req.ForceDeregister = sgsdkgo.Optional(*v)
			}
			return c.RunnerGroups.DeregisterRunner(ctx, a.org, args[0], &req)
		})
	deregister.Flags().String("runner-id", "", "runner id")
	deregister.Flags().StringSlice("container-instance-ids", nil, "container instance ids (comma-separated)")
	deregister.Flags().Bool("force", false, "force deregistration")

	state := a.command("update-runner-state RUNNER_GROUP", "Update the state of a runner (ACTIVE or DRAINING)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			status := getString(cmd, "status")
			if err := checkEnum("status", status, "ACTIVE", "DRAINING"); err != nil {
				return err
			}
			req := sgsdkgo.RunnerStatus{Status: sgsdkgo.StatusEnum(status)}
			if v := flagString(cmd, "runner-id"); v != nil {
				req.RunnerId = sgsdkgo.Optional(*v)
			}
			if v := flagStringSlice(cmd, "container-instance-ids"); v != nil {
				req.ContainerInstanceIds = sgsdkgo.Optional(v)
			}
			return c.RunnerGroups.UpdateRunnerState(ctx, a.org, args[0], &req)
		})
	requiredString(state, "status", "new runner status (ACTIVE, DRAINING)")
	state.Flags().String("runner-id", "", "runner id")
	state.Flags().StringSlice("container-instance-ids", nil, "container instance ids (comma-separated)")

	list := a.command("list", "List runner groups", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := runnergroups.ListAllRunnerGroupsRequest{
				Limit:               flagInt(cmd, "limit"),
				Lastevaluatedkey:    flagString(cmd, "last-evaluated-key"),
				FilterResourceNames: flagString(cmd, "resource-names"),
			}
			_, err := c.RunnerGroups.ListAllRunnerGroups(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("resource-names", "", "filter by resource names")

	register := a.command("register RUNNER_GROUP", "Register a runner node (use the runner group's own token as the API key)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.RunnerGroups.RegisterRunner(ctx, a.org, args[0])
			return err
		})

	storageAuth := a.command("storage-backend-auth RUNNER_GROUP", "Generate short-lived credentials for the runner group's storage backend", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.RunnerGroups.GenerateStorageBackendAuth(ctx, a.org, args[0])
			return err
		})

	cmd.AddCommand(create, get, del, update, deregister, state, list, register, storageAuth)
	return cmd
}
