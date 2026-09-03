package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/workflowgroups"
)

func (a *app) workflowGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workflow-groups",
		Aliases: []string{"wfgrps", "wfgrp"},
		Short:   "Manage workflow groups (nested groups are addressed as parent/child)",
	}

	create := a.command("create", "Create a workflow group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req sgsdkgo.WorkflowGroup
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowGroups.CreateWorkflowGroup(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get WORKFLOW_GROUP", "Get a workflow group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.WorkflowGroups.ReadWorkflowGroup(ctx, a.org, args[0])
			return err
		})

	del := a.command("delete WORKFLOW_GROUP", "Delete a workflow group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.WorkflowGroups.DeleteWorkflowGroup(ctx, a.org, args[0])
			return err
		})

	update := a.command("update WORKFLOW_GROUP", "Update a workflow group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedWorkflowGroup
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowGroups.UpdateWorkflowGroup(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	createChild := a.command("create-child PARENT_GROUP", "Create a child workflow group inside a parent group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.WorkflowGroup
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowGroups.CreateChildWorkflowGroup(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(createChild)

	listChildren := a.command("list-children PARENT_GROUP", "List the child workflow groups of a parent group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := sgsdkgo.ListAllChildWorkflowGroupsRequest{
				Description:      flagString(cmd, "description"),
				ResourceNames:    flagString(cmd, "resource-names"),
				Tags:             flagString(cmd, "tags"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
			}
			_, err := c.WorkflowGroups.ListAllChildWorkflowGroups(ctx, a.org, args[0], &req)
			return err
		})
	addPaginationFlags(listChildren)
	listChildren.Flags().String("description", "", "filter by description")
	listChildren.Flags().String("resource-names", "", "filter by resource names")
	listChildren.Flags().String("tags", "", "filter by tags")

	list := a.command("list", "List workflow groups", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllWorkflowGroupsRequest{
				Description:      flagString(cmd, "description"),
				ResourceNames:    flagString(cmd, "resource-names"),
				Tags:             flagString(cmd, "tags"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				SearchQuery:      flagString(cmd, "search-query"),
			}
			_, err := c.WorkflowGroups.ListAllWorkflowGroups(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("description", "", "filter by description")
	list.Flags().String("resource-names", "", "filter by resource names")
	list.Flags().String("tags", "", "filter by tags")
	list.Flags().String("search-query", "", "search by resource name, tags or description")

	listResources := a.command("list-resources WORKFLOW_GROUP", "List the workflows, stacks and child groups inside a workflow group", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := workflowgroups.ListAllResourcesRequest{
				Limit:            flagInt(cmd, "limit"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				ResourceNames:    flagString(cmd, "resource-names"),
				Description:      flagString(cmd, "description"),
				Tags:             flagString(cmd, "tags"),
				SearchQuery:      flagString(cmd, "search-query"),
				LatestStatuses:   flagString(cmd, "latest-statuses"),
				ResourceTypes:    flagString(cmd, "resource-types"),
				ContextTags:      flagString(cmd, "context-tags"),
				IsActive:         flagString(cmd, "is-active"),
			}
			_, err := c.WorkflowGroups.ListAllResourcesInWorkflowGroup(ctx, a.org, args[0], &req)
			return err
		})
	addPaginationFlags(listResources)
	listResources.Flags().String("resource-names", "", "filter by resource names (comma-separated)")
	listResources.Flags().String("description", "", "filter by description")
	listResources.Flags().String("tags", "", "filter by tags (comma-separated)")
	listResources.Flags().String("search-query", "", "search by resource name, description or tags")
	listResources.Flags().String("latest-statuses", "", "filter by latest status (comma-separated)")
	listResources.Flags().String("resource-types", "", "filter by resource type: WORKFLOW, STACK, WORKFLOW_GROUP (comma-separated)")
	listResources.Flags().String("context-tags", "", "filter by context tags as key or key:value (comma-separated)")
	listResources.Flags().String("is-active", "", "filter by active state (0 or 1)")

	cmd.AddCommand(create, get, del, update, createChild, listChildren, list, listResources)
	return cmd
}
