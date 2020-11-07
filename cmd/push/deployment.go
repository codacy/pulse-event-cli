package push

import (
	"fmt"
	"time"

	"github.com/codacy/event-cli/pkg/ingestion/events"
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

		item := events.Deployment{Source: source, DeployID: identifier, TimeCreated: time.Unix(timestamp, 0), Changes: args, Type: "deployment"}
		apiClient.CreateEvent(&item)
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
