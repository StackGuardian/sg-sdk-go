package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/workflowtemplates"
)

func (a *app) workflowTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workflow-templates",
		Short: "Manage workflow (IaC) templates",
	}

	get := a.command("get TEMPLATE_ID", "Get a workflow template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.WorkflowTemplates.ReadWorkflowTemplate(ctx, a.org, args[0])
			return err
		})

	create := a.command("create", "Create a workflow template", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req workflowtemplates.CreateWorkflowTemplateRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowTemplates.CreateWorkflowTemplate(ctx, a.org, getBool(cmd, "create-first-revision"), &req)
			return err
		})
	addBodyFlags(create)
	addCreateFirstRevisionFlag(create)

	update := a.command("update TEMPLATE_ID", "Update a workflow template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflowtemplates.UpdateWorkflowTemplateRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowTemplates.UpdateWorkflowTemplate(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete TEMPLATE_ID", "Delete a workflow template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			return c.WorkflowTemplates.DeleteWorkflowTemplate(ctx, a.org, args[0])
		})

	cmd.AddCommand(get, create, update, del)
	return cmd
}

func (a *app) workflowTemplateRevisionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workflow-template-revisions",
		Short: "Manage revisions of workflow (IaC) templates",
	}

	create := a.command("create TEMPLATE_ID", "Create a new revision of a workflow template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflowtemplaterevisions.CreateWorkflowTemplateRevisionsRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowTemplatesRevisions.CreateWorkflowTemplateRevision(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(create)

	del := a.command("delete REVISION_ID", "Delete a workflow template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			return c.WorkflowTemplatesRevisions.DeleteWorkflowTemplateRevision(ctx, a.org, args[0], getBool(cmd, "keep-parent-template"))
		})
	addKeepParentTemplateFlag(del)

	get := a.command("get REVISION_ID", "Get a workflow template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.WorkflowTemplatesRevisions.ReadWorkflowTemplateRevision(ctx, a.org, args[0])
			return err
		})

	update := a.command("update REVISION_ID", "Update a workflow template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflowtemplaterevisions.UpdateWorkflowTemplateRevisionRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowTemplatesRevisions.UpdateWorkflowTemplateRevision(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	cmd.AddCommand(create, del, get, update)
	return cmd
}
