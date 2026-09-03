package cli

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// capture is an http.RoundTripper that records the most recent response body
// so the CLI can print the API's JSON verbatim instead of a lossy re-encoding
// of the SDK's typed structs (which drop extra properties on marshal).
type capture struct {
	next    http.RoundTripper
	verbose bool
	log     io.Writer

	mu   sync.Mutex
	body []byte
}

func (c *capture) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.verbose {
		fmt.Fprintf(c.log, "> %s %s\n", req.Method, req.URL)
	}
	resp, err := c.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	c.mu.Lock()
	c.body = body
	c.mu.Unlock()
	if c.verbose {
		fmt.Fprintf(c.log, "< %d (%d bytes)\n", resp.StatusCode, len(body))
	}
	return resp, nil
}

func (c *capture) lastBody() []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.body
}
