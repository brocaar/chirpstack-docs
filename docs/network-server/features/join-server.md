---
description: Configure ChirpStack with an external Join Server.
---

# Join Server

ChirpStack Network Server supports the activation of devices through an
external Join Server, using the API as specified in the LoRaWAN Backend
Interfaces specification. In this case the session-keys will not be derived by
the ChirpStack Application Server and it doesn't need to know the `AppKey` (and
`NwkKey` in case of a LoRaWAN 1.1 device).

Routing a join-request to a Join Server is based on the `JoinEUI` which is part
of the LoRaWAN join-request. If ChirpStack Network Server finds a server matching
the `JoinEUI`, then it will forward the join-request to this server. In any other
case, the ChirpStack Application Server will handle the join-request.

## Configuration

### Semtech LoRa Cloud Join Server

To configure the Semtech LoRa Cloud Join Server with ChirpStack obtain the TLS
certificates that you must use from the LoRa Cloud web-interface.

* Under **Your network servers** click **Add Server**
* Then click download credentials

Within the obtained credentials archive you will find three files that you need
to configure within the ChirpStack Network Server configuration:

* `.trust`: must be configured as `ca_cert`
* `.crt`: must be configured as `tls_cert`
* `.key`: must be configured as `tls_key`

**Note:** While the web-interface indicates that `https://js.loracloud.com:7009`
is the **LoRa Cloud Join Server**, this is not the endpoint that must be
configured as `server` within the ChirpStack Network Server configuration. The
correct endpoint is `https://js.loracloud.com:7009/api/v1/rens/rens-XYZ/lbi`.
You have to replace `rens-XYZ` with the actual Rens ID, you will find this value
in the LoRa Cloud web-interface.


Example configuration:

```toml
[[join_server.servers]]
join_eui="0016c001fffe0001"
server="https://js.loracloud.com:7009/api/v1/rens/rens-123/lbi"
ca_cert="/etc/chirpstack-network-server/certs/acct.trust"
tls_cert="/etc/chirpstack-network-server/certs/acct.crt"
tls_key="/etc/chirpstack-network-server/certs/acct.key"
```
