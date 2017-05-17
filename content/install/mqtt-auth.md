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

[Mosquitto](https://mosquitto.org) offers multiple ways to handle
authentication and authoriszation:

#### Password file

Using the `mosquitto_passwd` command, it is possible to create a password file
for authentication. Note that this does not handle authorization (which user
has permission to access which topic).

Example to create a password file and add add an username:

```bash
sudo mosquitto_passwd -c /etc/mosquitto/passwd <user_name>
```

Then edit the `/etc/mosquitto/mosquitto.conf` config file so that it contains
the following entries:

```
password_file /etc/mosquitto/passwd
allow_anonymous false
```

#### Mosquitto Auth Plugin

The [mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug)
project provides various backends to setup authentication
including authorization (defining which user has access to which MQTT topic).
You could use the (hashed) user passwords from
[LoRa App Server](/lora-app-server/) as the format of the hash is compatible
with this plugin. Please refer to the
[mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug)
project for more information and instructions.