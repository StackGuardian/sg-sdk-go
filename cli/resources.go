package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/resources"
)

func (a *app) resourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resources",
		Short: "Move and search resources across the organization",
	}

	move := a.command("move", "Move workflows, stacks or workflow groups to another parent", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req resources.MoveResourcesRequest
			if hasBody(cmd) {
				if err := readBody(cmd, &req); err != nil {
					return err
				}
			} else {
				req.SourceResourceIds = flagStringSlice(cmd, "source-ids")
				req.DestinationParentId = getString(cmd, "destination")
			}
			_, err := c.Resources.MoveResources(ctx, a.org, &req)
			return err
		})
	addBodyFlags(move)
	move.Flags().StringSlice("source-ids", nil, "resource ids to move (comma-separated)")
	move.Flags().String("destination", "", "id of the destination parent")
	move.MarkFlagsRequiredTogether("source-ids", "destination")
	move.MarkFlagsOneRequired("source-ids", "body", "body-file")

	search := a.command("search", "Search resources with a JSON filter body", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req resources.SearchResourcesRequest
			if err := readOptionalBody(cmd, &req); err != nil {
				return err
			}
			if v := flagStringSlice(cmd, "resource-types"); v != nil {
				req.ResourceTypes = v
			}
			if v := flagString(cmd, "search-query"); v != nil {
				req.SearchQuery = v
			}
			if v := flagInt(cmd, "limit"); v != nil {
				req.Limit = v
			}
			if v := flagString(cmd, "last-evaluated-key"); v != nil {
				req.LastEvaluatedKey = v
			}
			_, err := c.Resources.SearchResources(ctx, a.org, &req)
			return err
		})
	addBodyFlags(search)
	addPaginationFlags(search)
	search.Flags().String("search-query", "", "free-text search query")
	search.Flags().StringSlice("resource-types", nil, "resource types to search, e.g. WORKFLOW,STACK,WORKFLOW_GROUP (required by the API; comma-separated)")

	tags := a.command("tags", "List the distinct tags used on resources", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := resources.ListResourceTagsRequest{
				ContextTagKey: flagString(cmd, "context-tag-key"),
				Limit:         flagInt(cmd, "limit"),
				Query:         flagString(cmd, "query"),
				ResourceType:  flagString(cmd, "resource-type"),
				TagKind:       flagString(cmd, "tag-kind"),
			}
			_, err := c.Resources.ListResourceTags(ctx, a.org, &req)
			return err
		})
	tags.Flags().String("context-tag-key", "", "context tag key to list values for")
	tags.Flags().Int("limit", 0, "maximum number of tags to return")
	tags.Flags().String("query", "", "filter tags by prefix or substring")
	tags.Flags().String("resource-type", "", "restrict to a resource type")
	tags.Flags().String("tag-kind", "", "tag kind to list")

	cmd.AddCommand(move, search, tags)
	return cmd
}
