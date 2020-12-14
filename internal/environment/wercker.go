package environment

import "os"

type wercker struct{}

var werckerName = "wercker"

func (e wercker) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("WERCKER_GIT_BRANCH"); ok {
		return &werckerName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, wercker{})
}
