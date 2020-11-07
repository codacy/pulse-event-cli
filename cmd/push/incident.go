package push

import (
	"fmt"
	"time"

	"github.com/codacy/event-cli/pkg/ingestion/events"
	"github.com/spf13/cobra"
)

// incidentCmd represents the incident command
var incidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Push incident event",
	Long:  `Pushes a incident event to the Pulse service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		identifier, _ := cmd.Flags().GetString("identifier")
		timestampCreated, _ := cmd.Flags().GetInt64("timestampCreated")
		timestampResolved, _ := cmd.Flags().GetInt64("timestampResolved")
		source, _ := cmd.Flags().GetString("source")
		cmd.SilenceUsage = true

		apiClient, err := GetAPIClient(cmd)
		if err != nil {
			return err
		}

		fmt.Print("Pushing incident event with identifier ", identifier, ", created timestamp ", time.Unix(timestampCreated, 0), ", resolved timestamp ", time.Unix(timestampResolved, 0), ", source ", source, "\n")

		item := events.Incident{Source: source, IncidentID: identifier, TimeCreated: time.Unix(timestampCreated, 0), TimeResolved: time.Unix(timestampResolved, 0), Type: "incident"}
		return apiClient.CreateEvent(item)
	},
}

func init() {
	PushCmd.AddCommand(incidentCmd)
	incidentCmd.Flags().String("identifier", "", "Incident identifer (e.g.: commit sha)")
	incidentCmd.MarkFlagRequired("identifier")
	incidentCmd.Flags().Int64("timestampCreated", 0, "Incident created timestamp (e.g.: 1602253523)")
	incidentCmd.MarkFlagRequired("timestampCreated")
	incidentCmd.Flags().Int64("timestampResolved", 0, "Incident resolved timestamp (e.g.: 1602253524)")
	incidentCmd.MarkFlagRequired("timestampResolved")
	incidentCmd.Flags().String("source", "cli", "Incident source (e.g.: cli, git, GitHub)")
}
