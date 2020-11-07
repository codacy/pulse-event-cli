package main

import (
	"github.com/codacy/event-cli/cmd"
	_ "github.com/codacy/event-cli/cmd/push"
)

func main() {
	cmd.Execute()
}
