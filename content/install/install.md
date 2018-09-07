---
title: General Installation
menu:
    main:
        parent: install
        weight: 3
---

# Installing the LoRa Server project

The LoRa Server components can be setup in multiple ways. For all cases,
downloads can be found at the following URLs:

* [LoRa Gateway Bridge downloads](/lora-gateway-bridge/overview/downloads/)
* [LoRa Server downloads](/loraserver/overview/downloads/)
* [LoRa App Server downloads](/lora-app-server/overview/downloads/)

## Manual

In this case you need to download the pre-compiled binaries for each
component and setup scripts so that these components will be started on
boot (if desired).

For this you could use [systemd](https://en.wikipedia.org/wiki/Systemd)
unit-files or [init](https://en.wikipedia.org/wiki/Init) script, based
on the used Linux distribution.

## Debian / Ubuntu repository

The LoRa Server project provides Debian / Ubuntu `.deb` packages which
can be downloaded from the LoRa Server Deb repository. To guarantee
compatibility, for each major version a separate repository is provided.
Please refer to the [Debian / Ubuntu install guide]({{<ref "/guides/debian-ubuntu.md">}})
for a step-by-step guide.

## Docker

The LoRa Server project also provides Docker images which for example can be used with
[Docker Compose](https://docs.docker.com/compose/). See the
[Docker install]({{<relref "docker.md">}}) for more information.
