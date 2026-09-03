package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/stacktemplaterevisions"
	"github.com/StackGuardian/sg-sdk-go/stacktemplates"
)

func (a *app) stackTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack-templates",
		Short: "Manage stack (IaC group) templates",
	}

	get := a.command("get TEMPLATE_ID", "Get a stack template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.StackTemplates.ReadStackTemplate(ctx, a.org, args[0])
			return err
		})

	create := a.command("create", "Create a stack template", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req stacktemplates.CreateStackTemplateRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackTemplates.CreateStackTemplate(ctx, a.org, getBool(cmd, "create-first-revision"), &req)
			return err
		})
	addBodyFlags(create)
	addCreateFirstRevisionFlag(create)

	update := a.command("update TEMPLATE_ID", "Update a stack template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req stacktemplates.UpdateStackTemplateRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackTemplates.UpdateStackTemplate(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete TEMPLATE_ID", "Delete a stack template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			return c.StackTemplates.DeleteStackTemplate(ctx, a.org, args[0])
		})

	cmd.AddCommand(get, create, update, del)
	return cmd
}

func (a *app) stackTemplateRevisionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack-template-revisions",
		Short: "Manage revisions of stack (IaC group) templates",
	}

	create := a.command("create TEMPLATE_ID", "Create a new revision of a stack template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req stacktemplaterevisions.CreateStackTemplateRevisionRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackTemplateRevisions.CreateStackTemplateRevision(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(create)

	get := a.command("get REVISION_ID", "Get a stack template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.StackTemplateRevisions.ReadStackTemplateRevision(ctx, a.org, args[0])
			return err
		})

	update := a.command("update REVISION_ID", "Update a stack template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req stacktemplaterevisions.UpdateStackTemplateRevisionRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackTemplateRevisions.UpdateStackTemplateRevision(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete REVISION_ID", "Delete a stack template revision", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			return c.StackTemplateRevisions.DeleteStackTemplateRevision(ctx, a.org, args[0], getBool(cmd, "keep-parent-template"))
		})
	addKeepParentTemplateFlag(del)

	cmd.AddCommand(create, get, update, del)
	return cmd
}
