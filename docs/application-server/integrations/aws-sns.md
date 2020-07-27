---
description: Receiving LoRaWAN device-data using AWS Simple Notification Service (SNS).
---

# AWS SNS

The [Simple Notification Service (SNS)](https://aws.amazon.com/sns/) integration
publishes all the events to a SNS Topic to which other applications or AWS
services can subscribe for further processing.

## Events

The AWS Simple Notification Service integration exposes all events as
documented by [Event types](events.md).

### Message attributes

The following message attributes are added to each published message:

* `event` - the event type
* `dev_eui` - the device EUI
* `application_id` - the ChirpStack Application Server application ID

## Example code

The following code example demonstrates how to consume integration events using
an [AWS SQS](https://aws.amazon.com/sqs/) subscription.

!!! important
	Make sure the _Enable raw message delivery_ option is enabled on the subscription.
	If not enabled, the SQS messages will not have the expected attributes.

=== "main.go"

	```go
	--8<-- "examples/chirpstack-application-server/integrations/aws-sns/go/main.go"
	```
