package environment

import "os"

type magnum struct{}

var magnumName = "magnum"

func (e magnum) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("MAGNUM"); ok {
		return &magnumName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, magnum{})
}
