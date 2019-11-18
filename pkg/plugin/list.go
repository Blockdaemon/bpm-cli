package plugin

import (
	"os"

	"github.com/kataras/tablewriter"
	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) List() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NAME",
		"INSTALLED VERSION",
		"AVAILABLE VERSION",
	})

	client := pbr.New(p.RegistryURL)

	// Get name and version from the manifest
	for name, plugin := range p.Manifest.Plugins {
		// This is not exactly performant if there are lots of plugins installed but it works well enough for now
		// Plenty of room for improvement by doing just one request total instead of one request per plugin
		packageVersion, err := client.GetLatestPackageVersion(name, p.RuntimeOS)
		if err != nil {
			return err
		}
		latestVersion := packageVersion.Version

		table.Append([]string{
			name,
			plugin.Version,
			latestVersion,
		})
	}

	table.Render()
	return nil

}
