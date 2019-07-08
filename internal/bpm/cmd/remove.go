package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/tasks"
)

var removeCmd = &cobra.Command{
	Use:   "remove <plugin>",
	Short: "Removes a running blockchain client. Data and configuration will not be removed.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		output, err := tasks.Remove(apiKey, baseDir, pluginURL, pluginName, runnerVersion)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	addAPIKeyFlag(removeCmd)
}
