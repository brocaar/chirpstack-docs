---
description: Receive LoRaWAN device-data using a HTTP endpoint.
---

# HTTP

When configured, the HTTP integration will make `POST` requests to the
configured event endpoint or endpoints (multiple URLs can ben configured, comma
separated). The `event` URL query parameter indicates the type of the event.

## Events

The GCP Pub/Sub integration exposes all events as documented by [Event types](events.md).

## Example code

### Go

The following code example demonstrates how to implement an HTTP endpoint using
[Go](https://golang.org/) which decodes either a Protobuf or JSON payload. If
you run this example on the same host as ChirpStack Application Server, then
the endpoint for the HTTP integration is `http://localhost:8090`.

=== "main.go"

	```go
	--8<-- "examples/chirpstack-application-server/integrations/http/go/main.go"
	```
