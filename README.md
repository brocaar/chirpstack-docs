# docs.loraserver.io

This repository contains the global LoRa Server project documentation.

## Build

[Hugo](http://gohugo.io/) is needed to transform the Markdown formatted
documentation into HTML.

### Install Hugo

Please refer to the Hugo website for binaries, or install Hugo from source:

```bash
go get -v github.com/spf13/hugo
```

### Start Hugo

To start Hugo as a web-server (auto-refreshing on file-change), run the
following command from the root of this repository:

```bash
hugo server -w
```