package plugin

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/rs/xid"
)

func (p *PluginCmdContext) Configure(pluginName string, networkParam string, networkTypeParam string, protocolParam string, subtypeParam string, skipUpgradeCheck bool) error {
	// Generate instance id
	id := xid.New().String()

	if !p.isInstalled(pluginName) {
		return fmt.Errorf("The package %q is currently not installed.", pluginName)
	}

	// Check if plugin is using the latest version
	if !skipUpgradeCheck {
		needsUpgrade, latestVersion, err := p.needsUpgrade(pluginName)

		if err != nil {
			return err
		}

		if needsUpgrade {
			return fmt.Errorf("A new version of package %q is available. Please install using \"bpm install %s %s\".\n", pluginName, pluginName, latestVersion)
		}
	}

	parameterOptions, err := p.getParameterOptions(pluginName)
	if err != nil {
		return err
	}

	// Validate parameters
	network, err := validateParameter("network", networkParam, parameterOptions.Network)
	if err != nil {
		return err
	}
	protocol, err := validateParameter("protocol", protocolParam, parameterOptions.Protocol)
	if err != nil {
		return err
	}
	networkType, err := validateParameter("network-type", networkTypeParam, parameterOptions.NetworkType)
	if err != nil {
		return err
	}
	subtype, err := validateParameter("subtype", subtypeParam, parameterOptions.Subtype)
	if err != nil {
		return err
	}

	// Create node config
	n := node.New(config.NodesDir(p.HomeDir), id)
	n.Environment = network
	n.Protocol = protocol
	n.Subtype = subtype
	n.Version = p.getInstalledVersion(pluginName)
	n.NetworkType = networkType

	// Only temporary until we find a better solution to distribute the certs
	if config.FileExists(p.HomeDir, "beats") {
		n.Collection.Host = "dev-1.logstash.blockdaemon.com:5044"
		n.Collection.Cert = "~/.bpm/beats/beat.crt"
		n.Collection.CA = "~/.bpm/beats/ca.crt"
		n.Collection.Key = "~/.bpm/beats/beat.key"
	} else {
		fmt.Printf("\nNo credentials found in %q, skipping configuration of Blockdaemon monitoring. Please configure your own monitoring in the node configuration files.\n\n", filepath.Join(p.HomeDir, "beats"))
	}

	if err := n.Save(); err != nil {
		return err
	}

	// Secrets
	err = p.execPrintNodeCommand(n, "create-secrets")
	if err != nil {
		return err
	}

	// Config
	err = p.execPrintNodeCommand(n, "create-configurations")
	if err != nil {
		return err
	}

	fmt.Printf("\nNode with id %q has been initialized.\n\nTo change the configuration, modify the files here:\n    %s\nTo start the node, run:\n    bpm start %s\nTo see the status of configured nodes, run:\n    bpm status\n", id, n.ConfigsDirectory(), id)

	return nil
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
