package environment

import "os"

type buddy struct{}

var buddyName = "buddy"

func (e buddy) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("BUDDY"); ok {
		return &buddyName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, buddy{})
}
