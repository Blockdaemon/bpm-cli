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
		if err := tasks.DownloadVersionInfo(apiKey, pluginURL, baseDir); err != nil {
			return err
		}

		upgradable, err := tasks.CheckRunnerUpgradable(baseDir, runnerVersion)
		if err != nil {
			return err
		}
		if upgradable {
			fmt.Println("A new version of the runner is available, please upgrade!")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	addAPIKeyFlag(refreshCmd)
}
