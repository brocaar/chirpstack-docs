---
description: Configure the Tektelic gateways to work with the ChirpStack Gateway Bridge.
---

# Tektelic

## KONA Pico IoT Gateway

* [Product detail page](https://tektelic.com/iot/lorawan-gateways/)

The KONA Pico IoT Gateway is able to run different firmwares, each using a
different protocol. In order to work together with the ChirpStack Gateway Brige,
you must install the `semtech-vx.xx.bin` firmware. These steps have been tested
with version 1.2.1. The latest firmware can be [downloaded here](https://artifacts.chirpstack.io/vendor/tektelic/kona-pico/).

1. After turning on the gateway, find the IP that has been assigned to it
   (eg. by listing the devices connected to your router). In the examples below
   it is assumed that the IP of the gateway is `192.168.1.10`.

2. Upload the channel configuration to the gateway. For this you need to use
   TFTP. The example command below shows how to do this from the command-line.
   The KONA Pico IoT Gateway configuration user-guide contains instructions how
   to do this using using a UI (Windows).

		$ tftp 192.168.1.10
		tftp> put lorawan_conf.json
		Sent 1412 bytes in 0.0 seconds

   Note: the tftp command is invoked from the same directory as where the
   `lorawan_conf.json` is stored.

3. Upload the configuration file containing the IP of the ChirpStack Gateway Bridge
   instance and the used ports.

		$ tftp 192.168.1.10
		tftp> put customer_conf.json
		Sent 121 bytes in 0.1 seconds

   Note: the tftp command is invoked from the same directory as where the
   `customer_conf.json` is stored.

### Example configuration files

#### lorawan_conf.json for EU868

This configuration file contains exactly the same radio configuration as
[global_conf.json.PCB_E336.EU868.basic](https://github.com/Lora-net/packet_forwarder/blob/master/lora_pkt_fwd/cfg/global_conf.json.PCB_E336.EU868.basic)

```json
{
    "public": true,
    "radio": [
        {
            "enable": true,
            "freq": 867500000
        },
        {
            "enable": true,
            "freq": 868500000
        }
    ],
    "lora_multi": [
        {
            "enable": true,
            "radio": 1,
            "offset": -400000
        },
        {
            "enable": true,
            "radio": 1,
            "offset": -200000
        },
        {
            "enable": true,
            "radio": 1,
            "offset": 0
        },
        {
            "enable": true,
            "radio": 0,
            "offset": -400000
        },
        {
            "enable": true,
            "radio": 0,
            "offset": -200000
        },
        {
            "enable": true,
            "radio": 0,
            "offset": 0
        },
        {
            "enable": true,
            "radio": 0,
            "offset": 200000
        },
        {
            "enable": true,
            "radio": 0,
            "offset": 400000
        }
    ],
    "lora_std": {
        "enable": true,
        "radio": 1,
        "offset": -200000,
        "bandwidth": "250kHz",
        "spread_factor": "SF7"
    },
    "fsk": {
        "enable": true,
        "radio": 1,
        "offset": 300000,
        "bandwidth": "125kHz",
        "datarate": 50000
    }
}
```

#### customer_conf.json

The IP or hostname of the `network_server` key must match the IP or hostname
of the machine on which ChirpStack Gateway Bridge is running.

```json
{
    "network_server": "192.168.1.5",
    "network_service_up_port": 1700,
    "network_service_down_port": 1700
}
```

## Kona gateways

These instructions appy to the Linux based Tektelic Kona gateways.

* [Product detail page](https://tektelic.com/catalog//type[2])

### SSH into the gateway

The first step is to login into the gateway using ssh:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

The default password is the serial-number of the gateway which is printed on
the back of the gateway (the 9 characters above the 12V = 1A line).

### Download IPK package

Find the latest package at https://artifacts.chirpstack.io/vendor/tektelic/kona/
and copy the URL to your clipboard. Then on the gateway use `curl` and use the link
as argument. Example for `chirpstack-gateway-bridge_{{ gateway_bridge.version }}-r1_kona.ipk`:

```bash
# curl URL --output chirpstack-gateway-bridge.ipk
curl https://artifacts.chirpstack.io/vendor/tektelic/kona/chirpstack-gateway-bridge_{{ gateway_bridge.version }}-r1_kona.ipk --output chirpstack-gateway-bridge.ipk
```

### Install IPK package

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
opkg install chirpstack-gateway-bridge.ipk
```

### Edit the ChirpStack Gateway Bridge configuration

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

### Configure packet-forwarder

You must configure the packet-forwarder on the gateway to forward its data to
`127.0.0.1` at port `1700`. The file `/etc/default/config.json` must contain the
following lines:

```json
"server_address": "127.0.0.1",
"serv_port_up": 1700,
"serv_port_down": 1700,
```
