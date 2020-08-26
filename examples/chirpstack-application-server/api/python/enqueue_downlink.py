import os
import sys

import grpc
from chirpstack_api.as_pb.external import api

# Configuration.

# This must point to the API interface.
server = "localhost:8080"

# The DevEUI for which you want to enqueue the downlink.
dev_eui = bytes([0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01])

# The API token (retrieved using the web-interface).
api_token = "..."

if __name__ == "__main__":
  # Connect without using TLS.
  channel = grpc.insecure_channel(server)

  # Device-queue API client.
  client = api.DeviceQueueServiceStub(channel)

  # Define the API key meta-data.
  auth_token = [("authorization", "Bearer %s" % api_token)]

  # Construct request.
  req = api.EnqueueDeviceQueueItemRequest()
  req.device_queue_item.confirmed = False
  req.device_queue_item.data = bytes([0x01, 0x02, 0x03])
  req.device_queue_item.dev_eui = dev_eui.hex()
  req.device_queue_item.f_port = 10

  resp = client.Enqueue(req, metadata=auth_token)

  # Print the downlink frame-counter value.
  print(resp.f_cnt)
