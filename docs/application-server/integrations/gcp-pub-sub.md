---
description: Receive LoRaWAN device-data using Google Cloud Platform Pub/Sub.
---

# GCP Pub/Sub

The [Google Cloud Platform](https://cloud.google.com/) [Pub/Sub](https://cloud.google.com/pubsub/)
integration publishes all the events to a configurable GCP Pub/Sub topic.
Using the GCP console (or APIs) you are able to create one or multiple Pub/Sub
subscriptions, for integrating this with your application(s) or store the data
in one of the storage options provided by the Google Cloud Platform.

## Events

The GCP Pub/Sub integration exposes all events as documented by [Event types](events.md).

### Attributes

The following attributes are added to each Pub/Sub message:

* `event`: the event type
* `devEUI`: the device EUI to which the event relates

## Example code

The following code example demonstrates how to consume integration events using
a [GCP Pub/Sub Subscription](https://cloud.google.com/pubsub/docs/overview).

### Go

=== "main.go"

	```go
	--8<-- "examples/chirpstack-application-server/integrations/gcppubsub/go/main.go"
	```

=== "go.mod"

	```text
	--8<-- "examples/chirpstack-application-server/integrations/gcppubsub/go/go.mod"
	```

### Python

Please refer to the [Setting up authentication](https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-python)
section for creating a service-account and setting up the credentials.

=== "main.py"

	```python
	--8<-- "examples/chirpstack-application-server/integrations/gcppubsub/python/main.py"
	```

=== "requirements.txt"

	```text
	--8<-- "examples/chirpstack-application-server/integrations/gcppubsub/python/requirements.txt"
	```
