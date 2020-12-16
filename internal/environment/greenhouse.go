package environment

import "os"

type greenhouse struct{}

var greenhouseName = "greenhouse"

func (e greenhouse) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("GREENHOUSE"); ok {
		return &greenhouseName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, greenhouse{})
}
