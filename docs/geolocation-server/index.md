# Introduction

!!! warning
	The ChirpStack Geolocation Server has been deprecated. Geolocation capabilities
	have been merged into the [ChirpStack Application Server](../application-server/index.md)
	as [LoRa Cloud](../application-server/integrations/loracloud.md) integration.

ChirpStack Geolocation Server is an open-souce Geolocation server, part of the
[ChirpStack](https://www.chirpstack.io/) open-source LoRaWAN<sup>&reg;</sup> Network Server stack.
It can be used to resolve the location of devices based on TDoA (time-difference
of arrival) meta-data provided by LoRa<sup>&reg;</sup> gateways.

## Supported backends

The following backends are supported by ChirpStack Geolocation Server:

* [Semtech Collos](backends/semtech-collos.md)
* [Semtech LoRa Cloud](backends/lora-cloud.md)
