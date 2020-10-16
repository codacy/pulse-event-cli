package environment

import "os"

type codemagic struct{}

var codemagicName = "codemagic"

func (e codemagic) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("FCI_PROJECT_ID"); ok {
		return &codemagicName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, codemagic{})
}
