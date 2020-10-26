package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

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
	parsedBaseURL, err := url.Parse(baseURL)

	if err != nil {
		fmt.Printf("Invalid base URL: %s", baseURL)
		os.Exit(1)
	}

	endpointURL, _ := parsedBaseURL.Parse("/v1/ingestion/cli")

	err = addQueryParameter(endpointURL, "api_key", apiKey)

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

func addQueryParameter(u *url.URL, key string, value string) error {
	q, err := url.ParseQuery(u.RawQuery)

	if err != nil {
		return err
	}

	q.Add(key, value)
	u.RawQuery = q.Encode()
	return nil
}
