package testlib

import (
	"testing"

	"github.com/sniperkit/goreleaser/pipeline"
)

func TestAssertSkipped(t *testing.T) {
	AssertSkipped(t, pipeline.Skip("skip"))
}
