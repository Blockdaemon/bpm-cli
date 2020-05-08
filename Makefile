# Need to wrap in "bash -c" so env vars work in the compiler as well as on the cli to specify the output
BUILD_CMD:=bash -c 'go build -ldflags "-X main.version=$(VERSION)" -o bin/bpm-$(VERSION)-$$GOOS-$$GOARCH cmd/*'

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 $(BUILD_CMD)
	GOOS=darwin GOARCH=amd64 $(BUILD_CMD)
	GOOS=windows GOARCH=amd64 $(BUILD_CMD)

.PHONY: check
check: test lint

.PHONY: test
test:
	go test -v ./...
	./smoke-test.sh 1.1.0 # Testing compatibility with old plugins
	./smoke-test.sh 1.2.0

.PHONY: lint
lint:
	golangci-lint run --enable gofmt ./...
.PHONY: build

.PHONY: pre-release
pre-release: 
	@ test -n "$(version)" || (echo 'ERROR: version is not set. Call like this: make version=1.14.0-rc1 release'; exit 1) 

	@ test -z "$$(git status --porcelain)" || (echo "ERROR: git is dirty - clean up first"; exit 1)

	@ echo "CHANGELOG.md starting here"
	@ echo "--------------------------"
	@ cat CHANGELOG.md
	@ read -p "Press enter to continue if the changelog looks ok. CTRL+C to abort."

.PHONY: release
release: pre-release check
	# tag it
	git tag v$(version)
	git push origin v$(version)

	# finally run the actually release
	gorelaser release --rm-dist

