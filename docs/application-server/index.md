# Introduction

ChirpStack Application Server is an open-source LoRaWAN<sup>&reg;</sup>
Application Server, part of the [ChirpStack](https://www.chirpstack.io/) open-source
LoRaWAN Network Server stack. It is responsible for the device "inventory"
part of a LoRaWAN infrastructure, handling of join-request and the handling
and encryption of application payloads.

It offers a [web-interface](use/login.md) where users,
organizations, applications and devices can be managed. For integration with
external services, it offers a [RESTful](integrate/rest.md) 
and [gRPC](integrate/grpc.md) API.

Device data can be [sent and / or received](integrations/events.md) over
MQTT, HTTP and be written directly into InfluxDB.

See also the complete list of [ChirpStack Application Server features](features.md).

## Screenshots

![applications](/static/img/screenshots/web_applications.png)
![nodes](/static/img/screenshots/web_nodes.png)
![node details](/static/img/screenshots/web_node_details.png)
![swagger api](/static/img/screenshots/swagger.png)
