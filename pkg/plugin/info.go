package plugin

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) Info(pluginName string) error {
	client := pbr.New(p.RegistryURL)

	versions, err := client.ListVersions(p.RuntimeOS, pluginName)
	if err != nil {
		return err
	}

	fmt.Printf("Name:         %s\n", versions[0].Package.Name)
	fmt.Printf("Description:  %s\n", versions[0].Package.Description)

	if p.isInstalled(pluginName) {
		parameterOptions, err := p.getParameterOptions(pluginName)
		if err != nil {
			return err
		}

		for ix, protocol := range parameterOptions.Protocol {
			if ix == 0 {
				fmt.Printf("Protocol:     %s (default)\n", protocol)
			} else {
				fmt.Printf("              %s\n", protocol)
			}
		}

		for ix, network := range parameterOptions.Network {
			if ix == 0 {
				fmt.Printf("Network:      %s (default)\n", network)
			} else {
				fmt.Printf("              %s\n", network)
			}
		}

		for ix, networkType := range parameterOptions.NetworkType {
			if ix == 0 {
				fmt.Printf("Network Type: %s (default)\n", networkType)
			} else {
				fmt.Printf("              %s\n", networkType)
			}
		}

		for ix, subType := range parameterOptions.Subtype {
			if ix == 0 {
				fmt.Printf("Subtype:      %s (default)\n", subType)
			} else {
				fmt.Printf("              %s\n", subType)
			}
		}
	}

	installedVersion := p.getInstalledVersion(pluginName)
	for ix, version := range versions {
		installedPlaceholder := ""
		if version.Version == installedVersion {
			installedPlaceholder = " (installed)"
		}

		if ix == 0 {
			fmt.Printf("Versions:     %s%s\n", version.Version, installedPlaceholder)
		} else {
			fmt.Printf("              %s\n", version.Version)
		}
	}

	return nil
}
