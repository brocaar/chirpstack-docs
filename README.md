# ChirpStack documentation

This repository contains the global [chirpstack.io](https://www.chirpstack.io) documentation.

## Build

### Git submodules

Init and / or update the Git submodules:

```bash
git submodule init
git submodule update
```


### Install Hugo

[Hugo](http://gohugo.io/) is needed to transform the Markdown formatted
documentation into HTML.

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
