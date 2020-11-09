package git

import (
	"log"
	"os"

	"github.com/codacy/event-cli/cmd/push"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:              "git",
	Short:            "Git events",
	Long:             "Utility command to get events from git",
	TraverseChildren: true,
}

func init() {
	push.PushCmd.AddCommand(gitCmd)

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	deploymentCmd.PersistentFlags().String("directory", path, "The directory where the git repository can be found")
}

func getGitDirectory(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("directory")
}
