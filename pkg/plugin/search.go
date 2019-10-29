package plugin

import (
	"bytes"

	"github.com/kataras/tablewriter"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
)

func Search(registry string, query string, os string, m config.Manifest) (string, error) {
	client := pbr.New(registry)

	packages, err := client.SearchPackages(query, os)
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
