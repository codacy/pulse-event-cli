package environment

import "os"

type docker struct{}

var dockerName = "docker"

func (e docker) GetName() (*string, bool) {
	_, hasSourceBranch := os.LookupEnv("SOURCE_BRANCH")
	_, hasSourceCommit := os.LookupEnv("SOURCE_COMMIT")
	_, hasDockerRepo := os.LookupEnv("DOCKER_REPO")
	_, hasDockerfilePath := os.LookupEnv("DOCKERFILE_PATH")
	if hasSourceBranch && hasSourceCommit && hasDockerRepo && hasDockerfilePath {
		return &dockerName, true
	}

	return nil, false
}
