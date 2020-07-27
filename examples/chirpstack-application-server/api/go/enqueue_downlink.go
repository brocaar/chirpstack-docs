package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
	"github.com/brocaar/lorawan"
)

// configuration
var (
	// This must point to the API interface
	server = "localhost:8080"

	// The DevEUI for which we want to enqueue the downlink
	devEUI = lorawan.EUI64{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}

	// The API token (retrieved using the web-interface)
	apiToken = "..."
)

type APIToken string

func (a APIToken) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", a),
	}, nil
}

func (a APIToken) RequireTransportSecurity() bool {
	return false
}

func main() {
	// define gRPC dial options
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(APIToken(apiToken)),
		grpc.WithInsecure(), // remove this when using TLS
	}

	// connect to the gRPC server
	conn, err := grpc.Dial(server, dialOpts...)
	if err != nil {
		panic(err)
	}

	// define the DeviceQueueService client
	queueClient := api.NewDeviceQueueServiceClient(conn)

	// make an Enqueue api call
	resp, err := queueClient.Enqueue(context.Background(), &api.EnqueueDeviceQueueItemRequest{
		DeviceQueueItem: &api.DeviceQueueItem{
			DevEui:    devEUI.String(),
			FPort:     10,
			Confirmed: false,
			Data:      []byte{0x01, 0x02, 0x03},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("The downlink has been enqueued with FCnt: %d\n", resp.FCnt)
}
