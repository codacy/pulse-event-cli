package environment

import "os"

type awsCodebuild struct{}

var awsCodebuildName = "aws-codebuild"

func (e awsCodebuild) GetName() (*string, bool) {
	if _, ok := os.LookupEnv("CODEBUILD_BUILD_ID"); ok {
		return &awsCodebuildName, true
	}

	return nil, false
}

func init() {
	environments = append(environments, awsCodebuild{})
}
