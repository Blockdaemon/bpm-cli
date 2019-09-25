package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
)

func newUninstallCmd(c *command) *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall <package>",
		Short: "Uninstall a package. Data and configuration will not be removed.",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := strings.ToLower(args[0])

			// Check if plugin is installed
			if _, ok := m.Plugins[pluginName]; !ok {
				fmt.Printf("The package %q is currently not installed.\n", pluginName)
				return nil
			}

			// Delete the plugin
			if err := config.DeleteFile(
				config.PluginsDir(homeDir),
				pluginName,
			); err != nil {
				return err
			}

			// Remove plugin from manifest
			delete(m.Plugins, pluginName)

			if err := config.WriteFile(
				homeDir,
				config.ManifestFilename,
				m,
			); err != nil {
				return err
			}

			fmt.Printf("The package %q has been uninstalled.\n", pluginName)

			return nil
		}),
	}
}
