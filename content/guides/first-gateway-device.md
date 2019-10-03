---
title: First gateway and device
menu:
    main:
        parent: guides
        weight: 4
description: When you have the LoRa Server project components up & running, this guide helps you getting started with your first gateway and device.
---

# First gateway and device

After all components are installed, you should be able to navigate to the
LoRa App Server web-interface. 

To access the LoRa App Server web-interface, enter the IP address or hostname
of you server, followed by port `8080` (this is a default configuration which
can be modified through the `lora-app-server.toml` configuration file).

If a TLS certificate has been configured (optional), use http**s://**
else use the http:// option (default). Examples:

* **http://** [http://localhost:8080/](http://localhost:8080/)
* **https://** [https://localhost:8080/](https://localhost:8080/)

## Login

The default login credentials are:

* Username: admin
* Password: admin

## Add a gateway

There are two steps involved when adding a gateway. First of all, you need
to configure your gateway so that it sends data to the
[LoRa Gateway Bridge](/lora-gateway-bridge/)
component. In the [packet-forwarder](https://github.com/Lora-net/packet_forwarder)
configuration, modify the following configuration keys:

* `server_address` to the IP address / hostname of the LoRa Gateway Bridge
* `serv_port_up` to `1700` (the default port that LoRa Gateway Bridge is using)
* `serv_port_down` to `1700` (same)

After restarting the packet-forwarder process, you should see log-lines
appearing in the LoRa Gateway Bridge logs.

The second step is to configure the gateway in your LoRa Server network. For
this, log in into the [LoRa App Server](/lora-app-server/)
web-interface and add the gateway to your organization. In case your gateway
does not have a GPS, you can set the location manually.

## Setting up your first device

Now all components are installed, you should be able to navigate to the
LoRa App Server web-interface. 

To access the LoRa App Server web-interface, enter the IP address or hostname
of you server, followed by port `8080` (this is a default configuration which
can be modified through the `lora-app-server.toml` configuration file).

Example: [https://localhost:8080/](https://localhost:8080/). 

Once you have the interface working, you are ready to add the configurations required in order to receive data from a device.

### Add network-server

In order to connect your LoRa App Server instance with the LoRa Server instance, click *Network servers* and then *Add*. As the LoRa Server is installed on the same host as the LoRa App Server in this guide, use `localhost:8000` as network-server name (port `8000` is the default port used by LoRa Server, this can be modified through `loraserver.toml`). 

Note that the LoRa App Server can connect to multiple LoRa Server instances. For example each LoRa Server instance could support a different region.

### Service-profile

The service-profile defines the features that can be used by an organization.

Click on *Service-profiles* and then *Create* to create a service-profile for the LoRa Server organization. This will also associate the organization with the network-server instance.

### Device-profile

The device-profile defines the device properties of a device. For example it defines the activation type (OTAA vs ABP), the implemented LoRaWAN version etc...

Click on *Device-profiles* and then *Create* to create a device-profile for the LoRa Server organization.

### Application

Now that there is a LoRa App Server/LoRa Server association, a service-profile for the organization and device-profile, it is time to create your first application.

Click on *Applications*, then click on *Create*.

Next, click on the created application to see the list of devices associated with this application. This will be an empty list until you complete the next step...

### Device

Click on the *Devices* tab (found under Application/_YourApp_ if you aren't there already), then click on the *Create* button to create a new device.
    
After the creation of an Over the Air Activation (OTAA) device, you will be redirected to a page where you can enter the root key(s). After the creation of an Activation By Personalization (ABP) device, you will be redirected to a page where you can enter the session keys. The selected Device Profile that was created in the steps above determines whether the device uses OTAA or ABP.

### Check that you are receiving data

It is possible to stream all LoRaWAN frames (raw and encrypted data) and
device data from the web-interface. Click on the created device and click on
the *live data* or *live LoRaWAN frames* tab. Now it is time to turn on your
device and start receiving data!

Besides seeing the data in the web-interface, you can also subscribe to the
MQTT topic to receive data, for example using the `mosquitto_sub` utility:

{{<highlight bash>}}
mosquitto_sub -v -t "#" -h localhost -p 1883
{{< /highlight >}}

Where:

* `-v` - verbose output - includes the *topic* of the message
* `-t "#"` - any message. `"#"` is a multi-level wildcard. Other possibilities
  include:
    * `"gateway/#"` - any gateway messages
    * `"application/#"` - any application messages
* `-u` - The user to log into mosquitto with
* `-P` - The password for the user
* `-h` - The host to log in to
* `-p` - The mosquitto port


Read more more about sending and receiving data in the
[LoRa App Server](/lora-app-server/) documentation.

In case you don't see any data confirm (in the logs) that:

* [LoRa Gateway Bridge](/lora-gateway-bridge/) received data from your gateway
* [LoRa Server](/loraserver/) did not return any errors
* [LoRa App Server](/lora-app-server/) did not return any errors
