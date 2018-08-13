---
title: Features
menu:
    main:
        parent: overview
        weight: 3
---

# Features

## ISM bands

LoRa Server provides support for all regions specified by
[LoRaWAN Regional Parameters](https://www.lora-alliance.org/For-Developers/LoRaWANDevelopers)
specification. Note that each region has different frequencies and settings on
which the LoRaWAN protocol operates, to comply to local regulations.

## Devices classes

LoRa Server implements LoRaWAN Class-A, Class-B and Class-C, making it ideal for
different purposes. Note that for all supported device-classes both
unconfirmed as confirmed data is supported (in the latter case an
acknowledgment will be sent by the receiving party).

### Class A

End-devices of Class A allow for bi-directional communications whereby each end-device‘s uplink transmission is followed by two short downlink receive windows. The transmission slot scheduled by the end-device is based on its own communication needs with a small variation based on a random time basis (ALOHA-type of protocol).

### Class B

End-devices of Class B allow for more receive slots. In addition to the Class A random receive windows, Class B devices open extra receive windows at scheduled times. In order for the End-device to open it receive window at the scheduled time it receives a time synchronized Beacon from the gateway.

### Class C

End-devices of Class C have nearly continuously open receive windows, only closed when transmitting. Class C end-device will use more power to operate than Class A or Class B but they offer the lowest latency for server to end-device communication.

## Adaptive data-rate (ADR)

LoRa Server implements adaptive data-rate. For devices that
support ADR, this will make sure that they are operating at the best data-rate
possible using the lowest possible transmission power.

## Channel re-configuration

For networks where only a sub-set of channels is used of the channels defined
by the LoRaWAN Regional Parameters (e.g. for the US region), LoRa Server
supports channel re-configuration. It will automatically disable channels
on the node that are not being used by the network.

## Device re-configuration

LoRa Server keeps a copy of network related parameters per device.  When it
detects that you have made changes, it will automatically signal these
configuration updates to your devices.

## Web-interface

LoRa App Server provides a web-interface to manage users, organizations,
applications, gateways and devices. It also provides debugging capabilities
like inspecting the LoRaWAN frames sent and received by gateways and
devices and inspecting the application payloads.

## APIs

Both LoRa Server (the network-server) and LoRa App Server
(the application-server) provide APIs to make them integratable in your own
infrastructure. By using the LoRa Server API, you could even implement your own
node inventory management system and fully replace LoRa App Server if needed.

## Gateway management

LoRa Server offers gateway management so that you are able to manage your
gateways and their GPS location and exposes statistics to track their 
performance.
