---
description: Instructions and examples how to configure the ChirpStack Gateway Bridge service.
---

# Configuration

The `chirpstack-gateway-bridge` has the following command-line flags:

```text
ChirpStack Gateway Bridge abstracts the packet_forwarder protocol into Protobuf or JSON over MQTT
        > documentation & support: https://www.chirpstack.io/gateway-bridge/
        > source & copyright information: https://github.com/brocaar/chirpstack-gateway-bridge

Usage:
  chirpstack-gateway-bridge [flags]
  chirpstack-gateway-bridge [command]

Available Commands:
  configfile  Print the ChirpStack Gateway Bridge configuration file
  help        Help about any command
  version     Print the ChirpStack Gateway Bridge version

Flags:
  -c, --config string   path to configuration file (optional)
  -h, --help            help for chirpstack-gateway-bridge
      --log-level int   debug=5, info=4, error=2, fatal=1, panic=0 (default 4)

Use "chirpstack-gateway-bridge [command] --help" for more information about a command.
```

## Configuration file

By default `chirpstack-gateway-bridge` will look in the following order for a
configuration at the following paths when `--config` / `-c` is not set:

* `chirpstack-gateway-bridge.toml` (current working directory)
* `$HOME/.config/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`
* `/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`

To load configuration from a different location, use the `--config` flag.

To generate a new configuration file `chirpstack-gateway-bridge.toml`, execute the following command:

```bash
chirpstack-gateway-bridge configfile > chirpstack-gateway-bridge.toml
```

Note that this configuration file will be pre-filled with the current configuration
(either loaded from the paths mentioned above, or by using the `--config` flag).
This makes it possible when new fields get added to upgrade your configuration file
while preserving your old configuration. Example:

```bash
chirpstack-gateway-bridge configfile --config chirpstack-gateway-bridge-old.toml > chirpstack-gateway-bridge-new.toml
```

Example configuration file:

```toml
--8<-- "examples/chirpstack-gateway-bridge/configuration/chirpstack-gateway-bridge.toml"
```

## Environment variables

Although using the configuration file is recommended, it is also possible
to use environment variables to set configuration variables. Configuration 
dots `.` are replaced with double underscores `__`.

Example:

```toml
[backend.semtech_udp]
udp_bind="0.0.0.0:1700"
```

Can be set using the environment variable:

```text
BACKEND__SEMTECH_UDP__UDP_BIND="0.0.0.0:1700"
```
