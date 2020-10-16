package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var identifier string
var timestamp int64
var source string

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Push deployment event",
	Long:  `Pushes a deployment event to the Pulse service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Pushing deployment event with identifier ", identifier, ", timestamp ", time.Unix(timestamp, 0), ", source ", source, " and changes ", args, "\n")

		credentials, credentialsBytes := getCredentials()

		ctx := context.Background()
		clientOptions := option.WithCredentialsJSON(credentialsBytes)
		client, err := bigquery.NewClient(ctx, credentials.ProjectID, clientOptions)
		if err != nil {
			fmt.Println(err)
		}

		ins := client.Dataset(credentials.DataSet).Table("deployments").Inserter()
		items := []*deployment{{Source: source, DeployID: identifier, TimeCreated: time.Unix(timestamp, 0), Changes: args}}
		if err := ins.Put(ctx, items); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	pushCmd.AddCommand(deploymentCmd)
	deploymentCmd.Flags().StringVar(&identifier, "identifier", "", "Deployment identifer (e.g.: commit sha)")
	deploymentCmd.MarkFlagRequired("identifier")
	deploymentCmd.Flags().Int64Var(&timestamp, "timestamp", 0, "Deployment timestamp (e.g.: 1602253523)")
	deploymentCmd.MarkFlagRequired("timestamp")
	deploymentCmd.Flags().StringVar(&source, "source", "cli", "Deployment source (e.g.: cli, git, GitHub)")
}

func getCredentials() (CredentialsType, []byte) {
	var credentialsBytes []byte
	var credentials CredentialsType
	credentialsBytes, err := base64.StdEncoding.DecodeString(CredentialsString)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(credentialsBytes, &credentials)

	return credentials, credentialsBytes
}

// BigQuery Deployments table schema
type deployment struct {
	Source      string    `bigquery:"source"`
	DeployID    string    `bigquery:"deploy_id"`
	TimeCreated time.Time `bigquery:"time_created"`
	Changes     []string  `bigquery:"changes"`
}

// CredentialsType authenticates the user
type CredentialsType struct {
	ProjectID string `json:"project_id"`
	DataSet   string `json:"data_set"`
}
