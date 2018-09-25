---
title: Network-server troubleshooting
menu:
    main:
        parent: troubleshooting
description: Troubleshooting network-server (LoRa Server) related issues.
---

# Troubleshooting network-server issues

This guide helps you to troubleshoot network-server related connectivity
issues. This guide assumes that your gateway is connected and that the LoRa
Gateway Bridge is publishing the received data. If you are not sure, please
refer to the [Debugging gateway connectivity]({{<relref "gateway.md">}}) guide.
It also assumes that you have setup the [LoRa Server](/loraserver/) and
[LoRa App Server](/lora-app-server/) component and that you have
[setup your gateway and device]({{<ref "/guides/first-gateway-device.md">}})
through the web-interface.

In this guide we will validate:

* If the [LoRa Server](/loraserver/) component receives uplink data
* If the received data is valid (MIC and frame-counters)
* If the received payload is forwarded to [LoRa App Server]({{<relref "application-server.md">}})

## LoRa Server receives data?

To find out if [LoRa Server](/loraserver/) is receiving messages from your
gateway, you should refer to the logs. Depending how LoRa Server was installed,
one of the following commands will show you the logs:

* `journalctl -f -n 100 -u loraserver`
* `tail -f -n 100 /var/log/loraserver/loraserver.log`

### Expected log output

{{<highlight text>}}
INFO[0163] backend/gateway: uplink frame received
INFO[0164] device gateway rx-info meta-data saved        dev_eui=0101010101010101
INFO[0164] device-session saved                          dev_addr=018f5aa9 dev_eui=0101010101010101
INFO[0164] finished client unary call                    grpc.code=OK grpc.method=HandleUplinkData grpc.service=as.ApplicationServerService grpc.time_ms=52.204 span.kind=client system=grpc
{{< /highlight >}}

In this log, LoRa Server has processed the uplink frame, and forwarded the
application payload to the `ApplicationServerService` API
(provided by [LoRa App Server](/lora-app-server/)).

### Invalid MIC or frame-counter

{{<highlight text>}}
INFO[0581] backend/gateway: uplink frame received
ERRO[0581] processing uplink frame error                 data_base64="QKlajwGCAgADBwF6Eabhjw==" error="get device-session error: device-session does not exist or invalid fcnt or mic"
{{< /highlight >}}

In this log, LoRa Server was unable to authenticate the received data and
map this to a device-session. The cause could be:

* The device-session (activation) does not exist the device
* The session-keys do not match the session-keys as known by the network-server

#### OTAA device

Try to re-activate your device. The activation did not match any of the
device-sessions stored in the database. After a long time of inactivity, it
could have expired. Also make sure that the LoRaWAN version selected in the
Device-profile matches the implemented LoRaWAN version by the device!

#### ABP device

In case of an ABP device, there are a couple of things that could case this
error:

First of all, make sure that the session-keys are entered in the correct
byte-order. Some devices expect little-endian, some big-endian and in some
cases these are mixed. One suggestion to find this out is by starting with
a key where `LSB(value)` equals `MSB(value)`. Example: `01010101...`, where
the order does not matter.

When the error is raised after the device was reset (e.g. power-cycle) then
the issue is likely related to a frame-counter reset. A LoRaWAN network-server
must check that the frame-counter is incremented, if not it will be rejected.
This is to protect the network against replay-attacks. The problem is that most
ABP devices reset their frame-counters to `0` on a power-cycle, which is not
in line with the LoRaWAN specification.
