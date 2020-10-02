---
description: Write received LoRaWAN device-data into Pilot Things.
---

# Pilot Things

When configured, the Pilot Things integration will send raw device data to the configured [Pilot Things](https://www.pilot-things.com/) instance.

## Requirements

Pilot Things requires an _Authentication Token_. This token must be correctly set in the device integration for data submission to be successful.

## Attributes

For each uplink message, ChirpStack Application Server will send Pilot Things the following data:

* RSSI
* LoRa SNR
* RF chain
* Antenna ID
* Board ID
* Device name
* Raw message data
* Device EUI
* LoRa port
* Device address
* Frame count
