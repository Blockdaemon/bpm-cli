package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/models"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <plugin name> [<version>]",
	Short: "Installs or upgrades a plugin",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		versionInfoExists, err := models.CheckVersionInfoExists(baseDir)
		if err != nil {
			return err
		}

		if !versionInfoExists {
			fmt.Println(VERSION_INFO_MISSING)
			return nil
		}

		plugin, err := models.LoadPlugin(baseDir, pluginURL, pluginName)
		if err != nil {
			return err
		}

		if len(args) > 1 {
			version := args[1]

			return plugin.InstallVersion(apiKey, version)
		}

		return plugin.InstallLatest(apiKey)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	addAPIKeyFlag(installCmd)
}
