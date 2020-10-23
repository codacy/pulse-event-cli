package cmd

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var apiKey string
var baseURL string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:              "push",
	Short:            "Push events",
	Long:             `Pushes all the different events necessary to generate your metrics in the Pulse service.`,
	TraverseChildren: true,
	Run:              func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "The API key to authenticate the organization/system")
	pushCmd.MarkFlagRequired("api-key")
	pushCmd.PersistentFlags().StringVar(&baseURL, "base-url", "https://ingestion.acceleratedevops.net", "The API base url")
	pushCmd.MarkFlagRequired("base-url")
}

// createEvent creates the events in the data store
func createEvent(json []byte) {
	url := strings.TrimSuffix(baseURL, "/") + "/v1/ingestion/cli?api_key=" + apiKey

	_, err := http.Post(url, "application/json", bytes.NewBuffer(json))

	if err != nil {
		panic(err)
	}
}
