---
title: Architecture
menu:
    main:
        parent: overview
        weight: 2
---

# System architecture

A LoRa Server architecture consists of multiple components:

![architecture](/img/architecture.png)

## LoRaWAN devices

The IoT devices (not pictured in the image above) are the devices
sending data to the LoRa network (through the LoRa gateways). These devices
could be for example sensors measuring air quality, temperature, humidity,
location...

## LoRa gateway

The gateways are receiving data from the devices and typically run an
implementation of the [packet-forwarder](https://github.com/Lora-net/packet_forwarder)
software. This software is responsible for the interface with the LoRa hardware
on the gateway.

## LoRa Gateway Bridge

The [LoRa Gateway Bridge](/lora-gateway-bridge/)
component is responsible for the communication with
your gateway. It "transforms" the packet-forwarder UDP protocol into messages
over MQTT. The advantages over directly working with the UDP protocol are:

* It makes debugging easier
* Sending downlink data only requires knowledge about the corresponding MQTT
  topic of the gateway, the MQTT broker will route it to the LoRa Gateway
  Bridge instance responsible for the gateway
* It enables a secure connection between your gateways and the network
  (using MQTT over TLS)

## LoRa Server

The [LoRa Server](/loraserver/) component provides the LoRaWAN network-server
component, responsible for managing the state of the network.
It has knowledge of devices active on the network and is able to handle
join-requests when devices want to join the network. 

When data is received by multiple gateways, LoRa Server will de-duplicate
this data and forward it once to the LoRaWAN application-server. When an
application-server needs to send data back to a device, LoRa Server will
keep these items in queue, until it is able to send to one of the gateways.

LoRa Server provides an API which can be used for integration or when
implementing your own application-server.

## LoRa Geo Server

The [LoRa Geo Server](/lora-geo-server/) component (optional) provides
geolocation services for resolving location of each device.

## LoRa App Server

The [LoRa App Server](/lora-app-server/) component
implements a LoRaWAN application-server compatible
with the LoRa Server component. It provides a web-interface and APIs for
management of users, organizations, applications, gateways and devices.

To receive the application payload sent by one of your devices, you can
use one of the integrations that LoRa App Server provides (e.g. MQTT, HTTP
or directly write to an InfluxDB database).

## Application

The end-application (not provided) handles the application-payloads sent by
your devices. It receives this data from LoRa App Server, using one of the
possible integrations.
