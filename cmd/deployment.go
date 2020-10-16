package cmd

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/spf13/cobra"
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

		ctx := context.Background()
		client, err := bigquery.NewClient(ctx, "pulse-poc-1")
		if err != nil {
			fmt.Println(err)
		}

		// TODO: Remove after we do not need this example for testing
		// q := client.Query(`SELECT * FROM ` + "`pulse-poc-1.rodrigopessoa.deployments`" + ` LIMIT 1000`)
		// it, err := q.Read(ctx)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// for {
		// 	var values []bigquery.Value
		// 	err := it.Next(&values)
		// 	if err == iterator.Done {
		// 		break
		// 	}
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	fmt.Println(values)
		// }

		ins := client.Dataset("rodrigopessoa").Table("deployments").Inserter()
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

// BigQuery Deployments table schema
type deployment struct {
	Source      string    `bigquery:"source"`
	DeployID    string    `bigquery:"deploy_id"`
	TimeCreated time.Time `bigquery:"time_created"`
	Changes     []string  `bigquery:"changes"`
}

// Save method for the Deployment entity
func (deployment *deployment) save() (map[string]bigquery.Value, string, error) {
	return map[string]bigquery.Value{
		"source":       deployment.Source,
		"deploy_id":    deployment.DeployID,
		"time_created": deployment.TimeCreated,
		"changes":      deployment.Changes,
	}, "", nil
}
