package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/templates"
	"github.com/StackGuardian/sg-sdk-go/workflows"
)

var (
	templateTypes     = []string{"IAC", "IAC_GROUP", "IAC_POLICY", "WORKFLOW_STEP"}
	subscriptionTypes = []string{"IACGroupSubscriptions", "IACSubscriptions", "PolicySubscriptions", "WfStepSubscriptions"}
)

// ownerOrg returns --owner-org, falling back to the current organization.
func (a *app) ownerOrg(cmd *cobra.Command) string {
	if v := getString(cmd, "owner-org"); v != "" {
		return v
	}
	return a.org
}

func addOwnerOrgFlag(cmd *cobra.Command) {
	cmd.Flags().String("owner-org", "", "organization that owns the template (defaults to --org)")
}

func addTemplateTypeFlag(cmd *cobra.Command) {
	requiredString(cmd, "type", "template type (IAC, IAC_GROUP, IAC_POLICY, WORKFLOW_STEP)")
}

func templateType(cmd *cobra.Command) (string, error) {
	v := getString(cmd, "type")
	return v, checkEnum("type", v, templateTypes...)
}

func (a *app) templatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "templates",
		Short: "Manage templates and template revisions across all template types",
		Long: `Manage templates and template revisions across all template types.

A template revision is addressed as NAME or NAME:REVISION (for example
my-template:5). Without a revision the base template is used.`,
	}

	list := a.command("list", "List templates of a type", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			req := sgsdkgo.ListAllTemplatesRequest{
				SgOrgid:          a.org,
				IsPublic:         flagString(cmd, "is-public"),
				OwnerOrgs:        flagString(cmd, "owner-orgs"),
				SearchQuery:      flagString(cmd, "search-query"),
				TemplateId:       flagString(cmd, "template-id"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
			}
			_, err = c.Templates.ListAllTemplates(ctx, sgsdkgo.ListAllTemplatesRequestTemplateType(t), &req)
			return err
		})
	addTemplateTypeFlag(list)
	list.Flags().String("is-public", "", "0 for private, 1 for public templates")
	list.Flags().String("owner-orgs", "", "organizations that own the templates (comma-separated)")
	list.Flags().String("search-query", "", "search by name, tags or description")
	list.Flags().String("template-id", "", "parent template id to list its revisions")
	list.Flags().String("last-evaluated-key", "", "pagination token from a previous response")

	listByOwner := a.command("list-by-owner OWNER_ORG", "List templates owned by an organization", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := sgsdkgo.ListAllTemplatesBasedOnOwnerOrgRequest{
				SgOrgid:          a.org,
				IsPublic:         flagString(cmd, "is-public"),
				SearchQuery:      flagString(cmd, "search-query"),
				IsSharedTemplate: flagBool(cmd, "shared"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
			}
			_, err := c.Templates.ListAllTemplatesBasedOnOwnerOrg(ctx, args[0], &req)
			return err
		})
	listByOwner.Flags().String("is-public", "", "0 for private, 1 for public templates")
	listByOwner.Flags().String("search-query", "", "search by name, tags or description")
	listByOwner.Flags().Bool("shared", false, "list shared templates")
	listByOwner.Flags().String("last-evaluated-key", "", "pagination token from a previous response")

	create := a.command("create", "Create a template or a new revision of it (body must include TemplateType: IAC, IAC_GROUP or IAC_POLICY)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var body sgsdkgo.CreateTemplateRequest
			if err := readBody(cmd, &body); err != nil {
				return err
			}
			_, err := c.Templates.CreateTemplateRevision(ctx, &sgsdkgo.CreateTemplateRevisionRequest{SgOrgid: a.org, Body: &body})
			return err
		})
	addBodyFlags(create)

	get := a.command("get TEMPLATE[:REVISION]", "Get a template or one of its revisions", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			_, err = c.Templates.ReadTemplateRevision(ctx, a.ownerOrg(cmd), args[0],
				sgsdkgo.ReadTemplateRevisionRequestTemplateType(t), &sgsdkgo.ReadTemplateRevisionRequest{SgOrgid: a.org})
			return err
		})
	addTemplateTypeFlag(get)
	addOwnerOrgFlag(get)

	del := a.command("delete TEMPLATE[:REVISION]", "Delete a template or one of its revisions", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			return c.Templates.DeleteTemplateRevision(ctx, a.ownerOrg(cmd), args[0], t, &sgsdkgo.DeleteTemplateRevisionRequest{SgOrgid: a.org})
		})
	addTemplateTypeFlag(del)
	addOwnerOrgFlag(del)

	update := a.command("update TEMPLATE[:REVISION]", "Update a template or one of its revisions", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			var req sgsdkgo.PatchedTemplateUpdate
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			req.SgOrgid = a.org
			_, err = c.Templates.UpdateTemplateRevision(ctx, a.ownerOrg(cmd), args[0], t, &req)
			return err
		})
	addTemplateTypeFlag(update)
	addOwnerOrgFlag(update)
	addBodyFlags(update)

	groupIac := a.command("get-group-iac GROUP_TEMPLATE IAC_TEMPLATE_ID", "Get an IaC template that belongs to an IaC group template", cobra.ExactArgs(2),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Templates.ReadIacGroupsIacTemplate(ctx, a.ownerOrg(cmd), args[1], args[0], &sgsdkgo.ReadIacGroupsIacTemplateRequest{SgOrgid: a.org})
			return err
		})
	addOwnerOrgFlag(groupIac)

	inputSchema := a.command("input-schema TEMPLATE[:REVISION]", "Decode one of a template's input schemas (requires a user token)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			_, err = c.Templates.GetTemplateInputSchema(ctx, a.ownerOrg(cmd), t, args[0], &templates.GetTemplateInputSchemaRequest{SchemaType: getString(cmd, "schema-type")})
			return err
		})
	addTemplateTypeFlag(inputSchema)
	addOwnerOrgFlag(inputSchema)
	requiredString(inputSchema, "schema-type", "schema to decode (FORM_JSONSCHEMA, RAW_JSON, NO_CODE_JSON, TIRITH_JSON)")

	vcsTriggers := a.command("vcs-triggers TEMPLATE", "Create VCS webhook triggers for a template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			var req workflows.CreateVcsTriggersRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err = c.Templates.CreateTemplateVcsTriggers(ctx, a.ownerOrg(cmd), t, args[0], &req)
			return err
		})
	addTemplateTypeFlag(vcsTriggers)
	addOwnerOrgFlag(vcsTriggers)
	addBodyFlags(vcsTriggers)

	cmd.AddCommand(list, listByOwner, create, get, del, update, groupIac, inputSchema, vcsTriggers,
		a.templateSubscriptionsCmd(), a.templateArtifactsCmd(), a.publicTemplatesCmd())
	return cmd
}

