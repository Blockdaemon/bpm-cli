package command

import (
	"fmt"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/fileutil"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/rs/xid"
)

// ConfigureHelp provides the logic for the `configure` command without parameters
//
// Since we cannot do much without a plugin as first parameter, it just prints help information
func (p *CmdContext) ConfigureHelp(pluginName string) error {
	if pluginName == "" {
		if !p.Manifest.HasPluginsInstalled() {
			return fmt.Errorf("cannot configure without an installed package")
		}

		return fmt.Errorf("no package specified. See `--help` for details")
	}

	if !p.isInstalled(pluginName) {
		return fmt.Errorf("the package %q is currently not installed", pluginName)
	}

	return nil
}

// Configure provides the logic for configuring a node using a particular plugin
func (p *CmdContext) Configure(pluginName string, strParameters map[string]string, boolParameters map[string]bool, skipUpgradeCheck bool) error {
	// Generate instance id
	id := xid.New().String()

	if !p.isInstalled(pluginName) {
		return fmt.Errorf("the package %q is currently not installed", pluginName)
	}

	// Check if plugin is using the latest version
	if !skipUpgradeCheck {
		needsUpgrade, err := p.needsUpgrade(pluginName)

		if err != nil {
			// During development it often happens that a plugin is being used that hasn't been uploaded
			// to the registry. Ergo, the upgrade check will fail. In order to speed up development we
			// we just tell the user about it but don't stop here
			fmt.Printf("Upgrade check failed: %s\n", err)
		} else {
			if needsUpgrade {
				return fmt.Errorf("a new version of package %q is available. Please install using \"bpm install %s\" or skip this check using \"--skip-upgrade-check\"", pluginName, pluginName)
			}
		}
	}

	// Create node config
	n := node.New(config.NodeFile(p.HomeDir, id))
	n.ID = id
	n.PluginName = pluginName
	n.StrParameters = strParameters
	n.BoolParameters = boolParameters
	n.Version = p.getInstalledVersion(pluginName)

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

	// Secrets have been removed but for compatibility reasons we still need to create the secrets directory for older plugins
	meta, err := p.getMeta(pluginName)
	if err != nil {
		return err
	}
	if meta.ProtocolVersion == "1.0.0" {
		_, err = fileutil.MakeDirectory(filepath.Join(n.NodeDirectory(), "secrets"))
		if err != nil {
			return err
		}

		err := p.execCmd(n, "create-secrets")
		if err != nil {
			return err
		}
	}

	// Config
	err = p.execCmd(n, "create-configurations")
	if err != nil {
		return err
	}

	fmt.Printf("\nNode with id %q has been initialized.\n\nTo change the configuration, modify the files here:\n    %s\nTo start the node, run:\n    bpm nodes start %s\nTo see the status of configured nodes, run:\n    bpm nodes status\n", id, n.ConfigsDirectory(), id)

	return nil
}
