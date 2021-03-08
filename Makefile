NOW=`date +%Y%m%d%H%M%S`
OS=`uname -n -m`
AFTER_COMMIT=`git rev-parse HEAD`
GOPATH_DIR=`go env GOPATH`

install:
	go install -ldflags "-X 'main.BuildTime=$(NOW)' -X 'main.BuildOSUname=$(OS)' -X 'main.BuildCommit=$(AFTER_COMMIT)'" .

build-release:
	go run ./_script/release.go -build-time="$(NOW)" -build-uname="$(OS)" -build-commit="$(AFTER_COMMIT)"

check:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH_DIR)/bin v1.30.0
	@echo "running linters..."
	@$(GOPATH_DIR)/bin/golangci-lint run ./src/...
	@echo "running tests..."
	@go test -tags tracing -count 3 -race -v ./src/...
	@go test -race ./example/custom
	@echo "everything is OK"

stat:
	go run . check -output=reports.json -output-json ./src/tests/golden/testdata

stat-master:
	git clone https://github.com/VKCOM/noverify.git ./noverify-master
	cd ./noverify-master
	go run . check -output=reports-master.json -output-json ./src/tests/golden/testdata
	ls
	cd ..

.PHONY: check build-release stat stat-master
