package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient handles HTTP requests with retry logic.
type HTTPClient struct {
	client       *http.Client
	baseURL      string
	apiKey       string
	userAgent    string
	maxRetries   int
	retryWaitMin time.Duration
	retryWaitMax time.Duration
}

// NewHTTPClient creates a new HTTP client with retry logic.
func NewHTTPClient(client *http.Client, baseURL, apiKey, userAgent string, maxRetries int, retryWaitMin, retryWaitMax time.Duration) *HTTPClient {
	if client == nil {
		client = &http.Client{Timeout: 60 * time.Second}
	}
	return &HTTPClient{
		client:       client,
		baseURL:      strings.TrimSuffix(baseURL, "/"),
		apiKey:       apiKey,
		userAgent:    userAgent,
		maxRetries:   maxRetries,
		retryWaitMin: retryWaitMin,
		retryWaitMax: retryWaitMax,
	}
}

// RequestOptions holds options for making an HTTP request.
type RequestOptions struct {
	Method      string
	Path        string
	QueryParams url.Values
	Body        interface{}
	Response    interface{}
}

// Do executes an HTTP request with retry logic.
func (c *HTTPClient) Do(ctx context.Context, opts *RequestOptions) error {
	var lastErr error

	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			// Calculate backoff duration
			waitDuration := c.calculateBackoff(attempt)

			select {
			case <-time.After(waitDuration):
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		err := c.doRequest(ctx, opts)
		if err == nil {
			return nil
		}

		lastErr = err

		// Don't retry on context cancellation or certain error types
		if ctx.Err() != nil {
			return err
		}

		// Check if error is retryable
		if !c.isRetryable(err) {
			return err
		}
	}

	return lastErr
}

// doRequest executes a single HTTP request.
func (c *HTTPClient) doRequest(ctx context.Context, opts *RequestOptions) error {
	// Build URL
	u, err := url.Parse(c.baseURL + opts.Path)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	if opts.QueryParams != nil && len(opts.QueryParams) > 0 {
		u.RawQuery = opts.QueryParams.Encode()
	}

	// Prepare request body
	var bodyReader io.Reader
	if opts.Body != nil {
		bodyBytes, err := json.Marshal(opts.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, opts.Method, u.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("User-Agent", c.userAgent)
	if opts.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	// Execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return &HTTPError{
			Message:    "request failed",
			StatusCode: 0,
			Err:        err,
			Retryable:  true,
		}
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			Message:    string(bodyBytes),
			StatusCode: resp.StatusCode,
			RequestID:  resp.Header.Get("X-Request-ID"),
			Retryable:  c.isStatusCodeRetryable(resp.StatusCode),
		}
	}

	// Decode response
	if opts.Response != nil {
		if err := json.NewDecoder(resp.Body).Decode(opts.Response); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// calculateBackoff calculates the backoff duration for a retry attempt.
func (c *HTTPClient) calculateBackoff(attempt int) time.Duration {
	// Exponential backoff with jitter
	backoff := float64(c.retryWaitMin) * math.Pow(2, float64(attempt-1))

	// Add jitter
	jitter := rand.Float64() * backoff * 0.1
	backoff += jitter

	// Cap at max wait time
	if backoff > float64(c.retryWaitMax) {
		backoff = float64(c.retryWaitMax)
	}

	return time.Duration(backoff)
}

// isRetryable determines if an error is retryable.
func (c *HTTPClient) isRetryable(err error) bool {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr.Retryable
	}
	return false
}

// isStatusCodeRetryable determines if a status code is retryable.
func (c *HTTPClient) isStatusCodeRetryable(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	default:
		return false
	}
}

// HTTPError represents an HTTP error.
type HTTPError struct {
	Message    string
	StatusCode int
	RequestID  string
	Retryable  bool
	Err        error
}

// Error implements the error interface.
func (e *HTTPError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
	}
	return e.Message
}

// Unwrap returns the underlying error.
func (e *HTTPError) Unwrap() error {
	return e.Err
}
