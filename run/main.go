package main

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/paketo-buildpacks/unicorn"
)

func main() {
	parser := unicorn.NewGemfileParser()
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))

	packit.Run(
		unicorn.Detect(parser),
		unicorn.Build(logger),
	)
}
