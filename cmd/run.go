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

		versionInfoExists, err := tasks.CheckVersionInfoExists(baseDir)
		if err != nil {
			return err
		}

		if !versionInfoExists {
			fmt.Println(VERSION_INFO_MISSING)
			return nil
		}

		plugin, err := tasks.LoadPlugin(baseDir, pluginURL, pluginName)
		if err != nil {
			return err
		}

		return plugin.RunPlugin()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	addAPIKeyFlag(runCmd)
}
