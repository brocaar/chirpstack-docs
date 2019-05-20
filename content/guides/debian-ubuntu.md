---
title: Quickstart Debian / Ubuntu
menu:
  main:
    parent: guides
    weight: 1
description: Quickstart tutorial on how to install all components on a Debian / Ubuntu system.
---

# Quickstart on Debian or Ubuntu

This tutorial describes the steps needed to setup the LoRa Server project
**including all requirements** on a single machine. It has been tested on
the following distributions:

* Ubuntu 16.04 LTS
* Ubuntu 18.04 LTS
* Debian 9 (Stretch)

Please refer to the other install pages for more generic installation
instructions.


## Assumptions

Many configurations of these packages is possible. Dependent software packages
could be installed on any number of remote servers, and the packages themselves
could be on their own servers. However, in order to simplify the initial
installation, we will assume the following deployment architecture:

* All LoRa Server components and their dependencies will be installed on a
   single server instance.
* The [LoRa Gateway Bridge](/lora-gateway-bridge/) component will be installed
   on the server, but can also be installed on the gateway itself.
* No firewall rules are setup.

Of course, optimizations may need to be made depending on the performance of
your systems. You may opt to move the PostgreSQL database to another server.
Or you may decide to put your MQTT server on a different system, or even use a
different server than the one recommended in this document. These and other
installation changes are beyond the scope of this document. However, you
should be able to find the information here that would make these changes
relatively straight-forward.

## Install requirements

