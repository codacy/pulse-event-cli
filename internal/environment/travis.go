package environment

import "os"

type travis struct{}

var travisName = "travis"

func (e travis) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("TRAVIS"); ok {
		return &travisName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, travis{})
}