func (a *app) templateSubscriptionsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "subscriptions", Short: "Manage the organization's template subscriptions"}

	get := a.command("get", "Get the subscriptions of a type", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			v := getString(cmd, "type")
			if err := checkEnum("type", v, subscriptionTypes...); err != nil {
				return err
			}
			_, err := c.Templates.ReadSubscription(ctx, a.org, &sgsdkgo.ReadSubscriptionRequest{
				SubscriptionType: sgsdkgo.ReadSubscriptionRequestSubscriptionType(v),
			})
			return err
		})
	requiredString(get, "type", "subscription type (IACGroupSubscriptions, IACSubscriptions, PolicySubscriptions, WfStepSubscriptions)")

	create := a.command("create", "Create a subscription set (all four subscription maps are required)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req templates.CreateSubscriptionRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Templates.CreateSubscription(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	update := a.command("update SUBSCRIPTION", "Subscribe to or unsubscribe from templates", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			action := getString(cmd, "action")
			if err := checkEnum("action", action, "subscribe", "unsubscribe"); err != nil {
				return err
			}
			var req templates.UpdateSubscriptionRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Templates.UpdateSubscription(ctx, a.org, args[0], action, &req)
			return err
		})
	requiredString(update, "action", "subscribe or unsubscribe")
	addBodyFlags(update)

	cmd.AddCommand(get, create, update)
	return cmd
}

func (a *app) templateArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "artifacts", Short: "Work with template artifacts through presigned URLs"}
	scope := func(cmd *cobra.Command) {
		addTemplateTypeFlag(cmd)
		addOwnerOrgFlag(cmd)
	}

	list := a.command("list TEMPLATE", "List the artifacts of a template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			req := templates.ListAllTemplateArtifactsRequest{
				SgOrgid:        a.org,
				ArtifactPrefix: flagString(cmd, "prefix"),
				MaxKeys:        flagInt(cmd, "max-keys"),
				StartAfterKey:  flagString(cmd, "start-after-key"),
			}
			_, err = c.Templates.ListAllTemplateArtifacts(ctx, a.ownerOrg(cmd), t, args[0], &req)
			return err
		})
	scope(list)
	list.Flags().String("prefix", "", "only list artifacts under this prefix")
	list.Flags().Int("max-keys", 0, "maximum number of keys")
	list.Flags().String("start-after-key", "", "pagination: start after this key")

	downloadURL := a.command("download-url TEMPLATE ARTIFACT_ID", "Get a presigned download URL for an artifact", cobra.ExactArgs(2),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			_, err = c.Templates.GetTemplateArtifactDownloadUrl(ctx, a.ownerOrg(cmd), t, args[0], args[1], &templates.TemplateArtifactRequest{SgOrgid: a.org})
			return err
		})
	scope(downloadURL)

	uploadURL := a.command("upload-url TEMPLATE ARTIFACT_ID", "Get a presigned upload URL for an artifact", cobra.ExactArgs(2),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			_, err = c.Templates.GetTemplateArtifactUploadUrl(ctx, a.ownerOrg(cmd), t, args[0], args[1], &templates.TemplateArtifactRequest{SgOrgid: a.org})
			return err
		})
	scope(uploadURL)

	del := a.command("delete TEMPLATE ARTIFACT_ID", "Delete an artifact", cobra.ExactArgs(2),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			_, err = c.Templates.DeleteTemplateArtifact(ctx, a.ownerOrg(cmd), t, args[0], args[1], &templates.TemplateArtifactRequest{SgOrgid: a.org})
			return err
		})
	scope(del)

	cmd.AddCommand(list, downloadURL, uploadURL, del)
	return cmd
}

