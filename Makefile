# parameters

BINARY_NAME=ken-all
VETPKGS = $(shell go list ./... | grep -v -e vendor)

export GO111MODULE=on

.PHONY: build
build:
	gox --osarch "darwin/amd64 linux/amd64 windows/amd64" -output="bin/{{.OS}}_{{.Arch}}/$(BINARY_NAME)"

.PHONY: clean
clean:
	go clean
	rm -rf bin/

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: deps
deps:
	go get -d -v .
	go mod tidy

.PHONY: vet
vet:
	go vet $(VETPKGS)
