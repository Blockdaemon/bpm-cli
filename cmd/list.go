package cmd

import (
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/runner/tasks"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available and installed blockchain protocols",
	RunE: func(cmd *cobra.Command, args []string) error {
		versionInfoExists, err := tasks.CheckVersionInfoExists(baseDir)
		if err != nil {
			return err
		}

		if !versionInfoExists {
			fmt.Println(VERSION_INFO_MISSING)
			return nil
		}

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
