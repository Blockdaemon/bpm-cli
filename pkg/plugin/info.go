package plugin

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
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

	buf.WriteString("Name:         " + versions[0].Package.Name + "\n")
	buf.WriteString("Description:  " + versions[0].Package.Description + "\n")
	buf.WriteString("Protocol:     " + strings.Join(parameterOptions.Protocol, ", ") + "\n")
	buf.WriteString("Network:      " + strings.Join(parameterOptions.Network, ", ") + "\n")
	buf.WriteString("Network Type: " + strings.Join(parameterOptions.NetworkType, ", ") + "\n")
	buf.WriteString("Subtype:      " + strings.Join(parameterOptions.Subtype, ", ") + "\n")
	prefix := "Versions:     "
	for ix, version := range versions {
		buf.WriteString(prefix + version.Version + "\n")

		if ix == 0 {
			prefix = "              "
		}
	}

	return buf.String(), nil
}
