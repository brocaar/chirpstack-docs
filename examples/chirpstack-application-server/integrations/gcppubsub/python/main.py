from chirpstack_api.as_pb import integration
from google.cloud import pubsub_v1
from google.protobuf.json_format import Parse


class Handler:
    def __init__(self, json, project_id, subscription_name):
        self.json = json
        self.project_id = project_id
        self.subscription_name = subscription_name

    def receive(self):
        subscriber = pubsub_v1.SubscriberClient()
        sub_path = subscriber.subscription_path(self.project_id, self.subscription_name)

        while True:
            resp = subscriber.pull(sub_path, max_messages=10)
            for msg in resp.received_messages:
                event = msg.message.attributes['event']
                subscriber.acknowledge(sub_path, [msg.ack_id])

                if event == 'up':
                    self.up(msg.message.data)
                elif event == 'join':
                    self.join(msg.message.data)
                else:
                    print('handler for event %s is not implemented' % event)

    def up(self, body):
        up = self.unmarshal(body, integration.UplinkEvent())
        print('Uplink received from: %s with payload: %s' % (up.dev_eui.hex(), up.data.hex()))

    def join(self, body):
        join = self.unmarshal(body, integration.JoinEvent())
        print('Device: %s joined with DevAddr: %s' % (join.dev_eui.hex(), join.dev_addr.hex()))

    def unmarshal(self, body, pl):
        if self.json:
            return Parse(body, pl)
        
        pl.ParseFromString(body)
        return pl


h = Handler(
    # True -  JSON marshaler
    # False - Protobuf marshaler (binary)
    False,

    # GCP Project ID
    "project-id",

    # GCP Pub/Sub Subsciption name
    "subscription-name",
)
h.receive()
