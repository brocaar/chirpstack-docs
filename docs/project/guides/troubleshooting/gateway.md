---
description: Troubleshooting LoRa Gateway and ChirpStack Gateway Bridge issues.
---

# LoRa Gateway troubleshooting

This guide helps you to troubleshoot gateway connectivity issues and to find the
underlying cause. This guide assumes you already have the
[ChirpStack Gateway Bridge](../../../gateway-bridge/index.md) component installed and running.

We will validate in this guide:

* If the [Packet Forwarder](https://github.com/lora-net/packet_forwarder) is receiving device data
* If the [ChirpStack Gateway Bridge](../../../gateway-bridge/index.md) is receiving data from the packet-forwarder
* If ChirpStack Gateway Bridge is publishing the data to the MQTT broker

## Semtech Packet Forwarder sends data?

The first step starts at the source of the data, the gateway. To make sure your
gateway is actually receiving device data, you can use `tcpdump` to monitor
the data that is sent by your gateway.

When the ChirpStack Gateway Bridge is running on the gateway itself, then you need
to run the following command **on the gateway** (it will monitor the loopback interface):

```bash
sudo tcpdump -AUq -i lo port 1700
```

When the ChirpStack Gateway Bridge is installed on a separate machine / server, the
you need to run the following command:

```bash
sudo tcpdump -AUq port 1700
```

The above command can be executed on the gateway (if possible) or on the
machine where the ChirpStack Gateway Bridge component is installed. Running it
on the gateway will show the data sent by the gateway, running it on your
machine will show the data received by your machine.

When you see that data is sent by the gateway, but not received by your
machine / server, then likely there is a firewall inbetween that is blocking
the data.

### Expected tcpdump output

As the packet-forwarder sends regular "ping" messages, there is no need for
your device to send data. You should see a regular data-exchange in the
`tcpdump` output, for example:

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

What we see in this log:

* `localhost.34268 > localhost.1700`: packet sent from the packet-forwarder to the ChirpStack Gateway Bridge
* `localhost.1700 > localhost.34268`: packet sent from the ChirpStack Gateway Bridge to the Packet Forwarder


### No tcpdump output?

When `tcpdump` does not show anything, then likely the Packet Forwarder on
the gateway is not running, or the Packet Forwarder is misconfigured.

Inspect the `local_conf.json` of the Packet Forwarder running on your gateway.
You might need to refer to your gateway manual to find out where you can locate
this file. This file could contain the following content:

```json
{
    "gateway_conf": {
        ...
        "serv_port_down": 1700,
        "serv_port_up": 1700,
        "server_address": "localhost",
        ...
    }
}
```

What we learn from this file is that:

* It uses port `1700` (default port used by ChirpStack Gateway Bridge)
* It sends data to `localhost` (when ChirpStack Gateway Bridge is installed on the same device)

Make sure the ports and `server_address` are correct. In case ChirpStack Gateway
Bridge is not running on the same device, you need to replace this with
the correct hostname or IP of your ChirpStack Gateway Bridge instance. After making
any changes, don’t forget to restart the Packet Forwarder.

See [https://github.com/Lora-net/packet_forwarder/tree/master/lora_pkt_fwd](https://github.com/Lora-net/packet_forwarder/tree/master/lora_pkt_fwd)
for more information about the packet-forwarder.

## ChirpStack Gateway Bridge receives data?

When the previous step confirmed that the gateway is sending data, you need to
confirm that the [ChirpStack Gateway Bridge](../../../gateway-bridge/index.md) is receiving
data from your gateway.

The first indications you will find in the ChirpStack Gateway Bridge logs. Depending
how the ChirpStack Gateway Bridge was installed on your system, one of the following
commands will show you the logs:

* `journalctl -f -n 100 -u chirpstack-gateway-bridge`
* `tail -f -n 100 /var/log/chirpstack-gateway-bridge/chirpstack-gateway-bridge.log`

When the [Packet Forwarder](https://github.com/lora-net/packet_forwarder) sends
data to the ChirpStack Gateway Bridge (this could be a "ping"), you will see the following logs:

```text
INFO[0013] mqtt: subscribing to topic qos=0 topic=gateway/7276ff002e062c18/command/#
```

When your device sends an uplink message, you will see something like:

```text
INFO[0267] mqtt: publishing message qos=0 topic=gateway/7276ff002e062c18/event/up
```

If you see these logs, then this indicates that the ChirpStack Gateway Bridge
components receives the data sent by the packet-forwarder.

### No log output?

#### Double-check gateway

When you don't see any logs printed when your device sends an uplink message
double-check if the packet-forwarder sends data (previous section).

#### ChirpStack Gateway Bridge active?

You also need to make sure that the ChirpStack Gateway Bridge is actually active.
You can use the following command to check this:

```
ps aux |grep chirpstack-gateway-bridge
```

The output should look like:

```text
root      6403  0.0  0.2  12944  1088 pts/0    S+   12:53   0:00 grep --color=auto chirpstack-gateway-bridge
gateway+ 23060  0.1  2.1 214260 10664 ?        Ssl  Aug30  47:55 /usr/bin/chirpstack-gateway-bridge
```

If no ChirpStack Gateway Bridge process is active, please refer to the
[ChirpStack Gateway Bridge](../../../gateway-bridge/index.md) install
documentation.

#### ChirpStack Gateway Bridge configuration

When you have completed the previous steps successfully, then packet-forwarder
data is received by your machine / server, but is not seen by the ChirpStack Gateway Bridge service.
This probably means that ChirpStack Gateway Bridge is binding on a different network
interface and / or port. Please check your `chirpstack-gateway-bridge.toml` [Configuration](../../../gateway-bridge/install/config.md).

## ChirpStack Gateway Bridge publishes data?

If you have confirmed that the ChirpStack Gateway Bridge component receives the
data sent by the packet-forwarder, it is time to confirm ChirpStack Gateway Brige is
successfully publishing this data to the MQTT broker.

To validate that the ChirpStack Gateway Bridge is publishing LoRa<sup>&reg;</sup> frames to the MQTT
broker, you can subscribe to the `gateway/#` MQTT topic. When using the
`mosquitto_sub` utility, you can execute the following command:

```bash
mosquitto_sub -v -t "gateway/#"
```

When you do not see any data appear when your device sends data, then make sure
the ChirpStack Gateway Bridge is authorized to publish to the MQTT topic **and**
the `mosquitto_sub` client is authorized to subscribe to the given MQTT topic.
This issue usually happens when you have configured your MQTT broker so that
clients need to authenticate when connecting.
