.PHONY: version verify test vet lint deps

# Info
NAME=scaleway
VERSION := $(shell git describe --abbrev=0 --tags 2> /dev/null || echo unknown)
REVISION := $(shell git rev-parse --short HEAD || echo unknown)

export GO15VENDOREXPERIMENT := 1

help:
	# make version - show information about current version
	# make deps - install dependencies
	# make verify - run vet, lint and test
version:
	@echo Name: $(NAME)
	@echo Version: $(VERSION)
	@echo Revision: $(REVISION)

verify: gover vet lint test

gover:
	@go version

test:
	@go test -v -covermode=count -coverprofile=coverage.out

vet:
	@go vet *.go

lint:
	@golint *.go

deps:
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls
	@go get -u github.com/golang/lint/golint
