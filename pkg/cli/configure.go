package cli

import (
	"fmt"

	sdkplugin "github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/Blockdaemon/bpm/pkg/plugin"
	"github.com/spf13/cobra"
)

func newConfigureCmd(cmdContext plugin.PluginCmdContext) *cobra.Command {
	var skipUpgradeCheck bool

	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure a new blockchain node",
	}

	for name, meta := range cmdContext.Manifest.Plugins {
		pluginCmd := &cobra.Command{
			Use:   name,
			Short: fmt.Sprintf("Configure a new blockchain node using the %q package", name),
			RunE: func(cmd *cobra.Command, args []string) error {
				// Read dynamic parameters
				strParameters := map[string]string{}
				boolParameters := map[string]bool{}

				for _, parameter := range meta.Parameters {
					if parameter.Type == sdkplugin.ParameterTypeString {
						value, err := cmd.Flags().GetString(parameter.Name)
						if err != nil {
							return err
						}
						strParameters[parameter.Name] = value
					} else {
						value, err := cmd.Flags().GetBool(parameter.Name)
						if err != nil {
							return err
						}
						boolParameters[parameter.Name] = value
					}
				}

				return cmdContext.Configure(name, strParameters, boolParameters, skipUpgradeCheck)
			},
		}
		pluginCmd.Flags().BoolVar(&skipUpgradeCheck, "skip-upgrade-check", false, "Skip checking whether a new version of the package is available")

		// Add dynamic configuration parameters
		for _, parameter := range meta.Parameters {
			pluginCmd.Flags().String(parameter.Name, parameter.Default, parameter.Description)
			if parameter.Mandatory {
				pluginCmd.MarkFlagRequired(parameter.Name)
			}
		}

		cmd.AddCommand(pluginCmd)
	}

	return cmd
}
