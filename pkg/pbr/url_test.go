package pbr

import (
	"testing"
)

func TestIsTarGz(t *testing.T) {
	testFunc := func(urlUnderTest string, expected bool) {
		archive, err := URLPointsToArchive(urlUnderTest)
		if err != nil {
			t.Fatalf("unexpected error during test: %s", err)
		}
		if archive != expected {
			t.Errorf("URLPointsToArchive returned %t for %q; expected: %t", archive, urlUnderTest, expected)
		}
	}

	// Not an archive
	testFunc("https://runner-test.sfo2.digitaloceanspaces.com/polkadot-0.6.0-darwin-amd64", false)
	testFunc("https://runner-test.sfo2.digitaloceanspaces.com/polkadot-0.6.0-darwin-amd64?abitrary_parameter", false)
	testFunc("https://runner-test.sfo2.digitaloceanspaces.com/polkadot-0.6.0-darwin-amd64?abitrary_parameter&param2=3", false)
	testFunc("https://runner-test.sfo2.digitaloceanspaces.com/polkadot-0.6.0-darwin-amd64#anchor", false)
	// tar.gz archive
	testFunc("https://gitlab.com/blockdaemon/bpm-cli/uploads/e63d81151d00a71a24b498dc58518192/bpm-cli_0.14.0-rc7_darwin_386.tar.gz", true)
	testFunc("https://gitlab.com/blockdaemon/bpm-cli/uploads/e63d81151d00a71a24b498dc58518192/bpm-cli_0.14.0-rc7_darwin_386.tar.gz?arbitrary_parameter", true)
	testFunc("https://gitlab.com/blockdaemon/bpm-cli/uploads/e63d81151d00a71a24b498dc58518192/bpm-cli_0.14.0-rc7_darwin_386.tar.gz?arbitrary_parameter&param2=3", true)
	testFunc("https://gitlab.com/blockdaemon/bpm-cli/uploads/e63d81151d00a71a24b498dc58518192/bpm-cli_0.14.0-rc7_darwin_386.tar.gz#anchor", true)
	// tgz archive
	testFunc("https://gitlab.com/blockdaemon/bpm-cli/uploads/e63d81151d00a71a24b498dc58518192/bpm-cli_0.14.0-rc7_darwin_386.tgz", true)
}
