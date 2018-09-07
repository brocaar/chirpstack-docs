---
title: MQTT authentication
menu:
    main:
        parent: guides
        weight: 3
---

# MQTT authentication & authorization

The LoRa Server project does not handle MQTT authentication and authorization
for you, this is the responsibility of the MQTT broker. To restrict gateways
and applications so that they can only publish and subscribe to their own
MQTT topics, it is recommended to setup MQTT authentication and authorization.

For example, you could give every gateway its own login restricted to its
own set of MQTT topics and you could give each user its own login, restricted
to a set of applications.

## Mosquitto

For [Mosquitto](https://mosquitto.org) there are multiple ways to setup
authentication and authorization. This can be pre-configured in so called
password and ACL (access control list) files and / or can be retrieved
dynamically from the [LoRa App Server](/lora-app-server/) user table (stored
in the database). In the latter case, LoRa App Server users are able to login
with their own credentials when connecting the MQTT broker and are limited to
the applications to which they have access (in the LoRa App Server
web-interface).

### Static password and ACL file

These steps describe how to setup Mosquitto with a static password and ACL
file. In case you would like to setup Mosquitto so that users and permissions
are retrieved from LoRa App Server, go to the next section (Mosquitto Auth Plugin).

#### Passwords

Using the `mosquitto_passwd` command, it is possible to create a password file
for authentication.

Example to create a password file and add add an username (use the `-c` only
the first time as it will create a new file):

```bash
# Create a password file, with users loraserver_gw, loraserver_ns, loraserver_as
# and bob.
sudo mosquitto_passwd -c /etc/mosquitto/passwd loraserver_gw
sudo mosquitto_passwd /etc/mosquitto/passwd loraserver_ns
sudo mosquitto_passwd /etc/mosquitto/passwd loraserver_as
sudo mosquitto_passwd /etc/mosquitto/passwd bob

# Secure the password file
sudo chmod 600 /etc/mosquitto/passwd
```

#### ACLs

The access control list file will map usernames to a set of topics.
Write this file to `/etc/mosquitto/acls`. An
example is:

```
user loraserver_gw
topic write gateway/+/stats
topic write gateway/+/rx
topic read gateway/+/tx

user loraserver_ns
topic read gateway/+/stats
topic write gateway/+/tx
topic read gateway/+/rx

user loraserver_as
topic write application/+/device/+/rx
topic write application/+/device/+/join
topic write application/+/device/+/ack
topic write application/+/device/+/error
topic read application/+/device/+/tx

user bob
topic read application/123/device/+/+
topic write application/123/device/+/tx
```

The access parameter for each topic can be `read`, `write` or `readwrite`.
Note that `+` is a wildcard character (e.g. all gateways, applications or
devices in the above example).

#### Mosquitto configuration

Then add a new configuration file called `/etc/mosquitto/conf.d/auth.conf` with
the following configuration:

```
allow_anonymous false
password_file /etc/mosquitto/passwd
acl_file /etc/mosquitto/acls
```

### Mosquitto Auth Plugin

To setup Mosquitto so that it retrieves the users and permissions from the
[LoRa App Server](/lora-app-server/) database, you need to setup the
[mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug)
plugin. This project provides authentication and authorization to Mosquitto
using various backends. In our case we're interested in the PostgreSQL and
files backend.

#### Compile mosquitto-auth-plug

Before the mosquitto-auth-plugin can be compiled, you need to install the
following requirements (can be installed using `sudo apt-get install ...`):

* git
* mosquitto-dev
* libmosquitto-dev
* postgresql-server-dev-9.6
* build-essential
* libssl-dev

Next clone the [mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug)
repository:

```bash
cd /opt
sudo git clone https://github.com/jpmens/mosquitto-auth-plug.git
```

Write the following content to `/opt/mosquitto-auth-plug/config.mk`:

```
# Select your backends from this list
BACKEND_CDB ?= no
BACKEND_MYSQL ?= no
BACKEND_SQLITE ?= no
BACKEND_REDIS ?= no
BACKEND_POSTGRES ?= yes
BACKEND_LDAP ?= no
BACKEND_HTTP ?= no
BACKEND_JWT ?= no
BACKEND_MONGO ?= no
BACKEND_FILES ?= yes
BACKEND_MEMCACHED ?= no

# Specify the path to the Mosquitto sources here
MOSQUITTO_SRC =

# Specify the path the OpenSSL here
OPENSSLDIR = /usr

# Specify optional/additional linker/compiler flags here
# On macOS, add 
#	CFG_LDFLAGS = -undefined dynamic_lookup
# as described in https://github.com/eclipse/mosquitto/issues/244
CFG_LDFLAGS =
CFG_CFLAGS = -DRAW_SALT
```

Compile the plugin:

```bash
cd /opt/mosquitto-auth-plug
sudo make
```

#### Configure mosquitto-auth-plug

Create a directory and empty files for additional static passwords and ACLs:

```bash
sudo mkdir /etc/mosquitto/mosquitto-auth-plug
sudo touch /etc/mosquitto/mosquitto-auth-plug/passwords
sudo touch /etc/mosquitto/mosquitto-auth-plug/acls
```

Write the following content to `/etc/mosquitto/conf.d/mosquitto-auth-plug.conf`:

```text
allow_anonymous false

auth_plugin /opt/mosquitto-auth-plug/auth-plug.so
auth_opt_backends files,postgres

auth_opt_host localhost
auth_opt_port 5432
auth_opt_dbname loraserver_as
auth_opt_user loraserver_as
auth_opt_pass loraserver_as
auth_opt_userquery select password_hash from "user" where username = $1 and is_active = true limit 1
auth_opt_superquery select count(*) from "user" where username = $1 and is_admin = true
auth_opt_aclquery select distinct 'application/' || a.id || '/#' from "user" u inner join organization_user ou on ou.user_id = u.id inner join organization o on o.id = ou.organization_id inner join application a on a.organization_id = o.id where u.username = $1 and $2 = $2

auth_opt_password_file /etc/mosquitto/mosquitto-auth-plug/passwords
auth_opt_acl_file /etc/mosquitto/mosquitto-auth-plug/acls
```

You might want to change the following configuration, to match your
[LoRa App Server](/lora-app-server/) configuration:

* `auth_opt_host`: database hostname
* `auth_opt_port`: database port
* `auth_opt_dbname`: database name
* `auth_opt_user`: database username
* `auth_opt_pass`: database password

#### Static passwords

As [LoRa Gateway Bridge](/lora-gateway-bridge/), [LoRa Server](/loraserver/)
and [LoRa App Server](/lora-app-server/) also make use of MQTT, you might want
to configure static passwords for these services.

To generate a password readable by mosquitto-auth-plug, use the following
command:

```bash
/opt/mosquitto-auth-plug/np
```

This will prompt to enter the password, and returns a hashed version of it,
example:

```text
root@ubuntu-xenial:/# /opt/mosquitto-auth-plug/np
Enter password:
Re-enter same password:
PBKDF2$sha256$901$MLnQ3iGgLsHQd8Ym$fXD7OFInOk31jAc28O9xSHoMue0zF1SR
```

You now need to write this output to the `/etc/mosquitto/mosquitto-auth-plug/passwords`
file, where each line is in the format `USERNAME:PASSWORDHASH`. In the end your
passwords file should look like this:

```text
loraserver_gw:PBKDF2$sha256$100000$7GKPpz5FmcthzI8P$hNljou3w7CIoZMoIN7cj/H8CHnP9770t
loraserver_ns:PBKDF2$sha256$100000$jXjd9LKwjkLhec/m$qwhGxiPON/tKCXcfS6fpfAr1xQec8AQI
loraserver_as:PBKDF2$sha256$100000$AC51663HqjWlPisA$uV4WQmy0c6nMsLwEffXUeVqIFRDb4Y+h
```

#### Static ACLs

For the static passwords created in the previous step, you probably want to
limit these logins to a certain set of topics. For this you can add ACL rules
to limit the set of topics per username in the file
`/etc/mosquitto/mosquitto-auth-plug/acls`. An example:

```
user loraserver_gw
topic write gateway/+/stats
topic write gateway/+/rx
topic read gateway/+/tx

user loraserver_ns
topic read gateway/+/stats
topic write gateway/+/tx
topic read gateway/+/rx

user loraserver_as
topic write application/+/device/+/rx
topic write application/+/device/+/join
topic write application/+/device/+/ack
topic write application/+/device/+/error
topic read application/+/device/+/tx
```