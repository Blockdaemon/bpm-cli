module github.com/Blockdaemon/bpm

go 1.13

require (
	github.com/Blockdaemon/bpm-sdk v0.0.0-20190923132945-53b6830dfb4d
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883 // indirect
	github.com/coreos/go-semver v0.2.0
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/kami-zh/go-capturer v0.0.0-20171211120116-e492ea43421d
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rs/xid v1.2.1
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/xanzy/go-gitlab v0.22.0 // indirect
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/Blockdaemon/bpm-sdk => ../bpm-sdk
