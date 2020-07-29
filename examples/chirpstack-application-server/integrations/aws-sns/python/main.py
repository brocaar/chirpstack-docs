import boto3
from chirpstack_api.as_pb import integration
from google.protobuf.json_format import Parse


class Handler:
    def __init__(self, json, queue_url):
        self.json = json
        self.queue_url = queue_url

    def receive(self):
        sqs = boto3.client('sqs')

        while True:
            resp = sqs.receive_message(
                QueueUrl=self.queue_url,
                MessageAttributeNames=[
                    'All',
                ],
                MaxNumberOfMessages=1,
                WaitTimeSeconds=10,
            )

            if not 'Messages' in resp:
                continue

            msg = resp['Messages'][0]
            receipt_handle = msg['ReceiptHandle']

            sqs.delete_message(
                QueueUrl=self.queue_url,
                ReceiptHandle=receipt_handle,
            )

            event = msg['MessageAttributes']['event']['StringValue']

            if event == "up":
                self.up(msg['Body'])
            elif event == "join":
                self.join(msg['Body'])
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

    # SQS queue url
    'https://sqs....',
)
h.receive()
