package command

import (
	"fmt"
	"path/filepath"

	"go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/cli/pkg/manager"
	"go.blockdaemon.com/bpm/cli/pkg/pbr"
	"go.blockdaemon.com/bpm/cli/pkg/version"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"

	"gopkg.in/yaml.v2"
)

// CmdContext encapsulates data used by commands
type CmdContext struct {
	HomeDir     string
	Manifest    config.Manifest
	RuntimeOS   string
	RegistryURL string
	Debug       bool
}

// The following functions contain common functionality used by various commands
func (p *CmdContext) getInstalledVersion(pluginName string) string {
	if !p.isInstalled(pluginName) {
		return ""
	}

	return p.Manifest.Plugins[pluginName].Version
}

func (p *CmdContext) getLatestVersion(pluginName string) (string, error) {
	client := pbr.New(p.RegistryURL)

	packageVersion, err := client.GetLatestPackageVersion(pluginName, p.RuntimeOS)
	if err != nil {
		return "", err
	}

	return packageVersion.Version, nil
}

func (p *CmdContext) isInstalled(pluginName string) bool {
	_, ok := p.Manifest.Plugins[pluginName]

	return ok
}

func (p *CmdContext) needsUpgrade(pluginName string) (bool, error) {
	latestVersion, err := p.getLatestVersion(pluginName)
	if err != nil {
		return false, err
	}

	needsUpgrade, err := version.NeedsUpgrade(p.getInstalledVersion(pluginName), latestVersion)
	if err != nil {
		return false, err
	}

	return needsUpgrade, nil
}

func (p *CmdContext) execCmdCapture(n node.Node, cmd string) (string, error) {
	pluginName := n.PluginName

	// Check if plugin is installed
	if !p.isInstalled(pluginName) {
		return "", fmt.Errorf("the package %q is currently not installed", pluginName)
	}

	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	return manager.ExecCmdCapture(p.Debug, pluginFilename, cmd, n.NodeFile())
}

func (p *CmdContext) execCmd(n node.Node, cmd string) error {
	pluginName := n.PluginName

	// Check if plugin is installed
	if !p.isInstalled(pluginName) {
		return fmt.Errorf("the package %q is currently not installed", pluginName)
	}

	// Run plugin commands
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	return manager.ExecCmd(p.Debug, pluginFilename, cmd, n.NodeFile())
}

func (p *CmdContext) getMeta(pluginName string) (plugin.MetaInfo, error) {
	meta := plugin.MetaInfo{}
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)

	// Get parameter options
	output, err := manager.ExecCmdCapture(p.Debug, pluginFilename, "meta")
	if err != nil {
		return meta, err
	}

	err = yaml.Unmarshal([]byte(output), &meta)
	return meta, err
}
