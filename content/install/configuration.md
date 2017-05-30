---
title: Configuration
menu:
    main:
        parent: install
        weight: 6
---

## Configuring the LoRa Server project applications

The LoRa Server project applications can be configured via the command line or via environment variables.  The following outlines the options for each of the applications.  More information on these parameters can be found on the
documentation pages for the applications accessible via the links above.

### LoRa App Server

To list all configuration options, start lora-app-server with the --help flag. This will display:

```bash
GLOBAL OPTIONS:
   --postgres-dsn value        postgresql dsn (e.g.: postgres://user:password@hostname/database?sslmode=disable) (default: "postgres://localhost/loraserver?sslmode=disable") [$POSTGRES_DSN]
   --db-automigrate            automatically apply database migrations [$DB_AUTOMIGRATE]
   --redis-url value           redis url (e.g. redis://user:password@hostname/0) (default: "redis://localhost:6379") [$REDIS_URL]
   --mqtt-server value         mqtt server (e.g. scheme://host:port where scheme is tcp, ssl or ws) (default: "tcp://localhost:1883") [$MQTT_SERVER]
   --mqtt-username value       mqtt server username (optional) [$MQTT_USERNAME]
   --mqtt-password value       mqtt server password (optional) [$MQTT_PASSWORD]
   --ca-cert value             ca certificate used by the api server (optional) [$CA_CERT]
   --tls-cert value            tls certificate used by the api server (optional) [$TLS_CERT]
   --tls-key value             tls key used by the api server (optional) [$TLS_KEY]
   --bind value                ip:port to bind the api server (default: "0.0.0.0:8001") [$BIND]
   --http-bind value           ip:port to bind the (user facing) http server to (web-interface and REST / gRPC api) (default: "0.0.0.0:8080") [$HTTP_BIND]
   --http-tls-cert value       http server TLS certificate [$HTTP_TLS_CERT]
   --http-tls-key value        http server TLS key [$HTTP_TLS_KEY]
   --jwt-secret value          JWT secret used for api authentication / authorization [$JWT_SECRET]
   --ns-server value           hostname:port of the network-server api server (default: "127.0.0.1:8000") [$NS_SERVER]
   --ns-ca-cert value          ca certificate used by the network-server client (optional) [$NS_CA_CERT]
   --ns-tls-cert value         tls certificate used by the network-server client (optional) [$NS_TLS_CERT]
   --ns-tls-key value          tls key used by the network-server client (optional) [$NS_TLS_KEY]
   --pw-hash-iterations value  the number of iterations used to generate the password hash (default: 100000) [$PW_HASH_ITERATIONS]
   --log-level value           debug=5, info=4, warning=3, error=2, fatal=1, panic=0 (default: 4) [$LOG_LEVEL]
   --help, -h                  show help
   --version, -v               print the version
```

### LoRa Server

To list all configuration options, start loraserver with the --help flag. This will display:

```bash
GLOBAL OPTIONS:
   --net-id value                          network identifier (NetID, 3 bytes) encoded as HEX (e.g. 010203) [$NET_ID]
   --band value                            ism band configuration to use (options: AS_923, AU_915_928, CN_470_510, CN_779_787, EU_433, EU_863_870, KR_920_923, RU_864_869, US_902_928) [$BAND]
   --band-dwell-time-400ms                 band configuration takes 400ms dwell-time into account [$BAND_DWELL_TIME_400ms]
   --band-repeater-compatible              band configuration takes repeater encapsulation layer into account [$BAND_REPEATER_COMPATIBLE]
   --ca-cert value                         ca certificate used by the api server (optional) [$CA_CERT]
   --tls-cert value                        tls certificate used by the api server (optional) [$TLS_CERT]
   --tls-key value                         tls key used by the api server (optional) [$TLS_KEY]
   --bind value                            ip:port to bind the api server (default: "0.0.0.0:8000") [$BIND]
   --redis-url value                       redis url (e.g. redis://user:password@hostname:port/0) (default: "redis://localhost:6379") [$REDIS_URL]
   --postgres-dsn value                    postgresql dsn (e.g.: postgres://user:password@hostname/database?sslmode=disable) (default: "postgres://localhost/loraserver_ns?sslmode=disable") [$POSTGRES_DSN]
   --db-automigrate                        automatically apply database migrations [$DB_AUTOMIGRATE]
   --gw-mqtt-server value                  mqtt broker server used by the gateway backend (e.g. scheme://host:port where scheme is tcp, ssl or ws) (default: "tcp://localhost:1883") [$GW_MQTT_SERVER]
   --gw-mqtt-username value                mqtt username used by the gateway backend (optional) [$GW_MQTT_USERNAME]
   --gw-mqtt-password value                mqtt password used by the gateway backend (optional) [$GW_MQTT_PASSWORD]
   --as-server value                       hostname:port of the application-server api server (optional) (default: "127.0.0.1:8001") [$AS_SERVER]
   --as-ca-cert value                      ca certificate used by the application-server client (optional) [$AS_CA_CERT]
   --as-tls-cert value                     tls certificate used by the application-server client (optional) [$AS_TLS_CERT]
   --as-tls-key value                      tls key used by the application-server client (optional) [$AS_TLS_KEY]
   --nc-server value                       hostname:port of the network-controller api server (optional) [$NC_SERVER]
   --nc-ca-cert value                      ca certificate used by the network-controller client (optional) [$NC_CA_CERT]
   --nc-tls-cert value                     tls certificate used by the network-controller client (optional) [$NC_TLS_CERT]
   --nc-tls-key value                      tls key used by the network-controller client (optional) [$NC_TLS_KEY]
   --deduplication-delay value             time to wait for uplink de-duplication (default: 200ms) [$DEDUPLICATION_DELAY]
   --get-downlink-data-delay value         delay between uplink delivery to the app server and getting the downlink data from the app server (if any) (default: 100ms) [$GET_DOWNLINK_DATA_DELAY]
   --gw-stats-aggregation-intervals value  aggregation intervals to use for aggregating the gateway stats (valid options: second, minute, hour, day, week, month, quarter, year) (default: "minute", "hour", "day") [$GW_STATS_AGGREGATION_INTERVALS]
   --timezone value                        timezone to use when aggregating data (e.g. 'Europe/Amsterdam') (optional, by default the local timezone is used) [$TIMEZONE]
   --gw-create-on-stats                    create non-existing gateways on receiving of stats [$GW_CREATE_ON_STATS]
   --help, -h                              show help
   --version, -v                           print the version
```
### LoRa Gateway Bridge

To list all configuration options, start loraserver with the --help flag. This will display:

```bash
GLOBAL OPTIONS:
   --udp-bind value       ip:port to bind the UDP listener to (default: "0.0.0.0:1700") [$UDP_BIND]
   --mqtt-server value    mqtt server (e.g. scheme://host:port where scheme is tcp, ssl or ws) (default: "tcp://127.0.0.1:1883") [$MQTT_SERVER]
   --mqtt-username value  mqtt server username (optional) [$MQTT_USERNAME]
   --mqtt-password value  mqtt server password (optional) [$MQTT_PASSWORD]
   --skip-crc-check       skip the CRC status-check of received packets [$SKIP_CRC_CHECK]
   --log-level value      debug=5, info=4, warning=3, error=2, fatal=1, panic=0 (default: 4) [$LOG_LEVEL]
   --help, -h             show help
   --version, -v          print the version
```
