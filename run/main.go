package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
	"github.com/paketo-community/unicorn"
)

func main() {
	parser := unicorn.NewGemfileParser()
	logger := scribe.NewLogger(os.Stdout)

	packit.Run(
		unicorn.Detect(parser),
		unicorn.Build(logger),
	)
}
