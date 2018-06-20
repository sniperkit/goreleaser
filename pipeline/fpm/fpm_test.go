package fpm

import (
	"testing"

	"github.com/sniperkit/goreleaser/config"
	"github.com/sniperkit/goreleaser/context"
	"github.com/stretchr/testify/assert"
)

func TestDescription(t *testing.T) {
	assert.NotEmpty(t, Pipe{}.String())
}

func TestDefault(t *testing.T) {
	var ctx = &context.Context{
		Config: config.Project{
			FPM: config.NFPM{
				Formats: []string{"deb"},
			},
		},
	}
	assert.NoError(t, Pipe{}.Default(ctx))
	assert.Equal(t, ctx.Config.FPM, ctx.Config.NFPM)
}

func TestDefaultSet(t *testing.T) {
	var ctx = &context.Context{
		Config: config.Project{
			NFPM: config.NFPM{
				Formats: []string{"deb"},
			},
		},
	}
	assert.NoError(t, Pipe{}.Default(ctx))
	assert.Equal(t, config.NFPM{}, ctx.Config.FPM)
}
