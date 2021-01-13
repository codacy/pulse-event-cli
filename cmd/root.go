package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/codacy/pulse-event-cli/internal/build"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var versionTemplate = fmt.Sprintf("pulse-event-cli version %v (%v)\nhttps://github.com/codacy/pulse-event-cli/releases/latest\n", build.Version, build.Date)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pulse-event-cli",
	Short: "Pulse command line interface",
	Long: `This command line interface is a client for the Pulse service.
							For more information see https://docs.pulse.codacy.com`,
	Version:          build.Version,
	TraverseChildren: true,
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version",
	Long:  `Version of the CLI binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionTemplate)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "configuration file (the default is $HOME/.event-cli.yaml)")
	RootCmd.SetVersionTemplate(versionTemplate)
	RootCmd.AddCommand(versionCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".event-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".event-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
