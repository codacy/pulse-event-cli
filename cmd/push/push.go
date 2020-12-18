package push

import (
	"github.com/codacy/pulse-event-cli/cmd"
	"github.com/codacy/pulse-event-cli/internal/build"
	"github.com/codacy/pulse-event-cli/internal/environment"
	"github.com/codacy/pulse-event-cli/pkg/ingestion"
	"github.com/spf13/cobra"
)

var PushCmd = &cobra.Command{
	Use:              "push",
	Short:            "Push events",
	Long:             `Pushes all the different events necessary to generate your metrics in the Pulse service.`,
	TraverseChildren: true,
}

func init() {
	cmd.RootCmd.AddCommand(PushCmd)
	PushCmd.PersistentFlags().String("api-key", "", "API authentication key for the organization")
	PushCmd.MarkPersistentFlagRequired("api-key")
	PushCmd.PersistentFlags().String("base-url", "https://ingestion.pulse.codacy.com", "base URL of the API endpoint")
	PushCmd.PersistentFlags().String("system", "", "repository or component to associate with the event")
}

// GetAPIClient returns an API client created from the push command flags
func GetAPIClient(cmd *cobra.Command) (*ingestion.PulseIngestionAPIClient, error) {
	apiKey, _ := cmd.Flags().GetString("api-key")
	baseURL, _ := cmd.Flags().GetString("base-url")
	system, _ := cmd.Flags().GetString("system")

	environment := environment.GetEnvironmentName()
	cliVersion := build.Version

	return ingestion.NewPulseIngestionAPIClient(baseURL, apiKey, system, environment, cliVersion)
}
