package git

import (
	"github.com/codacy/event-cli/cmd/push"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:              "git",
	Short:            "Git events",
	Long:             "Utility command to get events from git",
	TraverseChildren: true,
	Run:              func(cmd *cobra.Command, args []string) {},
}

func init() {
	push.PushCmd.AddCommand(gitCmd)

	// TODO add git directory config
	// deploymentCmd.PersistentFlags().String("git-directory", "./", "The location of the git repo")
}
