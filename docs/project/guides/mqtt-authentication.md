---
description: Guide on how to setup MQTT (Mosquitto) topic-based authentication and authorization.
---

# MQTT authentication & authorization

The ChirpStack Network Server and Application Server do not handle MQTT authentication and authorization
for you as this is the responsibility of the MQTT broker. To restrict gateways
and applications so that they can only publish and subscribe to their own
MQTT topics, it is recommended to setup MQTT authentication and authorization.

For example, you could give every gateway its own login restricted to its
own set of MQTT topics and you could give each user its own login, restricted
to a set of applications.

## Mosquitto

For [Mosquitto](https://mosquitto.org) there are multiple ways to setup
authentication and authorization. This can be pre-configured in so called
password and ACL (access control list) files and / or can be retrieved
dynamically from the [ChirpStack Application Server](../../application-server/index.md) user table (stored
in the database). In the latter case, ChirpStack Application Server users are able to login
with their own credentials when connecting the MQTT broker and are limited to
the applications to which they have access (in the ChirpStack Application Server
web-interface).

### Static password and ACL file

These steps describe how to setup Mosquitto with a static password and ACL
file. In case you would like to setup Mosquitto so that users and permissions
are retrieved from ChirpStack Application Server, go to the next sections for instruction on
how to configure Mosquitto Auth Plugin or the alternative Mosquitto Go Auth.

#### Passwords

Using the `mosquitto_passwd` command, it is possible to create a password file
for authentication.

Example to create a password file and add add an username (use the `-c` only
the first time as it will create a new file):

```bash
# Create a password file, with users chirpstack_gw, chirpstack_ns, chirpstack_as
# and bob.
sudo mosquitto_passwd -c /etc/mosquitto/passwd chirpstack_gw
sudo mosquitto_passwd /etc/mosquitto/passwd chirpstack_ns
sudo mosquitto_passwd /etc/mosquitto/passwd chirpstack_as
sudo mosquitto_passwd /etc/mosquitto/passwd bob

# Secure the password file
sudo chmod 600 /etc/mosquitto/passwd
```

#### ACLs

The access control list file will map usernames to a set of topics.
Write this file to `/etc/mosquitto/acls`. An
example is:

```text
user chirpstack_gw
topic write gateway/+/event/+
topic read gateway/+/command/+

user chirpstackns
topic read gateway/+/event/+
topic write gateway/+/command/+

user chirpstack_as
topic write application/+/device/+/event/+
topic read application/+/device/+/command/+

user bob
topic read application/123/device/+/event/+
topic write application/123/device/+/command/+
```

The access parameter for each topic can be `read`, `write` or `readwrite`.
Note that `+` is a wildcard character (e.g. all gateways, applications or
devices in the above example).

#### Mosquitto configuration

Then add a new configuration file called `/etc/mosquitto/conf.d/auth.conf` with
the following configuration:

```text
allow_anonymous false
password_file /etc/mosquitto/passwd
acl_file /etc/mosquitto/acls
```

### Mosquitto Auth Plugin

To setup Mosquitto so that it retrieves the users and permissions from the
[ChirpStack Application Server](../../application-server/index.md) database, you need to setup the
[mosquitto-auth-plug](https://github.com/jpmens/mosquitto-auth-plug)
plugin. This project provides authentication and authorization to Mosquitto
using various backends. In our case we're interested in the PostgreSQL and
files backend.

#### Compile mosquitto-auth-plug

Before the mosquitto-auth-plugin can be compiled, you need to install the
following requirements (can be installed using `sudo apt install ...`):

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

```text
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
auth_opt_dbname chirpstack_as
auth_opt_user chirpstack_as
auth_opt_pass chirpstack_as
auth_opt_userquery select password_hash from "user" where email = $1 and is_active = true limit 1
auth_opt_superquery select count(*) from "user" where email = $1 and is_admin = true
auth_opt_aclquery select distinct 'application/' || a.id || '/#' from "user" u inner join organization_user ou on ou.user_id = u.id inner join organization o on o.id = ou.organization_id inner join application a on a.organization_id = o.id where u.email = $1 and $2 = $2

auth_opt_password_file /etc/mosquitto/mosquitto-auth-plug/passwords
auth_opt_acl_file /etc/mosquitto/mosquitto-auth-plug/acls
```

You might want to change the following configuration, to match your
[ChirpStack Application Server](/appplication-server/) configuration:

* `auth_opt_host`: database hostname
* `auth_opt_port`: database port
* `auth_opt_dbname`: database name
* `auth_opt_user`: database username
* `auth_opt_pass`: database password

#### Static passwords

As [ChirpStack Gateway Bridge](/gateway-bridge/), [ChirpStack Network Server](/network-server/)
and [ChirpStack Application Server](/application-server/) also make use of MQTT, you might want
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
chirpstack_gw:PBKDF2$sha256$100000$7GKPpz5FmcthzI8P$hNljou3w7CIoZMoIN7cj/H8CHnP9770t
chirpstack_ns:PBKDF2$sha256$100000$jXjd9LKwjkLhec/m$qwhGxiPON/tKCXcfS6fpfAr1xQec8AQI
chirpstack_as:PBKDF2$sha256$100000$AC51663HqjWlPisA$uV4WQmy0c6nMsLwEffXUeVqIFRDb4Y+h
```

#### Static ACLs

For the static passwords created in the previous step, you probably want to
limit these logins to a certain set of topics. For this you can add ACL rules
to limit the set of topics per username in the file
`/etc/mosquitto/mosquitto-auth-plug/acls`. An example:

```text
user chirpstack_gw
topic write gateway/+/event/+
topic read gateway/+/command/+

user chirpstack_ns
topic read gateway/+/event/+
topic write gateway/+/command/+

user chirpstack_as
topic write application/+/device/+/event/+
topic read application/+/device/+/command/+
```


### Alternative plugin: Mosquitto Go Auth

An alternative to the mosquitto auth plugin is [mosquitto-go-auth](https://github.com/iegomez/mosquitto-go-auth).
It also provides authentication and authorization to Mosquitto, and the most relevant differences are that it's
written in Go (easy to extend and build) and that it provides a local JWT backend. It may be used **instead**
of mosquitto-auth-plug.

#### Build

This package needs Go to be built. Check https://golang.org/dl/ for instructions on installing Go.

Start by cloning the plugin and then installing requirements (dep is used to manage dependencies and may be
installed with `make dev-requirements`):

```bash
cd go/src/github.com/iegomez/
git clone https://github.com/iegomez/mosquitto-go-auth.git
cd mosquitto-go-auth
make requirements
```

Now we need to create the `go-auth.so` shared object and the `pw` binary utility.

##### Build for mosquitto 1.4.x

Compile the plugin with:

```bash
make
```

##### Build for mosquitto 1.5.x


Export needed flags and compile the plugin with:

```bash
export CGO_CFLAGS="-I/usr/local/include -fPIC"
export CGO_LDFLAGS="-shared"
make
```

#### Configure mosquitto-go-auth

Create a directory and empty files for additional static passwords and ACLs:

```bash
sudo mkdir /etc/mosquitto/mosquitto-go-auth
sudo touch /etc/mosquitto/mosquitto-go-auth/passwords
sudo touch /etc/mosquitto/mosquitto-go-auth/acls
```

This guide assumes that you have Redis running in your host as it's a
requirement of ChirpStack Network Server. Redis is used by the plugin for cache purposes.
The cache may be disabled or configured differently, check the repo for more
details.  

Also, we'll configure the plugin with the JWT backend in local mode using
chirpstack-application-server's DB. This allows you to connect a client using a
chirpstack-application-server user's JWT token.

Write the following content to `/etc/mosquitto/conf.d/mosquitto-go-auth.conf`:

```text
auth_plugin /home/your-user/go/src/github.com/iegomez/mosquitto-go-auth/go-auth.so
auth_opt_backends files, postgres, jwt
auth_opt_check_prefix false
allow_anonymous false

auth_opt_password_path /etc/mosquitto/auth/passwords
auth_opt_acl_path /etc/mosquitto/auth/acls

auth_opt_cache true
auth_opt_cache_reset true
#Change to whatever redis DB you want to avoid messing with other services.
auth_opt_cache_db 4

auth_opt_pg_host localhost
auth_opt_pg_port 5432
auth_opt_pg_dbname chirpstack_as
auth_opt_pg_user chirpstack_as
auth_opt_pg_password chirpstack_as_password
auth_opt_pg_userquery select password_hash from "user" where email = $1 and is_active = true limit 1
auth_opt_pg_superquery select count(*) from "user" where email = $1 and is_admin = true
auth_opt_pg_aclquery select distinct 'application/' || a.id || '/#' from "user" u inner join organization_user ou on ou.user_id = u.id inner join organization o on o.id = ou.organization_id inner join application a on a.organization_id = o.id where u.email = $1 and $2 = $2

auth_opt_jwt_remote false
auth_opt_jwt_db postgres
auth_opt_jwt_secret application-server-jwt-secret
auth_opt_jwt_userquery select count(*) from "user" where email = $1 and is_active = true limit 1
auth_opt_jwt_superquery select count(*) from "user" where email = $1 and is_admin = true
auth_opt_jwt_aclquery select distinct 'application/' || a.id || '/#' from "user" u inner join organization_user ou on ou.user_id = u.id inner join organization o on o.id = ou.organization_id inner join application a on a.organization_id = o.id where u.email = $1 and $2 = $2
auth_opt_jwt_userfield Username
```

You might want to change the following configuration, to match your
[ChirpStack Application Server](../../application-server/index.md) configuration:

* `auth_plugin`: path to the generated `go-auth.so` shared object

* `auth_opt_pg_host`: database hostname
* `auth_opt_pg_port`: database port
* `auth_opt_pg_dbname`: database name
* `auth_opt_pg_user`: database username
* `auth_opt_pg_password`: database password

* `auth_opt_jwt_secret`: ChirpStack Application Server jwt secret
* `auth_opt_redis_db`: redis db to use as cache

Finally, add the following to the end of `/etc/mosquitto/mosquitto.conf` to
include the `conf.d` directory.

```text
include_dir /etc/mosquitto/conf.d
```

#### Static passwords

As [ChirpStack Gateway Bridge](../../gateway-bridge/index.md), [ChirpStack Network Server](../../network-server/index.md)
and [ChirpStack Application Server](../../application-server/index.md) also make use of MQTT, you might want
to configure static passwords for these services.

To generate a password readable by mosquitto-go-auth, you may use the `pw`
utility generated when building the plugin. It will print a hash version
of the password given to the `-p` flag.

```bash
cd go/src/github.com/iegomez/
./pw -p your-password
PBKDF2$sha512$100000$dUI9gW3+7pUXXTh49ZVT3w==$FJIajqgLKPgDaa78cJ7HkqKTsSiXxmjlpgbqDjKuAZYcSIt5x73ZPAL06b6q0gJhChIrkifD1fFiVZoae4LQgQ==
```

You may change number of iterations and algorithm (sha512 or sha256) with the
`-i` and `-a` flags. Defaults are sha512 and 1000 iterations.

```bash
cd go/src/github.com/iegomez/
./pw -p your-password -a sha256 -i 500
PBKDF2$sha256$500$MEyshu2cmqV7T4oVzvn39g==$asVpX5WPRwDy2EGTN5P6fRN+jlkg6VoSiWAGz6+9AZ4=
```

You now need to write this output to the `/etc/mosquitto/mosquitto-go-auth/passwords`
file, where each line is in the format `USERNAME:PASSWORDHASH`. In the end your
passwords file should look like this:

```text
chirpstack_gw:PBKDF2$sha256$100000$7GKPpz5FmcthzI8P$hNljou3w7CIoZMoIN7cj/H8CHnP9770t
chirpstack_ns:PBKDF2$sha256$100000$jXjd9LKwjkLhec/m$qwhGxiPON/tKCXcfS6fpfAr1xQec8AQI
chirpstack_as:PBKDF2$sha256$100000$AC51663HqjWlPisA$uV4WQmy0c6nMsLwEffXUeVqIFRDb4Y+h
```

#### Static ACLs

For the static passwords created in the previous step, you probably want to
limit these logins to a certain set of topics. For this you can add ACL rules
to limit the set of topics per username in the file
`/etc/mosquitto/mosquitto-go-auth/acls`. An example:

```text
user chirpstack_gw
topic write gateway/+/event/+
topic read gateway/+/command/+

user chirpstack_ns
topic read gateway/+/event/+
topic write gateway/+/command/+

user chirpstack_as
topic write application/+/device/+/rx
topic write application/+/device/+/join
topic write application/+/device/+/ack
topic write application/+/device/+/error
topic read application/+/device/+/tx
topic write application/+/device/+/status
topic write application/+/device/+/location
```
