module github.com/paketo-buildpacks/unicorn

go 1.16

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/onsi/gomega v1.33.1
	github.com/paketo-buildpacks/occam v0.18.5
	github.com/paketo-buildpacks/packit/v2 v2.13.0
	github.com/sclevine/spec v1.4.0
	gotest.tools/v3 v3.5.1 // indirect
)

replace github.com/CycloneDX/cyclonedx-go => github.com/CycloneDX/cyclonedx-go v0.6.0
