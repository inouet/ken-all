# parameters

BINARY_NAME=ken-all
VETPKGS = $(shell go list ./... | grep -v -e vendor)

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
	go list ./... | xargs golint -set_exit_status | grep -v 'should have comment or be unexported' | grep -v 'comment on exported function'

.PHONY: deps
deps:
	go get -d -v .
	dep ensure -update

.PHONY: vet
vet:
	go vet $(VETPKGS)
