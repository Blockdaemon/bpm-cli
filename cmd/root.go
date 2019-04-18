package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

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
	cobra.OnInitialize(tasks.CheckVersion)
	// pf := rootCmd.PersistentFlags()
	// rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "The API key from the Blockdaemon dashboard [REQUIRED]")
	// cobra.MarkFlagRequired(pf, "api-key")
}
