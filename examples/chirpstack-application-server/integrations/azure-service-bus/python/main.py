from azure.servicebus import ServiceBusClient
from chirpstack_api.as_pb import integration
from google.protobuf.json_format import Parse


class Handler:
    def __init__(self, json, connection_string, queue_name):
        self.json = json
        self.connection_string = connection_string
        self.queue_name = queue_name

    def receive(self):
        client = ServiceBusClient.from_connection_string(self.connection_string)
        queue_client = client.get_queue(self.queue_name)
        messages = queue_client.get_receiver()

        for message in messages:
            message.complete()
            body = b''.join(message.body)

            event = message.user_properties[b'Event']

            if event == b'up':
                self.up(body)
            elif event == b'join':
                self.join(body)
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

    # Service-Bus connection string
    'Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=example-policy;SharedAccessKey=...',

    # Service-Bus queue name
    'events',
)
h.receive()
