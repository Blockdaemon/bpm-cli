package pbr

import (
	"net/url"
	"strings"
)

// URLPointsToArchive returns true if the URL points to a tar.gz file
func URLPointsToArchive(rawurl string) (bool, error) {
	parsedurl, err := url.Parse(rawurl)
	if err != nil {
		return false, err
	}

	if strings.HasSuffix(parsedurl.Path, "tar.gz") {
		return true, nil
	}

	if strings.HasSuffix(parsedurl.Path, "tgz") {
		return true, nil
	}

	return false, nil
}
