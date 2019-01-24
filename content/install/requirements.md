---
title: Requirements
menu:
    main:
        parent: install
        weight: 2
description: Information about the LoRa Server requirements and how to install these.
---

# Requirements

Before getting started with the LoRa Server project, there are a couple of
requirements that needs to be satisfied. 

## MQTT broker

LoRa Server makes use of MQTT for publishing and receiving application
payloads. [Mosquitto](http://mosquitto.org/) is a popular open-source MQTT
server, but any MQTT broker implementing MQTT 3.1.1 should work.
In case you install Mosquitto, make sure you install a **recent** version.

MQTT is used by LoRa Gateway Bridge, LoRa Server, and LoRa App Server.

### Install

#### Debian / Ubuntu

To install Mosquitto:

{{<highlight bash>}}
sudo apt-get install mosquitto
{{< /highlight >}}

#### Other platforms

Please refer to the [Mosquitto download](https://mosquitto.org/download/) page
for information about how to setup Mosquitto for your platform.

## PostgreSQL database

The LoRa Server components are using [PostgreSQL](https://www.postgresql.org)
for persistent data-storage. Note that PostgreSQL 9.5+ is required and that
each component requires its own database to avoid schema conflicts. When
running multiple LoRa Server instances to support multiple LoRaWAN regions,
you must create a database for each region!

There is no need to run multiple PostgreSQL instances as a single instance
can host multiple databases.

PostgreSQL is used by LoRa Server and LoRa App Server.

### Install

#### Debian / Ubuntu

To install the PostgreSQL:

{{<highlight bash>}}
sudo apt-get install postgresql
{{< /highlight >}}

#### Other platforms

Please refer to the [PostgreSQL download](https://www.postgresql.org/download/)
page for information how to setup PostgreSQL on your platform.

## Redis database

The LoRa Server components are storing all non-persistent data into a
[Redis](http://redis.io/) datastore. Note that at least Redis 2.6.0
is required.

Redis is used by LoRa Server.

### Install

#### Debian / Ubuntu

To Install Redis:

{{<highlight bash>}}
sudo apt-get install redis-server
{{< /highlight >}}

#### Other platforms

Please refer to the [Redis](https://redis.io/) documentation for information
about how to setup Redis for your platform.
