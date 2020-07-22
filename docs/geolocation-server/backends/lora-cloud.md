---
description: LoRa Cloud Geolocation service by Semtech.
---

# LoRa Cloud<sup>&trade;</sup> Geolocation

[LoRa Cloud](https://www.loracloud.com/) is a service provided by [Semtech](https://www.semtech.com/)
which provides a geolocation API as a service. Please refer to the LoRa Cloud
website for more information, pricing and for signing up.

## Implementation

ChirpStack Geolocation Server integrates with the LoRa Cloud Geolocation v2 TDoA localization
algorithm. Both single-frame and multi-frame TDOA are implemented.

## Metrics

For more information on metrics, see [Prometheus metrics](../metrics/prometheus.md).
The Collos backend exposes the following metrics:

### backend_lora_cloud_api_duration_seconds

A [Histogram](https://prometheus.io/docs/concepts/metric_types/#histogram) type
metrics tracking the total number of API calls to the Collos API endpoint and
their durations (using buckets).
