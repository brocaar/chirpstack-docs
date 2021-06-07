---
description: Installation of the ChirpStack Gateway Bridge on Kerlink gateways.
---

# Kerlink

## Kerlink KerOS based gateways

KerOS is used by Kerlink on the following gateways:

* [Kerlink iBTS](https://www.kerlink.com/product/wirnet-ibts/)
* [Kerlink iFemtoCell](https://www.kerlink.com/product/wirnet-ifemtocell/)
* [Kerlink iStation](https://www.kerlink.com/product/wirnet-istation/)

### Requirements

Please make sure that you have installed the latest KerOS
firmware version. The steps below have been tested with KerOS v4.3.3. For
upgrade instructions, please refer to the [Kerlink wiki](http://wikikerlink.fr/wirnet-productline).

### SSH into the gateway

The first step is to login into the gateway using ssh:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

Please refer to the [Kerlink wiki](http://wikikerlink.fr/wirnet-productline)
for login instructions.

### Enable Kerlink CPF

By default, the Kerlink Common Packet-Forwarder (CPF) is disabled. Please
make sure it is enabled. The following command can be used to enable the CPF:

```bash
klk_apps_config --activate-cpf
```

### Configure Kerlink CPF

You must configure the packet-forwarder on the gateway to forward its data to
`127.0.0.1` at port `1700`. The file `/etc/lorafwd/lorafwd.toml` must contain the
following lines under the `[ gwmp ]` section:

```toml
node = "127.0.0.1"
service.uplink = 1700
service.downlink = 1700
```

After updating this configuration file, make sure to restart the `lorafwd` service:

```bash
monit restart lorafwd
```

### Install ChirpStack Gateway Bridge

Find the latest package at [https://artifacts.chirpstack.io/vendor/kerlink/keros-gws/](https://artifacts.chirpstack.io/vendor/kerlink/keros-gws/)
and copy the URL to your clipboard. Then on the gateway use `wget` to download
the package into a folder named `/user/.updates`. Example for `chirpstack-gateway-bridge_{{ gateway_bridge.version }}-r1_klkgw.ipk`:

```bash
cd /user/.updates
wget https://artifacts.chirpstack.io/vendor/kerlink/keros-gws/chirpstack-gateway-bridge_{{ gateway_bridge.version }}-r1_klkgw.ipk
```

To trigger the gateway to install / update the package, run the following commands:

```bash
sync
kerosd -u
reboot
```

Please refer to the [Kerlink wiki](http://wikikerlink.fr/wirnet-productline)
for more information about installing and updating packages.

### Configure ChirpStack Gateway Bridge

To connect the ChirpStack Gateway Bridge with your MQTT broker, you must update
the ChirpStack Gateway Bridge configuration file, which is located at:
`/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`.

### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack Gateway Bridge Service:

```bash
# status
monit status chirpstack-gateway-bridge

# start
monit start chirpstack-gateway-bridge

# stop
monit stop chirpstack-gateway-bridge

# restart
monit restart chirpstack-gateway-bridge
```

## Kerlink IOT station

* [Product detail page](https://www.kerlink.com/product/wirnet-station/)

The Kerlink IOT station has a mechanism to start "custom" application on boot.
These steps will install the LoRa Gateway Bridge ARM build on the Kerlink.

1. Add the Semtech Packet Forwarder (SPF) depending on version of your Wirnet station.

	Semtech Packet Forwarder v3.1.0-klk16 (May 2018):
	This Packet forwarder is only compatible with HAL 4.1.3-klk8, with firmware v3.x and upper and with 27 dBm Wirnet™ stations.
	(dota_spf_3.1.0-klk16_4.1.3-klk8_wirnet.tar.gz)
	Semtech Packet Forwarder v3.1.0-klk11 (April 2017):

	Packet forwarder is only compatible with HAL 4.1.3.
	The packet forwarder source code is only compatible with firmware v2.2 and upper. This source is only compatible with 27 dBm Wirnet™ stations.
	Packet Forwarder - DOTA.
	Wirgrid DOTA is for Wirgrid v2.x Firmware (dota_spf_3.1.0-klk11_4.1.3-klk3_wirgrid_31_03_2017.tar.gz).
	Wirnet DOTA is for Wirnet v3.x Firmware (dota_spf_3.1.0-klk11_4.1.3-klk3_wirnet_31_03_2017.tar.gz).

	Please refer to the [Kerlink wiki](http://wikikerlink.fr/lora-station/) for the complete procedure as well as to recover the dota_spf_xxx files.

2. Untar the dota_spf_xxx.tar.gz (example):

		root@Debian02:~# tar zxvf dota_spf_3.1.0-klk11_4.1.3-klk3_wirgrid_31_03_2017.tar.gz
		./
		./end_dota.sh
		./mnt/
		./mnt/fsuser-1/
		./mnt/fsuser-1/spf/
		./mnt/fsuser-1/spf/etc/
		./mnt/fsuser-1/spf/etc/global_conf_US903.json
		./mnt/fsuser-1/spf/etc/global_conf_EU868.json
		./mnt/fsuser-1/spf/etc/global_conf_JP923.json
		./mnt/fsuser-1/spf/manifest.xml
		./mnt/fsuser-1/spf/bin/
		./mnt/fsuser-1/spf/bin/execute_spf.sh
		./mnt/fsuser-1/spf/bin/spf

3. Edit global_conf_XXXXX.json and add your gateway conf:

		{
			"gateway_conf": {
				"gateway_ID": "0000000000000000",
				"serv_port_up": 1700,
				"serv_port_down": 1700,
				"server_address": "localhost",
				"forward_crc_valid": true,
				"forward_crc_error": false,
				"forward_crc_disabled": true,
				"gps": true
			}
		}

4. Create the the directories needed:

	mkdir -p /mnt/fsuser-1/chirpstack-gateway-bridge/bin

5. Download and extract the ChirpStack Gateway Bridge ARMv5 binary into the above
   directory. See [downloads](../downloads.md).
   Make sure the binary is marked as executable.

6. Save the following content as `/mnt/fsuser-1/chirpstack-gateway-bridge/start.sh`:

		#!/bin/bash

		LOGGER="logger -p local1.notice"

		cd /mnt/fsuser-1/spf/bin/.
		./execute_spf.sh

		iptables -A INPUT -p tcp --sport 1883 -j ACCEPT
		cd /mnt/fsuser-1/chirpstack-gateway-bridge/bin/.
		./chirpstack-gateway-bridge --config /var/config/chirpstack-gateway-bridge.toml

7. Add your chirpstack-gateway-bridge.toml [configuration](../install/config.md) in the /var/config/. directory.

8. Save the following content as `/mnt/fsuser-1/chirpstack-gateway-bridge/manifest.xml`:

		<?xml version="1.0"?>
		<manifest>
			<app name="chirpstack-gateway-bridge" appid="1" binary="start.sh" >
				<start param="" autostart="y"/>
				<stop kill="9"/>
			</app>
		</manifest>

Reboot your system.
