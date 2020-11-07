package push

import (
	"github.com/codacy/event-cli/cmd"
	"github.com/codacy/event-cli/pkg/ingestion"
	"github.com/spf13/cobra"
)

var apiClient *ingestion.PulseIngestionAPIClient

var pushCmd = &cobra.Command{
	Use:              "push",
	Short:            "Push events",
	Long:             `Pushes all the different events necessary to generate your metrics in the Pulse service.`,
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ := cmd.Flags().GetString("api-key")
		baseURL, _ := cmd.Flags().GetString("base-url")
		system, _ := cmd.Flags().GetString("system")

		var err error
		apiClient, err = ingestion.NewPulseIngestionAPIClient(baseURL, apiKey, system)

		return err
	},
}

func init() {
	cmd.RootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().String("api-key", "", "The API key to authenticate the organization/system")
	pushCmd.MarkFlagRequired("api-key")
	pushCmd.PersistentFlags().String("base-url", "https://ingestion.acceleratedevops.net", "The API base url")
	pushCmd.MarkFlagRequired("base-url")
	pushCmd.PersistentFlags().String("system", "", "The system the data refers to (e.g.: webapp, backend)")
}
