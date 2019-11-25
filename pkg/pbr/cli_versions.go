package pbr

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetCLIVersion(os string) (CLIVersion, error) {
	v := url.Values{}

	r, err := newRequest(
		http.MethodGet,
		c.host,
		fmt.Sprintf("cli/%s", os),
		v.Encode(),
		nil,
	)
	if err != nil {
		return CLIVersion{}, err
	}

	ver := CLIVersion{}

	if err := c.do(r, &ver); err != nil {
		return CLIVersion{}, err
	}

	return ver, nil

}

