package unicorn

import (
	"fmt"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/fs"
)

//go:generate faux --interface Parser --output fakes/parser.go
type Parser interface {
	Parse(path string) (hasUnicorn bool, err error)
}

type BuildPlanMetadata struct {
	Launch bool `toml:"launch"`
}

func Detect(gemfileParser Parser) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		exists, err := fs.Exists(filepath.Join(context.WorkingDir, "config.ru"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to stat 'config.ru': %w", err)
		}

		if !exists {
			return packit.DetectResult{}, packit.Fail.WithMessage("no 'config.ru' file found")
		}

		hasUnicorn, err := gemfileParser.Parse(filepath.Join(context.WorkingDir, "Gemfile"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to parse Gemfile: %w", err)
		}

		if !hasUnicorn {
			return packit.DetectResult{}, packit.Fail.WithMessage("unicorn was not found in the Gemfile")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "gems",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
					{
						Name: "bundler",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
					{
						Name: "mri",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
				},
			},
		}, nil
	}
}
