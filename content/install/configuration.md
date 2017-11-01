---
title: Configuration
menu:
    main:
        parent: install
        weight: 7
---

## Configuring the LoRa Server project applications

The LoRa Server project applications can be configured via the command line or via environment variables.  The following outlines the options for each of the applications.  More information on these parameters can be found on the
documentation pages for the applications accessible via the links above.

### LoRa App Server

To list all configuration options, start lora-app-server with the --help flag. This will display:

```text
   --postgres-dsn value             postgresql dsn (e.g.: postgres://user:password@hostname/database?sslmode=disable) (default: "postgres://localhost/loraserver?sslmode=disable") [$POSTGRES_DSN]
   --db-automigrate                 automatically apply database migrations [$DB_AUTOMIGRATE]
   --redis-url value                redis url (e.g. redis://user:password@hostname/0) (default: "redis://localhost:6379") [$REDIS_URL]
   --mqtt-server value              mqtt server (e.g. scheme://host:port where scheme is tcp, ssl or ws) (default: "tcp://localhost:1883") [$MQTT_SERVER]
   --mqtt-username value            mqtt server username (optional) [$MQTT_USERNAME]
   --mqtt-password value            mqtt server password (optional) [$MQTT_PASSWORD]
   --mqtt-ca-cert value             mqtt CA certificate file used by the gateway backend (optional) [$MQTT_CA_CERT]
   --as-public-server value         ip:port of the application-server api (used by LoRa Server to connect back to LoRa App Server) (default: "localhost:8001") [$AS_PUBLIC_SERVER]
   --as-public-id value             random uuid defining the id of the application-server installation (used by LoRa Server as routing-profile id) (default: "6d5db27e-4ce2-4b2b-b5d7-91f069397978") [$AS_PUBLIC_ID]
   --bind value                     ip:port to bind the api server (default: "0.0.0.0:8001") [$BIND]
   --ca-cert value                  ca certificate used by the api server (optional) [$CA_CERT]
   --tls-cert value                 tls certificate used by the api server (optional) [$TLS_CERT]
   --tls-key value                  tls key used by the api server (optional) [$TLS_KEY]
   --http-bind value                ip:port to bind the (user facing) http server to (web-interface and REST / gRPC api) (default: "0.0.0.0:8080") [$HTTP_BIND]
   --http-tls-cert value            http server TLS certificate [$HTTP_TLS_CERT]
   --http-tls-key value             http server TLS key [$HTTP_TLS_KEY]
   --jwt-secret value               JWT secret used for api authentication / authorization [$JWT_SECRET]
   --ns-ca-cert value               ca certificate used by the network-server client (optional) [$NS_CA_CERT]
   --ns-tls-cert value              tls certificate used by the network-server client (optional) [$NS_TLS_CERT]
   --ns-tls-key value               tls key used by the network-server client (optional) [$NS_TLS_KEY]
   --pw-hash-iterations value       the number of iterations used to generate the password hash (default: 100000) [$PW_HASH_ITERATIONS]
   --log-level value                debug=5, info=4, warning=3, error=2, fatal=1, panic=0 (default: 4) [$LOG_LEVEL]
   --disable-assign-existing-users  when set, existing users can't be re-assigned (to avoid exposure of all users to an organization admin) [$DISABLE_ASSIGN_EXISTING_USERS]
   --gw-ping                        enable sending gateway pings [$GW_PING]
   --gw-ping-interval value         the interval used for each gateway to send a ping (default: 24h0m0s) [$GW_PING_INTERVAL]
   --gw-ping-frequency value        the frequency used for transmitting the gateway ping (in Hz) (default: 0) [$GW_PING_FREQUENCY]
   --gw-ping-dr value               the data-rate to use for transmitting the gateway ping (default: 0) [$GW_PING_DR]
   --js-bind value                  ip:port to bind the join-server api interface to (default: "0.0.0.0:8003") [$JS_BIND]
   --js-ca-cert value               ca certificate used by the join-server api server (optional) [$JS_CA_CERT]
   --js-tls-cert value              tls certificate used by the join-server api server (optional) [$JS_TLS_CERT]
   --js-tls-key value               tls key used by the join-server api server (optional) [$JS_TLS_KEY]
   --help, -h                       show help
   --version, -v                    print the version
```

### LoRa Server

To list all configuration options, start loraserver with the --help flag. This will display:

