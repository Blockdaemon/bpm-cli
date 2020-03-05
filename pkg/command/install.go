package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Blockdaemon/bpm/pkg/config"
	"github.com/Blockdaemon/bpm/pkg/manager"
	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *CmdContext) addPluginToManifest(pluginName string) error {
	// Add plugin to manifest
	meta, err := p.getMeta(pluginName)
	if err != nil {
		return err
	}
	return p.Manifest.UpdatePlugin(pluginName, meta)
}

// InstallFile installs a plugin from a local file.
//
// This is very useful during development to avoid having to upload a plugin
// to the registry every time we want to test a change.
func (p *CmdContext) InstallFile(pluginName string, sourcePath string) error {
	if p.Debug {
		fmt.Printf("Installing package %q from file %q\n", pluginName, sourcePath)
	}
	targetPath := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	if err := config.CopyFile(sourcePath, targetPath); err != nil {
		return err
	}

	if p.Debug {
		fmt.Printf("Changing %q to be executable\n", targetPath)
	}
	if err := os.Chmod(targetPath, 0700); err != nil {
		return err
	}

	if p.Debug {
		fmt.Printf("Adding package %q to manifest\n", pluginName)
	}
	if err := p.addPluginToManifest(pluginName); err != nil {
		return err
	}

	fmt.Printf("The package %q has been installed.\n", pluginName)
	return nil
}

// InstallLatest installs the latest version of a plugin
func (p *CmdContext) InstallLatest(pluginName string) error {
	latestVersion, err := p.getLatestVersion(pluginName)
	if err != nil {
		return err
	}

	return p.Install(pluginName, latestVersion)
}

// Install installs a particular version of a plugin
func (p *CmdContext) Install(pluginName, versionToInstall string) error {
	// Check if this version is already installed
	if p.getInstalledVersion(pluginName) == versionToInstall {
		return fmt.Errorf("%q version %q has already been installed", pluginName, versionToInstall)
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

	if err := p.addPluginToManifest(pluginName); err != nil {
		return err
	}

	fmt.Printf("The package %q has been installed.\n", pluginName)
	return nil
}
