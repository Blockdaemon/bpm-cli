package pbr

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetPackageVersion(packageName, version, os string) (Version, error) {
	v := url.Values{}
	v.Set("expand", "package")
	v.Set("os", os)

	r, err := newRequest(
		http.MethodGet,
		c.host,
		fmt.Sprintf("packages/%s/versions/%s", packageName, version),
		v.Encode(),
		nil,
	)
	if err != nil {
		return Version{}, err
	}

	ver := Version{}

	if err := c.do(r, &ver); err != nil {
		return Version{}, err
	}

	return ver, nil

}
