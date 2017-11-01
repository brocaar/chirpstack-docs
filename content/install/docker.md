---
title: Docker install
menu:
    main:
        parent: install
        weight: 4
---

## Docker install

The LoRa Server project provides [Docker](https://www.docker.com) containers
for all project components. An overview of available containers can be found
at: [https://hub.docker.com/u/loraserver/](https://hub.docker.com/u/loraserver/).

### Versioning

* `latest` refers to the latest version from the `master` branch
* All other tags refer to tagged versions

### Docker Compose

A [Docker Compose](https://docs.docker.com/compose/) example can be found
below. Please use this `docker-compose.yml` file as a starting point, not
as a ready to use solution.

```yaml
version: "2"

services:
  loraserver:
    image: loraserver/loraserver
    environment:
      - DB_AUTOMIGRATE=true
      - LOG_NODE_FRAMES=true
      - NET_ID=010203
      - BAND=EU_863_870
      - REDIS_URL=redis://redis:6379
      - GW_MQTT_SERVER=tcp://mosquitto:1883
      - GW_SERVER_JWT_SECRET=verysecret
      - POSTGRES_DSN=postgres://loraserver_ns:loraserver_ns@postgresql_ns/loraserver_ns?sslmode=disable
      - AS_SERVER=appserver:8001

  appserver:
    image: loraserver/lora-app-server
    ports:
      - 8080:8080
    environment:
      - DB_AUTOMIGRATE=true
      - REDIS_URL=redis://redis:6379
      - POSTGRES_DSN=postgres://loraserver_as:loraserver_as@postgresql_as/loraserver_as?sslmode=disable
      - MQTT_SERVER=tcp://mosquitto:1883
      - JS_SERVER=loraserver:8003
      - JWT_SECRET=verysecret
      - HTTP_TLS_CERT=/etc/lora-app-server/certs/http.pem
      - HTTP_TLS_KEY=/etc/lora-app-server/certs/http-key.pem

  gatewaybridge:
    ports:
      - 1700:1700/udp
    image: loraserver/lora-gateway-bridge
    environment:
      - MQTT_SERVER=tcp://mosquitto:1883

  postgresql_ns:
    image: postgres:9.6-alpine
    ports:
      - 5432
    environment:
      - POSTGRES_PASSWORD=loraserver_ns
      - POSTGRES_USER=loraserver_ns
      - POSTGRES_DB=loraserver_ns

  postgresql_as:
    image: postgres:9.6-alpine
    ports:
      - 5432
    environment:
      - POSTGRES_PASSWORD=loraserver_as
      - POSTGRES_USER=loraserver_as
      - POSTGRES_DB=loraserver_as

  redis:
    ports:
      - 6379
    image: redis:4-alpine

  mosquitto:
    ports:
      - 1883
    image: eclipse-mosquitto
```
