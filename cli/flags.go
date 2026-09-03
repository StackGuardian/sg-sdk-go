package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Optional query/body parameters are only sent when the flag was set
// explicitly, so a default value never overrides the API's own default.

func flagString(cmd *cobra.Command, name string) *string {
	if !cmd.Flags().Changed(name) {
		return nil
	}
	v, _ := cmd.Flags().GetString(name)
	return &v
}

func flagInt(cmd *cobra.Command, name string) *int {
	if !cmd.Flags().Changed(name) {
		return nil
	}
	v, _ := cmd.Flags().GetInt(name)
	return &v
}

func flagBool(cmd *cobra.Command, name string) *bool {
	if !cmd.Flags().Changed(name) {
		return nil
	}
	v, _ := cmd.Flags().GetBool(name)
	return &v
}

func flagStringSlice(cmd *cobra.Command, name string) []string {
	if !cmd.Flags().Changed(name) {
		return nil
	}
	v, _ := cmd.Flags().GetStringSlice(name)
	return v
}

func getString(cmd *cobra.Command, name string) string {
	v, _ := cmd.Flags().GetString(name)
	return v
}

func getBool(cmd *cobra.Command, name string) bool {
	v, _ := cmd.Flags().GetBool(name)
	return v
}

// requiredString registers a required string flag.
func requiredString(cmd *cobra.Command, name, usage string) {
	cmd.Flags().String(name, "", usage)
	_ = cmd.MarkFlagRequired(name)
}

// Path-scoping flags shared by the workflow-group / stack / workflow hierarchy.

func addWfGrpFlag(cmd *cobra.Command) {
	requiredString(cmd, "wfgrp", "workflow group name (nested groups use parent/child)")
}

func addStackFlag(cmd *cobra.Command) {
	requiredString(cmd, "stack", "stack name")
}

func addWfFlag(cmd *cobra.Command) {
	requiredString(cmd, "wf", "workflow name")
}

// addPaginationFlags registers the --limit / --last-evaluated-key pair used by list commands.
func addPaginationFlags(cmd *cobra.Command) {
	cmd.Flags().Int("limit", 0, "maximum number of results to return")
	cmd.Flags().String("last-evaluated-key", "", "pagination token from a previous response")
}

// checkEnum validates that value is one of allowed (case-sensitive).
func checkEnum(flag, value string, allowed ...string) error {
	for _, v := range allowed {
		if value == v {
			return nil
		}
	}
	return fmt.Errorf("invalid value %q for --%s: must be one of %s", value, flag, strings.Join(allowed, ", "))
}

// addWorkflowListFlags registers the filters shared by workflow list commands.
func addWorkflowListFlags(cmd *cobra.Command) {
	addPaginationFlags(cmd)
	cmd.Flags().String("description", "", "filter by description")
	cmd.Flags().String("iac-template-id", "", "filter by IaC template id")
	cmd.Flags().Int("latest-wf-run-statuses", 0, "filter by latest workflow run status")
	cmd.Flags().String("resource-names", "", "filter by resource names")
	cmd.Flags().String("runner-names", "", "filter by runner names")
	cmd.Flags().String("tags", "", "filter by tags")
}
