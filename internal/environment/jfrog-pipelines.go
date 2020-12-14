package environment

import "os"

type jfrogPipelines struct{}

var jfrogPipelinesName = "jfrog-pipelines"

func (e jfrogPipelines) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("JFROG_CLI_BUILD_NUMBER"); ok {
		return &jfrogPipelinesName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, jfrogPipelines{})
}
