package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/codacy/event-cli/cmd"
	"github.com/spf13/cobra"
)

var apiKey string
var baseURL string
var system string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:              "push",
	Short:            "Push events",
	Long:             `Pushes all the different events necessary to generate your metrics in the Pulse service.`,
	TraverseChildren: true,
	Run:              func(cmd *cobra.Command, args []string) {},
}

func init() {
	cmd.RootCmd.AddCommand(pushCmd)
	pushCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "The API key to authenticate the organization/system")
	pushCmd.MarkFlagRequired("api-key")
	pushCmd.PersistentFlags().StringVar(&baseURL, "base-url", "https://ingestion.acceleratedevops.net", "The API base url")
	pushCmd.MarkFlagRequired("base-url")
	pushCmd.PersistentFlags().StringVar(&system, "system", "", "The system the data refers to (e.g.: webapp, backend)")
}

// createEvent creates the events in the data store
func createEvent(json []byte) {
	parsedBaseURL, err := url.Parse(baseURL)

	if err != nil {
		fmt.Printf("Invalid base URL: %s", baseURL)
		os.Exit(1)
	}

	endpointURL, _ := parsedBaseURL.Parse("/v1/ingestion/cli")

	queryParameters := map[string]string{"api_key": apiKey}

	if system != "" {
		queryParameters["system"] = system
	}

	err = addQueryParameters(endpointURL, queryParameters)

	if err != nil {
		fmt.Printf("Invalid query parameters in base URL: %s", baseURL)
		os.Exit(1)
	}

	resp, err := http.Post(endpointURL.String(), "application/json", bytes.NewBuffer(json))

	if err != nil {
		fmt.Print("Unexpected error pushing event.\n")
		os.Exit(1)
	}

	statusOk := resp.StatusCode >= 200 && resp.StatusCode <= 299

	if !statusOk {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if err != nil {
			bodyStr = "Failed to read http response body"
		}

		fmt.Printf("Failed to push event, status code %s.\n%s\n", resp.Status, bodyStr)
		os.Exit(1)
	}

}

func addQueryParameters(u *url.URL, parameters map[string]string) error {
	q, err := url.ParseQuery(u.RawQuery)

	if err != nil {
		return err
	}

	for key, value := range parameters {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	return nil
}
