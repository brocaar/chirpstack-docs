---
title: Application-server troubleshooting
menu:
    main:
        parent: troubleshooting
description: Troubleshooting LoRa App Server (application-server) related issues.
---

# Troubleshooting application-server issues

This guide helps you to troubleshoot application-server related connectivity issues.
This guide assumes that your gateway is connected, the LoRa Gateway Bridge
is publishing the received data and that data is succesfully processed by the
LoRa Server component. If you are not sure, please refer to the
[Debugging network-server issues]({{<relref "network-server.md">}}) guide.
It also assumes that you have installed the [LoRa Server](/loraserver/)
and [LoRa App Server](/lora-app-server/) components and that you have
[setup your gateway and device]({{<ref "/guides/first-gateway-device.md">}})
through the web-interface.

In this guide we will validate:

* If [LoRa App Server](/lora-app-server/) receives data from [LoRa Server](/loraserver/)
* If the application payload is decrypted correctly
* If the application payload is succesfully published over MQTT
* If the web-interface is able to connect to the WebSocket API

## LoRa App Server receives data?


To find out if [LoRa App Server](/lora-app-server/) is receiving messages from 
[LoRa Server](/loraserver/), you should refer to the logs. Depending how LoRa
App Server was installed, one of the following commands will show you the logs:

* `journalctl -f -n 100 -u lora-app-server`
* `tail -f -n 100 /var/log/lora-app-server/lora-app-server.log`

Something like the following log is expected on an uplink device payload:

{{<highlight text>}}
INFO[0186] handler/mqtt: publishing message              qos=0 topic=application/1/device/0101010101010101/rx
INFO[0186] finished unary call with code OK              grpc.code=OK grpc.method=HandleUplinkData grpc.request.deadline="2018-09-24T10:54:37+02:00" grpc.service=as.ApplicationServerService grpc.start_time="2018-09-24T10:54:36+02:00" grpc.time_ms=6.989 peer.address="[::1]:63536" span.kind=server system=grpc
{{< /highlight >}}

In the above log, LoRa App Server received an uplink application-payload from
LoRa Server and published this payload to the `application/1/device/0101010101010101/rx`
MQTT topic.

### No log output?

Please refer to the [Network-server debugging guide]({{<relref "network-server.md">}})
to confirm data is sent to LoRa App Server on receiving an uplink payload.

## LoRa App Server publised data?

If you have confirmed that LoRa App Server has received the uplink payload
from LoRa Server and has published this to the MQTT broker, then you can
validate this by subscribing to this topic. When using the `mosquitto_sub`
utility, you can execute the following command:

{{<highlight bash>}}
mosquitto_sub -v -t "application/#"
{{< /highlight >}}

When you do not see any data appear when LoRa App Server receives uplink device
payloads, then make sure the LoRa App Server instance is authorized to publish
to the MQTT topic **and** the `mosquitto_sub` client is authorized to subscribe
to the given MQTT topic. This issue usually happens when you have configured
your MQTT broker so that clients need to authenticate when connecting.

## Published payload is invalid

First make sure you understand that the published payload is in [Base64](https://en.wikipedia.org/wiki/Base64)
encoding. This is an encoding to represent bytes as a string. Therefore you
first need to Base64 decode the payload to get the original slice of bytes.

In case the received payload still does not match the payload sent, make sure
the `AppSKey` is set correctly in case of an ABP device.

## Live device data / LoRaWAN frames issues

### Frames under gateway, but not under device

Please note that **all** received LoRaWAN frames are displayed on the gateway
page, but only the authenticated LoRaWAN frames are displayed on the device
page. When you **do** see frames on the gateway page, but **don't** see these
on the device page, then there is likely a MIC or frame-counter error in you
LoRa Server logs. Please refer to the [Debugging network-server issues]({{<relref "network-server.md">}})
guide.

### Not connected to WebSocket API error

This means that your browser is unable to connect to the WebSocket API.
When your browser is able to render the web-interface, then this is likely
because you have a proxy inbetween your browser and the LoRa App Server API
which is not properly configured to forward the WebSocket requests, or is
forwarding the WebSocket requests but without the authentication headers.

In the LoRa App Server logs, you will see the following error:

{{<highlight text>}}
level=info msg=“finished streaming call with code Unauthenticated” error=“rpc error: code = Unauthenticated desc = authentication failed: get token from context error: no authorization-data in metadata” grpc.code=Unauthenticated grpc.method=StreamFrameLogs grpc.service=api.Gateway grpc.start_time=“2018-04-04T09:48:20+07:00” grpc.time_ms=0.07 peer.address=“127.0.0.1:60048” span.kind=server system=grpc
{{< /highlight >}}

The exampe below shows how to properly configure [NGINX](http://nginx.org/)
to proxy WebSocket requests (note that you might have to change paths or
ports):

{{<highlight nginx>}}
server {
	listen 443 ssl;
	server_name localhost;

	ssl_certificate /etc/lora-app-server/certs/http.pem;
	ssl_certificate_key /etc/lora-app-server/certs/http-key.pem;

	location ~ ^/api/(gateways|devices)/(\w+)/(frames|events)$ {
		proxy_pass http://localhost:8080/api/$1/$2/$3;
		proxy_http_version 1.1;
		proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection "Upgrade";
		proxy_read_timeout 86400s;
		proxy_send_timeout 86400s;
	}

	location / {
		proxy_pass http://localhost:8080/;
	}
}
{{< /highlight >}}
