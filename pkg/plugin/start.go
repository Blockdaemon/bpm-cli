package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm/pkg/config"
)

func (p *PluginCmdContext) Start(nodeID string) (string, error) {
	// Check if manual intervention is necessary in the configs
	// This is the case if a string like, e.g. {% ADD NODE KEY HERE %} is found in the files.
	// Until we have a better way of getting this information via the CLI, the users can edit the files manually.
	var ff = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // skip dirs
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var substitute = regexp.MustCompile(`{%[^%]*%}`)

		matches := substitute.FindSubmatch(content)

		if len(matches) > 0 {
			return fmt.Errorf("The string %q needs to be replaced with a suitable value in %q", string(matches[0]), path)
		}

		return nil
	}

	n, err := node.Load(config.NodesDir(p.HomeDir), nodeID)
	if err != nil {
		return "", err
	}

	if err := filepath.Walk(n.ConfigsDirectory(), ff); err != nil {
		return "", err
	}

	return p.execNodeCommand(n, "start")
}
