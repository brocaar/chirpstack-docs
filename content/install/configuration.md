---
title: Configuration
menu:
    main:
        parent: install
        weight: 8
description: Instructions on how to configure the ChirpStack components.
---

# Configuring the ChirpStack components

The ChirpStack components are configured by configuration files,
which are by default loaded from (in this order):

* `NAME.toml` (current working directory)
* `~/.config/NAME/NAME.toml`
* `/etc/NAME/NAME.toml`

`NAME` must be replaced by the executable name (e.g. `chirpstack-network-server`, 
`chirpstack-application-server`, `chirpstack-gateway-bridge`).

To load a configuration file from an alternative location, use
the `--config` or `-c` flag.

To print a (new) configuration file, use the `configfile` sub-command. This
can not only be used to generate a new configuration file containing all the
default, but can also be used to update an existing configuration file to
include the latest defaults (maintaining the already set variables) Example:

{{<highlight bash>}}
# generate new configuration file
chirpstack-network-server configfile > chirpstack-network-server.toml

# migrate configuration file
chirpstack-network-server configfile -c chirpstack-network-server-old.toml > chirpstack-network-server.toml
{{< /highlight >}}

## Configuration reference

* [ChirpStack Gateway Bridge configuration](/gateway-bridge/install/config/)
* [ChirpStack Network Server configuration](/network-server/install/config/)
* [ChirpStack Application Server configuration](/application-server/install/config/)
* [ChirpStack Geolocation Server configuration](/geolocation-server/install/config/)
