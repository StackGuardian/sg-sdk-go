package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeURL(t *testing.T) {
	tests := []struct {
		desc string
		args []interface{}
		want string
	}{
		{
			desc: "plain args escape slashes",
			args: []interface{}{"org", "a/b"},
			want: "/orgs/org/wfgrps/a%2Fb/",
		},
		{
			desc: "PathWithSlashes keeps slashes",
			args: []interface{}{"org", pathWithSlashes("wfg-module-testing/wfg-automatedTesting-3irgxu3n")},
			want: "/orgs/org/wfgrps/wfg-module-testing/wfg-automatedTesting-3irgxu3n/",
		},
		{
			desc: "PathWithSlashes escapes within segments",
			args: []interface{}{"org", pathWithSlashes("a b/c?d")},
			want: "/orgs/org/wfgrps/a%20b/c%3Fd/",
		},
		{
			desc: "PathWithSlashes without slashes",
			args: []interface{}{"org", pathWithSlashes("grp")},
			want: "/orgs/org/wfgrps/grp/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert.Equal(t, tt.want, EncodeURL("/orgs/%v/wfgrps/%v/", tt.args...))
		})
	}
}

func TestPathWithSlashesValidate(t *testing.T) {
	tests := []struct {
		desc    string
		path    pathWithSlashes
		wantErr bool
	}{
		{desc: "single segment", path: "grp", wantErr: false},
		{desc: "nested segments", path: "parent/child", wantErr: false},
		{desc: "deeply nested segments", path: "a/b/c", wantErr: false},
		{desc: "empty value", path: "", wantErr: true},
		{desc: "doubled slash", path: "parent//child", wantErr: true},
		{desc: "leading slash", path: "/child", wantErr: true},
		{desc: "trailing slash", path: "parent/", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := tt.path.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewPathWithSlashes(t *testing.T) {
	t.Run("valid value is returned unchanged", func(t *testing.T) {
		p, err := NewPathWithSlashes("parent/child")
		assert.NoError(t, err)
		assert.Equal(t, pathWithSlashes("parent/child"), p)
	})

	t.Run("invalid value returns an error and the zero value", func(t *testing.T) {
		p, err := NewPathWithSlashes("parent//child")
		assert.Error(t, err)
		assert.Equal(t, pathWithSlashes(""), p)
	})
}
