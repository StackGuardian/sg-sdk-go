package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
	"github.com/StackGuardian/sg-sdk-go/workflows"
)

func (a *app) workflowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workflows",
		Aliases: []string{"wf", "wfs"},
		Short:   "Manage workflows inside a workflow group",
	}

	create := a.command("create", "Create a workflow", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req workflows.Workflow
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Workflows.CreateWorkflow(ctx, a.org, getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(create)
	addBodyFlags(create)

	get := a.command("get WORKFLOW", "Get a workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.ReadWorkflow(ctx, a.org, args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(get)

	del := a.command("delete WORKFLOW", "Delete a workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.DeleteWorkflow(ctx, a.org, args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(del)

	update := a.command("update WORKFLOW", "Update a workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflows.PatchedWorkflow
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			var mode *workflows.UpgradeModeEnum
			if v := flagString(cmd, "upgrade-mode"); v != nil {
				if err := checkEnum("upgrade-mode", *v, "PRESERVE_SETTINGS", "RESET_TO_TEMPLATE", "MANUAL"); err != nil {
					return err
				}
				m := workflows.UpgradeModeEnum(*v)
				mode = &m
			}
			_, err := c.Workflows.UpdateWorkflow(ctx, a.org, args[0], getString(cmd, "wfgrp"), mode, &req)
			return err
		})
	addWfGrpFlag(update)
	addBodyFlags(update)
	update.Flags().String("upgrade-mode", "", "template upgrade mode (PRESERVE_SETTINGS, RESET_TO_TEMPLATE, MANUAL)")

	outputs := a.command("outputs WORKFLOW", "Get the outputs of a workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.Outputs(ctx, a.org, args[0], getString(cmd, "wfgrp"))
			return err
		})
	addWfGrpFlag(outputs)

	uploadURL := a.command("tfstate-upload-url WORKFLOW", "Get a signed URL to upload a Terraform state file", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := sgsdkgo.GetSignedUrlToUploadTfstateFileRequest{Filename: flagString(cmd, "filename")}
			_, err := c.Workflows.GetSignedUrlToUploadTfstateFile(ctx, a.org, args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(uploadURL)
	uploadURL.Flags().String("filename", "", "name of the file to upload (default tfstate.json)")

	fileUploadURL := a.command("file-upload-url WORKFLOW", "Get a signed URL to upload a file to a workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := workflows.GetFileUploadUrlRequest{Filename: getString(cmd, "filename"), Folder: flagString(cmd, "folder")}
			_, err := c.Workflows.GetFileUploadUrl(ctx, a.org, args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(fileUploadURL)
	requiredString(fileUploadURL, "filename", "name of the file to upload")
	fileUploadURL.Flags().String("folder", "", "folder to upload the file into")

	compare := a.command("compare WORKFLOW", "Compare a workflow against a template without applying changes", cobra.ExactArgs(1),
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
			_, err := c.Workflows.CompareWorkflow(ctx, a.org, args[0], getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(compare)
	addBodyFlags(compare)
	compare.Flags().String("target-template-id", "", "template to compare against")
	compare.Flags().String("upgrade-mode", "", "upgrade mode to simulate")
	compare.MarkFlagsOneRequired("target-template-id", "body", "body-file")

	vcsTriggers := a.command("vcs-triggers WORKFLOW", "Create VCS webhook triggers for a workflow", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var req workflows.CreateVcsTriggersRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Workflows.CreateVcsTriggers(ctx, a.org, getString(cmd, "wfgrp"), args[0], &req)
			return err
		})
	addWfGrpFlag(vcsTriggers)
	addBodyFlags(vcsTriggers)

	list := a.command("list", "List the workflows in a workflow group", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ListAllWorkflowsRequest{
				Description:         flagString(cmd, "description"),
				IacTemplateId:       flagString(cmd, "iac-template-id"),
				LatestWfRunStatuses: flagInt(cmd, "latest-wf-run-statuses"),
				ResourceNames:       flagString(cmd, "resource-names"),
				RunnerNames:         flagString(cmd, "runner-names"),
				Tags:                flagString(cmd, "tags"),
				Lastevaluatedkey:    flagString(cmd, "last-evaluated-key"),
				Limit:               flagInt(cmd, "limit"),
			}
			_, err := c.Workflows.ListAllWorkflows(ctx, a.org, getString(cmd, "wfgrp"), &req)
			return err
		})
	addWfGrpFlag(list)
	addWorkflowListFlags(list)

	cmd.AddCommand(create, get, del, update, outputs, uploadURL, fileUploadURL, compare, vcsTriggers, list, a.workflowArtifactsCmd())
	return cmd
}

func (a *app) workflowArtifactsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "artifacts", Short: "Manage the artifacts of a workflow (state files, outputs, ...)"}
	scope := func(cmd *cobra.Command) {
		addWfGrpFlag(cmd)
		addWfFlag(cmd)
	}

	list := a.command("list", "List the artifacts of a workflow", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			_, err := c.Workflows.ListAllWorkflowArtifacts(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"))
			return err
		})
	scope(list)

	getURL := a.command("get-url", "Get a signed download URL for an artifact", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := workflows.GetWorkflowArtifactRequest{ArtifactPath: getString(cmd, "artifact-path"), VersionId: flagString(cmd, "version-id")}
			_, err := c.Workflows.GetWorkflowArtifactUrl(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), &req)
			return err
		})
	scope(getURL)
	requiredString(getURL, "artifact-path", "path of the artifact")
	getURL.Flags().String("version-id", "", "specific version of the artifact")

	get := a.command("get ARTIFACT_ID", "Read an artifact's content (or a signed URL on private runners)", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.ReadArtifact(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(get)

	create := a.command("create ARTIFACT_ID", "Create or replace an artifact with the JSON body as its content", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			var content interface{}
			if err := readBody(cmd, &content); err != nil {
				return err
			}
			_, err := c.Workflows.CreateArtifact(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], content)
			return err
		})
	scope(create)
	addBodyFlags(create)

	del := a.command("delete ARTIFACT_ID", "Delete an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.DeleteArtifact(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(del)

	versions := a.command("versions ARTIFACT_ID", "List the versions of an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			req := workflows.ListArtifactVersionsRequest{Lastevaluatedkey: flagString(cmd, "last-evaluated-key"), Limit: flagInt(cmd, "limit")}
			_, err := c.Workflows.ListArtifactVersions(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], &req)
			return err
		})
	scope(versions)
	addPaginationFlags(versions)

	rollback := a.command("rollback ARTIFACT_ID", "Roll an artifact back to a previous version", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.RollbackArtifactVersion(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0], &workflows.RollbackArtifactVersionRequest{VersionId: getString(cmd, "version-id")})
			return err
		})
	scope(rollback)
	requiredString(rollback, "version-id", "version to roll back to")

	lock := a.command("lock ARTIFACT_ID", "Lock the workflow on an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.LockArtifact(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(lock)

	unlock := a.command("unlock ARTIFACT_ID", "Release the workflow's lock on an artifact", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, args []string) error {
			_, err := c.Workflows.UnlockArtifact(ctx, a.org, getString(cmd, "wf"), getString(cmd, "wfgrp"), args[0])
			return err
		})
	scope(unlock)

	cmd.AddCommand(list, getURL, get, create, del, versions, rollback, lock, unlock)
	return cmd
}
