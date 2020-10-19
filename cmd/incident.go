package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var incidentIdentifier string
var incidentTimestampCreated int64
var incidentTimestampResolved int64
var incidentSource string

// incidentCmd represents the incident command
var incidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Push incident event",
	Long:  `Pushes a incident event to the Pulse service.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Pushing incident event with identifier ", incidentIdentifier, ", created timestamp ", time.Unix(incidentTimestampCreated, 0), ", resolved timestamp ", time.Unix(incidentTimestampResolved, 0), ", source ", incidentSource, "\n")

		items := []*incident{{Source: incidentSource, IncidentID: incidentIdentifier, TimeCreated: time.Unix(incidentTimestampCreated, 0), TimeResolved: time.Unix(incidentTimestampResolved, 0)}}
		CreateEvent("incidents", items)
	},
}

func init() {
	pushCmd.AddCommand(incidentCmd)
	incidentCmd.Flags().StringVar(&incidentIdentifier, "identifier", "", "Incident identifer (e.g.: commit sha)")
	incidentCmd.MarkFlagRequired("identifier")
	incidentCmd.Flags().Int64Var(&incidentTimestampCreated, "timestampCreated", 0, "Incident created timestamp (e.g.: 1602253523)")
	incidentCmd.MarkFlagRequired("timestampCreated")
	incidentCmd.Flags().Int64Var(&incidentTimestampResolved, "timestampResolved", 0, "Incident resolved timestamp (e.g.: 1602253524)")
	incidentCmd.MarkFlagRequired("timestampResolved")
	incidentCmd.Flags().StringVar(&incidentSource, "source", "cli", "Incident source (e.g.: cli, git, GitHub)")
}

// BigQuery Incident table schema
type incident struct {
	Source       string    `bigquery:"source"`
	IncidentID   string    `bigquery:"incident_id"`
	TimeCreated  time.Time `bigquery:"time_created"`
	TimeResolved time.Time `bigquery:"time_resolved"`
}
