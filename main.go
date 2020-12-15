package main

import (
	"github.com/codacy/pulse-event-cli/cmd"
	// These imports are required in order to load the commands subpackages
	_ "github.com/codacy/pulse-event-cli/cmd/push"
	_ "github.com/codacy/pulse-event-cli/cmd/push/git"
)

func main() {
	cmd.Execute()
}
