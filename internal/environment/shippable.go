package environment

import "os"

type shippable struct{}

var shippableName = "shippable"

func (e shippable) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("SHIPPABLE"); ok {
		return &shippableName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, shippable{})
}
