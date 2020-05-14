package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.blockdaemon.com/bpm/cli/pkg/command"
	sdkplugin "go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

func newConfigureCmd(cmdContext command.CmdContext) *cobra.Command {
	var skipUpgradeCheck bool
	var nodeName string

	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configure a blockchain node, creates a new node if it doesn't exist yet",
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
	for pluginName, meta := range cmdContext.Manifest.Plugins {
		// copy to break closure
		pluginNameCopy := pluginName
		metaCopy := meta

		pluginCmd := &cobra.Command{
			Use:   pluginNameCopy,
			Short: fmt.Sprintf("Configure a new blockchain node using the %q package", pluginNameCopy),
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

				return cmdContext.Configure(pluginNameCopy, nodeName, strParameters, boolParameters, skipUpgradeCheck)
			},
		}
		pluginCmd.Flags().BoolVar(&skipUpgradeCheck, "skip-upgrade-check", false, "Skip checking whether a new version of the package is available")
		pluginCmd.Flags().StringVar(&nodeName, "name", "", "The name of the node. Will be chosen automatically if not specified")

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
