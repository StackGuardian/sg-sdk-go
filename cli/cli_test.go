package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recorded struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   string
}

// fakeAPI records every request and answers with a canned status and body.
type fakeAPI struct {
	srv      *httptest.Server
	mu       sync.Mutex
	requests []recorded
	status   int
	response string
}

func newFakeAPI(t *testing.T) *fakeAPI {
	t.Helper()
	f := &fakeAPI{status: http.StatusOK, response: "{}"}
	f.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		f.mu.Lock()
		f.requests = append(f.requests, recorded{r.Method, r.URL.Path, r.URL.Query(), r.Header.Clone(), string(body)})
		status, response := f.status, f.response
		f.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, response)
	}))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeAPI) last(t *testing.T) recorded {
	t.Helper()
	f.mu.Lock()
	defer f.mu.Unlock()
	require.NotEmpty(t, f.requests, "expected at least one HTTP request")
	return f.requests[len(f.requests)-1]
}

// execCLI runs the CLI against f with test credentials and returns stdout, stderr and the error.
func execCLI(t *testing.T, f *fakeAPI, stdin string, args ...string) (string, string, error) {
	t.Helper()
	root := NewRootCmd()
	var out, errOut bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&errOut)
	root.SetIn(strings.NewReader(stdin))
	root.SetArgs(append([]string{"--base-url", f.srv.URL, "--api-key", "test-token", "--org", "demo-org"}, args...))
	err := root.Execute()
	return out.String(), errOut.String(), err
}

// cmdCase describes one CLI invocation and the HTTP request it must produce.
type cmdCase struct {
	name     string
	args     []string
	stdin    string
	method   string
	path     string
	query    url.Values        // nil asserts that no query string is sent
	body     string            // "" asserts that no body is sent
	headers  map[string]string // extra headers that must be present
	response string            // server response, defaults to {}
}

func runCases(t *testing.T, cases []cmdCase) {
	t.Helper()
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			f := newFakeAPI(t)
			if tc.response != "" {
				f.response = tc.response
			}
			out, _, err := execCLI(t, f, tc.stdin, tc.args...)
			require.NoError(t, err)
			r := f.last(t)
			assert.Equal(t, tc.method, r.Method, "method")
			assert.Equal(t, tc.path, r.Path, "path")
			if tc.query == nil {
				assert.Empty(t, r.Query, "query")
			} else {
				assert.Equal(t, tc.query, r.Query, "query")
			}
			if tc.body == "" {
				assert.Empty(t, r.Body, "body")
			} else {
				assert.JSONEq(t, tc.body, r.Body, "body")
			}
			assert.Equal(t, "apikey test-token", r.Header.Get("Authorization"))
			for k, v := range tc.headers {
				assert.Equal(t, v, r.Header.Get(k), "header "+k)
			}
			assert.Equal(t, indentJSON(t, f.response), out, "stdout")
		})
	}
}

func indentJSON(t *testing.T, s string) string {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, json.Indent(&buf, []byte(s), "", "  "))
	return buf.String() + "\n"
}

