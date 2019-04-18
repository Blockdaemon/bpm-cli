package cmd

import (
	"gitlab.com/Blockdaemon/runner/tasks"

	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh <api-key>",
	Short: "Download the plugin list",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := args[0]
		return tasks.DownloadPluginList(apiKey)
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
