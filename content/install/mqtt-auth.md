---
title: MQTT authentication
menu:
    main:
        parent: install
        weight: 3
---

## MQTT authentication & authorization

The LoRa Server project does not handle MQTT authentication and authorization
(currently). To make sure that not all data is exposed to all uses, it is
advised to setup MQTT authentication & authorization.

For example, you could give every gateway its own login restricted to its
own set of MQTT topics and you could give each user its own login, restricted
to a set of applications.

### Mosquitto

[Mosquitto](https://mosquitto.org) supports authentication / authorization
through the [mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug)
plugin. **Note:** the (hashed) user password format stored by
[LoRa App Server](/lora-app-server) is compatible with the PostgreSQL backend
of this plugin.