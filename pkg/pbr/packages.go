package pbr

import (
	"net/http"
	"net/url"
)

func (c *Client) SearchPackages(query, os string) ([]Package, error) {
	v := url.Values{}
	v.Set("query", query)
	v.Set("os", os)

	packages := []Package{}

	r, err := newRequest(
		http.MethodGet,
		c.host,
		"packages",
		v.Encode(),
		nil,
	)
	if err != nil {
		return packages, err
	}

	if err := c.do(r, &packages); err != nil {
		return packages, err
	}

	return packages, nil

}
