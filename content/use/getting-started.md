---
title: Getting started
menu:
    main:
        parent: use
        weight: 1
---

## Getting started

After you succesfully setup the project, follow the steps below to get started
with receiving and sending from / to your node:

### Add a gateway

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

### Create an application and add a node

A node is always part of an application, therefore you first need to create
an application in [LoRa App Server](/lora-app-server/).
An application contains one or multiple nodes that have the same purpose, for
example a weather-station. Each node could then be a weather-station at a
different location.

After creating the application, you need to create the node in LoRa App Server.
When creating the node, you can choose between OTAA (over the air activation)
and ABP (activation by personalization). In the latter case you first create
the node and after creation, click the ABP actionvation button to activate
it in the network.

In case you have chosen your own `DevEUI`, `AppEUI` and `AppKey`, don't forget
to update these settings on your node.

Read more about the management of your nodes in the
[LoRa App Server](/lora-app-server/) documentation.

### Receive data

To receive data (and events) from your node(s), you need to subscribe to the
topic of your node and / or application. For this example we will be using
the MQTT client that comes with [Mosquitto](https://mosquitto.org). However,
every MQTT client will do. To subscribe to all data / events, subscribe to the
topic: `application/+/node/+/+`. The `+` is a wildcard (thus all applications,
nodes and events).

```bash
mosquitto_sub -v -t "application/+/node/+/+"
```

In case you configured your node for OTAA, perform a join (from your node).
You should see a join event being published over MQTT.

Read more more about sending and receiving data in the
[LoRa App Server](/lora-app-server/use/data/) documentation.

In case you don't see any data confirm (in the logs) that:

* [LoRa Gateway Bridge](/lora-gateway-bridge/) received data from your gateway
* [LoRa Server](/loraserver/) did not return any errors
* [LoRa App Server](/lora-app-server/) did not return any errors
