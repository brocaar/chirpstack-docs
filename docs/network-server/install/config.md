---
description: Instructions and examples how to configure the ChirpStack Network Server service.
---

# Configuration

To list all configuration options, start `chirpstack-network-server` with the `--help`
flag. This will display:

The `chirpstack-network-server` binary has the following command-line flags:

```text
ChirpStack Network Server is an open-source network-server, part of the ChirpStack Network Server stack
        > documentation & support: https://www.chirpstack.io/network-server/
        > source & copyright information: https://github.com/brocaar/network-server/

Usage:
  chirpstack-network-server [flags]
  chirpstack-network-server [command]

Available Commands:
  configfile  Print the ChirpStack Network Server configuration file
  help        Help about any command
  version     Print the ChirpStack Network Server version

Flags:
  -c, --config string   path to configuration file (optional)
  -h, --help            help for chirpstack-network-server
      --log-level int   debug=5, info=4, error=2, fatal=1, panic=0 (default 4)

Use "chirpstack-network-server [command] --help" for more information about a command.
```

## Configuration file

By default `chirpstack-network-server` will look in the following order for a
configuration file at the following paths when `--config` is not:

* `chirpstack-network-server.toml` (current working directory)
* `$HOME/.config/chirpstack-network-server/chirpstack-network-server.toml`
* `/etc/chirpstack-network-server/chirpstack-network-server.toml`

To load configuration from a different location, use the `--config` flag.

To generate a new configuration file `chirpstack-network-server.toml`, execute the following command:

```bash
chirpstack-network-server configfile > chirpstack-network-server.toml
```

Note that this configuration file will be pre-filled with the current configuration
(either loaded from the paths mentioned above, or by using the `--config` flag).
This makes it possible when new fields get added to upgrade your configuration file
while preserving your old configuration. Example:

```bash
chirpstack-network-server configfile --config chirpstack-network-server-old.toml > chirpstack-network-server-new.toml
```

Example configuration file:

```toml
--8<-- "examples/chirpstack-network-server/configuration/chirpstack-network-server.toml"
```

## Securing the Network Server API

In order to protect the Network Server API (`network_server.api`) against
unauthorized access and to encrypt all communication, it is advised to use
TLS certificates. Once the `ca_cert`, `tls_cert` and `tls_key` are set,
the API will enforce client certificate validation on all incoming connections.
This means that when configuring this Network Server instance in [ChirpStack Application Server](../../application-server/index.md)
you must provide the CA and TLS client certificate. See also ChirpStack Application Server
[Network Server Management](../../application-server/use/network-servers.md).

See [https://github.com/brocaar/chirpstack-certificates](https://github.com/brocaar/chirpstack-certificates)
for a set of scripts to generate such certificates.

## Join Server API configuration

In the current implementation ChirpStack Network Server uses a fixed join-server URL
(provided by ChirpStack Application Server) which is used as a Join Server backend (`join_server.default`).

In case this endpoint is secured using a TLS certificate and expects a client
certificate, you must set `ca_cert`, `tls_cert` and `tls_key`.
Also don't forget to change `server` from `http://...` to `https://...`.

See [https://github.com/brocaar/chirpstack-certificates](https://github.com/brocaar/chirpstack-certificates)
for a set of scripts to generate such certificates.

## Environment variables

Although using the configuration file is recommended, it is also possible
to use environment variables to set configuration variables. Configuration 
dots `.` are replaced with double underscores `__`.

Example:

```toml
[network_server.scheduler]
scheduler_interval="1s"
```

Can be set using the environment variable:

```text
NETWORK_SERVER__SCHEDULER__SCHEDULER_INTERVAL="1s"
```

