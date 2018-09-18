---
title: Features
menu:
    main:
        parent: overview
        weight: 3
description: Overview of the features provided by the LoRa Server project.
---

# Features

## Gateway connectivity

Gateway connectivity (over MQTT) is provided by the [LoRa Gateway Bridge](/lora-gateway-bridge/)
component. It abstracts the [Semtech packet-forwarder](https://github.com/lora-net/packet_forwarder)
into messages published over MQTT.

## Network-server

All LoRaWAN related features are provided by the [LoRa Server](/loraserver/)
component. Please refer to the [LoRa Server features](/loraserver/features/)
page for a complete overview and documentation of all the implemented features.

## Geolocation-server

[LoRa Geo Server](/lora-geo-server/) provides a LoRaWAN geolocation server.
For more information about its features, please refer to the
[LoRa Geo Server](/lora-geo-server/) documentation page.

## Application-server

[LoRa App Server](/lora-app-server/) provides the LoRaWAN application-server
features. It also provides a web-interface for the management and configuratio of
network-servers, users, gateways and devices. Please refer to the
[Using LoRa App Server](/lora-app-server/use/) documentation for more information.
