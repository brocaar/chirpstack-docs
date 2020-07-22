---
description: Decrypts the fine-timestamp of geolocation capable LoRa gateways and resolves the device location using a Geolocation Server.
---

# Geolocation

Geolocation is implemented as an [ChirpStack Application Server](../../application-server/index.md)
integration. ChirpStack Network Server simply forwards the 'fine-timestamp' (if available)
when it receives an uplink.

## Requirements

For using geolocation, you need to use gateways that are capable of providing
a "fine-timestamp". Some of these gateways encrypt this fine-timestamp.

## Decrypt fine-timestamp

When configuring the per gateway and board specific decryption key, ChirpStack Network Server
will decrypt the fine-timestamp, before forwarding it to the geolocation-server.
For getting this fine-timestamp decryption key, please contact your gateway vendor
or Semtech.
