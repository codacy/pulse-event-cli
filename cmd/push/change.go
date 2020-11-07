package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Push change event",
	Long:  `Pushes a change event to the Pulse service.`,
	Run: func(cmd *cobra.Command, args []string) {
		identifier, _ := cmd.Flags().GetString("identifier")
		timestamp, _ := cmd.Flags().GetInt64("timestamp")
		source, _ := cmd.Flags().GetString("source")
		eventType, _ := cmd.Flags().GetString("event_type")

		fmt.Print("Pushing change event with identifier ", identifier, ", timestamp ", time.Unix(timestamp, 0), ", source ", source, " and event type ", eventType, "\n")

		item := change{Source: source, ChangeID: identifier, TimeCreated: time.Unix(timestamp, 0), EventType: eventType, Type: "change"}
		itemBytes, _ := json.Marshal(item)
		createEvent(itemBytes)
	},
}

func init() {
	pushCmd.AddCommand(changeCmd)
	changeCmd.Flags().String("identifier", "", "Change identifer (e.g.: commit sha)")
	changeCmd.MarkFlagRequired("identifier")
	changeCmd.Flags().Int64("timestamp", 0, "Change timestamp (e.g.: 1602253523)")
	changeCmd.MarkFlagRequired("timestamp")
	changeCmd.Flags().String("source", "cli", "Change source (e.g.: cli, git, GitHub)")
	changeCmd.Flags().String("event_type", "commit", "Event type (e.g.: commit, push)")
}

type change struct {
	Source      string    `json:"source"`
	ChangeID    string    `json:"change_id"`
	TimeCreated time.Time `json:"time_created"`
	EventType   string    `json:"event_type"`
	Type        string    `json:"$type"`
}
