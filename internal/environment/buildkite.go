package environment

import "os"

type buildkite struct{}

var buildkiteName = "buildkite"

func (e buildkite) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("BUILDKITE"); ok {
		return &buildkiteName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, buildkite{})
}
