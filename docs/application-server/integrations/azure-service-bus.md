---
description: Receive LoRaWAN device-data using the Azure Service Bus.
---

# Azure Service Bus

The [Azure Service Bus](https://azure.microsoft.com/en-us/services/service-bus/)
integration publishes all the events to a Service Bus [Topic or Queue](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-messaging-overview)
to which applications can subscribe.

## Events

The Azure Service Bus integration exposes all events as documented by [Event types](events.md).

### User properties

The following user properties are added to each published message:

* `event` - the event type
* `dev_eui` - the device EUI
* `application_id` - the ChirpStack Application Server application ID

## Example code

The following code example demonstrates how to consume integration events using
an [Azure Service-Bus Queue](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-queues-topics-subscriptions).

### Go

=== "main.go"

	```go
	--8<-- "examples/chirpstack-application-server/integrations/azure-service-bus/go/main.go"
	```

=== "go.mod"

	```text
	--8<-- "examples/chirpstack-application-server/integrations/azure-service-bus/go/go.mod"
	```

### Python

=== "main.py"

	```python
	--8<-- "examples/chirpstack-application-server/integrations/azure-service-bus/python/main.py"
	```

=== "requirements.txt"

	```text
	--8<-- "examples/chirpstack-application-server/integrations/azure-service-bus/python/requirements.txt"
	```
