---
description: General installation instructions to setup the ChirpStack LoRaWAN Network Server components.
---

# General install

The ChirpStack LoRaWAN<sup>&reg;</sup> Network Server components can be setup in
multiple ways. For all cases, downloads can be found at the following URLs:

* [ChirpStack Gateway Bridge downloads](../../gateway-bridge/downloads.md)
* [ChirpStack Network Server downloads](../../network-server/downloads.md)
* [ChirpStack App Server downloads](../../application-server/downloads.md)

## Manual

In this case you need to download the pre-compiled binaries for each
component and setup scripts so that these components will be started on
boot (if desired).

For this you could use [systemd](https://en.wikipedia.org/wiki/Systemd)
unit-files or [init](https://en.wikipedia.org/wiki/Init) script, based
on the used Linux distribution.

## Debian / Ubuntu repository

ChirpStack provides Debian / Ubuntu `.deb` packages which
can be downloaded from the ChirpStack Deb repository. To guarantee
compatibility, for each major ChirpStack version a separate repository is provided.
Please refer to the [Debian / Ubuntu install guide](../guides/debian-ubuntu.md)
for a step-by-step guide.

## Docker

ChirpStack also provides Docker images which for example can be used with
[Docker Compose](https://docs.docker.com/compose/). See the
[Docker install](docker.md) for more information.
