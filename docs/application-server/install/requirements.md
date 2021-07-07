---
description: Instructions how to setup the ChirpStack Application Server requirements.
---

# Requirements


## MQTT broker

ChirpStack Application Server makes use of MQTT for publishing and receiving application
payloads. [Mosquitto](http://mosquitto.org/) is a popular open-source MQTT
server, but any MQTT broker implementing MQTT 3.1.1 should work.
In case you install Mosquitto, make sure you install a **recent** version.

### Install

#### Debian / Ubuntu

In order to install Mosquitto, execute the following command:

```bash
sudo apt install mosquitto
```

#### Other platforms

Please refer to the [Mosquitto download](https://mosquitto.org/download/) page
for information about how to setup Mosquitto for your platform.

## PostgreSQL database

ChirpStack Application Server persists the gateway data into a
[PostgreSQL](https://www.postgresql.org) database. Note that PostgreSQL 9.5+
is required.

### pq_trgm and hstore extension

You also need to enable the [pg_trgm](https://www.postgresql.org/docs/current/static/pgtrgm.html)
(trigram) and [hstore](https://www.postgresql.org/docs/current/hstore.html)
extensions. Example to enable this extension (assuming your ChirpStack Application Server
database is named `chirpstack_as`):

Start the PostgreSQL prompt as the `postgres` user:

```bash
sudo -u postgres psql
```

Within the PostgreSQL prompt, enter the following queries:

```sql
-- change to the ChirpStack Application Server database
\c chirpstack_as

-- enable the extensions
create extension pg_trgm;
create extension hstore;

-- exit the prompt
\q
```

### Install

#### Debian / Ubuntu

To install the latest PostgreSQL:

```bash
sudo apt-get install postgresql
```

#### Other platforms

Please refer to the [PostgreSQL download](https://www.postgresql.org/download/)
page for information how to setup PostgreSQL on your platform.

## Redis database

ChirpStack Application Server stores all non-persistent data into a
[Redis](http://redis.io/) datastore. Note that at least Redis 5.0.0
is required.

### Install

#### Debian / Ubuntu

To Install Redis:

```bash
sudo apt-get install redis-server
```

#### Other platforms

Please refer to the [Redis](https://redis.io/) documentation for information
about how to setup Redis for your platform.
