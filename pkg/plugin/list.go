package plugin

import (
	"bytes"

	"github.com/kataras/tablewriter"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
)

func List(m config.Manifest) (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NAME",
		"PROTOCOL",
		"ENVIRONMENT",
		"NETWORK TYPE",
		"SUBTYPE",
		"INSTALLED VERSION",
	})

	// Get name and version from the manifest
	for name, p := range m.Plugins {
		table.Append([]string{
			name,
			p.Protocol,
			p.Environment,
			p.NetworkType,
			p.Subtype,
			p.Version,
		})
	}

	table.Render()

	return buf.String(), nil

}
