This was a side project to learn some Go. Don't use this in production. Support not given, use at your own risk, etc. 


# bakeit

[![CircleCI][img-circleci-badge]][cirlce-ci]
[![PRs Welcome][img-prs-welcome-badge]][prs-welcome]
[![Go Report Card][img-reportcard-badge]][reportcard]

`bakeit` is a platform agnostic chef bootstrap tool. Unlike `knife bootstrap`,
bakeit is written for end points not servers. All required configuration
data is compiled into a single static binary that can be ran on the end point
for bootstrapping.

## Table of Contents

* [**Features**](#features)
* [**Basic Usage**](#basic-usage)
* [**Developers**](#developers)
* [**Changelog**](#changelog)
* [**License**](#license)

## Features

* No extra dependencies necessary on your end points
* Shared code between each platform where possible
* One file, [config.go][], holds all configuration settings
* Works on Mac, Linux (coming soon) and Windows

## Basic Usage

This section is designed for users new to the Go ecosystem that just want to
build and use this project.

1. Download and install [Go 1.9][download-go]
1. Set required Go variables
    ```bash
    export GOPATH=$(go env GOPATH)
    PATH=$PATH:${GOPATH}/bin
    ```
1. Create the proper go path
    ```bash
    mkdir -p $GOPATH/src/github.com/clburlison/bakeit
    ```
1. Clone this repo
    ```bash
    git clone git@github.com:clburlison/bakeit $GOPATH/src/github.com/clburlison/bakeit
    ```
1. Download dependencies
    ```bash
    make deps
    ```
1. Modify the [config.go][] file. Make sure to modify the following keys:
    * `ChefClientChefServerURL` - The URL to your chef server
    * `ChefClientValidationClientName` - The name of your validator file
    * `ValidationPEM` - The contents of your validator certificate
    * `OrgCert` - (Optional) The contents of your organization certificate
1. Build
    ```bash
    make build-all
    ```
1. Copy the correct output file from build/ to a machine and run it.
1. Running instructions
    ```bash
    # Linux/macOS
    sudo /path/bakeit

    # Windows. Open a command prompt as administrator
    /path/bakeit.exe
    ```

## Developers
Coming soon!
<!-- https://golang.org/doc/code.html, https://blog.golang.org/cover -->

## Contributing

See [CONTRIBUTING.md](.github/CONTRIBUTING.md)

## Changelog

See [CHANGELOG.md](CHANGELOG.md)

## License

[MIT](LICENSE) © Clayton Burlison

<!--
Link References
-->

[img-circleci-badge]:https://circleci.com/gh/clburlison/bakeit.svg?style=shield&circle-token=e56e3ca96a10956ff58dc8f504601d28778cb7c2
[img-prs-welcome-badge]:https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat
[img-reportcard-badge]:https://goreportcard.com/badge/github.com/clburlison/bakeit
[cirlce-ci]:https://circleci.com/gh/clburlison/bakeit
[prs-welcome]:http://makeapullrequest.com
[reportcard]:https://goreportcard.com/report/github.com/clburlison/bakeit
[download-go]: https://golang.org/dl/
[config.go]: https://github.com/clburlison/bakeit/blob/master/src/config/config.go
