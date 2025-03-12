// This file was auto-generated by Fern from our API Definition.

package api

type GetBenchmarkReportsRequest struct {
	// Specifies the date for the report in `DD_MM_YYYY` format. Defaults to today's date if not provided.
	Date *string `json:"-" url:"date,omitempty"`
	// Boolean flag that, if set to `true`, provides detailed results instead of grouped data.
	Detailed *bool `json:"-" url:"detailed,omitempty"`
	// Filter by cloud service provider. Possible values: AZURE, GCP, AWS.
	// Support multiple values separated by comma
	FilterCsp *string `json:"-" url:"filter:CSP,omitempty"`
	// Filter by account ID.
	// Support multiple values separated by comma
	FilterAccountId *string `json:"-" url:"filter:accountId,omitempty"`
	// Filter by benchmark. Possible values: nist_sp_800_53_rev_5, soc_2, gdpr, cis_v200, etc.
	// Support multiple values separated by comma.
	FilterBenchmark *string `json:"-" url:"filter:benchmark,omitempty"`
	// Filter by control ID. There are over 800 possible control IDs.
	FilterControlId *string `json:"-" url:"filter:controlId,omitempty"`
	// Filter by control title.
	FilterControlTitle *string `json:"-" url:"filter:controlTitle,omitempty"`
	// Filter by region. Example values: eu-central-1, global.
	// Support multiple values separated by comma
	FilterRegion *string `json:"-" url:"filter:region,omitempty"`
	// Filter by severity level. Possible values: Critical, High, Medium, Low, Info, NA.
	// Support multiple values separated by comma.
	FilterSeverity *string `json:"-" url:"filter:severity,omitempty"`
	// Filter by status. Possible values: fails, passes, info, skips, error.
	// Support multiple values separated by comma.
	FilterStatus *string `json:"-" url:"filter:status,omitempty"`
	// Defines the column(s) by which the data should be grouped. Required when `detailed=false`.
	GroupBy *string `json:"-" url:"groupBy,omitempty"`
	// Limits the number of returned records. Required for pagination.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The page number for paginated results. Requires `limit` to be set.
	Page *int `json:"-" url:"page,omitempty"`
	// Specifies the column names to include in the detailed report. Can have multiple values separated by commas.
	// Available columns: CSP, severity, status, benchmark, integrations, resource, reason, region, controlTitle, controlId, observedAt.
	// Required when `detailed=true`.
	RequiredColumns *string `json:"-" url:"requiredColumns,omitempty"`
}
