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

## Debian / Ubuntu

The LoRa Server project provides pre-compiled binaries packaged as Debian (.deb)
packages. In order to activate this repository, execute the following
commands:

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

export DISTRIB_ID=`lsb_release -si`
export DISTRIB_CODENAME=`lsb_release -sc`
sudo echo "deb https://repos.loraserver.io/${DISTRIB_ID,,} ${DISTRIB_CODENAME} testing" | sudo tee /etc/apt/sources.list.d/loraserver.list
sudo apt-get update
```

Then to install all components:

```bash
sudo apt-get install lora-gateway-bridge loraserver lora-app-server
```

Note that after installing, you still need to configure each component.
Configuration files are located at `/etc/NAME/NAME.toml` where `NAME` must 
be substituted by the component name.

Please refer to the documentation of each component for more details about
setting up and configuration:

* [LoRa Gateway Bridge](/lora-gateway-bridge/)
* [LoRa Server](/loraserver/)
* [LoRa App Server](/lora-app-server/)
