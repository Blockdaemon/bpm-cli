package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiKey string
var pluginURL string
var baseDir string
var runnerVersion string

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
	// cobra.OnInitialize(tasks.CheckVersion)

	rootCmd.PersistentFlags().StringVar(&pluginURL, "plugin-url", "https://runner-test.sfo2.digitaloceanspaces.com/", "The URL used to download the plugins")
	rootCmd.PersistentFlags().StringVar(&baseDir, "base-dir", "~/.blockdaemon/", "The directory in which plugins and configuration is stored")

	if runnerVersion == "" {
		runnerVersion = "development"
	}
}
