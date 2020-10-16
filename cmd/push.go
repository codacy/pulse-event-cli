package cmd

import (
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push events",
	Long:  `Pushes all the different events necessary to generate your metrics in the Pulse service.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
