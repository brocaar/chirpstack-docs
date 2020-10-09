---
description: Connecting LoRaWAN devices with ChirpStack and steps for troubleshooting issues.
---

# Connecting a device

This guide describes how to connect your LoRaWAN device with ChirpStack
and how to validate that it can successfully activate. At this point it is
expected that you have the [ChirpStack Network Server](../../network-server/index.md)
and [ChirpStack Application Server](../../application-server/index.md) components
installed and that you have successfully [connected a LoRa gateway](connect-gateway.md)
to it.

## Requirements

Before continuing, there are a couple things you need to know about your device.
This information is usually provided by the device vendor.

* DevEUI
* LoRaWAN MAC version implemented by the device
* Regional Parameters revision implemented by the device

### ABP

* DevAddr
* Session-keys

### OTAA

* Device root-keys (when no external join-server is used)

## Adding the device

[Login](../../application-server/use/login.md) into the [ChirpStack Application Server](../../application-server/index.md) web-interface.
The default credentials are:

* Username: admin
* Password: admin

### Device-profile

Before you can add the device to ChirpStack, you have to create a [Device-profile](../../application-server/use/device-profiles.md)
if you haven't done this already. In general it is a good practice to create
separate device-profiles for different types of devices.
A device-profile contains the capabilities of your device. For example if it
uses ABP or OTAA for activation, which LoRaWAN version and Regional Parameters
revision is implemented by the device, etc... It can also be configured with a
function to decode the payloads sent by the devices using the device-profile.

Within the ChirpStack Application Server web-interface, click
**Service-profiles** and then **Create**. Fill in the required fields and
save the device-profile.


### Application

Devices are grouped by applications. For example you could group your
temperature sensors under one application and weather stations under an other
application.

If you haven't created an application yet to which you want to add the device,
click **Applications**, then click **Create**. Fill in the required fields
and save the application.

### Device

Click the (newly created) application to which you want to add your device.
Under the **Devices** tab, click **Create**. Fill in the required fields and
select the device-profile that you want to associate with your device and
save the device.

Depending the device-profile is configured for OTAA or ABP, the next page
will ask you to enter the device root-keys (OTAA) or device session-keys (ABP).

In case your [ChirpStack Network Server](../../network-server/index.md) is
configured with a join-server and your (OTAA) device will use this join-server
for activation, then there is no need to enter the root-keys.

## Validate

After adding your LoRaWAN device to ChirpStack, validate that your device
is able activate (in case of OTAA) and send data.

When clicking the device in the [ChirpStack Application Server](../../application-server/index.md)
web-interface, open in one window the **Device data** and in an other window
the **LoRaWAN frames** tab.

Then turn on your device or trigger an uplink transmission. In case of an OTAA
device you should first see a _JoinRequest_ followed by a _JoinAccept_ message in
the **LoRaWAN frames** tab.

When the device sends its first data payload, you should also see a _Join_ and
_Up_ event in the **Device data** tab.


## Troubleshooting

When the device is not able to activate there are a several troubleshooting
steps that you can perform.

### Gateway LoRaWAN frames

After navigating to the gateway details page of a gateway close to the device, click
the **LoRaWAN frames** tab.

#### OTAA

For OTAA devices, confirm that when the device tries to OTAA activate, you see a
_JoinRequest_ message followed by a _JoinAccept_ message.

If you do not see a _JoinRequest_ and _JoinAccept_, make sure that the device
sends an OTAA request and that your gateway is correctly configured. Refer to
the [Connecting a gateway](connect-gateway.md) for validating and troubleshooting
instructions.

If you only see a _JoinRequest_ message, this means that either the OTAA
request is ignored or rejected by ChirpStack **or** that the gateway is
rejecting the _JoinAccept_ transmission. Continue with the
_Device LORaWAN frames_ section.

#### Uplink data

For ABP and OTAA devices, when the device sends an uplink payload, confirm
that you see an _UnconfirmedDataUp_ or _ConfirmedDataUp_ (depending the
uplink is of type confirmed).

If you do not see such message after the device sends an uplink, verify that
your gateway is correctly configured and able to communicate with ChirpStack.
Refer to the [Connecting a gateway](connect-gateway.md) for validating and
troubleshooting instructions.

### Device LoRaWAN frames

After you have confirmed that the gateway receives uplink frames sent by
your device, navigate to the device details page of your device and click the
**LoRaWAN frames** tab.

#### OTAA

For OTAA devices, confirm that when the device tries to OTAA activate, you see a
_JoinRequest_ message followed by a _JoinAccept_ message.

If you see a _JoinRequest_ but no _JoinAccept_, then this means that the OTAA
request is rejected by ChirpStack or that the _JoinAccept_ message is rejected
by the gateway. Continue with the _Device data_ section.

If you do not see a _JoinRequest_, but you did see a _JoinRequest_ in the
**LoRaWAN frames** tab of your gateway, then it is likely that you have
mis-configured the DevEUI of your device.

#### Uplink data

For ABP and OTAA devices, when the device sends an uplink payload, confirm
that you see an _UnconfirmedDataUp_ or _ConfirmedDataUp_ (depending the
uplink is of type confirmed).

If you do not see such message, but you did see it under the gateway
**LoRaWAN frames** tab, then it is likely that ChirpStack is unable to
authenticate the uplink frame. Continue with the next section.

### Device data

In case of a failed OTAA activation or if uplink frames are seen under the
gateway **LoRaWAN frames** but not under the device **LoRaWAN frames**, it
is very likely that there is a misconfiguration. In such case, you will find
the error message under the **Device data** tab.

#### MIC error

This means that the root or session-keys are incorrectly configured. ChirpStack
is unable to validate the Message Integrity Code (MIC) of the LoRaWAN payload.

#### Re-transmission error

This means that the device used an uplink frame-counter which has already been
seen by the [ChirpStack Network Server](../../network-server/index.md).
When a device is configured to re-transmit an uplink multiple times, then it is
likely that the first uplink has already been processed, in which case you
can ignore this error.

A common error with ABP devices is that they "forget" their frame-counters
after a power-cycle. This is against the latest specifications, but
unfortunately happens with many devices. In this case you can enable the
**Disable frame-counter check** option in the device configuration.
