---
description: Installing ChirpStack Gateway OS on a Wifx LORIX One gateway.
hide: true
---

# Wifx LORIX One

!!! warning
	Wifx gateways are no longer supported by the ChirpStack Gateway OS. Please
	note thate the latest LORIX One firmware comes with the ChirpStack
	Gateway Bridge embedded.

## LORIX One

**[LORIX One product page](https://www.lorixone.io/)**.

The LORIX One gateway is capable of booting from a SD Card. Because of this,
you don't need to overwrite the factory firmware to use the ChirpStack Gateway OS
image. To "revert" to the factory firmware, you simple remove the SD Card.

### Installation

* Download one of the provided SD Card images from the LORIX One images folder.
  Please note, there are two LORIX One versions, one with 256MB and one with
  512MB flash. Make sure you download the right version:
  * [256MB version](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/wifx/lorix-one-sd/)
  * [512MB version](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/wifx/lorix-one-512-sd/)
* Flash the SD Card image using for example [Etcher](https://www.balena.io/etcher/) on a SD Card.
* Continue with [Using the ChirpStack Gateway OS images](../use/getting-started.md).
