package cli

import (
	"context"

	"github.com/spf13/cobra"

	sgsdkgo "github.com/StackGuardian/sg-sdk-go"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) benchmarkReportsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "benchmark-reports",
		Short: "Read compliance benchmark reports",
	}

	get := a.command("get", "Get the benchmark report", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := sgsdkgo.GetBenchmarkReportsRequest{
				Date:               flagString(cmd, "date"),
				Detailed:           flagBool(cmd, "detailed"),
				FilterCsp:          flagString(cmd, "filter-csp"),
				FilterAccountId:    flagString(cmd, "filter-account-id"),
				FilterBenchmark:    flagString(cmd, "filter-benchmark"),
				FilterControlId:    flagString(cmd, "filter-control-id"),
				FilterControlTitle: flagString(cmd, "filter-control-title"),
				FilterRegion:       flagString(cmd, "filter-region"),
				FilterSeverity:     flagString(cmd, "filter-severity"),
				FilterStatus:       flagString(cmd, "filter-status"),
				GroupBy:            flagString(cmd, "group-by"),
				Limit:              flagInt(cmd, "limit"),
				Page:               flagInt(cmd, "page"),
				RequiredColumns:    flagString(cmd, "required-columns"),
			}
			_, err := c.BenchmarkReports.GetBenchmarkReports(ctx, a.org, &req)
			return err
		})
	get.Flags().String("date", "", "report date in DD_MM_YYYY format (defaults to today)")
	get.Flags().Bool("detailed", false, "return detailed results instead of grouped data")
	get.Flags().String("filter-csp", "", "filter by cloud provider (AWS, AZURE, GCP; comma-separated)")
	get.Flags().String("filter-account-id", "", "filter by account id (comma-separated)")
	get.Flags().String("filter-benchmark", "", "filter by benchmark, e.g. cis_v200 (comma-separated)")
	get.Flags().String("filter-control-id", "", "filter by control id")
	get.Flags().String("filter-control-title", "", "filter by control title")
	get.Flags().String("filter-region", "", "filter by region (comma-separated)")
	get.Flags().String("filter-severity", "", "filter by severity (Critical, High, Medium, Low, Info, NA)")
	get.Flags().String("filter-status", "", "filter by status (fails, passes, info, skips, error)")
	get.Flags().String("group-by", "", "column(s) to group by (required when not --detailed)")
	get.Flags().Int("limit", 0, "maximum number of records (required for pagination)")
	get.Flags().Int("page", 0, "page number (requires --limit)")
	get.Flags().String("required-columns", "", "columns to include in a detailed report (comma-separated)")

	cmd.AddCommand(get)
	return cmd
}