func (a *app) publicTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "public", Short: "Browse public marketplace templates (requires a user token)"}

	list := a.command("list", "List public templates of a type", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			req := templates.ListAllPublicTemplatesRequest{
				Limit:             flagInt(cmd, "limit"),
				Lastevaluatedkey:  flagString(cmd, "last-evaluated-key"),
				TemplateId:        flagString(cmd, "template-id"),
				OwnerOrgs:         flagString(cmd, "owner-orgs"),
				SearchQuery:       flagString(cmd, "search-query"),
				ContextTags:       flagString(cmd, "context-tags"),
				SourceConfigKinds: flagString(cmd, "source-config-kinds"),
			}
			_, err = c.Templates.ListAllPublicTemplates(ctx, t, &req)
			return err
		})
	list.Annotations = map[string]string{annotationOrg: orgOptional}
	addTemplateTypeFlag(list)
	addPaginationFlags(list)
	list.Flags().String("template-id", "", "parent template id (with owner org) to list its revisions")
	list.Flags().String("owner-orgs", "", "owner organizations (comma-separated)")
	list.Flags().String("search-query", "", "search by name, description or tags")
	list.Flags().String("context-tags", "", "filter by context tags as key or key:value (comma-separated)")
	list.Flags().String("source-config-kinds", "", "filter by source config kind (comma-separated)")

	listAll := a.command("list-all", "List public templates across all types", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := templates.ListAllPublicTemplatesV2Request{
				Limit:             flagInt(cmd, "limit"),
				Lastevaluatedkey:  flagString(cmd, "last-evaluated-key"),
				IsActive:          flagString(cmd, "is-active"),
				TemplateName:      flagString(cmd, "template-name"),
				GitHubComRepoID:   flagString(cmd, "github-repo-id"),
				SearchQuery:       flagString(cmd, "search-query"),
				Tags:              flagString(cmd, "tags"),
				TemplateTypes:     flagString(cmd, "template-types"),
				SourceConfigKinds: flagString(cmd, "source-config-kinds"),
				ContextTags:       flagString(cmd, "context-tags"),
				OwnerOrgs:         flagString(cmd, "owner-orgs"),
			}
			_, err := c.Templates.ListAllPublicTemplatesV2(ctx, &req)
			return err
		})
	listAll.Annotations = map[string]string{annotationOrg: orgOptional}
	addPaginationFlags(listAll)
	listAll.Flags().String("is-active", "", "filter by active state (0 or 1)")
	listAll.Flags().String("template-name", "", "filter by template name")
	listAll.Flags().String("github-repo-id", "", "filter by GitHub repository id")
	listAll.Flags().String("search-query", "", "search by name, description, tags or id")
	listAll.Flags().String("tags", "", "filter by tags (comma-separated)")
	listAll.Flags().String("template-types", "", "filter by template types (comma-separated)")
	listAll.Flags().String("source-config-kinds", "", "filter by source config kind (comma-separated)")
	listAll.Flags().String("context-tags", "", "filter by context tags as key or key:value (comma-separated)")
	listAll.Flags().String("owner-orgs", "", "owner organizations (comma-separated)")

	get := a.command("get TEMPLATE[:REVISION]", "Get a public template", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			t, err := templateType(cmd)
			if err != nil {
				return err
			}
			_, err = c.Templates.ReadPublicTemplate(ctx, getString(cmd, "owner-org"), t, args[0])
			return err
		})
	get.Annotations = map[string]string{annotationOrg: orgOptional}
	addTemplateTypeFlag(get)
	requiredString(get, "owner-org", "organization that owns the template")

	cmd.AddCommand(list, listAll, get)
	return cmd
}
