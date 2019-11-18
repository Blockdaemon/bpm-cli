package plugin

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

// PlguinCmdContext encapsulates data used by commands related to plugins
type PluginCmdContext struct {
	HomeDir     string
	Manifest    config.Manifest
	RuntimeOS   string
	RegistryURL string
	Debug       bool
}

// The following functions contain common functionality used by various commands
func (p *PluginCmdContext) getInstalledVersion(pluginName string) string {
	if !p.isInstalled(pluginName) {
		return ""
	}

	return p.Manifest.Plugins[pluginName].Version
}

func (p *PluginCmdContext) getLatestVersion(pluginName string) (string, error) {
	client := pbr.New(p.RegistryURL)

	packageVersion, err := client.GetLatestPackageVersion(pluginName, p.RuntimeOS)
	if err != nil {
		return "", err
	}

	return packageVersion.Version, nil
}

func (p *PluginCmdContext) isInstalled(pluginName string) bool {
	_, ok := p.Manifest.Plugins[pluginName]

	return ok
}

func (p *PluginCmdContext) needsUpgrade(pluginName string) (bool, string, error) {
	latestVersion, err := p.getLatestVersion(pluginName)
	if err != nil {
		return false, "", err
	}

	needsUpgrade, err := version.NeedsUpgrade(p.getInstalledVersion(pluginName), latestVersion)
	if err != nil {
		return false, "", err
	}

	return needsUpgrade, latestVersion, nil
}

func (p *PluginCmdContext) execCmdCapture(n node.Node, cmd string) (string, error) {
	pluginName := n.Protocol // TODO: Wrong variable, should be pluginName

	// Check if plugin is installed
	if !p.isInstalled(pluginName) {
		return "", fmt.Errorf("The package %q is currently not installed.\n", pluginName)
	}

	// Run plugin commands
	baseDir := config.NodesDir(p.HomeDir)
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	return manager.ExecCmdCapture(p.Debug, pluginFilename, "--base-dir", baseDir, cmd, n.ID)
}

func (p *PluginCmdContext) execCmd(n node.Node, cmd string) error {
	pluginName := n.Protocol // TODO: Wrong variable, should be pluginName

	// Check if plugin is installed
	if !p.isInstalled(pluginName) {
		return fmt.Errorf("The package %q is currently not installed.\n", pluginName)
	}

	// Run plugin commands
	baseDir := config.NodesDir(p.HomeDir)
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	return manager.ExecCmd(p.Debug, pluginFilename, "--base-dir", baseDir, cmd, n.ID)
}

func (p *PluginCmdContext) getParameterOptions(pluginName string) (plugin.Parameters, error) {
	parameterOptions := plugin.Parameters{}
	pluginFilename := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	baseDirArgs := []string{"--base-dir", config.NodesDir(p.HomeDir)}

	// Get parameter options
	configArgs := append([]string{"parameters"}, baseDirArgs...)
	output, err := manager.ExecCmdCapture(false, pluginFilename, configArgs...)
	if err != nil {
		return parameterOptions, err
	}

	err = yaml.Unmarshal([]byte(output), &parameterOptions)
	return parameterOptions, err
}
