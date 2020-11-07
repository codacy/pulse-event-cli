package main

import (
	"github.com/codacy/event-cli/cmd"
	// This imports are required in order to load the commands subpackages
	_ "github.com/codacy/event-cli/cmd/push"
	_ "github.com/codacy/event-cli/cmd/push/git"
)

func main() {
	cmd.Execute()
}
