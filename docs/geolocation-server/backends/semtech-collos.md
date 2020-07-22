---
description: Collos geolocation service by Semtech.
---

# Semtech Collos

Collos is a platform provided by Semtech which provides localization algorithms
as a service. Please refer to [http://preview.collos.org/](http://preview.collos.org/)
for more information and for signing up.


## Implementation

ChirpStack Geolocation Server integrates with the Collos v2 TDoA localization algorithm. Both
single-frame and multi-frame TDOA are implemented.

## Metrics

For more information on metrics, see [Prometheus metrics](../metrics/prometheus.md).
The Collos backend exposes the following metrics:


### backend_collos_api_duration_seconds

A [Histogram](https://prometheus.io/docs/concepts/metric_types/#histogram) type
metrics tracking the total number of API calls to the Collos API endpoint and
their durations (using buckets).
