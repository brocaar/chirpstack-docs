---
title: Getting started
menu:
    main:
        parent: use
        weight: 1
---

# Getting started

After you successfully setup the project, follow the steps below to get started
with receiving and sending from / to your node:

## Add a gateway

There are two steps involved when adding a gateway. First of all, you need
to configure your gateway so that it sends data to the
[LoRa Gateway Bridge](/lora-gateway-bridge/)
component. In the [packet-forwarder](https://github.com/Lora-net/packet_forwarder)
configuration, modify the following configuration keys:

* `server_address` to the IP address / hostname of the LoRa Gateway Bridge
* `serv_port_up` to `1700` (the default port that LoRa Gateway Bridge is using)
* `serv_port_down` to `1700` (same)

After restarting the packet-forwarder process, you should see log-lines
appearing in the LoRa Gateway Bridge logs.

The second step is to configure the gateway in your LoRa Server network. For
this, log in into the [LoRa App Server](/lora-app-server/)
web-interface and add the gateway to your organization. In case your gateway
does not have a GPS, you can set the location manually.

## Adding your first device

The following steps must be performed in the
[LoRa App Server](/lora-app-server/) web-interface.

### Associating the network-server

LoRa App Server must know to which network-server(s) to connect. Therefore
the first action is to add a network-server (LoRa Server instance)
to your LoRa App Server installation. See
[network-servers](/lora-app-server/use/network-servers/) for more information.

### Create a service-profile

To define what features can be used by users assigned to an organization
you must create one or multiple service-profiles for each organization.
See [service-profiles](/lora-app-server/use/service-profiles/) for more
information.

### Create a device-profile

To define the capabilities of the device you are going to add, you must
create one or multiple device-profiles for each organization. See
[device-profiles](/lora-app-server/use/device-profiles/) for more information.

### Create an application

A device is always part of an application, therefore you need to create
an application. An application contains one or multiple devices that serve the
same purpose, for example a weather-station. See
[applications](/lora-app-server/use/applications/) for more information.

### Create a device

After creating the application, you need to create the device and assign a
device-profile to it. After creating the device, don't forget to add the
application-key to it (OTAA) or activate the device (ABP).

## Receive data

To receive data (and events) from your devices(s), you need to subscribe to the
topic of your device and / or application. For this example we will be using
the MQTT client that comes with [Mosquitto](https://mosquitto.org). However,
every MQTT client will do. To subscribe to all data / events, subscribe to the
topic: `application/#`. The `#` is a milti-level wildcard (thus everything under
the `application/` prefix).

```bash
mosquitto_sub -v -t "application/#"
```

In case you configured your node for OTAA, perform a join (from your node).
You should see a join event being published over MQTT.

Read more more about sending and receiving data in the
[LoRa App Server](/lora-app-server/use/data/) documentation.

In case you don't see any data confirm (in the logs) that:

* [LoRa Gateway Bridge](/lora-gateway-bridge/) received data from your gateway
* [LoRa Server](/loraserver/) did not return any errors
* [LoRa App Server](/lora-app-server/) did not return any errors
