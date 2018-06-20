// Package defaults implements the Pipe interface providing default values
// for missing configuration.
package defaults

import (
	"fmt"

	"github.com/apex/log"
	"github.com/sniperkit/goreleaser/context"
	"github.com/sniperkit/goreleaser/pipeline/archive"
	"github.com/sniperkit/goreleaser/pipeline/artifactory"
	"github.com/sniperkit/goreleaser/pipeline/brew"
	"github.com/sniperkit/goreleaser/pipeline/build"
	"github.com/sniperkit/goreleaser/pipeline/checksums"
	"github.com/sniperkit/goreleaser/pipeline/docker"
	"github.com/sniperkit/goreleaser/pipeline/env"
	"github.com/sniperkit/goreleaser/pipeline/fpm"
	"github.com/sniperkit/goreleaser/pipeline/nfpm"
	"github.com/sniperkit/goreleaser/pipeline/project"
	"github.com/sniperkit/goreleaser/pipeline/release"
	"github.com/sniperkit/goreleaser/pipeline/s3"
	"github.com/sniperkit/goreleaser/pipeline/scoop"
	"github.com/sniperkit/goreleaser/pipeline/sign"
	"github.com/sniperkit/goreleaser/pipeline/snapcraft"
	"github.com/sniperkit/goreleaser/pipeline/snapshot"
)

// Pipe that sets the defaults
type Pipe struct{}

func (Pipe) String() string {
	return "setting defaults for:"
}

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

var defaulters = []Defaulter{
	env.Pipe{},
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	archive.Pipe{},
	build.Pipe{},
	fpm.Pipe{},
	nfpm.Pipe{},
	snapcraft.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	docker.Pipe{},
	artifactory.Pipe{},
	s3.Pipe{},
	brew.Pipe{},
	scoop.Pipe{},
}

// Run the pipe
func (Pipe) Run(ctx *context.Context) error {
	if ctx.Config.Dist == "" {
		ctx.Config.Dist = "dist"
	}
	if ctx.Config.GitHubURLs.Download == "" {
		ctx.Config.GitHubURLs.Download = "https://github.com"
	}
	for _, defaulter := range defaulters {
		log.Info(defaulter.String())
		if err := defaulter.Default(ctx); err != nil {
			return err
		}
	}
	return nil
}
