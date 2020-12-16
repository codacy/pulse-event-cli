package environment

import "os"

type jenkins struct{}

var jenkinsName = "jenkins"

func (e jenkins) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("JENKINS_URL"); ok {
		return &jenkinsName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, jenkins{})
}
