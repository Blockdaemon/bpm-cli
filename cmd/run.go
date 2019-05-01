package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

var runCmd = &cobra.Command{
	Use:   "run <plugin>",
	Short: "Run an installed plugin",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		upgradable, err := tasks.CheckPluginUpgradable(baseDir, pluginName)
		if err != nil {
			return err
		}
		if upgradable {
			fmt.Printf("Please upgrade the plugin first by running: runner install %s\n", pluginName)
			return nil
		}

		return tasks.RunPlugin(baseDir, pluginName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
