package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/organizations"
)

func (a *app) organizationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "organizations",
		Aliases: []string{"org", "orgs"},
		Short:   "Manage organizations",
	}

	get := a.command("get", "Get the organization identified by --org", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Organizations.ReadOrganization(ctx, a.org)
			return err
		})

	create := a.command("create", "Create an organization (requires a user token, not an API key)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req organizations.CreateOrganizationRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Organizations.CreateOrganization(ctx, &req)
			return err
		})
	create.Annotations = map[string]string{annotationOrg: orgOptional}
	addBodyFlags(create)

	list := a.command("list", "List the organizations the caller belongs to (requires a user token)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Organizations.ListAllOrganizations(ctx)
			return err
		})
	list.Annotations = map[string]string{annotationOrg: orgOptional}

	update := a.command("update", "Update the organization's settings", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req organizations.PatchedOrganization
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Organizations.UpdateOrganization(ctx, a.org, &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete", "Archive the organization and all of its resources", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Organizations.DeleteOrganization(ctx, a.org)
			return err
		})

	orgWorkflowsRequest := func(cmd *cobra.Command) *organizations.ListAllOrganizationWorkflowsRequest {
		return &organizations.ListAllOrganizationWorkflowsRequest{
			Limit:                           flagInt(cmd, "limit"),
			Lastevaluatedkey:                flagString(cmd, "last-evaluated-key"),
			FetchSgOwnedWfs:                 flagBool(cmd, "fetch-sg-owned-wfs"),
			LatestWfRunStatuses:             flagString(cmd, "latest-wf-run-statuses"),
			LatestDriftRunStatuses:          flagString(cmd, "latest-drift-run-statuses"),
			RepoId:                          flagString(cmd, "repo-id"),
			UseMarketplaceTemplate:          flagString(cmd, "use-marketplace-template"),
			IacTemplateId:                   flagString(cmd, "iac-template-id"),
			ParentIacTemplateId:             flagString(cmd, "parent-iac-template-id"),
			TerraformConfigTerraformVersion: flagString(cmd, "terraform-version"),
			RunnerNames:                     flagString(cmd, "runner-names"),
			TerraformConfigDriftCheck:       flagString(cmd, "drift-check"),
			ResourceNames:                   flagString(cmd, "resource-names"),
			Description:                     flagString(cmd, "description"),
			Tags:                            flagString(cmd, "tags"),
			IsActive:                        flagString(cmd, "is-active"),
		}
	}
	addOrgWorkflowFlags := func(cmd *cobra.Command) {
		addPaginationFlags(cmd)
		cmd.Flags().Bool("fetch-sg-owned-wfs", true, "include workflows owned by StackGuardian")
		cmd.Flags().String("latest-wf-run-statuses", "", "filter by latest run status (comma-separated)")
		cmd.Flags().String("latest-drift-run-statuses", "", "filter by latest drift run status (comma-separated)")
		cmd.Flags().String("repo-id", "", "filter by VCS repository")
		cmd.Flags().String("use-marketplace-template", "", "filter by marketplace template usage (true/false)")
		cmd.Flags().String("iac-template-id", "", "filter by IaC template id")
		cmd.Flags().String("parent-iac-template-id", "", "filter by parent IaC template id (any revision)")
		cmd.Flags().String("terraform-version", "", "filter by Terraform version")
		cmd.Flags().String("runner-names", "", "filter by runner names (comma-separated)")
		cmd.Flags().String("drift-check", "", "any value restricts to workflows with drift information")
		cmd.Flags().String("resource-names", "", "filter by resource names (comma-separated)")
		cmd.Flags().String("description", "", "filter by description (comma-separated)")
		cmd.Flags().String("tags", "", "filter by tags (comma-separated)")
		cmd.Flags().String("is-active", "", "filter by active state (0 or 1)")
	}

	workflows := a.command("workflows", "List workflows across every workflow group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.Organizations.ListAllWorkflows(ctx, a.org, orgWorkflowsRequest(cmd))
			return err
		})
	addOrgWorkflowFlags(workflows)

	count := a.command("count-workflows", "Count workflows across every workflow group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.Organizations.CountWorkflows(ctx, a.org, orgWorkflowsRequest(cmd))
			return err
		})
	addOrgWorkflowFlags(count)

	bulk := a.command("bulk-action", "Activate or deactivate workflows and stacks in bulk", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			action := getString(cmd, "action")
			if err := checkEnum("action", action, "activate", "deactivate"); err != nil {
				return err
			}
			req := organizations.WorkflowsBulkActionRequest{Action: action, ResourceIds: flagStringSlice(cmd, "resource-ids")}
			_, err := c.Organizations.BulkActionWorkflows(ctx, a.org, &req)
			return err
		})
	requiredString(bulk, "action", "activate or deactivate")
	bulk.Flags().StringSlice("resource-ids", nil, "resource paths relative to the organization, e.g. /wfgrps/g/wfs/w (comma-separated)")
	_ = bulk.MarkFlagRequired("resource-ids")

	cmd.AddCommand(get, create, list, update, del, workflows, count, bulk)
	return cmd
}
