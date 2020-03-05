package cli

import (
	"fmt"

	sdkplugin "github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/Blockdaemon/bpm/pkg/command"
	"github.com/spf13/cobra"
)

func newConfigureCmd(cmdContext command.CmdContext) *cobra.Command {
	var skipUpgradeCheck bool

	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure a new blockchain node",
		RunE: func(cmd *cobra.Command, args []string) error {
			// The configure command doesn't provide any functionality by itself.
			// It has two purposes:
			//
			// 1. Act as parent command for all `configure <plugin>` commands
			// 2. If no plugin is installed yet, provide a useful help message
			pluginName := ""

			if len(args) > 0 {
				pluginName = args[0]
			}

			return cmdContext.ConfigureHelp(pluginName)
		},
	}

	// Create `configure <plugin>` command for evert installed plugin. With parameters specific
	// to that pugin
	for name, meta := range cmdContext.Manifest.Plugins {
		// copy to break closure
		nameCopy := name
		metaCopy := meta

		pluginCmd := &cobra.Command{
			Use:   nameCopy,
			Short: fmt.Sprintf("Configure a new blockchain node using the %q package", nameCopy),
			RunE: func(cmd *cobra.Command, args []string) error {
				// Read dynamic parameters
				strParameters := map[string]string{}
				boolParameters := map[string]bool{}

				for _, parameter := range metaCopy.Parameters {
					if parameter.Type == sdkplugin.ParameterTypeString {
						value, err := cmd.Flags().GetString(parameter.Name)
						if err != nil {
							return fmt.Errorf("Cannot read parameter %q: ", err)
						}
						strParameters[parameter.Name] = value
					} else {
						value, err := cmd.Flags().GetBool(parameter.Name)
						if err != nil {
							return fmt.Errorf("Cannot read parameter %q: %s", parameter.Name, err)
						}
						boolParameters[parameter.Name] = value
					}
				}

				return cmdContext.Configure(nameCopy, strParameters, boolParameters, skipUpgradeCheck)
			},
		}
		pluginCmd.Flags().BoolVar(&skipUpgradeCheck, "skip-upgrade-check", false, "Skip checking whether a new version of the package is available")

		// Add dynamic configuration parameters
		for _, parameter := range metaCopy.Parameters {
			if parameter.Type == sdkplugin.ParameterTypeString {
				pluginCmd.Flags().String(parameter.Name, parameter.Default, parameter.Description)
			} else {
				defaultValue := false
				if parameter.Default == "true" || parameter.Default == "yes" || parameter.Default == "on" {
					defaultValue = true
				}
				pluginCmd.Flags().Bool(parameter.Name, defaultValue, parameter.Description)
			}

			if parameter.Mandatory {
				if err := pluginCmd.MarkFlagRequired(parameter.Name); err != nil {
					exitWithError(err, pluginCmd)
				}
			}
		}

		cmd.AddCommand(pluginCmd)
	}

	return cmd
}
