package pbr

import (
	"net/http"
	"net/url"
)

func (c *Client) ListVersions(os string, packageNames ...string) ([]Version, error) {
	v := url.Values{}
	v.Set("expand", "package")
	v.Set("os", os)

	for _, name := range packageNames {
		v.Set("package", name)
	}

	r, err := newRequest(
		http.MethodGet,
		c.host,
		"versions",
		v.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	versions := []Version{}

	if err := c.do(r, &versions); err != nil {
		return nil, err
	}

	return versions, nil
}
