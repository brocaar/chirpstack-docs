---
title: Project
menu:
    main:
        parent: overview
        weight: 1
---

# The LoRa Server project

## What is LoRa?

LoRaWAN is a long range, low power wireless protocol that is intended for use
in building IoT networks.  IoT devices ("nodes") send small data packets to
any number of "gateways" that may be in the several-kilometer range of a node
via the LoRaWAN wireless protocol. The gateways then use more traditional
communications such as wired Internet connections to forward the messages
to a network-server which validates the packets and forwards the application
payload to an application-server.

The nature of the LoRa network potentially allows IoT devices to run for years
on small batteries, occasionally sending out small packets of data, waiting for
a short time for response messages, and then closing the connection until more
data needs to be sent. Devices can also be set up so that they are always
listening for messages from their applications, though this obviously requires
more power and may be more appropriate for devices that are, say, plugged in
to a wall socket.

Of course there is much more to LoRaWAN than is described here. The LoRaWAN
protocol is defined and managed by the [LoRa Alliance](https://www.lora-alliance.org/).
There is a great deal of information available there.

## About the Lora Server Project

The Lora Server project is an open-source set of applications that fill the
gap between the gateways receiving messages from the nodes to just before the
applications receiving the data. It provides mechanisms for managing the
gateways on the LoRa network, the applications supported, and the devices
associated with the applications.

The project is designed so that it may be used in a very flexible manner.
For example the [LoRa App Server](/lora-app-server/) component implements
the application-server component and offers a Web UI for users to access and
modify their gateways, applications and nodes. The system can also be accessed
via programmatic interfaces implemented in [gRPC](http://www.grpc.io/) and
JSON REST APIs. Further, the APIs are designed such that the subsystems may
be replaced by other software implementing the same interfaces.

For a more technical understanding of the parts of the Lora Server software
system and how they work together, please refer to the
[architecture]({{< relref "architecture.md" >}}) page.
