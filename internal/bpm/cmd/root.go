package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pluginURL string
var baseDir string
var runnerVersion string
var apiKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "runner",
	Short: "The Blockdaemon runner manages Blockchain nodes in your own infrastructure",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	pf := rootCmd.PersistentFlags()
	pf.StringVar(&pluginURL, "plugin-url", "https://runner-test.sfo2.digitaloceanspaces.com/", "The URL used to download the plugins")
	pf.StringVar(&baseDir, "base-dir", "~/.blockdaemon/", "The directory in which plugins and configuration is stored")

	if runnerVersion == "" {
		runnerVersion = "development"
	}
}

// addAPIKeyFlag adds the flag "--api-key"
//
// Most commands have this flag but some don't. Therefore we need to specify
// it for each command that supports it.
func addAPIKeyFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&apiKey, "api-key", "", "The API key from the Blockdaemon dashboard")
	if err := cmd.MarkFlagRequired("api-key"); err != nil {
		panic(err) // Not much we can do here
	}

}
