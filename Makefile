all: build

.PHONY: build

GOPATH ?= $(HOME)/go

PATH := $(GOPATH)/bin:$(PATH)
VERSION = $(shell git describe --tags --always --dirty)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)
PKGDIR_TMP = ${TMPDIR}golang
WORKSPACE = ${GOPATH}/src/github.com/clburlison/bakeit

-include config.mk

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
	-X github.com/clburlison/bakeit/src/version.appName=${APP_NAME} \
	-X github.com/clburlison/bakeit/src/version.version=${VERSION} \
	-X github.com/clburlison/bakeit/src/version.branch=${BRANCH} \
	-X github.com/clburlison/bakeit/src/version.buildUser=${USER} \
	-X github.com/clburlison/bakeit/src/version.buildDate=${NOW} \
	-X github.com/clburlison/bakeit/src/version.revision=${REVISION} \
	-X github.com/clburlison/bakeit/src/version.goVersion=${GOVERSION}"

define HELP_TEXT

  Makefile commands

	make deps         - Install dependent programs and libraries
	make clean        - Delete all build artifacts

	make build        - Build the code for current platform
	make build-all    - Build the code for all platforms
	make package      - Build macOS package (Not yet implemented)
	make generate     - Build updated resource.syso for windows binary

	make test         - Run the Go tests
	make lint         - Run the Go linters
	make test-ci      - Run the Go tests with circleci locally (Linux based)
	make test-cover   - Create test coverage and preview via web browser

  Administrative commands

	changelog         - Update the CHANGELOG.md

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
	go get -u github.com/golang/lint/golint
	go get -u github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	dep ensure -vendor-only -v

test:
	go test -cover -race -v `go list ./...`

test-ci:
	circleci build --job build-go1.9

# More info can be found https://blog.golang.org/cover & https://github.com/codecov/example-go
# This sed command will only work on macOS
test-cover:
	for d in $$(go list ./... | grep -v vendor); do \
		go test -race -coverprofile=profile.out $$d; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.out ; \
			rm profile.out ; \
		fi \
	done
	$$(sed -i '' '1!{/^mode: atomic/d;}' ./coverage.out)
	@go tool cover -html=coverage.out

lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	@go vet `go list ./...`
	golint -set_exit_status `go list ./... | grep -v /vendor/`

build: bakeit
build-all: xp-bakeit

clean:
	rm -rf build/
	rm -f *.zip
	rm -rf ${PKGDIR_TMP}_darwin
	rm -rf ${PKGDIR_TMP}_linux
	rm -rf ${PKGDIR_TMP}_windows

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux
	mkdir -p build/windows

APP_NAME = bakeit

.pre-bakeit:
	$(eval APP_NAME = bakeit)

bakeit: .pre-build .pre-bakeit
	go build -i -o build/$(CURRENT_PLATFORM)/${OUTPUT} -ldflags ${BUILD_VERSION} ./cmd/bakeit

install-bakeit: .pre-bakeit
	go install -ldflags ${BUILD_VERSION} ./cmd/bakeit

generate:
	go generate cmd/bakeit/main.go

xp-bakeit: .pre-build .pre-bakeit
	GOOS=darwin go build -i -o build/darwin/${APP_NAME} -pkgdir ${PKGDIR_TMP}_darwin -ldflags ${BUILD_VERSION} ./cmd/bakeit
	GOOS=linux go build -i -o build/linux/${APP_NAME} -pkgdir ${PKGDIR_TMP}_linux -ldflags ${BUILD_VERSION} ./cmd/bakeit
	GOOS=windows GOARCH=386 go build -i -o build/windows/${APP_NAME}.exe -pkgdir ${PKGDIR_TMP}_windows -ldflags ${BUILD_VERSION} ./cmd/bakeit

release-zip: xp-bakeit
	zip -r bakeit_${VERSION}.zip build/

changelog:
	docker run -it --rm -v "$(shell pwd)":/usr/local/src/your-app \
	clburlison/github-changelog-generator \
	-u clburlison -p bakeit \
	-t ${CHANGELOG_GITHUB_TOKEN}
	git add "CHANGELOG.md"
	git commit -m "chore: Update changelog"