* **MQTT broker** - A publish/subscribe protocol that allows users to publish
  information under topics that others can subscribe to. A popular
  implementation of the MQTT protocol is [Mosquitto](https://mosquitto.org/).
* **Redis** - An in-memory database used to store relatively transient data.
* **PostgreSQL** - The long-term storage database used by the open source
  packages.

Use the package manager `apt` to install these dependencies:

{{<highlight bash>}}
sudo apt install mosquitto mosquitto-clients redis-server redis-tools postgresql 
{{< /highlight >}}

### Setup PostgreSQL databases and users

To enter the command line utility for PostgreSQL:

{{<highlight bash>}}
sudo -u postgres psql
{{< /highlight >}}

Inside this prompt, execute the following queries to set up the databases
that are used by the LoRa Server components. It is recommended to change the
usernames and passwords. Just remember to use these other values when updating
the `loraserver.toml` and `lora-app-server.toml` configuration files. Since these
two applications both use the same table to track database upgrades, they must
have separate databases.

{{<highlight sql>}}
-- set up the users and the passwords
-- (note that it is important to use single quotes and a semicolon at the end!)
create role loraserver_as with login password 'dbpassword';
create role loraserver_ns with login password 'dbpassword';

-- create the database for the servers
create database loraserver_as with owner loraserver_as;
create database loraserver_ns with owner loraserver_ns;

-- change to the LoRa App Server database
\c loraserver_as

-- enable the pq_trgm extension
-- (this is needed to facilidate the search feature)
create extension pg_trgm;

-- exit psql
\q
{{< /highlight >}}

## Setup LoRa Server software repository

The LoRa Server project provides a repository that is compatible with the
Ubuntu apt package system. First make sure that both `dirmngr` and
`apt-transport-https` are installed:

{{<highlight bash>}}
sudo apt install apt-transport-https dirmngr
{{< /highlight >}}

Set up the key for this new repository:

{{<highlight bash>}}
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00
{{< /highlight >}}

Add the repository to the repository list by creating a new file:

{{<highlight bash>}}
sudo echo "deb https://artifacts.loraserver.io/packages/3.x/deb stable main" | sudo tee /etc/apt/sources.list.d/loraserver.list
{{< /highlight >}}

Update the apt package cache:

{{<highlight bash>}}
sudo apt update
{{< /highlight >}}

## Install LoRa Gateway Bridge

**Note:** when you intent to run the [LoRa Gateway Bridge](/lora-gateway-bridge/)
only on the gateways itself, you can skip this step.

Install the package using `apt`:

{{<highlight bash>}}
sudo apt install lora-gateway-bridge
{{< /highlight >}}

The configuration file is located at `/etc/lora-gateway-bridge/lora-gateway-bridge.toml`.
The default configuration is sufficient for this guide.

Start the LoRa Gateway Bridge service:

{{<highlight bash>}}
# start lora-gateway-bridge
sudo systemctl start lora-gateway-bridge

# start lora-gateway-bridge on boot
sudo systemctl enable lora-gateway-bridge
{{< /highlight >}}

## Installing LoRa Server

Install the package using apt:

{{<highlight bash>}}
sudo apt install loraserver
{{< /highlight >}}

The configuration file is located at `/etc/loraserver/loraserver.toml` and
must be updated to match the database and band configuration. See below
two examples for the EU868 and US915 band. For more information about all
the LoRa Server configuration options, see
[LoRa Server configuration](/loraserver/install/config/).

After updating the configuration, you need to restart LoRa Server and validate
that there are no errors.

Start the LoRa Server service:

{{<highlight bash>}}
# start loraserver
sudo systemctl start loraserver

# start loraserver on boot
sudo systemctl enable loraserver
{{< /highlight >}}

Print the LoRa Server log-output:

{{<highlight bash>}}
sudo journalctl -f -n 100 -u loraserver
{{< /highlight >}}

### EU868 configuration example

{{<highlight toml>}}
[general]
log_level=4

[postgresql]
dsn="postgres://loraserver_ns:dbpassword@localhost/loraserver_ns?sslmode=disable"

[network_server]
net_id="000000"

  [network_server.band]
  name="EU_863_870"

  [[network_server.network_settings.extra_channels]]
  frequency=867100000
  min_dr=0
  max_dr=5

  [[network_server.network_settings.extra_channels]]
  frequency=867300000
  min_dr=0
  max_dr=5

  [[network_server.network_settings.extra_channels]]
  frequency=867500000
  min_dr=0
  max_dr=5

  [[network_server.network_settings.extra_channels]]
  frequency=867700000
  min_dr=0
  max_dr=5

  [[network_server.network_settings.extra_channels]]
  frequency=867900000
  min_dr=0
  max_dr=5
{{< /highlight >}}

### US915 configuration example (channels 0 - 7)

{{<highlight toml>}}
[general]
log_level=4

[postgresql]
dsn="postgres://loraserver_ns:dbpassword@localhost/loraserver_ns?sslmode=disable"

[network_server]
net_id="000000"

[network_server.band]
name="US_902_928"

[network_server.network_settings]
enabled_uplink_channels=[0, 1, 2, 3, 4, 5, 6, 7]
{{< /highlight >}}

### US915 configuration example (channels 8 - 15, same as The Things Network)

{{<highlight toml>}}
[general]
log_level=4

[postgresql]
dsn="postgres://loraserver_ns:dbpassword@localhost/loraserver_ns?sslmode=disable"

[network_server]
net_id="000000"

[network_server.band]
name="US_902_928"

[network_server.network_settings]
enabled_uplink_channels=[8, 9, 10, 11, 12, 13, 14, 15]
{{< /highlight >}}

## Installing LoRa App Server

Install the package using apt:

{{<highlight bash>}}
sudo apt install lora-app-server
{{< /highlight >}}

The configuration file is located at `/etc/lora-app-server/lora-app-server.toml` and
must be updated to match the database configuration. See below a configuration
example which matches the database which we have created in one of the previous steps.
For more information about the LoRa App Server configuration options, see
[LoRa App Server configuration](/lora-app-server/install/config/).

{{<highlight toml>}}
[general]
log_level=4

[postgresql]
dsn="postgres://loraserver_as:dbpassword@localhost/loraserver_as?sslmode=disable"

  [application_server.external_api]
  jwt_secret="verysecret"
{{< /highlight >}}

**Note:** you **must** replace the `jwt_secret` with a secure secret!
You could use the command `openssl rand -base64 32` to generate a random secret.

Start the LoRa App Server service:

{{<highlight bash>}}
# start lora-app-server
sudo systemctl start lora-app-server

# start lora-app-server on boot
sudo systemctl enable lora-app-server
{{< /highlight >}}

Print the LoRa App Server log-output:

{{<highlight bash>}}
sudo journalctl -f -n 100 -u lora-app-server
{{< /highlight >}}

## Optional: install LoRa Gateway Bridge on the gateway

It is advised to run LoRa Gateway Bridge on each gateway itself, to enable a
secure connection between your gateways and your server.

As there are many types of gateways available, please refer to the
LoRa Gateway Bridge instructions for
[installing LoRa Gateway Bridge on the gateway](/lora-gateway-bridge/gateway/).

## Setting up your first device

To setup your first device, please refer to the [First gateway and device]({{<relref "first-gateway-device.md">}})
guide.
