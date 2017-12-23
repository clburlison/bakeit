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
PKGDIR_TMP = "/private/tmp/bakeit_pkgdir_tmp"
WORKSPACE = ${GOPATH}/src/github.com/clburlison/bakeit

ifneq ($(OS), Windows_NT)
	CURRENT_PLATFORM = linux
	# If on macOS, set the shell to bash explicitly
	ifeq ($(shell uname), Darwin)
		SHELL := /bin/bash
		CURRENT_PLATFORM = darwin
	endif

	# The output binary name is different on Windows, so we're explicit here
	OUTPUT = bakeit

	# To populate version metadata, we use unix tools to get certain data
	GOVERSION = $(shell go version | awk '{print $$3}')
	NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
else
	# The output binary name is different on Windows, so we're explicit here
	OUTPUT = bakeit.exe
	CURRENT_PLATFORM = windows

	# To populate version metadata, we use windows tools to get the certain data
	GOVERSION_CMD = "(go version).Split()[2]"
	GOVERSION = $(shell powershell $(GOVERSION_CMD))
	NOW	= $(shell powershell Get-Date -format s)
endif

BUILD_VERSION = "\
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.appName=${APP_NAME} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.version=${VERSION} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.branch=${BRANCH} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.buildUser=${USER} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.buildDate=${NOW} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.revision=${REVISION} \
	-X github.com/clburlison/bakeit/vendor/github.com/bakeit/go4/version.goVersion=${GOVERSION}"

define HELP_TEXT

  Makefile commands

	make deps         - Install dependent programs and libraries
	make clean        - Delete all build artifacts

	make build        - Build the code
	make build-all    - Build the code for all platforms
	make package      - Build macOS package (Not yet implemented)

	make test         - Run the Go tests
	make lint         - Run the Go linters

endef

help:
	$(info $(HELP_TEXT))

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

lint:
	go vet ./...

build: bakeit
build-all: xp-bakeit

clean:
	rm -rf build/
	rm -f *.zip

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux
	mkdir -p build/windows

APP_NAME = bakeit

.pre-bakeit:
	$(eval APP_NAME = bakeit)

bakeit: .pre-build .pre-bakeit
	go build -i -o build/$(CURRENT_PLATFORM)/bakeit -ldflags ${BUILD_VERSION} ./cmd/bakeit

install-bakeit: .pre-bakeit
	go install -ldflags ${BUILD_VERSION} ./cmd/bakeit

xp-bakeit: .pre-build .pre-bakeit
	GOOS=darwin go build -i -o build/darwin/${OUTPUT} -ldflags ${BUILD_VERSION} ./cmd/bakeit
	GOOS=linux go build -i -o build/linux/${OUTPUT} -pkgdir ${PKGDIR_TMP}_linux -ldflags ${BUILD_VERSION} ./cmd/bakeit
	GOOS=windows GOARCH=386 go build -i -o build/windows/${OUTPUT}.exe -pkgdir ${PKGDIR_TMP}_windows -ldflags ${BUILD_VERSION} ./cmd/bakeit

release-zip: xp-bakeit
	zip -r bakeit_${VERSION}.zip build/
