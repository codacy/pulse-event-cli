package environment

import "os"

type gitHubActions struct{}

var gitHubActionsName = "github-actions"

func (gha gitHubActions) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("GITHUB_ACTIONS"); ok {
		return &gitHubActionsName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, gitHubActions{})
}
