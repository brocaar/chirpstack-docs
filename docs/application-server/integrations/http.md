---
description: Receive LoRaWAN device-data using a HTTP endpoint.
---

# HTTP

When configured, the HTTP integration will make `POST` requests to the
configured event endpoint or endpoints (multiple URLs can ben configured, comma
separated). The `event` URL query parameter indicates the type of the event.

## Events

The GCP Pub/Sub integration exposes all events as documented by [Event types](events.md).
