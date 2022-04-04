---
description: Integrate ChirpStack with Semtech LoRa Cloud services.
---

# LoRa Cloud

The [LoRa Cloud](https://www.loracloud.com/) integration provides an integration
with the LoRa Cloud Modem & Geolocation Services. This service can be used for
handling LoRa Edge&trade; LR1110 and LoRa Basics&trade; Modem-E devices and / or
different forms of geolocation:

* TDOA (Time Difference Of Arrival)
* RSSI (signal strength)
* Wifi (Wifi access-point bssid scanning)
* GNSS (using the [LR1110](https://www.semtech.com/products/wireless-rf/lora-transceivers/lr1110))

## LoRa Edge&trade; LR1110 and LoRa Basics&trade; Modem-E devices

If this option is enabled, you must configure the GNSS port (default 198) and Modem port
(default 199). If ChirpStack receives an uplink on these ports, then it will forward the
uplink as the respective message type to LoRa Cloud.

If your device adheres to the LoRa Edge™ Tracker Reference Design protocol, then you can
enable this option as well. In this case ChirpStack will know how to parse the protocol
payloads to retrieve the geolocation payload from it, after which it will try to resolve
it using the LoRa Cloud geolocation resolver.

## Advanced options

For other devices, you can configure geolocation using the advanced options.

### TDOA

For TDOA based geolocation, at least three gateways are required that
implement fine-timestamping. Please refer to the ChirpStack Network Server
[Geolocation](../../network-server/features/geolocation.md) page for more
information about the fine-timestamp.

### Wifi

When using Wifi based geolocation, you must configure a payload decoder for
decoding the FRMPayload into an object expected by the LoRa Cloud integration.

Example payload decoder:

```js
--8<--- "examples/chirpstack-application-server/codecs/wifi-geolocation/decode.js"
```

The an example of the output produced by the above codec:
2

```json
{
	"access_points": [
		{
			"macAddress": "...", // base64 encoded
			"signalStrength": -70
		},
		{
			"macAddress": "...", // base64 encoded
			"signalStrength": -80
		},
		{
			"macAddress": "...", // base64 encoded
			"signalStrength": -75
		}
	]
}
```

In this case you must configure the **Wifi payload field** in the LoRa Cloud
Geolocation configuration to _access_points_.

### GNSS

When using GNSS based geolocation, you must configure a payload decoder to split
the GNSS payload part from the uplink FRMPayload. For example, your application
might not only send the GNSS payload, but also sends the number of satellites.

Example payload decoder:

```js
--8<--- "examples/chirpstack-application-server/codecs/gnss-geolocation/decode.js"
```

In the above case, you must configure the **GNSS payload field** in the LoRa Cloud
Geolocation configuration to _lr1110_gnss_.

When using the GNSS resolver, either the timestamp included the GNSS payload
or the receive timestamp (of the uplink) can be used.
