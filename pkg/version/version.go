package version

import (
	"github.com/coreos/go-semver/semver"
)

func NeedsUpgrade(currentVersionStr, availableVersionStr string) (bool, error) {
	currentVersion, err := semver.NewVersion(currentVersionStr)
	if err != nil {
		return false, err
	}

	availableVersion, err := semver.NewVersion(availableVersionStr)
	if err != nil {
		return false, err
	}

	if currentVersion.LessThan(*availableVersion) {
		return true, nil
	}

	return false, nil
}
