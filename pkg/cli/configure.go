package cli

import (
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
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
			pluginName := args[0]

			// TODO: Why do we have three ways of passing down variables?
			cmdContext := plugin.PluginCmdContext{
				HomeDir: homeDir,
				Manifest: m,
				RuntimeOS: runtimeOS,
				RegistryURL: c.registry,
				Debug: c.debug,
			}

			return cmdContext.Configure(pluginName, network, networkType, protocol, subtype, skipUpgradeCheck)
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
