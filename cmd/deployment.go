package cmd

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var deploymentIdentifier string
var deploymentTimestamp int64
var deploymentSource string

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Push deployment event",
	Long:  `Pushes a deployment event to the Pulse service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Pushing deployment event with identifier ", deploymentIdentifier, ", timestamp ", time.Unix(deploymentTimestamp, 0), ", source ", deploymentSource, " and changes ", args, "\n")

		credentials, credentialsBytes := GetCredentials()

		ctx := context.Background()
		clientOptions := option.WithCredentialsJSON(credentialsBytes)
		client, err := bigquery.NewClient(ctx, credentials.ProjectID, clientOptions)
		if err != nil {
			fmt.Println(err)
		}

		ins := client.Dataset(credentials.DataSet).Table("deployments").Inserter()
		items := []*deployment{{Source: deploymentSource, DeployID: deploymentIdentifier, TimeCreated: time.Unix(deploymentTimestamp, 0), Changes: args}}
		if err := ins.Put(ctx, items); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	pushCmd.AddCommand(deploymentCmd)
	deploymentCmd.Flags().StringVar(&deploymentIdentifier, "identifier", "", "Deployment identifer (e.g.: commit sha)")
	deploymentCmd.MarkFlagRequired("identifier")
	deploymentCmd.Flags().Int64Var(&deploymentTimestamp, "timestamp", 0, "Deployment timestamp (e.g.: 1602253523)")
	deploymentCmd.MarkFlagRequired("timestamp")
	deploymentCmd.Flags().StringVar(&deploymentSource, "source", "cli", "Deployment source (e.g.: cli, git, GitHub)")
}

// BigQuery Deployments table schema
type deployment struct {
	Source      string    `bigquery:"source"`
	DeployID    string    `bigquery:"deploy_id"`
	TimeCreated time.Time `bigquery:"time_created"`
	Changes     []string  `bigquery:"changes"`
}
