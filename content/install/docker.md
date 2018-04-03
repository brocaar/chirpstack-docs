---
title: Docker install
menu:
    main:
        parent: install
        weight: 4
---

# Docker install

The LoRa Server project provides [Docker](https://www.docker.com) containers
for all project components. An overview of available containers can be found
at: [https://hub.docker.com/u/loraserver/](https://hub.docker.com/u/loraserver/).

## Versioning

* `latest` refers to the latest version from the `master` branch
* All other tags refer to tagged versions

## Docker Compose

A complete [docker-compose](https://docs.docker.com/compose/) example configuration
is provided by [https://github.com/brocaar/loraserver-docker](https://github.com/brocaar/loraserver-docker).

Example to get started (with the default EU868 band configuration):

```bash
$ git clone https://github.com/brocaar/loraserver-docker.git
$ cd loraserver-docker
$ docker-compose up
```

For more information refer to the `README.md` of
[https://github.com/brocaar/loraserver-docker](https://github.com/brocaar/loraserver-docker).

### Add network-server

When adding the network-server in the LoRa App Server web-interface
(see [network-servers](https://docs.loraserver.io/lora-app-server/use/network-servers/)),
you must enter `loraserver:8000` as the network-server `hostname:IP`.
