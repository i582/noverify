NOW=`date +%Y%m%d%H%M%S`
OS=`uname -n -m`
AFTER_COMMIT=`git rev-parse HEAD`
GOPATH_DIR=`go env GOPATH`

install:
	go install -ldflags "-X 'main.BuildTime=$(NOW)' -X 'main.BuildOSUname=$(OS)' -X 'main.BuildCommit=$(AFTER_COMMIT)'" .

check:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_DIR)/bin v1.30.0
	@echo "running linters..."
	@$(GOPATH_DIR)/bin/golangci-lint run ./src/...
	@echo "running tests..."
	@go test -coverprofile=coverage.txt -covermode=atomic -count 3 -race -v ./src/...
	@go test -race ./example/custom
	@echo "everything is OK"
	
cover:
	go install github.com/mattn/goveralls
	$(GOPATH_DIR)/bin/goveralls -package github.com/i582/noverify -coverprofile=coverage.out -service travis-ci -repotoken $(COVERALLS_TOKEN)

.PHONY: check cover
