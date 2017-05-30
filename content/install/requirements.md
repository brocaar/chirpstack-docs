---
title: Requirements
menu:
    main:
        parent: install
        weight: 2
---

## Requirements

Before getting started with the LoRa Server project, there are a couple of
requirements that needs to be satisfied:

### MQTT broker

LoRa Server makes use of MQTT for publishing and receiving application
payloads. [Mosquitto](http://mosquitto.org/) is a popular open-source MQTT
server, but any MQTT broker implementing MQTT 3.1.1 should work.
In case you install Mosquitto, make sure you install a **recent** version.

MQTT is used by LoRa Gateway Bridge, LoRa Server, and Lora App Server.

#### Install

##### Debian / Ubuntu

For Ubuntu Trusty (14.04), execute the following command in order to add the
Mosquitto Apt repository, for Ubuntu Xenial and Debian Jessie you can skip
this step:

```bash
sudo apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
sudo apt-get update
```

In order to install Mosquitto, execute the following command:

```bash
sudo apt-get install mosquitto
```

##### Other platforms

Please refer to the [Mosquitto download](https://mosquitto.org/download/) page
for information about how to setup Mosquitto for your platform.

### PostgreSQL database

The LoRa Server components are using [PostgreSQL](https://www.postgresql.org)
for persistent data-storage. Note that PostgreSQL 9.5+ is required and that
each component requires its own database to avoid schema conflicts
(there is no need to run multiple database instances).

PostgreSQL is used by LoRa Server and LoRa App Server.

#### Install

##### Debian / Ubuntu

To install the latest PostgreSQL:

```bash
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

export DISTRIB_CODENAME=`lsb_release -sc`
sudo echo "deb http://apt.postgresql.org/pub/repos/apt/ ${DISTRIB_CODENAME}-pgdg main" | sudo tee /etc/apt/sources.list.d/pgdg.list
sudo apt-get update

sudo apt-get install postgresql-9.6
```

##### Other platforms

Please refer to the [PostgreSQL download](https://www.postgresql.org/download/)
page for information how to setup PostgreSQL on your platform.

### Redis database

The LoRa Server components are storing all non-persistent data into a
[Redis](http://redis.io/) datastore. Note that at least Redis 2.6.0
is required.

Redis is used by LoRa Server.

#### Install

##### Debian / Ubuntu

To Install Redis:

```bash
sudo apt-get install redis-server
```

##### Other platforms

Please refer to the [Redis](https://redis.io/) documentation for information
about how to setup Redis for your platform.
