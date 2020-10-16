package environment

import "os"

type herokuCI struct{}

var herokuCIName = "heroku-ci"

func (e herokuCI) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("HEROKU_TEST_RUN_ID"); ok {
		return &herokuCIName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, herokuCI{})
}
