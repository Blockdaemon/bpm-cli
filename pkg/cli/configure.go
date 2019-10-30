package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newConfigureCmd(c *command, runtimeOS string) *cobra.Command {
	var fields []string
	var skipUpgradeCheck bool

	cmd := &cobra.Command{
		Use:   "configure <package>",
		Short: "Configure a new blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := strings.ToLower(args[0])

			output, err := plugin.Configure(pluginName, homeDir, m, runtimeOS, c.registry, fields, skipUpgradeCheck, c.debug)
			if err != nil {
				return err
			}

			fmt.Println(output)
			return nil
		}),
	}

	cmd.Flags().StringSliceVar(&fields, "field", []string{}, "Custom fields to add to node.json")
	cmd.Flags().BoolVar(&skipUpgradeCheck, "skip-upgrade-check", false, "Skip checking whether a new version of the package is available")

	return cmd
}
