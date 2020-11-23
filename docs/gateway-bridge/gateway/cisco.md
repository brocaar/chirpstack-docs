---
description: Installation of the ChirpStack Gateway Bridge on Cisco gateway.
---

# Cisco

## Cisco Wireless Gateway

* [Product detail page](https://www.cisco.com/c/en/us/products/routers/wireless-gateway-lorawan/index.html)
* [Product documentation](https://www.cisco.com/c/en/us/support/routers/interface-module-lorawan/tsd-products-support-series-home.html)
* [Firmware downloads](https://software.cisco.com/download/home/286311296/type/286311234/release/)

### Preparation

Before proceeding with the following steps, make sure you have connected the
antennas and (PoE) ethernet interface as documented by the Cisco manual.

The following steps are executed using the Cisco Console interface, for which
you need a special USB (connected to your computer) to RJ45 (connected to the
gateway) console cable. 

**Note:** the following instructions only reflect the configuration to get you
started. Please consult the Cisco Wireless Gateway documentation for a complete
manual.


### Connect to console

You can use `screen` to connect to the Cisco (serial) Console. Example:

```bash
# replace /dev/ttyUSB with the serial device
screen /dev/ttyUSB0 115200
```

After the gateway has been fully started, you will see the folliwing line:

```text
Press RETURN to get started
```

Press _RETURN_ and you will see `Gateway>` as prompt. For the configuration of
the Cisco, you need to turn on the _Privileged commands_. To do so, enter the
following command:

```shell
enable
```

The prompt should now have changed to `Gateway#`.

### Firmware version

Enter the following command to display the installed firmware version of the
gateway:

```shell
show version
```

Make sure that the version is (at least) 2.1.0. If your gateway has an older
version installed, please update it first.

#### Upgrade

To upgrade the firmware, put the firmware image on an USB key, then insert
the USB key into the gateway. In the following example, the firmware image
is named `ixm_mdm_i_k9-2.1.0.2.tar.gz`.

Mount the USB key and change directory and list its content:

```shell
usb enable
cd usb:/
dir
```

The output should look simmilar to:

```shell
Directory of usb:/

  -rwx    87371566  Nov 17 2020 20:06:06  ixm_mdm_i_k9-2.1.0.2.tar.gz
```

To only update the firmware and keep the user data:

```shell
archive download-sw firmware /normal /save-reload ixm_mdm_i_k9-2.1.0.2.tar.gz
```

To upgrade the firmware and delete the user data:

```shell
archive download-sw firmware /factory /force-reload ixm_mdm_i_k9-2.1.0.2.tar.gz
```

### Network setup

Enter the following commands to configure the Gateway network interface:

```shell
# Configure gateway from the terminal
configure terminal

# Select interface to configure
interface FastEthernet 0/1
```

To automatically assign an IP address using DHCP:

```shell
ip address dhcp
```

To assign a static IP address to the gateway:

```shell
ip address <ip-address> <subnet-mask>
```

To save the network interface configuration:

```shell
# Set interface specific description
description Ethernet

# Exit interface configuration
exit

# Exit configuration mode
exit

# Save the configuration
copy running-config startup-config
```

To test that the ethernet interface has been configured properly, you can use
the `ping ip` command:

```shell
ping ip <ip-address>
```

### Enable GPS

Enter the following commands to enable the GPS module:

```shell
# Configure gateway from the terminal
configure terminal

# Enable UBX data in UART output
gps ubx enable

# Exit configuration mode
exit

# Save the configuration
copy running-config startup-config
```

### Radio status

Enter the following commands to make sure the radio is enabled:

```shell
show radio
```

When it is turned of, turn it on with the following commands:

```shell
# Configure gateway from the terminal
configure terminal

# Enable the radio
no radio off

# Exit configuration mode
exit

# Save the configuration
copy running-config startup-config
```

### Common Packet Forwarder

The Cisco Wireless Gateway comes with a Common Packet Forwarder which is
compatible with the [Semtech Basic Station](https://doc.sm.tc/station/).
In this case, the ChirpStack Gateway Bridge will not run on the gateway, but
must be installed on a separate server, with the Basic Station backend enabled.

#### ChirpStack Gateway Bridge configuration

Below you will find a simplified configuration example for the EU868 band.
Refer to the [Configuration](../install/config.md) page for a full configuration
example.

```toml
# Gateway backend configuration.
[backend]

# Backend type.
type="basic_station"

  # Basic Station backend.
  [backend.basic_station]

  # ip:port to bind the Websocket listener to.
  bind=":3001"

  # Region.
  region="EU868"

  # Minimal frequency (Hz).
  frequency_min=863000000

  # Maximum frequency (Hz).
  frequency_max=870000000

  # Concentrator configuration.
  # Note: this is defined twice as the Cisco gateway has two SX1301 chips.
  [[backend.basic_station.concentrators]]
    # Multi-SF channel configuration.
    [backend.basic_station.concentrators.multi_sf]
 
    # Frequencies (Hz).
    frequencies=[
      868100000,
      868300000,
      868500000,
      867100000,
      867300000,
      867500000,
      867700000,
      867900000,
    ]
  
    # LoRa STD channel.
    [backend.basic_station.concentrators.lora_std]
  
    # Frequency (Hz).
    frequency=868300000
  
    # Bandwidth (Hz).
    bandwidth=250000
  
    # Spreading factor.
    spreading_factor=7
  
    # FSK channel.
    [backend.basic_station.concentrators.fsk]
  
    # Frequency (Hz).
    frequency=868800000

  [[backend.basic_station.concentrators]]
    # Multi-SF channel configuration.
    [backend.basic_station.concentrators.multi_sf]
 
    # Frequencies (Hz).
    frequencies=[
      868100000,
      868300000,
      868500000,
      867100000,
      867300000,
      867500000,
      867700000,
      867900000,
    ]
  
    # LoRa STD channel.
    [backend.basic_station.concentrators.lora_std]
  
    # Frequency (Hz).
    frequency=868300000
  
    # Bandwidth (Hz).
    bandwidth=250000
  
    # Spreading factor.
    spreading_factor=7
  
    # FSK channel.
    [backend.basic_station.concentrators.fsk]
  
    # Frequency (Hz).
    frequency=868800000


# Integration configuration.
[integration]

# Payload marshaler.
marshaler="protobuf"

  # MQTT authentication.
  [integration.mqtt.auth]
  type="generic"

    # Generic MQTT authentication.
    [integration.mqtt.auth.generic]
    # MQTT servers.
    #
    # Configure one or multiple MQTT server to connect to. Each item must be in
    # the following format: scheme://host:port where scheme is tcp, ssl or ws.
    servers=[
      "tcp://127.0.0.1:1883",
    ]
```

#### Common Packet Forwarder configuration

!!! info
	Common Packet Forwarder will use the GPS module to determine the country
	and region. In case `cpf enable` will result in an error that the region is
	undefined, make sure that the gateway is able to obtain a GPS position and
	try again.

Enter the following commands to configure the Common Packet Forwarder:

```shell
# Configure gateway from the terminal
configure terminal

# Enter the Common Packet Forwarder configuration
common-packet-forwarder profile

# Configure antenna gain and loss value
antenna 1 omni gain 4.3 loss 0.1

# Configure gateway ID
# Replace <GatewayID> with your gateway ID
gatewayid <GATEWAYID>

# Enable GPS usage
gps enable

# Configure IP and port to which the Common Packet Forwarder must connect
# Replace <IP> with the IP of the ChirpStack Gateway Bridge
# Replace <PORT> with the port on which ChirpStack Gateway Bridge is listening
ipaddr <IP> port <PORT>

# Enable Common Packet Forwarder
cpf enable

# Exit the (Common Packet Forwarder) configuration mode
exit
exit

# Save the configuration
copy running-config startup-config
```
