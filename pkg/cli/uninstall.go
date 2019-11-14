package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
)

func newUninstallCmd(c *command, runtimeOS string) *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall <package>",
		Short: "Uninstall a package. Data and configuration will not be removed.",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := args[0]

			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: runtimeOS,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			output, err := cmdContext.Uninstall(pluginName)
			fmt.Println(output)
			return err
		}),
	}
}
