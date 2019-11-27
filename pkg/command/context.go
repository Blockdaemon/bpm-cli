package command

import (
	"fmt"
	"path/filepath"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
	"github.com/Blockdaemon/bpm/pkg/pbr"
	"github.com/Blockdaemon/bpm/pkg/version"

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
		return "", fmt.Errorf("The package %q is currently not installed.\n", pluginName)
	}

	// Run plugin commands
	baseDir := config.NodesDir(p.HomeDir)
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	return manager.ExecCmdCapture(p.Debug, pluginFilename, "--base-dir", baseDir, cmd, n.ID)
}

func (p *CmdContext) execCmd(n node.Node, cmd string) error {
	pluginName := n.PluginName

	// Check if plugin is installed
	if !p.isInstalled(pluginName) {
		return fmt.Errorf("The package %q is currently not installed.\n", pluginName)
	}

	// Run plugin commands
	baseDir := config.NodesDir(p.HomeDir)
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	return manager.ExecCmd(p.Debug, pluginFilename, "--base-dir", baseDir, cmd, n.ID)
}

func (p *CmdContext) getMeta(pluginName string) (plugin.MetaInfo, error) {
	meta := plugin.MetaInfo{}
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	baseDirArgs := []string{"--base-dir", config.NodesDir(p.HomeDir)}

	// Get parameter options
	configArgs := append([]string{"meta"}, baseDirArgs...)
	output, err := manager.ExecCmdCapture(p.Debug, pluginFilename, configArgs...)
	if err != nil {
		return meta, err
	}

	err = yaml.Unmarshal([]byte(output), &meta)
	return meta, err
}
