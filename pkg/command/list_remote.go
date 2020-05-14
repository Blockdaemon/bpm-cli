package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/kataras/tablewriter"
	"go.blockdaemon.com/bpm/cli/pkg/pbr"
)

// Search searches the registry for plugins
func (p *CmdContext) Search(query string) error {
	client := pbr.New(p.RegistryURL)

	packages, err := client.SearchPackages(strings.ToLower(query), p.RuntimeOS)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NAME",
		"DESCRIPTION",
		"INSTALLED VERSION",
		"RECOMMENDED VERSION",
	})

	for _, pkg := range packages {
		// This is not exactly performant if there are lots of plugins available but it works well enough for now
		// Plenty of room for improvement by doing just one request total instead of one request per plugin
		latestVersion := "unknown"
		packageVersion, err := client.GetLatestPackageVersion(pkg.Name, p.RuntimeOS)
		if err != nil {
			fmt.Printf("Cannot get latest version for package %q\n", pkg.Name)
		} else {
			latestVersion = packageVersion.Version
		}

		installedVersion := ""
		pluginManifest, ok := p.Manifest.Plugins[pkg.Name]
		if ok {
			installedVersion = pluginManifest.Version
		}

		table.Append([]string{
			pkg.Name,
			pkg.Description,
			installedVersion,
			latestVersion,
		})
	}

	table.Render()

	return nil

}
