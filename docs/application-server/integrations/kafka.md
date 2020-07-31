---
description: Receive LoRaWAN device-data using Kafka.
---

# Kafka

The [Kafka](https://kafka.apache.org/) integration publishes events to a Kafka
topic, using a configurable event key. Kafka uses the key for distributing
messages over partitions. You can use this to ensure some subset of messages
end up in the same partition, so they can be consumed in-order. And Kafka
can use the key for data retention decisions.  A header "event" with the
event type is included in each message. There is no need to parse it from
the key.

## Events

The GCP Pub/Sub integration exposes all events as documented by [Event types](events.md).

## Configuration examples

### Azure Event Hub

[Azure Event Hub](https://docs.microsoft.com/en-us/azure/event-hubs/event-hubs-about)
provides a [Kafka endpoint](https://docs.microsoft.com/en-us/azure/event-hubs/event-hubs-for-kafka-ecosystem-overview)
which can be used to ingest data into the Event Hub using the Kafka protocol.

You can configure the Kafka integration with an Azure Event Hub using the
following configuration example:

```toml
[application_server.integration.kafka]

# Fill in the name of your Event Hubs Namespace
brokers=["<EVENT HUBS NAMESPACE>.servicebus.windows.net:9093"]

# Fill in the name of your Event Hub
topic="<EVENT HUB>"

# TLS must be set to true
tls=true

# Leave this as-is
username="$ConnectionString"

# SAS Policy connection-string for the Event Hub
password="Endpoint=sb://.."
```
