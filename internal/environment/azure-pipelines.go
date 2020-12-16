package environment

import "os"

type azurePipelines struct{}

var azurePipelinesName = "azure-pipelines"

func (e azurePipelines) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("TF_BUILD"); ok {
		return &azurePipelinesName, true
	}

	return nil, false
}
