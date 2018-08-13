---
title: Docker install
menu:
    main:
        parent: install
        weight: 4
---

# Docker install

The LoRa Server project provides [Docker](https://www.docker.com) images
for all project components. An overview of available images can be found
at: [https://hub.docker.com/u/loraserver/](https://hub.docker.com/u/loraserver/).

## Versioning

* `latest` refers to the latest version from the `master` branch. It is not
  recommeded to use this for production!
* All other tags refer to tagged versions.

## Docker Compose

[Docker Compose](https://docs.docker.com/compose/) (part of Docker) makes
it possible to orchestrate the configuration of multiple Docker containers
at once using a `docker-compose.yml` file.

The LoRa Server project provides an example `docker-compose.yml` file that
you can use as a starting-point. This example can be found at
[https://github.com/brocaar/loraserver-docker](https://github.com/brocaar/loraserver-docker)
and also contains more documentation.

Example to get started (with the default EU868 band configuration):

```bash
$ git clone https://github.com/brocaar/loraserver-docker.git
$ cd loraserver-docker
$ docker-compose up
```

### Add network-server

As each container has its own hostname, you must use the hostname of the 
`loraserver` container when adding the network-server in the LoRa App Server
web-interface.

When using the above example, it means that you must enter `loraserver:8000`
as the network-server hostname:IP. See [network-servers](https://docs.loraserver.io/lora-app-server/use/network-servers/)
for more information.
