package environment

import "os"

type gitlabCI struct{}

var gitlabCIName = "gitlab-ci"

func (e gitlabCI) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("GITLAB_CI"); ok {
		return &gitlabCIName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, gitlabCI{})
}
