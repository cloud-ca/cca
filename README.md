# cca [![Build Status](https://github.com/cloud-ca/cca/workflows/build/badge.svg)](https://github.com/cloud-ca/cca/actions) [![GoDoc](https://godoc.org/github.com/cloud-ca/cca?status.svg)](https://godoc.org/github.com/cloud-ca/cca) [![Go Report Card](https://goreportcard.com/badge/github.com/cloud-ca/cca)](https://goreportcard.com/report/github.com/cloud-ca/cca)

This tool allows you to interact with a [cloud.ca](https://cloud.ca/`) services via a command line interface.

## Installation

The latest version can be installed using `go get`:

``` bash
GO111MODULE="on" go get github.com/cloud-ca/cca@v0.0.1
```

**NOTE:** please use the latest go to do this, ideally go 1.12.9 or greater.

This will put `cca` in `$(go env GOPATH)/bin`. If you encounter the error `cca: command not found` after installation then you may need to either add that directory to your `$PATH` as shown [here](https://golang.org/doc/code.html#GOPATH) or do a manual installation by cloning the repo and run `make build` from the repository which will put `cca` in:

```bash
$(go env GOPATH)/src/github.com/cloud-ca/cca/bin/$(uname | tr '[:upper:]' '[:lower:]')-amd64/cca
```

Stable binaries are also available on the [releases](https://github.com/cloud-ca/cca/releases) page. To install, download the binary for your platform from "Assets" and place this into your `$PATH`:

```bash
curl -Lo ./cca.tar.gz https://github.com/cloud-ca/cca/releases/download/v0.0.1/cca-$(uname)-amd64.tar.gz
tar -xzf ./cca.tar.gz
chmod +x ./cca
rm ./cca.tar.gz
mv ./cca /some-dir-in-your-PATH/cca
```

**NOTE:** Windows releases are compressed in `ZIP` format.

## Code Completion

The code completion for `bash` or `zsh` can be installed using:

**Note:** Shell auto-completion is not available for Windows users.

### bash

``` bash
cca completion bash > ~/.cca-completion
source ~/.cca-completion

# or simply the one-liner below
source <(cca completion bash)
```

### zsh

``` bash
cca completion zsh > /usr/local/share/zsh/site-functions/_cca
autoload -U compinit && compinit
```

To make this change permenant, the above commands can be added to your `~/.profile` file.
