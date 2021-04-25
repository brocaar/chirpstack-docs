# Configuration

ChirpStack Concentratord provides different binaries in order to target different
hardware platforms. Currently these are SX1301/8 and SX1302 based gateways.

ChirpStack Concentratord makes it possible to use multiple configuration files,
by repeating the `-c` argument. For example:

* `global.toml` for generic configuration
* `region.toml` for region specific settings (e.g. Class-B beacon configuration)
* `channels.toml` to configure the channels

When using ChirpStack Concentratord within a region supporting multiple
sub-bands, this means that only the `channels.toml` file needs to be updated
when changing the sub-band, avoiding duplication.

## `chirpstack-concentratord-sx1301`

The `chirpstack-concentratord-sx1301` binary has the following command-line
flags:

```text
concentratord {{ concentratord.version }}
Orne Brocaar <info@brocaar.com>
LoRa concentrator HAL daemon (v1)

USAGE:
    chirpstack-concentratord-sx1301 [OPTIONS]

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -c, --config <FILE>...    Path to configuration file
```

### Configuration example

```toml
# Concentratord configuration.
[concentratord]

# Log level.
#
# Valid options are:
#   * TRACE
#   * DEBUG
#   * INFO
#   * WARN
#   * ERROR
#   * OFF
log_level="INFO"

# Log to syslog.
#
# When set to true, log messages are being written to syslog instead of stdout.
log_to_syslog=false

# Statistics interval.
stats_interval="30s"

  # Configuration for the (ZeroMQ based) API.
  [concentratord.api]

  # Event PUB socket bind.
  event_bind="ipc:///tmp/concentratord_event"

  # Command REP socket bind.
  command_bind="ipc:///tmp/concentratord_command"


# LoRa gateway configuration.
[gateway]

# Antenna gain (dB).
antenna_gain=0

# Public LoRaWAN network.
lorawan_public=true

# Gateway vendor / model.
#
# This configures various vendor and model specific settings like the min / max
# frequency and TX gain table.
model=""

# Gateway vendor / model flags.
#
# Flag can be used to configure additional vendor / model features. The
# following flags can be used:
#
#   Global flags:
#     GNSS - Enable GNSS / GPS support
#
#   Multitech:
#     AP1  - Module is in AP1 slot (default)
#     AP2  - Module is in AP2 slot
model_flags=[]

# Gateway ID.
gateway_id="0202030405060708"


  # LoRa concentrator configuration.
  [gateway.concentrator]

  # Multi spreading-factor channels (LoRa).
  multi_sf_channels=[
    868100000,
    868300000,
    868500000,
    867100000,
    867300000,
    867500000,
    867700000,
    867900000,
  ]

  # LoRa std channel (single spreading-factor).
  [gateway.concentrator.lora_std]
  frequency=868300000
  bandwidth=250000
  spreading_factor=7

  # FSK channel.
  [gateway.concentrator.fsk]
  frequency=868800000
  bandwidth=125000
  datarate=50000


  # Beacon configuration.
  #
  # This requires a gateway with GPS / GNSS.
  #
  # Please note that the beacon settings are region dependent. The correct
  # settings can be found in the LoRaWAN Regional Parameters specification.
  [gateway.beacon]

  # Compulsory RFU size.
  compulsory_rfu_size=2

  # Beacon frequency / frequencies (Hz).
  frequencies=[
  	869525000,
  ]

  # Bandwidth (Hz).
  bandwidth=125000

  # Spreading factor.
  spreading_factor=9

  # TX power.
  tx_power=14


  # Static gateway location.
  [gateway.location]

  # When set to non-zero values, the static gateway location will be reported
  # when the gateway does not have a GNSS module or when no GNSS location fix
  # is available.
  latitude=0
  longitude=0
  altitude=0
```

## `chirpstack-concentratord-sx1302`

The `chirpstack-concentratord-sx1302` binary has the following command-line
flags:

```text
chirpstack-concentratord-sx1302 {{ concentratord.version }}
Orne Brocaar <info@brocaar.com>
LoRa concentrator HAL daemon for SX1302

USAGE:
    chirpstack-concentratord-sx1302 [OPTIONS]

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -c, --config <FILE>...    Path to configuration file
```

### Configuration example

!!! info

	Unlike the `-sx1301` binary, there is no option for configuring the
	`gateway_id`. The unique gateway ID is embedded in the SX1302 and read
	by the Concentratord daemon.

```toml
# Concentratord configuration.
[concentratord]

# Log level.
#
# Valid options are:
#   * TRACE
#   * DEBUG
#   * INFO
#   * WARN
#   * ERROR
#   * OFF
log_level="INFO"

# Log to syslog.
#
# When set to true, log messages are being written to syslog instead of stdout.
log_to_syslog=false

# Statistics interval.
stats_interval="30s"

  # Configuration for the (ZeroMQ based) API.
  [concentratord.api]

  # Event PUB socket bind.
  event_bind="ipc:///tmp/concentratord_event"

  # Command REP socket bind.
  command_bind="ipc:///tmp/concentratord_command"


# LoRa gateway configuration.
[gateway]

# Antenna gain (dB).
antenna_gain=0

# Public LoRaWAN network.
lorawan_public=true

# Gateway vendor / model.
#
# This configures various vendor and model specific settings like the min / max
# frequency and TX gain table.
model=""

# Gateway vendor / model flags.
#
# Flag can be used to configure additional vendor / model features. The
# following flags can be used:
#
#   Global flags:
#     GNSS - Enable GNSS / GPS support
#     USB  - Use USB for concentrator communication (default is SPI)
model_flags=[]


  # LoRa concentrator configuration.
  [gateway.concentrator]

  # Multi spreading-factor channels (LoRa).
  multi_sf_channels=[
    868100000,
    868300000,
    868500000,
    867100000,
    867300000,
    867500000,
    867700000,
    867900000,
  ]

  # LoRa std channel (single spreading-factor).
  [gateway.concentrator.lora_std]
  frequency=868300000
  bandwidth=250000
  spreading_factor=7

  # FSK channel.
  [gateway.concentrator.fsk]
  frequency=868800000
  bandwidth=125000
  datarate=50000


  # Static gateway location.
  [gateway.location]

  # When set to non-zero values, the static gateway location will be reported
  # when the gateway does not have a GNSS module or when no GNSS location fix
  # is available.
  latitude=0
  longitude=0
  altitude=0
```
