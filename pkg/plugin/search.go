package plugin

import (
	"bytes"
	"strings"

	"github.com/kataras/tablewriter"
	"github.com/Blockdaemon/bpm/pkg/pbr"
)

func (p *PluginCmdContext) Search(query string) (string, error) {
	client := pbr.New(p.RegistryURL)

	packages, err := client.SearchPackages(strings.ToLower(query), p.RuntimeOS)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
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

	return buf.String(), nil

}
