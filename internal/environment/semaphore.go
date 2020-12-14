package environment

import "os"

type semaphore struct{}

var semaphoreName = "semaphore"

func (e semaphore) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("SEMAPHORE"); ok {
		return &semaphoreName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, semaphore{})
}
