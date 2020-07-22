---
description: Backend which uses the Google Cloud Platform Pub/Sub broker for communication between the LoRa gateways and the ChirpStack Network Server.
---

# GCP Pub/Sub

The [Google Cloud Platform](https://cloud.google.com/) [Pub/Sub](https://cloud.google.com/pubsub/)
backend uses a Pub/Sub queue for receiving gateway events and a Pub/Sub topic
for publishing gateway commands.

In order to connect the gateways (running the [ChirpStack Gateway Bridge](../../gateway-bridge/index.md)),
the [Cloud IoT Core](https://cloud.google.com/iot-core/) service is used, which
provides an [MQTT bridge](https://cloud.google.com/iot/docs/how-tos/mqtt-bridge).

Gateway events received by the Cloud IoT Core service are forwarded to a Pub/Sub
queue which is consumed by the ChirpStack Network Server instance or instances.

Downlink gateway commands are published by the ChirpStack Network Server to a
Pub/Sub topic. A [Cloud Function](https://cloud.google.com/functions/) then
calls the Cloud IoT API to forward the command to the ChirpStack Gateway Bridge
over MQTT.


## Architecture

[![architecture](/static/img/network-server/graphs/gcp_pub_sub.dot.png)](/static/img/network-server/graphs/gcp_pub_sub.dot.png)
