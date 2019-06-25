export GO111MODULE=on

# Vendoring is a bit controversial. I'm doing it for these reasons:
#
# - Being able to build without internet
# - Avoid dependencies dissapearing (https://www.theregister.co.uk/2016/03/23/npm_left_pad_chaos/)
# - Avoid attacks by injecting malicious code in dependencies (https://www.theregister.co.uk/2018/11/26/npm_repo_bitcoin_stealer/)
#
# The last point is admittedly unlikely because collision attacks with sha1 are still very expensive
# but with potentially lots of money at stake we shouldn't take changes.
#
# Linus Torvalds: There's a big difference between using a cryptographic hash for things like security signing,
#                 and using one for generating a 'content identifier' for a content-addressable system like git.
#
# (https://www.theregister.co.uk/2017/02/26/git_fscked_by_sha1_collision_not_so_fast_says_linus_torvalds/)
export GOFLAGS=-mod=vendor

VERSION:=$(CI_COMMIT_REF_NAME)

ifeq ($(VERSION),)
	# Looks like we are not running in the CI so default to current branch
	VERSION:=$(shell git rev-parse --abbrev-ref HEAD)
endif

# Need to wrap in "bash -c" so env vars work in the compiler as well as on the cli to specify the output
BUILD_CMD:=bash -c 'go build -ldflags "-X gitlab.com/Blockdaemon/bpm/internal/bpm/cmd.runnerVersion=$(VERSION)" -o binaries/bpm-$(VERSION)-$$GOOS-$$GOARCH cmd/bpm/main.go'

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 $(BUILD_CMD)
	GOOS=darwin GOARCH=amd64 $(BUILD_CMD)

.PHONY: check
check: lint test

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...
.PHONY: build

