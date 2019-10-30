package plugin

import (
	"bytes"

	"github.com/kataras/tablewriter"
	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
)

func List(registry string, m config.Manifest, os string) (string, error) {
	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetBorder(false)
	table.SetHeader([]string{
		"NAME",
		"INSTALLED VERSION",
		"AVAILABLE VERSION",
	})

	client := pbr.New(registry)

	// Get name and version from the manifest
	for name, p := range m.Plugins {
		// This is not exactly performant if there are lots of plugins installed but it works well enough for now
		// Plenty of room for improvement by doing just one request total instead of one request per plugin
		packageVersion, err := client.GetLatestPackageVersion(name, os)
		if err != nil {
			return "", err
		}
		latestVersion := packageVersion.Version

		table.Append([]string{
			name,
			p.Version,
			latestVersion,
		})
	}

	table.Render()

	return buf.String(), nil

}
