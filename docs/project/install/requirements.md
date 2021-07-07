---
description: Information about the ChirpStack LoRaWAN Network Server requirements and how to install these.
---

# Requirements

Before getting started with the ChirpStack LoRaWAN<sup>&reg;</sup> Network Server components,
there are a couple of requirements that needs to be satisfied. 

## MQTT broker

ChirpStack makes by default use of MQTT for publishing and receiving application
payloads. [Mosquitto](http://mosquitto.org/) is a popular open-source MQTT
server, but any MQTT broker implementing MQTT 3.1.1 should work.
In case you install Mosquitto, make sure you install a **recent** version.

MQTT is used by ChirpStack Gateway Bridge, ChirpStack Network Server, and ChirpStack Application Server.

### Install

#### Debian / Ubuntu

To install Mosquitto:

```bash
sudo apt install mosquitto
```

#### Other platforms

Please refer to the [Mosquitto download](https://mosquitto.org/download/) page
for information about how to setup Mosquitto for your platform.

## PostgreSQL database

The ChirpStack components are using [PostgreSQL](https://www.postgresql.org)
for persistent data-storage. Note that PostgreSQL 9.5+ is required and that
each component requires its own database to avoid schema conflicts. When
running multiple ChirpStack Network Server instances to support multiple LoRaWAN<sup>&reg;</sup> regions,
you must create a database for each region!

There is no need to run multiple PostgreSQL instances as a single instance
can host multiple databases.

PostgreSQL is used by ChirpStack Network Server and ChirpStack Application Server.

### Install

#### Debian / Ubuntu

To install the PostgreSQL:

```bash
sudo apt install postgresql
```

#### Other platforms

Please refer to the [PostgreSQL download](https://www.postgresql.org/download/)
page for information how to setup PostgreSQL on your platform.

## Redis database

The ChirpStack components are storing all non-persistent data into a
[Redis](http://redis.io/) datastore. Note that at least Redis 5.0.0
is required.

Redis is used by ChirpStack Network Server and ChirpStack Application Server.

### Install

#### Debian / Ubuntu

To Install Redis:

```bash
sudo apt install redis-server
```

#### Other platforms

Please refer to the [Redis](https://redis.io/) documentation for information
about how to setup Redis for your platform.
