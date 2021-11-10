package unicorn

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		logger.Process("Writing start command")
		// 8080 is the default unicorn port
		command := `bundle exec unicorn --listen "${PORT:-8080}"`
		logger.Subprocess(command)

		return packit.BuildResult{
			Launch: packit.LaunchMetadata{
				Processes: []packit.Process{
					{
						Type:    "web",
						Command: command,
						Default: true,
					},
				},
			},
		}, nil
	}
}