```text
GLOBAL OPTIONS:
   --net-id value                          network identifier (NetID, 3 bytes) encoded as HEX (e.g. 010203) [$NET_ID]
   --band value                            ism band configuration to use (options: AS_923, AU_915_928, CN_470_510, CN_779_787, EU_433, EU_863_870, IN_865_867, KR_920_923, US_902_928) [$BAND]
   --band-dwell-time-400ms                 band configuration takes 400ms dwell-time into account [$BAND_DWELL_TIME_400ms]
   --band-repeater-compatible              band configuration takes repeater encapsulation layer into account [$BAND_REPEATER_COMPATIBLE]
   --ca-cert value                         ca certificate used by the api server (optional) [$CA_CERT]
   --tls-cert value                        tls certificate used by the api server (optional) [$TLS_CERT]
   --tls-key value                         tls key used by the api server (optional) [$TLS_KEY]
   --bind value                            ip:port to bind the api server (default: "0.0.0.0:8000") [$BIND]
   --gw-server-ca-cert value               ca certificate used by the gateway api server (optional) [$GW_SERVER_CA_CERT]
   --gw-server-tls-cert value              tls certificate used by the gateway api server (optional) [$GW_SERVER_TLS_CERT]
   --gw-server-tls-key value               tls key used by the gateway api server (optional) [$GW_SERVER_TLS_KEY]
   --gw-server-jwt-secret value            JWT secret used by the gateway api server for gateway authentication / authorization [$GW_SERVER_JWT_SECRET]
   --gw-server-bind value                  ip:port to bind the gateway api server (default: "0.0.0.0:8002") [$GW_SERVER_BIND]
   --redis-url value                       redis url (e.g. redis://user:password@hostname:port/0) (default: "redis://localhost:6379") [$REDIS_URL]
   --postgres-dsn value                    postgresql dsn (e.g.: postgres://user:password@hostname/database?sslmode=disable) (default: "postgres://localhost/loraserver_ns?sslmode=disable") [$POSTGRES_DSN]
   --db-automigrate                        automatically apply database migrations [$DB_AUTOMIGRATE]
   --gw-mqtt-server value                  mqtt broker server used by the gateway backend (e.g. scheme://host:port where scheme is tcp, ssl or ws) (default: "tcp://localhost:1883") [$GW_MQTT_SERVER]
   --gw-mqtt-username value                mqtt username used by the gateway backend (optional) [$GW_MQTT_USERNAME]
   --gw-mqtt-password value                mqtt password used by the gateway backend (optional) [$GW_MQTT_PASSWORD]
   --gw-mqtt-ca-cert value                 mqtt CA certificate file used by the gateway backend (optional) [$GW_MQTT_CA_CERT]
   --as-ca-cert value                      ca certificate used by the application-server client (optional) [$AS_CA_CERT]
   --as-tls-cert value                     tls certificate used by the application-server client (optional) [$AS_TLS_CERT]
   --as-tls-key value                      tls key used by the application-server client (optional) [$AS_TLS_KEY]
   --nc-server value                       hostname:port of the network-controller api server (optional) [$NC_SERVER]
   --nc-ca-cert value                      ca certificate used by the network-controller client (optional) [$NC_CA_CERT]
   --nc-tls-cert value                     tls certificate used by the network-controller client (optional) [$NC_TLS_CERT]
   --nc-tls-key value                      tls key used by the network-controller client (optional) [$NC_TLS_KEY]
   --deduplication-delay value             time to wait for uplink de-duplication (default: 200ms) [$DEDUPLICATION_DELAY]
   --get-downlink-data-delay value         delay between uplink delivery to the app server and getting the downlink data from the app server (if any) (default: 100ms) [$GET_DOWNLINK_DATA_DELAY]
   --gw-stats-aggregation-intervals value  aggregation intervals to use for aggregating the gateway stats (valid options: second, minute, hour, day, week, month, quarter, year) (default: "minute,hour,day") [$GW_STATS_AGGREGATION_INTERVALS]
   --timezone value                        timezone to use when aggregating data (e.g. 'Europe/Amsterdam') (optional, by default the db timezone is used) [$TIMEZONE]
   --gw-create-on-stats                    create non-existing gateways on receiving of stats [$GW_CREATE_ON_STATS]
   --extra-frequencies value               extra frequencies to use for ISM bands that implement the CFList [$EXTRA_FREQUENCIES]
   --enable-uplink-channels value          enable only a given sub-set of channels (e.g. '0-7,8-15') [$ENABLE_UPLINK_CHANNELS]
   --node-session-ttl value                the ttl after which a node-session expires after no activity (default: 744h0m0s) [$NODE_SESSION_TTL]
   --log-node-frames                       log uplink and downlink frames to the database [$LOG_NODE_FRAMES]
   --log-level value                       debug=5, info=4, warning=3, error=2, fatal=1, panic=0 (default: 4) [$LOG_LEVEL]
   --js-server value                       hostname:port of the default join-server (default: "http://localhost:8003") [$JS_SERVER]
   --js-ca-cert value                      ca certificate used by the default join-server client (optional) [$JS_CA_CERT]
   --js-tls-cert value                     tls certificate used by the default join-server client (optional) [$JS_TLS_CERT]
   --js-tls-key value                      tls key used by the default join-server client (optional) [$JS_TLS_KEY]
   --installation-margin value             installation margin (dB) used by the ADR engine (default: 10) [$INSTALLATION_MARGIN]
   --rx1-delay value                       class a rx1 delay (default: 1) [$RX1_DELAY]
   --rx1-dr-offset value                   rx1 data-rate offset (valid options documented in the LoRaWAN Regional Parameters specification) (default: 0) [$RX1_DR_OFFSET]
   --rx2-dr value                          rx2 data-rate (when set to -1, the default rx2 data-rate will be used) (default: -1) [$RX2_DR]
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
