package push

import (
	"github.com/codacy/event-cli/cmd"
	"github.com/codacy/event-cli/internal/environment"
	"github.com/codacy/event-cli/pkg/ingestion"
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
	PushCmd.PersistentFlags().String("api-key", "", "The API key to authenticate the organization/system")
	PushCmd.MarkPersistentFlagRequired("api-key")
	PushCmd.PersistentFlags().String("base-url", "https://ingestion.pulse.codacy.com", "The API base url")
	PushCmd.PersistentFlags().String("system", "", "The system the data refers to (e.g.: webapp, backend)")
}

// GetAPIClient returns an API client created from the push command flags
func GetAPIClient(cmd *cobra.Command) (*ingestion.PulseIngestionAPIClient, error) {
	apiKey, _ := cmd.Flags().GetString("api-key")
	baseURL, _ := cmd.Flags().GetString("base-url")
	system, _ := cmd.Flags().GetString("system")

	environment := environment.GetEnvironmentName()

	return ingestion.NewPulseIngestionAPIClient(baseURL, apiKey, system, environment)
}
