---
description: Guide on getting started with ChirpStack Application Server and Pilot Things.
---

# Getting started with Pilot Things

[Pilot Things](https://www.pilot-things.com/) is a platform to manage, visualize, automate and integrate your IoT device fleet.

After following this guide, ChirpStack Application Server will be setup to forward device data to [Pilot Things](https://www.pilot-things.com/).

## Setting up gateway and device

The steps to setup your first gateway and device are covered in the [First Gateway and Device](first-gateway-device.md) guide.

Before you continue, make sure everything works up to this point.

## Integrate ChirpStack Application Server with ThingsBoard

### Get Auth Token

In order to let ChirpStack Application Server push data to your Pilot Things server, you need to obtain a Pilot Things _Authentication Token_. To obtain an authentication token, send an email to [contact@pilot-things.com](mailto:contact@pilot-things.com).

### Setup ChirpStack Application Server ThingsBoard integration

In the ChirpStack Application Server web-interface, navigate to the Application to you want to add the integration. Find the Pilot Things integration under **Integration**, then click **Add**. Fill in the two required fields:

* Pilot Things server: This is the URL you normally would use to access the Pilot Things user interface. For example, https://kerlink.pilot-things.com/.
* Authentication token: This is the token you got in an [earlier step](#get-auth-token).

## Validate integration

If you completed all the steps, then Pilot Things is ready to receive uplink data and ChirpStack Application Server is setup to forward data for your Device, using the _Authorization Token_ for authentication.

The last step is to let your device send some data and validate that this data is received by Pilot Things. You will find this data under the _Activity_ tab when navigating to the Device within Pilot Things.
