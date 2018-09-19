---
title: Configuration
menu:
    main:
        parent: install
        weight: 8
description: Instructions on how to configure the LoRa Server components.
---

# Configuring the LoRa Server project applications

The LoRa server project components are configured by configuration files,
which are by default loaded from (in this order):

* `NAME.toml` (current working directory)
* `~/.config/NAME/NAME.toml`
* `/etc/NAME/NAME.toml`

`NAME` must be replaced by the executable name (e.g. `loraserver`, 
`lora-app-server`, `lora-gateway-bridge`).

To load a configuration file from an alternative location, use
the `--config` or `-c` flag.

To print a (new) configuration file, use the `configfile` sub-command. This
can not only be used to generate a new configuration file containing all the
default, but can also be used to update an existing configuration file to
include the latest defaults (maintaining the already set variables) Example:

{{<highlight bash>}}
# generate new configuration file
loraserver configfile > loraserver.toml

# migrate configuration file
loraserver configfile -c loraserver-old.toml > loraserver.toml
{{< /highlight >}}

## Configuration reference

* [LoRa App Server configuration](/lora-app-server/install/config/)
* [LoRa Server configuration](/loraserver/install/config/)
* [LoRa Gateway Bridge configuration](/lora-gateway-bridge/install/config/)
