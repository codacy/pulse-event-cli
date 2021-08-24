package git

import (
	"bytes"
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
		teams, _ := cmd.Flags().GetStringSlice("teams")
		gitDirectory, _ := getGitDirectory(cmd)
		cmd.SilenceUsage = true

		apiClient, err := push.GetAPIClient(cmd)
		if err != nil {
			return err
		}

		changes, err := getChanges(gitDirectory, previousDeploymentRef, args)
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

		deployment := events.Deployment{
			Source:      "git",
			DeployID:    identifier,
			TimeCreated: time.Unix(timestamp, 0),
			Changes:     changesIds,
			Teams:       teams,
			Type:        "deployment",
		}
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

func getChanges(gitDirectory string, previousDeploymentRef string, paths []string) ([]*gitmodule.Commit, error) {
	revisionInterval := fmt.Sprintf("%s..HEAD", previousDeploymentRef)
	commits, err := RepoLog(gitDirectory, revisionInterval, paths)
	if err != nil {
		return []*gitmodule.Commit{}, fmt.Errorf("failed to get git repository: %v", err)
	}

	return commits, nil
}

// RepoLog returns a list of commits in the state of given revision of the repository
// in given path. The returned list is in reverse chronological order.
//
// This method is sourced from "github.com/gogs/git-module" to add support for multiple paths
func RepoLog(repoPath string, rev string, paths []string) ([]*gitmodule.Commit, error) {
	r, err := gitmodule.Open(repoPath)
	if err != nil {
		return nil, fmt.Errorf("open: %v", err)
	}

	cmd := gitmodule.NewCommand("log", "--pretty="+gitmodule.LogFormatHashOnly, rev)

	if len(paths) > 0 {
		cmd.AddArgs("--")
		for _, path := range paths {
			cmd.AddArgs(path)
		}
	}

	logOptions := gitmodule.LogOptions{}

	stdout, err := cmd.RunInDirWithTimeout(logOptions.Timeout, repoPath)
	if err != nil {
		return nil, err
	}

	// parsePrettyFormatLogToList returns a list of commits parsed from given logs that are
	// formatted in LogFormatHashOnly.
	parsePrettyFormatLogToList := func(timeout time.Duration, logs []byte) ([]*gitmodule.Commit, error) {
		if len(logs) == 0 {
			return []*gitmodule.Commit{}, nil
		}

		var err error
		ids := bytes.Split(logs, []byte{'\n'})
		commits := make([]*gitmodule.Commit, len(ids))
		for i, id := range ids {
			commits[i], err = r.CatFileCommit(string(id), gitmodule.CatFileCommitOptions{Timeout: timeout})
			if err != nil {
				return nil, err
			}
		}
		return commits, nil
	}
	return parsePrettyFormatLogToList(logOptions.Timeout, stdout)
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
	deploymentCmd.Flags().StringSlice("teams", []string{}, "A comma separated list of teams responsible for the changes in this deployment (e.g.: jupiter,mercury)")

	deploymentCmd.Flags().Bool("dry-run", false, "do not push the events, only print them to stdout")
}
