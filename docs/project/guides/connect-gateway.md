---
description: Connecting LoRa gateways to ChirpStack and steps for troubleshooting issues.
---

# Connecting a gateway

This guide describes how to connect your gateway to ChirpStack and how
to validate that it is succesfully communicating with the [ChirpStack Network Server](../../network-server/index.md).
At this point it is expected that you have the [ChirpStack Application Server](../../application-server/index.md)
and [ChirpStack Network Server](../../network-server/index.md) components up
and running.

!!! warning
	This guide does not cover the configuration of MQTT credentials and / or
	(client) certificates.

## Requirements

Before continuing, please make sure that you have installed both a
packet-forwarder and the [ChirpStack Gateway Bridge](../../gateway-bridge/index.md).
The packet-forwarder that is installed on your gateway and
the steps needed to install the ChirpStack Gateway Bridge vary per gateway vendor
and model. In some cases you must also install the ChirpStack Gateway Bridge on the
gateway. Please refer to the ChirpStack Gateway Bridge [Gateway installation](../../gateway-bridge/gateway/index.md)
documentation, which contains instructions for various gateway models.

### Packet-forwarders

There are different packet-forwarder implementations.
The packet-forwarder that is installed on your gateway depends on the gateway vendor and model.
The packet-forwarders that are compatible with ChirpStack:

