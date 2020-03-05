package command

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/pbr"
)

// Info prints information about a particular plugin
func (p *CmdContext) Info(pluginName string) error {
	client := pbr.New(p.RegistryURL)

	versions, err := client.ListVersions(p.RuntimeOS, pluginName)
	if err != nil {
		return err
	}

	fmt.Printf("Name:         %s\n", versions[0].Package.Name)
	fmt.Printf("Description:  %s\n", versions[0].Package.Description)

	if p.isInstalled(pluginName) {
		meta, err := p.getMeta(pluginName)
		if err != nil {
			return err
		}

		for ix, parameter := range meta.Parameters {
			separator := ""
			appendStr := ""
			if parameter.Mandatory {
				appendStr = "mandatory"
				separator = ", "
			}
			if parameter.Default != "" {
				appendStr = fmt.Sprintf("%s%sdefault: %q", appendStr, separator, parameter.Default)
				separator = ", "
			}
			appendStr = fmt.Sprintf("%s%stype: %q", appendStr, separator, parameter.Type)

			if ix == 0 {
				fmt.Printf("Parameters:   %s - %s (%s)\n", parameter.Name, parameter.Description, appendStr)
			} else {
				fmt.Printf("              %s - %s (%s)\n", parameter.Name, parameter.Description, appendStr)
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
