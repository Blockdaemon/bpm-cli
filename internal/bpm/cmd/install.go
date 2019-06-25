package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/tasks"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <plugin name> [<version>]",
	Short: "Installs or upgrades a plugin",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]
		version := ""
		if len(args) > 1 {
			version = args[1]
		}

		output, err := tasks.Install(apiKey, baseDir, pluginURL, pluginName, version)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	addAPIKeyFlag(installCmd)
}
