package environment

import "os"

type solano struct{}

var solanoName = "solano"

func (e solano) GetName() (*string, bool) {
	_, hasSolano := os.LookupEnv("SOLANO")
	_, hasTddium := os.LookupEnv("TDDIUM")
	if hasSolano || hasTddium {
		return &solanoName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, solano{})
}
