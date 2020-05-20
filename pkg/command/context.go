package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

// getExecutable returns the executable for the plugin
//
// There are two places in which we can find the executable:
// - If the plugin itself is a simple executable it's under `<bpm-dir>/plugins/pluginname`
// - If the plugin was a tar.gz it is in a subdirectory: `<bpm-dir>/plugins/pluginname/pluginname`
func (p *CmdContext) getExecutable(n node.Node) (string, error) {
	pluginName := n.PluginName

	// Check if plugin is installed
	if !p.isInstalled(pluginName) {
		return "", fmt.Errorf("the package %q is currently not installed", pluginName)
	}

	executableOrDirectory := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)

	fi, err := os.Stat(executableOrDirectory)
	if err != nil {
		return "", err
	}
	if fi.Mode().IsDir() {
		return filepath.Join(executableOrDirectory, pluginName), nil
	}

	return executableOrDirectory, nil
}

func (p *CmdContext) execCmdCapture(n node.Node, cmd string) (string, error) {
	executable, err := p.getExecutable(n)
	if err != nil {
		return "", err
	}

	// Run plugin commands
	return manager.ExecCmdCapture(p.Debug, executable, cmd, n.NodeFile())
}

func (p *CmdContext) execCmd(n node.Node, cmd string) error {
	executable, err := p.getExecutable(n)
	if err != nil {
		return err
	}

	// Run plugin commands
	return manager.ExecCmd(p.Debug, executable, cmd, n.NodeFile())
}

// getMetaFromManifest returns the meta information for a plugin from the manifest
//
// This is typically faster than getMetaFromExcutable because we have the data in memory already
func (p *CmdContext) getMetaFromManifest(pluginName string) (plugin.MetaInfo, error) {
	metaInfo, ok := p.Manifest.Plugins[pluginName]
	if !ok {
		return plugin.MetaInfo{}, fmt.Errorf("could not find %q in the manifest, is it installed?", pluginName)
	}

	return metaInfo, nil
}

// getMetaFromManifest returns the meta information for a plugin by calling the `meta` command on the plugin executable
func (p *CmdContext) getMetaFromExecutable(executablePath string) (plugin.MetaInfo, error) {
	meta := plugin.MetaInfo{}

	// Get parameter options
	output, err := manager.ExecCmdCapture(p.Debug, executablePath, "meta")
	if err != nil {
		return meta, err
	}

	err = yaml.Unmarshal([]byte(output), &meta)
	return meta, err
}

func (p *CmdContext) printfDebug(format string, a ...interface{}) {
	if p.Debug {
		if !strings.HasSuffix(format, "\n") {
			format = format + "\n"
		}

		fmt.Printf(format, a...)
	}
}
