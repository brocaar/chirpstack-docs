---
title: Components
menu:
    main:
        parent: lorawan
        weight: 3
---

## LoRaWAN components

A typical LoRaWAN network has the following components:

### Nodes

The (low-power) devices / sensors communicating with an application over
the LoRaWAN network

### Gateways

Gateways are able to listen constantly on multiple channels and multiple
data-rates. The data they receive from nodes is forwarded to the
network-server, the data that is received from the network-server is
transmitted to the nodes.

### Network-server

The network-server manages the state of nodes active in the network and
validates the authenticity of the data that has been received. When a node
sends a join-request, it will as the application-server if the node is allowed
to join the network. [LoRa Server](https://docs.loraserver.io/loraserver/)
performs the role of network-server.

### Application-server

The application-server handles the data received by the network-server. It
is also responsible for encrypting / decrypting the application payload.
Note that only the application-server has the application session-key, so that
the network can not learn about what data is sent over the network.
[LoRa App Server](https://docs.loraserver.io/lora-app-server/) performs the
role of an application-server.