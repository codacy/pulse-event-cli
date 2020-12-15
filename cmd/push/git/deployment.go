package git

import (
	"fmt"
	"time"

	"github.com/codacy/pulse-event-cli/cmd/push"
	"github.com/codacy/pulse-event-cli/pkg/ingestion/events"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
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

		changesIterator, err := getChanges(gitDirectory, previousDeploymentRef)
		if err != nil {
			return err
		}

		var changesIds []string

		err = changesIterator.ForEach(func(commit *object.Commit) error {
			change := events.Change{
				Source:      "git",
				ChangeID:    commit.Hash.String(),
				TimeCreated: commit.Author.When,
				EventType:   "commit",
				Type:        "change",
			}

			changesIds = append(changesIds, change.ChangeID)

			fmt.Printf("Found change %s\n", change.ChangeID)

			if !dryRun {
				return apiClient.CreateEvent(change)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("failed to upload changes: %v", err)
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

func getChanges(gitDirectory string, previousDeploymentRef string) (object.CommitIter, error) {
	repo, err := git.PlainOpen(gitDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to get git repository: %v", err)
	}

	previousDeploymentHash, err := repo.ResolveRevision(plumbing.Revision(previousDeploymentRef))
	if err != nil {
		return nil, fmt.Errorf("could not get previous deployment revision: %v", err)
	}

	headReference, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("could not get repository HEAD: %v", err)
	}

	currentCommit, err := repo.CommitObject(headReference.Hash())
	if err != nil {
		return nil, fmt.Errorf("could not get repository HEAD: %v", err)
	}

	commitStopFilter := object.CommitFilter(func(commit *object.Commit) bool { return commit.Hash == *previousDeploymentHash })

	return object.NewFilterCommitIter(currentCommit, nil, &commitStopFilter), nil
}

func init() {
	gitCmd.AddCommand(deploymentCmd)

	deploymentCmd.Flags().String("previous-deployment-ref", "", "Git reference of the previous deployment (commit SHA or tag)")
	deploymentCmd.MarkFlagRequired("previous-deployment-ref")

	deploymentCmd.Flags().String("identifier", "", "Deployment identifer (e.g.: commit SHA)")
	deploymentCmd.MarkFlagRequired("identifier")
	deploymentCmd.Flags().Int64("timestamp", 0, "Deployment timestamp (e.g.: 1602253523)")
	deploymentCmd.MarkFlagRequired("timestamp")
	deploymentCmd.Flags().String("source", "cli", "Deployment source (e.g.: cli, git, GitHub)")

	deploymentCmd.Flags().Bool("dry-run", false, "Do not push the events, just print them to stdout")
}
