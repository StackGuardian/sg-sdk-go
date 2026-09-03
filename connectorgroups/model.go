package connectorgroups

type DiscoveryScanRequest struct {
	// Comma-separated benchmark names to scan. Default is inventory.
	Benchmark *string `json:"-" url:"benchmark,omitempty"`
}

type DiscoveryScanResult struct {
	Triggered int `json:"triggered" url:"-"`
	Skipped   int `json:"skipped" url:"-"`
	Total     int `json:"total" url:"-"`
}

type DiscoveryScanResponse struct {
	Msg  *string              `json:"msg,omitempty" url:"-"`
	Data *DiscoveryScanResult `json:"data,omitempty" url:"-"`
}
