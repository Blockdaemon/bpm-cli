package command

import (
	"fmt"
	"os"

	"github.com/Blockdaemon/bpm/pkg/pbr"
	"github.com/kataras/tablewriter"
)

func (p *CmdContext) List() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NAME",
		"DESCRIPTION",
		"INSTALLED VERSION",
		"RECOMMENDED VERSION",
	})

	client := pbr.New(p.RegistryURL)

	// Get name and version from the manifest
	for name, plugin := range p.Manifest.Plugins {
		// This is not exactly performant if there are lots of plugins installed but it works well enough for now
		// Plenty of room for improvement by doing just one request total instead of one request per plugin
		latestVersion := "unknown"
		packageVersion, err := client.GetLatestPackageVersion(name, p.RuntimeOS)
		if err != nil {
			fmt.Printf("Cannot get latest version for package %q\n", name)
		} else {
			latestVersion = packageVersion.Version
		}

		table.Append([]string{
			name,
			plugin.Description,
			plugin.Version,
			latestVersion,
		})
	}

	table.Render()
	return nil

}
