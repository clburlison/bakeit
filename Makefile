all: build

.PHONY: build

ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif

PATH := $(GOPATH)/bin:$(PATH)
VERSION = $(shell git describe --tags --always --dirty)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)
GOVERSION = $(shell go version | awk '{print $$3}')
NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
SHELL = /bin/bash

ifneq ($(OS), Windows_NT)
	CURRENT_PLATFORM = linux
	ifeq ($(shell uname), Darwin)
		SHELL := /bin/bash
		CURRENT_PLATFORM = darwin
	endif
else
	CURRENT_PLATFORM = windows
endif

BUILD_VERSION = "\
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.appName=${APP_NAME} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.version=${VERSION} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.branch=${BRANCH} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.buildUser=${USER} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.buildDate=${NOW} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.revision=${REVISION} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.goVersion=${GOVERSION}"

WORKSPACE = ${GOPATH}/src/github.com/clburlison/bakeit
check-deps:
ifneq ($(shell test -e ${WORKSPACE}/Gopkg.lock && echo -n yes), yes)
	@echo "folder is cloned in the wrong place, copying to a Go Workspace"
	@echo "See: https://golang.org/doc/code.html#Workspaces"
	@git clone git@github.com:clburlison/bakeit ${WORKSPACE}
	@echo "cd to ${WORKSPACE} and run make deps again."
	@exit 1
endif
ifneq ($(shell pwd), $(WORKSPACE))
	@echo "cd to ${WORKSPACE} and run make deps again."
	@exit 1
endif

deps: check-deps
	go get -u github.com/golang/dep/...
	dep ensure -vendor-only

test:
	go test -cover -race -v $(shell go list ./... | grep -v /vendor/)

build: bakeit

clean:
	rm -rf build/
	rm -f *.zip

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux

INSTALL_STEPS := \
	install-bakeit

install-local: $(INSTALL_STEPS)

APP_NAME = bakeit

.pre-bakeit:
	$(eval APP_NAME = bakeit)

bakeit: .pre-build .pre-bakeit
	go build -i -o build/$(CURRENT_PLATFORM)/bakeit -ldflags ${BUILD_VERSION} ./cmd/bakeit

install-bakeit: .pre-bakeit
	go install -ldflags ${BUILD_VERSION} ./cmd/bakeit

xp-bakeit: .pre-build .pre-bakeit
	GOOS=darwin go build -i -o build/darwin/bakeit -ldflags ${BUILD_VERSION} ./cmd/bakeit
	GOOS=linux CGO_ENABLED=0 go build -i -o build/linux/bakeit  -ldflags ${BUILD_VERSION} ./cmd/bakeit

release-zip: xp-bakeit xp-mdmctl
	zip -r bakeit_${VERSION}.zip build/
