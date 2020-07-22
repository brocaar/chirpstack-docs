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
