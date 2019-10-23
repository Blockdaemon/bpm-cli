module gitlab.com/Blockdaemon/bpm

go 1.13

require (
	github.com/Blockdaemon/bpm-sdk v0.0.0-20190923132945-53b6830dfb4d
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rs/xid v1.2.1
	github.com/spf13/cobra v0.0.5
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7
)

replace github.com/Blockdaemon/bpm-sdk => ../bpm-sdk
