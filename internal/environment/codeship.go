package environment

import "os"

type codeship struct{}

var codeshipName = "codeship"

func (e codeship) GetName() (*string, bool) {
	if os.Getenv("CI_NAME") == "codeship" {
		return &codeshipName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, codeship{})
}
