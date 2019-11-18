package plugin

import (
	"fmt"
	"strings"

	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) Info(pluginName string) error {
	client := pbr.New(p.RegistryURL)

	versions, err := client.ListVersions(p.RuntimeOS, pluginName)
	if err != nil {
		return err
	}

	parameterOptions, err := p.getParameterOptions(pluginName)
	if err != nil {
		return err
	}

	fmt.Printf("Name:         %s\n", versions[0].Package.Name)
	fmt.Printf("Description:  %s\n", versions[0].Package.Description)
	fmt.Printf("Protocol:     %s\n", strings.Join(parameterOptions.Protocol, ", "))
	fmt.Printf("Network:      %s\n", strings.Join(parameterOptions.Network, ", "))
	fmt.Printf("Network Type: %s\n", strings.Join(parameterOptions.NetworkType, ", "))
	fmt.Printf("Subtype:      %s\n", strings.Join(parameterOptions.Subtype, ", "))
	prefix := "Versions:     "
	for ix, version := range versions {
		fmt.Printf("%s%s\n", prefix, version.Version)

		if ix == 0 {
			prefix = "              "
		}
	}

	return nil
}
