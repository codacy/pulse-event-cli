package environment

import "os"

type bitrise struct{}

var bitriseName = "bitrise"

func (e bitrise) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("BITRISE_IO"); ok {
		return &bitriseName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, bitrise{})
}
