---
title: Architecture
menu:
    main:
        parent: overview
        weight: 2
---

## System architecture

A LoRa Server architecture constists of multiple components:

![architecture](/img/architecture.png)

### LoRa nodes

The IoT devices or "nodes" (not pictured on the above image) are the devices
sending data to the LoRa network (through the LoRa gateways). These devices
could be for example sensors measturing air quality, temperature, humitidy,
location...

### LoRa gateway

The gateways are receiving data from the nodes and typtically runs an
implementation of the [packet-forwarder](https://github.com/Lora-net/packet_forwarder)
software. This software is responsible for the interface with the LoRa hardware
on the gateway.

### LoRa Gateway Bridge

The [LoRa Gateway Bridge](/lora-gateway-bridge/)
component is responsible for the communication with
your gateway. It "transforms" the packet-forwarder UDP protocol into JSON
over MQTT. The advantages over directly working with the UDP protocol are:

* It makes debugging easier
* Sending downlink data only requires knowledge about the corresponding MQTT
  topic of the gateway, the MQTT broker will route it to the LoRa Gateway
  Bridge instance responsible for the gateway
* It enables a secure connection between your gateways and the network
  (using MQTT over TLS)
* In the future, different bridge versions could handle different gateway
  protocol, so that the rest of the infrastructure only needs to know about the
  JSON over MQTT format

### LoRa Server

The [LoRa Server](/loraserver/) component is
responsible for the network. It knows about active
node sessions and when a new node joins the network, it will ask the
application-server if the node is allowed to your the network and if so,
which settings to use for this node.

For the active node-sessions, it de-duplicate the received data (which is
potentially received by multiple gateways), it authenticates this date (to
make sure that these are no replay-attacks), it forwards this (encrypted)
data to the application-server and it will ask the application-server if it
should send anything back.

Besides managing the data-flows, it also manages the state of the node through
so called mac-commands (e.g. to change the data-rate, channels, ...).

LoRa Server implements a gRPC API so that you could easily build your own
application-server.

### LoRa App Server

The [LoRa App Server](/lora-app-server/) component
implements an application-server compatible
with the LoRa Server component. It offers node management per application,
per organization and gateway management per organization. It also offers user
management and the possibility to assign users to organizations and / or
applications. Communication with the application is using JSON over MQTT and
using the exposed APIs.

LoRa App Servers offers a web-interface that can be used for gateway, node
and gateway management, but also offers API endpoints so that it can be
integrated with your own products.

### Application

The application subscribes to the MQTT topic for receiving the data of the
node(s) within the application. It is also able to send data back over MQTT.
If needed, it could also interact with LoRa App Server using the gRPC or
JSON REST api.
