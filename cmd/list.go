package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available and installed blockchain protocols",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return tasks.DownloadPluginList(apiKey)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
