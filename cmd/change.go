package cmd

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var changeIdentifier string
var changeTimestamp int64
var changeSource string
var changeEventType string

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Push change event",
	Long:  `Pushes a change event to the Pulse service.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Pushing change event with identifier ", changeIdentifier, ", timestamp ", time.Unix(changeTimestamp, 0), ", source ", changeSource, " and event type ", changeEventType, "\n")

		credentials, credentialsBytes := GetCredentials()

		ctx := context.Background()
		clientOptions := option.WithCredentialsJSON(credentialsBytes)
		client, err := bigquery.NewClient(ctx, credentials.ProjectID, clientOptions)
		if err != nil {
			fmt.Println(err)
		}

		ins := client.Dataset(credentials.DataSet).Table("changes").Inserter()
		items := []*change{{Source: changeSource, ChangeID: changeIdentifier, TimeCreated: time.Unix(changeTimestamp, 0), EventType: changeEventType}}
		if err := ins.Put(ctx, items); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	pushCmd.AddCommand(changeCmd)
	changeCmd.Flags().StringVar(&changeIdentifier, "identifier", "", "Change identifer (e.g.: commit sha)")
	changeCmd.MarkFlagRequired("identifier")
	changeCmd.Flags().Int64Var(&changeTimestamp, "timestamp", 0, "Change timestamp (e.g.: 1602253523)")
	changeCmd.MarkFlagRequired("timestamp")
	changeCmd.Flags().StringVar(&changeSource, "source", "cli", "Change source (e.g.: cli, git, GitHub)")
	changeCmd.Flags().StringVar(&changeEventType, "event_type", "commit", "Event type (e.g.: commit, push)")
}

// BigQuery Change table schema
type change struct {
	Source      string    `bigquery:"source"`
	ChangeID    string    `bigquery:"change_id"`
	TimeCreated time.Time `bigquery:"time_created"`
	EventType   string    `bigquery:"event_type"`
}
