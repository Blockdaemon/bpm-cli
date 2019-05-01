package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <plugin name> <api-key> [<version>]",
	Short: "Installs or upgrades a plugin",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]
		apiKey := args[1]

		if len(args) > 2 {
			version := args[2]
			return tasks.InstallPluginVersion(baseDir, pluginURL, apiKey, pluginName, version)
		}

		return tasks.InstallPluginLatest(baseDir, pluginURL, apiKey, pluginName)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
