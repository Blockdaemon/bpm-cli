package plugin

import (
	"fmt"

	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) InstallLatest(pluginName string) error {
	latestVersion, err := p.getLatestVersion(pluginName)
	if err != nil {
		return err
	}

	return p.Install(pluginName, latestVersion)
}

func (p *PluginCmdContext) Install(pluginName, versionToInstall string) error {
	// Check if this version is already installed
	if p.getInstalledVersion(pluginName) == versionToInstall {
		return fmt.Errorf("%q version %q has already been installed.", pluginName, versionToInstall)
	}

	// Download the plugin file
	client := pbr.New(p.RegistryURL)
	ver, err := client.GetPackageVersion(pluginName, versionToInstall, p.RuntimeOS)
	if err != nil {
		return err
	}
	if err := manager.DownloadToFile(config.PluginsDir(p.HomeDir), pluginName, ver.RegistryURL); err != nil {
		return err
	}

	// Add plugin to manifest
	if err := p.Manifest.UpdatePlugin(pluginName, versionToInstall); err != nil {
		return err
	}

	fmt.Printf("The package %q has been installed.\n", pluginName)
	return nil
}
