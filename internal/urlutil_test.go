package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		template string
		params   []interface{}
		want     string
		wantErr  bool
	}{
		{
			name:     "simple path with one parameter",
			baseURL:  "",
			template: "/api/v1/orgs/%v/",
			params:   []interface{}{"my-org"},
			want:     "/api/v1/orgs/my-org/",
			wantErr:  false,
		},
		{
			name:     "path with multiple parameters",
			baseURL:  "",
			template: "/api/v1/orgs/%v/wfgrps/%v/",
			params:   []interface{}{"my-org", "my-group"},
			want:     "/api/v1/orgs/my-org/wfgrps/my-group/",
			wantErr:  false,
		},
		{
			name:     "nested workflow group (preserves slashes)",
			baseURL:  "",
			template: "/api/v1/orgs/%v/wfgrps/%v/",
			params:   []interface{}{"my-org", "parent/child"},
			want:     "/api/v1/orgs/my-org/wfgrps/parent/child/",
			wantErr:  false,
		},
		{
			name:     "deeply nested workflow group",
			baseURL:  "",
			template: "/api/v1/orgs/%v/wfgrps/%v/",
			params:   []interface{}{"my-org", "parent/child/grandchild"},
			want:     "/api/v1/orgs/my-org/wfgrps/parent/child/grandchild/",
			wantErr:  false,
		},
		{
			name:     "parameter with special characters (URL encoded)",
			baseURL:  "",
			template: "/api/v1/orgs/%v/",
			params:   []interface{}{"org name with spaces"},
			want:     "/api/v1/orgs/org%20name%20with%20spaces/",
			wantErr:  false,
		},
		{
			name:     "with base URL",
			baseURL:  "https://api.app.stackguardian.io",
			template: "/api/v1/orgs/%v/",
			params:   []interface{}{"my-org"},
			want:     "https://api.app.stackguardian.io/api/v1/orgs/my-org/",
			wantErr:  false,
		},
		{
			name:     "not enough parameters",
			baseURL:  "",
			template: "/api/v1/orgs/%v/wfgrps/%v/",
			params:   []interface{}{"my-org"},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "too many parameters",
			baseURL:  "",
			template: "/api/v1/orgs/%v/",
			params:   []interface{}{"my-org", "extra"},
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildURL(tt.baseURL, tt.template, tt.params...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestEncodeQueryParams(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]interface{}
		want   map[string][]string // url.Values is map[string][]string
	}{
		{
			name: "simple params",
			params: map[string]interface{}{
				"page":     1,
				"pageSize": 10,
			},
			want: map[string][]string{
				"page":     {"1"},
				"pageSize": {"10"},
			},
		},
		{
			name: "string params",
			params: map[string]interface{}{
				"filter": "status:active",
				"sort":   "name",
			},
			want: map[string][]string{
				"filter": {"status:active"},
				"sort":   {"name"},
			},
		},
		{
			name: "nil values are skipped",
			params: map[string]interface{}{
				"filter": nil,
				"page":   1,
			},
			want: map[string][]string{
				"page": {"1"},
			},
		},
		{
			name:   "empty params",
			params: map[string]interface{}{},
			want:   map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeQueryParams(tt.params)
			assert.Equal(t, tt.want, map[string][]string(got))
		})
	}
}
