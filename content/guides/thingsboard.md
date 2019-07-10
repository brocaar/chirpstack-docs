---
title: ThingsBoard getting started
menu:
  main:
    parent: guides
    weight: 4
description: Guide on getting started with LoRa Server and ThingsBoard.
---

# Getting started with ThingsBoard

[ThingsBoard](https://www.thingsboard.io) is an Open-source IoT Platform for
Device management, data collection, processing and visualization for your IoT
solution.

This guide describes how to setup LoRa Server so that it forwards device-data
to [ThingsBoard](https://www.thingsboard.io) for processing and visualization.

## Installing LoRa Server

The installation of LoRa Server is not covered in this guide. There are
different [Guides](/guides/) describing different deployment options. However,
the easiest option is the [Docker Compose]({{<relref "docker-compose.md">}})
guide, when you are not sure where to start.

## Setting up gateway and device

The steps to setup your first gateway and device are covered in the
[First Gateway and Device]({{<relref "first-gateway-device.md">}}) guide.

It is also recommended to setup a Payload Codec. A Payload Codec will decode
the received payload into "readable" information like `temperature: 21.5`.
This is important, as this is the information that will be forwarded to
ThingsBoard. The Payload Codec can be set up in the [Device Profile](/lora-app-server/use/device-profiles/).

Before you continue, make sure everything works up to this point.

## Install ThingsBoard

ThingsBoard can be installed in many different ways. Please refer to the
[ThingsBoard Installation Documentation](https://thingsboard.io/docs/installation/)
for more information.

### Docker Compose setup

When you have installed LoRa Server using the Docker Compose guilde, the easiest
option to integrate with ThingsBoard is by adding the Thiingsboard Docker
container to the LoRa Server `docker-compose.yml` file.

After making the modifications mentioned below, ThingsBoard will be started
together with LoRa Server when running a `docker-compose up`.

**Note:** The ThingsBoard server used in one of the next steps will be
`http://thingsboard:9090/`.

#### Service

Add the following snippet under `services`:

{{<highlight yaml>}}
  thingsboard:
    image: thingsboard/tb-postgres
    volumes:
      - thingsboarddata:/data
    ports:
      - 9090:9090
{{</highlight>}}

Your `docker-compose.yml` file now looks like:

{{<highlight yaml>}}
version: "3"

services:
  loraserver:
    image: loraserver/loraserver:3
    volumes:
      - ./configuration/loraserver:/etc/loraserver

  [...]

  thingsboard:
    image: thingsboard/tb-postgres
    volumes:
      - thingsboarddata:/data
    ports:
      - 9090:9090

  [...]
{{</highlight>}}

#### Volume

Add the following under `volumes`:

{{<highlight yaml>}}
  thingsboarddata:
{{</highlight>}}

Your `docker-compose.yml` file now looks like:

{{<highlight yaml>}}
[...]

volumes:
  postgresqldata:
  redisdata:
  thingsboarddata:
{{</highlight>}}

## Setup ThingsBoard

For getting started with ThingsBoard, please refer to the
[ThingsBoard Getting Started](https://thingsboard.io/docs/getting-started-guides/helloworld/)
guide. The important thing is that you have created Device within ThingsBoard.

## Integrate LoRa App Server with ThingsBoard

### Get Device Auth Token

In order to let LoRa App Server push data to your ThingsBoard device, you need
to obtain the ThingsBoard Device _Access Token_. Within ThingsBoard, open your
Device and click the **Copy Access Token** button. This will copy the
_Access Token_ to your clipboard.

### Set Device Auth Token in LoRa App Server

Now open the LoRa App Server web-interface and navigate to the Device. Click
**Configuration**, then click **Variables**.

Add a variable named **ThingsBoardAccessToken** and with as value the content
from your clipboard (containing the ThingsBoard Device _Access Token_).
This step is also documented in the LoRa App Server [ThingsBoard Integration](/lora-app-server/integrate/sending-receiving/thingsboard/)
documentation.

### Setup LoRa App Server ThingsBoard integration

Still in the LoRa App Server web-interface, navigate to the Application to
which the Device belongs. Click **Integrations**, then click **+ Create**.

* Integration kind: **ThingsBoard.io**.
* ThingsBoard.io server: Usually this is **http://host:9090** (where **host**
  is replaced by the hostname of the server serving ThingsBoard). When you are
  using the Docker Compose instructions, set this to **http://thingsboard:9090**.

## Validate integration

If you completed all the steps, then ThingsBoard is ready to receive uplink
data (or telemetry) and LoRa App Server is setup to forward data for your
Device, using the _Access Token_ for authentication.

The last step is to let your device send some data and validate that this data
is received by Thingsboard. You will find this data under the _Latest Telemetry_
tab when navigating to the Device within Thingsboard.
