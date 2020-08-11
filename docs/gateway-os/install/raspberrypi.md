---
description: Installing ChirpStack Gateway OS on a Raspberry Pi.
---

# Raspberry Pi

## Image types

There are two file types:

* `.wic.gz` - Image to use for an initial installation
* `.swu` - Software update file, see [Software update](../use/software-update.md)

## Image links

* [Raspberry Zero W](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi0-wifi/{{ gateway_os.version }}/)
* [Raspberry Pi 1](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi/{{ gateway_os.version }}/)
* [Raspberry Pi 3](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi3/{{ gateway_os.version }}/)
* [Raspberry Pi 4](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi4/{{ gateway_os.version }}/)

## SD card flashing

### Using Balena Etcher

* Download the SD card image for your Raspberry Pi (ending with `.wic.gz`).
* Click the **Flash from file** option in the [Balena Etcher](https://www.balena.io/etcher/) interface.
* Flash the SD card.
* Continue with [Using the ChirpStack Gateway OS images](../use/getting-started.md).

### Using Win32DiskImager

* Download the SD card image for your Raspberry Pi (ending with `.wic.gz`).
* Extract the `.wic.gz` image using for example [7-Zip](https://www.7-zip.org/).
* Within the [Win32DiskImager](http://sourceforge.net/projects/win32diskimager/) interface, select the extracted `.wic` file.
  Please note that you must select the `*.*` filter in the _Select a disk image_ popup before you can select the `.wic` file.
* Flash the SD card.
* Continue with [Using the ChirpStack Gateway OS images](../use/getting-started.md).

## Supported shields

The Raspberry Pi images provide out-of-the-box support for the following
concentrator shields / gateway kits:

* [IMST - iC880A](https://wireless-solutions.de/products/long-range-radio/ic880a.html)
* [IMST - iC980A](http://www.imst.com/)
* [IMST - Lite Gateway](https://wireless-solutions.de/products/long-range-radio/lora-lite-gateway.html)
* [Pi Supply - LoRa Gateway Hat](https://uk.pi-supply.com/products/iot-lora-gateway-hat-for-raspberry-pi)
* [RAK - RAK2245](https://store.rakwireless.com/products/rak2245-pi-hat)
* [RAK - RAK2246 / RAK2246G](https://store.rakwireless.com/products/rak7246-lpwan-developer-gateway)
* [RAK - RAK831 Gateway Developer Kit](https://store.rakwireless.com/products/rak831-gateway-module?variant=22375114801252)
* [RisingHF - RHF0M301 LoRaWAN IoT Discovery Kit](http://risinghf.com/#/product-details?product_id=9&lang=en)
* [Sandbox Electronics - LoRaGo PORT](https://sandboxelectronics.com/?product=lorago-port-multi-channel-lorawan-gateway)
* [Semtech - SX1302 CoreCell](https://www.semtech.com/products/wireless-rf/lora-gateways/sx1302cxxxgw1)
