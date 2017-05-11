---
title: Project
menu:
    main:
        parent: overview
        weight: 1
---

## The LoRa Server project

The LoRa Server project is an open-source LoRaWAN solution offering all
components needed to setup your own (private) LoRaWAN networks.
It consists of three components:

* [LoRa Gateway Bridge](/lora-gateway-bridge/) for connectivity with the
  gateway(s)
* [LoRa Server](/loraserver/) for maintaining the state of the nodes on the
  network
* [LoRa App Server](/lora-app-server/) for managing the "inventory" of nodes
  and providing a web-interface to the user and APIs to the end-application

See [Architecture]({{< relref "architecture.md" >}}) for more information
about each component.

### Features

#### ISM bands

LoRa Server implements all ISM bands specified by the LoRa Alliance LoRaWAN
[Regional Parameters 1.0](https://www.lora-alliance.org/For-Developers/LoRaWANDevelopers)
document.

#### Devices classes

LoRa Server implements both LoRaWAN Class-A and Class-C, making it ideal for
low-power devices sending data in (regular) intervals and devices which are
always turned on / listening for data. Both unconfirmed and confirmed data is
supported. Class-B support will be added in the future.

#### Adaptive data-rate (ADR)

LoRa Server implements adaptive data-rate for all ISM bands. For devices that
support ADR, this will make sure that they are operating at the best data-rate
possible with the least transmission power. This not only saves battery power,
but uses the network in the most ideal way.

#### Channel re-configuration

For networks where only a sub-set of channels is used of the channels defined
by the LoRaWAN Regional Parameters (e.g. for the US ism band), LoRa Server
supports channel re-configuration. It will automatically disable channels
on the node that are not being used.

#### Web-interface

LoRa App Server offers a web-interface to manage your nodes per application,
per organization. By being able to assigning users to organizations and / or
applications, The LoRa Server project is ideal for multi-team or
multi-organization setups.

#### API

Both LoRa Server (the network-server) and LoRa App Server
(the application-server) provide APIs to make them integratable in your own
infrastructure. By using the LoRa Server API, you could even implement your own
node inventory management system and fully replace LoRa App Server if needed.

#### Gateway management

LoRa Server offers gateway management so that you are able to manage your
gateways and their GPS location and exposes statistics to track their 
performance.