package plugin

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/rs/xid"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/manager"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
	"gitlab.com/Blockdaemon/bpm/pkg/version"
	"gopkg.in/yaml.v2"
)

// TODO: Too many parameters, need to clean this up
func Configure(pluginName string, homeDir string, m config.Manifest, runtimeOS string, registry string, networkParam string, networkTypeParam string, protocolParam string, subtypeParam string, skipUpgradeCheck bool, debug bool) (string, error) {
	// Generate instance id
	id := xid.New().String()

	// Check if plugin is installed
	p, ok := m.Plugins[pluginName]
	if !ok {
		return fmt.Sprintf("The package %q is currently not installed.\n", pluginName), nil
	}

	// Check if plugin is using the latest version
	if !skipUpgradeCheck {
		client := pbr.New(registry)
		packageVersion, err := client.GetLatestPackageVersion(pluginName, runtimeOS)
		if err != nil {
			return "", err
		}
		latestVersion := packageVersion.Version

		needsUpgrade, err := version.NeedsUpgrade(p.Version, latestVersion)
		if err != nil {
			return "", err
		}

		if needsUpgrade {
			return fmt.Sprintf("A new version of package %q is available. Please install using \"bpm install %s %s\".\n", pluginName, pluginName, latestVersion), nil
		}
	}

	// TODO: Duplicated code in info.go - remove when we simplify the plugin cli interface
	// Prepare running the plugin
	pluginFilename := filepath.Join(config.PluginsDir(homeDir), pluginName)
	baseDirArgs := []string{"--base-dir", config.NodesDir(homeDir)}

	// Get parameter options
	configArgs := append([]string{"parameters", id}, baseDirArgs...)
	output, err := manager.ExecCmd(debug, pluginFilename, configArgs...)
	if err != nil {
		return "", err
	}

	parameterOptions := plugin.Parameters{}
	err = yaml.Unmarshal([]byte(output), &parameterOptions)
	if err != nil {
		return "", err
	}

	// Validate parameters
	network, err := validateParameter("network", networkParam, parameterOptions.Network)
	if err != nil {
		return "", err
	}
	protocol, err := validateParameter("protocol", protocolParam, parameterOptions.Protocol)
	if err != nil {
		return "", err
	}
	networkType, err := validateParameter("network-type", networkTypeParam, parameterOptions.NetworkType)
	if err != nil {
		return "", err
	}
	subtype, err := validateParameter("subtype", subtypeParam, parameterOptions.Subtype)
	if err != nil {
		return "", err
	}

	// Create node config
	n := node.New(config.NodesDir(homeDir), id)
	n.Environment = network
	n.Protocol = protocol
	n.Subtype = subtype
	n.Version = p.Version
	n.NetworkType = networkType

	// Only temporary until we find a better solution to distribute the certs
	if config.FileExists(homeDir, "beats") {
		n.Collection.Host = "dev-1.logstash.blockdaemon.com:5044"
		n.Collection.Cert = "~/.bpm/beats/beat.crt"
		n.Collection.CA = "~/.bpm/beats/ca.crt"
		n.Collection.Key = "~/.bpm/beats/beat.key"
	} else {
		fmt.Printf("No credentials found in %q, skipping configuration of Blockdaemon monitoring. Please configure your own monitoring in the node configuration files.\n\n", filepath.Join(homeDir, "beats"))
	}

	if err := n.Save(); err != nil {
		return "", err
	}

	// Secrets
	secretArgs := append([]string{"create-secrets", id}, baseDirArgs...)
	output, err = manager.ExecCmd(debug, pluginFilename, secretArgs...)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	// Config
	configArgs = append([]string{"create-configurations", id}, baseDirArgs...)
	output, err = manager.ExecCmd(debug, pluginFilename, configArgs...)
	if err != nil {
		return "", err
	}

	fmt.Println(output)

	return fmt.Sprintf("\nNode with id %q has been initialized.\n\nTo change the configuration, modify the files here:\n    %s\nTo start the node, run:\n    bpm start %s\nTo see the status of configured nodes, run:\n    bpm status\n", id, n.ConfigsDirectory(), id), nil
}

func validateParameter(name string, value string, options []string) (string, error) {
	if len(value) == 0 {
		return options[0], nil // default to first option
	}

	if !stringInSlice(value, options) {
		return "", fmt.Errorf("%s must be one of: %s", name, strings.Join(options, ", "))
	}

	return value, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
