package plugin

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Uninstall(pluginName string) (string, error) {
	if !p.isInstalled(pluginName) {
		return "", fmt.Errorf("The package %q is currently not installed.", pluginName)
	}

	// Delete the plugin
	if err := config.DeleteFile( config.PluginsDir(p.HomeDir), pluginName); err != nil {
		return "", err
	}

	// Remove plugin from manifest
	if err := p.Manifest.RemovePlugin(pluginName); err != nil {
		return "", err
	}

	return fmt.Sprintf("The package %q has been uninstalled.", pluginName), nil
}