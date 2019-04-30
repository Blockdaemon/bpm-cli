package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install <plugin name> <api-key>",
	Short: "Installs or upgrades a plugin",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]
		apiKey := args[1]

		return tasks.InstallPlugin(baseDir, pluginURL, apiKey, pluginName)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