func TestGlobalConfig(t *testing.T) {
	t.Run("missing api key", func(t *testing.T) {
		f := newFakeAPI(t)
		root := NewRootCmd()
		root.SetArgs([]string{"--base-url", f.srv.URL, "--org", "o", "organizations", "get"})
		t.Setenv(envAPIToken, "")
		err := root.Execute()
		require.Error(t, err)
		assert.Contains(t, err.Error(), envAPIToken)
	})
	t.Run("missing org", func(t *testing.T) {
		f := newFakeAPI(t)
		root := NewRootCmd()
		root.SetArgs([]string{"--base-url", f.srv.URL, "--api-key", "k", "organizations", "get"})
		t.Setenv(envOrg, "")
		err := root.Execute()
		require.Error(t, err)
		assert.Contains(t, err.Error(), envOrg)
	})
	t.Run("env vars and apikey prefix", func(t *testing.T) {
		f := newFakeAPI(t)
		t.Setenv(envAPIToken, "apikey from-env")
		t.Setenv(envBaseURL, f.srv.URL+"/")
		t.Setenv(envOrg, "env-org")
		root := NewRootCmd()
		var out bytes.Buffer
		root.SetOut(&out)
		root.SetArgs([]string{"organizations", "get"})
		require.NoError(t, root.Execute())
		r := f.last(t)
		assert.Equal(t, "/api/v1/orgs/env-org/", r.Path)
		assert.Equal(t, "apikey from-env", r.Header.Get("Authorization"))
	})
	t.Run("compact output preserves extra fields", func(t *testing.T) {
		f := newFakeAPI(t)
		f.response = `{"msg": {"ResourceName": "demo-org", "Unknown": [1, 2]}}`
		out, _, err := execCLI(t, f, "", "--compact", "organizations", "get")
		require.NoError(t, err)
		assert.Equal(t, `{"msg":{"ResourceName":"demo-org","Unknown":[1,2]}}`+"\n", out)
	})
	t.Run("verbose logs request", func(t *testing.T) {
		f := newFakeAPI(t)
		_, errOut, err := execCLI(t, f, "", "--verbose", "organizations", "get")
		require.NoError(t, err)
		assert.Contains(t, errOut, "> GET "+f.srv.URL+"/api/v1/orgs/demo-org/")
		assert.Contains(t, errOut, "< 200")
	})
	t.Run("api error surfaces status", func(t *testing.T) {
		f := newFakeAPI(t)
		f.status = http.StatusNotFound
		f.response = `{"message": "no such org"}`
		_, _, err := execCLI(t, f, "", "organizations", "get")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "404")
	})
	t.Run("missing body", func(t *testing.T) {
		f := newFakeAPI(t)
		_, _, err := execCLI(t, f, "", "secrets", "create")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "--body")
	})
	t.Run("invalid body", func(t *testing.T) {
		f := newFakeAPI(t)
		_, _, err := execCLI(t, f, "", "secrets", "create", "--body", "{not json")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid JSON")
	})
	t.Run("unknown body field", func(t *testing.T) {
		f := newFakeAPI(t)
		_, _, err := execCLI(t, f, "", "stacks", "create", "--wfgrp", "g", "-b", `{"ResourceName":"s","Descriptionn":"typo","RunOnCreate":true}`)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown field(s) in request body: Descriptionn, RunOnCreate")
		f.mu.Lock()
		defer f.mu.Unlock()
		assert.Empty(t, f.requests, "no request should be sent")
	})
	t.Run("body field case insensitive", func(t *testing.T) {
		f := newFakeAPI(t)
		_, _, err := execCLI(t, f, "", "secrets", "create", "-b", `{"resourceName":"s","resourcevalue":"v"}`)
		require.NoError(t, err)
		assert.JSONEq(t, `{"ResourceName":"s","ResourceValue":"v"}`, f.last(t).Body)
	})
	t.Run("body from file", func(t *testing.T) {
		f := newFakeAPI(t)
		path := filepath.Join(t.TempDir(), "body.json")
		require.NoError(t, os.WriteFile(path, []byte(`{"ResourceName":"s1","ResourceValue":"v"}`), 0o600))
		_, _, err := execCLI(t, f, "", "secrets", "create", "--body-file", path)
		require.NoError(t, err)
		assert.JSONEq(t, `{"ResourceName":"s1","ResourceValue":"v"}`, f.last(t).Body)
	})
	t.Run("max attempts retries", func(t *testing.T) {
		f := newFakeAPI(t)
		f.status = http.StatusInternalServerError
		_, _, err := execCLI(t, f, "", "--max-attempts", "2", "organizations", "get")
		require.Error(t, err)
		f.mu.Lock()
		defer f.mu.Unlock()
		assert.Len(t, f.requests, 2)
	})
	t.Run("explicit null is preserved for optional fields", func(t *testing.T) {
		f := newFakeAPI(t)
		_, _, err := execCLI(t, f, "", "secrets", "update", "s", "-b", `{"ResourceValue":null, "resourcename": null}`)
		require.NoError(t, err)
		assert.JSONEq(t, `{"ResourceName":null,"ResourceValue":null}`, f.last(t).Body)
	})
	t.Run("body from stdin", func(t *testing.T) {
		f := newFakeAPI(t)
		_, _, err := execCLI(t, f, `{"ResourceName":"s1"}`, "secrets", "create", "-f", "-")
		require.NoError(t, err)
		assert.JSONEq(t, `{"ResourceName":"s1","ResourceValue":""}`, f.last(t).Body)
	})
}
