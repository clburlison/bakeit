# bakeit

[![CircleCI][img-circleci-badge]][cirlce-ci]
[![PRs Welcome][img-prs-welcome-badge]][prs-welcome]

`bakeit` is a platform agnostic chef bootstrap tool. Unlike `knife bootstrap`,
bakeit is written for end points not servers. All required configuration
data is compiled into a single static binary that can be ran on the end point
for bootstrapping.

## Table of Contents

[**Features**](#features)
[**Basic Usage**](#basic-usage)
[**Developers**](#developers)
[**Changelog**](#changelog)
[**License**](#license)

## Features

* No extra dependencies necessary on your end points
* Shared code between each platform where possible
* One file holds all configuration settings
* Works on Mac, Linux (coming soon) and Windows

## Basic Usage

This section is designed for users new to the Go ecosystem that just want to
build and use this project.

1. Download and install [Go 1.9][download-go]
1. Clone this repo
    ```bash
    export GOPATH=$(go env GOPATH)
    PATH=$PATH:${GOPATH}/bin
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

[img-circleci-badge]:https://circleci.com/gh/clburlison/bakeit.svg?style=svg&circle-token=e56e3ca96a10956ff58dc8f504601d28778cb7c2
[img-prs-welcome-badge]:https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square
[cirlce-ci]:https://circleci.com/gh/clburlison/bakeit
[prs-welcome]:http://makeapullrequest.com
[download-go]: https://golang.org/dl/
[config.go]: https://github.com/clburlison/bakeit/blob/master/src/config/config.go
