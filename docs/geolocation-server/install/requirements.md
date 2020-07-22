---
description: Instructions how to setup the ChirpStack Geolocation Server requirements.
---

# Requirement

## Gateway

When using geolocation, you need a LoRa<sup>&reg;</sup> gateway which provides geolocation
capabilities. These gateways provide accurate timestamps, which is a
requirement when doing geolocation based on time-difference of arrival.

### Tested gateways

The following gateways have been tested:

* [Kerlink iBTS](https://www.kerlink.com/product/wirnet-ibts/)


## Decryption key

Gateways implementing the Semtech v2 reference design will encrypt the
fine-timestamp. Before being able to use this timestamp for geolocation,
you must therefore request a decryption key. Contact your gateway vendor
or Semtech for more information.

One you have this decryption key, you must set this in the gateway configuration
so that ChirpStack Network Server is able to decrypt the timestamp before sending it
to ChirpStack Geolocation Server.
Please refer to [ChirpStack Application Server Gateway Management](/application-server/use/gateways/) for
more information.
