---
description: Instructions and examples how to configure the ChirpStack Geolocation Server service.
---

# Configuration

To list all configuration options, start `chirpstack-geolocation-server` with the `--help`
flag. This will display:

```text
ChirpStack Geolocation Server provides geolocation services for ChirpStack Network Server
        > documentation & support: https://www.chirpstack.io/geolocation-server/
        > source & copyright information: https://github.com/brocaar/chirpstack-geolocation-server/

Usage:
  chirpstack-geolocation-server [flags]
  chirpstack-geolocation-server [command]

Available Commands:
  configfile                    Print the ChirpStack Geolocation Server configuration file
  help                          Help about any command
  test-resolve-multi-frame-tdoa Runs the resolve multi-frame TDOA request from the given directory
  test-resolve-tdoa             Runs the resolve TDOA request from the given directory
  version                       Print the ChirpStack Geolocation Server version

Flags:
  -c, --config string   path to configuration file (optional)
  -h, --help            help for chirpstack-geolocation-server
      --log-level int   debug=5, info=4, error=2, fatal=1, panic=0 (default 4)

Use "chirpstack-geolocation-server [command] --help" for more information about a command.
```

## Configuration file

By default `chirpstack-geolocation-server` will look in the following order for a
configuration file at the following paths when `--config` is not set:

* `chirpstack-geolocation-server.toml` (current working directory)
* `$HOME/.config/chirpstack-geolocation-server/chirpstack-geolocation-server.toml`
* `/etc/chirpstack-geolocation-server/chirpstack-geolocation-server.toml`

To load configuration from a different location, use the `--config` flag.

To generate a new configuration file `chirpstack-geolocation-server.toml`, execute the following command:

```bash
chirpstack-geolocation-server configfile > chirpstack-geolocation-server.toml
```

Note that this configuration file will be pre-filled with the current configuration
(either loaded from the paths mentioned above, or by using the `--config` flag).
This makes it possible when new fields get added to upgrade your configuration file
while preserving your old configuration. Example:

```bash
chirpstack-geolocation-server configfile --config chirpstack-geolocation-server-old.toml > chirpstack-geolocation-server-new.toml
```

Example configuration file:

```bash
[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level=4

# Geolocation-server configuration.
[geo_server]
  # Geolocation API.
  #
  # This is the geolocation API that can be used by ChirpStack Network Server.
  [geo_server.api]
  # ip:port to bind the api server
  bind="0.0.0.0:8005"

  # CA certificate used by the api server (optional)
  ca_cert=""

  # TLS certificate used by the api server (optional)
  tls_cert=""

  # TLS key used by the api server (optional)
  tls_key=""


  # Geolocation backend configuration.
  [geo_server.backend]
  # Type.
  #
  # The backend type to use.
  type="collos"

  # Request log directory.
  #
  # Logging requests can be used to "replay" geolocation requests and to compare
  # different geolocation backends. When left blank, logging will be disabled.
  request_log_dir=""

    # Collos backend.
    [geo_server.backend.collos]
    # Collos subscription key.
    #
    # This key can be retrieved after creating a Collos account at:
    # http://preview.collos.org/
    subscription_key=""

    # Request timeout.
    #
    # This defines the request timeout when making calls to the Collos API.
    request_timeout="1s"


    # LoRa Cloud backend.
    #
    # Please see https://www.loracloud.com/ for more information about this
    # geolocation service.
    [geo_server.backend.lora_cloud]
    # API URI.
    #
    # The URI of the Geolocation API. This URI can be found under
    # 'Token Management'.
    uri=""

    # API token.
    token=""

    # Request timeout.
    #
    # This defines the request timeout when making calls to the LoRa Cloud API.
    request_timeout="1s"


# Prometheus metrics settings.
[metrics.prometheus]
# Enable Prometheus metrics endpoint.
endpoint_enabled=false

# The ip:port to bind the Prometheus metrics server to for serving the
# metrics endpoint.
bind=""

# API timing histogram.
#
# By setting this to true, the API request timing histogram will be enabled.
# See also: https://github.com/grpc-ecosystem/go-grpc-prometheus#histograms
api_timing_histogram=false
```

## Securing the geolocation API

In order to protect the geolocation API (`geo_server.api`) against
unauthorized access and to encrypt all communication, it is advised to use
TLS certificates. Once the `ca_cert`, `tls_cert` and `tls_key` are set,
the API will enforce client certificate validation on all incoming connections.
This means that when configuring this geolocation-server instance in ChirpStack Network Server,
you must provide the CA and TLS client certificate. See also
[ChirpStack Network Server configuration](/network-server/install/config/).

See [https://github.com/brocaar/chirpstack-certificates](https://github.com/brocaar/chirpstack-certificates)
for a set of scripts to generate such certificates.
