---
description: Receive LoRaWAN device-data using a HTTP endpoint.
---

# HTTP

When configured, the HTTP integration will make `POST` requests to the
configured event endpoint or endpoints (multiple URLs can ben configured, comma
separated). The `event` URL query parameter indicates the type of the event.

## Events

The HTTP integration exposes all events as documented by [Event types](events.md).

## Example code

!!! important
	The following code examples are for demonstration purposes only to
	demonstrate how integration events can be decoded it the most simple way
	without taking performance or security in mind. Additional code might be
	required for production usage.

### Go

The following code example demonstrates how to implement an HTTP endpoint using
[Go](https://golang.org/) which decodes either a Protobuf or JSON payload. If
you run this example on the same host as ChirpStack Application Server, then
the endpoint for the HTTP integration is `http://localhost:8090`.

=== "main.go"

	```go
	--8<-- "examples/chirpstack-application-server/integrations/http/go/main.go"
	```

=== "go.mod"

	```text
	--8<-- "examples/chirpstack-application-server/integrations/http/go/go.mod"
	```

### Python

The following code example demonstrates how to implement an HTTP endpoint using
[Python 3](https://www.python.org/) which decodes either a Protobuf or JSON
payload. If you run this example on the same host as ChirpStack Application Server,
then the endpoint for the HTTP integration is `http://localhost:8090`.

=== "main.py"

	```python
	--8<-- "examples/chirpstack-application-server/integrations/http/python/main.py"
	```

=== "requirements.txt"

	```text
	--8<-- "examples/chirpstack-application-server/integrations/http/python/requirements.txt"
	```
