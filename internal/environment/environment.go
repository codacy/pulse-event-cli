package environment

type Environment interface {
	GetName() (*string, bool)
}

var environments = []Environment{}

// Environments where the validation is generic and might clash with others (e.g.: Environment variables used use generic names).
// These environments will only be checked if none in [environments] matches.
var otherEnvironments = []Environment{azurePipelines{}, docker{}, googleCloudBuild{}}

func GetEnvironmentName() *string {
	for i := range environments {
		if name, ok := environments[i].GetName(); ok {
			return name
		}
	}
	for i := range otherEnvironments {
		if name, ok := otherEnvironments[i].GetName(); ok {
			return name
		}
	}

	return nil
}
