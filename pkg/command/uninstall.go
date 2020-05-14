package command

import (
	"fmt"

	"go.blockdaemon.com/bpm/cli/pkg/config"
)

func (p *CmdContext) Uninstall(pluginName string) error {
	if !p.isInstalled(pluginName) {
		return fmt.Errorf("The package %q is currently not installed.", pluginName)
	}

	// Delete the plugin
	if err := config.DeleteFile(config.PluginsDir(p.HomeDir), pluginName); err != nil {
		return err
	}

	// Remove plugin from manifest
	if err := p.Manifest.RemovePlugin(pluginName); err != nil {
		return err
	}

	fmt.Printf("The package %q has been uninstalled.\n", pluginName)
	return nil
}
