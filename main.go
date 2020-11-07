package main

import (
	"github.com/codacy/event-cli/cmd"
	_ "github.com/codacy/event-cli/cmd/push"
	_ "github.com/codacy/event-cli/cmd/push/git"
)

func main() {
	cmd.Execute()
}
