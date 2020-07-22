---
description: Read metrics from the Prometheus metrics endpoint.
---

# Prometheus metrics

ChirpStack Geolocation Server provides a [Prometheus](https://prometheus.io/) metrics endpoint
for monitoring the performance of the ChirpStack Geolocation Server service. Please refer to
the [Prometheus](https://prometheus.io/) website for more information on
setting up and using Prometheus.

## Configuration

Please refer to the [Configuration documentation](../install/config.md).

## Metrics

### Go runtime metrics

These metrics are prefixed with `go_` and provide general information about
the process like:

* Garbage-collector statistics
* Memory usage
* Go go-routines

### gRPC API metrics

These metrics are prefixed with `grpc_` and provide metrics about the gRPC
API (used by [ChirpStack Network Server](/network-server/)), e.g.:

* The number of times each API was called
* The duration of each API call (if enabled in the [Configuration](../install/config.md)

### Geolocation backend metrics

Please refer to the backend documentation for information about
the exposed metrics per backend.
