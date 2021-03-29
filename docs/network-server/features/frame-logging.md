---
description: LoRaWAN frame logging.
---

# Frame logging

ChirpStack Network Server provides an option to log all uplink and downlink
frames to a Redis stream, which can be consumed by external application(s)
for monitoring or logging purposes.

There are two streams that can be configured:

1. Frames received and sent by all gateways
2. Frames received from and sent to devices

The difference between these two is that in case of 1., it will include
**all** frames, thus also frames from devices from neighboring networks and
frames which did not pass the frame-counter and MIC checks (and thus could
have been replay-attacks). As in case of 2. the device is known, ChirpStack
Network Server will always decrypt the mac-commands before logging the frame
to the Redis stream in case this is needed.

## Configuration

Please refer to the `monitoring.device_frame_log_max_history` and
`monitoring.gateway_frame_log_max_history` [Configuration](../install/config.md)
options.

## Redis keys

The following Redis keys are used:

* `lora:ns:gw:stream:frame`: for all gateway uplink and downlink frames
* `lora:ns:device:stream:frame`: for all device uplink and downlink frames

## Message content

### Uplink

Uplink frames are published under the `up` field of the Redis Stream. The content
of this frame is a [Protobuf](https://developers.google.com/protocol-buffers/)
encoded payload. The description of the `UplinkFrameLog` message can be found
in the [ns.proto](https://github.com/brocaar/chirpstack-api/blob/master/protobuf/ns/ns.proto)
file.

### Downlink

Downlink frames are published under the `down` field of the Redis Stream. The
content of this frame is a [Protobuf](https://developers.google.com/protocol-buffers/)
encoded payload. The description of the `DownlinkFrameLog` message can be found
in the [ns.proto](https://github.com/brocaar/chirpstack-api/blob/master/protobuf/ns/ns.proto)
file.

## Redis Streams

Redis Streams is a feature of [Redis](https://redis.io/) which provides a
log data structure. It makes it possible to write and persist a configurable
amount of uplink / downlink frames, which can be consumed by one or multiple
consumers part of a consumer group. Refer to [https://redis.io/topics/streams-intro](https://redis.io/topics/streams-intro)
for an introduction and examples how this feature can be used.
