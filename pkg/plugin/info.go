package plugin

import (
	"bytes"

	"gitlab.com/Blockdaemon/bpm/pkg/config"
	"gitlab.com/Blockdaemon/bpm/pkg/pbr"
)

func Info(registry string, packageName string, os string, m config.Manifest) (string, error) {
	client := pbr.New(registry)

	versions, err := client.ListVersions(os, packageName)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	buf.WriteString("Name:        " + versions[0].Package.Name + "\n")
	buf.WriteString("Description: " + versions[0].Package.Description + "\n")
	buf.WriteString("Protocol:    " + versions[0].Package.Protocol + "\n")
	prefix := "Versions:    "
	for ix, version := range versions {
		buf.WriteString(prefix + version.Version + "\n")

		if ix == 0 {
			prefix = "             "
		}
	}

	return buf.String(), nil
}
