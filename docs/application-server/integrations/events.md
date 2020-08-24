---
description: Events that are sent by the ChirpStack Application Server to one or multiple integrations.
---

# Event types

## Encoding

Depending the integration, it is possible to encode events in several ways:

* JSON: JSON based on the Protocol Buffers JSON mapping
* Protobuf: Protocol Buffers binary encoding
* JSON (v3): Legacy JSON format, this option will be removed in the next major release and exists for backwards compatibility

The way in which a payload is encoded can be either configured per integration
in the web-interface, or through the `marshaler` configuration under
`[application_server.integations]` in the [configuration](../install/config.md) file.

For the [Protobuf](https://developers.google.com/protocol-buffers/)
message definitions, please refer to [proto/as/integration/integration.proto](https://github.com/brocaar/chirpstack-api/blob/master/protobuf/as/integration/integration.proto).
in the [chirpstack-api](https://github.com/brocaar/chirpstack-api/) repository.


!!! Important
	The Protocol Buffers [JSON Mapping](https://developers.google.com/protocol-buffers/docs/proto3#json)
	defines that bytes must be encoded as base64 strings. This applies to the
	devEUI field for example. When re-encoding this filed to HEX encoding, you
	will find the expected devEUI string.

## up

Contains the data and meta-data for an uplink application payload.

### JSON

```json
{
    "applicationID": "123",
    "applicationName": "temperature-sensor",
    "deviceName": "garden-sensor",
    "devEUI": "AgICAgICAgI=",
    "rxInfo": [
        {
            "gatewayID": "AwMDAwMDAwM=",
            "time": "2019-11-08T13:59:25.048445Z",
            "timeSinceGPSEpoch": null,
            "rssi": -48,
            "loRaSNR": 9,
            "channel": 5,
            "rfChain": 0,
            "board": 0,
            "antenna": 0,
            "location": {
                "latitude": 52.3740364,
                "longitude": 4.9144401,
                "altitude": 10.5
            },
            "fineTimestampType": "NONE",
            "context": "9u/uvA==",
            "uplinkID": "jhMh8Gq6RAOChSKbi83RHQ=="
        }
    ],
    "txInfo": {
        "frequency": 868100000,
        "modulation": "LORA",
        "loRaModulationInfo": {
            "bandwidth": 125,
            "spreadingFactor": 11,
            "codeRate": "4/5",
            "polarizationInversion": false
        }
    },
    "adr": true,
    "dr": 1,
    "fCnt": 10,
    "fPort": 5,
    "data": "...",
    "objectJSON": "{\"temperatureSensor\":25,\"humiditySensor\":32}",
    "tags": {
        "key": "value"
    }
}
```

### Protobuf

This message is defined by the `UplinkEvent` Protobuf message.

## status

Event for battery and margin status received from devices.

The interval in which the Network Server will request the device-status is
configured by the [service-profile](../use/service-profiles.md).

### JSON

```json
{
    "applicationID": "123",
    "applicationName": "temperature-sensor",
    "deviceName": "garden-sensor",
    "devEUI": "AgICAgICAgI=",
    "margin": 6,
    "externalPowerSource": false,
    "batteryLevelUnavailable": false,
    "batteryLevel": 75.5,
    "tags": {
        "key": "value"
    }
}
```

### Protobuf

This message is defined by the `StatusEvent` Protobuf message.

## join

Event published when a device joins the network. Please note that this is sent
after the first received uplink (data) frame.


### JSON

```json
{
    "applicationID": "123",
    "applicationName": "temperature-sensor",
    "deviceName": "garden-sensor",
    "devEUI": "AgICAgICAgI=",
    "devAddr": "AFE5Qg==",
    "rxInfo": [
        {
            "gatewayID": "AwMDAwMDAwM=",
            "time": "2019-11-08T13:59:25.048445Z",
            "timeSinceGPSEpoch": null,
            "rssi": -48,
            "loRaSNR": 9,
            "channel": 5,
            "rfChain": 0,
            "board": 0,
            "antenna": 0,
            "location": {
                "latitude": 52.3740364,
                "longitude": 4.9144401,
                "altitude": 10.5
            },
            "fineTimestampType": "NONE",
            "context": "9u/uvA==",
            "uplinkID": "jhMh8Gq6RAOChSKbi83RHQ=="
        }
    ],
    "txInfo": {
        "frequency": 868100000,
        "modulation": "LORA",
        "loRaModulationInfo": {
            "bandwidth": 125,
            "spreadingFactor": 11,
            "codeRate": "4/5",
            "polarizationInversion": false
        }
    },
    "dr": 1,
    "tags": {
        "key": "value"
    }
}
```

### Protobuf

This message is defined by the `JoinEvent` Protobuf message.

## ack

Event published on downlink frame acknowledgements.


### JSON

```json
{

    "applicationID": "123",
    "applicationName": "temperature-sensor",
    "deviceName": "garden-sensor",
    "devEUI": "AgICAgICAgI=",
    "rxInfo": [
        {
            "gatewayID": "AwMDAwMDAwM=",
            "time": "2019-11-08T13:59:25.048445Z",
            "timeSinceGPSEpoch": null,
            "rssi": -48,
            "loRaSNR": 9,
            "channel": 5,
            "rfChain": 0,
            "board": 0,
            "antenna": 0,
            "location": {
                "latitude": 52.3740364,
                "longitude": 4.9144401,
                "altitude": 10.5
            },
            "fineTimestampType": "NONE",
            "context": "9u/uvA==",
            "uplinkID": "jhMh8Gq6RAOChSKbi83RHQ=="
        }
    ],
    "txInfo": {
        "frequency": 868100000,
        "modulation": "LORA",
        "loRaModulationInfo": {
            "bandwidth": 125,
            "spreadingFactor": 11,
            "codeRate": "4/5",
            "polarizationInversion": false
        }
    },
	"acknowledged": true,
	"fCnt": 15,
    "tags": {
        "key": "value"
    }
}
```

### Protobuf

This message is defined by the `AckEvent` Protobuf message.

## txack

Event published when a downlink frame has been acknowledged by the gateway
for transmission.


### JSON

```json
{
    "applicationID": "123",
    "applicationName": "temperature-sensor",
    "deviceName": "garden-sensor",
    "devEUI": "0202020202020202",
    "fCnt": 12,
    "tags": {
        "key": "value"
    }
}
```

### Protobuf

This message is defined by the `TxAckEvent` Protobuf message.

## error

Event published in case of an error related to payload scheduling or handling.
E.g. in case when a payload could not be scheduled as it exceeds the maximum
payload-size.

### JSON

```json
{
    "applicationID": "123",
    "applicationName": "temperature-sensor",
    "deviceName": "garden-sensor",
    "devEUI": "AgICAgICAgI=",
	"type": "UPLINK_CODEC",
	"error": "...",
    "tags": {
        "key": "value"
    }
}
```

### Protobuf

This message is defined by the `ErrorEvent` Protobuf message.
