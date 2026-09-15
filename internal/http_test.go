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
			args: []interface{}{"org", PathWithSlashes("wfg-module-testing/wfg-automatedTesting-3irgxu3n")},
			want: "/orgs/org/wfgrps/wfg-module-testing/wfg-automatedTesting-3irgxu3n/",
		},
		{
			desc: "PathWithSlashes escapes within segments",
			args: []interface{}{"org", PathWithSlashes("a b/c?d")},
			want: "/orgs/org/wfgrps/a%20b/c%3Fd/",
		},
		{
			desc: "PathWithSlashes without slashes",
			args: []interface{}{"org", PathWithSlashes("grp")},
			want: "/orgs/org/wfgrps/grp/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert.Equal(t, tt.want, EncodeURL("/orgs/%v/wfgrps/%v/", tt.args...))
		})
	}
}
