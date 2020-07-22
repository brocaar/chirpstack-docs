---
description: Receive LoRaWAN device-data using AMQP / RabbitMQ.
---

# AMQP / RabbitMQ

The [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol) /
[RabbitMQ](https://www.rabbitmq.com/) integration publishes events to an AMQP
routing-key. By creating one or multiple bindings to one or multiple queues, it
is possible to subscribe to all data, or just a sub-set.

## Events

The GCP Pub/Sub integration exposes all events as documented by [Event types](events.md).
