package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/internal/bpm/tasks"
)

var purge bool

var removeCmd = &cobra.Command{
	Use:   "remove <plugin>",
	Short: "Removes a running blockchain client. Data and configuration will not be removed.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		output, err := tasks.Remove(baseDir, pluginURL, pluginName, runnerVersion, purge)
		if err != nil {
			return err
		}

		fmt.Println(output)
		return nil
	},
}

func init() {
	removeCmd.Flags().BoolVar(&purge, "purge", false, "Purge all data and configuration files")
	rootCmd.AddCommand(removeCmd)
}
