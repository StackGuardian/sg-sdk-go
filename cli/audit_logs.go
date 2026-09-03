package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) auditLogsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "audit-logs",
		Short: "Read organization audit logs",
	}

	list := a.command("list", "List audit log entries", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.ReadAuditLogsRequest{
				Effect:           flagString(cmd, "effect"),
				EndTime:          flagInt(cmd, "end-time"),
				Lastevaluatedkey: flagString(cmd, "last-evaluated-key"),
				Limit:            flagInt(cmd, "limit"),
				PrincipalEmail:   flagString(cmd, "principal-email"),
				RequestMethod:    flagString(cmd, "request-method"),
				Resource:         flagString(cmd, "resource"),
				SourceIp:         flagString(cmd, "source-ip"),
				StartTime:        flagInt(cmd, "start-time"),
			}
			_, err := c.AccessManagement.ReadAuditLogs(ctx, a.org, &req)
			return err
		})
	addPaginationFlags(list)
	list.Flags().String("effect", "", "filter by effect (Allow, Deny)")
	list.Flags().Int("start-time", 0, "start time as a Unix timestamp in milliseconds")
	list.Flags().Int("end-time", 0, "end time as a Unix timestamp in milliseconds")
	list.Flags().String("principal-email", "", "filter by principal email")
	list.Flags().String("request-method", "", "filter by HTTP method (GET, POST, ...)")
	list.Flags().String("resource", "", "filter by resource name, e.g. SG_SIGN_IN or /orgs/<org>/wfgrps/<grp>/wfs/<wf>")
	list.Flags().String("source-ip", "", "filter by source IP address")

	cmd.AddCommand(list)
	return cmd
}
