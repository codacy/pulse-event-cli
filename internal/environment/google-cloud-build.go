package environment

import "os"

type googleCloudBuild struct{}

var googleCloudBuildName = "google-cloud-build"

func (e googleCloudBuild) GetName() (*string, bool) {
	_, hasProjectID := os.LookupEnv("PROJECT_ID")
	_, hasBuildID := os.LookupEnv("BUILD_ID")
	_, hasCommitSHA := os.LookupEnv("COMMIT_SHA")
	_, hasRevisionID := os.LookupEnv("REVISION_ID")
	if hasProjectID && hasBuildID && hasCommitSHA && hasRevisionID {
		return &googleCloudBuildName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, googleCloudBuild{})
}
