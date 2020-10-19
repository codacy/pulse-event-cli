package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// credentialsString authenticates the user
var credentialsString string

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
	pushCmd.PersistentFlags().StringVar(&credentialsString, "credentials", "", "Cedentials to authenticate the user")
	pushCmd.MarkFlagRequired("credentials")
}

// createEvent creates the events in the data store
func createEvent(tableName string, events interface{}) {
	credentials, credentialsBytes := getCredentials()

	ctx := context.Background()
	clientOptions := option.WithCredentialsJSON(credentialsBytes)
	client, err := bigquery.NewClient(ctx, credentials.ProjectID, clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	ins := client.Dataset(credentials.DataSet).Table(tableName).Inserter()
	if err := ins.Put(ctx, events); err != nil {
		fmt.Println(err)
	}
}

// getCredentials parses credentials
func getCredentials() (credentialsType, []byte) {
	var credentialsBytes []byte
	var credentials credentialsType
	credentialsBytes, err := base64.StdEncoding.DecodeString(credentialsString)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(credentialsBytes, &credentials)

	return credentials, credentialsBytes
}

// credentialsType authenticates the user
type credentialsType struct {
	ProjectID string `json:"project_id"`
	DataSet   string `json:"data_set"`
}
