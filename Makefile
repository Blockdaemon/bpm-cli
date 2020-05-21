# Need to wrap in "bash -c" so env vars work in the compiler as well as on the cli to specify the output
BUILD_CMD:=bash -c 'go build -ldflags "-X main.version=$(VERSION)" -o bin/bpm-$(VERSION)-$$GOOS-$$GOARCH cmd/*'

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 $(BUILD_CMD)
	GOOS=darwin GOARCH=amd64 $(BUILD_CMD)
	GOOS=windows GOARCH=amd64 $(BUILD_CMD)

.PHONY: check
check: unit-test lint

.PHONY: unit-test
unit-test:
	go test -v ./...

.PHONY: smoke-test
smoke-test:
	./smoke-test.sh polkadot 1.1.0 # Testing compatibility with old plugins
	./smoke-test.sh polkadot 1.2.0
	./smoke-test.sh tezos 1.2.0

.PHONY: lint
lint:
	golangci-lint run --enable gofmt ./...

.PHONY: pre-release
pre-release: 
	@ test -n "$(version)" || (echo 'ERROR: version is not set. Call like this: make version=1.14.0-rc1 release'; exit 1) 

	@ test -n "$(GITLAB_TOKEN)" || (echo 'ERROR: GITLAB_TOKEN is not set. See: https://goreleaser.com/quick-start/'; exit 1) 

	@ test -z "$$(git status --porcelain)" || (echo "ERROR: git is dirty - clean up first"; exit 1)

	@ echo "CHANGELOG.md starting here"
	@ echo "--------------------------"
	@ cat CHANGELOG.md
	@ read -p "Press enter to continue if the changelog looks ok. CTRL+C to abort."

.PHONY: dev-release
dev-release: smoke-test check
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY: release
release: pre-release smoke-test check
	# tag it
	git tag v$(version)
	git push origin v$(version)

	# finally run the actually release
	goreleaser release --rm-dist

