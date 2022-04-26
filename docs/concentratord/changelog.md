---
description: ChirpStack Concentratord changelog.
---

# Changelog

## v3.3.2

This release adds the following gateway model configurations:

* pi_supply_lora_gateway_hat_as923
* pi_supply_lora_gateway_hat_in865
* pi_supply_lora_gateway_hat_kr920
* pi_supply_lora_gateway_hat_ru864
* rak_2246_cn470
* rak_2246_eu433
* rak_2247_as923
* rak_2247_au915
* rak_2247_cn433
* rak_2247_eu433
* rak_2247_eu868
* rak_2247_in865
* rak_2247_kr920
* rak_2247_ru864
* rak_2247_us915
* rak_2287_cn470
* rak_2287_eu433
* rak_5146_as923
* rak_5146_au915
* rak_5146_cn470
* rak_5146_eu433
* rak_5146_eu868
* rak_5146_in865
* rak_5146_kr920
* rak_5146_ru864
* rak_5146_us915

## v3.3.1

This release adds the following gateway model configurations:

* imst_ic880a_ru864
* imst_ic880a_in865
* pi_supply_lora_gateway_hat_au915
* risinghf_rhf0m301_eu868
* risinghf_rhf0m301_us915

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
