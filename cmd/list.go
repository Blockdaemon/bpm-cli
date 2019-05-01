package cmd

import (
	"os"

	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available and installed blockchain protocols",
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginListItems, err := tasks.ListPlugins(baseDir)
		if err != nil {
			return err
		}

		tableprinter.Print(os.Stdout, pluginListItems)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
