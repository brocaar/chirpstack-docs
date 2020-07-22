---
description: Installation of the ChirpStack Gateway Bridge on Cisco gateway.
---

# Cisco

## Cisco Wireless Gateway

* [Product detail page](https://www.cisco.com/c/en/us/products/routers/wireless-gateway-lorawan/index.html)
* [Product documentation](https://www.cisco.com/c/en/us/support/routers/interface-module-lorawan/tsd-products-support-series-home.html)
* [Firmware downloads](https://software.cisco.com/download/home/286311296/type/286311234/release/)

### Preparation

Before proceeding with the following steps, make sure you have connected the
antennas and (PoE) ethernet interface as documented by the Cisco manual.

The following steps are executed using the Cisco Console interface, for which
you need a special USB (connected to your computer) to RJ45 (connected to the
gateway) console cable. 

**Note:** the following instructions only reflect the configuration to get you
started. Please consult the Cisco Wireless Gateway documentation for a complete
manual.

### Connect to console

You can use `screen` to connect to the Cisco (serial) Console. Example:

```bash
# replace /dev/ttyUSB with the serial device
screen /dev/ttyUSB0 115200
```

After the gateway has been fully started, you will see the folliwing line:

```text
Press RETURN to get started
```

Press _RETURN_ and you will see `Gateway>` as prompt. For the configuration of
the Cisco, you need to turn on the _Privileged commands_. To do so, enter the
following command:

```shell
enable
```

The prompt should now have changed to `Gateway#`.

### Firmware version

Enter the following command to display the installed firmware version of the
gateway:

```shell
show version
```

Make sure that the version is (at least) 2.0.32. If your gateway has an older
version installed, please update it first.

### Network setup

Enter the following commands to configure the Gateway network interface:

```shell
# Configure gateway from the terminal
configure terminal

# Select interface to configure
interface FastEthernet 0/1
```

To automatically assign an IP address using DHCP:

```shell
ip address dhcp
```

To assign a static IP address to the gateway:

```shell
ip address <ip-address> <subnet-mask>
```

To save the network interface configuration:

```shell
# Set interface specific description
description Ethernet

# Exit interface configuration
exit

# Exit configuration mode
exit

# Save the configuration
copy running-config startup-config
```

To test that the ethernet interface has been configured properly, you can use
the `ping ip` command:

```shell
ping ip <ip-address>
```

### Enable GPS

Enter the following commands to enable the GPS module:

```shell
# Configure gateway from the terminal
configure terminal

# Enable UBX data in UART output
gps ubx enable

# Exit configuration mode
exit

# Save the configuration
copy running-config startup-config
```

### Packet Forwarder

The Cisco Wireless Gateway comes with an UDP Packet Forwarder for testing
purposes. This Packet Forwarder is running in an isolated LXC container. To
access the shell for this LXC container:

```shell
# Request lxc-container console access
request shell container-console
```

You will be requested to enter the _System Password_. By default this is `admin`.

#### Configuration

Create the directory to store the Packet Forwarder configuration:

```bash
mkdir /etc/lora-packet-forwarder
```

Cisco provides a list of configuration templates. To list the available
templates:

```bash
ls -l /tools/templates
```

Copy the configuration template to `/etc/lora-packet-forwarder`:

```bash
cp /tools/templates/<template> /etc/lora-packet-forwarder/config.json
```

Next update the configuration file so that it forwards the UDP data to
`localhost` on port `1700`.

```bash
vi /etc/lora-packet-forwarder/config.json
```

Under `gateway_conf` update the following keys:

* `server_address`: `"localhost"`
* `serv_port_up`: `1700`
* `serv_port_down`: `1700`

To test the Packet Forwarder, you can run the following command:

```bash
/tools/pkt_forwarder -c /etc/lora-packet-forwarder/config.json -g /dev/ttyS1
```

#### Init script

To start the Packet Forwarder automatically, you need to create an init script.
To create this file:

```bash
vi /etc/init.d/S60lora-packet-forwarder
```

Then paste the following content:

```bash
#!/bin/sh

start() {
  echo "Starting lora-packet-forwarder"
  start-stop-daemon \
  	--start \
	--background \
	--make-pidfile \
	--pidfile /var/run/lora-packet-forwarder.pid \
	--exec /tools/pkt_forwarder -- -c /etc/lora-packet-forwarder/config.json -g /dev/ttyS1
}

stop() {
  echo "Stopping lora-packet-forwarder"
  start-stop-daemon \
  	--stop \
	--oknodo \
	--quiet \
	--pidfile /var/run/lora-packet-forwarder.pid
}

restart() {
  stop
  sleep 1
  start
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  restart|reload)
    restart
    ;;
  *)
    echo "Usage: $0 {start|stop|restart}"
    exit 1
esac

exit $?
```

Then make the init script executable:

```bash
chmod +x /etc/init.d/S60lora-packet-forwarder
```

To start the Packet Forwarder manually:

```bash
/etc/init.d/S60lora-packet-forwarder start
```

The next time when the Wireless Gateway is (re)started, the Packet Forwarder
will be started automatically.

### ChirpStack Gateway Bridge

By installing the ChirpStack Gateway Bridge directly on the Wireless Gateway, it can
be directly connected to a MQTT broker. When you have exited the LXC shell,
enter it again:

```shell
# Request lxc-container console access
request shell container-console
```

#### Download

Copy the link to the latest ChirpStack Gateway Bridge **armv7.tar.gz** package from
the [Downloads](../downloads.md) page. Then:

```bash
# Create directories
mkdir -p /opt/chirpstack-gateway-bridge

# Download ChirpStack Gateway Bridge
cd /opt/chirpstack-gateway-bridge
wget <download-link>

# Uncompress archive
tar zxf *.tar.gz

# Remove archive file
rm *.tar.gz
```

#### Configuration

The ChirpStack Gateway Bridge uses a file for configuration. Please refer to
[Configuration](../install/config.md) for a full example.
Below you will find a minimal configuration example to get you started.

To create the configuration directory and create the configuration file:

```bash
mkdir /etc/chirpstack-gateway-bridge
vi /etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml
```

Then paste the following configuration and make modifications where needed:

```toml
{% raw %}
# Gateway backend configuration.
[backend]
# Backend type.
type="semtech_udp"

  # Semtech UDP packet-forwarder backend.
  [backend.semtech_udp]

  # ip:port to bind the UDP listener to
  #
  # Example: 0.0.0.0:1700 to listen on port 1700 for all network interfaces.
  # This is the listener to which the packet-forwarder forwards its data
  # so make sure the 'serv_port_up' and 'serv_port_down' from your
  # packet-forwarder matches this port.
  udp_bind = "0.0.0.0:1700"


# Integration configuration.
[integration]
# Payload marshaler.
#
# This defines how the MQTT payloads are encoded. Valid options are:
# * protobuf:  Protobuf encoding
# * json:      JSON encoding
marshaler="protobuf"

  # MQTT integration configuration.
  [integration.mqtt]
  # Event topic template.
  event_topic_template="gateway/{{ .GatewayID }}/event/{{ .EventType }}"

  # Command topic template.
  command_topic_template="gateway/{{ .GatewayID }}/command/#"

  # MQTT authentication.
  [integration.mqtt.auth]
  # Type defines the MQTT authentication type to use.
  #
  # Set this to the name of one of the sections below.
  type="generic"

    # Generic MQTT authentication.
    [integration.mqtt.auth.generic]
    # MQTT server (e.g. scheme://host:port where scheme is tcp, ssl or ws)
    server="tcp://127.0.0.1:1883"

    # Connect with the given username (optional)
    username=""

    # Connect with the given password (optional)
    password=""
{% endraw %}
```

To test the ChirpStack Gateway Bridge, you can run the following command:

```bash
/opt/chirpstack-gateway-bridge/chirpstack-gateway-bridge -c /etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml
```

#### Init script

To start the ChirpStack Gateway Bridge automatically, you need to create an init script.
To create this file:

```bash
vi /etc/init.d/S60chirpstack-gateway-bridge
```

Then paste the following content:

```bash
#!/bin/sh

start() {
  echo "Starting chirpstack-gateway-bridge"
  start-stop-daemon \
  	--start \
	--background \
	--make-pidfile \
	--pidfile /var/run/chirpstack-gateway-bridge.pid \
	--exec /opt/chirpstack-gateway-bridge/chirpstack-gateway-bridge -- -c /etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml
}

stop() {
  echo "Stopping chirpstack-gateway-bridge"
  start-stop-daemon \
  	--stop \
	--oknodo \
	--quiet \
	--pidfile /var/run/chirpstack-gateway-bridge.pid
}

restart() {
  stop
  sleep 1
  start
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  restart|reload)
    restart
    ;;
  *)
    echo "Usage: $0 {start|stop|restart}"
    exit 1
esac

exit $?
```

Then make the init script executable:

```bash
chmod +x /etc/init.d/S60chirpstack-gateway-bridge
```

To start the ChirpStack Gateway Bridge manually:

```bash
/etc/init.d/S60chirpstack-gateway-bridge start
```

The next time when the Wireless Gateway is (re)started, the ChirpStack Gateway Bridge
will be started automatically.
