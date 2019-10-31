package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/plugin"
)

func newConfigureCmd(c *command, runtimeOS string) *cobra.Command {
	var (
		skipUpgradeCheck bool
		network          string
		networkType      string
		protocol         string
		subtype          string
	)

	cmd := &cobra.Command{
		Use:   "configure <package>",
		Short: "Configure a new blockchain node",
		Args:  cobra.MinimumNArgs(1),
		RunE: c.Wrap(func(homeDir string, m config.Manifest, args []string) error {
			pluginName := strings.ToLower(args[0])

			output, err := plugin.Configure(pluginName, homeDir, m, runtimeOS, c.registry, network, networkType, protocol, subtype, skipUpgradeCheck, c.debug)
			if err != nil {
				return err
			}

			fmt.Println(output)
			return nil
		}),
	}

	cmd.Flags().BoolVar(&skipUpgradeCheck, "skip-upgrade-check", false, "Skip checking whether a new version of the package is available")

	// For simplicty sake the parameters are hardcoded here. In the future we may want to add them dynamically. This would allow plugins to specify
	// arbitrary parameters.
	cmd.Flags().StringVar(&network, "network", "", "The network this node should connect to (Examples: 'mainnet', 'testnet', 'goerli', ...")
	cmd.Flags().StringVar(&networkType, "network-type", "", "The network-type specifies whether this is a public or private network")
	cmd.Flags().StringVar(&protocol, "protocol", "", "The protocol this node should use (Examples: 'ethereum', 'polkadot', ...")
	cmd.Flags().StringVar(&subtype, "subtype", "", "The subtype specifies how this node should be configured (Examples: 'validator', 'watcher', 'archive', ...")

	return cmd
}
