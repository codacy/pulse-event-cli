package environment

import "os"

type drone struct{}

var droneName = "drone"

func (e drone) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("DRONE"); ok {
		return &droneName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, drone{})
}
