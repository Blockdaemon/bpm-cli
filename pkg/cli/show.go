package cli

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func newShowCmd(c *command, runtimeOS string) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show <resource>",
		Short: "Print a resource to stdout",
	}

	showConfigCmd := &cobra.Command{
		Use:   "config <id>",
		Short: "Display config files for a node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			id := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: runtimeOS,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			output, err := cmdContext.ShowConfig(id)
			fmt.Println(output)
			return err
		}),
	}

	showNodeCmd := &cobra.Command{
		Use:   "node <id>",
		Short: "Display the node.json config",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			id := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: runtimeOS,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			output, err := cmdContext.ShowNode(id)
			fmt.Println(output)
			return err
		}),
	}

	showCmd.AddCommand(showConfigCmd)
	showCmd.AddCommand(showNodeCmd)

	return showCmd
}
