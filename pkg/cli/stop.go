package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newStopCmd(c *command) *cobra.Command {
	var purge bool

	cmd := &cobra.Command{
		Use:   "stop <package>",
		Short: "Removes a running blockchain client. Data and configuration will not be removed.",
		Args:  cobra.MinimumNArgs(2),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := strings.ToLower(args[0])
			id := args[1]

			// Check if plugin is installed
			if _, ok := m.Plugins[pluginName]; !ok {
				fmt.Printf("The package %q is currently not installed.\n", pluginName)
				return nil
			}

			// Remove plugin
			if err := plugin.Stop(homeDir, pluginName, id, purge, c.debug); err != nil {
				return err
			}

			fmt.Println("\nThe directory node has not been removed. It may contain private keys that protect sensitive information/funds. Please remove it manually if it is not needed anymore.")

			return nil
		}),
	}

	cmd.Flags().BoolVar(&purge, "purge", false, "Purge all data and configuration files. Secrets (e.g. private keys) will not be removed because they may protect sensitive information/funds.")

	return cmd
}
