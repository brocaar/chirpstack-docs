---
description: Getting started with the ChirpStack Gateway OS.
---

# Getting started

These steps describe how to get started with ChirpStack Gateway OS **after** you
have installed ChirpStack Gateway OS on your gateway.

**Important:** The **chirpstack-gateway-os-full** image will setup the PostgreSQL
on its first boot. This could take a couple of minutes and during this time,
the gateway will be less responsive!

After booting the gateway, you need to login using SSH. In case the gateway
running the ChirpStack Gateway OS supports WIFI, then it will be configured
as access-point with the name `ChirpStackAP` and password `ChirpStackAP`.
Once connected with `ChirpStackAP` the IP of the gateway is `192.168.0.1`.

If you are connected using ethernet, then it uses DHCP to obtain an IP address.
Many internet routers provide a web-interface with the IP addresses of connected
devices.

If the IP of your gateway is `192.168.0.1`:

```bash
ssh admin@192.168.0.1
```

The default username is `admin`, the default password is `admin`.

This will prompt the following message:

```text
   ________    _           _____ __             __     _     
  / ____/ /_  (_)________ / ___// /_____ ______/ /__  (_)___ 
 / /   / __ \/ / ___/ __ \\__ \/ __/ __ `/ ___/ //_/ / / __ \
/ /___/ / / / / /  / /_/ /__/ / /_/ /_/ / /__/ ,< _ / / /_/ /
\____/_/ /_/_/_/  / .___/____/\__/\__,_/\___/_/|_(_)_/\____/ 
                 /_/

Documentation and copyright information:
> www.chirpstack.io

Commands:
> sudo gateway-config  - configure the gateway
> sudo monit status    - display service monitor
```

Then execute the `sudo gateway-config` to configure the channel-configuration
that the gateway must use.

## chirpstack-gateway-os-base

After the board and channel-plan have been configured, you must update the
ChirpStack Gateway Bridge configuration, such that it connects to your MQTT
broker. This can be done using the **Edit ChirpStack Gateway Bridge config**
option in the `gateway-config` menu.

Use the **Edit configuration file** option to edit the configuration file or
the **MQTT connection wizzard**. In case you are using (client-)certificate
authentication / authorization, the latter is recommended as it allows you
do directly paste the certificate files.

## chirpstack-gateway-os-full

Unlike the **chirpstack-gateway-os-base** image, you **should not** update the
ChirpStack Gateway Bridge configuration. It is configured to point to the MQTT broker
which comes with the **chirpstack-gateway-os-full** image.

When using the **chirpstack-gateway-os-full** image and after configuring the
channel-plan, proceed with the following documentation to connect your first
gateway and device:

* [Connecting a gateway](../../project/guides/connect-gateway.md)
* [Connecting a device](../../project/guides/connect-device.md)

## Important to know

### SD Card wearout

Although ChirpStack Network Server and ChirpStack Application Server will try to
minimize the number of database writes, there will be regular writes to
the SD Card (PostgreSQL and Redis snapshots).
According to [Is it true that a SD/MMC Card does wear levelling with its own controller?](https://electronics.stackexchange.com/questions/27619/is-it-true-that-a-sd-mmc-card-does-wear-levelling-with-its-own-controller)
it might make a difference which SD Card brand you use.
