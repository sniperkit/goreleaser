package nametemplate

import (
	"testing"

	"github.com/sniperkit/goreleaser/config"
	"github.com/sniperkit/goreleaser/context"
	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  string
	}{
		{
			desc: "with env",
			in:   "{{ .Env.FOO }}",
			out:  "BAR",
		},
		{
			desc: "with env",
			in:   "{{ .Env.BAR }}",
			out:  "",
		},
	}
	var ctx = context.New(config.Project{})
	ctx.Env = map[string]string{
		"FOO": "BAR",
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out, _ := Apply(ctx, tC.in)
			assert.Equal(t, tC.out, out)
		})
	}
}

func TestFuncMap(t *testing.T) {
	var ctx = context.New(config.Project{
		ProjectName: "proj",
	})
	for _, tc := range []struct {
		Template string
		Name     string
	}{
		{
			Template: `{{ time "2006-01-02" }}`,
			Name:     "YYYY-MM-DD",
		},
		{
			Template: `{{ time "01/02/2006" }}`,
			Name:     "MM/DD/YYYY",
		},
		{
			Template: `{{ time "01/02/2006" }}`,
			Name:     "MM/DD/YYYY",
		},
	} {
		out, err := Apply(ctx, tc.Template)
		assert.NoError(t, err)
		assert.NotEmpty(t, out)
	}
}

func TestInvalidTemplate(t *testing.T) {
	_, err := Apply(context.New(config.Project{}), "{{{.Foo}")
	assert.EqualError(t, err, "template: release:1: unexpected \"{\" in command")
}
