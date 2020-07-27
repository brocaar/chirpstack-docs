package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/brocaar/chirpstack-api/go/v3/as/integration"
)

type handler struct {
	json bool

	sqs      *sqs.SQS
	queueURL string
}

func (h *handler) receive() error {
	for {
		result, err := h.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &h.queueURL,
			MaxNumberOfMessages: aws.Int64(1),
		})
		if err != nil {
			return err
		}

		for _, msg := range result.Messages {
			_, err := h.sqs.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &h.queueURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("delete message error: %s", err)
			}

			event, ok := msg.MessageAttributes["event"]
			if !ok || event.StringValue == nil {
				log.Printf("event attribute is missing")
				continue
			}

			switch *event.StringValue {
			case "up":
				err = h.up(*msg.Body)
			case "join":
				err = h.join(*msg.Body)
			default:
				log.Printf("handler for event %s is not implemented", *event.StringValue)
				err = nil
			}

			if err != nil {
				log.Printf("handling event '%s' returned error: %s", *event.StringValue, err)
			}

		}
	}

	return nil
}

func (h *handler) up(body string) error {
	var up integration.UplinkEvent
	if err := h.unmarshal(body, &up); err != nil {
		return err
	}
	log.Printf("Uplink received from %s with payload: %s", hex.EncodeToString(up.DevEui), hex.EncodeToString(up.Data))
	return nil
}

func (h *handler) join(body string) error {
	var join integration.JoinEvent
	if err := h.unmarshal(body, &join); err != nil {
		return err
	}
	log.Printf("Device %s joined with DevAddr %s", hex.EncodeToString(join.DevEui), hex.EncodeToString(join.DevAddr))
	return nil
}

func (h *handler) unmarshal(body string, v proto.Message) error {
	if h.json {
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true, // we don't want to fail on unknown fields
		}
		return unmarshaler.Unmarshal(bytes.NewReader([]byte(body)), v)
	}

	b, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, v)
}

func newHandler(json bool, accessKeyID, secretAccessKey, region, queueURL string) (*handler, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return &handler{
		json:     json,
		sqs:      sqs.New(sess),
		queueURL: queueURL,
	}, nil
}

func main() {
	h, err := newHandler(
		// set true when using JSON encoding
		false,

		// AWS AccessKeyID
		"...",

		// AWS SecretAccessKey
		"...",

		// AWS region
		"eu-west-1",

		// SQS queue url
		"https://sqs...",
	)
	if err != nil {
		panic(err)
	}

	panic(h.receive())
}
