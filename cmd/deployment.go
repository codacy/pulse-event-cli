package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Push deployment event",
	Long:  `Pushes a deployment event to the Pulse service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		identifier, _ := cmd.Flags().GetString("identifier")
		timestamp, _ := cmd.Flags().GetInt64("timestamp")
		source, _ := cmd.Flags().GetString("source")

		fmt.Print("Pushing deployment event with identifier ", identifier, ", timestamp ", time.Unix(timestamp, 0), ", source ", source, " and changes ", args, "\n")

		items := []*deployment{{Source: source, DeployID: identifier, TimeCreated: time.Unix(timestamp, 0), Changes: args}}
		CreateEvent("deployments", items)
	},
}

func init() {
	pushCmd.AddCommand(deploymentCmd)
	deploymentCmd.Flags().String("identifier", "", "Deployment identifer (e.g.: commit sha)")
	deploymentCmd.MarkFlagRequired("identifier")
	deploymentCmd.Flags().Int64("timestamp", 0, "Deployment timestamp (e.g.: 1602253523)")
	deploymentCmd.MarkFlagRequired("timestamp")
	deploymentCmd.Flags().String("source", "cli", "Deployment source (e.g.: cli, git, GitHub)")
}

// BigQuery Deployments table schema
type deployment struct {
	Source      string    `bigquery:"source"`
	DeployID    string    `bigquery:"deploy_id"`
	TimeCreated time.Time `bigquery:"time_created"`
	Changes     []string  `bigquery:"changes"`
}
