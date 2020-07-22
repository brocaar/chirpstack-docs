---
description: Quickstart guide on hosting the ChirpStack Network Server components on Azure.
---

# Quickstart Microsoft Azure

This tutorial describes the steps needed to setup the ChirpStack Network Server project
on the [Azure Platform](https://azure.microsoft.com/). After completing
this guide, the following Azure service will be used:

* [IoT Hub](https://azure.microsoft.com/en-us/services/iot-hub/) is used to
  connect your LoRa<sup>&reg;</sup> gateways with the Azure platform.
* [Service Bus](https://azure.microsoft.com/en-us/services/service-bus/) is
  used for messaging between IoT Hub and ChirpStack Network Server (uplink).
* [Database for PostgreSQL](https://azure.microsoft.com/en-us/services/postgresql/)
  is used as hosted PostgreSQL service.
* [Cache for Redis](https://azure.microsoft.com/en-us/services/cache/) is used
  as hosted Redis solution.
* [Virtual Machines](https://azure.microsoft.com/en-us/services/virtual-machines/)
  is used for launching a VM instance to host the ChirpStack stack.

## Assumptions

* In this tutorial we will assume that the [ChirpStack Gateway Bridge](../../gateway-bridge/index.md)
  component will be installed on the gateway. 
* [ChirpStack Network Server](../../network-server/index.md) and [ChirpStack Application Server](../../application-server/index.md) will be installed
  on a single Virtual Machine instance to simplify this tutorial.
* The LoRaWAN<sup>&reg;</sup> region used in this tutorial will be EU868. Configuration
  examples for US915 are also given for US915.
* In this tutorial names are given for various entities. These can (and in some
  cases must) be replaced by something different (e.g. when they are already in use).
* When creating / configuring entities, suggestions are given for **some**
  of the fields, there might be mandatory fields missing. You should be fine
  to enter these without further instructions (e.g. usernames, passwords, ...).

## Requirements

* Azure account. You can create one [here](https://azure.microsoft.com/).
* LoRa gateway.
* LoRaWAN device.

## Create Service Bus

To create a Service Bus namespace, click **Create a resource**, select
**Service Bus** and click **Create**.

* **Name**: We will name this `chirpstack`.
* **Pricing tier**: For testing select **Basic**.
* **Location**: Select a location close to you.

Click **Create**.

After the Service Bus namespace has been created (this might take a minute
or two), click **Go To Resource** (hint: you can also search for `chirpstack`
in the search and then click on the resource link).

Under **Settings** click **Shared Access Policies**. Then click **+ Add**.

* **Policy Name:** We will name this `application-server`.
* **Manage:** Select this field.

Click **Create**.

Click **application-server**. Write down the following information:

* **Device Events Connection String:** The **Primary Connection String**.

### Create gateway events queue

This queue will be used by the IoT Hub to publish the received gateway events.

Under **Entities** click **Queues**. Then click **+Queue** to create a new queue.

* **Name:** We will name this `eu868-gateway-events`.

Click **Create**.

Then click **eu868-gateway-events**. Under **Settings** click
**Shared Access Policies** and click **+ Add**. 

* **Policy Name:** We name this `chirpstack`.
* **Listen:** select this field.

Click **Create**. 

Click **chirpstack**, then write down the following information:

* **Gateway Events Connection String:** The **Primary Connection String**.

### Create ChirpStack Application Server event queue

This queue will be used by ChirpStack Application Server to publish the device events.
Repeat the above steps to create an other queue named `device-events`.

## Create IoT Hub

The IoT Hub will be used by the gateway(s) to communicate with ChirpStack Network Server.
The gateway(s) connect to the IoT Hub using the IoT Hub MQTT interface.
Gateway events are written by IoT Hub to a Service Bus Queue.

To create an IoT Hub instance, click **Create a resource**, select **IoT Hub**
and click **Create**.

_Basics_

* **Region**: Select a region close to you.
* **IoT Hub Name**: We will name this `eu868-gateways` in this guide.

_Size and scale_

* **Pricing and scale tier**: For testing you can select the **F1: Free tier**.

Click **Review + create** then **Create**.

After the IoT Hub instance has been created (this might take a few minutes),
click **Go to resource** to open the overview and options. 

Under **Settings** click **Shared Access Policies** and click **+ Add**.

* **Access Policy Name:** We name this `chirpstack`.
* **Service Connect:** Select this field.

Click **Create**.

Click **chirpstack**, then write down the following information:

* **Gateway Commands Connection String:** The **Connection string - primary key**.

### Setup message routing

Under **Messaging** click **Message routing**. Click the **Custom endpoints**
tab and then **+Add > Service Bus Queue**. 

* **Endpoint name:** We name this endpoint `eu868-gateway-events`.
* **Service Bus Namespace:** Select **chirpstack-devel**.
* **Service Bus Queue:** Select **eu868-gateway-events**.

Click **Create**.

Click the **Routes** tab, then click **+Add**.

* **Name:** We name this route `eu868-gateway-events-route`.
* **Endpoint:** Select the **eu868-gateway-events** under **Service Bus Queues**.

Click **Create**. You can then click **Disable fallback route**.

### Add gateway to IoT Hub

Under **Explorers** (IoT Hub overview) click on the option **IoT Devices**.

Click **+ Add** then as **Device ID** enter the _Gateway ID_ (e.g. `0102030405060708`).
This must be entered in lowercase as the IoT Hub Device ID is case-sensitive.
You will find this value in your packet-forwarder configuration. Click **Save**.

Click on the created device (LoRa gateway) to obtain its **Connection string**.
This string looks like: `HostName=iot-hub-name.azure-devices.net;DeviceId=0102030405060708;SharedAccessKey=...`.
This **Connection string** will be needed in the next step.

### Configure ChirpStack Gateway Bridge

As there are different ways to install the [ChirpStack Gateway Bridge](../../gateway-bridge/index.md)
on your gateway, only the configuration is covered here. For installation
instructions, please refer to [ChirpStack Gateway Bridge gateway installation & configuration](../../gateway-bridge/gateway/index.md).

As ChirpStack Gateway Bridge will forwards its data to the Azure IoT Hub MQTT
interface, you need update the `chirpstack-gateway-bridge.toml` [Configuration file](../../gateway-bridge/install/config.md).
The `device_connection_string` needs to be replaced with the obtained
**Connection string**.

Minimal configuration example:

```toml
[integration.mqtt.auth]
type="azure_iot_hub"

  [integration.mqtt.auth.azure_iot_hub]
  device_connection_string="HostName=iot-hub-name.azure-devices.net;DeviceId=0102030405060708;SharedAccessKey=..."
```

In short:

* This will configure the Azure IoT Hub authentication
* Configures the **Connection string** so that ChirpStack Gateway Bridge knows how to
  connect to the IoT Hub.

After applying the above configuration changes, validate that the
ChirpStack Gateway Bridge connects to the Azure IoT Hub. On publishing events, the
log output should look like:

```text
INFO[0005] integration/mqtt: connected to mqtt broker   
INFO[0007] integration/mqtt: subscribing to topic        qos=0 topic="devices/00800000a00016b6/messages/devicebound/#"
INFO[0018] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0048] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0078] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0108] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0118] integration/mqtt: publishing event            event=up qos=0 topic=devices/00800000a00016b6/messages/events/up
```

In case you see something like below (on publishing an event), then this might
indicate that there is a configuration issue. Again, make sure that your
Gateway ID (which is also part of the **topic**) matches the IoT Hub
**Device ID** exactly.

```text
ERRO[0019] mqtt: connection error                        error=EOF
```

## Setup Redis database

Redis is used by both ChirpStack Network Server and ChirpStack Application Server for storing transient
data (session data, cache, etc...). 

To create a Redis instance, click **Create a resource**, select
**Azure Cache for Redis** and click **Create**.

* **DNS Name:** We will name this `chirpstack`.
* **Location:** Select a location close to you.
* **Pricing tier:** For testing, **Basic C0** should be sufficient.
* **Unblock port 6379 (not SSL encrypted):** Select this option.

Click **Create**. The deployment of the Redis instance might take a couple
of minutes. Click **Go to resource**.

Write down the following information:

* **Redis Host Name** (from the **Overview** page).
* **Redis Password** (from the **Settings > Acess Keys** page, either **Primary** or **Secondary**).

## Setup PostgreSQL databases

PostgreSQL is used by ChirpStack Network Server and ChirpStack Application Server for storing data that
must be persisted.

To create a PostgreSQL instance, click **Create a resource**, select
**Azure Database for PostgreSQL** and click **Create**. Under
**Single Server** click **Create**.

* **Server name:** We name this instance `chirpstack`.
* **Location:** Select a region close to you.
* **Version:** Select **10** (the latest version available from the dropdown).

Click **Configure server**.

For this tutorial, we scale down the PostgreSQL instance to its minimum.
You are free to select other options.

Click the **Basic** tab, slide the number of **vCores** to **1 vCore**,
slide the **Storage** to **5 GB** and click **OK**.

Click **Review + create** then click **Create**. The deployment of the
PostgreSQL instance might take a couple of minutes.

Once deployed, click **Go to resource**.

Under **Settings** click **Connection security**. Click **Allow access to
Azure services: ON** and click **Save**.

Write down the following information (from the **Overview** page):

* **PostgreSQL Server Name**
* **PostgreSQL Admin Username**

## Setup Virtual Machine

The Virtual Machine will be used to install ChirpStack Network Server and ChirpStack Application Server.
To keep this tutorial simple, we will use a single instance. However, you are
free to use other deployment options, including Kubernetes for example.

Under the default **Favorites** click **Virtual Machines**. Click **+ Add**.

_Instance details_

* **Virtual machine name:** We name this instance `chirpstack`.
* **Region:** Select a region close to you.
* **Image:** The default **Debian 9 Stretch** is a fine option.
* **Size:** For testing, select **Standard B1ls**.

_Inboud port rules_

* **Public inbound ports:** Select **Allow selected ports**.
* **Select inbound ports:** Select SSH (22).

Click **Review + create**, then click **Create**. The deployment of this
instance might take a few minutes. Then click **Go to resource**.

Under **Settings** click **Networking** and then click
**Add inbound port rule**. 

* **Source port range:** `*`.
* **Destination port ranges:** `8080` (this is the default ChirpStack Application Server port).

Click **Add**.

### Login using SSH

Under **Networking** (after the previous step, you are already on this page)
find the **NIC Public IP**.

Use SSH to login into your VM instance, example:

```bash
ssh [Administrator Username]@[NIC Public IP]
```

Depending you have setup SSH public key authentication or a password this step
might prompt you for a password.

### Initialize databases

In this step we are going to initialize the PostgreSQL database for ChirpStack Network Server
and ChirpStack Application Server.

First we need to install the PostgreSQL client utilities:

```bash
# update the apt cache
sudo apt update

# install PostgreSQL client
sudo apt install postgresql-client
```


Then we need to initialize the databases. Use the following command to enter
the PostgreSQL console:

```bash
psql -U '[PostgreSQL Admin Username]' -W -h [PostgreSQL Server Name] postgres
```

Inside this prompt, execute the following queries to set up the databases that
are used by the ChirpStack Network Server components. It is recommended to change the
passwords. Just remember to use these other passwords when updating
the `chirpstack-network-server.toml` and `chirpstack-application-server.toml` configuration files. Since
these two applications both use the same table to track database upgrades,
they must have separate databases.

```sql
-- set up the users and the passwords
-- (note that it is important to use single quotes and a semicolon at the end!)
create role chirpstack_as with login password 'dbpassword';
create role chirpstack_ns with login password 'dbpassword';

-- here the [PostgreSQL Admin Username] should only be the username before
-- the @, thus if this is admin@chirpstack, you must only use admin.
grant chirpstack_as to [PostgreSQL Admin Username];
grant chirpstack_ns to [PostgreSQL Admin Username];

-- create the database for the servers
create database chirpstack_as with owner chirpstack_as;
create database chirpstack_ns with owner chirpstack_ns;

-- change to the ChirpStack Application Server database
\c chirpstack_as

-- enable the pq_trgm and hstore extensions
-- (this is needed to facilidate the search feature)
create extension pg_trgm;
-- (this is needed to store additional k/v meta-data)
create extension hstore;

-- exit psql
\q
```

### Install ChirpStack Network Server

The following commands will be executed on the VM. After the previous step
you should still be connected using SSH.

```bash
# add required packages
sudo apt install apt-transport-https dirmngr

# import ChirpStack Network Server key
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

# add the repository to apt configuration
sudo echo "deb https://artifacts.chirpstack.io/packages/3.x/deb stable main" | sudo tee /etc/apt/sources.list.d/chirpstack.list

# update the package cache
sudo apt update

# install ChirpStack Network Server
sudo apt install chirpstack-network-server
```

### Configure ChirpStack Network Server

The ChirpStack Network Server configuration file is located at
`/etc/chirpstack-network-server/chirpstack-network-server.toml`. Below you will find two (minimal but working)
configuration examples. Please refer to the ChirpStack Network Server
[Configuration](/network-server/install/config/) documentation for all the
available options.

To test if there are no errors, you can execute the following command:


```bash
sudo chirpstack-network-server
```

This should output something like (it is important that there are no errors):

```text
INFO[0003] gateway/azure_iot_hub: setting up service-bus namespace 
INFO[0003] gateway/azure_iot_hub: starting queue consumer  queue=eu868-gateway-events
INFO[0003] no geolocation-server configured             
INFO[0003] configuring join-server client                ca_cert= server="http://localhost:8003" tls_cert= tls_key=
INFO[0003] gateway/azure_iot_hub: negotiating amqp cbs claim 
INFO[0006] starting api server                           bind="0.0.0.0:8000" ca-cert= tls-cert= tls-key=
INFO[0006] starting downlink device-queue scheduler     
INFO[0006] starting multicast scheduler                 
INFO[0010] gateway/azure_iot_hub: event received from gateway  event=stats gateway_id=00800000a00016b6
ERRO[0010] update gateway state error                    error="get gateway error: get gateway error: object does not exist"
INFO[0010] metrics saved                                 aggregation="[MINUTE HOUR DAY MONTH]" name="gw:00800000a00016b6"
```

If all is well (ignore the **get gateway error**), then you can start the
service in the background using:

```bash
sudo systemctl start chirpstack-network-server
```

#### Configuration examples

Important note for **[chirpstack_ns username]**: Remember that your
**[PostgreSQL Admin Username]** contained a _@_. In case this is 
**admin@chirpstack**, then replace **admin** with **chirpstack_ns**.
To continue this example, then the **chirpstack_ns username** will be
**chirpstack_ns**@chirpstack.

##### EU868 example

```toml
[postgresql]
dsn="postgres://[chirpstack_ns username]:[chirpstack_ns database password]@[PostgreSQL Server Name]/chirpstack_ns"

[redis]
url="redis://:[Redis Password]@[Redis Host Name]:6379"

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

  [network_server.gateway.backend]
  type="azure_iot_hub"

    [network_server.gateway.backend.azure_iot_hub]
    events_connection_string="[Gateway Events Connection String]"
    commands_connection_string="[Gateway Commands Connection String]"
```

##### US915 example sub-band 1 (125kHz channels 0 - 7 & 500kHz channel 64)

```toml
[postgresql]
dsn="postgres://[chirpstack_ns username]:[chirpstack_ns database password]@[PostgreSQL Server Name]/chirpstack_ns"

[redis]
url="redis://:[Redis Password]@[Redis Host Name]:6379"

[network_server]
net_id="000000"

  [network_server.band]
  name="US_902_928"

  [network_server.network_settings]
  enabled_uplink_channels=[0, 1, 2, 3, 4, 5, 6, 7, 64]

  [network_server.gateway.backend]
  type="azure_iot_hub"

    [network_server.gateway.backend.azure_iot_hub]
    events_connection_string="[Gateway Events Connection String]"
    commands_connection_string="[Gateway Commands Connection String]"
```

##### US915 example sub-band 2 (125kHz channels 8 - 15 & 500kHz channel 65)

```toml
[postgresql]
dsn="postgres://[chirpstack_ns username]:[chirpstack_ns database password]@[PostgreSQL Server Name]/chirpstack_ns"

[redis]
url="redis://:[Redis Password]@[Redis Host Name]:6379"

[network_server]
net_id="000000"

  [network_server.band]
  name="US_902_928"

  [network_server.network_settings]
  enabled_uplink_channels=[8, 9, 10, 11, 12, 13, 14, 15, 65]

  [network_server.gateway.backend]
  type="azure_iot_hub"

    [network_server.gateway.backend.azure_iot_hub]
    events_connection_string="[Gateway Events Connection String]"
    commands_connection_string="[Gateway Commands Connection String]"
```

### Install ChirpStack Application Server

The following commands will be executed on the VM. After the previous step
you should still be connected using SSH.

```bash
sudo apt install chirpstack-application-server
```

### Configure ChirpStack Application Server

The ChirpStack Application Server configuration file is located at
`/etc/chirpstack-application-server/chirpstack-application-server.toml`. Below you will find a minimal
but working configuration example. Please refer to the ChirpStack Application Server
[Configuration](/application-server/install/config/) documentation for all the
available options.

To test if there are no errors, you can execute the following command:

```bash
sudo chirpstack-application-server
```

This should output something like (it is important that there are no errors):

```text
INFO[0003] integration/azureservicebus: setting up namespace 
INFO[0003] integration/azureservicebus: testing if queue exists  queue=device-events
INFO[0004] api/as: starting application-server api       bind="0.0.0.0:8001" ca_cert= tls_cert= tls_key=
INFO[0004] api/external: starting api server             bind="0.0.0.0:8080" tls-cert= tls-key=
INFO[0005] api/external: registering rest api handler and documentation endpoint  path=/api
INFO[0005] api/js: starting join-server api              bind="0.0.0.0:8003" ca_cert= tls_cert= tls_key=
```

If all is well, then you can start the service in the background using:

```bash
sudo systemctl start chirpstack-application-server
```

#### Configuration example

Important note for **[chirpstack_as username]**: Remember that your
**[PostgreSQL Admin Username]** contained a _@_. In case this is
**admin@chirpstack**, then replace **admin** with **chirpstack_as**.
To continue this example, then the **chirpstack_as username** will be
**chirpstack_as@chirpstack.

```toml
[postgresql]
dsn="postgres://[chirpstack_as username]:[chirpstack_as database password]@[PostgreSQL Server Name]/chirpstack_as"

[redis]
url="redis://:[Redis Password]@[Redis Host Name]:6379"

[application_server]

  [application_server.integration]
  backend="azure_service_bus"

  [application_server.integration.azure_service_bus]
  connection_string="[Device Events Connection String]"
  publish_mode="queue"
  publish_name="device-events"

  [application_server.external_api]
  bind="0.0.0.0:8080"
  # You could generate this by executing 'openssl rand -base64 32'
  jwt_secret="[JWT Secret]"
```

## Getting started

### Setup your first gateway and device

To get started with the ChirpStack Network Server stack, please follow the [First gateway and device](first-gateway-device.md)
guide. It will explain how to login into the web-interface and add your first
gateway and device.

### Integrate your applications

We have configured ChirpStack Application Server to send data to an Azure Service Bus Queue.

For more information about Azure Service Bus, please refer to the following
pages:

* [Azure Service Bus product page](https://azure.microsoft.com/en-us/services/service-bus/)
* [Azure Service Bus documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/)
