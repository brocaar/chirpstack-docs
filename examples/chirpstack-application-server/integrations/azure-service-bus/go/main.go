package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"log"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/brocaar/chirpstack-api/go/v3/as/integration"
)

type handler struct {
	json bool

	ns    *servicebus.Namespace
	queue *servicebus.Queue
}

func (h *handler) receive() error {
	for {
		if err := h.queue.Receive(context.TODO(), servicebus.HandlerFunc(h.receiveFunc)); err != nil {
			return err
		}
	}
}

func (h *handler) receiveFunc(ctx context.Context, msg *servicebus.Message) error {
	ev, ok := msg.UserProperties["Event"]
	if !ok {
		log.Println("event attribute is missing")
	}

	event, ok := ev.(string)
	if !ok {
		log.Println("event must be of type string")
	}

	var err error

	switch event {
	case "up":
		err = h.up(msg.Data)
	case "join":
		err = h.join(msg.Data)
	default:
		log.Printf("handler for event %s is not implemented", event)
	}

	if err != nil {
		log.Printf("handling event '%s' returned error: %s", event, err)
	}

	return msg.Complete(ctx)
}

func (h *handler) up(b []byte) error {
	var up integration.UplinkEvent
	if err := h.unmarshal(b, &up); err != nil {
		return err
	}
	log.Printf("Uplink received from %s with payload: %s", hex.EncodeToString(up.DevEui), hex.EncodeToString(up.Data))
	return nil
}

func (h *handler) join(b []byte) error {
	var join integration.JoinEvent
	if err := h.unmarshal(b, &join); err != nil {
		return err
	}
	log.Printf("Device %s joined with DevAddr %s", hex.EncodeToString(join.DevEui), hex.EncodeToString(join.DevAddr))
	return nil
}

func (h *handler) unmarshal(b []byte, v proto.Message) error {
	if h.json {
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true, // we don't want to fail on unknown fields
		}
		return unmarshaler.Unmarshal(bytes.NewReader(b), v)
	}
	return proto.Unmarshal(b, v)
}

func newHandler(json bool, connStr string, queueName string) (*handler, error) {
	ns, err := servicebus.NewNamespace(
		servicebus.NamespaceWithConnectionString(connStr),
	)
	if err != nil {
		panic(err)
	}

	queue, err := ns.NewQueue(queueName)
	if err != nil {
		panic(err)
	}

	return &handler{
		json:  json,
		ns:    ns,
		queue: queue,
	}, nil

}

func main() {
	h, err := newHandler(
		// set true when using JSON encoding
		false,

		// service-bus connection string
		"Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=example-policy;SharedAccessKey=...",

		// queue name
		"events",
	)
	if err != nil {
		panic(err)
	}
	panic(h.receive())
}
