package environment

import "os"

type gocd struct{}

var gocdName = "gocd"

func (e gocd) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("GO_PIPELINE_NAME"); ok {
		return &gocdName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, gocd{})
}
