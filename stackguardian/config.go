package stackguardian

import (
	"net/http"
	"time"

	api "github.com/StackGuardian/sg-sdk-go"
)

// Config holds the configuration for the StackGuardian client.
type Config struct {
	// APIKey is the API key for authentication with StackGuardian.
	// Format: "apikey YOUR_API_KEY"
	APIKey string

	// BaseURL is the base URL for the StackGuardian API.
	// Default: https://api.app.stackguardian.io
	BaseURL string

	// HTTPClient is the HTTP client to use for requests.
	// If nil, a default client with sensible timeouts will be used.
	HTTPClient *http.Client

	// MaxRetries is the maximum number of retry attempts for failed requests.
	// Default: 3
	MaxRetries int

	// RetryWaitMin is the minimum time to wait between retries.
	// Default: 1 second
	RetryWaitMin time.Duration

	// RetryWaitMax is the maximum time to wait between retries.
	// Default: 30 seconds
	RetryWaitMax time.Duration

	// UserAgent is the User-Agent header to send with requests.
	// Default: "sg-sdk-go/v{version}"
	UserAgent string
}

// DefaultConfig returns a Config with default values.
func DefaultConfig() *Config {
	return &Config{
		BaseURL:      "https://api.app.stackguardian.io",
		HTTPClient:   defaultHTTPClient(),
		MaxRetries:   3,
		RetryWaitMin: 1 * time.Second,
		RetryWaitMax: 30 * time.Second,
		UserAgent:    "sg-sdk-go/v" + Version,
	}
}

// defaultHTTPClient returns a default HTTP client with sensible timeouts.
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.APIKey == "" {
		return &api.Error{
			Type:    api.ErrorTypeConfiguration,
			Message: "APIKey is required",
		}
	}
	if c.BaseURL == "" {
		return &api.Error{
			Type:    api.ErrorTypeConfiguration,
			Message: "BaseURL is required",
		}
	}
	return nil
}
