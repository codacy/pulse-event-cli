package ingestion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PulseIngestionAPIClient a client to call Pulse's ingestion API
type PulseIngestionAPIClient struct {
	baseURL     *url.URL
	apiKey      string
	system      string
	environment *string
}

// NewPulseIngestionAPIClient creates an API client validating the provided baseURL
func NewPulseIngestionAPIClient(baseURL string, apiKey string, system string, environment *string) (*PulseIngestionAPIClient, error) {
	parsedBaseURL, err := url.Parse(baseURL)

	if err != nil {
		err := fmt.Errorf("invalid base URL: %s", baseURL)
		return nil, err
	}

	return &PulseIngestionAPIClient{baseURL: parsedBaseURL, apiKey: apiKey, system: system, environment: environment}, nil
}

// CreateEvent creates the events in the data store
func (client *PulseIngestionAPIClient) CreateEvent(event interface{}) error {
	json, _ := json.Marshal(event)

	endpointURL, _ := client.baseURL.Parse("/v1/ingestion/cli")

	queryParameters := map[string]string{"api_key": client.apiKey}

	if client.system != "" {
		queryParameters["system"] = client.system
	}

	err := addQueryParameters(endpointURL, queryParameters)

	if err != nil {
		return fmt.Errorf("invalid query parameters in base URL: %s", client.baseURL)
	}

	req, err := http.NewRequest("POST", endpointURL.String(), bytes.NewBuffer(json))
	if err != nil {
		return fmt.Errorf("unexpected error preparing push event request:\n%v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if client.environment != nil {
		req.Header.Set("Environment", *client.environment)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("unexpected error pushing event:\n%v", err)
	}

	statusOk := resp.StatusCode >= 200 && resp.StatusCode <= 299

	if !statusOk {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)

		if err != nil {
			bodyStr = "Failed to read http response body"
		}

		return fmt.Errorf("failed to push event, status code %s.\n%s", resp.Status, bodyStr)
	}

	return nil
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
