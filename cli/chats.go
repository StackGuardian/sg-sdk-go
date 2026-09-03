package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/chats"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) chatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chats",
		Short: "Manage AI assistant chats (infra2code and kyro) and their messages",
	}

	create := a.command("create", "Create a chat", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req chats.Chat
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Chats.CreateChat(ctx, a.org, &req)
			return err
		})
	addBodyFlags(create)

	list := a.command("list", "List chats", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := chats.ListAllChatsRequest{
				Limit:            flagInt(cmd, "limit"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				SearchQuery:      flagString(cmd, "search-query"),
				RefreshPrStatus:  flagBool(cmd, "refresh-pr-status"),
				ChatType:         flagString(cmd, "chat-type"),
			}
			_, err := c.Chats.ListAllChats(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("search-query", "", "comma-separated terms matched against name, description and tags")
	list.Flags().Bool("refresh-pr-status", false, "refresh pull request statuses from the VCS provider")
	list.Flags().String("chat-type", "", "filter by chat type (infra2code, kyro)")

	get := a.command("get CHAT", "Get a chat", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Chats.ReadChat(ctx, a.org, args[0], &chats.ReadChatRequest{RefreshPrStatus: flagBool(cmd, "refresh-pr-status")})
			return err
		})
	get.Flags().Bool("refresh-pr-status", false, "refresh the pull request status from the VCS provider")

	update := a.command("update CHAT", "Update a chat", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req chats.PatchedChat
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Chats.UpdateChat(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(update)

	del := a.command("delete CHAT", "Archive a chat", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Chats.DeleteChat(ctx, a.org, args[0])
			return err
		})

	commit := a.command("commit-bundle CHAT", "Commit an uploaded git bundle", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			v, _ := cmd.Flags().GetInt("expected-version")
			_, err := c.Chats.CommitChatBundle(ctx, a.org, args[0], &chats.CommitChatBundleRequest{ExpectedVersion: v})
			return err
		})
	commit.Flags().Int("expected-version", 0, "the BundleVersion last seen by the caller")
	_ = commit.MarkFlagRequired("expected-version")

	createPr := a.command("create-pr CHAT", "Push the generated code and open a pull request", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req chats.CreateChatPrRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Chats.CreateChatPr(ctx, a.org, args[0], &req)
			return err
		})
	addBodyFlags(createPr)

	prStatus := a.command("pr-status CHAT", "Show the status of the chat's pull request", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Chats.GetChatPrStatus(ctx, a.org, args[0])
			return err
		})

	prSync := a.command("pr-sync-status CHAT", "Compare the local bundle with the pull request", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Chats.GetChatPrSyncStatus(ctx, a.org, args[0])
			return err
		})

	cmd.AddCommand(create, list, get, update, del, commit, createPr, prStatus, prSync, a.chatArtifactsCmd(), a.chatMessagesCmd())
	return cmd
}

func (a *app) chatArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "artifacts", Short: "Work with a chat's artifacts through presigned URLs"}

	list := a.command("list CHAT", "Get a presigned URL that lists the artifacts", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := chats.ListAllChatArtifactsRequest{
				Prefix:        flagString(cmd, "prefix"),
				MaxKeys:       flagInt(cmd, "max-keys"),
				StartAfterKey: flagString(cmd, "start-after-key"),
			}
			_, err := c.Chats.ListAllChatArtifacts(ctx, a.org, args[0], &req)
			return err
		})
	list.Flags().String("prefix", "", "only list artifacts under this prefix")
	list.Flags().Int("max-keys", 0, "maximum number of keys")
	list.Flags().String("start-after-key", "", "full S3 key to start after")

	getURL := a.command("get-url CHAT", "Get a presigned download URL for an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Chats.GetChatArtifact(ctx, a.org, args[0], &chats.GetChatArtifactRequest{ArtifactPath: getString(cmd, "artifact-path")})
			return err
		})
	requiredString(getURL, "artifact-path", "artifact path relative to the chat's artifact store")

	uploadURL := a.command("upload-url CHAT", "Get a presigned upload URL for an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := chats.SaveChatArtifactRequest{ArtifactPath: getString(cmd, "artifact-path"), ContentType: flagString(cmd, "content-type")}
			_, err := c.Chats.SaveChatArtifact(ctx, a.org, args[0], &req)
			return err
		})
	requiredString(uploadURL, "artifact-path", "artifact path relative to the chat's artifact store")
	uploadURL.Flags().String("content-type", "", "content type the upload must use (default application/octet-stream)")

	cmd.AddCommand(list, getURL, uploadURL)
	return cmd
}

func (a *app) chatMessagesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "messages", Short: "Send and read the messages of a kyro chat"}
	scope := func(cmd *cobra.Command) { requiredString(cmd, "chat", "chat id") }

	create := a.command("create", "Send a message", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req chats.Message
			if hasBody(cmd) {
				if err := readBody(cmd, &req); err != nil {
					return err
				}
			} else {
				text := getString(cmd, "text")
				req.Content = []*chats.ContentBlock{{Type: "text", Text: &text}}
			}
			req.ClientRequestId = flagString(cmd, "client-request-id")
			_, err := c.Chats.CreateMessage(ctx, a.org, getString(cmd, "chat"), &req)
			return err
		})
	scope(create)
	addBodyFlags(create)
	create.Flags().String("text", "", "text of the message (shortcut for a single text content block)")
	create.Flags().String("client-request-id", "", "idempotency key")
	create.MarkFlagsOneRequired("text", "body", "body-file")
	create.MarkFlagsMutuallyExclusive("text", "body")
	create.MarkFlagsMutuallyExclusive("text", "body-file")

	list := a.command("list", "List the messages of a chat, oldest first", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := chats.ListAllMessagesRequest{Limit: flagInt(cmd, "limit"), Lastevaluatedkey: flagString(cmd, "last-evaluated-key")}
			_, err := c.Chats.ListAllMessages(ctx, a.org, getString(cmd, "chat"), &req)
			return err
		})
	scope(list)
	addPaginationFlags(list)

	get := a.command("get MESSAGE", "Get a message", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Chats.ReadMessage(ctx, a.org, getString(cmd, "chat"), args[0])
			return err
		})
	scope(get)

	cancel := a.command("cancel MESSAGE", "Stop a pending or streaming assistant message", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Chats.CancelMessage(ctx, a.org, getString(cmd, "chat"), args[0])
			return err
		})
	scope(cancel)

	retry := a.command("retry MESSAGE", "Retry an errored or stopped assistant message", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Chats.RetryMessage(ctx, a.org, getString(cmd, "chat"), args[0])
			return err
		})
	scope(retry)

	cmd.AddCommand(create, list, get, cancel, retry)
	return cmd
}
