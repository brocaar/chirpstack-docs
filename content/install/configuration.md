---
title: Configuration
menu:
    main:
        parent: install
        weight: 7
---

## Configuring the LoRa Server project applications

The LoRa server project components are configured by configuration files,
which are by default loaded from (in this order):

* `NAME.toml` (current working directory)
* `~/.config/NAME/NAME.toml`
* `/etc/NAME/NAME.toml`

`NAME` must be replaced by the executable name (e.g. `loraserver`, 
`lora-app-server`, `lora-gateway-bridge`).

To load a configuration file from an alternative location, use
the `--config` or `-c` flag.

To print a (new) configuration file, use the `configfile` sub-command. Example:

```bash
loraserver configfile > loraserver.toml
```

### Configuration reference

* [LoRa App Server configuration](/lora-app-server/install/config/)
* [LoRa Server configuration](/loraserver/install/config/)
* [LoRa Gateway Bridge configuration](/lora-gateway-bridge/config/)
* [LoRa Channel Manager configuration](/lora-channel-manager/config/)
