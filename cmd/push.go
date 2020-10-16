package cmd

import (
	"github.com/spf13/cobra"
)

// CredentialsString authenticates the user
var CredentialsString string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:              "push",
	Short:            "Push events",
	Long:             `Pushes all the different events necessary to generate your metrics in the Pulse service.`,
	TraverseChildren: true,
	Run:              func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().StringVar(&CredentialsString, "credentials", "", "Cedentials to authenticate the user")
	pushCmd.MarkFlagRequired("credentials")
}
