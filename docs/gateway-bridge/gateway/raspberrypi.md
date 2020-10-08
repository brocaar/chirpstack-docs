---
description: Installation of the ChirpStack Gateway Bridge on Raspberry Pi based gateways.
---

# Raspberry Pi

## ChirpStack Gateway OS

Rasberry Pi based gateways are supported by the
[ChirpStack Gateway OS](../../gateway-os/index.md) images. This is the preferred
and easiest option to get started. Please refer the ChirpStack Gateway OS
documentation for installation instructions.

## Raspberry Pi OS

### ChirpStack Gateway Bridge

ChirpStack provides a package repository that is compatible with Raspberry Pi OS.
In order to add this repository, execute the following command:

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

sudo echo "deb https://artifacts.chirpstack.io/packages/3.x/deb stable main" | sudo tee /etc/apt/sources.list.d/chirpstack.list
sudo apt update
```

Then in order to install the ChirpStack Gateway Bridge:

```bash
sudo apt install chirpstack-gateway-bridge
```

If the above step has been completed, update the ChirpStack Gateway Bridge
configuration file. This file is located at `/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`.
By default it is configured to use with the [Semtech UDP Packet Forwarder](https://github.com/lora-net/packet_forwarder).
The configuration that you most likely want to update is related to the MQTT integration
configuration.

To (re)start, stop and retrieve the status of the ChirpStack Gateway Bridge
process, use the following commands:

```bash
sudo systemctl [start|stop|restart|status] chirpstack-gateway-bridge
```

To retrieve the ChirpStack Gateway Bridge logs:

```bash
sudo journalctl -u chirpstack-gateway-bridge -f -n 50
```

### Packet Forwarder

#### Semtech UDP Packet Forwarder

The Semtech UDP Packet Forwarder must be configured so that it forwards its
data to the ChirpStack Gateway Bridge that has been installed in the previous
step. Where the configuration file of the Semtech UDP Packet Forwarder is
located depends on how this Packet Forwarder was installed (which is outside
the scope of this documentation). Typically this file is called `local_conf.json`
and / or `global_conf.json` (in case both are present, values set in `local_conf.json`
have priority over variables set in `global_conf.json`).

Find the configuration section called `gateway_conf` and update the
`server_address`, `serv_port_up` and `serv_port_down` accordingly:


```json
"gateway_conf": {
	...
	"server_address": "localhost",
	"serv_port_up": 1700,
	"serv_port_down": 1700,
	...
}
```
