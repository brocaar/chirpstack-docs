---
description: Instructions and examples how to configure the ChirpStack Application Server service.
---

# Configuration

The `chirpstack-application-server` binary has the following command-line flags:

```text
ChirpStack Application Server is an open-source application-server, part of the ChirpStack Network Server project
	> documentation & support: https://www.chirpstack.io/application-server/
	> source & copyright information: https://github.com/brocaar/chirpstack-application-server

Usage:
  chirpstack-application-server [flags]
  chirpstack-application-server [command]

Available Commands:
  configfile  Print the LoRa Application Server configuration file
  help        Help about any command
  version     Print the ChirpStack Application Server version

Flags:
  -c, --config string   path to configuration file (optional)
  -h, --help            help for chirpstack-application-server
      --log-level int   debug=5, info=4, error=2, fatal=1, panic=0 (default 4)

Use "chirpstack-application-server [command] --help" for more information about a command.
```

## Configuration file

By default `chirpstack-application-server` will look in the following order for a
configuration file at the following paths when `--config` is not set:

* `chirpstack-application-server.toml` (current working directory)
* `$HOME/.config/chirpstack-application-server/chirpstack-application-server.toml`
* `/etc/chirpstack-application-server/chirpstack-application-server.toml`

To load configuration from a different location, use the `--config` flag.

To generate a new configuration file `chirpstack-application-server.toml`, execute the following command:

```bash
chirpstack-application-server configfile > chirpstack-application-server.toml
```

Note that this configuration file will be pre-filled with the current configuration
(either loaded from the paths mentioned above, or by using the `--config` flag).
This makes it possible when new fields get added to upgrade your configuration file
while preserving your old configuration. Example:

```bash
chirpstack-application-server configfile --config chirpstack-application-server-old.toml > chirpstack-application-server-new.toml
```

Example configuration file:

```toml
--8<-- "examples/chirpstack-application-server/configuration/chirpstack-application-server.toml"
```

## Securing the application-server internal API

In order to protect the application-server internal API (`[application_server.internal_api]`) against
unauthorized access and to encrypt all communication, it is advised to use TLS
certificates. Once the `ca_cert`, `tls_cert` and `tls_key` are set, the
API will enforce client certificate validation on all incoming connections.
This means that when configuring a network-server instance in ChirpStack Application Server,
you must provide the CA and TLS client certificate in order to let the
network-server to connect to ChirpStack Application Server. See also
[network-server management]().

See [https://github.com/brocaar/chirpstack-certificates](https://github.com/brocaar/chirpstack-certificates)
for a set of script to generate such certificates.

## Securing the join-server API

In order to protect the join-server API (`[join_server]`) against
unauthorized access and to encrypt all communication, it is advised to use TLS
certificates. Once the `ca_cert`, `tls_cert` and `tls_key` are
set, the API will enforce client certificate validation on all incoming connections.
When the `ca_cert` is left blank, TLS will still be configured, but the server
will not require and validate the client-certificate.

Please note that you also need to configure ChirpStack Network Server so that it uses a
client certificate for its join-server API client. See
[ChirpStack Network Server configuration](https://www.chirpstack.io/network-server/install/config/).

## Securing the web-interface and external API

The web-interface and public api (`[application_server.external_api]`) can be
secured using a TLS certificate and key. Once the `tls_cert` and `tls_key`
are set (`[application_server.external_api]`), TLS will be activated.

### Self-signed certificate

A self-signed certificate can be generated with the following command:

```bash
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 90 -nodes
```

### Let's Encrypt

For generating a certificate with [Let's Encrypt](https://letsencrypt.org/),
first follow the [getting started](https://letsencrypt.org/getting-started/)
instructions. When the `letsencrypt` cli tool has been installed, execute:

```bash
letsencrypt certonly --standalone -d DOMAINNAME.HERE 
```

## Environment variables

Although using the configuration file is recommended, it is also possible
to use environment variables to set configuration variables. Configuration 
dots `.` are replaced with double underscores `__`.

Example:

```toml
[application_server.user_authentication.openid_connect]
client_id="my_client_id"
```

Can be set using the environment variable:

```text
APPLICATION_SERVER__USER_AUTHENTICATION__OPENID_CONNECT__CLIENT_ID="my_client_id"
```
