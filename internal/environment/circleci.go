package environment

import "os"

type circleci struct{}

var circleciName = "circleci"

func (e circleci) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("CIRCLECI"); ok {
		return &circleciName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, circleci{})
}
