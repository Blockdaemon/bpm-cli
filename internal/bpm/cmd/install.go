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
		pluginVersion := ""
		if len(args) > 1 {
			pluginVersion = args[1]
		}

		output, err := tasks.Install(baseDir, pluginURL, pluginName, pluginVersion, runnerVersion)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
