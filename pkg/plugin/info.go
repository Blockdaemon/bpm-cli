package plugin

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) Info(pluginName string) (string, error) {
	client := pbr.New(p.RegistryURL)

	versions, err := client.ListVersions(p.RuntimeOS, pluginName)
	if err != nil {
		return "", err
	}

	parameterOptions, err := p.getParameterOptions(pluginName)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("Name:         %s\n", versions[0].Package.Name))
	buf.WriteString(fmt.Sprintf("Description:  %s\n", versions[0].Package.Description))
	buf.WriteString(fmt.Sprintf("Protocol:     %s\n", strings.Join(parameterOptions.Protocol, ", ")))
	buf.WriteString(fmt.Sprintf("Network:      %s\n", strings.Join(parameterOptions.Network, ", ")))
	buf.WriteString(fmt.Sprintf("Network Type: %s\n", strings.Join(parameterOptions.NetworkType, ", ")))
	buf.WriteString(fmt.Sprintf("Subtype:      %s\n", strings.Join(parameterOptions.Subtype, ", ")))
	prefix := "Versions:     "
	for ix, version := range versions {
		buf.WriteString(fmt.Sprintf("%s%s\n", prefix, version.Version))

		if ix == 0 {
			prefix = "              "
		}
	}

	return buf.String(), nil
}
