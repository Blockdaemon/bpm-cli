package plugin

import (
	"os"
	"strings"

	"github.com/kataras/tablewriter"
	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) Search(query string) error {
	client := pbr.New(p.RegistryURL)

	packages, err := client.SearchPackages(strings.ToLower(query), p.RuntimeOS)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NAME",
		"PROTOCOL",
		"DESCRIPTION",
	})

	for _, pkg := range packages {
		table.Append([]string{
			pkg.Name,
			pkg.Protocol,
			pkg.Description,
		})
	}

	table.Render()

	return nil

}
