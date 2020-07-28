---
description: Integrate ChirpStack with Semtech LoRa Cloud services.
---

# LoRa Cloud

The [LoRa Cloud](https://www.loracloud.com/) integration provides integrations
with several LoRa Cloud services.

## LoRa Cloud Geolocation

The LoRa Cloud Geolocation service provides a geolocation resolver for:

* TDOA (Time Difference Of Arrival)
* RSSI (signal strength)
* Wifi (Wifi access-point bssid scanning)
* GNSS (using the [LR1110](https://www.semtech.com/products/wireless-rf/lora-transceivers/lr1110))

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

## LoRa Cloud Device & Application Services (DAS)

LoRa Cloud Device & Application Services comprise a set of full life-cycle
management features for LoRa-based devices. When enabled, ChirpStack Application
Server will:

* Forward uplink meta-data for each received uplink frame
* Forward uplink meta-data + decrypted payload for each uplink received on the
  configured _DAS Modem port_.
