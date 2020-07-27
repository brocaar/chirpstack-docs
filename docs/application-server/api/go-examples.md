---
description: Integrate with the ChirpStack Application Server API using Go.
---

# Go examples

* [Go gRPC reference](https://pkg.go.dev/google.golang.org/grpc)
* [ChirpStack Go SDK reference](https://pkg.go.dev/github.com/brocaar/chirpstack-api/go/v3/as/external/api)

## Enqueue downlink

The example below demonstrates:

* Configuration of gRPC _dial options_ including API token
* Connect to a gRPC API
* Define service client (in this case for the `DeviceQueueService`)
* Perform an API call for a service (in this case `Enqueue`)

=== "main.go"

	```go
	--8<--- "examples/chirpstack-application-server/api/go/enqueue_downlink.go"
	```
