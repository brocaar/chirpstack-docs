---
title: Quickstart Microsoft Azure
menu:
  main:
    parent: guides
    weight: 1
description: Quickstart guide on hosting the LoRa Server components on Azure.
---

# Quickstart Microsoft Azure

This tutorial describes the steps needed to setup the LoRa Server project
on the [Azure Platform](https://azure.microsoft.com/). After completing
this guide, the following Azure service will be used:

* [IoT Hub](https://azure.microsoft.com/en-us/services/iot-hub/) is used to
  connect your LoRa gateways with the Azure platform.
* [Service Bus](https://azure.microsoft.com/en-us/services/service-bus/) is
  used for messaging between IoT Hub and LoRa Server (uplink).
* [Database for PostgreSQL](https://azure.microsoft.com/en-us/services/postgresql/)
  is used as hosted PostgreSQL service.
* [Cache for Redis](https://azure.microsoft.com/en-us/services/cache/) is used
  as hosted Redis solution.
* [Virtual Machines](https://azure.microsoft.com/en-us/services/virtual-machines/)
  is used for launching a VM instance to host LoRa (App) Server.

## Assumptions

* In this tutorial we will assume that the [LoRa Gateway Bridge](/lora-gateway-bridge/)
  component will be installed on the gateway. 
* [LoRa Server](/loraserver/) and [LoRa App Server](/lora-app-server) will be installed
  on a single Virtual Machine instance to simplify this tutorial.
* The LoRaWAN region used in this tutorial will be EU868. Configuration
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

* **Name**: We will name this `loraserver`.
* **Pricing tier**: For testing select **Basic**.
* **Location**: Select a location close to you.

Click **Create**.

After the Service Bus namespace has been created (this might take a minute
or two), click **Go To Resource** (hint: you can also search for `loraserver`
in the search and then click on the resource link).

Under **Settings** click **Shared Access Policies**. Then click **+ Add**.

* **Policy Name:** We will name this `lora-app-server`.
* **Manage:** Select this field.

Click **Create**.

Click **lora-app-server**. Write down the following information:

* **Device Events Connection String:** The **Primary Connection String**.

### Create gateway events queue

This queue will be used by the IoT Hub to publish the received gateway events.

Under **Entities** click **Queues**. Then click **+Queue** to create a new queue.

* **Name:** We will name this `eu868-gateway-events`.

Click **Create**.

Then click **eu868-gateway-events**. Under **Settings** click
**Shared Access Policies** and click **+ Add**. 

* **Policy Name:** We name this `loraserver`.
* **Listen:** select this field.

Click **Create**. 

Click **loraserver**, then write down the following information:

* **Gateway Events Connection String:** The **Primary Connection String**.

### Create LoRa App Server event queue

This queue will be used by LoRa App Server to publish the device events.
Repeat the above steps to create an other queue named `device-events`.

## Create IoT Hub

The IoT Hub will be used by the gateway(s) to communicate with LoRa Server.
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

* **Access Policy Name:** We name this `loraserver`.
* **Service Connect:** Select this field.

Click **Create**.

Click **loraserver**, then write down the following information:

* **Gateway Commands Connection String:** The **Connection string - primary key**.

### Setup message routing

Under **Messaging** click **Message routing**. Click the **Custom endpoints**
tab and then **+Add > Service Bus Queue**. 

* **Endpoint name:** We name this endpoint `eu868-gateway-events`.
* **Service Bus Namespace:** Select **loraserver-devel**.
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

### Configure LoRa Gateway Bridge

As there are different ways to install the [LoRa Gateway Bridge](/lora-gateway-bridge/)
on your gateway, only the configuration is covered here. For installation
instructions, please refer to [LoRa Gateway Bridge gateway installation & configuration](/lora-gateway-bridge/install/gateway/).

As LoRa Gateway Bridge will forwards its data to the Azure IoT Hub MQTT
interface, you need update the `lora-gateway-bridge.toml` [Configuration file](/lora-gateway-bridge/install/config/).
The `device_connection_string` needs to be replaced with the obtained
**Connection string**.

Minimal configuration example:

{{<highlight toml>}}
[integration.mqtt.auth]
type="azure_iot_hub"

  [integration.mqtt.auth.azure_iot_hub]
  device_connection_string="HostName=iot-hub-name.azure-devices.net;DeviceId=0102030405060708;SharedAccessKey=..."
{{</highlight>}}

In short:

* This will configure the Azure IoT Hub authentication
* Configures the **Connection string** so that LoRa Gateway Bridge knows how to
  connect to the IoT Hub.

After applying the above configuration changes, validate that the
LoRa Gateway Bridge connects to the Azure IoT Hub. On publishing events, the
log output should look like:

{{<highlight text>}}
INFO[0005] integration/mqtt: connected to mqtt broker   
INFO[0007] integration/mqtt: subscribing to topic        qos=0 topic="devices/00800000a00016b6/messages/devicebound/#"
INFO[0018] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0048] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0078] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0108] integration/mqtt: publishing event            event=stats qos=0 topic=devices/00800000a00016b6/messages/events/stats
INFO[0118] integration/mqtt: publishing event            event=up qos=0 topic=devices/00800000a00016b6/messages/events/up
{{</highlight>}}

In case you see something like below (on publishing an event), then this might
indicate that there is a configuration issue. Again, make sure that your
Gateway ID (which is also part of the **topic**) matches the IoT Hub
**Device ID** exactly.

{{<highlight text>}}
ERRO[0019] mqtt: connection error                        error=EOF
{{</highlight>}}

## Setup Redis database

Redis is used by both LoRa Server and LoRa App Server for storing transient
data (session data, cache, etc...). 

To create a Redis instance, click **Create a resource**, select
**Azure Cache for Redis** and click **Create**.

* **DNS Name:** We will name this `loraserver`.
* **Location:** Select a location close to you.
* **Pricing tier:** For testing, **Basic C0** should be sufficient.
* **Unblock port 6379 (not SSL encrypted):** Select this option.

Click **Create**. The deployment of the Redis instance might take a couple
of minutes. Click **Go to resource**.

Write down the following information:

* **Redis Host Name** (from the **Overview** page).
* **Redis Password** (from the **Settings > Acess Keys** page, either **Primary** or **Secondary**).

## Setup PostgreSQL databases

PostgreSQL is used by LoRa Server and LoRa App Server for storing data that
must be persisted.

To create a PostgreSQL instance, click **Create a resource**, select
**Azure Database for PostgreSQL** and click **Create**. Under
**Single Server** click **Create**.

* **Server name:** We name this instance `loraserver`.
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

The Virtual Machine will be used to install LoRa Server and LoRa App Server.
To keep this tutorial simple, we will use a single instance. However, you are
free to use other deployment options, including Kubernetes for example.

Under the default **Favorites** click **Virtual Machines**. Click **+ Add**.

_Instance details_

* **Virtual machine name:** We name this instance `loraserver`.
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
* **Destination port ranges:** `8080` (this is the default LoRa App Server port).

Click **Add**.

### Login using SSH

Under **Networking** (after the previous step, you are already on this page)
find the **NIC Public IP**.

Use SSH to login into your VM instance, example:

{{<highlight bash>}}
ssh [Administrator Username]@[NIC Public IP]
{{</highlight>}}

Depending you have setup SSH public key authentication or a password this step
might prompt you for a password.

### Initialize databases

In this step we are going to initialize the PostgreSQL database for LoRa Server
and LoRa App Server.

First we need to install the PostgreSQL client utilities:

{{<highlight bash>}}
# update the apt cache
sudo apt update

# install PostgreSQL client
sudo apt install postgresql-client
{{</highlight>}}


Then we need to initialize the databases. Use the following command to enter
the PostgreSQL console:

{{<highlight bash>}}
psql -U '[PostgreSQL Admin Username]' -W -h [PostgreSQL Server Name] postgres
{{</highlight>}}

Inside this prompt, execute the following queries to set up the databases that
are used by the LoRa Server components. It is recommended to change the
passwords. Just remember to use these other passwords when updating
the `loraserver.toml` and `lora-app-server.toml` configuration files. Since
these two applications both use the same table to track database upgrades,
they must have separate databases.

{{<highlight sql>}}
-- set up the users and the passwords
-- (note that it is important to use single quotes and a semicolon at the end!)
create role loraserver_as with login password 'dbpassword';
create role loraserver_ns with login password 'dbpassword';

-- here the [PostgreSQL Admin Username] should only be the username before
-- the @, thus if this is admin@loraserver, you must only use admin.
grant loraserver_as to [PostgreSQL Admin Username];
grant loraserver_ns to [PostgreSQL Admin Username];

-- create the database for the servers
create database loraserver_as with owner loraserver_as;
create database loraserver_ns with owner loraserver_ns;

-- change to the LoRa App Server database
\c loraserver_as

-- enable the pq_trgm and hstore extensions
-- (this is needed to facilidate the search feature)
create extension pg_trgm;
-- (this is needed to store additional k/v meta-data)
create extension hstore;

-- exit psql
\q
{{</highlight>}}

### Install LoRa Server

The following commands will be executed on the VM. After the previous step
you should still be connected using SSH.

{{<highlight bash>}}
# add required packages
sudo apt install apt-transport-https dirmngr

# import LoRa Server key
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

# add the repository to apt configuration
sudo echo "deb https://artifacts.loraserver.io/packages/3.x/deb stable main" | sudo tee /etc/apt/sources.list.d/loraserver.list

# update the package cache
sudo apt update

# install loraserver
sudo apt install loraserver
{{< /highlight >}}

### Configure LoRa Server

The LoRa Server configuration file is located at
`/etc/loraserver/loraserver.toml`. Below you will find two (minimal but working)
configuration examples. Please refer to the LoRa Server
[Configuration](/loraserver/install/config/) documentation for all the
available options.

To test if there are no errors, you can execute the following command:


{{<highlight bash>}}
sudo loraserver
{{< /highlight >}}

This should output something like (it is important that there are no errors):

{{<highlight text>}}
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
{{</highlight>}}

If all is well (ignore the **get gateway error**), then you can start the
service in the background using:

{{<highlight bash>}}
sudo systemctl start loraserver
{{< /highlight >}}

#### Configuration examples

Important note for **[loraserver_ns username]**: Remember that your
**[PostgreSQL Admin Username]** contained a _@_. In case this is 
**admin@loraserver**, then replace **admin** with **loraserver_ns**.
To continue this example, then the **loraserver_ns username** will be
**loraserver_ns**@loraserver.

##### EU868 example

{{<highlight toml>}}
[postgresql]
dsn="postgres://[loraserver_ns username]:[loraserver_ns database password]@[PostgreSQL Server Name]/loraserver_ns"

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
{{</highlight>}}

##### US915 example (channels 0 - 7)

{{<highlight toml>}}
[postgresql]
dsn="postgres://[loraserver_ns username]:[loraserver_ns database password]@[PostgreSQL Server Name]/loraserver_ns"

[redis]
url="redis://:[Redis Password]@[Redis Host Name]:6379"

[network_server]
net_id="000000"

  [network_server.band]
  name="US_902_928"

  [network_server.network_settings]
  enabled_uplink_channels=[0, 1, 2, 3, 4, 5, 6, 7]

  [network_server.gateway.backend]
  type="azure_iot_hub"

    [network_server.gateway.backend.azure_iot_hub]
    events_connection_string="[Gateway Events Connection String]"
    commands_connection_string="[Gateway Commands Connection String]"
{{</highlight>}}

##### US915 example (channels 8 - 15)

{{<highlight toml>}}
[postgresql]
dsn="postgres://[loraserver_ns username]:[loraserver_ns database password]@[PostgreSQL Server Name]/loraserver_ns"

[redis]
url="redis://:[Redis Password]@[Redis Host Name]:6379"

[network_server]
net_id="000000"

  [network_server.band]
  name="US_902_928"

  [network_server.network_settings]
  enabled_uplink_channels=[8, 9, 10, 11, 12, 13, 14, 15]

  [network_server.gateway.backend]
  type="azure_iot_hub"

    [network_server.gateway.backend.azure_iot_hub]
    events_connection_string="[Gateway Events Connection String]"
    commands_connection_string="[Gateway Commands Connection String]"
{{</highlight>}}

### Install LoRa App Server

The following commands will be executed on the VM. After the previous step
you should still be connected using SSH.

{{<highlight bash>}}
sudo apt install lora-app-server
{{</highlight>}}

### Configure LoRa App Server

The LoRa App Server configuration file is located at
`/etc/lora-app-server/lora-app-server.toml`. Below you will find a minimal
but working configuration example. Please refer to the LoRa App Server
[Configuration](/lora-app-server/install/config/) documentation for all the
available options.

To test if there are no errors, you can execute the following command:

{{<highlight bash>}}
sudo lora-app-server
{{< /highlight >}}

This should output something like (it is important that there are no errors):

{{<highlight text>}}
INFO[0003] integration/azureservicebus: setting up namespace 
INFO[0003] integration/azureservicebus: testing if queue exists  queue=device-events
INFO[0004] api/as: starting application-server api       bind="0.0.0.0:8001" ca_cert= tls_cert= tls_key=
INFO[0004] api/external: starting api server             bind="0.0.0.0:8080" tls-cert= tls-key=
INFO[0005] api/external: registering rest api handler and documentation endpoint  path=/api
INFO[0005] api/js: starting join-server api              bind="0.0.0.0:8003" ca_cert= tls_cert= tls_key=
{{< /highlight >}}

If all is well, then you can start the service in the background using:

{{<highlight bash>}}
sudo systemctl start lora-app-server
{{< /highlight >}}

#### Configuration example

Important note for **[loraserver_as username]**: Remember that your
**[PostgreSQL Admin Username]** contained a _@_. In case this is
**admin@loraserver**, then replace **admin** with **loraserver_as**.
To continue this example, then the **loraserver_as username** will be
**loraserver_as@loraserver.

{{<highlight toml>}}
[postgresql]
dsn="postgres://[loraserver_as username]:[loraserver_as database password]@[PostgreSQL Server Name]/loraserver_as"

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
{{< /highlight >}}

## Getting started

### Setup your first gateway and device

To get started with LoRa (App) Server, please follow the [First gateway and device]({{<relref "first-gateway-device.md">}})
guide. It will explain how to login into the web-interface and add your first
gateway and device.

### Integrate your applications

We have configured LoRa App Server to send data to an Azure Service Bus Queue.

For more information about Azure Service Bus, please refer to the following
pages:

* [Azure Service Bus product page](https://azure.microsoft.com/en-us/services/service-bus/)
* [Azure Service Bus documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/)
