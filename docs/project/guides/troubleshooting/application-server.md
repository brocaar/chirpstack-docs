---
description: Troubleshooting ChirpStack Application Server related issues.
---

# ChirpStack Application Server

This guide helps you to troubleshoot application-server related connectivity issues.
This guide assumes that your gateway is connected, the ChirpStack Gateway Bridge
is publishing the received data and that data is succesfully processed by the
ChirpStack Network Server component. If you are not sure, please refer to the
[ChirpStack Network Server troubleshooting](network-server.md) guide.
It also assumes that you have installed the [ChirpStack Network Server](../../../network-server/index.md)
and [ChirpStack Application Server](../../../application-server/index.md) components and that you have
[setup your Gateway and Device](../first-gateway-device.md)
using the web-interface.

In this guide we will validate:

* If [ChirpStack Application Server](../../../application-server/index.md) receives data from [ChirpStack Network Server](../../../network-server/index.md)
* If the application payload is decrypted correctly
* If the application payload is succesfully published over MQTT
* If the web-interface is able to connect to the WebSocket API

## ChirpStack Application Server receives data?


To find out if [ChirpStack Application Server](../../../application-server/index.md) is receiving messages from 
[ChirpStack Network Server](../../../network-server/index.md), you should refer to the logs. Depending how ChirpStack
Application Server was installed, one of the following commands will show you the logs:

* `journalctl -f -n 100 -u chirpstack-application-server`
* `tail -f -n 100 /var/log/chirpstack-application-server/chirpstack-application-server.log`

Something like the following log is expected on an uplink device payload:

```text
INFO[0186] handler/mqtt: publishing message              qos=0 topic=application/1/device/0101010101010101/rx
INFO[0186] finished unary call with code OK              grpc.code=OK grpc.method=HandleUplinkData grpc.request.deadline="2018-09-24T10:54:37+02:00" grpc.service=as.ApplicationServerService grpc.start_time="2018-09-24T10:54:36+02:00" grpc.time_ms=6.989 peer.address="[::1]:63536" span.kind=server system=grpc
```

In the above log, ChirpStack Application Server received an uplink application-payload from
ChirpStack Network Server and published this payload to the `application/1/device/0101010101010101/rx`
MQTT topic.

### No log output?

Please refer to the [Network Server troubleshooting](network-server.md)
to confirm data is sent to ChirpStack Application Server on receiving an uplink payload.

## ChirpStack Applicaiton Server published data?

If you have confirmed that ChirpStack Application Server has received the uplink payload
from ChirpStack Network Server and has published this to the MQTT broker, then you can
validate this by subscribing to this topic. When using the `mosquitto_sub`
utility, you can execute the following command:

```bash
mosquitto_sub -v -t "application/#"
```

When you do not see any data appear when ChirpStack Application Server receives uplink device
payloads, then make sure the ChirpStack Application Server instance is authorized to publish
to the MQTT topic **and** the `mosquitto_sub` client is authorized to subscribe
to the given MQTT topic. This issue usually happens when you have configured
your MQTT broker so that clients need to authenticate when connecting.

## Published payload is invalid

First make sure you understand that the published payload is in [Base64](https://en.wikipedia.org/wiki/Base64)
encoding. This is an encoding to represent bytes as a string. Therefore you
first need to Base64 decode the payload to get the original slice of bytes.

In case the received payload still does not match the payload sent, make sure
the `AppSKey` is set correctly in case of an ABP device.

## Live device data / LoRaWAN<sup>&reg;</sup> frames issues

### Frames under gateway, but not under device

Please note that **all** received LoRaWAN<sup>&reg;</sup> frames are displayed on the gateway
page, but only the authenticated LoRaWAN<sup>&reg;</sup> frames are displayed on the device
page. When you **do** see frames on the gateway page, but **don't** see these
on the device page, then there is likely a MIC or frame-counter error in you
ChirpStack Network Server logs. Please refer to the [Network Server troubleshooting](network-server.md)
guide.

### Not connected to WebSocket API error

This means that your browser is unable to connect to the WebSocket API.
When your browser is able to render the web-interface, then this is likely
because you have a proxy inbetween your browser and the ChirpStack Application Server API
which is not properly configured to forward the WebSocket requests, or is
forwarding the WebSocket requests but without the authentication headers.

In the ChirpStack Application Server logs, you will see the following error:

```text
level=info msg=“finished streaming call with code Unauthenticated” error=“rpc error: code = Unauthenticated desc = authentication failed: get token from context error: no authorization-data in metadata” grpc.code=Unauthenticated grpc.method=StreamFrameLogs grpc.service=api.Gateway grpc.start_time=“2018-04-04T09:48:20+07:00” grpc.time_ms=0.07 peer.address=“127.0.0.1:60048” span.kind=server system=grpc
```

The exampe below shows how to properly configure [NGINX](http://nginx.org/)
to proxy WebSocket requests (note that you might have to change paths or
ports):

```nginx
server {
	listen 443 ssl;
	server_name localhost;

	ssl_certificate /etc/chirpstack-application-server/certs/http.pem;
	ssl_certificate_key /etc/chirpstack-application-server/certs/http-key.pem;

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
```
