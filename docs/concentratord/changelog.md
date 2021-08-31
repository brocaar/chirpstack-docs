---
description: ChirpStack Concentratord changelog.
---

# Changelog

## v3.3.0

### Features

* Implement and expose various gateway stats aggregations (uplinks / downlinks per frequency and modulation parameters and downlinks per ack status).

### Bugfixes

* Remove Class-B beacon frequency correction on enqueue.

## v3.2.0

### Features

* Implement support for 2.4 GHz concentrator.
* Add `configfile` sub-command to binaries for printing the configuration template.

## v3.1.0

### Features

* Add support for setting static gateway location.
* Add support for RAK2287 module (SPI & USB). ([#16](https://github.com/brocaar/chirpstack-concentratord/pull/16))

### Improvements

* Update SX1302 HAL to v2.1.0.

## v3.0.3

### Bugfixes

* Fix `channel_max` calculation.

### Features

* Implement overwriting reset-pin for RPi shields.

## v3.0.2

### Bugfixes

* Fix sending 0, 0 coordinates when GPS is unavailable.

## v3.0.1

### Bugfixes

* Fix beacon loop termination on re-configuration and improve debug logging.

## v3.0.0

Initial stable release.
