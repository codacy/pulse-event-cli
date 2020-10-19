package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// incidentCmd represents the incident command
var incidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Push incident event",
	Long:  `Pushes a incident event to the Pulse service.`,
	Run: func(cmd *cobra.Command, args []string) {
		identifier, _ := cmd.Flags().GetString("identifier")
		timestampCreated, _ := cmd.Flags().GetInt64("timestampCreated")
		timestampResolved, _ := cmd.Flags().GetInt64("timestampResolved")
		source, _ := cmd.Flags().GetString("source")

		fmt.Print("Pushing incident event with identifier ", identifier, ", created timestamp ", time.Unix(timestampCreated, 0), ", resolved timestamp ", time.Unix(timestampResolved, 0), ", source ", source, "\n")

		items := []*incident{{Source: source, IncidentID: identifier, TimeCreated: time.Unix(timestampCreated, 0), TimeResolved: time.Unix(timestampResolved, 0)}}
		createEvent("incidents", items)
	},
}

func init() {
	pushCmd.AddCommand(incidentCmd)
	incidentCmd.Flags().String("identifier", "", "Incident identifer (e.g.: commit sha)")
	incidentCmd.MarkFlagRequired("identifier")
	incidentCmd.Flags().Int64("timestampCreated", 0, "Incident created timestamp (e.g.: 1602253523)")
	incidentCmd.MarkFlagRequired("timestampCreated")
	incidentCmd.Flags().Int64("timestampResolved", 0, "Incident resolved timestamp (e.g.: 1602253524)")
	incidentCmd.MarkFlagRequired("timestampResolved")
	incidentCmd.Flags().String("source", "cli", "Incident source (e.g.: cli, git, GitHub)")
}

// BigQuery Incident table schema
type incident struct {
	Source       string    `bigquery:"source"`
	IncidentID   string    `bigquery:"incident_id"`
	TimeCreated  time.Time `bigquery:"time_created"`
	TimeResolved time.Time `bigquery:"time_resolved"`
}
