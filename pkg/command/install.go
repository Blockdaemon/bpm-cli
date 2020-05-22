package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"go.blockdaemon.com/bpm/cli/pkg/config"
	"go.blockdaemon.com/bpm/cli/pkg/manager"
	"go.blockdaemon.com/bpm/cli/pkg/pbr"
	"go.blockdaemon.com/bpm/sdk/pkg/fileutil"
)

func (p *CmdContext) addPluginToManifest(pluginName string, executablePath string) error {
	// Add plugin to manifest
	meta, err := p.getMetaFromExecutable(executablePath)
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
	p.printfDebug("Installing package %q from file %q\n", pluginName, sourcePath)

	targetPath := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	p.printfDebug("Remove previously installed version of there is one")
	p.clearInstallationDestination(targetPath)

	if err := config.CopyFile(sourcePath, targetPath); err != nil {
		return err
	}

	p.printfDebug("Changing %q to be executable\n", targetPath)
	if err := os.Chmod(targetPath, 0700); err != nil {
		return err
	}

	p.printfDebug("Adding package %q to manifest\n", pluginName)
	if err := p.addPluginToManifest(pluginName, targetPath); err != nil {
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

func (p *CmdContext) clearInstallationDestination(dstPath string) error {
	// If something (file or directory!) already exists, delete it
	exists, err := fileutil.FileExists(dstPath)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(dstPath)
	}

	return nil
}

// Install installs a particular version of a plugin
func (p *CmdContext) Install(pluginName, versionToInstall string) error {
	// Check if this version is already installed
	if p.getInstalledVersion(pluginName) == versionToInstall {
		fmt.Printf("%q version %q has already been installed", pluginName, versionToInstall)
		return nil
	}

	// Download the plugin file
	client := pbr.New(p.RegistryURL)
	ver, err := client.GetPackageVersion(pluginName, versionToInstall, p.RuntimeOS)
	if err != nil {
		return err
	}

	// If it's an archive we need to extract it first
	isArchive, err := pbr.URLPointsToArchive(ver.RegistryURL)
	if err != nil {
		return err
	}
	executablePath := ""
	if isArchive {
		p.printfDebug("Found tar.gz file for %q", pluginName)

		tempDir, err := ioutil.TempDir("", "bpm-")
		if err != nil {
			return err
		}
		defer os.RemoveAll(tempDir)

		const tmpFilename = "archive.tar.gz"

		p.printfDebug("Downloading %q", ver.RegistryURL)
		if err := manager.DownloadToFile(tempDir, tmpFilename, ver.RegistryURL); err != nil {
			return err
		}

		dstPath := filepath.Join(config.PluginsDir(p.HomeDir), pluginName)

		if err := p.clearInstallationDestination(dstPath); err != nil {
			return err
		}

		if err := os.Mkdir(dstPath, os.FileMode(0755)); err != nil {
			return err
		}
		p.printfDebug("Extracting %q", filepath.Join(tempDir, tmpFilename))
		if err := fileutil.ExtractTarGz(filepath.Join(tempDir, tmpFilename), dstPath); err != nil {
			return err
		}

		executablePath = filepath.Join(dstPath, pluginName)
	} else {
		p.printfDebug("Donwloading %q", ver.RegistryURL)
		if err := manager.DownloadToFile(config.PluginsDir(p.HomeDir), pluginName, ver.RegistryURL); err != nil {
			return err
		}

		executablePath = filepath.Join(config.PluginsDir(p.HomeDir), pluginName)
	}

	p.printfDebug("Changing permissions on %q to make sure it is executable", executablePath)
	if err := os.Chmod(executablePath, os.FileMode(0766)); err != nil {
		return err
	}

	if err := p.addPluginToManifest(pluginName, executablePath); err != nil {
		return err
	}

	fmt.Printf("The package %q has been installed.\n", pluginName)
	return nil
}
