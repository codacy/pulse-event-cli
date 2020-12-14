package environment

type Environment interface {
	GetName() (*string, bool)
}

var environments = []Environment{}

func GetEnvironmentName() *string {
	for i := range environments {
		if name, ok := environments[i].GetName(); ok {
			return name
		}
	}

	return nil
}
