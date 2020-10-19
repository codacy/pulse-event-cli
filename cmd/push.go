package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// CredentialsString authenticates the user
var CredentialsString string

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
	pushCmd.PersistentFlags().StringVar(&CredentialsString, "credentials", "", "Cedentials to authenticate the user")
	pushCmd.MarkFlagRequired("credentials")
}

// GetCredentials parses credentials
func GetCredentials() (CredentialsType, []byte) {
	var credentialsBytes []byte
	var credentials CredentialsType
	credentialsBytes, err := base64.StdEncoding.DecodeString(CredentialsString)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(credentialsBytes, &credentials)

	return credentials, credentialsBytes
}

// CredentialsType authenticates the user
type CredentialsType struct {
	ProjectID string `json:"project_id"`
	DataSet   string `json:"data_set"`
}
