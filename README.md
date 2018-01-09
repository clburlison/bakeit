# bakeit [![CircleCI](https://circleci.com/gh/clburlison/bakeit.svg?style=svg&circle-token=e56e3ca96a10956ff58dc8f504601d28778cb7c2)](https://circleci.com/gh/clburlison/bakeit)

### Overview

bakeit is designed to be a platform agnostic chef bootstrap tool. Unlike `knife bootstrap`, bakeit is written to be ran on end points.

### Requirements

1. [Go 1.9](https://golang.org/dl/)
1. Chef validator key

### Building

```bash
# if GOPATH is unset:
# export GOPATH=$(HOME)/go

# Clone the repo into GOPATH:
git clone git@github.com:clburlison/bakeit $GOPATH/src/github.com/clburlison/bakeit
cd $GOPATH/src/github.com/clburlison/bakeit

# Download dependencies and build:
make deps
make
```

* More Info on [`GOPATH`](https://golang.org/doc/code.html#GOPATH)
