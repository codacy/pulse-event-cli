package environment

import "os"

type appveyor struct{}

var appveyorName = "appveyor"

func (e appveyor) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("APPVEYOR"); ok {
		return &appveyorName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, appveyor{})
}
