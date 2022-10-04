---
description: Installation of the ChirpStack Gateway Bridge service on the Multitech Conduit gateway.
---

# Multitech

## Multitech Conduit

* [Product detail page](https://www.multitech.com/brands/multiconnect-conduit)
* [Product documentation page](http://www.multitech.net/developer/products/multiconnect-conduit-platform/)

After completing these steps, you will have a Multitech Conduit running both the
packet-forwarder and ChirpStack Gateway Bridge. The packet-forwarder will forward
the UDP data to `localhost:1700` and the ChirpStack Gateway Bridge will forward
this data over MQTT to a MQTT broker.

<!-- toc -->

### Requirements

Before you continue, please make confirm you have the latest [mPower](http://www.multitech.net/developer/software/aep/)
firmware installed. Please refer to the [mLinux software guide](https://www.multitech.com/documents/publications/software-guides/s000727--mPower-Edge-Intelligence-Conduit-AEP-software-guide.pdf)
for more information on the firmware upgrade process.

It is also asumed that you already have completed the network setup of your
gateway and that you have obtained the IP address. In case your gateway has been
setup with DHCP, this IP address usually can be obtained by logging into your
internet router.

### Web-interface log in

The mPower web-interface can be accessed by entering `https://IP-ADDRESS/` into
your browser. As the web-interface uses a self-signed certificate, your browser
will probably raise a warning.

### Packet-forwarder setup

1. In the left menu, click _LoRaWAN&reg;_
2. Under _LoRa mode_, select **PACKET FORWARDER**
3. Under _LoRa Packet Forwarder Configuration_ enter / select the following settings:
	* Network Settings
		* Network: **Manual**
		* Channel Plan: The desired channel-plan
	* Server Settings
		* Server Address: **127.0.0.1**
		* Upstream Port: **1700**
		* Downstream Port: **1700**
4. Click _Submit_, and then _Save and Appy_ in the left menu.

After completing these steps, you should see _RUNNING_ under the Packet Forwarder Status.

### Enable SSH

In order to install ChirpStack packages on the gateway, you must first enable SSH.

1. In the left menu, click _Administration_ and then _Access configuration_. 
2. Under _SSH Settings_, make sure this option is **Enabled**.
3. In case of changes, click _Submit_ and then _Save and Apply_ in the left menu.

### SSH log in

To log in into your gateway, use the following command:

```bash
ssh USERNAME@IP-ADDRESS
```

Where `USERNAME` is the username that you use to gain access to the web-interface
of the gateway and `IP-ADDRESS` with the IP address of the gateway.

### ChirpStack Gateway Bridge install

1. Log in using SSH if you haven't done so yet (see above step).
2. Download the latest `chirpstack-gateway-bridge` `.ipk` package from [https://artifacts.chirpstack.io/vendor/multitech/conduit/](https://artifacts.chirpstack.io/vendor/multitech/conduit/).
   For example, execute the following command on the gateway:

		wget https://artifacts.chirpstack.io/vendor/multitech/conduit/chirpstack-gateway-bridge_{{ gateway_bridge.version }}-r1_arm926ejste.ipk

3. Now that the .ipk package is stored on the Conduit, you can install it using the opkg package-manager utility. Example (assuming the same .ipk file):

		sudo opkg install chirpstack-gateway-bridge_{{ gateway_bridge.version }}-r1_arm926ejste.ipk

4. Update the ChirpStack Gateway Bridge configuration. You will find the
   configuration file in the `/var/config/chirpstack-gateway-bridge` directory.
   Most likely, you want to change the following configuration:

	* MQTT connection details (with the hostname of your MQTT broker)
	* MQTT topic templates (with the correct topic prefix)

5. Restart ChirpStack Gateway Bridge:

		sudo monit restart chirpstack-gateway-bridge
