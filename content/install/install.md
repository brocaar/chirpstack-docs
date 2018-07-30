---
title: General Installation
menu:
    main:
        parent: install
        weight: 3
---

# Installing the LoRa Server project

There are multiple ways how the LoRa Server project could be setup:

## Manual

In this case you need to download the binaries for each component (matching
your platform and architecture) and make sure they are started on boot.
For this you could use [systemd](https://en.wikipedia.org/wiki/Systemd)
unit-files or [init](https://en.wikipedia.org/wiki/Init) script, based
on the platform you're using.

Pre-compiled binaries and packages can be found at:

* [LoRa Gateway Bridge downloads](/lora-gateway-bridge/overview/downloads/)
* [LoRa Server downloads](/loraserver/overview/downloads/)
* [LoRa App Server downloads](/lora-app-server/overview/downloads/)

## Debian / Ubuntu repository

As all packages are signed using a PGP key, you first need to import this key:

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00
```

### Distributions

#### Testing

The testing distribution contains the latest (test) releases.

```bash
sudo echo "deb https://artifacts.loraserver.io/packages/2.x/deb testing main" | sudo tee /etc/apt/sources.list.d/loraserver.list
sudo apt-get update
```

#### Stable

The stable distribution contains releases that are considered stable.

```bash
sudo echo "deb https://artifacts.loraserver.io/packages/2.x/deb stable main" | sudo tee /etc/apt/sources.list.d/loraserver.list
sudo apt-get update
```

### Installation

To install all components:

```bash
sudo apt-get install lora-gateway-bridge loraserver lora-app-server
```

Note that after installing, you still need to [configure]({{<relref "configuration.md">}}) each component.
