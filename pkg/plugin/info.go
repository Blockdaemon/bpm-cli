package plugin

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
	"github.com/Blockdaemon/bpm/pkg/pbr"
	"gopkg.in/yaml.v2"
)

func Info(registry string, packageName string, os string, m config.Manifest, homeDir string, debug bool) (string, error) {
	client := pbr.New(registry)

	versions, err := client.ListVersions(os, packageName)
	if err != nil {
		return "", err
	}

	// TODO: Duplicated code in configure.go - remove when we simplify the plugin cli interface
	// Prepare running the plugin
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), packageName)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Get parameter options
	configArgs := append([]string{"parameters"}, baseDirArgs...)
	output, err := manager.ExecCmd(debug, pluginFilename, configArgs...)
	if err != nil {
		return "", err
	}

	parameterOptions := plugin.Parameters{}
	err = yaml.Unmarshal([]byte(output), &parameterOptions)
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
