package plugin

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Uninstall(pluginName string) (string, error) {
	if !p.isInstalled(pluginName) {
		return "", fmt.Errorf("The package %q is currently not installed.\n", pluginName)
	}

	// Delete the plugin
	if err := config.DeleteFile( config.PluginsDir(p.HomeDir), pluginName); err != nil {
		return "", err
	}

	// Remove plugin from manifest
	delete(p.Manifest.Plugins, pluginName)

	if err := config.WriteFile(p.HomeDir, config.ManifestFilename, p.Manifest); err != nil {
		return "", err
	}

	return fmt.Sprintf("The package %q has been uninstalled.\n", pluginName), nil
}
