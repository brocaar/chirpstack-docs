---
description: Configuration of a NGINX proxy.
---

# NGINX proxy

When you are using [NGINX](http://nginx.org) in front of ChirpStack Application Server, you must
add additional configuration for the WebSocket endpoints. Without this
additional configuration, you are going to find the following messages in the
ChirpStack Application Server logs:

```text
level=info msg=“finished streaming call with code Unauthenticated” error=“rpc error: code = Unauthenticated desc = authentication failed: get token from context error: no authorization-data in metadata” grpc.code=Unauthenticated grpc.method=StreamFrameLogs grpc.service=api.Gateway grpc.start_time=“2018-04-04T09:48:20+07:00” grpc.time_ms=0.07 peer.address=“127.0.0.1:60048” span.kind=server system=grpc
```

## Config example

The following is a configuration example. Adapt it to your own environment.

```nginx
server {
	listen 443 ssl;
	server_name localhost;

	ssl_certificate /etc/chirpstack-application-server/certs/http.pem;
	ssl_certificate_key /etc/chirpstack-application-server/certs/http-key.pem;

	# WebSocket configuration
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
