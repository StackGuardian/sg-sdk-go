package cli

import (
	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
)

// addApprovalFlags registers the flags shared by the workflow-run approve commands.
func addApprovalFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("reject", false, "reject instead of approve")
	cmd.Flags().String("message", "", "message to attach to the decision")
	cmd.Flags().String("approval-step", "", "name of the approval step")
	cmd.Flags().String("reason", "", "reason the approval was required")
}

func approvalFromFlags(cmd *cobra.Command) *sgsdkgo.WorkflowRunApproval {
	approve := !getBool(cmd, "reject")
	return &sgsdkgo.WorkflowRunApproval{
		Approve:                   &approve,
		Message:                   flagString(cmd, "message"),
		ApprovalStep:              flagString(cmd, "approval-step"),
		ReasonForApprovalRequired: flagString(cmd, "reason"),
	}
}
