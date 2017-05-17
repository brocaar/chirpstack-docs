---
title: Features
menu:
    main:
        parent: overview
        weight: 3
---

## Features

### ISM bands

LoRa Server implements all regional bands that are defined by the
ISM bands specified by the LoRa Alliance LoRaWAN
[Regional Parameters 1.0](https://www.lora-alliance.org/For-Developers/LoRaWANDevelopers)
document. Note that each region has different frequencies and settings on
which the LoRaWAN protocol operates, to comply to local regulations.

### Devices classes

LoRa Server implements both LoRaWAN Class-A and Class-C, making it ideal for
different purposes. Note that for all supported device-classes both
unconfirmed as confirmed data is supported (in the latter case an
acknowledgment will be sent by the receiving party).

#### Class A

The node / device is always in sleep mode. Only when it has data to send
it wakes up to transmit, after which it opens two receive-windows in order
to receive a downlink transmission. After it received a downlink or after the
second receive window expired, it goes back to sleep.

#### Class B

Not implemented yet.

#### Class C

The node / devices is always listening, making it possible for the application
to send data to the node any time. When running in Class-C mode, a node will
use significantly more energy.

### Adaptive data-rate (ADR)

LoRa Server implements adaptive data-rate. For devices that
support ADR, this will make sure that they are operating at the best data-rate
possible with the least transmission power. This not only saves battery power,
but uses the network in the most ideal way as airtime of the node decreases.

### Channel re-configuration

For networks where only a sub-set of channels is used of the channels defined
by the LoRaWAN Regional Parameters (e.g. for the US region), LoRa Server
supports channel re-configuration. It will automatically disable channels
on the node that are not being used by the network.

### Web-interface

LoRa App Server offers a web-interface to manage your nodes per application,
per organization. By being able to assigning users to organizations and / or
applications, The LoRa Server project is ideal for multi-team or
multi-organization setups.

### API

Both LoRa Server (the network-server) and LoRa App Server
(the application-server) provide APIs to make them integratable in your own
infrastructure. By using the LoRa Server API, you could even implement your own
node inventory management system and fully replace LoRa App Server if needed.

### Gateway management

LoRa Server offers gateway management so that you are able to manage your
gateways and their GPS location and exposes statistics to track their 
performance.