package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Download the plugin list and check if the runner is up-to-date",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := tasks.Refresh(apiKey, baseDir, pluginURL, runnerVersion)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	addAPIKeyFlag(refreshCmd)
}
