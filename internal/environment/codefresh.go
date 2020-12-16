package environment

import "os"

type codefresh struct{}

var codefreshName = "codefresh"

func (e codefresh) GetName() (*string, bool) {
	_, hasBuildID := os.LookupEnv("CF_BUILD_ID")
	_, hasBuildURL := os.LookupEnv("CF_BUILD_URL")
	if hasBuildID && hasBuildURL {
		return &codefreshName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, codefresh{})
}
