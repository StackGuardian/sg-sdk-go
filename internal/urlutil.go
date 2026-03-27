package internal

import (
	"fmt"
	"net/url"
	"strings"
)

// BuildURL builds a URL with path parameters and optional nested resource support.
// For nested workflow groups (containing "/"), the slashes are preserved in the path.
func BuildURL(baseURL, pathTemplate string, params ...interface{}) (string, error) {
	// Replace %v placeholders with URL-encoded parameters
	parts := strings.Split(pathTemplate, "/")
	paramIndex := 0

	for i, part := range parts {
		if strings.Contains(part, "%v") {
			if paramIndex >= len(params) {
				return "", fmt.Errorf("not enough parameters for path template")
			}

			param := fmt.Sprintf("%v", params[paramIndex])

			// Check if this is a nested resource (contains "/")
			// This is important for workflow groups which can be nested like "parent/child"
			if strings.Contains(param, "/") {
				// For nested resources, don't encode the slashes
				parts[i] = strings.ReplaceAll(part, "%v", param)
			} else {
				// For regular parameters, URL encode them
				parts[i] = strings.ReplaceAll(part, "%v", url.PathEscape(param))
			}
			paramIndex++
		}
	}

	if paramIndex != len(params) {
		return "", fmt.Errorf("too many parameters for path template")
	}

	path := strings.Join(parts, "/")
	return baseURL + path, nil
}

// EncodeQueryParams encodes a map of query parameters into a url.Values.
func EncodeQueryParams(params map[string]interface{}) url.Values {
	values := url.Values{}
	for key, value := range params {
		if value != nil {
			values.Add(key, fmt.Sprintf("%v", value))
		}
	}
	return values
}
