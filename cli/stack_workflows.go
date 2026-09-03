package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/workflows"
)

func (a *app) stackWorkflowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stack-workflows",
		Short: "Manage the workflows that belong to a stack",
	}
	scope := func(cmd *cobra.Command) {
		addWfGrpFlag(cmd)
		addStackFlag(cmd)
	}

	create := a.command("create", "Create a workflow inside a stack", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req workflows.Workflow
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflows.CreateStackWorkflow(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(create)
	addBodyFlags(create)

	get := a.command("get WORKFLOW", "Get a stack workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflows.ReadStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"))
			return err
		})
	scope(get)

	del := a.command("delete WORKFLOW", "Delete a stack workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			return c.StackWorkflows.DeleteStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"))
		})
	scope(del)

	update := a.command("update WORKFLOW", "Update a stack workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req sgsdkgo.PatchedWorkflow
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflows.UpdateStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(update)
	addBodyFlags(update)

	outputs := a.command("outputs WORKFLOW", "Get the outputs of a stack workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflows.StackWorkflowOutputs(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"))
			return err
		})
	scope(outputs)

	uploadURL := a.command("tfstate-upload-url WORKFLOW", "Get a signed URL to upload a Terraform state file", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := sgsdkgo.GetSignedUrlToUploadTfstateFileForStackWorkflowRequest{Filename: flagString(cmd, "filename")}
			_, err := c.StackWorkflows.GetSignedUrlToUploadTfstateFileForStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(uploadURL)
	uploadURL.Flags().String("filename", "", "name of the file to upload (default tfstate.json)")

	fileUploadURL := a.command("file-upload-url WORKFLOW", "Get a signed URL to upload a file to a stack workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := workflows.GetFileUploadUrlRequest{Filename: getString(cmd, "filename"), Folder: flagString(cmd, "folder")}
			_, err := c.StackWorkflows.GetFileUploadUrlForStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(fileUploadURL)
	requiredString(fileUploadURL, "filename", "name of the file to upload")
	fileUploadURL.Flags().String("folder", "", "folder to upload the file into")

	compare := a.command("compare WORKFLOW", "Compare a stack workflow against a template without applying changes", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflows.CompareWorkflowRequest
			if hasBody(cmd) {
				if err := readBody(cmd, &req); err != nil {
					return err
				}
			} else {
				req.TargetTemplateId = getString(cmd, "target-template-id")
				req.UpgradeMode = flagString(cmd, "upgrade-mode")
			}
			_, err := c.StackWorkflows.CompareStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(compare)
	addBodyFlags(compare)
	compare.Flags().String("target-template-id", "", "template to compare against")
	compare.Flags().String("upgrade-mode", "", "upgrade mode to simulate")
	compare.MarkFlagsOneRequired("target-template-id", "body", "body-file")

	vcsTriggers := a.command("vcs-triggers WORKFLOW", "Create VCS webhook triggers for a stack workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflows.CreateVcsTriggersRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.StackWorkflows.CreateVcsTriggersForStackWorkflow(ctx, a.org, getString(cmd, "stack"), args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(vcsTriggers)
	addBodyFlags(vcsTriggers)

	list := a.command("list", "List the workflows of a stack", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllStackWorkflowsRequest{
				Description:         flagString(cmd, "description"),
				IacTemplateId:       flagString(cmd, "iac-template-id"),
				LatestWfRunStatuses: flagInt(cmd, "latest-wf-run-statuses"),
				ResourceNames:       flagString(cmd, "resource-names"),
				RunnerNames:         flagString(cmd, "runner-names"),
				Tags:                flagString(cmd, "tags"),
				Lastevaluatedkey:    flagString(cmd, "last-evaluated-key"),
				Limit:               flagInt(cmd, "limit"),
			}
			_, err := c.StackWorkflows.ListAllStackWorkflows(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(list)
	addWorkflowListFlags(list)

	cmd.AddCommand(create, get, del, update, outputs, uploadURL, fileUploadURL, compare, vcsTriggers, list, a.stackWorkflowArtifactsCmd())
	return cmd
}

func (a *app) stackWorkflowArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "artifacts", Short: "Manage the artifacts of a stack workflow"}
	scope := func(cmd *cobra.Command) {
		addWfGrpFlag(cmd)
		addStackFlag(cmd)
		addWfFlag(cmd)
	}

	list := a.command("list", "List the artifacts of a stack workflow", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.StackWorkflows.ListAllStackWorkflowsArtifacts(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"))
			return err
		})
	scope(list)

	getURL := a.command("get-url", "Get a signed download URL for an artifact", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := workflows.GetWorkflowArtifactRequest{ArtifactPath: getString(cmd, "artifact-path")}
			_, err := c.StackWorkflows.GetStackWorkflowArtifactUrl(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(getURL)
	requiredString(getURL, "artifact-path", "path of the artifact")

	get := a.command("get ARTIFACT_ID", "Read an artifact's content (or a signed URL on private runners)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflows.ReadStackWorkflowArtifact(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(get)

	create := a.command("create ARTIFACT_ID", "Create or replace an artifact with the JSON body as its content", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var content interface{}
			if err := readBody(cmd, &content); err != nil {
				return err
			}
			_, err := c.StackWorkflows.CreateStackWorkflowArtifact(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], content)
			return err
		})
	scope(create)
	addBodyFlags(create)

	del := a.command("delete ARTIFACT_ID", "Delete an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflows.DeleteStackWorkflowArtifact(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(del)

	lock := a.command("lock ARTIFACT_ID", "Lock the stack workflow on an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflows.LockStackWorkflowArtifact(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(lock)

	unlock := a.command("unlock ARTIFACT_ID", "Release the stack workflow's lock on an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.StackWorkflows.UnlockStackWorkflowArtifact(ctx, a.org, getString(cmd, "stack"), getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(unlock)

	cmd.AddCommand(list, getURL, get, create, del, lock, unlock)
	return cmd
}
