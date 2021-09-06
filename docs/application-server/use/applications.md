---
description: Manage applications, integrations and payload decoders.
---

# Applications

An application is a collection of devices with the same purpose / of the same type.
Think of a weather station collecting data at different locations for example.

When creating an Application, you need to select the [Service Profile](service-profiles.md)
which will be used for the devices created under this Application. Note that
once a Service Profile has been selected, it can't be changed.

## Payload codecs

The payload codec options have moved to the [Device Profile](device-profiles.md).
For backward compatibility, existing codec configuration on the application is still accessible
and functional, but this will be removed fully in the next major release update.

## Integrations

For documentation on the available integrations, please refer to the
**Integrations** section in the left menu.

## Devices

Multiple [Devices](devices.md) can be added to the Application.
