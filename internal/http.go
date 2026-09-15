package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient is an interface for a subset of the *http.Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ResolveBaseURL resolves the base URL from the given arguments,
// preferring the first non-empty value.
func ResolveBaseURL(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

// pathWithSlashes is a path parameter whose "/" are kept as separators
// (e.g. nested workflow groups); each segment is still escaped.
// DO NOT REVERT - the API cannot resolve nested workflow groups sent as %2F.
//
// Unexported on purpose: the only way to obtain one is NewPathWithSlashes,
// which validates. A bare conversion from outside this package isn't
// possible, so validation can't be bypassed by accident.
type pathWithSlashes string

// NewPathWithSlashes validates value and returns it as a pathWithSlashes,
// or an error if invalid. This is the only way to construct one.
func NewPathWithSlashes(value string) (pathWithSlashes, error) {
	p := pathWithSlashes(value)
	if err := p.Validate(); err != nil {
		return "", err
	}
	return p, nil
}

// Validate returns an error if p contains an empty segment - an empty
// value, or a leading, trailing, or doubled "/" - which would otherwise
// silently escape to a malformed request URL instead of a clear failure.
// The error doesn't name the field being validated; callers know which
// parameter they're validating and should add that context themselves.
func (p pathWithSlashes) Validate() error {
	for _, s := range strings.Split(string(p), "/") {
		if s == "" {
			return fmt.Errorf("invalid value %q: can't be empty or contain a leading, trailing, or doubled \"/\"", string(p))
		}
	}
	return nil
}

// EncodeURL encodes the given arguments into the URL, escaping
// values as needed.
func EncodeURL(urlFormat string, args ...interface{}) string {
	escapedArgs := make([]interface{}, 0, len(args))
	for _, arg := range args {
		if p, ok := arg.(pathWithSlashes); ok {
			escapedArgs = append(escapedArgs, escapePathWithSlashes(string(p)))
			continue
		}
		escapedArgs = append(escapedArgs, url.PathEscape(fmt.Sprintf("%v", arg)))
	}
	return fmt.Sprintf(urlFormat, escapedArgs...)
}

func escapePathWithSlashes(path string) string {
	segments := strings.Split(path, "/")
	for i, s := range segments {
		segments[i] = url.PathEscape(s)
	}
	return strings.Join(segments, "/")
}

// MergeHeaders merges the given headers together, where the right
// takes precedence over the left.
func MergeHeaders(left, right http.Header) http.Header {
	for key, values := range right {
		if len(values) > 1 {
			left[key] = values
			continue
		}
		if value := right.Get(key); value != "" {
			left.Set(key, value)
		}
	}
	return left
}
