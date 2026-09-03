package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/workflowsteptemplate"
	"github.com/StackGuardian/sg-sdk-go/workflowsteptemplaterevision"
)

func addCreateFirstRevisionFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("create-first-revision", true, "also create the first revision of the template")
}

func addKeepParentTemplateFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("keep-parent-template", false, "keep the parent template when deleting its last revision")
}

func (a *app) workflowStepTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workflow-step-templates",
		Short: "Manage workflow step templates (WORKFLOW_STEP)",
	}

	get := a.command("get TEMPLATE_ID", "Get a workflow step template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.WorkflowStepTemplate.ReadWorkflowStepTemplate(ctx, a.org, args[0])
			return err
		})

	create := a.command("create", "Create a workflow step template", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req workflowsteptemplate.CreateWorkflowStepTemplate
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowStepTemplate.CreateWorkflowStepTemplate(ctx, a.org, getBool(cmd, "create-first-revision"), &req)
			return err
		})
	addBodyFlags(create)
	addCreateFirstRevisionFlag(create)

	update := a.command("update TEMPLATE_ID", "Update a workflow step template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflowsteptemplate.UpdateWorkflowStepTemplateRequestModel
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowStepTemplate.UpdateWorkflowStepTemplate(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete TEMPLATE_ID", "Delete a workflow step template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			return c.WorkflowStepTemplate.DeleteWorkflowStepTemplate(ctx, a.org, args[0])
		})

	cmd.AddCommand(get, create, update, del)
	return cmd
}

func (a *app) workflowStepTemplateRevisionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workflow-step-template-revisions",
		Short: "Manage revisions of workflow step templates",
	}

	create := a.command("create TEMPLATE_ID", "Create a new revision of a workflow step template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflowsteptemplaterevision.CreateWorkflowStepTemplateRevisionModel
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowStepTemplateRevision.CreateWorkflowStepTemplateRevision(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(create)

	update := a.command("update REVISION_ID", "Update a workflow step template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflowsteptemplaterevision.UpdateWorkflowStepTemplateRevisionModel
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.WorkflowStepTemplateRevision.UpdateWorkflowStepTemplateRevision(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	get := a.command("get REVISION_ID", "Get a workflow step template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.WorkflowStepTemplateRevision.ReadWorkflowStepTemplateRevision(ctx, a.org, args[0])
			return err
		})

	del := a.command("delete REVISION_ID", "Delete a workflow step template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			return c.WorkflowStepTemplateRevision.DeleteWorkflowStepTemplateRevision(ctx, a.org, args[0], getBool(cmd, "keep-parent-template"))
		})
	addKeepParentTemplateFlag(del)

	cmd.AddCommand(create, update, get, del)
	return cmd
}