* [ChirpStack Concentratord](../../concentratord/index.md)
* [Semtech UDP Packet Forwarder](https://github.com/lora-net/packet_forwarder)
* [Semtech BasicStation](https://doc.sm.tc/station/)

### ChirpStack Gateway Bridge

The [ChirpStack Gateway Bridge](../../gateway-bridge/index.md) component
acts as a backend for the above packet-forwarders. It communicates with the
[ChirpStack Network Server](../../network-server/index.md) using [MQTT](https://mqtt.org/).

## Configuration

There are two components that you need to configure. This section covers a
summary. Please refer to the ChirpStack Gateway Bridge [Gateway installation](../../gateway-bridge/gateway/index.md)
for instructions specific to your gateway model.

### Packet-forwarder

The packet-forwarder that is configured on your gateway must forward its data
to the ChirpStack Gateway Bridge. As it controls the LoRa<sup>&reg;</sup> chipset of the
gateway, it also must be configured for the correct frequencies. A mismatch
in frequencies means that the gateway will not receive uplinks sent by a device
and / or is unable to send downlink payloads when the downlink frequency is
outside the configured frequency range. Usually gateway vendors provide
configuration examples for various bands. Please validate that the configuration
matches the band and channels in the ChirpStack Network Server [Configuration](../../network-server/install/config.md).

### ChirpStack Gateway Bridge

The [ChirpStack Gateway Bridge](../../gateway-bridge/index.md) must be configured
with the correct packet-forwarder backend. Please refer to the ChirpStack Gateway Bridge [Gateway installation](../../gateway-bridge/gateway/index.md)
for instructions specific to your gateway model. When you have installed a
vendor specific package, this has already been pre-configured for you.

What you still need to configure is to which MQTT broker the ChirpStack Gateway
Bridge will connect. This is configured in the following [Configuration](../../gateway-bridge/install/config.md)
section of the ChirpStack Gateway Bridge:

```toml
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

Please refer to the ChirpStack Gateway Bridge [Configuration](../../gateway-bridge/install/config.md)
documentation for a full configuration example. After making changes to the
configuration file, do not forget to restart the ChirpStack Gateway Bridge.

## Adding the gateway to ChirpStack

[Login](../../application-server/use/login.md) into the [ChirpStack Application Server](../../application-server/index.md) web-interface.
The default credentials are:

* Username: admin
* Password: admin

### Add gateway

!!! info
	If you have not yet connected your [ChirpStack Application Server](../../application-server/)
	instance with a [ChirpStack Network Server](../../network-server/) instance,
	you must do this first. See [Network servers](../../application-server/use/network-servers.md).
	Also you must connect the organization with the network-server by creating
	a [Service profile](../../application-server/use/service-profiles.md).

Navigate to **Gateways** in the web-interface, and click ** + Create** and
complete the form. Make sure that the **Gateway ID** field is equal to the
Gateway ID of your gateway. If this value is incorrectly configured, data
received by your gateway will be rejected.

## Validate

There are a few ways to validate if your gateway is correctly configured.

### Last seen at

Event when no LoRa(WAN) data is received by the gateway, it will send gateway
statistics periodically. Usually this stats interval is configured to 30 seconds.
As [ChirpStack Application Server](../../application-server/index.md) will update
the gateway **Last seen at** timestamp when it receives statistics, this is the
easiest way to validate that the gateway is correctly configured. **Note:** it
might take a short while before statistics are sent by your gateway. You must
refresh the page in order to see the (new) **Last seen at** value.

### LoRaWAN frames

After opening the overview page of your gateway, you will see a **LoRaWAN frames**
tab. This will show all LoRaWAN frames that are received and sent by your
gateway. In case of received frames, this means that you will also see received
frames from devices that are not yours and / or that are not yet configured.
Therefore this screen is useful to validate if your gateway is able to receive
LoRaWAN frames and forward these to ChirpStack.

### MQTT

As the [ChirpStack Gateway Bridge](../../gateway-bridge/index.md) uses MQTT
for communication with the [ChirpStack Network Server](../../network-server/index.md)
it is also possible to subscribe to the gateway MQTT topics. This will validate
that data from your gateway arrives at least up to your MQTT broker.

An example command to subscribe to the gateway MQTT topic using `mosquitto_sub`:

```bash
mosquitto_sub -v -t "gateway/#"
```

If you see data like below, then this does not indicate that there is an issue.
The ChirpStack Gateway Bridge can publish (and receive) events either as JSON
or [Protobuf](https://developers.google.com/protocol-buffers), in which case
the payloads are binary encoded. This option can be configured in the
ChirpStack Gateway Bridge [Configuration](../../gateway-bridge/install/config.md).

```text
gateway/00800000a00016b6/event/up 
��b��鄞▒       }
                ▒4/5▒X
���
  ѣ�����_▒
�$~ �    䨡����^(���������1"@8@zv��d���\�A3Al��
```

## Troubleshooting

In case no data is received by [ChirpStack Application Server](../../application-server/index.md),
follow the steps below to troubleshoot.

### Packet forwarder

The first component to troubleshoot is the packet-forwarder on the gateway.
By sending an uplink message from a LoRaWAN device (it does not have to be
configured in ChirpStack) and checking the packet-forwarder logs, you can
verify if the gateway is able to receive the uplinks. Not receiving any uplink
messages could indicate a mis-configuration of the gateway.

Depending the gateway model, the commands to retrieve the packet-forwarder logs
can vary. Please refer to the ChirpStack Gateway Bridge [Gateway installation](../../gateway-bridge/gateway/index.md)
for model specific instructions.

#### ChirpStack Concentratord

##### Received uplink

On receiving an uplink message, it is expected to see a log item similar to:

```text
Frame received, uplink_id: 648c85f9-19e8-4186-af96-f5a45eae7ba5, count_us: 105814115, freq: 868500000, bw: 125000, mod: LoRa, dr: SF8
```

#### Semtech UDP Packet Forwarder

##### PULL_ACK

it is expected to periodically see log items similar to the following:

```text
INFO: [down] PULL_ACK received in 1 ms
```

This indicates that the UDP data sent by the packet-forwarder
to the ChirpStack Gateway Bridge is acknowledged.

##### Received uplink

On receiving an uplink message, it is expected to see a log item similar to:

```text
JSON up: {"rxpk":[{"tmst":661201908,"time":"2020-10-06T13:25:02.523074Z","tmms":1286025921523,"chan":1,"rfch":1,"freq":868.300000,"stat":1,"modu":"LORA","datr":"SF11BW125","codr":"4/5","lsnr":8.0,"rssi":-42,"size":23,"data":"AAEAAAAAAAAAAQEBAQEBAQGheWrMbQo="}]}
```

##### tcpdump

An other way to troubleshoot the functioning of the Semtech UDP Packet Forwarder
is to use `tcpdump`.

If the ChirpStack Gateway Bridge is installed on the gateway itself, then the
following command must be executed on the gateway itself:

```bash
sudo tcpdump -AUq -i lo port 1700
```

If the ChirpStack Gateway Bridge is installed on a server, then you must execute
the following command (either on the gateway or on the receiving server):

```bash
sudo tcpdump -AUq port 1700
```

As the above command can be either executed on the gateway or on the server,
this allows you to validate not only if the gateway is forwarding data, but
also if the server is able to receive the data. If the gateway is correctly
forwarding but the server is not receiving, then this could indicate a network
related issue (e.g. firewall).


Output similar to the following is expected:

```text
11:42:00.114726 IP localhost.34268 > localhost.1700: UDP, length 12
E..(..@.@."................'.....UZ.....
11:42:00.130292 IP localhost.1700 > localhost.34268: UDP, length 4
E.. ..@.@.".....................
11:42:10.204723 IP localhost.34268 > localhost.1700: UDP, length 12
E..(.&@.@..................'.x...UZ.....
11:42:10.206503 IP localhost.1700 > localhost.34268: UDP, length 4
E.. .'@.@....................x..
11:42:10.968420 IP localhost.43827 > localhost.1700: UDP, length 113
E....h@.@............3...y.......UZ.....{"stat":{"time":"2017-09-11 11:42:10 GMT","rxnb":0,"rxok":0,"rxfw":0,"ackr":100.0,"dwnb":0,"txnb":0}}
11:42:10.970702 IP localhost.1700 > localhost.43827: UDP, length 4
E.. .i@.@..b...........3........
11:42:20.284752 IP localhost.34268 > localhost.1700: UDP, length 12
E..(..@.@..................'.....UZ.....
11:42:20.289256 IP localhost.1700 > localhost.34268: UDP, length 4
E.. ..@.@.......................
11:42:30.364780 IP localhost.34268 > localhost.1700: UDP, length 12
E..( .@.@..................'.S7..UZ.....
11:42:30.366310 IP localhost.1700 > localhost.34268: UDP, length 4
E..  .@.@....................S7.
```

Explanation:

* `localhost.34268 > localhost.1700`: packet sent from the packet-forwarder to the ChirpStack Gateway Bridge
* `localhost.1700 > localhost.34268`: packet sent from the ChirpStack Gateway Bridge to the Packet Forwarder

#### Semtech BasicStation

##### Received uplink

On receiving an uplink message, it is expected to see a log item similar to:

```text
RX 867.9MHz DR1 SF11/BW125 snr=10.8 rssi=-55 xtime=0x91000001A21C4B - updf mhdr=40 DevAddr=00D35EC3 FCtrl=00 FCnt=0 FOpts=[] 01A5 mic=643170495 (14 bytes)
```

### ChirpStack Gateway Bridge

When you have verified that the packet-forwarder is receiving LoRa(WAN) frames
sent by your device(s), you can validate that it is communicating
succesfully with the [ChirpStack Gateway Bridge](../../gateway-bridge/index.md).
You can do this by inspecting the ChirpStack Gateway Bridge logs.

#### Gateway detected

When ChirpStack Gateway Bridge detects data from a gateway, it will subscribe
to a MQTT topic containing the gateway ID. You should see a log message
similar to:

```text
integration/mqtt: subscribing to topic        qos=0 topic="gateway/00800000a00016b6/command/#"
```

Note that this log message is only printed once. If you do not see this log
message, you need to confirm that:

1. The packet-forwarder is correctly configured to forward its data to your
   ChirpStack Gateway Bridge instance.
2. In case the packet-forwarder and the ChirpStack Gateway Bridge are not
   installed on the same device, that there are no firewalls blocking the
   communication between the gateway and the server.
3. The ChirpStack Gateway Bridge is configured with the correct backend
   configuration for the used packet-forwarder. Please refer to the
   ChirpStack Gateway Bridge [Configuration](../../gateway-bridge/install/config.md)
   documentation.

#### Uplink received

When ChirpStack Gateway Bridge receives an uplink message from the gateway,
it will print a message similar to:

```text
publishing event            event=up qos=0 topic=gateway/00800000a00016b6/event/up uplink_id=c2f855a9-7f74-4093-98ad-7d2b06a79398
```

### MQTT broker

If you have verified the above components, you should be able to see messages
being published to your MQTT broker. For this, please repeat the MQTT validate
step.

In case the ChirpStack Gateway Bridge does log that it is publishing events,
but you are unable to receive data from the MQTT broker, then verify if you
configured ACLs that prevent the ChirpStack Gateway Bridge from publishing
events and / or prevent your MQTT client from receiving events.

### ChirpStack Network Server

If you have verified that (uplink) events are published to your MQTT broker,
you can verify that these are received by the [ChirpStack Network Server](../../network-server/index.md).
When the ChirpStack Network Server receives an uplink frame from one of your
gateways through the MQTT broker, it will print a log item similar to:

```
gateway/mqtt: uplink frame received           gateway_id=00800000a00016b6 uplink_id=be864b60-1392-47a1-98ef-8cd1115b6f05
```

If you do not see similar log items, then validate:

1. That in case you have configured MQTT ACLs, the ChirpStack Network Server
   has access to the configured MQTT topics.
2. That in case you have modified the MQTT topic templates in the
   ChirpStack Gateway Bridge [Configuration](../../gateway-bridge/install/config.md)
   and / or ChirpStack Network Server [Configuration](../../network-server/install/config.md),
   this configuration is aligned. With the default topic configuration, the
   ChirpStack Gateway Bridge publishes events on topics expected by the
   ChirpStack Network Server.
