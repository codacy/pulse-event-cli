package environment

import "os"

type teamcity struct{}

var teamcityName = "teamcity"

func (e teamcity) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("TEAMCITY_VERSION"); ok {
		return &teamcityName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, teamcity{})
}
