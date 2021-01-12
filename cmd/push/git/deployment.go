package git

import (
	"fmt"
	"time"

	"github.com/codacy/pulse-event-cli/cmd/push"
	"github.com/codacy/pulse-event-cli/pkg/ingestion/events"
	gitmodule "github.com/gogs/git-module"
	"github.com/spf13/cobra"
)

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Push events based on git history",
	Long:  "Retrieves changes for a deployment from git and pushes the deployment event and its changes to the Pulse service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		previousDeploymentRef, _ := cmd.Flags().GetString("previous-deployment-ref")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		identifier, _ := cmd.Flags().GetString("identifier")
		timestamp, _ := cmd.Flags().GetInt64("timestamp")
		gitDirectory, _ := getGitDirectory(cmd)
		cmd.SilenceUsage = true

		apiClient, err := push.GetAPIClient(cmd)
		if err != nil {
			return err
		}

		changes, err := getChanges(gitDirectory, previousDeploymentRef)
		if err != nil {
			return err
		}

		var changesIds []string
		for _, commit := range changes {
			change := events.Change{
				Source:      "git",
				ChangeID:    commit.ID.String(),
				TimeCreated: commit.Author.When,
				EventType:   "commit",
				Type:        "change",
			}

			changesIds = append(changesIds, change.ChangeID)

			fmt.Printf("Found change %s\n", change.ChangeID)

			if !dryRun {
				err := apiClient.CreateEvent(change)
				if err != nil {
					return fmt.Errorf("failed to upload change: %v", err)
				}
			}
		}

		deployment := events.Deployment{Source: "git", DeployID: identifier, TimeCreated: time.Unix(timestamp, 0), Changes: changesIds, Type: "deployment"}
		fmt.Printf("Deployment %s\n", deployment.DeployID)
		if !dryRun {
			err = apiClient.CreateEvent(&deployment)

			if err != nil {
				return fmt.Errorf("failed to upload deployment: %v", err)
			}
		}

		return nil
	},
}

func getChanges(gitDirectory string, previousDeploymentRef string) ([]*gitmodule.Commit, error) {
	revisionInterval := fmt.Sprintf("%s..HEAD", previousDeploymentRef)
	commits, err := gitmodule.RepoLog(gitDirectory, revisionInterval, gitmodule.LogOptions{})
	if err != nil {
		return []*gitmodule.Commit{}, fmt.Errorf("failed to get git repository: %v", err)
	}

	return commits, nil
}

func init() {
	gitCmd.AddCommand(deploymentCmd)

	deploymentCmd.Flags().String("previous-deployment-ref", "", "git reference of the previous deployment (commit SHA or tag)")
	deploymentCmd.MarkFlagRequired("previous-deployment-ref")

	deploymentCmd.Flags().String("identifier", "", "deployment identifer (e.g.: commit SHA)")
	deploymentCmd.MarkFlagRequired("identifier")
	deploymentCmd.Flags().Int64("timestamp", 0, "deployment timestamp (e.g.: 1602253523)")
	deploymentCmd.MarkFlagRequired("timestamp")
	deploymentCmd.Flags().String("source", "cli", "deployment source (e.g.: cli, git, GitHub)")

	deploymentCmd.Flags().Bool("dry-run", false, "do not push the events, only print them to stdout")
}
